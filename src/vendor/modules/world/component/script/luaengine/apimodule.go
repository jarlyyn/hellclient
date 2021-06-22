package luaengine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/api"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	lua "github.com/yuin/gopher-lua"
)

func createApi(b *bus.Bus) *luaapi {
	return &luaapi{
		API: &api.API{
			Bus: b,
		},
	}
}

type luaapi struct {
	API *api.API
}

func (a *luaapi) InstallAPIs(l *lua.LState) {
	l.SetGlobal("Note", l.NewFunction(a.Note()))
	l.SetGlobal("SendImmediate", l.NewFunction(a.SendImmediate()))
	l.SetGlobal("Send", l.NewFunction(a.Send()))

}
func (a *luaapi) Note() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		info := L.ToString(1)
		a.API.Note(info)
		return 0
	}
}
func (a *luaapi) SendImmediate() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		info := L.ToString(1)
		a.API.SendImmediate(info)
		return 0
	}
}
func (a *luaapi) Send() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		info := L.ToString(1)
		a.API.Send(info)
		return 0
	}
}
func NewAPIModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("worldapi",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			createApi(b).InstallAPIs(l)
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
