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
		NewAPIModule(b),
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

func (e *LuaEngine) OnTrigger(*bus.Bus) {

}
func (e *LuaEngine) OnAlias(*bus.Bus) {

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