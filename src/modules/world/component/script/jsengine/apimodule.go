package jsengine

import (
	"context"
	"modules/world"
	"modules/world/bus"
	"modules/world/component/script/api"
	"strconv"
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
type World struct {
	API     *api.API
	Runtime *goja.Runtime
	Map     map[string]goja.Value
}

func CreateWorld(a *api.API, r *goja.Runtime) *World {
	return &World{
		API:     a,
		Runtime: r,
		Map:     map[string]goja.Value{},
	}
}

// Get a property value for the key. May return nil if the property does not exist.
func (w *World) Get(key string) goja.Value {

	lkey := strings.ToLower(key)
	switch lkey {
	case "speedwalkdelay":
		return w.Runtime.ToValue(w.API.SpeedWalkDelay())
	}
	v, ok := w.Map[lkey]
	if !ok {
		return nil
	}
	return v
}

// Set a property value for the key. Return true if success, false otherwise.
func (w *World) Set(key string, val goja.Value) bool {
	lkey := strings.ToLower(key)
	switch lkey {
	case "speedwalkdelay":
		w.API.SetSpeedWalkDelay(int(val.ToInteger()))
	default:
		w.Map[lkey] = val
	}
	return true
}

// Has should return true if and only if the property exists.
func (w *World) Has(key string) bool {
	lkey := strings.ToLower(key)
	switch lkey {
	case "speedwalkdelay":
		return true
	}
	_, ok := w.Map[lkey]
	return ok

}

// Delete the property for the key. Returns true on success (note, that includes missing property).
func (w *World) Delete(key string) bool {
	delete(w.Map, strings.ToLower(key))
	return true
}

// Keys returns a list of all existing property keys. There are no checks for duplicates or to make sure
// that the order conforms to https://262.ecma-international.org/#sec-ordinaryownpropertykeys
func (w *World) Keys() []string {
	keys := []string{}
	for k := range w.Map {
		keys = append(keys, strings.ToLower(k))
	}
	return keys
}
func AppendToWorld(r *goja.Runtime, world *goja.Object, name string, call func(call goja.FunctionCall, r *goja.Runtime) goja.Value) {
	r.Set(name, call)
	world.Set(strings.ToLower(name), call)
}
func (a *jsapi) InstallAPIs(p herbplugin.Plugin) {
	jp := p.(*jsplugin.Plugin)
	world := jp.Runtime.NewDynamicObject(CreateWorld(a.API, jp.Runtime))
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
	AppendToWorld(jp.Runtime, world, "GetVariableComment", a.GetVariableComment)
	AppendToWorld(jp.Runtime, world, "SetVariableComment", a.SetVariableComment)

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
	AppendToWorld(jp.Runtime, world, "WorldProxy", a.WorldProxy)

	AppendToWorld(jp.Runtime, world, "Trim", a.Trim)
	AppendToWorld(jp.Runtime, world, "GetUniqueNumber", a.GetUniqueNumber)
	AppendToWorld(jp.Runtime, world, "GetUniqueID", a.GetUniqueID)
	AppendToWorld(jp.Runtime, world, "CreateGUID", a.CreateGUID)
	AppendToWorld(jp.Runtime, world, "FlashIcon", a.FlashIcon)
	AppendToWorld(jp.Runtime, world, "SetStatus", a.SetStatus)
	AppendToWorld(jp.Runtime, world, "Execute", a.Execute)
	AppendToWorld(jp.Runtime, world, "DeleteCommandHistory", a.DeleteCommandHistory)
	AppendToWorld(jp.Runtime, world, "DiscardQueue", a.DiscardQueue)
	AppendToWorld(jp.Runtime, world, "LockQueue", a.LockQueue)
	AppendToWorld(jp.Runtime, world, "GetQueue", a.GetQueue)
	AppendToWorld(jp.Runtime, world, "Queue", a.Queue)

	AppendToWorld(jp.Runtime, world, "DoAfter", a.DoAfter)
	AppendToWorld(jp.Runtime, world, "DoAfterNote", a.DoAfterNote)
	AppendToWorld(jp.Runtime, world, "DoAfterSpeedWalk", a.DoAfterSpeedWalk)
	AppendToWorld(jp.Runtime, world, "DoAfterSpecial", a.DoAfterSpecial)

	AppendToWorld(jp.Runtime, world, "DeleteGroup", a.DeleteGroup)

	AppendToWorld(jp.Runtime, world, "AddTimer", a.AddTimer)
	AppendToWorld(jp.Runtime, world, "DeleteTimer", a.DeleteTimer)
	AppendToWorld(jp.Runtime, world, "DeleteTemporaryTimers", a.DeleteTemporaryTimers)
	AppendToWorld(jp.Runtime, world, "DeleteTimerGroup", a.DeleteTimerGroup)
	AppendToWorld(jp.Runtime, world, "EnableTimer", a.EnableTimer)
	AppendToWorld(jp.Runtime, world, "EnableTimerGroup", a.EnableTimerGroup)
	AppendToWorld(jp.Runtime, world, "GetTimerList", a.GetTimerList)
	AppendToWorld(jp.Runtime, world, "IsTimer", a.IsTimer)
	AppendToWorld(jp.Runtime, world, "ResetTimer", a.ResetTimer)
	AppendToWorld(jp.Runtime, world, "ResetTimers", a.ResetTimers)
	AppendToWorld(jp.Runtime, world, "GetTimerOption", a.GetTimerOption)
	AppendToWorld(jp.Runtime, world, "SetTimerOption", a.SetTimerOption)
	AppendToWorld(jp.Runtime, world, "AddAlias", a.AddAlias)
	AppendToWorld(jp.Runtime, world, "DeleteAlias", a.DeleteAlias)
	AppendToWorld(jp.Runtime, world, "DeleteTemporaryAliases", a.DeleteTemporaryAliases)
	AppendToWorld(jp.Runtime, world, "DeleteAliasGroup", a.DeleteAliasGroup)
	AppendToWorld(jp.Runtime, world, "EnableAlias", a.EnableAlias)
	AppendToWorld(jp.Runtime, world, "EnableAliasGroup", a.EnableAliasGroup)
	AppendToWorld(jp.Runtime, world, "GetAliasList", a.GetAliasList)
	AppendToWorld(jp.Runtime, world, "IsAlias", a.IsAlias)
	AppendToWorld(jp.Runtime, world, "GetAliasOption", a.GetAliasOption)
	AppendToWorld(jp.Runtime, world, "SetAliasOption", a.SetAliasOption)

	AppendToWorld(jp.Runtime, world, "AddTrigger", a.AddTrigger)
	AppendToWorld(jp.Runtime, world, "AddTriggerEx", a.AddTrigger)
	AppendToWorld(jp.Runtime, world, "DeleteTrigger", a.DeleteTrigger)
	AppendToWorld(jp.Runtime, world, "DeleteTemporaryTriggers", a.DeleteTemporaryTriggers)
	AppendToWorld(jp.Runtime, world, "DeleteTriggerGroup", a.DeleteTriggerGroup)
	AppendToWorld(jp.Runtime, world, "EnableTrigger", a.EnableTrigger)
	AppendToWorld(jp.Runtime, world, "EnableTriggerGroup", a.EnableTriggerGroup)
	AppendToWorld(jp.Runtime, world, "GetTriggerList", a.GetTriggerList)
	AppendToWorld(jp.Runtime, world, "IsTrigger", a.IsTrigger)
	AppendToWorld(jp.Runtime, world, "GetTriggerOption", a.GetTriggerOption)
	AppendToWorld(jp.Runtime, world, "SetTriggerOption", a.SetTriggerOption)
	AppendToWorld(jp.Runtime, world, "StopEvaluatingTriggers", a.StopEvaluatingTriggers)
	AppendToWorld(jp.Runtime, world, "GetTriggerWildcard", a.GetTriggerWildcard)

	AppendToWorld(jp.Runtime, world, "ColourNameToRGB", a.ColourNameToRGB)
	AppendToWorld(jp.Runtime, world, "SetSpeedWalkDelay", a.SetSpeedWalkDelay)
	AppendToWorld(jp.Runtime, world, "GetSpeedWalkDelay", a.GetSpeedWalkDelay)

	AppendToWorld(jp.Runtime, world, "HasFile", a.NewHasFileAPI(p))
	AppendToWorld(jp.Runtime, world, "ReadFile", a.NewReadFileAPI(p))
	AppendToWorld(jp.Runtime, world, "ReadLines", a.NewReadLinesAPI(p))

	AppendToWorld(jp.Runtime, world, "HasModFile", a.NewHasModFileAPI(p))
	AppendToWorld(jp.Runtime, world, "ReadModFile", a.NewReadModFileAPI(p))
	AppendToWorld(jp.Runtime, world, "ReadModLines", a.NewReadModLinesAPI(p))
	AppendToWorld(jp.Runtime, world, "GetModInfo", a.NewGetModInfoAPI(p))

	AppendToWorld(jp.Runtime, world, "MakeHomeFolder", a.NewMakeHomeFolderAPI(p))

	AppendToWorld(jp.Runtime, world, "HasHomeFile", a.NewHasHomeFileAPI(p))
	AppendToWorld(jp.Runtime, world, "ReadHomeFile", a.NewReadHomeFileAPI(p))
	AppendToWorld(jp.Runtime, world, "ReadHomeLines", a.NewReadHomeLinesAPI(p))
	AppendToWorld(jp.Runtime, world, "WriteHomeFile", a.NewWriteHomeFileAPI(p))

	AppendToWorld(jp.Runtime, world, "SplitN", a.SplitNfunc)
	AppendToWorld(jp.Runtime, world, "UTF8Len", a.UTF8Len)
	AppendToWorld(jp.Runtime, world, "UTF8Index", a.UTF8Index)
	AppendToWorld(jp.Runtime, world, "UTF8Sub", a.UTF8Sub)
	AppendToWorld(jp.Runtime, world, "ToUTF8", a.ToUTF8)
	AppendToWorld(jp.Runtime, world, "FromUTF8", a.FromUTF8)

	AppendToWorld(jp.Runtime, world, "Info", a.Info)
	AppendToWorld(jp.Runtime, world, "InfoClear", a.InfoClear)

	AppendToWorld(jp.Runtime, world, "GetAlphaOption", a.GetAlphaOption)
	AppendToWorld(jp.Runtime, world, "SetAlphaOption", a.SetAlphaOption)

	AppendToWorld(jp.Runtime, world, "GetLinesInBufferCount", a.GetLinesInBufferCount)
	AppendToWorld(jp.Runtime, world, "DeleteOutput", a.DeleteOutput)
	AppendToWorld(jp.Runtime, world, "DeleteLines", a.DeleteLines)
	AppendToWorld(jp.Runtime, world, "GetLineCount", a.GetLineCount)
	AppendToWorld(jp.Runtime, world, "GetRecentLines", a.GetRecentLines)
	AppendToWorld(jp.Runtime, world, "GetLineInfo", a.GetLineInfo)
	AppendToWorld(jp.Runtime, world, "BoldColour", a.BoldColour)
	AppendToWorld(jp.Runtime, world, "NormalColour", a.NormalColour)
	AppendToWorld(jp.Runtime, world, "GetStyleInfo", a.GetStyleInfo)

	AppendToWorld(jp.Runtime, world, "GetInfo", a.GetInfo)

	AppendToWorld(jp.Runtime, world, "GetTimerInfo", a.GetTimerInfo)
	AppendToWorld(jp.Runtime, world, "GetTriggerInfo", a.GetTriggerInfo)
	AppendToWorld(jp.Runtime, world, "GetAliasInfo", a.GetAliasInfo)

	AppendToWorld(jp.Runtime, world, "WriteLog", a.WriteLog)
	AppendToWorld(jp.Runtime, world, "CloseLog", a.CloseLog)
	AppendToWorld(jp.Runtime, world, "OpenLog", a.OpenLog)
	AppendToWorld(jp.Runtime, world, "FlushLog", a.FlushLog)

	AppendToWorld(jp.Runtime, world, "Broadcast", a.Broadcast)
	AppendToWorld(jp.Runtime, world, "Notify", a.Notify)
	AppendToWorld(jp.Runtime, world, "Request", a.Request)

	AppendToWorld(jp.Runtime, world, "GetGlobalOption", a.GetGlobalOption)

	AppendToWorld(jp.Runtime, world, "CheckPermissions", a.CheckPermissions)
	AppendToWorld(jp.Runtime, world, "RequestPermissions", a.RequestPermissions)
	AppendToWorld(jp.Runtime, world, "CheckTrustedDomains", a.CheckTrustedDomains)
	AppendToWorld(jp.Runtime, world, "RequestTrustDomains", a.RequestTrustDomains)

	AppendToWorld(jp.Runtime, world, "Encrypt", a.Encrypt)
	AppendToWorld(jp.Runtime, world, "Decrypt", a.Decrypt)

	AppendToWorld(jp.Runtime, world, "DumpOutput", a.DumpOutput)
	AppendToWorld(jp.Runtime, world, "ConcatOutput", a.ConcatOutput)
	AppendToWorld(jp.Runtime, world, "SliceOutput", a.SliceOutput)
	AppendToWorld(jp.Runtime, world, "OutputToText", a.OutputToText)
	AppendToWorld(jp.Runtime, world, "FormatOutput", a.FormatOutput)
	AppendToWorld(jp.Runtime, world, "PrintOutput", a.PrintOutput)

	AppendToWorld(jp.Runtime, world, "Simulate", a.Simulate)
	AppendToWorld(jp.Runtime, world, "SimulateOutput", a.SimulateOutput)

	AppendToWorld(jp.Runtime, world, "DumpTriggers", a.DumpTriggers)
	AppendToWorld(jp.Runtime, world, "RestoreTriggers", a.RestoreTriggers)
	AppendToWorld(jp.Runtime, world, "DumpTimers", a.DumpTimers)
	AppendToWorld(jp.Runtime, world, "RestoreTimers", a.RestoreTimers)
	AppendToWorld(jp.Runtime, world, "DumpAliases", a.DumpAliases)
	AppendToWorld(jp.Runtime, world, "RestoreAliases", a.RestoreAliases)

	AppendToWorld(jp.Runtime, world, "SetHUDSize", a.SetHUDSize)
	AppendToWorld(jp.Runtime, world, "GetHUDContent", a.GetHUDContent)
	AppendToWorld(jp.Runtime, world, "GetHUDSize", a.GetHUDSize)
	AppendToWorld(jp.Runtime, world, "UpdateHUD", a.UpdateHUD)
	AppendToWorld(jp.Runtime, world, "NewLine", a.NewLine)
	AppendToWorld(jp.Runtime, world, "NewWord", a.NewWord)

	AppendToWorld(jp.Runtime, world, "SetPriority", a.SetPriority)
	AppendToWorld(jp.Runtime, world, "GetPriority", a.GetPriority)
	AppendToWorld(jp.Runtime, world, "SetSummary", a.SetSummary)
	AppendToWorld(jp.Runtime, world, "GetSummary", a.GetSummary)
	AppendToWorld(jp.Runtime, world, "Save", a.Save)
}

func (a *jsapi) Print(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	msg := []string{}
	for _, v := range call.Arguments {
		msg = append(msg, v.String())
	}
	a.API.Note(strings.Join(msg, " "))
	return goja.Null()
}
func (a *jsapi) Request(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	msgtype := call.Argument(0).String()
	msg := call.Argument(1).String()
	id := a.API.Request(msgtype, msg)
	return r.ToValue(id)
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
	result := make([]string, len(list))
	for k := range list {
		result = append(result, k)
	}
	return r.ToValue(result)
}
func (a *jsapi) GetVariableComment(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	val := a.API.GetVariableComment(call.Argument(0).String())
	return r.ToValue(val)
}
func (a *jsapi) SetVariableComment(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	value := call.Argument(1).String()
	return r.ToValue(a.API.SetVariableComment(name, value))
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
func (a *jsapi) WorldProxy(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.WorldProxy())
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
	return r.ToValue(a.API.DiscardQueue(call.Argument(0).ToBoolean()))
}
func (a *jsapi) LockQueue(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.LockQueue)
}
func (a *jsapi) GetQueue(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	cmds := a.API.GetQueue()
	return r.ToValue(cmds)
}
func (a *jsapi) Queue(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.Queue(call.Argument(0).String(), call.Argument(1).ToBoolean()))
}
func (a *jsapi) DoAfter(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(0).ToFloat()
	send := call.Argument(1).String()
	return r.ToValue(a.API.DoAfter(seconds, send))
}
func (a *jsapi) DoAfterNote(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(0).ToFloat()
	send := call.Argument(1).String()
	return r.ToValue(a.API.DoAfterNote(seconds, send))

}
func (a *jsapi) DoAfterSpeedWalk(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(0).ToFloat()
	send := call.Argument(1).String()
	return r.ToValue(a.API.DoAfterSpeedWalk(seconds, send))
}
func (a *jsapi) DoAfterSpecial(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	seconds := call.Argument(0).ToFloat()
	send := call.Argument(1).String()
	sendto := int(call.Argument(2).ToInteger())
	return r.ToValue(a.API.DoAfterSpecial(seconds, send, sendto))

}

func (a *jsapi) DeleteGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteGroup(name))
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
	enabled := call.Argument(1).ToBoolean()
	return r.ToValue(a.API.EnableTimer(name, enabled))
}
func (a *jsapi) EnableTimerGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	group := call.Argument(0).String()
	enabled := call.Argument(1).ToBoolean()
	return r.ToValue(a.API.EnableTimerGroup(group, enabled))
}

func (a *jsapi) GetTimerList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	list := a.API.GetTimerList()
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}
	return r.NewArray(result)
}
func (a *jsapi) IsTimer(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.IsTimer(name))
}

func (a *jsapi) ResetTimer(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.ResetTimer(name))
}

func (a *jsapi) ResetTimers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.ResetTimers()
	return goja.Null()
}

func (a *jsapi) GetTimerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	option := call.Argument(1).String()
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		return goja.Null()
	} else {
		switch option {
		case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
			return r.ToValue(result == world.StringYes)
		case "group", "name", "script", "send", "variable":
			return r.ToValue(result)
		case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "send_to", "user":
			i, _ := strconv.Atoi(result)
			return r.ToValue(i)
		case "second":
			i, _ := strconv.ParseFloat(result, 64)
			return r.ToValue(i)
		}
	}
	return goja.Null()
}
func (a *jsapi) SetTimerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	option := call.Argument(1).String()
	var value string
	switch option {
	case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
		if call.Argument(2).ToBoolean() {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "script", "send", "variable":
		value = call.Argument(2).String()
	case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "second", "send_to", "user":
		value = call.Argument(2).String()
	}
	return r.ToValue(a.API.SetTimerOption(name, option, value))
}

func (a *jsapi) AddAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	match := call.Argument(1).String()
	send := call.Argument(2).String()
	flags := int(call.Argument(3).ToInteger())
	script := call.Argument(4).String()
	return r.ToValue(a.API.AddAlias(name, match, send, flags, script))
}
func (a *jsapi) DeleteAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteAlias(name))
}
func (a *jsapi) DeleteTemporaryAliases(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.DeleteTemporaryAliases())

}
func (a *jsapi) DeleteAliasGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteAliasGroup(name))
}

func (a *jsapi) EnableAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	enabled := call.Argument(1).ToBoolean()
	return r.ToValue(a.API.EnableAlias(name, enabled))
}
func (a *jsapi) EnableAliasGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	group := call.Argument(0).String()
	enabled := call.Argument(1).ToBoolean()
	return r.ToValue(a.API.EnableAliasGroup(group, enabled))
}

func (a *jsapi) GetAliasList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	list := a.API.GetAliasList()
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}
	return r.NewArray(result...)
}
func (a *jsapi) IsAlias(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.IsAlias(name))
}

func (a *jsapi) GetAliasOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	option := call.Argument(1).String()
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		return goja.Null()
	} else {
		switch option {
		case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			return r.ToValue(result == world.StringYes)
		case "group", "name", "match", "script", "send", "variable":
			return r.ToValue(result)
		case "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			return r.ToValue(i)
		}
	}
	return goja.Null()
}
func (a *jsapi) SetAliasOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	option := call.Argument(1).String()
	var value string
	switch option {
	case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
		if call.Argument(2).ToBoolean() {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "match", "script", "send", "variable":
		value = call.Argument(2).String()
	case "send_to", "user", "sequence":
		value = call.Argument(2).String()
	}
	return r.ToValue(a.API.SetAliasOption(name, option, value))
}

func (a *jsapi) AddTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	match := call.Argument(1).String()
	send := call.Argument(2).String()
	flags := int(call.Argument(3).ToInteger())
	color := int(call.Argument(4).ToInteger())
	wildcard := int(call.Argument(5).ToInteger())
	sound := call.Argument(6).String()
	script := call.Argument(7).String()
	return r.ToValue(a.API.AddTrigger(name, match, send, flags, color, wildcard, sound, script))
}
func (a *jsapi) AddTriggerEx(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	match := call.Argument(1).String()
	send := call.Argument(2).String()
	flags := int(call.Argument(3).ToInteger())
	color := int(call.Argument(4).ToInteger())
	wildcard := int(call.Argument(5).ToInteger())
	sound := call.Argument(6).String()
	script := call.Argument(7).String()
	sendto := int(call.Argument(8).ToInteger())
	sequence := int(call.Argument(9).ToInteger())
	return r.ToValue(a.API.AddTriggerEx(name, match, send, flags, color, wildcard, sound, script, sendto, sequence))
}
func (a *jsapi) DeleteTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteTrigger(name))
}
func (a *jsapi) DeleteTemporaryTriggers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.DeleteTemporaryTimers())

}
func (a *jsapi) DeleteTriggerGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.DeleteTriggerGroup(name))
}

func (a *jsapi) EnableTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	enabled := call.Argument(1).ToBoolean()
	return r.ToValue(a.API.EnableTrigger(name, enabled))
}
func (a *jsapi) EnableTriggerGroup(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	group := call.Argument(0).String()
	enabled := call.Argument(1).ToBoolean()
	return r.ToValue(a.API.EnableTriggerGroup(group, enabled))
}

func (a *jsapi) GetTriggerList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	list := a.API.GetTriggerList()
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}
	return r.NewArray(result...)
}
func (a *jsapi) IsTrigger(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	return r.ToValue(a.API.IsTrigger(name))
}

func (a *jsapi) GetTriggerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	option := call.Argument(1).String()
	result, code := a.API.GetTriggerOption(name, option)
	if code != api.EOK {
		return goja.Null()
	} else {
		switch option {
		case "echo_trigger", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			return r.ToValue(result == world.StringYes)
		case "group", "name", "match", "script", "send", "variable":
			return r.ToValue(result)
		case "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			return r.ToValue(i)
		}
	}
	return goja.Null()
}
func (a *jsapi) SetTriggerOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	option := call.Argument(1).String()
	var value string
	switch option {
	case "echo_trigger", "multi_line", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
		if call.Argument(2).ToBoolean() {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "match", "script", "send", "variable":
		value = call.Argument(2).String()
	case "lines_to_match", "send_to", "user", "sequence":
		value = call.Argument(2).String()
	}
	return r.ToValue(a.API.SetTriggerOption(name, option, value))
}

func (a *jsapi) StopEvaluatingTriggers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.StopEvaluatingTriggers()
	return goja.Null()
}
func (a *jsapi) GetTriggerWildcard(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetTriggerWildcard(call.Argument(0).String(), call.Argument(1).String()))
}

func (a *jsapi) ColourNameToRGB(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	v := a.API.ColourNameToRGB(call.Argument(0).String())
	return r.ToValue(v)
}
func (a *jsapi) SetSpeedWalkDelay(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.SetSpeedWalkDelay(int(call.Argument(0).ToInteger()))
	return goja.Null()
}
func (a *jsapi) GetSpeedWalkDelay(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.SpeedWalkDelay())
}

func (a *jsapi) NewGetModInfoAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		mod := a.API.GetModInfo(p)
		result := r.NewObject()
		result.Set("Enabled", r.ToValue(mod.Enabled))
		result.Set("Exists", r.ToValue(mod.Exists))
		result.Set("FileList", r.ToValue(mod.FileList))
		result.Set("FolderList", r.ToValue(mod.FolderList))
		return result
	}
}
func (a *jsapi) NewHasFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		return r.ToValue(a.API.HasFile(p, call.Argument(0).String()))
	}
}
func (a *jsapi) NewReadFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		return r.ToValue(a.API.ReadFile(p, call.Argument(0).String()))
	}
}
func (a *jsapi) NewHasModFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		return r.ToValue(a.API.HasModFile(p, call.Argument(0).String()))
	}
}
func (a *jsapi) NewMakeHomeFolderAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		return r.ToValue(a.API.MakeHomeFolder(p, call.Argument(0).String()))
	}
}
func (a *jsapi) NewHasHomeFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		return r.ToValue(a.API.HasHomeFile(p, call.Argument(0).String()))
	}
}
func (a *jsapi) NewWriteHomeFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		a.API.WriteHomeFile(p, call.Argument(0).String(), []byte(call.Argument(1).String()))
		return nil
	}
}
func (a *jsapi) NewReadModFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		return r.ToValue(a.API.ReadModFile(p, call.Argument(0).String()))
	}
}
func (a *jsapi) NewReadHomeFileAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		return r.ToValue(a.API.ReadHomeFile(p, call.Argument(0).String()))
	}
}
func (a *jsapi) NewReadLinesAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		lines := a.API.ReadLines(p, call.Argument(0).String())
		t := []interface{}{}
		for _, v := range lines {
			t = append(t, v)
		}
		return r.NewArray(t...)

	}
}
func (a *jsapi) NewReadModLinesAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		lines := a.API.ReadModLines(p, call.Argument(0).String())
		t := []interface{}{}
		for _, v := range lines {
			t = append(t, v)
		}
		return r.NewArray(t...)

	}
}
func (a *jsapi) NewReadHomeLinesAPI(p herbplugin.Plugin) func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return func(call goja.FunctionCall, r *goja.Runtime) goja.Value {
		lines := a.API.ReadHomeLines(p, call.Argument(0).String())
		t := []interface{}{}
		for _, v := range lines {
			t = append(t, v)
		}
		return r.NewArray(t...)

	}
}
func (a *jsapi) SplitNfunc(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	sep := call.Argument(1).String()
	n := int(call.Argument(2).ToInteger())
	s := a.API.SplitN(text, sep, n)
	t := []interface{}{}
	for _, v := range s {
		t = append(t, v)
	}
	return r.NewArray(t...)
}

func (a *jsapi) UTF8Len(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	return r.ToValue(a.API.UTF8Len(text))
}
func (a *jsapi) UTF8Index(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	sub := call.Argument(1).String()
	return r.ToValue(a.API.UTF8Index(text, sub))
}
func (a *jsapi) ToUTF8(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	code := call.Argument(0).String()
	text := call.Argument(1).String()
	return r.ToValue(a.API.ToUTF8(code, text))
}
func (a *jsapi) FromUTF8(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	code := call.Argument(0).String()
	text := call.Argument(1).String()
	return r.ToValue(a.API.FromUTF8(code, text))
}
func (a *jsapi) UTF8Sub(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	start := int(call.Argument(1).ToInteger())
	end := int(call.Argument(2).ToInteger())
	return r.ToValue(a.API.UTF8Sub(text, start, end))
}
func (a *jsapi) Info(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	a.API.Info(text)
	return nil
}
func (a *jsapi) InfoClear(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.InfoClear()
	return nil
}

func (a *jsapi) GetAlphaOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetAlphaOption(call.Argument(0).String()))
}

func (a *jsapi) SetAlphaOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.SetAlphaOption(call.Argument(0).String(), call.Argument(1).String()))
}
func (a *jsapi) WriteLog(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.WriteLog(call.Argument(0).String()))
}

func (a *jsapi) CloseLog(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.CloseLog())
}
func (a *jsapi) OpenLog(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.OpenLog())
}
func (a *jsapi) FlushLog(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.FlushLog())
}

func (a *jsapi) GetLinesInBufferCount(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetLinesInBufferCount())
}
func (a *jsapi) DeleteOutput(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.DeleteOutput()
	return nil
}
func (a *jsapi) DeleteLines(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.DeleteLines(int(call.Argument(0).ToInteger()))
	return nil
}
func (a *jsapi) GetLineCount(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetLineCount())
}
func (a *jsapi) GetRecentLines(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetRecentLines(int(call.Argument(0).ToInteger())))
}
func (a *jsapi) GetLineInfo(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	val, ok := a.API.GetLineInfo(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()))
	if !ok {
		return nil
	}
	switch int(call.Argument(1).ToInteger()) {
	case 1:
		return r.ToValue(val)
	case 2:
		return r.ToValue(world.FromStringInt(val))
	case 3:
		return r.ToValue(world.FromStringInt(val))
	case 4:
		return r.ToValue(world.FromStringBool(val))
	case 5:
		return r.ToValue(world.FromStringBool(val))
	case 6:
		return r.ToValue(world.FromStringBool(val))
	case 7:
		return r.ToValue(world.FromStringBool(val))
	case 8:
		return r.ToValue(world.FromStringBool(val))
	case 9:
		return r.ToValue(world.FromStringInt(val))
	case 11:
		return r.ToValue(world.FromStringInt(val))
	}
	return nil
}
func (a *jsapi) BoldColour(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.BoldColour(int(call.Argument(0).ToInteger())))

}
func (a *jsapi) NormalColour(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.NormalColour(int(call.Argument(0).ToInteger())))
}

func (a *jsapi) GetStyleInfo(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	val, ok := a.API.GetStyleInfo(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()))
	if !ok {
		return nil
	}
	switch int(call.Argument(2).ToInteger()) {
	case 1:
		return r.ToValue(val)
	case 2:
		return r.ToValue(world.FromStringInt(val))
	case 3:
		return r.ToValue(world.FromStringInt(val))
	case 8:
		return r.ToValue(world.FromStringBool(val))
	case 9:
		return r.ToValue(world.FromStringBool(val))
	case 10:
		return r.ToValue(world.FromStringBool(val))
	case 11:
		return r.ToValue(world.FromStringBool(val))
	case 14:
		return r.ToValue(world.FromStringInt(val))
	case 15:
		return r.ToValue(world.FromStringInt(val))

	}
	return nil
}

func (a *jsapi) GetInfo(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetInfo(int(call.Argument(0).ToInteger())))
}
func (a *jsapi) GetTimerInfo(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	v, ok := a.API.GetTimerInfo(call.Argument(0).String(), int(call.Argument(1).ToInteger()))
	if ok != api.EOK {
		return nil
	}
	switch call.Argument(1).ToInteger() {
	case 1:
		return r.ToValue(world.FromStringInt(v))
	case 2:
		return r.ToValue(world.FromStringInt(v))
	case 3:
		return r.ToValue(world.FromStringInt(v))
	case 4:
		return r.ToValue(v)
	case 5:
		return r.ToValue(v)
	case 6:
		return r.ToValue(world.FromStringBool(v))
	case 7:
		return r.ToValue(world.FromStringBool(v))
	case 8:
		return r.ToValue(world.FromStringBool(v))
	case 14:
		return r.ToValue(world.FromStringBool(v))
	case 19:
		return r.ToValue(v)
	case 20:
		return r.ToValue(world.FromStringInt(v))
	case 21:
		return r.ToValue(world.FromStringInt(v))
	case 22:
		return r.ToValue(v)
	case 23:
		return r.ToValue(world.FromStringBool(v))
	case 24:
		return r.ToValue(world.FromStringBool(v))
	}
	return nil
}
func (a *jsapi) GetTriggerInfo(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	v, ok := a.API.GetTriggerInfo(call.Argument(0).String(), int(call.Argument(1).ToInteger()))
	if ok != api.EOK {
		return nil
	}
	switch call.Argument(1).ToInteger() {
	case 1:
		return r.ToValue(v)
	case 2:
		return r.ToValue(v)
	case 3:
		return r.ToValue(v)
	case 4:
		return r.ToValue(v)
	case 5:
		return r.ToValue(world.FromStringBool(v))
	case 6:
		return r.ToValue(world.FromStringBool(v))
	case 7:
		return r.ToValue(world.FromStringBool(v))
	case 8:
		return r.ToValue(world.FromStringBool(v))
	case 9:
		return r.ToValue(world.FromStringBool(v))
	case 10:
		return r.ToValue(world.FromStringBool(v))
	case 11:
		return r.ToValue(world.FromStringBool(v))
	case 13:
		return r.ToValue(world.FromStringBool(v))
	case 15:
		return r.ToValue(world.FromStringInt(v))
	case 16:
		return r.ToValue(world.FromStringInt(v))
	case 23:
		return r.ToValue(world.FromStringBool(v))
	case 25:
		return r.ToValue(world.FromStringBool(v))
	case 26:
		return r.ToValue(v)
	case 27:
		return r.ToValue(v)
	case 28:
		return r.ToValue(world.FromStringInt(v))
	case 31:
		return r.ToValue(world.FromStringInt(v))
	case 36:
		return r.ToValue(world.FromStringBool(v))
	}
	return nil
}

func (a *jsapi) GetAliasInfo(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	v, ok := a.API.GetAliasInfo(call.Argument(0).String(), int(call.Argument(1).ToInteger()))
	if ok != api.EOK {
		return nil
	}
	switch call.Argument(1).ToInteger() {
	case 1:
		return r.ToValue(v)
	case 2:
		return r.ToValue(v)
	case 3:
		return r.ToValue(v)
	case 4:
		return r.ToValue(v)
	case 5:
		return r.ToValue(v)
	case 6:
		return r.ToValue(world.FromStringBool(v))
	case 7:
		return r.ToValue(world.FromStringBool(v))
	case 8:
		return r.ToValue(world.FromStringBool(v))
	case 9:
		return r.ToValue(world.FromStringBool(v))
	case 14:
		return r.ToValue(world.FromStringBool(v))
	case 16:
		return r.ToValue(v)
	case 17:
		return r.ToValue(v)
	case 18:
		return r.ToValue(world.FromStringInt(v))
	case 19:
		return r.ToValue(world.FromStringBool(v))
	case 20:
		return r.ToValue(world.FromStringInt(v))
	case 22:
		return r.ToValue(world.FromStringBool(v))
	case 23:
		return r.ToValue(world.FromStringInt(v))
	case 29:
		return r.ToValue(world.FromStringBool(v))

	}
	return nil
}

func (a *jsapi) Broadcast(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.Broadcast(call.Argument(0).String(), call.Argument(1).ToBoolean())
	return nil
}
func (a *jsapi) Notify(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	a.API.Notify(call.Argument(0).String(), call.Argument(1).String())
	return nil
}
func (a *jsapi) GetGlobalOption(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	result := a.API.GetGlobalOption(call.Argument(0).String())
	switch call.Argument(0).String() {
	default:
		switch result {
		case "0":
			return r.ToValue(0)
		case "1":
			return r.ToValue(1)
		default:
			return r.ToValue(result)
		}
	}
}

func (a *jsapi) CheckPermissions(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	items := []string{}
	err := r.ExportTo(call.Argument(0), &items)
	if err != nil {
		panic(err)
	}
	return r.ToValue(a.API.CheckPermissions(items))
}
func (a *jsapi) RequestPermissions(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	items := []string{}
	err := r.ExportTo(call.Argument(0), &items)
	if err != nil {
		panic(err)
	}
	var reason = ""
	if !goja.IsUndefined(call.Argument(1)) {
		reason = call.Argument(1).String()
	}
	var script = ""
	if !goja.IsUndefined(call.Argument(2)) {
		script = call.Argument(2).String()
	}
	a.API.RequestPermissions(items, reason, script)
	return nil
}
func (a *jsapi) CheckTrustedDomains(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	items := []string{}
	err := r.ExportTo(call.Argument(0), &items)
	if err != nil {
		panic(err)
	}
	return r.ToValue(a.API.CheckTrustedDomains(items))
}

func (a *jsapi) RequestTrustDomains(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	items := []string{}
	err := r.ExportTo(call.Argument(0), &items)
	if err != nil {
		panic(err)
	}
	var reason = ""
	if !goja.IsUndefined(call.Argument(1)) {
		reason = call.Argument(1).String()
	}
	var script = ""
	if !goja.IsUndefined(call.Argument(2)) {
		script = call.Argument(2).String()
	}
	a.API.RequestTrustDomains(items, reason, script)
	return nil
}
func (a *jsapi) Encrypt(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	data := call.Argument(0).String()
	key := call.Argument(1).String()
	result := a.API.Encrypt(data, key)
	if result == nil {
		return nil
	}
	return r.ToValue(*result)
}
func (a *jsapi) Decrypt(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	data := call.Argument(0).String()
	key := call.Argument(1).String()
	result := a.API.Decrypt(data, key)
	if result == nil {
		return nil
	}
	return r.ToValue(*result)
}

func (a *jsapi) DumpOutput(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	length := int(call.Argument(0).ToInteger())
	offset := int(call.Argument(1).ToInteger())
	return r.ToValue(a.API.DumpOutput(length, offset))
}

func (a *jsapi) ConcatOutput(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	output1 := call.Argument(0).String()
	output2 := call.Argument(1).String()
	return r.ToValue(a.API.ConcatOutput(output1, output2))
}
func (a *jsapi) SliceOutput(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	output := call.Argument(0).String()
	start := int(call.Argument(1).ToInteger())
	end := int(call.Argument(2).ToInteger())
	return r.ToValue(a.API.SliceOutput(output, start, end))
}
func (a *jsapi) OutputToText(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	output := call.Argument(0).String()
	return r.ToValue(a.API.OutputToText(output))
}
func (a *jsapi) FormatOutput(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	output := call.Argument(0).String()
	return r.ToValue(a.API.FormatOutput(output))
}
func (a *jsapi) PrintOutput(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	output := call.Argument(0).String()
	return r.ToValue(a.API.PrintOutput(output))
}
func (a *jsapi) Simulate(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	a.API.Simulate(text)
	return nil
}
func (a *jsapi) SimulateOutput(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	output := call.Argument(0).String()
	a.API.SimulateOutput(output)
	return nil
}

func (a *jsapi) DumpTriggers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	byUser := call.Argument(0).ToBoolean()
	result := r.ToValue(a.API.DumpTriggers(byUser))
	return result
}
func (a *jsapi) RestoreTriggers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	data := call.Argument(0).String()
	byUser := call.Argument(1).ToBoolean()
	a.API.RestoreTriggers(data, byUser)
	return nil
}
func (a *jsapi) DumpTimers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	byUser := call.Argument(0).ToBoolean()
	result := r.ToValue(a.API.DumpTimers(byUser))
	return result
}
func (a *jsapi) RestoreTimers(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	data := call.Argument(0).String()
	byUser := call.Argument(1).ToBoolean()
	a.API.RestoreTimers(data, byUser)
	return nil
}
func (a *jsapi) DumpAliases(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	byUser := call.Argument(0).ToBoolean()
	result := r.ToValue(a.API.DumpAliases(byUser))
	return result

}
func (a *jsapi) RestoreAliases(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	data := call.Argument(0).String()
	byUser := call.Argument(1).ToBoolean()
	a.API.RestoreAliases(data, byUser)
	return nil
}
func (a *jsapi) SetHUDSize(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	size := call.Argument(0).ToInteger()
	a.API.SetHUDSize(int(size))
	return nil
}
func (a *jsapi) GetHUDContent(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	content := a.API.GetHUDContent()
	return r.ToValue(content)
}
func (a *jsapi) GetHUDSize(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	size := a.API.GetHUDSize()
	return r.ToValue(size)
}
func (a *jsapi) UpdateHUD(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	start := call.Argument(0).ToInteger()
	content := call.Argument(1).String()
	result := a.API.UpdateHUD(int(start), content)
	return r.ToValue(result)
}
func (a *jsapi) NewLine(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.NewLine())
}
func (a *jsapi) NewWord(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	text := call.Argument(0).String()
	return r.ToValue(a.API.NewWord(text))
}

func (a *jsapi) SetPriority(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	value := int(call.Argument(0).ToInteger())
	a.API.SetPriority(value)
	return goja.Null()
}
func (a *jsapi) GetPriority(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetPriority())
}
func (a *jsapi) SetSummary(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	content := call.Argument(0).String()
	a.API.SetSummary(content)
	return nil
}
func (a *jsapi) GetSummary(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.GetSummary())
}
func (a *jsapi) Save(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.API.Save())
}
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
