package jsengine

import (
	"context"
	"modules/world"
	"modules/world/bus"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
	"github.com/herb-go/plugins/addons/httpaddon/httpjs"
)

var asyncexecute = func(bus *bus.Bus, req *httpjs.Request) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		callback := ""
		if call.Argument(0).ToBoolean() {
			callback = call.Argument(0).String()
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
		return nil
	}
}

func NewHTTPModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("http",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			addon := httpjs.Create(plugin)
			addon.Builder = func(r *goja.Runtime, req *httpjs.Request) *goja.Object {
				obj := httpjs.DefaultBuilder(r, req)
				obj.Delete("Execute")
				obj.Set("AsyncExecute", asyncexecute(b, req))
				return obj
			}
			r.Set("HTTP", addon.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
