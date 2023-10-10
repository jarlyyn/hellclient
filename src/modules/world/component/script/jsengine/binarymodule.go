package jsengine

import (
	"context"
	"modules/world/bus"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
	"github.com/herb-go/plugins/addons/binaryaddon/binaryjs"
)

func NewBinaryModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("binary",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			addon := binaryjs.Create(plugin)
			r.Set("Binary", addon.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
