package luaengine

import (
	"context"
	"errors"
	"modules/app"
	"modules/world"
	"modules/world/bus"
	"modules/world/component/script/api"
	"strconv"
	"strings"
	"time"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	lua "github.com/yuin/gopher-lua"
)

func createApi(b *bus.Bus) *luaapi {
	return &luaapi{
		API: &api.API{
			Bus: b,
		},
	}
}

type luaapi struct {
	API *api.API
}

func (a *luaapi) optional(L *lua.LState, idx int) bool {
	var enabled bool
	vt := L.Get(idx).Type()
	if vt != lua.LTNil {
		enabled = L.ToBool(idx)
	} else {
		enabled = true
	}
	return enabled
}
func (a *luaapi) convertstrings(v lua.LValue) []string {
	var result []string
	switch v.Type() {
	case lua.LTNil:
	case lua.LTTable:
		t := v.(*lua.LTable)
		max := t.MaxN()
		for i := 1; i <= max; i++ {
			result = append(result, lua.LVAsString(t.RawGetInt(i)))
		}
	default:
		panic(errors.New("value must be table"))
	}
	return result
}
func (a *luaapi) getstrings(L *lua.LState) []string {
	t := L.GetTop()
	msg := make([]string, 0, t-1)
	for i := 0; i < t; i++ {
		msg = append(msg, L.Get(i+1).String())
	}
	return msg
}
func (a *luaapi) combine(L *lua.LState) string {
	return strings.Join(a.getstrings(L), " ")
}
func (a *luaapi) InstallAPIs(p herbplugin.Plugin, l *lua.LState) {
	l.SetGlobal("print", l.NewFunction(a.Print))
	l.SetGlobal("Milliseconds", l.NewFunction(a.Milliseconds))

	l.SetGlobal("Note", l.NewFunction(a.Print))

	l.SetGlobal("SendImmediate", l.NewFunction(a.SendImmediate))
	l.SetGlobal("Send", l.NewFunction(a.Send))
	l.SetGlobal("SendNoEcho", l.NewFunction(a.SendNoEcho))
	l.SetGlobal("GetVariable", l.NewFunction(a.GetVariable))
	l.SetGlobal("SetVariable", l.NewFunction(a.SetVariable))
	l.SetGlobal("DeleteVariable", l.NewFunction(a.DeleteVariable))
	l.SetGlobal("GetVariableList", l.NewFunction(a.GetVariableList))
	l.SetGlobal("GetVariableComment", l.NewFunction(a.GetVariable))
	l.SetGlobal("SetVariableComment", l.NewFunction(a.SetVariable))
	l.SetGlobal("Version", l.NewFunction(a.Version))
	l.SetGlobal("Hash", l.NewFunction(a.Hash))
	l.SetGlobal("Base64Encode", l.NewFunction(a.Base64Encode))
	l.SetGlobal("Base64Decode", l.NewFunction(a.Base64Decode))
	l.SetGlobal("Connect", l.NewFunction(a.Connect))
	l.SetGlobal("IsConnected", l.NewFunction(a.IsConnected))
	l.SetGlobal("Disconnect", l.NewFunction(a.Disconnect))
	l.SetGlobal("GetWorldById", l.NewFunction(a.GetWorldById))
	l.SetGlobal("GetWorld", l.NewFunction(a.GetWorld))
	l.SetGlobal("GetWorldID", l.NewFunction(a.GetWorldID))
	l.SetGlobal("GetWorldIdList", l.NewFunction(a.GetWorldIdList))
	l.SetGlobal("GetWorldList", l.NewFunction(a.GetWorldList))
	l.SetGlobal("WorldName", l.NewFunction(a.WorldName))
	l.SetGlobal("WorldAddress", l.NewFunction(a.WorldAddress))
	l.SetGlobal("WorldPort", l.NewFunction(a.WorldPort))
	l.SetGlobal("WorldProxy", l.NewFunction(a.WorldProxy))
	l.SetGlobal("Trim", l.NewFunction(a.Trim))
	l.SetGlobal("GetUniqueNumber", l.NewFunction(a.GetUniqueNumber))
	l.SetGlobal("GetUniqueID", l.NewFunction(a.GetUniqueID))
	l.SetGlobal("CreateGUID", l.NewFunction(a.CreateGUID))
	l.SetGlobal("FlashIcon", l.NewFunction(a.FlashIcon))
	l.SetGlobal("SetStatus", l.NewFunction(a.SetStatus))
	l.SetGlobal("Execute", l.NewFunction(a.Execute))
	l.SetGlobal("DeleteCommandHistory", l.NewFunction(a.DeleteCommandHistory))
	l.SetGlobal("DiscardQueue", l.NewFunction(a.DiscardQueue))
	l.SetGlobal("LockQueue", l.NewFunction(a.LockQueue))
	l.SetGlobal("GetQueue", l.NewFunction(a.GetQueue))
	l.SetGlobal("Queue", l.NewFunction(a.Queue))

	l.SetGlobal("DoAfter", l.NewFunction(a.DoAfter))
	l.SetGlobal("DoAfterNote", l.NewFunction(a.DoAfterNote))
	l.SetGlobal("DoAfterSpeedWalk", l.NewFunction(a.DoAfterSpeedWalk))
	l.SetGlobal("DoAfterSpecial", l.NewFunction(a.DoAfterSpecial))

	l.SetGlobal("DeleteGroup", l.NewFunction(a.DeleteGroup))

	l.SetGlobal("AddTimer", l.NewFunction(a.AddTimer))
	l.SetGlobal("DeleteTimer", l.NewFunction(a.DeleteTimer))
	l.SetGlobal("DeleteTemporaryTimers", l.NewFunction(a.DeleteTemporaryTimers))
	l.SetGlobal("DeleteTimerGroup", l.NewFunction(a.DeleteTimerGroup))
	l.SetGlobal("EnableTimer", l.NewFunction(a.EnableTimer))
	l.SetGlobal("EnableTimerGroup", l.NewFunction(a.EnableTimerGroup))
	l.SetGlobal("GetTimerList", l.NewFunction(a.GetTimerList))
	l.SetGlobal("IsTimer", l.NewFunction(a.IsTimer))
	l.SetGlobal("ResetTimer", l.NewFunction(a.ResetTimer))
	l.SetGlobal("ResetTimers", l.NewFunction(a.ResetTimers))
	l.SetGlobal("GetTimerOption", l.NewFunction(a.GetTimerOption))
	l.SetGlobal("SetTimerOption", l.NewFunction(a.SetTimerOption))
	l.SetGlobal("AddAlias", l.NewFunction(a.AddAlias))
	l.SetGlobal("DeleteAlias", l.NewFunction(a.DeleteAlias))
	l.SetGlobal("DeleteTemporaryAliases", l.NewFunction(a.DeleteTemporaryAliases))
	l.SetGlobal("DeleteAliasGroup", l.NewFunction(a.DeleteAliasGroup))
	l.SetGlobal("EnableAlias", l.NewFunction(a.EnableAlias))
	l.SetGlobal("EnableAliasGroup", l.NewFunction(a.EnableAliasGroup))
	l.SetGlobal("GetAliasList", l.NewFunction(a.GetAliasList))
	l.SetGlobal("IsAlias", l.NewFunction(a.IsAlias))
	l.SetGlobal("GetAliasOption", l.NewFunction(a.GetAliasOption))
	l.SetGlobal("SetAliasOption", l.NewFunction(a.SetAliasOption))

	l.SetGlobal("AddTrigger", l.NewFunction(a.AddTrigger))
	l.SetGlobal("AddTriggerEx", l.NewFunction(a.AddTrigger))
	l.SetGlobal("DeleteTrigger", l.NewFunction(a.DeleteTrigger))
	l.SetGlobal("DeleteTemporaryTriggers", l.NewFunction(a.DeleteTemporaryTriggers))
	l.SetGlobal("DeleteTriggerGroup", l.NewFunction(a.DeleteTriggerGroup))
	l.SetGlobal("EnableTrigger", l.NewFunction(a.EnableTrigger))
	l.SetGlobal("EnableTriggerGroup", l.NewFunction(a.EnableTriggerGroup))
	l.SetGlobal("GetTriggerList", l.NewFunction(a.GetTriggerList))
	l.SetGlobal("IsTrigger", l.NewFunction(a.IsTrigger))
	l.SetGlobal("GetTriggerOption", l.NewFunction(a.GetTriggerOption))
	l.SetGlobal("SetTriggerOption", l.NewFunction(a.SetTriggerOption))
	l.SetGlobal("StopEvaluatingTriggers", l.NewFunction(a.StopEvaluatingTriggers))
	l.SetGlobal("GetTriggerWildcard", l.NewFunction(a.GetTriggerWildcard))

	l.SetGlobal("ColourNameToRGB", l.NewFunction(a.ColourNameToRGB))
	l.SetGlobal("SetSpeedWalkDelay", l.NewFunction(a.SetSpeedWalkDelay))
	l.SetGlobal("GetSpeedWalkDelay", l.NewFunction(a.GetSpeedWalkDelay))

	l.SetGlobal("HasFile", l.NewFunction(a.NewHasFileAPI(p)))
	l.SetGlobal("ReadFile", l.NewFunction(a.NewReadFileAPI(p)))
	l.SetGlobal("ReadLines", l.NewFunction(a.NewReadLinesAPI(p)))

	l.SetGlobal("HasModFile", l.NewFunction(a.NewHasModFileAPI(p)))
	l.SetGlobal("ReadModFile", l.NewFunction(a.NewReadModFileAPI(p)))
	l.SetGlobal("ReadModLines", l.NewFunction(a.NewReadModLinesAPI(p)))
	l.SetGlobal("GetModInfo", l.NewFunction(a.NewGetModInfoAPI(p)))

	l.SetGlobal("MakeHomeFolder", l.NewFunction(a.NewMakeHomeFolderAPI(p)))
	l.SetGlobal("HasHomeFile", l.NewFunction(a.NewHasHomeFileAPI(p)))
	l.SetGlobal("ReadHomeFile", l.NewFunction(a.NewReadHomeFileAPI(p)))
	l.SetGlobal("WriteHomeFile", l.NewFunction(a.NewWriteHomeFileAPI(p)))
	l.SetGlobal("ReadHomeLines", l.NewFunction(a.NewReadHomeLinesAPI(p)))

	l.SetGlobal("SplitN", l.NewFunction(a.SplitNfunc))
	l.SetGlobal("UTF8Len", l.NewFunction(a.UTF8Len))
	l.SetGlobal("UTF8Index", l.NewFunction(a.UTF8Index))

	l.SetGlobal("UTF8Sub", l.NewFunction(a.UTF8Sub))
	l.SetGlobal("FromUTF8", l.NewFunction(a.FromUTF8))
	l.SetGlobal("ToUTF8", l.NewFunction(a.ToUTF8))

	l.SetGlobal("Info", l.NewFunction(a.Info))
	l.SetGlobal("InfoClear", l.NewFunction(a.InfoClear))
	l.SetGlobal("GetAlphaOption", l.NewFunction(a.GetAlphaOption))
	l.SetGlobal("SetAlphaOption", l.NewFunction(a.SetAlphaOption))

	l.SetGlobal("WriteLog", l.NewFunction(a.WriteLog))
	l.SetGlobal("CloseLog", l.NewFunction(a.CloseLog))
	l.SetGlobal("OpenLog", l.NewFunction(a.OpenLog))
	l.SetGlobal("FlushLog", l.NewFunction(a.FlushLog))

	l.SetGlobal("GetLinesInBufferCount", l.NewFunction(a.GetLinesInBufferCount))
	l.SetGlobal("DeleteOutput", l.NewFunction(a.DeleteOutput))
	l.SetGlobal("DeleteLines", l.NewFunction(a.DeleteLines))
	l.SetGlobal("GetLineCount", l.NewFunction(a.GetLineCount))
	l.SetGlobal("GetRecentLines", l.NewFunction(a.GetRecentLines))
	l.SetGlobal("GetLineInfo", l.NewFunction(a.GetLineInfo))
	l.SetGlobal("GetBoldColour", l.NewFunction(a.BoldColour))
	l.SetGlobal("SetBoldColour", l.NewFunction(a.BoldColour))
	l.SetGlobal("GetNormalColour", l.NewFunction(a.NormalColour))
	l.SetGlobal("SetNormalColour", l.NewFunction(a.NormalColour))
	l.SetGlobal("GetStyleInfo", l.NewFunction(a.GetStyleInfo))

	l.SetGlobal("GetInfo", l.NewFunction(a.GetInfo))
	l.SetGlobal("GetTimerInfo", l.NewFunction(a.GetTimerInfo))
	l.SetGlobal("GetTriggerInfo", l.NewFunction(a.GetTriggerInfo))
	l.SetGlobal("GetAliasInfo", l.NewFunction(a.GetAliasInfo))

	l.SetGlobal("Broadcast", l.NewFunction(a.Broadcast))
	l.SetGlobal("Notify", l.NewFunction(a.Notify))
	l.SetGlobal("Request", l.NewFunction(a.Request))

	l.SetGlobal("GetGlobalOption", l.NewFunction(a.GetGlobalOption))

	l.SetGlobal("CheckPermissions", l.NewFunction(a.CheckPermissions))
	l.SetGlobal("RequestPermissions", l.NewFunction(a.RequestPermissions))
	l.SetGlobal("CheckTrustedDomains", l.NewFunction(a.CheckTrustedDomains))
	l.SetGlobal("RequestTrustDomains", l.NewFunction(a.RequestTrustDomains))

	l.SetGlobal("Encrypt", l.NewFunction(a.Encrypt))
	l.SetGlobal("Decrypt", l.NewFunction(a.Decrypt))

	l.SetGlobal("DumpOutput", l.NewFunction(a.DumpOutput))
	l.SetGlobal("ConcatOutput", l.NewFunction(a.ConcatOutput))
	l.SetGlobal("SliceOutput", l.NewFunction(a.SliceOutput))
	l.SetGlobal("OutputToText", l.NewFunction(a.OutputToText))
	l.SetGlobal("FormatOutput", l.NewFunction(a.FormatOutput))
	l.SetGlobal("PrintOutput", l.NewFunction(a.PrintOutput))

	l.SetGlobal("Simulate", l.NewFunction(a.Simulate))
	l.SetGlobal("SimulateOutput", l.NewFunction(a.SimulateOutput))

	l.SetGlobal("DumpTriggers", l.NewFunction(a.DumpTriggers))
	l.SetGlobal("RestoreTriggers", l.NewFunction(a.RestoreTriggers))
	l.SetGlobal("DumpTimers", l.NewFunction(a.DumpTimers))
	l.SetGlobal("RestoreTimers", l.NewFunction(a.RestoreTimers))
	l.SetGlobal("DumpAliases", l.NewFunction(a.DumpAliases))
	l.SetGlobal("RestoreAliases", l.NewFunction(a.RestoreAliases))

	l.SetGlobal("SetHUDSize", l.NewFunction(a.SetHUDSize))
	l.SetGlobal("GetHUDContent", l.NewFunction(a.GetHUDContent))
	l.SetGlobal("GetHUDSize", l.NewFunction(a.GetHUDSize))
	l.SetGlobal("UpdateHUD", l.NewFunction(a.UpdateHUD))
	l.SetGlobal("NewLine", l.NewFunction(a.NewLine))
	l.SetGlobal("NewWord", l.NewFunction(a.NewWord))

	l.SetGlobal("SetPriority", l.NewFunction(a.SetPriority))
	l.SetGlobal("GetPriority", l.NewFunction(a.GetPriority))
	l.SetGlobal("SetSummary", l.NewFunction(a.SetSummary))
	l.SetGlobal("GetSummary", l.NewFunction(a.GetSummary))
	l.SetGlobal("Save", l.NewFunction(a.Save))

	l.SetGlobal("OmitOutput", l.NewFunction(a.OmitOutput))
	l.SetGlobal("PrintSystem", l.NewFunction(a.PrintSystem))

}

func (a *luaapi) Print(L *lua.LState) int {
	a.API.Note(a.combine(L))
	return 0
}
func (a *luaapi) PrintSystem(L *lua.LState) int {
	a.API.PrintSystem(a.combine(L))
	return 0
}
func (a *luaapi) SendImmediate(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.SendImmediate(a.combine(L))))
	return 1
}
func (a *luaapi) Send(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.Send(a.combine(L))))
	return 1
}
func (a *luaapi) Execute(L *lua.LState) int {
	info := L.ToString(1)
	L.Push(lua.LNumber(a.API.Execute(info)))
	return 1
}
func (a *luaapi) SendNoEcho(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.SendNoEcho(a.combine(L))))
	return 1
}
func (a *luaapi) GetVariable(L *lua.LState) int {
	name := L.ToString(1)
	val := a.API.GetVariable(name)
	L.Push(lua.LString(val))
	return 1
}
func (a *luaapi) DeleteVariable(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteVariable(name)))
	return 1
}
func (a *luaapi) SetVariable(L *lua.LState) int {
	name := L.ToString(1)
	value := L.ToString(2)
	L.Push(lua.LNumber(a.API.SetVariable(name, value)))
	return 0
}
func (a *luaapi) GetVariableList(L *lua.LState) int {
	list := a.API.GetVariableList()
	if len(list) == 0 {
		L.Push(lua.LNil)
	} else {
		result := L.NewTable()
		for k, v := range list {
			result.RawSetString(k, lua.LString(v))
		}
		L.Push(result)
	}
	return 1
}
func (a *luaapi) GetVariableComment(L *lua.LState) int {
	name := L.ToString(1)
	val := a.API.GetVariableComment(name)
	L.Push(lua.LString(val))
	return 1
}
func (a *luaapi) SetVariableComment(L *lua.LState) int {
	name := L.ToString(1)
	value := L.ToString(2)
	L.Push(lua.LNumber(a.API.SetVariableComment(name, value)))
	return 0
}
func (a *luaapi) Version(L *lua.LState) int {
	L.Push(lua.LString(a.API.Version()))
	return 1
}
func (a *luaapi) Hash(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LString(a.API.Hash(name)))
	return 1
}
func (a *luaapi) Base64Encode(L *lua.LState) int {
	var ok bool
	src := L.ToString(1)
	ml := L.Get(2)
	if ml.Type() == lua.LTBool {
		ok = bool(ml.(lua.LBool))
	}
	L.Push(lua.LString(a.API.Base64Encode(src, ok)))
	return 1
}
func (a *luaapi) Base64Decode(L *lua.LState) int {
	src := L.ToString(1)
	result := a.API.Base64Decode(src)
	if result == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(*result))
	}
	return 1
}
func (a *luaapi) Connect(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.Connect()))
	return 1
}
func (a *luaapi) IsConnected(L *lua.LState) int {
	L.Push(lua.LBool(a.API.IsConnected()))
	return 1
}
func (a *luaapi) Disconnect(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.Disconnect()))
	return 1
}

func (a *luaapi) GetWorldById(L *lua.LState) int {
	L.Push(lua.LNil)
	return 1
}

func (a *luaapi) GetWorld(L *lua.LState) int {
	L.Push(lua.LNil)
	return 1
}

func (a *luaapi) GetWorldID(L *lua.LState) int {
	L.Push(lua.LString(a.API.GetWorldID()))
	return 1
}
func (a *luaapi) GetWorldIdList(L *lua.LState) int {
	L.Push(L.NewTable())
	return 1
}
func (a *luaapi) GetWorldList(L *lua.LState) int {
	L.Push(L.NewTable())
	return 1
}
func (a *luaapi) WorldName(L *lua.LState) int {
	L.Push(lua.LString(a.API.WorldName()))
	return 1
}
func (a *luaapi) WorldAddress(L *lua.LState) int {
	L.Push(lua.LString(a.API.WorldAddress()))
	return 1
}
func (a *luaapi) WorldPort(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.WorldPort()))
	return 1
}
func (a *luaapi) WorldProxy(L *lua.LState) int {
	L.Push(lua.LString(a.API.WorldProxy()))
	return 1
}
func (a *luaapi) Trim(L *lua.LState) int {
	src := L.ToString(1)
	L.Push(lua.LString(a.API.Trim(src)))
	return 1
}
func (a *luaapi) GetUniqueNumber(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.GetUniqueNumber()))
	return 1
}
func (a *luaapi) GetUniqueID(L *lua.LState) int {
	L.Push(lua.LString(a.API.GetUniqueID()))
	return 1
}
func (a *luaapi) CreateGUID(L *lua.LState) int {
	L.Push(lua.LString(a.API.CreateGUID()))
	return 1
}
func (a *luaapi) FlashIcon(L *lua.LState) int {
	a.API.FlashIcon()
	return 0
}
func (a *luaapi) SetStatus(L *lua.LState) int {
	a.API.SetStatus(a.combine(L))
	return 0
}
func (a *luaapi) DeleteCommandHistory(L *lua.LState) int {
	a.API.DeleteCommandHistory()
	return 0
}
func (a *luaapi) DiscardQueue(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.DiscardQueue(L.ToBool(1))))
	return 1
}
func (a *luaapi) LockQueue(L *lua.LState) int {
	a.API.LockQueue()
	return 0
}
func (a *luaapi) GetQueue(L *lua.LState) int {
	cmds := a.API.GetQueue()
	t := L.NewTable()
	for k := range cmds {
		t.Append(lua.LString(cmds[k]))
	}
	L.Push(t)
	return 1
}

func (a *luaapi) DoAfter(L *lua.LState) int {
	seconds := float64(L.ToNumber(1))
	send := L.ToString(2)
	L.Push(lua.LNumber(a.API.DoAfter(seconds, send)))
	return 1
}
func (a *luaapi) DoAfterNote(L *lua.LState) int {
	seconds := float64(L.ToNumber(1))
	send := L.ToString(2)
	L.Push(lua.LNumber(a.API.DoAfterNote(seconds, send)))
	return 1
}
func (a *luaapi) DoAfterSpeedWalk(L *lua.LState) int {
	seconds := float64(L.ToNumber(1))
	send := L.ToString(2)
	L.Push(lua.LNumber(a.API.DoAfterSpeedWalk(seconds, send)))
	return 1
}
func (a *luaapi) DoAfterSpecial(L *lua.LState) int {
	seconds := float64(L.ToNumber(1))
	send := L.ToString(2)
	sendto := int(L.ToNumber(3))
	L.Push(lua.LNumber(a.API.DoAfterSpecial(seconds, send, sendto)))
	return 1
}

func (a *luaapi) DeleteGroup(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteGroup(name)))
	return 1
}

func (a *luaapi) AddTimer(L *lua.LState) int {
	name := L.ToString(1)
	hour := int(L.ToNumber(2))
	min := int(L.ToNumber(3))
	seconds := float64(L.ToNumber(4))
	send := L.ToString(5)
	flags := int(L.ToNumber(6))
	script := L.ToString(7)
	L.Push(lua.LNumber(a.API.AddTimer(name, hour, min, seconds, send, flags, script)))
	return 1
}
func (a *luaapi) DeleteTimer(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteTimer(name)))
	return 1
}
func (a *luaapi) DeleteTemporaryTimers(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.DeleteTemporaryTimers()))
	return 1

}
func (a *luaapi) DeleteTimerGroup(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteTimerGroup(name)))
	return 1
}

func (a *luaapi) EnableTimer(L *lua.LState) int {
	name := L.ToString(1)
	var enabled bool
	enabled = a.optional(L, 2)
	L.Push(lua.LNumber(a.API.EnableTimer(name, enabled)))
	return 1
}
func (a *luaapi) EnableTimerGroup(L *lua.LState) int {
	group := L.ToString(1)
	var enabled bool
	enabled = a.optional(L, 2)
	L.Push(lua.LNumber(a.API.EnableTimerGroup(group, enabled)))
	return 1
}

func (a *luaapi) GetTimerList(L *lua.LState) int {
	list := a.API.GetTimerList()
	reuslt := L.NewTable()
	for _, v := range list {
		reuslt.Append(lua.LString(v))
	}
	L.Push(reuslt)
	return 1
}
func (a *luaapi) IsTimer(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.IsTimer(name)))
	return 1
}

func (a *luaapi) ResetTimer(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.ResetTimer(name)))
	return 1
}

func (a *luaapi) ResetTimers(L *lua.LState) int {
	a.API.ResetTimers()
	return 0
}

func (a *luaapi) GetTimerOption(L *lua.LState) int {
	name := L.ToString(1)
	option := L.ToString(2)
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		L.Push(lua.LNil)
	} else {
		switch option {
		case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
			L.Push(lua.LBool(result == world.StringYes))
		case "group", "name", "script", "send", "variable":
			L.Push(lua.LString(result))
		case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "send_to", "user":
			i, _ := strconv.Atoi(result)
			L.Push(lua.LNumber(i))
		case "second":
			i, _ := strconv.ParseFloat(result, 64)
			L.Push(lua.LNumber(i))

		default:
			L.Push(lua.LNil)
		}
	}
	return 1
}
func (a *luaapi) SetTimerOption(L *lua.LState) int {
	name := L.ToString(1)
	option := L.ToString(2)
	var value string
	switch option {
	case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
		if L.ToBool(3) {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "script", "send", "variable":
		value = L.ToString(3)
	case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "second", "send_to", "user":
		value = L.ToString(3)
	}
	L.Push(lua.LNumber(a.API.SetTimerOption(name, option, value)))
	return 1
}

func (a *luaapi) AddAlias(L *lua.LState) int {
	name := L.ToString(1)
	match := L.ToString(2)
	send := L.ToString(3)
	flags := int(L.ToNumber(4))
	script := L.ToString(5)
	L.Push(lua.LNumber(a.API.AddAlias(name, match, send, flags, script)))
	return 1
}
func (a *luaapi) DeleteAlias(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteAlias(name)))
	return 1
}
func (a *luaapi) DeleteTemporaryAliases(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.DeleteTemporaryAliases()))
	return 1

}
func (a *luaapi) DeleteAliasGroup(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteAliasGroup(name)))
	return 1
}

func (a *luaapi) EnableAlias(L *lua.LState) int {
	name := L.ToString(1)
	var enabled bool
	enabled = a.optional(L, 2)
	L.Push(lua.LNumber(a.API.EnableAlias(name, enabled)))
	return 1
}
func (a *luaapi) EnableAliasGroup(L *lua.LState) int {
	group := L.ToString(1)
	var enabled bool
	enabled = a.optional(L, 2)
	L.Push(lua.LNumber(a.API.EnableAliasGroup(group, enabled)))
	return 1
}

func (a *luaapi) GetAliasList(L *lua.LState) int {
	list := a.API.GetAliasList()
	reuslt := L.NewTable()
	for _, v := range list {
		reuslt.Append(lua.LString(v))
	}
	L.Push(reuslt)
	return 1
}
func (a *luaapi) IsAlias(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.IsAlias(name)))
	return 1
}

func (a *luaapi) GetAliasOption(L *lua.LState) int {
	name := L.ToString(1)
	option := L.ToString(2)
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		L.Push(lua.LNil)
	} else {
		switch option {
		case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			L.Push(lua.LBool(result == world.StringYes))
		case "group", "name", "match", "script", "send", "variable":
			L.Push(lua.LString(result))
		case "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			L.Push(lua.LNumber(i))
		default:
			L.Push(lua.LNil)
		}
	}
	return 1
}
func (a *luaapi) SetAliasOption(L *lua.LState) int {
	name := L.ToString(1)
	option := L.ToString(2)
	var value string
	switch option {
	case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
		if L.ToBool(3) {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "match", "script", "send", "variable":
		value = L.ToString(3)
	case "send_to", "user", "sequence":
		value = L.ToString(3)
	}
	L.Push(lua.LNumber(a.API.SetTimerOption(name, option, value)))
	return 1
}

func (a *luaapi) AddTrigger(L *lua.LState) int {
	name := L.ToString(1)
	match := L.ToString(2)
	send := L.ToString(3)
	flags := int(L.ToNumber(4))
	color := int(L.ToNumber(5))
	wildcard := int(L.ToNumber(6))
	sound := L.ToString(7)
	script := L.ToString(8)
	L.Push(lua.LNumber(a.API.AddTrigger(name, match, send, flags, color, wildcard, sound, script)))
	return 1
}
func (a *luaapi) AddTriggerEx(L *lua.LState) int {
	name := L.ToString(1)
	match := L.ToString(2)
	send := L.ToString(3)
	flags := int(L.ToNumber(4))
	color := int(L.ToNumber(5))
	wildcard := int(L.ToNumber(6))
	sound := L.ToString(7)
	script := L.ToString(8)
	sendto := int(L.ToNumber(9))
	sequence := int(L.ToNumber(10))
	L.Push(lua.LNumber(a.API.AddTriggerEx(name, match, send, flags, color, wildcard, sound, script, sendto, sequence)))
	return 1
}
func (a *luaapi) DeleteTrigger(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteTrigger(name)))
	return 1
}
func (a *luaapi) DeleteTemporaryTriggers(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.DeleteTemporaryTimers()))
	return 1

}
func (a *luaapi) DeleteTriggerGroup(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.DeleteTriggerGroup(name)))
	return 1
}

func (a *luaapi) EnableTrigger(L *lua.LState) int {
	name := L.ToString(1)
	var enabled bool
	enabled = a.optional(L, 2)
	L.Push(lua.LNumber(a.API.EnableTrigger(name, enabled)))
	return 1
}
func (a *luaapi) EnableTriggerGroup(L *lua.LState) int {
	group := L.ToString(1)
	var enabled bool
	enabled = a.optional(L, 2)
	L.Push(lua.LNumber(a.API.EnableTriggerGroup(group, enabled)))
	return 1
}

func (a *luaapi) GetTriggerList(L *lua.LState) int {
	list := a.API.GetTriggerList()
	reuslt := L.NewTable()
	for _, v := range list {
		reuslt.Append(lua.LString(v))
	}
	L.Push(reuslt)
	return 1
}
func (a *luaapi) IsTrigger(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LNumber(a.API.IsTrigger(name)))
	return 1
}

func (a *luaapi) GetTriggerOption(L *lua.LState) int {
	name := L.ToString(1)
	option := L.ToString(2)
	result, code := a.API.GetTriggerOption(name, option)
	if code != api.EOK {
		L.Push(lua.LNil)
	} else {
		switch option {
		case "echo_trigger", "multi_line", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			L.Push(lua.LBool(result == world.StringYes))
		case "group", "name", "match", "script", "send", "variable":
			L.Push(lua.LString(result))
		case "lines_to_match", "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			L.Push(lua.LNumber(i))
		default:
			L.Push(lua.LNil)
		}
	}
	return 1
}
func (a *luaapi) SetTriggerOption(L *lua.LState) int {
	name := L.ToString(1)
	option := L.ToString(2)
	var value string
	switch option {
	case "echo_trigger", "multi_line", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
		if L.ToBool(3) {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "match", "script", "send", "variable":
		value = L.ToString(3)
	case "lines_to_match", "send_to", "user", "sequence":
		value = L.ToString(3)
	}
	L.Push(lua.LNumber(a.API.SetTriggerOption(name, option, value)))
	return 1
}

func (a *luaapi) StopEvaluatingTriggers(L *lua.LState) int {
	a.API.StopEvaluatingTriggers()
	return 0
}
func (a *luaapi) GetTriggerWildcard(L *lua.LState) int {
	v := a.API.GetTriggerWildcard(L.ToString(1), L.ToString(2))
	if v == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(*v))
	}
	return 1
}
func (a *luaapi) ColourNameToRGB(L *lua.LState) int {
	v := a.API.ColourNameToRGB(L.ToString(1))
	L.Push(lua.LNumber(v))
	return 1
}
func (a *luaapi) SetSpeedWalkDelay(L *lua.LState) int {
	a.API.SetSpeedWalkDelay(L.ToInt(1))
	return 0
}
func (a *luaapi) GetSpeedWalkDelay(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.SpeedWalkDelay()))
	return 1
}
func (a *luaapi) Queue(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.Queue(L.ToString(1), a.optional(L, 2))))
	return 1
}

func (a *luaapi) NewHasFileAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.Push(lua.LBool(a.API.HasFile(p, L.ToString(1))))
		return 1
	}
}

func (a *luaapi) NewReadFileAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.Push(lua.LString(a.API.ReadFile(p, L.ToString(1))))
		return 1
	}
}
func (a *luaapi) NewReadLinesAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		lines := a.API.ReadLines(p, L.ToString(1))
		t := L.NewTable()
		for _, v := range lines {
			t.Append(lua.LString(v))
		}
		L.Push(t)
		return 1
	}
}

func (a *luaapi) NewGetModInfoAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		mod := a.API.GetModInfo(p)
		result := L.NewTable()
		result.RawSetString("Enabled", lua.LBool(mod.Enabled))
		result.RawSetString("Exists", lua.LBool(mod.Exists))
		filelist := L.NewTable()
		for _, name := range mod.FileList {
			filelist.Append(lua.LString(name))
		}
		result.RawSetString("FileList", filelist)

		folderlist := L.NewTable()
		for _, name := range mod.FolderList {
			folderlist.Append(lua.LString(name))
		}
		result.RawSetString("FolderList", folderlist)
		L.Push(result)
		return 1
	}
}
func (a *luaapi) NewReadModFileAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.Push(lua.LString(a.API.ReadModFile(p, L.ToString(1))))
		return 1
	}
}
func (a *luaapi) NewReadHomeFileAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.Push(lua.LString(a.API.ReadHomeFile(p, L.ToString(1))))
		return 1
	}
}
func (a *luaapi) NewReadModLinesAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		lines := a.API.ReadModLines(p, L.ToString(1))
		t := L.NewTable()
		for _, v := range lines {
			t.Append(lua.LString(v))
		}
		L.Push(t)
		return 1
	}
}
func (a *luaapi) NewReadHomeLinesAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		lines := a.API.ReadHomeLines(p, L.ToString(1))
		t := L.NewTable()
		for _, v := range lines {
			t.Append(lua.LString(v))
		}
		L.Push(t)
		return 1
	}
}
func (a *luaapi) NewWriteHomeFileAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		a.API.WriteHomeFile(p, L.ToString(1), []byte(L.ToString(2)))
		return 0
	}
}
func (a *luaapi) NewHasModFileAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.Push(lua.LBool(a.API.HasModFile(p, L.ToString(1))))
		return 1
	}
}

func (a *luaapi) NewMakeHomeFolderAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.Push(lua.LBool(a.API.MakeHomeFolder(p, L.ToString(1))))
		return 1
	}
}

func (a *luaapi) NewHasHomeFileAPI(p herbplugin.Plugin) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		L.Push(lua.LBool(a.API.HasHomeFile(p, L.ToString(1))))
		return 1
	}
}
func (a *luaapi) SplitNfunc(L *lua.LState) int {
	text := L.ToString(1)
	sep := L.ToString(2)
	n := L.ToInt(3)
	s := a.API.SplitN(text, sep, n)
	t := L.NewTable()
	for _, v := range s {
		t.Append(lua.LString(v))
	}
	L.Push(t)
	return 1
}

func (a *luaapi) UTF8Len(L *lua.LState) int {
	text := L.ToString(1)
	L.Push(lua.LNumber(a.API.UTF8Len(text)))
	return 1
}
func (a *luaapi) UTF8Index(L *lua.LState) int {
	text := L.ToString(1)
	sub := L.ToString(2)
	L.Push(lua.LNumber(a.API.UTF8Index(text, sub)))
	return 1
}
func (a *luaapi) UTF8Sub(L *lua.LState) int {
	text := L.ToString(1)
	start := L.ToInt(2)
	end := L.ToInt(3)
	L.Push(lua.LString(a.API.UTF8Sub(text, start, end)))
	return 1
}
func (a *luaapi) ToUTF8(L *lua.LState) int {
	code := L.ToString(1)
	text := L.ToString(2)
	result := a.API.ToUTF8(code, text)
	if result != nil {
		L.Push(lua.LString(*result))
	} else {
		L.Push(lua.LNil)
	}
	return 1
}
func (a *luaapi) FromUTF8(L *lua.LState) int {
	code := L.ToString(1)
	text := L.ToString(2)
	result := a.API.FromUTF8(code, text)
	if result != nil {
		L.Push(lua.LString(*result))
	} else {
		L.Push(lua.LNil)
	}
	return 1
}
func (a *luaapi) Info(L *lua.LState) int {
	a.API.Info(L.ToString(1))
	return 0
}
func (a *luaapi) InfoClear(L *lua.LState) int {
	a.API.InfoClear()
	return 0
}
func (a *luaapi) GetAlphaOption(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LString(a.API.GetAlphaOption(name)))
	return 1
}
func (a *luaapi) SetAlphaOption(L *lua.LState) int {
	name := L.ToString(1)
	value := L.ToString(2)
	L.Push(lua.LNumber(a.API.SetAlphaOption(name, value)))
	return 1
}

func (a *luaapi) WriteLog(L *lua.LState) int {
	message := L.ToString(1)
	L.Push(lua.LNumber(a.API.WriteLog(message)))
	return 1
}
func (a *luaapi) CloseLog(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.CloseLog()))
	return 1
}
func (a *luaapi) OpenLog(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.OpenLog()))
	return 1
}
func (a *luaapi) FlushLog(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.FlushLog()))
	return 1
}

func (a *luaapi) GetLinesInBufferCount(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.GetLinesInBufferCount()))
	return 1
}
func (a *luaapi) DeleteOutput(L *lua.LState) int {
	a.API.DeleteOutput()
	return 0
}
func (a *luaapi) DeleteLines(L *lua.LState) int {
	a.API.DeleteOutput()
	return 0
}
func (a *luaapi) GetLineCount(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.GetLineCount()))
	return 1
}
func (a *luaapi) GetRecentLines(L *lua.LState) int {
	L.Push(lua.LString(a.API.GetRecentLines(L.ToInt(1))))
	return 1
}
func (a *luaapi) getLineInfo(L *lua.LState) int {
	val, ok := a.API.GetLineInfo(L.ToInt(1), L.ToInt(2))
	if !ok {
		L.Push(lua.LNil)
		return 1
	}
	switch L.ToInt(2) {
	case 1:
		L.Push(lua.LString(val))
	case 2:
		L.Push(lua.LNumber(world.FromStringInt(val)))
	case 3:
		L.Push(lua.LNumber(world.FromStringInt(val)))
	case 4:
		L.Push(lua.LBool(world.FromStringBool(val)))
	case 5:
		L.Push(lua.LBool(world.FromStringBool(val)))
	case 6:
		L.Push(lua.LBool(world.FromStringBool(val)))
	case 7:
		L.Push(lua.LBool(world.FromStringBool(val)))
	case 8:
		L.Push(lua.LBool(world.FromStringBool(val)))
	case 9:
		L.Push(lua.LNumber(world.FromStringInt(val)))
	case 11:
		L.Push(lua.LNumber(world.FromStringInt(val)))
	default:
		L.Push(lua.LNil)
	}
	return 1
}
func (a *luaapi) getLineTable(L *lua.LState) int {
	line := a.API.Bus.GetLine(L.ToInt(1))
	if line == nil {
		L.Push(lua.LNil)
		return 1
	}
	t := L.NewTable()
	t.RawSetString("text", lua.LString(line.Plain()))
	t.RawSetString("length", lua.LNumber(len(line.Plain())))
	t.RawSetString("newline", lua.LBool(line.IsNewline()))
	t.RawSetString("note", lua.LBool(line.Type == world.LineTypePrint))
	t.RawSetString("user", lua.LBool(line.Type == world.LineTypeEcho))
	t.RawSetString("log", lua.LBool(!line.OmitFromLog))
	t.RawSetString("time", lua.LNumber(line.Time))
	t.RawSetString("timestr", lua.LString(app.Time.Datetime(time.Unix(line.Time, 0))))
	t.RawSetString("line", lua.LString(line.ID))
	t.RawSetString("styles", lua.LNumber(len(line.Words)))

	L.Push(t)
	return 1

}
func (a *luaapi) GetLineInfo(L *lua.LState) int {
	if L.ToInt(2) != 0 {
		return a.getLineInfo(L)
	}
	return a.getLineTable(L)
}
func (a *luaapi) BoldColour(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.BoldColour(L.ToInt(1))))
	return 1

}
func (a *luaapi) NormalColour(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.NormalColour(L.ToInt(1))))
	return 1
}
func (a *luaapi) getStyleInfo(ln int, st int, it int) lua.LValue {
	val, ok := a.API.GetStyleInfo(ln, st, it)
	if !ok {
		return lua.LNil
	}
	switch it {
	case 1:
		return lua.LString(val)
	case 2:
		return lua.LNumber(world.FromStringInt(val))
	case 3:
		return lua.LNumber(world.FromStringInt(val))
	case 8:
		return lua.LBool(world.FromStringBool(val))
	case 9:
		return lua.LBool(world.FromStringBool(val))
	case 10:
		return lua.LBool(world.FromStringBool(val))
	case 11:
		return lua.LBool(world.FromStringBool(val))
	case 14:
		return lua.LNumber(world.FromStringInt(val))
	case 15:
		return lua.LNumber(world.FromStringInt(val))
	}
	return lua.LNil
}
func (a *luaapi) getStyleInfoTable(L *lua.LState, ln int, st int) lua.LValue {
	line := a.API.Bus.GetLine(ln)
	if line.IsEmpty() {
		return lua.LNil
	}
	if st <= 0 || st > len(line.Words) {
		return lua.LNil
	}
	word := line.Words[st-1]
	t := L.NewTable()
	t.RawSetString("text", lua.LString(word.Text))
	t.RawSetString("length", lua.LNumber(len(word.Text)))
	t.RawSetString("column", lua.LNumber(line.GetWordStartColumn(st)))
	t.RawSetString("bold", lua.LBool(word.Bold))
	t.RawSetString("ul", lua.LBool(word.Underlined))
	t.RawSetString("blink", lua.LBool(word.Blinking))
	t.RawSetString("inverse", lua.LBool(word.Inverse))
	t.RawSetString("foreground", lua.LNumber(word.GetColorRGB()))
	t.RawSetString("background", lua.LNumber(word.GetBGColorRGB()))
	return t
}
func (a *luaapi) GetStyleInfo(L *lua.LState) int {
	ln := L.ToInt(1)
	sn := L.ToInt(2)
	if sn < 0 {
		L.Push(lua.LNil)
		return 1
	}
	tp := L.ToInt(3)
	if sn == 0 {
		line := a.API.Bus.GetLine(ln)
		if line == nil {
			L.Push(lua.LNil)
			return 1
		}
		v := L.NewTable()
		for k := range line.Words {
			if tp == 0 {
				v.Append(a.getStyleInfoTable(L, ln, k+1))
			} else {
				v.Append(a.getStyleInfo(ln, k+1, tp))
			}
		}
		L.Push(v)
		return 1
	} else if tp == 0 {
		L.Push(a.getStyleInfoTable(L, ln, sn))
		return 1
	}
	L.Push(a.getStyleInfo(ln, sn, tp))
	return 1
}

func (a *luaapi) GetInfo(L *lua.LState) int {
	L.Push(lua.LString(a.API.GetInfo(L.ToInt(1))))
	return 1
}
func (a *luaapi) GetTimerInfo(L *lua.LState) int {
	v, ok := a.API.GetTimerInfo(L.ToString(1), L.ToInt(2))
	if ok != api.EOK {
		L.Push(lua.LNil)
		return 1
	}
	switch L.ToInt(2) {
	case 1:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 2:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 3:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 4:
		L.Push(lua.LString(v))
	case 5:
		L.Push(lua.LString(v))
	case 6:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 7:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 8:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 14:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 19:
		L.Push(lua.LString(v))
	case 20:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 21:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 22:
		L.Push(lua.LString(v))
	case 23:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 24:
		L.Push(lua.LBool(world.FromStringBool(v)))
	default:
		L.Push(lua.LNil)
	}
	return 1
}

func (a *luaapi) GetTriggerInfo(L *lua.LState) int {
	v, ok := a.API.GetTriggerInfo(L.ToString(1), L.ToInt(2))
	if ok != api.EOK {
		L.Push(lua.LNil)
		return 1
	}
	switch L.ToInt(2) {
	case 1:
		L.Push(lua.LString(v))
	case 2:
		L.Push(lua.LString(v))
	case 3:
		L.Push(lua.LString(v))
	case 4:
		L.Push(lua.LString(v))
	case 5:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 6:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 7:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 8:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 9:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 10:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 11:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 13:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 15:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 16:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 23:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 25:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 26:
		L.Push(lua.LString(v))
	case 27:
		L.Push(lua.LString(v))
	case 28:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 31:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 36:
		L.Push(lua.LBool(world.FromStringBool(v)))
	default:
		L.Push(lua.LNil)

	}
	return 1
}

func (a *luaapi) GetAliasInfo(L *lua.LState) int {
	v, ok := a.API.GetAliasInfo(L.ToString(1), L.ToInt(2))
	if ok != api.EOK {
		L.Push(lua.LNil)
		return 1
	}
	switch L.ToInt(2) {
	case 1:
		L.Push(lua.LString(v))
	case 2:
		L.Push(lua.LString(v))
	case 3:
		L.Push(lua.LString(v))
	case 4:
		L.Push(lua.LString(v))
	case 5:
		L.Push(lua.LString(v))
	case 6:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 7:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 8:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 9:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 14:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 16:
		L.Push(lua.LString(v))
	case 17:
		L.Push(lua.LString(v))
	case 18:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 19:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 20:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 22:
		L.Push(lua.LBool(world.FromStringBool(v)))
	case 23:
		L.Push(lua.LNumber(world.FromStringInt(v)))
	case 29:
		L.Push(lua.LBool(world.FromStringBool(v)))

	}
	return 1
}

func (a *luaapi) Broadcast(L *lua.LState) int {
	a.API.Broadcast(L.ToString(1), L.ToBool(2))
	return 0
}
func (a *luaapi) Notify(L *lua.LState) int {
	var link *string
	if L.Get(3) == lua.LNil {
		link = nil
	} else {
		data := L.ToString(3)
		link = &data
	}
	a.API.Notify(L.ToString(1), L.ToString(2), link)
	return 0
}
func (a *luaapi) Request(L *lua.LState) int {
	id := a.API.Request(L.ToString(1), L.ToString(2))
	L.Push(lua.LString(id))
	return 1
}
func (a *luaapi) GetGlobalOption(L *lua.LState) int {
	result := a.API.GetGlobalOption(L.ToString(1))
	switch L.ToString(1) {
	default:
		switch result {
		case "0":
			L.Push(lua.LNumber(0))
		case "1":
			L.Push(lua.LNumber(1))
		default:
			L.Push(lua.LString(result))
		}
	}
	return 1
}
func (a *luaapi) CheckTrustedDomains(L *lua.LState) int {
	items := a.convertstrings(L.Get((1)))
	L.Push(lua.LBool(a.API.CheckTrustedDomains(items)))
	return 1
}

func (a *luaapi) RequestPermissions(L *lua.LState) int {
	items := a.convertstrings(L.Get((1)))
	var reason string
	vreason := L.Get(2)
	if vreason.Type() != lua.LTNil {
		reason = vreason.String()
	}
	var script string
	vscript := L.Get(3)
	if vscript.Type() != lua.LTNil {
		script = vscript.String()
	}
	a.API.RequestPermissions(items, reason, script)
	return 0
}

func (a *luaapi) CheckPermissions(L *lua.LState) int {
	items := a.convertstrings(L.Get((1)))
	L.Push(lua.LBool(a.API.CheckPermissions(items)))
	return 1
}

func (a *luaapi) RequestTrustDomains(L *lua.LState) int {
	items := a.convertstrings(L.Get((1)))
	var reason string
	vreason := L.Get(2)
	if vreason.Type() != lua.LTNil {
		reason = vreason.String()
	}
	var script string
	vscript := L.Get(3)
	if vscript.Type() != lua.LTNil {
		script = vscript.String()
	}
	a.API.RequestTrustDomains(items, reason, script)
	return 0
}

func (a *luaapi) Encrypt(L *lua.LState) int {
	data := L.ToString(1)
	key := L.ToString(2)
	result := a.API.Encrypt(data, key)
	if result == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(*result))
	}
	return 1
}
func (a *luaapi) Decrypt(L *lua.LState) int {
	data := L.ToString(1)
	key := L.ToString(2)
	result := a.API.Decrypt(data, key)
	if result == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(*result))
	}
	return 1
}

func (a *luaapi) DumpOutput(L *lua.LState) int {
	length := L.ToInt(1)
	offset := L.ToInt(2)
	L.Push(lua.LString(a.API.DumpOutput(length, offset)))
	return 1
}

func (a *luaapi) ConcatOutput(L *lua.LState) int {
	output1 := L.ToString(1)
	output2 := L.ToString(2)
	L.Push(lua.LString(a.API.ConcatOutput(output1, output2)))
	return 1
}
func (a *luaapi) SliceOutput(L *lua.LState) int {
	output := L.ToString(1)
	start := L.ToInt(2)
	end := L.ToInt(3)
	L.Push(lua.LString(a.API.SliceOutput(output, start, end)))
	return 1
}
func (a *luaapi) OutputToText(L *lua.LState) int {
	output := L.ToString(1)
	L.Push(lua.LString(a.API.OutputToText(output)))
	return 1
}
func (a *luaapi) FormatOutput(L *lua.LState) int {
	output := L.ToString(1)
	L.Push(lua.LString(a.API.FormatOutput(output)))
	return 1
}
func (a *luaapi) PrintOutput(L *lua.LState) int {
	output := L.ToString(1)
	L.Push(lua.LString(a.API.PrintOutput(output)))
	return 1
}
func (a *luaapi) Simulate(L *lua.LState) int {
	text := L.ToString(1)
	a.API.Simulate(text)
	return 0
}
func (a *luaapi) SimulateOutput(L *lua.LState) int {
	output := L.ToString(1)
	a.API.SimulateOutput(output)
	return 0
}
func (a *luaapi) DumpTriggers(L *lua.LState) int {
	byUser := L.ToBool(1)
	L.Push(lua.LString(a.API.DumpTriggers(byUser)))
	return 1
}
func (a *luaapi) RestoreTriggers(L *lua.LState) int {
	data := L.ToString(1)
	byUser := L.ToBool(2)
	a.API.RestoreTriggers(data, byUser)
	return 0
}
func (a *luaapi) DumpTimers(L *lua.LState) int {
	byUser := L.ToBool(1)
	L.Push(lua.LString(a.API.DumpTimers(byUser)))
	return 1
}
func (a *luaapi) RestoreTimers(L *lua.LState) int {
	data := L.ToString(1)
	byUser := L.ToBool(2)
	a.API.RestoreTimers(data, byUser)
	return 0
}
func (a *luaapi) DumpAliases(L *lua.LState) int {
	byUser := L.ToBool(1)
	L.Push(lua.LString(a.API.DumpAliases(byUser)))
	return 1
}
func (a *luaapi) RestoreAliases(L *lua.LState) int {
	data := L.ToString(1)
	byUser := L.ToBool(2)
	a.API.RestoreAliases(data, byUser)
	return 0
}
func (a *luaapi) SetHUDSize(L *lua.LState) int {
	size := L.ToInt(1)
	a.API.SetHUDSize(size)
	return 0
}
func (a *luaapi) GetHUDContent(L *lua.LState) int {
	content := a.API.GetHUDContent()
	L.Push(lua.LString(content))
	return 1
}
func (a *luaapi) GetHUDSize(L *lua.LState) int {
	size := a.API.GetHUDSize()
	L.Push(lua.LNumber(size))
	return 1
}
func (a *luaapi) UpdateHUD(L *lua.LState) int {
	start := L.ToInt(1)
	content := L.ToString(2)
	result := a.API.UpdateHUD(start, content)
	L.Push(lua.LBool(result))
	return 1
}
func (a *luaapi) NewLine(L *lua.LState) int {
	L.Push(lua.LString(a.API.NewLine()))
	return 1
}
func (a *luaapi) NewWord(L *lua.LState) int {
	data := L.ToString(1)
	L.Push(lua.LString(a.API.NewWord(data)))
	return 1
}

func (a *luaapi) SetPriority(L *lua.LState) int {
	value := L.ToInt(1)
	a.API.SetPriority(value)
	return 0
}
func (a *luaapi) GetPriority(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.GetPriority()))
	return 1
}
func (a *luaapi) SetSummary(L *lua.LState) int {
	content := L.ToString(1)
	a.API.SetSummary(content)
	return 0
}
func (a *luaapi) GetSummary(L *lua.LState) int {
	L.Push(lua.LString(a.API.GetSummary()))
	return 1
}

func (a *luaapi) Save(L *lua.LState) int {
	L.Push(lua.LBool(a.API.Save()))
	return 1
}
func (a *luaapi) Milliseconds(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.Milliseconds()))
	return 1
}
func (a *luaapi) OmitOutput(L *lua.LState) int {
	a.API.OmitOutput()
	return 0
}
func NewAPIModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("worldapi",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			createApi(b).InstallAPIs(plugin, l)
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
