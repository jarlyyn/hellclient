package luaengine

import (
	"context"
	"modules/world"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	"github.com/jarlyyn/golang-pkg-pcre/src/pkg/pcre"
	lua "github.com/yuin/gopher-lua"
)

func RexFlags(L *lua.LState) int {
	L.Push(L.NewTable())
	return 1
}

type Rex struct {
	Regexp pcre.Regexp
}
type matchResult struct {
	start int
	end   int
	*world.MatchResult
}

func (r *Rex) getToMatch(text string, pos int) string {
	if pos < 0 {
		pos = len(text) + pos
	}
	if pos > len(text) {
		pos = len(text)
	}

	return text[pos:]
}
func (r *Rex) exec(text string, pos int, flag int) []int {

	m := r.Regexp.MatcherString(r.getToMatch(text, pos), flag)
	if !m.Matches() {
		return nil
	}
	groups := m.Groups()
	result := m.Ovector()
	return result[:(groups+1)*2]
}
func (r *Rex) match(text string, pos int, flag int) *matchResult {
	text = r.getToMatch(text, pos)
	if text == "" {
		return nil
	}
	m := r.Regexp.MatcherString(text, flag)
	if !m.Matches() {
		return nil
	}
	matched := m.Groups()
	o := m.Ovector()
	result := &matchResult{
		start:       o[0] + 1,
		end:         o[1],
		MatchResult: world.NewMatchResult(),
	}

	for i := 0; i <= matched; i++ {
		v := m.GroupString(i)
		result.List = append(result.List, v)
	}
	for k, v := range m.Names() {
		if v != "" {
			result.Named[v] = result.List[k]
		}
	}
	return result
}
func (r *Rex) Exec(L *lua.LState) int {
	_ = L.ToString(1) //self
	text := L.ToString(2)
	pos := int(L.ToNumber(3))
	flag := int(L.ToNumber(4))
	result := r.exec(text, pos, flag)
	if result == nil {
		return 0
	}
	L.Push(lua.LNumber(result[0] + 1))
	L.Push(lua.LNumber(result[1]))
	t := L.NewTable()
	for i := 2; i < len(result); i++ {
		t.RawSetInt(i-1, lua.LNumber(result[i]+((i-1)%2)))
	}
	L.Push(t)
	return 3
}

func (r *Rex) Match(L *lua.LState) int {
	_ = L.ToString(1) //self
	text := L.ToString(2)
	pos := int(L.ToNumber(3))
	flag := int(L.ToNumber(4))
	result := r.match(text, pos, flag)
	if result == nil {
		return 0
	}
	L.Push(lua.LNumber(result.start))
	L.Push(lua.LNumber(result.end))
	t := L.NewTable()
	for i := 1; i < len(result.List); i++ {
		v := result.List[i]
		if v == "" {
			t.RawSetInt(i, lua.LBool(false))
		} else {
			t.RawSetInt(i, lua.LString(v))
		}
	}
	for k, v := range result.Named {
		if v == "" {
			t.RawSetString(k, lua.LBool(false))
		} else {
			t.RawSetString(k, lua.LString(v))
		}
	}
	L.Push(t)
	return 3
}
func (r *Rex) Gmatch(L *lua.LState) int {
	_ = L.ToString(1) //self
	text := L.ToString(2)
	fun := L.ToFunction(3)
	count := int(L.ToNumber(4))
	flag := int(L.ToNumber(5))
	var matched int
	for {
		if count > 0 && matched >= count {
			break
		}
		result := r.match(text, 0, flag)
		if result == nil {
			break
		}
		matched++
		t := L.NewTable()
		for i := 1; i < len(result.List); i++ {
			v := result.List[i]
			t.RawSetInt(i, lua.LString(v))
		}
		for k, v := range result.Named {
			t.RawSetString(k, lua.LString(v))
		}
		L.CallByParam(lua.P{
			Fn:   fun,
			NRet: 1,
		}, lua.LString(result.List[0]), t)
		ret := L.Get(-1)
		L.Pop(1)
		if lua.LVAsBool(ret) {
			break
		}
		if result.end >= len(text) {
			break
		}
		text = text[result.end:]
	}

	L.Push(lua.LNumber(matched))
	return 1
}

func CreateRex(L *lua.LState) int {
	text := L.ToString(1)
	flag := int(L.ToNumber(1))
	re := pcre.MustCompile(text, flag)
	rex := &Rex{re}
	result := L.NewTable()
	result.RawSetString("exec", L.NewFunction(rex.Exec))
	result.RawSetString("match", L.NewFunction(rex.Match))
	result.RawSetString("gmatch", L.NewFunction(rex.Gmatch))

	L.Push(result)
	return 1
}

var ModuleRex = herbplugin.CreateModule("rex",
	func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
		luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
		l := luapluing.LState
		rex := l.NewTable()
		rex.RawSetString("flags", l.NewFunction(RexFlags))
		rex.RawSetString("new", l.NewFunction(CreateRex))
		l.SetGlobal("rex", rex)
		next(ctx, plugin)
	},
	nil,
	nil,
)
