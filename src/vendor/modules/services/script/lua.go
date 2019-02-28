package script

import (
	"fmt"
	"modules/services/mapper"
	"path"
	"regexp"

	"github.com/yuin/gopher-lua"
)

type Lua struct {
	script *Script
	path   string
	Lua    *lua.LState
}

func NewLua() *Lua {
	return &Lua{}
}
func (l *Lua) Script() *Script {
	return l.script
}
func (l *Lua) SetScript(script *Script) {
	l.script = script
}
func (l *Lua) Load(path string) error {
	l.path = path
	l.Lua = lua.NewState()
	l.Lua.SetGlobal("GetInfo", l.Lua.NewFunction(l.APIGetInfo))
	l.Lua.SetGlobal("GetVariable", l.Lua.NewFunction(l.APIGetVariable))
	l.Lua.SetGlobal("GetTriggerList", l.Lua.NewFunction(l.APIGetTriggerList))
	l.Lua.SetGlobal("SetTriggerOption", l.Lua.NewFunction(l.APISetTriggerOption))

	l.Lua.SetGlobal("Mapper", l.APIMapper())
	l.Lua.SetGlobal("rex", rex(l.Lua))
	return l.Lua.DoFile(path)
}
func (l *Lua) APIMapper() *lua.LTable {
	m := l.script.Mapper
	L := l.Lua
	t := L.NewTable()
	L.SetFuncs(t, map[string]lua.LGFunction{
		"openroomsall": func(L *lua.LState) int {
			filename := L.ToString(-1)
			err := mapper.CommonRoomAllIniLoader.Load(m, filename)
			if err != nil {
				fmt.Println(err)
				L.Push(lua.LBool(false))
				return 1
			}
			L.Push(lua.LBool(true))
			return 1
		},
		"newarea": func(L *lua.LState) int {
			size := L.ToInt(-1)
			result := m.NewArea(size)
			t := L.NewTable()
			for _, v := range result {
				t.Append(lua.LString(v))
			}
			L.Push(t)
			return 1
		},
	})
	return t
}
func (l *Lua) APIGetInfo(L *lua.LState) int {
	var result lua.LValue
	i := L.ToInt(-1)
	if i == 35 {
		result = lua.LString(path.Dir(l.path)) + "/"
	}
	L.Push(result)
	return 1
}
func (l *Lua) APISetTriggerOption(L *lua.LState) int {
	id := L.ToString(-1)
	trigger := l.script.Triggers.VMTriggers[id]
	if trigger == nil {
		return 0
	}
	switch L.ToString(-1) {
	case "":
	default:
		return 0
	}

	l.script.Triggers.Add(trigger)
	return 0
}
func (l *Lua) APIGetVariable(L *lua.LState) int {
	name := L.ToString(-1)
	value := l.script.Variables[name]
	L.Push(lua.LString(value))
	return 1
}
func (l *Lua) APIGetTriggerList(L *lua.LState) int {
	t := L.NewTable()
	triggers := l.script.Triggers.VMTriggers
	for _, v := range triggers {
		t.Append(lua.LString(v.Name))
	}
	L.Push(t)
	return 1
}

func rex(L *lua.LState) *lua.LTable {
	t := L.NewTable()
	L.SetFuncs(t, map[string]lua.LGFunction{
		"new": newRegexp,
	})
	return t
}
func newRegexp(L *lua.LState) int {
	pattern := L.ToString(-1)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return 0
	}
	r := rexInstance{
		regexp: re,
	}
	t := L.NewTable()
	L.SetFuncs(t, map[string]lua.LGFunction{
		"find": r.Find,
	})
	L.Push(t)
	return 1
}

type rexInstance struct {
	regexp *regexp.Regexp
}

func (r *rexInstance) Find(L *lua.LState) int {
	return 0
}

func (r *rexInstance) Gmatch(L *lua.LState) int {
	return 0
}
