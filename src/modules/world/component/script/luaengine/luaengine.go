package luaengine

import (
	"hellclient/modules/world"
	"hellclient/modules/world/bus"
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
		NewHTTPModule(b),
		NewMapperModule(b),
		NewMetronomeModule(b),
		NewUserinputModule(b),
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
	onResponse   string
	onHUDClick   string
	onSubneg     string
	onBuffer     string
	onFocus      string
	onBufferMin  int
	onBufferMax  int
}

func NewLuaEngine() *LuaEngine {
	return &LuaEngine{
		Plugin: lua51plugin.New(),
	}
}

func (e *LuaEngine) Open(b *bus.Bus) error {
	opt := b.GetPluginOptions()
	data := b.GetScriptData()
	e.onClose = data.OnClose
	e.onConnect = data.OnConnect
	e.onDisconnect = data.OnDisconnect
	e.onBroadCast = data.OnBroadcast
	e.onResponse = data.OnResponse
	e.onBuffer = data.OnBuffer
	e.onSubneg = data.OnSubneg
	e.onHUDClick = data.OnHUDClick
	e.onBufferMin = data.OnBufferMin
	e.onBufferMax = data.OnBufferMax
	e.onFocus = data.OnFocus
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
func (e *LuaEngine) OnResponse(b *bus.Bus, msg *world.Message) {
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(e.onResponse)
	e.Locker.Unlock()
	go e.Call(b, fn, lua.LString(msg.Type), lua.LString(msg.ID), lua.LString(msg.Data))
}

func (e *LuaEngine) OnHUDClick(b *bus.Bus, c *world.Click) {
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(e.onHUDClick)
	e.Locker.Unlock()
	go e.Call(b, fn, lua.LNumber(c.X), lua.LNumber(c.Y))
}
func (e *LuaEngine) OnBuffer(b *bus.Bus, data []byte) bool {
	e.Locker.Lock()
	if e.Plugin.LState == nil || e.onBuffer == "" {
		e.Locker.Unlock()
		return false
	}
	l := len(data)
	if l < e.onBufferMin || (e.onBufferMax > 0 && l > e.onBufferMax) {
		e.Locker.Unlock()
		return false
	}
	fn := e.Plugin.LState.GetGlobal(e.onBuffer)
	e.Locker.Unlock()
	v := e.Call(b, fn, lua.LString(data))
	return lua.LVAsBool(v)
}
func (e *LuaEngine) OnSubneg(b *bus.Bus, code byte, data []byte) bool {
	e.Locker.Lock()
	if e.Plugin.LState == nil || e.onSubneg == "" {
		e.Locker.Unlock()
		return false
	}
	fn := e.Plugin.LState.GetGlobal(e.onSubneg)
	e.Locker.Unlock()
	v := e.Call(b, fn, lua.LNumber(code), lua.LString(data))
	return lua.LVAsBool(v)
}
func (e *LuaEngine) OnFocus(b *bus.Bus) {
	e.Locker.Lock()
	if e.Plugin.LState == nil || e.onSubneg == "" {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(e.onSubneg)
	e.Locker.Unlock()
	e.Call(b, fn)
}
func (e *LuaEngine) OnCallback(b *bus.Bus, cb *world.Callback) {
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(cb.Script)
	e.Locker.Unlock()
	go e.Call(b, fn, lua.LString(cb.Script), lua.LString(cb.Name), lua.LString(cb.ID), lua.LNumber(cb.Code), lua.LString(cb.Data))

}
func (e *LuaEngine) OnAssist(b *bus.Bus, script string) {
	e.Locker.Lock()
	if e.Plugin.LState == nil {
		e.Locker.Unlock()
		return
	}
	fn := e.Plugin.LState.GetGlobal(script)
	e.Locker.Unlock()
	go e.Call(b, fn)
}
func (e *LuaEngine) Run(b *bus.Bus, cmd string) {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	b.HandleScriptError(e.Plugin.LState.DoString(cmd))
}

func (e *LuaEngine) Call(b *bus.Bus, fn lua.LValue, args ...lua.LValue) lua.LValue {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	L := e.Plugin.LState
	if L == nil {
		return nil
	}
	if fn.Type() == lua.LTNil {
		return nil
	}
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, args...); err != nil {
		b.HandleScriptError(err)
	}
	return L.Get(-1)
}
