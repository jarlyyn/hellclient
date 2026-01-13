package v8engine

import (
	"context"
	"modules/world"
	"modules/world/bus"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/v8local"
	"github.com/herb-go/v8local/plugins/httpaddon/httpv8"
	"github.com/herb-go/v8local/v8plugin"
)

func AsyncExecute(a *httpv8.Addon, b *builder) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		callback := ""
		if call.GetArg(0).Boolean() {
			callback = call.GetArg(0).String()
		}
		id := call.This().Get("id")
		req := a.LoadReq(id.String())
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

				b.bus.DoSendCallbackToScript(cb)
			}
		})
		return nil
	}
}

type builder struct {
	bus                  *bus.Bus
	AsyncExecuteFunction *v8local.JsValue
}

func (b *builder) Init(r *v8local.Local, a *httpv8.Addon) {
	b.AsyncExecuteFunction = r.NewFunction(AsyncExecute(a, b)).AsExported()
}
func (b *builder) Build(r *v8local.Local, a *httpv8.Addon, req *httpv8.Request) *v8local.JsValue {
	obj := httpv8.DefaultBuilder(r, a, req)
	obj.Delete("Execute")

	obj.Set("AsyncExecute", b.AsyncExecuteFunction)
	return obj

}

func NewHTTPModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("http",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Top
			addon := httpv8.Create(plugin)
			builder := &builder{
				bus: b,
			}
			builder.Init(r, addon)
			addon.Builder = builder.Build
			global := r.Global()
			global.Set("HTTP", addon.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
