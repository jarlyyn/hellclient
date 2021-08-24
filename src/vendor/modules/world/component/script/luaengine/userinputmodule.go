package luaengine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/userinput"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	lua "github.com/yuin/gopher-lua"
)

type List struct {
	List *userinput.List
	bus  *bus.Bus
}

func (l *List) Send(L *lua.LState) int {
	ui := l.List.Send(l.bus, L.ToString(1))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (l *List) Append(L *lua.LState) int {
	l.List.Append(L.ToString(1), L.ToString(2))
	return 0
}
func (l *List) SetValues(L *lua.LState) int {
	v := []string{}
	t := L.ToTable(1)
	m := t.MaxN()
	for i := 1; i <= m; i++ {
		v = append(v, t.RawGetInt(i).String())
	}
	l.List.SetValues(v)
	return 0
}
func (l *List) SetMutli(L *lua.LState) int {
	l.List.SetMutli(L.ToBool(1))
	return 0
}

func (l *List) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("append", L.NewFunction(l.Append))
	t.RawSetString("send", L.NewFunction(l.Send))
	t.RawSetString("setvalues", L.NewFunction(l.SetValues))
	t.RawSetString("setmutli", L.NewFunction(l.SetMutli))

	return t
}

type Userinput struct {
	bus *bus.Bus
}

func (u *Userinput) NewList(L *lua.LState) int {
	list := &List{
		List: userinput.CreateList(L.ToString(1), L.ToString(2), L.ToBool(3)),
		bus:  u.bus,
	}
	L.Push(list.Convert(L))
	return 1
}
func (u *Userinput) Prompt(L *lua.LState) int {
	ui := userinput.SendPrompt(u.bus, L.ToString(1), L.ToString(2), L.ToString(3), L.ToString(4))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Confirm(L *lua.LState) int {
	ui := userinput.SendConfirm(u.bus, L.ToString(1), L.ToString(2), L.ToString(3))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Alert(L *lua.LState) int {
	ui := userinput.SendAlert(u.bus, L.ToString(1), L.ToString(2), L.ToString(3))
	L.Push(lua.LString(ui.ID))
	return 1
}

func (u *Userinput) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("prompt", L.NewFunction(u.Prompt))
	t.RawSetString("confirm", L.NewFunction(u.Confirm))
	t.RawSetString("alert", L.NewFunction(u.Alert))
	t.RawSetString("newlist", L.NewFunction(u.NewList))
	return t
}
func NewUserinputModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("userinput",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			u := &Userinput{bus: b}
			l.SetGlobal("Userinput", u.Convert(l))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
