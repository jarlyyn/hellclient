package v8engine

import (
	"context"

	"github.com/herb-go/herbplugin"
	"github.com/jarlyyn/v8js"
	"github.com/jarlyyn/v8js/v8plugin"
)

type Eval struct {
	runtime *v8js.Context
}

func (e *Eval) Run(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	v := e.runtime.RunScript(call.GetArg(0).String(), call.GetArg(1).String())
	return v.Consume()
}
func initmodule(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
	jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
	r := jsp.Runtime
	eval := &Eval{runtime: r}
	global := r.Global()
	global.Set("eval", r.NewFunctionTemplate(eval.Run).GetFunction(r).Consume())
	global.Release()
	next(ctx, plugin)
}

var ModuleEval = herbplugin.CreateModule("eval",
	initmodule,
	nil,
	nil,
)
