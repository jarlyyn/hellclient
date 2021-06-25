package luaengine

import (
	"context"
	"modules/world"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	lua "github.com/yuin/gopher-lua"
)

var ModuleConstsTimerFlag = herbplugin.CreateModule("sendto",
	func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
		next(ctx, plugin)
		luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
		l := luapluing.LState
		consts := l.NewTable()
		consts.RawSetString("Enabled", lua.LNumber(world.TimerFlagEnabled))
		consts.RawSetString("AtTime", lua.LNumber(world.TimerFlagAtTime))
		consts.RawSetString("OneShot", lua.LNumber(world.TimerFlagOneShot))
		consts.RawSetString("TimerSpeedWalk", lua.LNumber(world.TimerFlagTimerSpeedWalk))
		consts.RawSetString("TimerNote", lua.LNumber(world.TimerFlagTimerNote))
		consts.RawSetString("ActiveWhenClosed", lua.LNumber(world.TimerFlagActiveWhenClosed))
		consts.RawSetString("Replace", lua.LNumber(world.TimerFlagReplace))
		consts.RawSetString("Temporary", lua.LNumber(world.TimerFlagTemporary))
		l.SetGlobal("timer_flag", consts)
		next(ctx, plugin)
	},
	nil,
	nil,
)
var ModuleConstsSendTo = herbplugin.CreateModule("sendto",
	func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
		next(ctx, plugin)
		luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
		l := luapluing.LState
		consts := l.NewTable()
		consts.RawSetString("world", lua.LNumber(world.SendtoWorld))
		consts.RawSetString("world", lua.LNumber(world.SendtoCommand))
		consts.RawSetString("output", lua.LNumber(world.SendtoOutput))
		consts.RawSetString("status", lua.LNumber(world.SendtoStatus))
		consts.RawSetString("notepad", lua.LNumber(world.SendtoNotepad))
		consts.RawSetString("notepadappend", lua.LNumber(world.SendtoNotepadAppend))
		consts.RawSetString("logfile", lua.LNumber(world.SendtoLogfile))
		consts.RawSetString("notepadreplace", lua.LNumber(world.SendtoNotepadReplace))
		consts.RawSetString("commandqueue", lua.LNumber(world.SendtoCommandqueue))
		consts.RawSetString("variable", lua.LNumber(world.SendtoVariable))
		consts.RawSetString("execute", lua.LNumber(world.SendtoExecute))
		consts.RawSetString("execute", lua.LNumber(world.SendtoExecute))
		consts.RawSetString("speedwalk", lua.LNumber(world.SendtoSpeedwalk))
		consts.RawSetString("script", lua.LNumber(world.SendtoScript))
		consts.RawSetString("immediate", lua.LNumber(world.SendtoImmediate))
		consts.RawSetString("scriptafteromit", lua.LNumber(world.SendtoScriptAfterOmit))
		l.SetGlobal("sendto", consts)
		next(ctx, plugin)
	},
	nil,
	nil,
)
