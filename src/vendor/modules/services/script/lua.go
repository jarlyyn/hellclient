package script

import (
	"path"

	"github.com/yuin/gopher-lua"
)

type Lua struct {
	path string
	Lua  *lua.LState
}

func NewLua() *Lua {
	return &Lua{}
}
func (l *Lua) Load(path string) error {
	l.path = path
	l.Lua = lua.NewState()
	l.Lua.SetGlobal("GetInfo", l.Lua.NewFunction(l.APIGetInfo))
	return l.Lua.DoFile(path)
}

func (l *Lua) APIGetInfo(L *lua.LState) int {
	var result lua.LValue
	i := L.ToInt(1)
	if i == 35 {
		result = lua.LString(path.Dir(l.path))
	}
	L.Push(result)
	return 1
}
