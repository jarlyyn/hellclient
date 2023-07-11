package luaengine

import (
	"context"
	"hellclient/modules/world/bus"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	"github.com/herb-go/plugins/addons/binaryaddon/binarylua"
)

func NewBinaryModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("http",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			addon := binarylua.Create(plugin)
			l.SetGlobal("Binary", addon.Convert(l))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
