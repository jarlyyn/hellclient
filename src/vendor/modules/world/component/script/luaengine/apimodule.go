package luaengine

import (
	"context"
	"modules/world/bus"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	lua "github.com/yuin/gopher-lua"
)

func createApi(b *bus.Bus) *api {
	return &api{
		Bus: b,
	}
}

type api struct {
	Bus *bus.Bus
}

func (a *api) InstallAPIs(l *lua.LState) {
	l.SetGlobal("Note", l.NewFunction(a.APINote()))
	l.SetGlobal("SendImmediate", l.NewFunction(a.APISendImmediate()))
	l.SetGlobal("Send", l.NewFunction(a.APISend()))

}
func (a *api) APINote() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		info := L.ToString(1)
		a.Bus.DoPrint(info)
		return 0
	}
}
func (a *api) APISendImmediate() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		info := L.ToString(1)
		a.Bus.DoSend([]byte(info))
		return 0
	}
}
func (a *api) APISend() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		info := L.ToString(1)
		a.Bus.DoSendToQueue([]byte(info))
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
