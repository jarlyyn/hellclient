package luaengine

import (
	"context"
	"hellclient/modules/world"
	"hellclient/modules/world/bus"

	lua "github.com/yuin/gopher-lua"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	"github.com/herb-go/plugins/addons/httpaddon/httplua"
)

var asyncexecute = func(bus *bus.Bus, req *httplua.Request) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		callback := ""
		if L.ToBool(1) {
			callback = L.ToString(1)
		}
		req.Request.AsyncExecute(func(err error) {
			if callback != "" {
				cb := world.NewCallback()
				cb.Name = "httpexecute"
				cb.ID = req.Request.ID
				cb.Script = callback
				if err != nil {
					cb.Code = -1
					cb.Data = err.Error()
				} else {
					cb.Code = 0
					cb.Data = req.Request.GetURL()
				}

				bus.DoSendCallbackToScript(cb)
			}
		})
		return 0
	}
}

func NewHTTPModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("http",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			addon := httplua.Create(plugin)
			addon.Builder = func(L *lua.LState, req *httplua.Request) *lua.LTable {
				t := httplua.DefaultBuilder(L, req)
				t.RawSetString("Execute", lua.LNil)
				t.RawSetString("AsyncExecute", L.NewFunction(asyncexecute(b, req)))
				return t
			}
			l.SetGlobal("HTTP", addon.Convert(l))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
