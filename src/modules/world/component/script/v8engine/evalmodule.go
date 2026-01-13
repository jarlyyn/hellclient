package v8engine

import (
	"context"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/v8local"
	"github.com/herb-go/v8local/v8plugin"
)

type Eval struct {
	runtime *v8local.Local
}

func (e *Eval) Run(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	v := e.runtime.NewLocal().RunScript(call.GetArg(0).String(), call.GetArg(1).String())
	return v
}
func initmodule(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
	jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
	r := jsp.Top
	eval := &Eval{runtime: r}
	global := r.Global()
	global.Set("eval", r.Context().NewFunctionTemplate(eval.Run).GetLocalFunction(r))
	next(ctx, plugin)
}

var ModuleEval = herbplugin.CreateModule("eval",
	initmodule,
	nil,
	nil,
)
