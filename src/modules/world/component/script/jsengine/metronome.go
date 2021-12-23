package jsengine

import (
	"context"
	"errors"
	"hellclient/modules/world"
	"hellclient/modules/world/bus"
	"time"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
)

func NewMetronomeModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("metronome",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			m := r.NewObject()
			m.Set("getbeats", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				return r.ToValue(b.GetMetronomeBeats())
			})
			m.Set("setbeats", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.SetMetronomeBeats(int(call.Argument(0).ToInteger()))
				return nil
			})
			m.Set("reset", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.DoResetMetronome()
				return nil
			})
			m.Set("getspace", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				return r.ToValue(b.GetMetronomeSpace())
			})
			m.Set("getqueue", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				return r.ToValue(b.GetMetronomeQueue())
			})
			m.Set("discard", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.DoDiscardMetronome(call.Argument(0).ToBoolean())
				return nil
			})
			m.Set("lockqueue", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.DoLockMetronomeQueue()
				return nil
			})
			m.Set("full", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.DoFullMetronome()
				return nil
			})
			m.Set("fulltick", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.DoFullTickMetronome()
				return nil
			})
			m.Set("getinterval", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				return r.ToValue(b.GetMetronomeInterval() / time.Millisecond)
			})
			m.Set("setinterval", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.SetMetronomeInterval(time.Duration(call.Argument(0).ToInteger()) * time.Millisecond)
				return nil
			})
			m.Set("gettick", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				return r.ToValue(b.GetMetronomeTick() / time.Millisecond)
			})
			m.Set("settick", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				b.SetMetronomeTick(time.Duration(call.Argument(0).ToInteger()) * time.Millisecond)
				return nil
			})
			m.Set("push", func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
				args := []string{}
				err := r.ExportTo(call.Argument(0), &args)
				if err != nil {
					panic(errors.New("pushed commands must be string array"))
				}
				grouped := call.Argument(1).ToBoolean()
				echo := call.Argument(2).ToBoolean()
				cmds := []*world.Command{}
				for k := range args {
					c := world.CreateCommand(args[k])
					c.Echo = echo
					cmds = append(cmds, c)
				}
				b.DoPushMetronome(cmds, grouped)
				return nil
			})
			r.Set("Metronome", m)
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
