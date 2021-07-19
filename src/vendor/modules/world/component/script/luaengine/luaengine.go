package luaengine

import (
	"modules/world"
	"modules/world/bus"
	"sync"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/util"
	lua "github.com/yuin/gopher-lua"

	"github.com/herb-go/herbplugin/lua51plugin"
)

func newLuaInitializer(b *bus.Bus) *lua51plugin.Initializer {
	i := lua51plugin.NewInitializer()
	i.Entry = "main.lua"
	i.Modules = []*herbplugin.Module{
		lua51plugin.ModuleOpenlib,
		ModuleConstsSendTo,
		ModuleConstsTimerFlag,
		ModuleConstsAliasFlag,
		ModuleConstsTriggersFlag,
		ModuleRex,
		NewAPIModule(b),
		NewMapperModule(b),
	}
	return i
}

type LuaEngine struct {
	Locker       sync.RWMutex
	Plugin       *lua51plugin.Plugin
	onClose      string
	onDisconnect string
	onConnect    string
	onBroadCast  string
}

func NewLuaEngine() *LuaEngine {
	return &LuaEngine{
		Plugin: lua51plugin.New(),
	}
}

func (e *LuaEngine) Open(b *bus.Bus) error {
	opt := b.GetScriptPluginOptions()
	data := b.GetScriptData()
	e.onClose = data.OnClose
	e.onConnect = data.OnConnect
	e.onDisconnect = data.OnDisconnect
	e.onBroadCast = data.OnBroadcast
	err := util.Catch(func() {
		newLuaInitializer(b).MustApplyInitializer(e.Plugin)
	})
	if err != nil {
		return err
	}
	err = util.Catch(func() {
		herbplugin.Lanuch(e.Plugin, opt)
	})
	if err != nil {
		return err
	}
	if data.OnOpen != "" {
		b.HandleScriptError(e.Plugin.LState.DoString(data.OnOpen + "()"))
	}
	return nil
}
func (e *LuaEngine) Close(b *bus.Bus) {
	if e.onClose != "" {
		e.Call(b, e.Plugin.LState.GetGlobal(e.onClose))
		// b.HandleScriptError(e.Plugin.LState.DoString(e.onClose + "()"))
	}
	b.HandleScriptError(util.Catch(func() {
		e.Plugin.MustClosePlugin()
	}))
}
func (e *LuaEngine) OnConnect(b *bus.Bus) {
	if e.onConnect != "" {
		e.Call(b, e.Plugin.LState.GetGlobal(e.onConnect))
		// b.HandleScriptError(e.Plugin.LState.DoString(e.onConnect + "()"))
	}
}
func (e *LuaEngine) OnDisconnect(b *bus.Bus) {
	if e.onDisconnect != "" {
		e.Call(b, e.Plugin.LState.GetGlobal(e.onDisconnect))
		// b.HandleScriptError(e.Plugin.LState.DoString(e.onDisconnect + "()"))
	}
}

func (e *LuaEngine) ConvertStyle(L *lua.LState, line *world.Line) *lua.LTable {
	result := L.NewTable()
	for _, v := range line.Words {
		style := L.NewTable()
		style.RawSetString("text", lua.LString(v.Text))
		style.RawSetString("textcolour", lua.LNumber(world.Colours[v.Color]))
		style.RawSetString("backcolour", lua.LNumber(world.Colours[v.Background]))
		style.RawSetString("length", lua.LNumber(len(v.Text)))
		var s int
		if v.Bold {
			s = s + 1
		}
		style.RawSetString("length", lua.LNumber(s))
		result.Append(style)
	}
	return result
}
func (e *LuaEngine) OnTrigger(b *bus.Bus, line *world.Line, trigger *world.Trigger, result *world.MatchResult) {
	if trigger.Script == "" {
		return
	}
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(trigger.Script)
	L := e.Plugin.LState
	t := L.NewTable()
	for k, v := range result.List {
		t.RawSetInt(k, lua.LString(v))
	}
	for k, v := range result.Named {
		t.RawSetString(k, lua.LString(v))
	}
	e.Locker.Unlock()
	e.Call(b, fn, lua.LString(trigger.Name), lua.LString(line.Plain()), t, e.ConvertStyle(L, line))

}
func (e *LuaEngine) OnAlias(b *bus.Bus, message string, alias *world.Alias, result *world.MatchResult) {
	if alias.Script == "" {
		return
	}
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(alias.Script)
	L := e.Plugin.LState
	t := L.NewTable()
	for k, v := range result.List {
		t.RawSetInt(k, lua.LString(v))
	}
	for k, v := range result.Named {
		t.RawSetString(k, lua.LString(v))
	}
	e.Locker.Unlock()
	go e.Call(b, fn, lua.LString(alias.Name), lua.LString(message), t)

}
func (e *LuaEngine) OnTimer(b *bus.Bus, timer *world.Timer) {
	if timer.Script == "" {
		return
	}
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(timer.Script)
	e.Locker.Unlock()
	go e.Call(b, fn, lua.LString(timer.Name))
}
func (e *LuaEngine) OnBroadCast(b *bus.Bus, bc *world.Broadcast) {
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(e.onBroadCast)
	e.Locker.Unlock()
	go e.Call(b, fn, lua.LString(bc.Message), lua.LBool(bc.Global), lua.LString(bc.Channel), lua.LString(bc.ID))
}
func (e *LuaEngine) Run(b *bus.Bus, cmd string) {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	b.HandleScriptError(e.Plugin.LState.DoString(cmd))
}

func (e *LuaEngine) Call(b *bus.Bus, fn lua.LValue, args ...lua.LValue) {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	L := e.Plugin.LState
	if L == nil {
		return
	}
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, args...); err != nil {
		b.HandleScriptError(err)
	}
}
