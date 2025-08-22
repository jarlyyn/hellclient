package v8engine

import (
	"context"
	"modules/world"
	"modules/world/bus"

	"github.com/herb-go/herbplugin"
	"github.com/jarlyyn/v8js"
	"github.com/jarlyyn/v8js/plugins/httpaddon/httpv8"
	"github.com/jarlyyn/v8js/v8plugin"
)

type asyncExecute struct {
	bus *bus.Bus
	req *httpv8.Request
}

func (e *asyncExecute) Execute(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	callback := ""
	if call.GetArg(0).Boolean() {
		callback = call.GetArg(0).String()
	}
	e.req.Request.AsyncExecute(func(err error) {
		if callback != "" {
			cb := world.NewCallback()
			cb.Name = "httpexecute"
			cb.ID = e.req.Request.ID
			cb.Script = callback
			if err != nil {
				cb.Code = -1
				cb.Data = err.Error()
			} else {
				cb.Code = 0
				cb.Data = e.req.Request.GetURL()
			}

			e.bus.DoSendCallbackToScript(cb)
		}
	})
	return nil
}

type builder struct {
	bus *bus.Bus
}

func (b *builder) Build(r *v8js.Context, req *httpv8.Request) *v8js.JsValue {
	obj := httpv8.DefaultBuilder(r, req)
	obj.Delete("Execute")
	a := &asyncExecute{
		bus: b.bus,
		req: req,
	}
	obj.Set("AsyncExecute", r.NewFunction(a.Execute).Consume())
	return obj

}

func NewHTTPModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("http",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			addon := httpv8.Create(plugin)
			builder := &builder{
				bus: b,
			}
			addon.Builder = builder.Build
			global := r.Global()
			global.Set("HTTP", addon.Convert(r).Consume())
			global.Release()
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
