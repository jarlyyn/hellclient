package v8engine

import (
	"context"
	"modules/world"
	"modules/world/bus"
	"time"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/v8local"
	"github.com/herb-go/v8local/v8plugin"
)

func NewMetronomeModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("metronome",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Top
			m := jsp.Top.NewObject()
			m.Set("getbeats", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeBeats()))
			}))
			m.Set("GetBeats", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeBeats()))
			}))
			m.Set("setbeats", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.SetMetronomeBeats(int(call.GetArg(0).Integer()))
				return nil
			}))
			m.Set("SetBeats", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.SetMetronomeBeats(int(call.GetArg(0).Integer()))
				return nil
			}))
			m.Set("reset", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoResetMetronome()
				return nil
			}))
			m.Set("Reset", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoResetMetronome()
				return nil
			}))
			m.Set("getspace", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeSpace()))
			}))
			m.Set("GetSpace", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeSpace()))
			}))
			m.Set("getqueue", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewStringArray(b.GetMetronomeQueue()...)
			}))
			m.Set("GetQueue", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewStringArray(b.GetMetronomeQueue()...)
			}))
			m.Set("discard", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoDiscardMetronome(call.GetArg(0).Boolean())
				return nil
			}))
			m.Set("Discard", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoDiscardMetronome(call.GetArg(0).Boolean())
				return nil
			}))
			m.Set("lockqueue", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoLockMetronomeQueue()
				return nil
			}))
			m.Set("Lockqueue", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoLockMetronomeQueue()
				return nil
			}))
			m.Set("full", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoFullMetronome()
				return nil
			}))
			m.Set("Full", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoFullMetronome()
				return nil
			}))
			m.Set("fulltick", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoFullTickMetronome()
				return nil
			}))
			m.Set("FullTick", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.DoFullTickMetronome()
				return nil
			}))
			m.Set("getinterval", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeInterval() / time.Millisecond))
			}))
			m.Set("GetInterval", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeInterval() / time.Millisecond))
			}))
			m.Set("setinterval", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.SetMetronomeInterval(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}))
			m.Set("SetInterval", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.SetMetronomeInterval(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}))
			m.Set("gettick", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeTick() / time.Millisecond))
			}))
			m.Set("GetTick", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				return call.Local().NewInt32(int32(b.GetMetronomeTick() / time.Millisecond))
			}))
			m.Set("settick", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.SetMetronomeTick(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}))
			m.Set("SetTick", jsp.Top.NewFunction(func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				b.SetMetronomeTick(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}))

			push := func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

				args := call.GetArg(0).StringArrry()
				grouped := call.GetArg(1).Boolean()
				echo := call.GetArg(2).Boolean()
				cmds := []*world.Command{}
				for k := range args {
					c := world.CreateCommand(args[k])
					c.Echo = echo
					cmds = append(cmds, c)
				}
				b.DoPushMetronome(cmds, grouped)
				return nil
			}
			m.Set("push", jsp.Top.NewFunction(push))
			m.Set("Push", jsp.Top.NewFunction(push))
			global := r.Global()
			global.Set("Metronome", m)
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
