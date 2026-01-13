package v8engine

import (
	"context"
	"modules/world/bus"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/v8local/plugins/binaryaddon/binaryv8"
	"github.com/herb-go/v8local/v8plugin"
)

func NewBinaryModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("binary",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime.NewLocal()
			addon := binaryv8.Create(plugin)
			global := r.Global()
			global.Set("Binary", addon.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
