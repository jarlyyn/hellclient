package jsengine

import (
	"context"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
	"github.com/herb-go/plugins/addons/httpaddon/httpjs"
)

var ModuleHTTP = herbplugin.CreateModule("http",
	func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
		jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
		r := jsp.Runtime
		r.Set("HTTP", httpjs.Create(jsp))
		next(ctx, plugin)
	},
	nil,
	nil,
)
