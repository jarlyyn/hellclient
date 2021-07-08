package jsengine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/api"
	"strings"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
)

func createApi(b *bus.Bus) *jsapi {
	return &jsapi{
		API: &api.API{
			Bus: b,
		},
	}
}

type jsapi struct {
	API *api.API
}

func AppendToWorld(r *goja.Runtime, world *goja.Object, name string, call func(call goja.FunctionCall, r *goja.Runtime) goja.Value) {
	r.Set(name, call)
	world.Set(name, call)
}
func (a *jsapi) InstallAPIs(p herbplugin.Plugin) {
	jp := p.(*jsplugin.Plugin)
	world := jp.Runtime.NewObject()
	jp.Runtime.Set("world", world)
	AppendToWorld(jp.Runtime, world, "print", a.Print)
	AppendToWorld(jp.Runtime, world, "Note", a.Note)
	AppendToWorld(jp.Runtime, world, "SendImmediate", a.SendImmediate)
	AppendToWorld(jp.Runtime, world, "Send", a.Send)
	AppendToWorld(jp.Runtime, world, "SendNoEcho", a.SendNoEcho)
	AppendToWorld(jp.Runtime, world, "GetVariable", a.GetVariable)
	AppendToWorld(jp.Runtime, world, "SetVariable", a.SetVariable)
	AppendToWorld(jp.Runtime, world, "DeleteVariable", a.DeleteVariable)
	AppendToWorld(jp.Runtime, world, "GetVariableList", a.GetVariableList)
	AppendToWorld(jp.Runtime, world, "Version", a.Version)
	AppendToWorld(jp.Runtime, world, "Hash", a.Hash)
	AppendToWorld(jp.Runtime, world, "Base64Encode", a.Base64Encode)
	AppendToWorld(jp.Runtime, world, "Base64Decode", a.Base64Decode)
	AppendToWorld(jp.Runtime, world, "Connect", a.Connect)
	AppendToWorld(jp.Runtime, world, "IsConnected", a.IsConnected)
	AppendToWorld(jp.Runtime, world, "Disconnect", a.Disconnect)
	AppendToWorld(jp.Runtime, world, "GetWorldById", a.GetWorldById)
	AppendToWorld(jp.Runtime, world, "GetWorld", a.GetWorld)
	AppendToWorld(jp.Runtime, world, "GetWorldID", a.GetWorldID)
	AppendToWorld(jp.Runtime, world, "GetWorldIdList", a.GetWorldIdList)
	AppendToWorld(jp.Runtime, world, "GetWorldList", a.GetWorldList)
	AppendToWorld(jp.Runtime, world, "WorldName", a.WorldName)
	AppendToWorld(jp.Runtime, world, "WorldAddress", a.WorldAddress)
	AppendToWorld(jp.Runtime, world, "WorldPort", a.WorldPort)
	AppendToWorld(jp.Runtime, world, "Trim", a.Trim)
	AppendToWorld(jp.Runtime, world, "GetUniqueNumber", a.GetUniqueNumber)
	AppendToWorld(jp.Runtime, world, "GetUniqueID", a.GetUniqueID)
	AppendToWorld(jp.Runtime, world, "CreateGUID", a.CreateGUID)
	AppendToWorld(jp.Runtime, world, "FlashIcon", a.FlashIcon)
	AppendToWorld(jp.Runtime, world, "SetStatus", a.SetStatus)
	AppendToWorld(jp.Runtime, world, "Execute", a.Execute)
	AppendToWorld(jp.Runtime, world, "DeleteCommandHistory", a.DeleteCommandHistory)
	AppendToWorld(jp.Runtime, world, "DiscardQueue", a.DiscardQueue)
	AppendToWorld(jp.Runtime, world, "GetQueue", a.GetQueue)
	AppendToWorld(jp.Runtime, world, "Queue", a.Queue)

	AppendToWorld(jp.Runtime, world, "DoAfter", a.DoAfter)
	AppendToWorld(jp.Runtime, world, "DoAfterNote", a.DoAfterNote)
	AppendToWorld(jp.Runtime, world, "DoAfterSpeedWalk", a.DoAfterSpeedWalk)
	AppendToWorld(jp.Runtime, world, "DoAfterSpecial", a.DoAfterSpecial)
	AppendToWorld(jp.Runtime, world, "AddTimer", a.AddTimer)
	AppendToWorld(jp.Runtime, world, "DeleteTimer", a.DeleteTimer)
	AppendToWorld(jp.Runtime, world, "DeleteTemporaryTimers", a.DeleteTemporaryTimers)
	AppendToWorld(jp.Runtime, world, "DeleteTimerGroup", a.DeleteTimerGroup)
	AppendToWorld(jp.Runtime, world, "EnableTimer", a.EnableTimer)
	AppendToWorld(jp.Runtime, world, "EnableTimerGroup", a.EnableTimerGroup)
	// AppendToWorld(jp,"GetTimerList", a.GetTimerList)
	// AppendToWorld(jp,"IsTimer", a.IsTimer)
	// AppendToWorld(jp,"ResetTimer", a.ResetTimer)
	// AppendToWorld(jp,"ResetTimers", a.ResetTimers)
	// AppendToWorld(jp,"GetTimerOption", a.GetTimerOption)
	// AppendToWorld(jp,"SetTimerOption", a.SetTimerOption)
	// AppendToWorld(jp,"AddAlias", a.AddAlias)
	// AppendToWorld(jp,"DeleteAlias", a.DeleteAlias)
	// AppendToWorld(jp,"DeleteTemporaryAliases", a.DeleteTemporaryAliases)
	// AppendToWorld(jp,"DeleteAliasGroup", a.DeleteAliasGroup)
	// AppendToWorld(jp,"EnableAlias", a.EnableAlias)
	// AppendToWorld(jp,"EnableAliasGroup", a.EnableAliasGroup)
	// AppendToWorld(jp,"GetAliasList", a.GetAliasList)
	// AppendToWorld(jp,"IsAlias", a.IsAlias)
	// AppendToWorld(jp,"GetAliasOption", a.GetAliasOption)
	// AppendToWorld(jp,"SetAliasOption", a.SetAliasOption)

	// AppendToWorld(jp,"AddTrigger", a.AddTrigger)
	// AppendToWorld(jp,"AddTriggerEx", a.AddTrigger)
	// AppendToWorld(jp,"DeleteTrigger", a.DeleteTrigger)
	// AppendToWorld(jp,"DeleteTemporaryTriggers", a.DeleteTemporaryTriggers)
	// AppendToWorld(jp,"DeleteTriggerGroup", a.DeleteTriggerGroup)
	// AppendToWorld(jp,"EnableTrigger", a.EnableTrigger)
	// AppendToWorld(jp,"EnableTriggerGroup", a.EnableTriggerGroup)
	// AppendToWorld(jp,"GetTriggerList", a.GetTriggerList)
	// AppendToWorld(jp,"IsTrigger", a.IsTrigger)
	// AppendToWorld(jp,"GetTriggerOption", a.GetTriggerOption)
	// AppendToWorld(jp,"SetTriggerOption", a.SetTriggerOption)
	// AppendToWorld(jp,"StopEvaluatingTriggers", a.StopEvaluatingTriggers)

	// AppendToWorld(jp,"ColourNameToRGB", a.ColourNameToRGB)
	// AppendToWorld(jp,"SetSpeedWalkDelay", a.SetSpeedWalkDelay)
	// AppendToWorld(jp,"GetSpeedWalkDelay", a.GetSpeedWalkDelay)

	// AppendToWorld(jp,"ReadFile", a.NewReadFileAPI(p))
	// AppendToWorld(jp,"ReadLines", a.NewReadLinesAPI(p))
	// AppendToWorld(jp,"SplitN", a.SplitNfunc)
	// AppendToWorld(jp,"UTF8Len", a.UTF8Len)
	// AppendToWorld(jp,"UTF8Sub", a.UTF8Sub)

}
func (a *jsapi) Print(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	msg := []string{}
	for _, v := range call.Arguments {
		msg = append(msg, v.String())
	}
	a.API.Note(strings.Join(msg, " "))
	return goja.Null()
}
func (a *jsapi) Note(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	info := call.Argument(0).String()

	a.API.Note(info)
	return goja.Null()
}
func (a *jsapi) SendImmediate(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	info := call.Argument(0).String()

	return r.ToValue(a.API.SendImmediate(info))

}
func (a *jsapi) Send(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	info := call.Argument(0).String()

	res := a.API.Send(info)
	return r.ToValue(res)
}
func (a *jsapi) Execute(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	info := call.Argument(0).String()
	return r.ToValue(a.API.Execute(info))
}
func (a *jsapi) SendNoEcho(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	info := call.Argument(0).String()

	return r.ToValue(a.API.SendNoEcho(info))
}
func (a *jsapi) GetVariable(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	val := a.API.GetVariable(call.Argument(0).String())
	return r.ToValue(val)
}
func (a *jsapi) DeleteVariable(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteVariable(name))
}
func (a *jsapi) SetVariable(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	value := call.Argument(1).String()
	return r.ToValue(a.API.SetVariable(name, value))
}
func (a *jsapi) GetVariableList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	list := a.API.GetVariableList()
	return r.ToValue(list)

}
func (a *jsapi) Version(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.Version())
}
func (a *jsapi) Hash(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.Hash(name))
}
func (a *jsapi) Base64Encode(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	src := call.Argument(0).String()
	ok := call.Argument(1).ToBoolean()
	return r.ToValue(a.API.Base64Encode(src, ok))
}
func (a *jsapi) Base64Decode(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	src := call.Argument(0).String()
	result := a.API.Base64Decode(src)
	if result == nil {
		return goja.Null()
	}
	return r.ToValue(*result)
}
func (a *jsapi) Connect(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.Connect())
}
func (a *jsapi) IsConnected(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.IsConnected())

}
func (a *jsapi) Disconnect(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.Disconnect())
}

func (a *jsapi) GetWorldById(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return goja.Null()
}

func (a *jsapi) GetWorld(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return goja.Null()
}

func (a *jsapi) GetWorldID(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetWorldID())

}
func (a *jsapi) GetWorldIdList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ar := r.NewArray()
	return r.ToValue(ar)
}
func (a *jsapi) GetWorldList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ar := r.NewArray()

	return r.ToValue(ar)
}
func (a *jsapi) WorldName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.WorldName())
}
func (a *jsapi) WorldAddress(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.WorldAddress())
}
func (a *jsapi) WorldPort(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.WorldPort())
}
func (a *jsapi) Trim(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	src := call.Argument(0).String()
	return r.ToValue(a.API.Trim(src))
}
func (a *jsapi) GetUniqueNumber(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetUniqueNumber())
}
func (a *jsapi) GetUniqueID(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetUniqueID())
}
func (a *jsapi) CreateGUID(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.CreateGUID())
}
func (a *jsapi) FlashIcon(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.FlashIcon()
	return goja.Null()
}
func (a *jsapi) SetStatus(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	a.API.SetStatus(text)
	return goja.Null()
}
func (a *jsapi) DeleteCommandHistory(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.DeleteCommandHistory()
	return goja.Null()
}
func (a *jsapi) DiscardQueue(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.DiscardQueue())
}
func (a *jsapi) GetQueue(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	cmds := a.API.GetQueue()
	return r.ToValue(cmds)
}
func (a *jsapi) Queue(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.Queue(call.Argument(0).String(), call.Argument(2).ToBoolean()))
}
func (a *jsapi) DoAfter(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(1).ToFloat()
	send := call.Argument(1).String()
	return r.ToValue(a.API.DoAfter(seconds, send))
}
func (a *jsapi) DoAfterNote(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(1).ToFloat()
	send := call.Argument(1).String()
	return r.ToValue(a.API.DoAfterNote(seconds, send))

}
func (a *jsapi) DoAfterSpeedWalk(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(1).ToFloat()
	send := call.Argument(1).String()
	return r.ToValue(a.API.DoAfterSpeedWalk(seconds, send))
}
func (a *jsapi) DoAfterSpecial(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(1).ToFloat()
	send := call.Argument(1).String()
	sendto := int(call.Argument(3).ToInteger())
	return r.ToValue(a.API.DoAfterSpecial(seconds, send, sendto))

}
func (a *jsapi) AddTimer(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	hour := int(call.Argument(1).ToInteger())
	min := int(call.Argument(2).ToInteger())
	seconds := call.Argument(3).ToFloat()
	send := call.Argument(4).String()
	flags := int(call.Argument(5).ToInteger())
	script := call.Argument(6).String()
	return r.ToValue(a.API.AddTimer(name, hour, min, seconds, send, flags, script))
}
func (a *jsapi) DeleteTimer(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteTimer(name))

}
func (a *jsapi) DeleteTemporaryTimers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.DeleteTemporaryTimers())

}
func (a *jsapi) DeleteTimerGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteTimerGroup(name))
}

func (a *jsapi) EnableTimer(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	enabled := call.Argument(2).ToBoolean()
	return r.ToValue(a.API.EnableTimer(name, enabled))
}
func (a *jsapi) EnableTimerGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	group := call.Argument(0).String()
	enabled := call.Argument(2).ToBoolean()
	return r.ToValue(a.API.EnableTimerGroup(group, enabled))
}

// func (a *jsapi) GetTimerList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	list := a.API.GetTimerList()
// 	reuslt := L.NewTable()
// 	for _, v := range list {
// 		reuslt.Append(lua.LString(v))
// 	}
// 	L.Push(reuslt)
// 	return 1
// }
// func (a *jsapi) IsTimer(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.IsTimer(name)))
// 	return 1
// }

// func (a *jsapi) ResetTimer(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.ResetTimer(name)))
// 	return 1
// }

// func (a *jsapi) ResetTimers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	a.API.ResetTimers()
// 	return goja.Null()
// }

// func (a *jsapi) GetTimerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	option := call.Argument(1).String()
// 	result, code := a.API.GetTimerOption(name, option)
// 	if code != api.EOK {
// 		L.Push(lua.LNil)
// 	} else {
// 		switch option {
// 		case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
// 			L.Push(lua.LBool(result == world.StringYes))
// 		case "group", "name", "script", "send", "variable":
// 			L.Push(lua.LString(result))
// 		case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "second", "send_to", "user":
// 			i, _ := strconv.Atoi(result)
// 			L.Push(lua.LNumber(i))
// 		default:
// 			L.Push(lua.LNil)
// 		}
// 	}
// 	return 1
// }
// func (a *jsapi) SetTimerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	option := call.Argument(1).String()
// 	var value string
// 	switch option {
// 	case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
// 		if MustBool(call.Argument(3) {
// 			value = world.StringYes
// 		} else {
// 			value = ""
// 		}
// 	case "group", "name", "script", "send", "variable":
// 		value = MustString(call.Argument(2))
// 	case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "second", "send_to", "user":
// 		value = MustString(call.Argument(2))
// 	}
// 	L.Push(lua.LNumber(a.API.SetTimerOption(name, option, value)))
// 	return 1
// }

// func (a *jsapi) AddAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	match := call.Argument(1).String()
// 	send := MustString(call.Argument(2))
// 	flags := int(L.ToNumber(4))
// 	script := L.ToString(5)
// 	L.Push(lua.LNumber(a.API.AddAlias(name, match, send, flags, script)))
// 	return 1
// }
// func (a *jsapi) DeleteAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.DeleteAlias(name)))
// 	return 1
// }
// func (a *jsapi) DeleteTemporaryAliases(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	L.Push(lua.LNumber(a.API.DeleteTemporaryTimers()))
// 	return 1

// }
// func (a *jsapi) DeleteAliasGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.DeleteAliasGroup(name)))
// 	return 1
// }

// func (a *jsapi) EnableAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	enabled := MustBool(call.Argument(2)
// 	L.Push(lua.LNumber(a.API.EnableAlias(name, enabled)))
// 	return 1
// }
// func (a *jsapi) EnableAliasGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	group := call.Argument(0).String()
// 	enabled := MustBool(call.Argument(2)
// 	L.Push(lua.LNumber(a.API.EnableAliasGroup(group, enabled)))
// 	return 1
// }

// func (a *jsapi) GetAliasList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	list := a.API.GetAliasList()
// 	reuslt := L.NewTable()
// 	for _, v := range list {
// 		reuslt.Append(lua.LString(v))
// 	}
// 	L.Push(reuslt)
// 	return 1
// }
// func (a *jsapi) IsAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.IsAlias(name)))
// 	return 1
// }

// func (a *jsapi) GetAliasOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	option := call.Argument(1).String()
// 	result, code := a.API.GetTimerOption(name, option)
// 	if code != api.EOK {
// 		L.Push(lua.LNil)
// 	} else {
// 		switch option {
// 		case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
// 			L.Push(lua.LBool(result == world.StringYes))
// 		case "group", "name", "match", "script", "send", "variable":
// 			L.Push(lua.LString(result))
// 		case "send_to", "user", "sequence":
// 			i, _ := strconv.Atoi(result)
// 			L.Push(lua.LNumber(i))
// 		default:
// 			L.Push(lua.LNil)
// 		}
// 	}
// 	return 1
// }
// func (a *jsapi) SetAliasOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	option := call.Argument(1).String()
// 	var value string
// 	switch option {
// 	case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
// 		if MustBool(call.Argument(3) {
// 			value = world.StringYes
// 		} else {
// 			value = ""
// 		}
// 	case "group", "name", "match", "script", "send", "variable":
// 		value = MustString(call.Argument(2))
// 	case "send_to", "user", "sequence":
// 		value = MustString(call.Argument(2))
// 	}
// 	L.Push(lua.LNumber(a.API.SetTimerOption(name, option, value)))
// 	return 1
// }

// func (a *jsapi) AddTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	match := call.Argument(1).String()
// 	send := MustString(call.Argument(2))
// 	flags := int(L.ToNumber(4))
// 	color := int(L.ToNumber(5))
// 	wildcard := int(L.ToNumber(6))
// 	sound := L.ToString(7)
// 	script := L.ToString(8)
// 	L.Push(lua.LNumber(a.API.AddTrigger(name, match, send, flags, color, wildcard, sound, script)))
// 	return 1
// }
// func (a *jsapi) AddTriggerEx(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	match := call.Argument(1).String()
// 	send := MustString(call.Argument(2))
// 	flags := int(L.ToNumber(4))
// 	color := int(L.ToNumber(5))
// 	wildcard := int(L.ToNumber(6))
// 	sound := L.ToString(7)
// 	script := L.ToString(8)
// 	sendto := int(L.ToNumber(9))
// 	sequence := int(L.ToNumber(10))
// 	L.Push(lua.LNumber(a.API.AddTriggerEx(name, match, send, flags, color, wildcard, sound, script, sendto, sequence)))
// 	return 1
// }
// func (a *jsapi) DeleteTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.DeleteTrigger(name)))
// 	return 1
// }
// func (a *jsapi) DeleteTemporaryTriggers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	L.Push(lua.LNumber(a.API.DeleteTemporaryTimers()))
// 	return 1

// }
// func (a *jsapi) DeleteTriggerGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.DeleteTriggerGroup(name)))
// 	return 1
// }

// func (a *jsapi) EnableTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	enabled := MustBool(call.Argument(2)
// 	L.Push(lua.LNumber(a.API.EnableTrigger(name, enabled)))
// 	return 1
// }
// func (a *jsapi) EnableTriggerGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	group := call.Argument(0).String()
// 	enabled := MustBool(call.Argument(2)
// 	L.Push(lua.LNumber(a.API.EnableTriggerGroup(group, enabled)))
// 	return 1
// }

// func (a *jsapi) GetTriggerList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	list := a.API.GetTriggerList()
// 	reuslt := L.NewTable()
// 	for _, v := range list {
// 		reuslt.Append(lua.LString(v))
// 	}
// 	L.Push(reuslt)
// 	return 1
// }
// func (a *jsapi) IsTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.IsTrigger(name)))
// 	return 1
// }

// func (a *jsapi) GetTriggerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	option := call.Argument(1).String()
// 	result, code := a.API.GetTriggerOption(name, option)
// 	if code != api.EOK {
// 		L.Push(lua.LNil)
// 	} else {
// 		switch option {
// 		case "echo_trigger", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
// 			L.Push(lua.LBool(result == world.StringYes))
// 		case "group", "name", "match", "script", "send", "variable":
// 			L.Push(lua.LString(result))
// 		case "send_to", "user", "sequence":
// 			i, _ := strconv.Atoi(result)
// 			L.Push(lua.LNumber(i))
// 		default:
// 			L.Push(lua.LNil)
// 		}
// 	}
// 	return 1
// }
// func (a *jsapi) SetTriggerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	name := call.Argument(0).String()
// 	option := call.Argument(1).String()
// 	var value string
// 	switch option {
// 	case "echo_trigger", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
// 		if MustBool(call.Argument(3) {
// 			value = world.StringYes
// 		} else {
// 			value = ""
// 		}
// 	case "group", "name", "match", "script", "send", "variable":
// 		value = MustString(call.Argument(2))
// 	case "send_to", "user", "sequence":
// 		value = MustString(call.Argument(2))
// 	}
// 	L.Push(lua.LNumber(a.API.SetTriggerOption(name, option, value)))
// 	return 1
// }

// func (a *jsapi) StopEvaluatingTriggers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	a.API.StopEvaluatingTriggers()
// 	return goja.Null()
// }
// func (a *jsapi) ColourNameToRGB(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	v := a.API.ColourNameToRGB(call.Argument(0).String())
// 	L.Push(lua.LString(v))
// 	return 1
// }
// func (a *jsapi) SetSpeedWalkDelay(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	a.API.SetSpeedWalkDelay(L.ToInt(1))
// 	return goja.Null()
// }
// func (a *jsapi) GetSpeedWalkDelay(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	L.Push(lua.LNumber(a.API.SpeedWalkDelay()))
// 	return 1
// }

// func (a *jsapi) NewReadFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 		L.Push(lua.LString(a.API.ReadFile(p, call.Argument(0).String())))
// 		return 1
// 	}
// }
// func (a *jsapi) NewReadLinesAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 		lines := a.API.ReadLines(p, call.Argument(0).String())
// 		t := L.NewTable()
// 		for _, v := range lines {
// 			t.Append(lua.LString(v))
// 		}
// 		L.Push(t)
// 		return 1
// 	}
// }

// func (a *jsapi) SplitNfunc(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	text := call.Argument(0).String()
// 	sep := call.Argument(1).String()
// 	n := L.ToInt(3)
// 	s := a.API.SplitN(text, sep, n)
// 	t := L.NewTable()
// 	for _, v := range s {
// 		t.Append(lua.LString(v))
// 	}
// 	L.Push(t)
// 	return 1
// }

// func (a *jsapi) UTF8Len(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	text := call.Argument(0).String()
// 	L.Push(lua.LNumber(a.API.UTF8Len(text)))
// 	return 1
// }
// func (a *jsapi) UTF8Sub(call goja.FunctionCall, r *goja.Runtime) goja.Value {
// 	text := call.Argument(0).String()
// 	start := L.ToInt(2)
// 	end := L.ToInt(3)
// 	L.Push(lua.LString(a.API.UTF8Sub(text, start, end)))
// 	return 1
// }

func NewAPIModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("worldapi",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			createApi(b).InstallAPIs(plugin)
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
