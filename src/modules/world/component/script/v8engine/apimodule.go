package v8engine

import (
	"context"
	"encoding/json"
	"modules/world"
	"modules/world/bus"
	"modules/world/component/script/api"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/uniqueid"
	"github.com/herb-go/v8go"
	"github.com/herb-go/v8local"
	"github.com/herb-go/v8local/v8plugin"
)

func createApi(b *bus.Bus) *jsapi {
	return &jsapi{
		API: &api.API{
			Bus: b,
		},
	}
}

type jsapi struct {
	API    *api.API
	Plugin herbplugin.Plugin
}

func AppendToWorld(local *v8local.Local, world *v8local.JsValue, name string, call func(call *v8local.FunctionCallbackInfo) *v8local.JsValue) {
	fn := local.Context().NewFunctionTemplate(call)
	global := local.Global()
	global.Set(name, fn.GetLocalFunction(local))

	world.Set(strings.ToLower(name), fn.GetLocalFunction(local))
	if name != strings.ToLower(name) {
		world.Set(name, fn.GetLocalFunction(local))
	}
}
func (a *jsapi) InstallAPIs(p herbplugin.Plugin) {
	a.Plugin = p
	jp := p.(*v8plugin.Plugin)
	local := jp.Runtime.NewLocal()
	defer local.Close()
	world := local.NewObject()
	AppendToWorld(local, world, "print", a.Print)
	AppendToWorld(local, world, "Note", a.Note)
	AppendToWorld(local, world, "SendImmediate", a.SendImmediate)
	AppendToWorld(local, world, "Send", a.Send)
	AppendToWorld(local, world, "SendNoEcho", a.SendNoEcho)
	AppendToWorld(local, world, "GetVariable", a.GetVariable)
	AppendToWorld(local, world, "SetVariable", a.SetVariable)
	AppendToWorld(local, world, "DeleteVariable", a.DeleteVariable)
	AppendToWorld(local, world, "GetVariableList", a.GetVariableList)
	AppendToWorld(local, world, "GetVariableComment", a.GetVariableComment)
	AppendToWorld(local, world, "SetVariableComment", a.SetVariableComment)
	AppendToWorld(local, world, "Version", a.Version)
	AppendToWorld(local, world, "Hash", a.Hash)
	AppendToWorld(local, world, "Base64Encode", a.Base64Encode)
	AppendToWorld(local, world, "Base64Decode", a.Base64Decode)
	AppendToWorld(local, world, "Connect", a.Connect)
	AppendToWorld(local, world, "IsConnected", a.IsConnected)
	AppendToWorld(local, world, "Disconnect", a.Disconnect)
	AppendToWorld(local, world, "GetWorldById", a.GetWorldById)
	AppendToWorld(local, world, "GetWorld", a.GetWorld)
	AppendToWorld(local, world, "GetWorldID", a.GetWorldID)
	AppendToWorld(local, world, "GetWorldIdList", a.GetWorldIdList)
	AppendToWorld(local, world, "GetWorldList", a.GetWorldList)
	AppendToWorld(local, world, "WorldName", a.WorldName)
	AppendToWorld(local, world, "WorldAddress", a.WorldAddress)
	AppendToWorld(local, world, "WorldPort", a.WorldPort)
	AppendToWorld(local, world, "WorldProxy", a.WorldProxy)
	AppendToWorld(local, world, "Trim", a.Trim)
	AppendToWorld(local, world, "GetUniqueNumber", a.GetUniqueNumber)
	AppendToWorld(local, world, "GetUniqueID", a.GetUniqueID)
	AppendToWorld(local, world, "CreateGUID", a.CreateGUID)
	AppendToWorld(local, world, "FlashIcon", a.FlashIcon)
	AppendToWorld(local, world, "SetStatus", a.SetStatus)
	AppendToWorld(local, world, "Execute", a.Execute)
	AppendToWorld(local, world, "DeleteCommandHistory", a.DeleteCommandHistory)
	AppendToWorld(local, world, "DiscardQueue", a.DiscardQueue)
	AppendToWorld(local, world, "LockQueue", a.LockQueue)
	AppendToWorld(local, world, "GetQueue", a.GetQueue)
	AppendToWorld(local, world, "Queue", a.Queue)
	AppendToWorld(local, world, "DoAfter", a.DoAfter)
	AppendToWorld(local, world, "DoAfterNote", a.DoAfterNote)
	AppendToWorld(local, world, "DoAfterSpeedWalk", a.DoAfterSpeedWalk)
	AppendToWorld(local, world, "DoAfterSpecial", a.DoAfterSpecial)
	AppendToWorld(local, world, "DeleteGroup", a.DeleteGroup)
	AppendToWorld(local, world, "AddTimer", a.AddTimer)
	AppendToWorld(local, world, "DeleteTimer", a.DeleteTimer)
	AppendToWorld(local, world, "DeleteTemporaryTimers", a.DeleteTemporaryTimers)
	AppendToWorld(local, world, "DeleteTimerGroup", a.DeleteTimerGroup)
	AppendToWorld(local, world, "EnableTimer", a.EnableTimer)
	AppendToWorld(local, world, "EnableTimerGroup", a.EnableTimerGroup)
	AppendToWorld(local, world, "GetTimerList", a.GetTimerList)
	AppendToWorld(local, world, "IsTimer", a.IsTimer)
	AppendToWorld(local, world, "ResetTimer", a.ResetTimer)
	AppendToWorld(local, world, "ResetTimers", a.ResetTimers)
	AppendToWorld(local, world, "GetTimerOption", a.GetTimerOption)
	AppendToWorld(local, world, "SetTimerOption", a.SetTimerOption)
	AppendToWorld(local, world, "AddAlias", a.AddAlias)
	AppendToWorld(local, world, "DeleteAlias", a.DeleteAlias)
	AppendToWorld(local, world, "DeleteTemporaryAliases", a.DeleteTemporaryAliases)
	AppendToWorld(local, world, "DeleteAliasGroup", a.DeleteAliasGroup)
	AppendToWorld(local, world, "EnableAlias", a.EnableAlias)
	AppendToWorld(local, world, "EnableAliasGroup", a.EnableAliasGroup)
	AppendToWorld(local, world, "GetAliasList", a.GetAliasList)
	AppendToWorld(local, world, "IsAlias", a.IsAlias)
	AppendToWorld(local, world, "GetAliasOption", a.GetAliasOption)
	AppendToWorld(local, world, "SetAliasOption", a.SetAliasOption)
	AppendToWorld(local, world, "AddTrigger", a.AddTrigger)
	AppendToWorld(local, world, "AddTriggerEx", a.AddTrigger)
	AppendToWorld(local, world, "DeleteTrigger", a.DeleteTrigger)
	AppendToWorld(local, world, "DeleteTemporaryTriggers", a.DeleteTemporaryTriggers)
	AppendToWorld(local, world, "DeleteTriggerGroup", a.DeleteTriggerGroup)
	AppendToWorld(local, world, "EnableTrigger", a.EnableTrigger)
	AppendToWorld(local, world, "EnableTriggerGroup", a.EnableTriggerGroup)
	AppendToWorld(local, world, "GetTriggerList", a.GetTriggerList)
	AppendToWorld(local, world, "IsTrigger", a.IsTrigger)
	AppendToWorld(local, world, "GetTriggerOption", a.GetTriggerOption)
	AppendToWorld(local, world, "SetTriggerOption", a.SetTriggerOption)
	AppendToWorld(local, world, "StopEvaluatingTriggers", a.StopEvaluatingTriggers)
	AppendToWorld(local, world, "GetTriggerWildcard", a.GetTriggerWildcard)
	AppendToWorld(local, world, "ColourNameToRGB", a.ColourNameToRGB)
	AppendToWorld(local, world, "SetSpeedWalkDelay", a.SetSpeedWalkDelay)
	AppendToWorld(local, world, "GetSpeedWalkDelay", a.GetSpeedWalkDelay)
	AppendToWorld(local, world, "HasFile", a.NewHasFileAPI)
	AppendToWorld(local, world, "ReadFile", a.NewReadFileAPI)
	AppendToWorld(local, world, "ReadLines", a.NewReadLinesAPI)
	AppendToWorld(local, world, "HasModFile", a.NewHasModFileAPI)
	AppendToWorld(local, world, "ReadModFile", a.NewReadModFileAPI)
	AppendToWorld(local, world, "ReadModLines", a.NewReadModLinesAPI)
	AppendToWorld(local, world, "GetModInfo", a.NewGetModInfoAPI)
	AppendToWorld(local, world, "MakeHomeFolder", a.NewMakeHomeFolderAPI)
	AppendToWorld(local, world, "HasHomeFile", a.NewHasHomeFileAPI)
	AppendToWorld(local, world, "ReadHomeFile", a.NewReadHomeFileAPI)
	AppendToWorld(local, world, "ReadHomeLines", a.NewReadHomeLinesAPI)
	AppendToWorld(local, world, "WriteHomeFile", a.NewWriteHomeFileAPI)
	AppendToWorld(local, world, "SplitN", a.SplitNfunc)
	AppendToWorld(local, world, "UTF8Len", a.UTF8Len)
	AppendToWorld(local, world, "UTF8Index", a.UTF8Index)
	AppendToWorld(local, world, "UTF8Sub", a.UTF8Sub)
	AppendToWorld(local, world, "ToUTF8", a.ToUTF8)
	AppendToWorld(local, world, "FromUTF8", a.FromUTF8)
	AppendToWorld(local, world, "Info", a.Info)
	AppendToWorld(local, world, "InfoClear", a.InfoClear)
	AppendToWorld(local, world, "GetAlphaOption", a.GetAlphaOption)
	AppendToWorld(local, world, "SetAlphaOption", a.SetAlphaOption)
	AppendToWorld(local, world, "GetLinesInBufferCount", a.GetLinesInBufferCount)
	AppendToWorld(local, world, "DeleteOutput", a.DeleteOutput)
	AppendToWorld(local, world, "DeleteLines", a.DeleteLines)
	AppendToWorld(local, world, "GetLineCount", a.GetLineCount)
	AppendToWorld(local, world, "GetRecentLines", a.GetRecentLines)
	AppendToWorld(local, world, "GetLineInfo", a.GetLineInfo)
	AppendToWorld(local, world, "BoldColour", a.BoldColour)
	AppendToWorld(local, world, "NormalColour", a.NormalColour)
	AppendToWorld(local, world, "GetStyleInfo", a.GetStyleInfo)
	AppendToWorld(local, world, "GetInfo", a.GetInfo)
	AppendToWorld(local, world, "GetTimerInfo", a.GetTimerInfo)
	AppendToWorld(local, world, "GetTriggerInfo", a.GetTriggerInfo)
	AppendToWorld(local, world, "GetAliasInfo", a.GetAliasInfo)
	AppendToWorld(local, world, "WriteLog", a.WriteLog)
	AppendToWorld(local, world, "CloseLog", a.CloseLog)
	AppendToWorld(local, world, "OpenLog", a.OpenLog)
	AppendToWorld(local, world, "FlushLog", a.FlushLog)
	AppendToWorld(local, world, "Broadcast", a.Broadcast)
	AppendToWorld(local, world, "Notify", a.Notify)
	AppendToWorld(local, world, "Request", a.Request)
	AppendToWorld(local, world, "GetGlobalOption", a.GetGlobalOption)
	AppendToWorld(local, world, "CheckPermissions", a.CheckPermissions)
	AppendToWorld(local, world, "RequestPermissions", a.RequestPermissions)
	AppendToWorld(local, world, "CheckTrustedDomains", a.CheckTrustedDomains)
	AppendToWorld(local, world, "RequestTrustDomains", a.RequestTrustDomains)
	AppendToWorld(local, world, "Encrypt", a.Encrypt)
	AppendToWorld(local, world, "Decrypt", a.Decrypt)
	AppendToWorld(local, world, "DumpOutput", a.DumpOutput)
	AppendToWorld(local, world, "ConcatOutput", a.ConcatOutput)
	AppendToWorld(local, world, "SliceOutput", a.SliceOutput)
	AppendToWorld(local, world, "OutputToText", a.OutputToText)
	AppendToWorld(local, world, "FormatOutput", a.FormatOutput)
	AppendToWorld(local, world, "PrintOutput", a.PrintOutput)
	AppendToWorld(local, world, "Simulate", a.Simulate)
	AppendToWorld(local, world, "SimulateOutput", a.SimulateOutput)
	AppendToWorld(local, world, "DumpTriggers", a.DumpTriggers)
	AppendToWorld(local, world, "RestoreTriggers", a.RestoreTriggers)
	AppendToWorld(local, world, "DumpTimers", a.DumpTimers)
	AppendToWorld(local, world, "RestoreTimers", a.RestoreTimers)
	AppendToWorld(local, world, "DumpAliases", a.DumpAliases)
	AppendToWorld(local, world, "RestoreAliases", a.RestoreAliases)
	AppendToWorld(local, world, "SetHUDSize", a.SetHUDSize)
	AppendToWorld(local, world, "GetHUDContent", a.GetHUDContent)
	AppendToWorld(local, world, "GetHUDSize", a.GetHUDSize)
	AppendToWorld(local, world, "UpdateHUD", a.UpdateHUD)
	AppendToWorld(local, world, "NewLine", a.NewLine)
	AppendToWorld(local, world, "NewWord", a.NewWord)
	AppendToWorld(local, world, "SetPriority", a.SetPriority)
	AppendToWorld(local, world, "GetPriority", a.GetPriority)
	AppendToWorld(local, world, "SetSummary", a.SetSummary)
	AppendToWorld(local, world, "GetSummary", a.GetSummary)
	AppendToWorld(local, world, "Save", a.Save)
	AppendToWorld(local, world, "Milliseconds", a.Milliseconds)
	AppendToWorld(local, world, "OmitOutput", a.OmitOutput)
	AppendToWorld(local, world, "PrintSystem", a.PrintSystem)
	AppendToWorld(local, world, "V8Debug", a.Debug)
	AppendToWorld(local, world, "Snapshot", a.Snapshot)
	global := local.Global()
	global.Set("world", world)

}
func (a *jsapi) Snapshot(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	p := filepath.Join(a.API.Bus.GetLogsPath(), a.API.Bus.ID+"."+uniqueid.MustGenerateID()+".heapsnapshot")
	v8go.WriteHeapSnapshot(call.Local().Context().Raw.Isolate(), p)
	a.API.Note("镜像文件写入" + p)
	return nil
}

func (a *jsapi) Debug(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	debug.FreeOSMemory()
	v8go.ForceV8GC(call.Local().Context().Raw.Isolate())
	a.API.Note("V8 Isolate Count:" + strconv.Itoa(call.Local().Context().Raw.Isolate().RetainedValueCount()))
	a.API.Note("V8 Context Count:" + strconv.Itoa(call.Local().Context().Raw.RetainedValueCount()))
	a.API.Note("V8 Local Count:" + strconv.Itoa(a.Plugin.(*v8plugin.Plugin).Top.RetainedValueCount()))
	bs, err := json.Marshal(call.Local().Context().Raw.Isolate().GetHeapStatistics())
	if err != nil {
		panic(err)
	}
	a.API.Note("V8版本:" + v8go.Version())
	a.API.Note("V8内存统计:" + string(bs))
	return nil
}
func (a *jsapi) Print(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	msg := []string{}
	for _, v := range call.Args() {
		msg = append(msg, v.String())
	}
	a.API.Note(strings.Join(msg, " "))
	return nil
}

func (a *jsapi) Request(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	msgtype := call.GetArg(0).String()
	msg := call.GetArg(1).String()
	id := a.API.Request(msgtype, msg)
	return call.Local().NewString(id)
}
func (a *jsapi) Note(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	info := call.GetArg(0).String()
	a.API.Note(info)
	return nil
}
func (a *jsapi) PrintSystem(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	info := call.GetArg(0).String()
	a.API.PrintSystem(info)
	return nil
}
func (a *jsapi) SendImmediate(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	info := call.GetArg(0).String()

	return call.Local().NewInt32(int32(a.API.SendImmediate(info)))

}
func (a *jsapi) Send(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	info := call.GetArg(0).String()

	res := a.API.Send(info)
	return call.Local().NewInt32(int32(res))
}
func (a *jsapi) Execute(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	info := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.Execute(info)))
}
func (a *jsapi) SendNoEcho(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	info := call.GetArg(0).String()

	return call.Local().NewInt32(int32(a.API.SendNoEcho(info)))
}
func (a *jsapi) GetVariable(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	val := a.API.GetVariable(call.GetArg(0).String())
	return call.Local().NewString(val)
}
func (a *jsapi) DeleteVariable(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteVariable(name)))
}
func (a *jsapi) SetVariable(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	value := call.GetArg(1).String()
	return call.Local().NewInt32(int32(a.API.SetVariable(name, value)))
}
func (a *jsapi) GetVariableList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	list := a.API.GetVariableList()
	result := make([]string, len(list))
	for k := range list {
		result = append(result, k)
	}
	return call.Local().NewStringArray(result...)
}
func (a *jsapi) GetVariableComment(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	val := a.API.GetVariableComment(call.GetArg(0).String())
	return call.Local().NewString(val)
}
func (a *jsapi) SetVariableComment(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	value := call.GetArg(1).String()
	return call.Local().NewInt32(int32(a.API.SetVariableComment(name, value)))
}
func (a *jsapi) Version(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.Version())
}
func (a *jsapi) Hash(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewString(a.API.Hash(name))
}
func (a *jsapi) Base64Encode(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	src := call.GetArg(0).String()
	ok := call.GetArg(1).Boolean()
	return call.Local().NewString(a.API.Base64Encode(src, ok))
}
func (a *jsapi) Base64Decode(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	src := call.GetArg(0).String()
	result := a.API.Base64Decode(src)
	if result == nil {
		return nil
	}
	return call.Local().NewString(*result)
}
func (a *jsapi) Connect(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.Connect()))
}
func (a *jsapi) IsConnected(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewBoolean(a.API.IsConnected())

}
func (a *jsapi) Disconnect(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.Disconnect()))
}

func (a *jsapi) GetWorldById(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return nil
}

func (a *jsapi) GetWorld(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return nil
}

func (a *jsapi) GetWorldID(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.GetWorldID())

}
func (a *jsapi) GetWorldIdList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewArray()
}
func (a *jsapi) GetWorldList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewArray()
}
func (a *jsapi) WorldName(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.WorldName())
}
func (a *jsapi) WorldAddress(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.WorldAddress())
}
func (a *jsapi) WorldPort(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.WorldPort()))
}
func (a *jsapi) WorldProxy(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.WorldProxy())
}

func (a *jsapi) Trim(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	src := call.GetArg(0).String()
	return call.Local().NewString(a.API.Trim(src))
}
func (a *jsapi) GetUniqueNumber(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt64(int64(a.API.GetUniqueNumber()))
}
func (a *jsapi) GetUniqueID(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.GetUniqueID())
}
func (a *jsapi) CreateGUID(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.CreateGUID())
}
func (a *jsapi) FlashIcon(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.FlashIcon()
	return nil
}
func (a *jsapi) SetStatus(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	text := call.GetArg(0).String()
	a.API.SetStatus(text)
	return nil
}
func (a *jsapi) DeleteCommandHistory(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.DeleteCommandHistory()
	return nil
}
func (a *jsapi) DiscardQueue(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.DiscardQueue(call.GetArg(0).Boolean())))
}
func (a *jsapi) LockQueue(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	a.API.LockQueue()
	return nil
}
func (a *jsapi) GetQueue(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	cmds := a.API.GetQueue()
	return call.Local().NewStringArray(cmds...)
}
func (a *jsapi) Queue(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.Queue(call.GetArg(0).String(), call.GetArg(1).Boolean())))
}
func (a *jsapi) DoAfter(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	return call.Local().NewInt32(int32(a.API.DoAfter(seconds, send)))
}
func (a *jsapi) DoAfterNote(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	return call.Local().NewInt32(int32(a.API.DoAfterNote(seconds, send)))

}
func (a *jsapi) DoAfterSpeedWalk(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	return call.Local().NewInt32(int32(a.API.DoAfterSpeedWalk(seconds, send)))
}
func (a *jsapi) DoAfterSpecial(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	sendto := int(call.GetArg(2).Integer())
	return call.Local().NewInt32(int32(a.API.DoAfterSpecial(seconds, send, sendto)))

}

func (a *jsapi) DeleteGroup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteGroup(name)))
}

func (a *jsapi) AddTimer(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	hour := int(call.GetArg(1).Integer())
	min := int(call.GetArg(2).Integer())
	seconds := call.GetArg(3).Number()
	send := call.GetArg(4).String()
	flags := int(call.GetArg(5).Integer())
	script := call.GetArg(6).String()
	return call.Local().NewInt32(int32(a.API.AddTimer(name, hour, min, seconds, send, flags, script)))
}
func (a *jsapi) DeleteTimer(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteTimer(name)))

}
func (a *jsapi) DeleteTemporaryTimers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.DeleteTemporaryTimers()))

}
func (a *jsapi) DeleteTimerGroup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteTimerGroup(name)))
}

func (a *jsapi) EnableTimer(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Local().NewInt32(int32(a.API.EnableTimer(name, enabled)))
}
func (a *jsapi) EnableTimerGroup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	group := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Local().NewInt32(int32(a.API.EnableTimerGroup(group, enabled)))
}

func (a *jsapi) GetTimerList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	list := a.API.GetTimerList()
	result := []*v8local.JsValue{}
	for _, v := range list {
		result = append(result, call.Local().NewString(v))
	}
	return call.Local().NewArray(result...)
}
func (a *jsapi) IsTimer(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.IsTimer(name)))
}

func (a *jsapi) ResetTimer(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.ResetTimer(name)))
}

func (a *jsapi) ResetTimers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.ResetTimers()
	return nil
}

func (a *jsapi) GetTimerOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		return nil
	} else {
		switch option {
		case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
			return call.Local().NewBoolean(result == world.StringYes)
		case "group", "name", "script", "send", "variable":
			return call.Local().NewString(result)
		case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "send_to", "user":
			i, _ := strconv.Atoi(result)
			return call.Local().NewInt32(int32(i))
		case "second":
			i, _ := strconv.ParseFloat(result, 64)
			return call.Local().NewNumber(i)
		}
	}
	return nil
}
func (a *jsapi) SetTimerOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	var value string
	switch option {
	case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
		if call.GetArg(2).Boolean() {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "script", "send", "variable":
		value = call.GetArg(2).String()
	case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "second", "send_to", "user":
		value = call.GetArg(2).String()
	}
	return call.Local().NewInt32(int32(a.API.SetTimerOption(name, option, value)))
}

func (a *jsapi) AddAlias(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	match := call.GetArg(1).String()
	send := call.GetArg(2).String()
	flags := int(call.GetArg(3).Integer())
	script := call.GetArg(4).String()
	return call.Local().NewInt32(int32(a.API.AddAlias(name, match, send, flags, script)))
}
func (a *jsapi) DeleteAlias(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteAlias(name)))
}
func (a *jsapi) DeleteTemporaryAliases(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.DeleteTemporaryAliases()))

}
func (a *jsapi) DeleteAliasGroup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteAliasGroup(name)))
}

func (a *jsapi) EnableAlias(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Local().NewInt32(int32(a.API.EnableAlias(name, enabled)))
}
func (a *jsapi) EnableAliasGroup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	group := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Local().NewInt32(int32(a.API.EnableAliasGroup(group, enabled)))
}

func (a *jsapi) GetAliasList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	list := a.API.GetAliasList()
	result := []*v8local.JsValue{}
	for _, v := range list {
		result = append(result, call.Local().NewString(v))
	}
	return call.Local().NewArray(result...)
}
func (a *jsapi) IsAlias(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.IsAlias(name)))
}

func (a *jsapi) GetAliasOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		return nil
	} else {
		switch option {
		case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			return call.Local().NewBoolean(result == world.StringYes)
		case "group", "name", "match", "script", "send", "variable":
			return call.Local().NewString(result)
		case "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			return call.Local().NewInt32(int32(i))
		}
	}
	return nil
}
func (a *jsapi) SetAliasOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	var value string
	switch option {
	case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
		if call.GetArg(2).Boolean() {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "match", "script", "send", "variable":
		value = call.GetArg(2).String()
	case "send_to", "user", "sequence":
		value = call.GetArg(2).String()
	}
	return call.Local().NewInt32(int32(a.API.SetAliasOption(name, option, value)))
}

func (a *jsapi) AddTrigger(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	match := call.GetArg(1).String()
	send := call.GetArg(2).String()
	flags := int(call.GetArg(3).Integer())
	color := int(call.GetArg(4).Integer())
	wildcard := int(call.GetArg(5).Integer())
	sound := call.GetArg(6).String()
	script := call.GetArg(7).String()
	return call.Local().NewInt32(int32(a.API.AddTrigger(name, match, send, flags, color, wildcard, sound, script)))
}
func (a *jsapi) AddTriggerEx(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	match := call.GetArg(1).String()
	send := call.GetArg(2).String()
	flags := int(call.GetArg(3).Integer())
	color := int(call.GetArg(4).Integer())
	wildcard := int(call.GetArg(5).Integer())
	sound := call.GetArg(6).String()
	script := call.GetArg(7).String()
	sendto := int(call.GetArg(8).Integer())
	sequence := int(call.GetArg(9).Integer())
	return call.Local().NewInt32(int32(a.API.AddTriggerEx(name, match, send, flags, color, wildcard, sound, script, sendto, sequence)))
}
func (a *jsapi) DeleteTrigger(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteTrigger(name)))
}
func (a *jsapi) DeleteTemporaryTriggers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.DeleteTemporaryTimers()))

}
func (a *jsapi) DeleteTriggerGroup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.DeleteTriggerGroup(name)))
}

func (a *jsapi) EnableTrigger(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Local().NewInt32(int32(a.API.EnableTrigger(name, enabled)))
}
func (a *jsapi) EnableTriggerGroup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	group := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Local().NewInt32(int32(a.API.EnableTriggerGroup(group, enabled)))
}

func (a *jsapi) GetTriggerList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	list := a.API.GetTriggerList()
	result := []*v8local.JsValue{}
	for _, v := range list {
		result = append(result, call.Local().NewString(v))
	}
	return call.Local().NewArray(result...)
}
func (a *jsapi) IsTrigger(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.IsTrigger(name)))
}

func (a *jsapi) GetTriggerOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	result, code := a.API.GetTriggerOption(name, option)
	if code != api.EOK {
		return nil
	} else {
		switch option {
		case "echo_trigger", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			return call.Local().NewBoolean(result == world.StringYes)
		case "group", "name", "match", "script", "send", "variable":
			return call.Local().NewString(result)
		case "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			return call.Local().NewInt32(int32(i))
		}
	}
	return nil
}
func (a *jsapi) SetTriggerOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	var value string
	switch option {
	case "echo_trigger", "multi_line", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "omit_from_log", "omit_from_output", "one_shot", "regexp":
		if call.GetArg(2).Boolean() {
			value = world.StringYes
		} else {
			value = ""
		}
	case "group", "name", "match", "script", "send", "variable":
		value = call.GetArg(2).String()
	case "lines_to_match", "send_to", "user", "sequence":
		value = call.GetArg(2).String()
	}
	return call.Local().NewInt32(int32(a.API.SetTriggerOption(name, option, value)))
}

func (a *jsapi) StopEvaluatingTriggers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.StopEvaluatingTriggers()
	return nil
}
func (a *jsapi) GetTriggerWildcard(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	result := a.API.GetTriggerWildcard(call.GetArg(0).String(), call.GetArg(1).String())
	if result == nil {
		return nil
	}
	return call.Local().NewString(*result)
}

func (a *jsapi) ColourNameToRGB(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	v := a.API.ColourNameToRGB(call.GetArg(0).String())
	return call.Local().NewInt32(int32(v))
}
func (a *jsapi) SetSpeedWalkDelay(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.SetSpeedWalkDelay(int(call.GetArg(0).Integer()))
	return nil
}
func (a *jsapi) GetSpeedWalkDelay(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.SpeedWalkDelay()))
}

func (a *jsapi) NewGetModInfoAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	mod := a.API.GetModInfo(a.Plugin)
	result := call.Local().NewObject()
	result.Set("Enabled", call.Local().NewBoolean(mod.Enabled))
	result.Set("Exists", call.Local().NewBoolean(mod.Exists))
	result.Set("FileList", call.Local().NewStringArray(mod.FileList...))
	result.Set("FolderList", call.Local().NewStringArray(mod.FolderList...))
	return result
}
func (a *jsapi) NewHasFileAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewBoolean(a.API.HasFile(a.Plugin, call.GetArg(0).String()))
}
func (a *jsapi) NewReadFileAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return call.Local().NewString(a.API.ReadFile(a.Plugin, call.GetArg(0).String()))
}
func (a *jsapi) NewHasModFileAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewBoolean(a.API.HasModFile(a.Plugin, call.GetArg(0).String()))
}
func (a *jsapi) NewMakeHomeFolderAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewBoolean(a.API.MakeHomeFolder(a.Plugin, call.GetArg(0).String()))
}
func (a *jsapi) NewHasHomeFileAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return call.Local().NewBoolean(a.API.HasHomeFile(a.Plugin, call.GetArg(0).String()))
}
func (a *jsapi) NewWriteHomeFileAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.WriteHomeFile(a.Plugin, call.GetArg(0).String(), []byte(call.GetArg(1).String()))
	return nil
}
func (a *jsapi) NewReadModFileAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return call.Local().NewString(a.API.ReadModFile(a.Plugin, call.GetArg(0).String()))
}
func (a *jsapi) NewReadHomeFileAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return call.Local().NewString(a.API.ReadHomeFile(a.Plugin, call.GetArg(0).String()))
}
func (a *jsapi) NewReadLinesAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	lines := a.API.ReadLines(a.Plugin, call.GetArg(0).String())
	t := []*v8local.JsValue{}
	for _, v := range lines {
		t = append(t, call.Local().NewString(v))
	}
	return call.Local().NewArray(t...)

}
func (a *jsapi) NewReadModLinesAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	lines := a.API.ReadModLines(a.Plugin, call.GetArg(0).String())
	t := []*v8local.JsValue{}
	for _, v := range lines {
		t = append(t, call.Local().NewString(v))
	}
	return call.Local().NewArray(t...)
}
func (a *jsapi) NewReadHomeLinesAPI(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	lines := a.API.ReadHomeLines(a.Plugin, call.GetArg(0).String())
	t := []*v8local.JsValue{}
	for _, v := range lines {
		t = append(t, call.Local().NewString(v))
	}
	return call.Local().NewArray(t...)
}
func (a *jsapi) SplitNfunc(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	text := call.GetArg(0).String()
	sep := call.GetArg(1).String()
	n := int(call.GetArg(2).Integer())
	s := a.API.SplitN(text, sep, n)
	return call.Local().NewStringArray(s...)
}

func (a *jsapi) UTF8Len(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	text := call.GetArg(0).String()
	return call.Local().NewInt32(int32(a.API.UTF8Len(text)))
}
func (a *jsapi) UTF8Index(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	text := call.GetArg(0).String()
	sub := call.GetArg(1).String()
	return call.Local().NewInt32(int32(a.API.UTF8Index(text, sub)))
}
func (a *jsapi) ToUTF8(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	code := call.GetArg(0).String()
	text := call.GetArg(1).String()
	result := a.API.ToUTF8(code, text)
	if result == nil {
		return nil
	}
	return call.Local().NewString(*result)
}
func (a *jsapi) FromUTF8(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	code := call.GetArg(0).String()
	text := call.GetArg(1).String()
	result := a.API.FromUTF8(code, text)
	if result == nil {
		return nil
	}
	return call.Local().NewString(*result)
}
func (a *jsapi) UTF8Sub(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	text := call.GetArg(0).String()
	start := int(call.GetArg(1).Integer())
	end := int(call.GetArg(2).Integer())
	return call.Local().NewString(a.API.UTF8Sub(text, start, end))
}
func (a *jsapi) Info(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	text := call.GetArg(0).String()
	a.API.Info(text)
	return nil
}
func (a *jsapi) InfoClear(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.InfoClear()
	return nil
}

func (a *jsapi) GetAlphaOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.GetAlphaOption(call.GetArg(0).String()))
}

func (a *jsapi) SetAlphaOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.SetAlphaOption(call.GetArg(0).String(), call.GetArg(1).String())))
}
func (a *jsapi) WriteLog(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.WriteLog(call.GetArg(0).String())))
}

func (a *jsapi) CloseLog(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.CloseLog()))
}
func (a *jsapi) OpenLog(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.OpenLog()))
}
func (a *jsapi) FlushLog(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.FlushLog()))
}

func (a *jsapi) GetLinesInBufferCount(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.GetLinesInBufferCount()))
}
func (a *jsapi) DeleteOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.DeleteOutput()
	return nil
}
func (a *jsapi) DeleteLines(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.DeleteLines(int(call.GetArg(0).Integer()))
	return nil
}
func (a *jsapi) GetLineCount(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.GetLineCount()))
}
func (a *jsapi) GetRecentLines(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.GetRecentLines(int(call.GetArg(0).Integer())))
}
func (a *jsapi) GetLineInfo(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	val, ok := a.API.GetLineInfo(int(call.GetArg(0).Integer()), int(call.GetArg(1).Integer()))
	if !ok {
		return nil
	}
	switch int(call.GetArg(1).Integer()) {
	case 1:
		return call.Local().NewString(val)
	case 2:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))
	case 3:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))
	case 4:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 5:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 6:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 7:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 8:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 9:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))
	case 11:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))
	}
	return nil
}
func (a *jsapi) BoldColour(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.BoldColour(int(call.GetArg(0).Integer()))))

}
func (a *jsapi) NormalColour(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.NormalColour(int(call.GetArg(0).Integer()))))
}

func (a *jsapi) GetStyleInfo(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	val, ok := a.API.GetStyleInfo(int(call.GetArg(0).Integer()), int(call.GetArg(1).Integer()), int(call.GetArg(2).Integer()))
	if !ok {
		return nil
	}
	switch int(call.GetArg(2).Integer()) {
	case 1:
		return call.Local().NewString(val)
	case 2:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))
	case 3:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))
	case 8:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 9:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 10:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 11:
		return call.Local().NewBoolean(world.FromStringBool(val))
	case 14:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))
	case 15:
		return call.Local().NewInt32(int32(world.FromStringInt(val)))

	}
	return nil
}

func (a *jsapi) GetInfo(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.GetInfo(int(call.GetArg(0).Integer())))
}
func (a *jsapi) GetTimerInfo(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	v, ok := a.API.GetTimerInfo(call.GetArg(0).String(), int(call.GetArg(1).Integer()))
	if ok != api.EOK {
		return nil
	}
	switch call.GetArg(1).Integer() {
	case 1:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 2:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 3:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 4:
		return call.Local().NewString(v)
	case 5:
		return call.Local().NewString(v)
	case 6:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 7:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 8:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 14:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 19:
		return call.Local().NewString(v)
	case 20:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 21:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 22:
		return call.Local().NewString(v)
	case 23:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 24:
		return call.Local().NewBoolean(world.FromStringBool(v))
	}
	return nil
}
func (a *jsapi) GetTriggerInfo(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	v, ok := a.API.GetTriggerInfo(call.GetArg(0).String(), int(call.GetArg(1).Integer()))
	if ok != api.EOK {
		return nil
	}
	switch call.GetArg(1).Integer() {
	case 1:
		return call.Local().NewString(v)
	case 2:
		return call.Local().NewString(v)
	case 3:
		return call.Local().NewString(v)
	case 4:
		return call.Local().NewString(v)
	case 5:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 6:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 7:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 8:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 9:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 10:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 11:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 13:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 15:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 16:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 23:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 25:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 26:
		return call.Local().NewString(v)
	case 27:
		return call.Local().NewString(v)
	case 28:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 31:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 36:
		return call.Local().NewBoolean(world.FromStringBool(v))
	}
	return nil
}

func (a *jsapi) GetAliasInfo(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	v, ok := a.API.GetAliasInfo(call.GetArg(0).String(), int(call.GetArg(1).Integer()))
	if ok != api.EOK {
		return nil
	}
	switch call.GetArg(1).Integer() {
	case 1:
		return call.Local().NewString(v)
	case 2:
		return call.Local().NewString(v)
	case 3:
		return call.Local().NewString(v)
	case 4:
		return call.Local().NewString(v)
	case 5:
		return call.Local().NewString(v)
	case 6:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 7:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 8:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 9:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 14:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 16:
		return call.Local().NewString(v)
	case 17:
		return call.Local().NewString(v)
	case 18:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 19:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 20:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 22:
		return call.Local().NewBoolean(world.FromStringBool(v))
	case 23:
		return call.Local().NewInt32(int32(world.FromStringInt(v)))
	case 29:
		return call.Local().NewBoolean(world.FromStringBool(v))

	}
	return nil
}

func (a *jsapi) Broadcast(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.Broadcast(call.GetArg(0).String(), call.GetArg(1).Boolean())
	return nil
}
func (a *jsapi) Notify(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	var link *string
	if call.GetArg(2).IsNullOrUndefined() {
		link = nil
	} else {
		data := call.GetArg(2).String()
		link = &data
	}
	a.API.Notify(call.GetArg(0).String(), call.GetArg(1).String(), link)
	return nil
}
func (a *jsapi) GetGlobalOption(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	result := a.API.GetGlobalOption(call.GetArg(0).String())
	switch call.GetArg(0).String() {
	default:
		switch result {
		case "0":
			return call.Local().NewInt32(int32(0))
		case "1":
			return call.Local().NewInt32(int32(1))
		default:
			return call.Local().NewString(result)
		}
	}
}

func (a *jsapi) CheckPermissions(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	items := call.GetArg(0).StringArrry()
	return call.Local().NewBoolean(a.API.CheckPermissions(items))
}
func (a *jsapi) RequestPermissions(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	items := call.GetArg(0).StringArrry()
	var reason = ""
	if !call.GetArg(1).IsUndefined() {
		reason = call.GetArg(1).String()
	}
	var script = ""
	if call.GetArg(2).IsUndefined() {
		script = call.GetArg(2).String()
	}
	a.API.RequestPermissions(items, reason, script)
	return nil
}
func (a *jsapi) CheckTrustedDomains(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	items := call.GetArg(0).StringArrry()
	return call.Local().NewBoolean(a.API.CheckTrustedDomains(items))
}

func (a *jsapi) RequestTrustDomains(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	items := call.GetArg(0).StringArrry()
	var reason = ""
	if !call.GetArg(1).IsUndefined() {
		reason = call.GetArg(1).String()
	}
	var script = ""
	if !call.GetArg(2).IsUndefined() {
		script = call.GetArg(2).String()
	}
	a.API.RequestTrustDomains(items, reason, script)
	return nil
}
func (a *jsapi) Encrypt(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	data := call.GetArg(0).String()
	key := call.GetArg(1).String()
	result := a.API.Encrypt(data, key)
	if result == nil {
		return nil
	}
	return call.Local().NewString(*result)
}
func (a *jsapi) Decrypt(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	data := call.GetArg(0).String()
	key := call.GetArg(1).String()
	result := a.API.Decrypt(data, key)
	if result == nil {
		return nil
	}
	return call.Local().NewString(*result)
}

func (a *jsapi) DumpOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	length := int(call.GetArg(0).Integer())
	offset := int(call.GetArg(1).Integer())
	return call.Local().NewString(a.API.DumpOutput(length, offset))
}

func (a *jsapi) ConcatOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	output1 := call.GetArg(0).String()
	output2 := call.GetArg(1).String()
	return call.Local().NewString(a.API.ConcatOutput(output1, output2))
}
func (a *jsapi) SliceOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	output := call.GetArg(0).String()
	start := int(call.GetArg(1).Integer())
	end := int(call.GetArg(2).Integer())
	return call.Local().NewString(a.API.SliceOutput(output, start, end))
}
func (a *jsapi) OutputToText(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	output := call.GetArg(0).String()
	return call.Local().NewString(a.API.OutputToText(output))
}
func (a *jsapi) FormatOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	output := call.GetArg(0).String()
	return call.Local().NewString(a.API.FormatOutput(output))
}
func (a *jsapi) PrintOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	output := call.GetArg(0).String()
	return call.Local().NewString(a.API.PrintOutput(output))
}
func (a *jsapi) Simulate(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	text := call.GetArg(0).String()
	a.API.Simulate(text)
	return nil
}
func (a *jsapi) SimulateOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	output := call.GetArg(0).String()
	a.API.SimulateOutput(output)
	return nil
}

func (a *jsapi) DumpTriggers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	byUser := call.GetArg(0).Boolean()
	return call.Local().NewString(a.API.DumpTriggers(byUser))
}
func (a *jsapi) RestoreTriggers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	data := call.GetArg(0).String()
	byUser := call.GetArg(1).Boolean()
	a.API.RestoreTriggers(data, byUser)
	return nil
}
func (a *jsapi) DumpTimers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	byUser := call.GetArg(0).Boolean()

	return call.Local().NewString(a.API.DumpTimers(byUser))
}
func (a *jsapi) RestoreTimers(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	data := call.GetArg(0).String()
	byUser := call.GetArg(1).Boolean()
	a.API.RestoreTimers(data, byUser)
	return nil
}
func (a *jsapi) DumpAliases(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	byUser := call.GetArg(0).Boolean()
	return call.Local().NewString(a.API.DumpAliases(byUser))
}
func (a *jsapi) RestoreAliases(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	data := call.GetArg(0).String()
	byUser := call.GetArg(1).Boolean()
	a.API.RestoreAliases(data, byUser)
	return nil
}
func (a *jsapi) SetHUDSize(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	size := call.GetArg(0).Integer()
	a.API.SetHUDSize(int(size))
	return nil
}
func (a *jsapi) GetHUDContent(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	content := a.API.GetHUDContent()
	return call.Local().NewString(content)
}
func (a *jsapi) GetHUDSize(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	size := a.API.GetHUDSize()
	return call.Local().NewInt32(int32(size))
}
func (a *jsapi) UpdateHUD(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	start := call.GetArg(0).Integer()
	content := call.GetArg(1).String()
	result := a.API.UpdateHUD(int(start), content)
	return call.Local().NewBoolean(result)
}
func (a *jsapi) NewLine(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.NewLine())
}
func (a *jsapi) NewWord(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	text := call.GetArg(0).String()
	return call.Local().NewString(a.API.NewWord(text))
}

func (a *jsapi) SetPriority(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	value := int(call.GetArg(0).Integer())
	a.API.SetPriority(value)
	return nil
}
func (a *jsapi) GetPriority(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt32(int32(a.API.GetPriority()))
}
func (a *jsapi) SetSummary(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	content := call.GetArg(0).String()
	a.API.SetSummary(content)
	return nil
}
func (a *jsapi) GetSummary(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewString(a.API.GetSummary())
}
func (a *jsapi) Save(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewBoolean(a.API.Save())
}
func (a *jsapi) Milliseconds(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewInt64(int64(a.API.Milliseconds()))
}

func (a *jsapi) OmitOutput(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	a.API.OmitOutput()
	return nil
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
