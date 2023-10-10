package luaengine

import (
	"context"
	"modules/world"
	"modules/world/bus"
	"time"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	lua "github.com/yuin/gopher-lua"
)

func NewMetronomeModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("metronome",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			m := l.NewTable()
			m.RawSetString("getbeats", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeBeats()))
				return 1
			}))
			m.RawSetString("GetBeats", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeBeats()))
				return 1
			}))
			m.RawSetString("setbeats", l.NewFunction(func(L *lua.LState) int {
				_ = l.Get(1) //this
				b.SetMetronomeBeats(l.ToInt(2))
				return 0
			}))

			m.RawSetString("SetBeats", l.NewFunction(func(L *lua.LState) int {
				_ = l.Get(1) //this
				b.SetMetronomeBeats(l.ToInt(2))
				return 0
			}))
			m.RawSetString("reset", l.NewFunction(func(L *lua.LState) int {
				b.DoResetMetronome()
				return 0
			}))
			m.RawSetString("Reset", l.NewFunction(func(L *lua.LState) int {
				b.DoResetMetronome()
				return 0
			}))
			m.RawSetString("getspace", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeSpace()))
				return 1
			}))
			m.RawSetString("GetSpace", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeSpace()))
				return 1
			}))
			getqueue := l.NewFunction(func(L *lua.LState) int {
				q := b.GetMetronomeQueue()
				t := l.NewTable()
				for k := range q {
					t.Append(lua.LString(q[k]))
				}
				l.Push(t)
				return 1
			})
			m.RawSetString("getqueue", getqueue)
			m.RawSetString("GetQueue", getqueue)
			m.RawSetString("discard", l.NewFunction(func(L *lua.LState) int {
				b.DoDiscardMetronome(L.ToBool(1))
				return 0
			}))
			m.RawSetString("Discard", l.NewFunction(func(L *lua.LState) int {
				b.DoDiscardMetronome(L.ToBool(1))
				return 0
			}))
			m.RawSetString("lockqueue", l.NewFunction(func(L *lua.LState) int {
				b.DoLockMetronomeQueue()
				return 0
			}))
			m.RawSetString("LockQueue", l.NewFunction(func(L *lua.LState) int {
				b.DoLockMetronomeQueue()
				return 0
			}))
			m.RawSetString("full", l.NewFunction(func(L *lua.LState) int {
				b.DoFullMetronome()
				return 0
			}))
			m.RawSetString("Full", l.NewFunction(func(L *lua.LState) int {
				b.DoFullMetronome()
				return 0
			}))
			m.RawSetString("fulltick", l.NewFunction(func(L *lua.LState) int {
				b.DoFullTickMetronome()
				return 0
			}))
			m.RawSetString("FullTick", l.NewFunction(func(L *lua.LState) int {
				b.DoFullTickMetronome()
				return 0
			}))
			m.RawSetString("getinterval", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeInterval() / time.Millisecond))
				return 1
			}))
			m.RawSetString("GetInterval", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeInterval() / time.Millisecond))
				return 1
			}))
			m.RawSetString("setinterval", l.NewFunction(func(L *lua.LState) int {
				_ = l.Get(1) //this
				b.SetMetronomeInterval(time.Duration(l.ToInt64(2)) * time.Millisecond)
				return 0
			}))
			m.RawSetString("SetInterval", l.NewFunction(func(L *lua.LState) int {
				_ = l.Get(1) //this
				b.SetMetronomeInterval(time.Duration(l.ToInt64(2)) * time.Millisecond)
				return 0
			}))
			m.RawSetString("gettick", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeTick() / time.Millisecond))
				return 1
			}))
			m.RawSetString("GetTick", l.NewFunction(func(L *lua.LState) int {
				l.Push(lua.LNumber(b.GetMetronomeTick() / time.Millisecond))
				return 1
			}))
			m.RawSetString("settick", l.NewFunction(func(L *lua.LState) int {
				_ = l.Get(1) //this
				b.SetMetronomeTick(time.Duration(l.ToInt64(2)) * time.Millisecond)
				return 0
			}))
			m.RawSetString("SetTick", l.NewFunction(func(L *lua.LState) int {
				_ = l.Get(1) //this
				b.SetMetronomeTick(time.Duration(l.ToInt64(2)) * time.Millisecond)
				return 0
			}))
			push := l.NewFunction(func(L *lua.LState) int {
				_ = l.Get(1) //this
				args := l.ToTable(2)
				grouped := l.ToBool(3)
				echo := l.ToBool(4)
				cmds := []*world.Command{}
				args.ForEach(func(key lua.LValue, value lua.LValue) {
					c := world.CreateCommand(value.String())
					c.Echo = echo
					cmds = append(cmds, c)
				})
				b.DoPushMetronome(cmds, grouped)
				return 0
			})
			m.RawSetString("push", push)
			m.RawSetString("Push", push)
			l.SetGlobal("Metronome", m)
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
