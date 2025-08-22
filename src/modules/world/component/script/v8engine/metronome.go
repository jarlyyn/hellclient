package v8engine

import (
	"context"
	"modules/world"
	"modules/world/bus"
	"time"

	"github.com/herb-go/herbplugin"
	"github.com/jarlyyn/v8js"
	"github.com/jarlyyn/v8js/v8plugin"
)

func NewMetronomeModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("metronome",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			m := jsp.Runtime.NewObject()
			m.Set("getbeats", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeBeats())).Consume()
			}).Consume())
			m.Set("GetBeats", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeBeats())).Consume()
			}).Consume())
			m.Set("setbeats", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.SetMetronomeBeats(int(call.GetArg(0).Integer()))
				return nil
			}).Consume())
			m.Set("SetBeats", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.SetMetronomeBeats(int(call.GetArg(0).Integer()))
				return nil
			}).Consume())
			m.Set("reset", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoResetMetronome()
				return nil
			}).Consume())
			m.Set("Reset", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoResetMetronome()
				return nil
			}).Consume())
			m.Set("getspace", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeSpace())).Consume()
			}).Consume())
			m.Set("GetSpace", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeSpace())).Consume()
			}).Consume())
			m.Set("getqueue", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewStringArray(b.GetMetronomeQueue()...).Consume()
			}).Consume())
			m.Set("GetQueue", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewStringArray(b.GetMetronomeQueue()...).Consume()
			}).Consume())
			m.Set("discard", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoDiscardMetronome(call.GetArg(0).Boolean())
				return nil
			}).Consume())
			m.Set("Discard", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoDiscardMetronome(call.GetArg(0).Boolean())
				return nil
			}).Consume())
			m.Set("lockqueue", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoLockMetronomeQueue()
				return nil
			}).Consume())
			m.Set("Lockqueue", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoLockMetronomeQueue()
				return nil
			}).Consume())
			m.Set("full", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoFullMetronome()
				return nil
			}).Consume())
			m.Set("Full", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoFullMetronome()
				return nil
			}).Consume())
			m.Set("fulltick", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoFullTickMetronome()
				return nil
			}).Consume())
			m.Set("FullTick", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.DoFullTickMetronome()
				return nil
			}).Consume())
			m.Set("getinterval", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeInterval() / time.Millisecond)).Consume()
			}).Consume())
			m.Set("GetInterval", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeInterval() / time.Millisecond)).Consume()
			}).Consume())
			m.Set("setinterval", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.SetMetronomeInterval(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}).Consume())
			m.Set("SetInterval", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.SetMetronomeInterval(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}).Consume())
			m.Set("gettick", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeTick() / time.Millisecond)).Consume()
			}).Consume())
			m.Set("GetTick", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				return call.Context().NewInt32(int32(b.GetMetronomeTick() / time.Millisecond)).Consume()
			}).Consume())
			m.Set("settick", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.SetMetronomeTick(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}).Consume())
			m.Set("SetTick", jsp.Runtime.NewFunction(func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

				b.SetMetronomeTick(time.Duration(call.GetArg(0).Integer()) * time.Millisecond)
				return nil
			}).Consume())

			push := func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
			m.Set("push", jsp.Runtime.NewFunction(push).Consume())
			m.Set("Push", jsp.Runtime.NewFunction(push).Consume())
			global := r.Global()
			global.Set("Metronome", m.Consume())
			global.Release()
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
