package luaengine

import (
	"modules/world"
	"modules/world/bus"

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
	Plugin       *lua51plugin.Plugin
	onClose      string
	onDisconnect string
	onConnect    string
}

func NewLuaEngeine() *LuaEngine {
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
		b.HandleScriptError(e.Plugin.LState.DoString(e.onClose + "()"))
	}
	b.HandleScriptError(util.Catch(func() {
		e.Plugin.MustClosePlugin()
	}))
}
func (e *LuaEngine) OnConnect(b *bus.Bus) {
	if e.onConnect != "" {
		b.HandleScriptError(e.Plugin.LState.DoString(e.onConnect + "()"))
	}
}
func (e *LuaEngine) OnDisconnect(b *bus.Bus) {
	if e.onDisconnect != "" {
		b.HandleScriptError(e.Plugin.LState.DoString(e.onDisconnect + "()"))
	}
}

func (e *LuaEngine) ConvertStyle(L *lua.LState, line *world.Line) *lua.LTable {
	result := L.NewTable()
	for _, v := range line.Words {
		style := L.NewTable()
		style.RawSetString("text", lua.LString(v.Text))
		style.RawSetString("textcolour", lua.LString(v.Color))
		style.RawSetString("backcolour", lua.LString(v.Background))
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
	L := e.Plugin.LState
	t := L.NewTable()
	for k, v := range result.List {
		t.RawSetInt(k, lua.LString(v))
	}
	for k, v := range result.Named {
		t.RawSetString(k, lua.LString(v))
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(trigger.Script),
		NRet:    0,
		Protect: true,
	}, lua.LString(trigger.Name), lua.LString(line.Plain()), t, e.ConvertStyle(L, line)); err != nil {
		b.HandleScriptError(err)
	}

}
func (e *LuaEngine) OnAlias(b *bus.Bus, message string, alias *world.Alias, result *world.MatchResult) {
	if alias.Script == "" {
		return
	}
	L := e.Plugin.LState
	t := L.NewTable()
	for _, v := range result.List {
		t.Append(lua.LString(v))
	}
	for k, v := range result.Named {
		t.RawSetString(k, lua.LString(v))
	}
	if err := L.CallByParam(lua.P{
		Fn:   L.GetGlobal(alias.Script),
		NRet: 0,
	}, lua.LString(alias.Name), lua.LString(message), t); err != nil {
		b.HandleScriptError(err)
	}

}
func (e *LuaEngine) OnTimer(b *bus.Bus, timer *world.Timer) {
	if timer.Script == "" {
		return
	}

	L := e.Plugin.LState
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(timer.Script),
		NRet:    0,
		Protect: true,
	}, lua.LString(timer.Name)); err != nil {
		b.HandleScriptError(err)
	}
}
func (e *LuaEngine) Run(b *bus.Bus, cmd string) {
	b.HandleScriptError(e.Plugin.LState.DoString(cmd))
}
