package jsengine

import (
	"context"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
)

var ModuleEval = herbplugin.CreateModule("eval",
	func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
		jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
		r := jsp.Runtime
		r.Set("eval", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
			v, err := r.RunScript(call.Argument(1).String(), call.Argument(0).String())
			if err != nil {
				panic(err)
			}
			return v
		})
		next(ctx, plugin)
	},
	nil,
	nil,
)
