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

	"github.com/jarlyyn/v8js"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/uniqueid"
	"github.com/herb-go/v8go"
	"github.com/jarlyyn/v8js/v8plugin"
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

func AppendToWorld(ctx *v8js.Context, world *v8js.JsValue, name string, call func(call *v8js.FunctionCallbackInfo) *v8js.Consumed) {
	fn := ctx.NewFunctionTemplate(call)
	global := ctx.Global()
	global.Set(name, fn.GetFunction(ctx).Consume())

	world.Set(strings.ToLower(name), fn.GetFunction(ctx).Consume())
	if name != strings.ToLower(name) {
		world.Set(name, fn.GetFunction(ctx).Consume())
	}
	global.Release()
}
func (a *jsapi) InstallAPIs(p herbplugin.Plugin) {
	a.Plugin = p
	jp := p.(*v8plugin.Plugin)
	world := jp.Runtime.NewObject()

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

	AppendToWorld(jp.Runtime, world, "HasFile", a.NewHasFileAPI)
	AppendToWorld(jp.Runtime, world, "ReadFile", a.NewReadFileAPI)
	AppendToWorld(jp.Runtime, world, "ReadLines", a.NewReadLinesAPI)

	AppendToWorld(jp.Runtime, world, "HasModFile", a.NewHasModFileAPI)
	AppendToWorld(jp.Runtime, world, "ReadModFile", a.NewReadModFileAPI)
	AppendToWorld(jp.Runtime, world, "ReadModLines", a.NewReadModLinesAPI)
	AppendToWorld(jp.Runtime, world, "GetModInfo", a.NewGetModInfoAPI)

	AppendToWorld(jp.Runtime, world, "MakeHomeFolder", a.NewMakeHomeFolderAPI)

	AppendToWorld(jp.Runtime, world, "HasHomeFile", a.NewHasHomeFileAPI)
	AppendToWorld(jp.Runtime, world, "ReadHomeFile", a.NewReadHomeFileAPI)
	AppendToWorld(jp.Runtime, world, "ReadHomeLines", a.NewReadHomeLinesAPI)
	AppendToWorld(jp.Runtime, world, "WriteHomeFile", a.NewWriteHomeFileAPI)

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
	AppendToWorld(jp.Runtime, world, "Milliseconds", a.Milliseconds)

	AppendToWorld(jp.Runtime, world, "OmitOutput", a.OmitOutput)
	AppendToWorld(jp.Runtime, world, "PrintSystem", a.PrintSystem)

	AppendToWorld(jp.Runtime, world, "V8Debug", a.Debug)
	AppendToWorld(jp.Runtime, world, "Snapshot", a.Snapshot)
	global := jp.Runtime.Global()
	global.Set("world", world.Consume())
	global.Release()

}
func (a *jsapi) Snapshot(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	p := filepath.Join(a.API.Bus.GetLogsPath(), a.API.Bus.ID+"."+uniqueid.MustGenerateID()+".heapsnapshot")
	v8go.WriteHeapSnapshot(call.Context().Raw.Isolate(), p)
	a.API.Note("镜像文件写入" + p)
	return nil
}

func (a *jsapi) Debug(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	debug.FreeOSMemory()
	v8go.ForceV8GC(call.Context().Raw.Isolate())
	a.API.Note(strconv.Itoa(call.Context().Raw.RetainedValueCount()))
	bs, err := json.Marshal(call.Context().Raw.Isolate().GetHeapStatistics())
	if err != nil {
		panic(err)
	}
	a.API.Note("V8版本:" + v8go.Version())
	a.API.Note("V8内存统计:" + string(bs))
	return nil
}
func (a *jsapi) Print(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	msg := []string{}
	for _, v := range call.Args() {
		msg = append(msg, v.String())
	}
	a.API.Note(strings.Join(msg, " "))
	return nil
}

func (a *jsapi) Request(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	msgtype := call.GetArg(0).String()
	msg := call.GetArg(1).String()
	id := a.API.Request(msgtype, msg)
	return call.Context().NewString(id).Consume()
}
func (a *jsapi) Note(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	info := call.GetArg(0).String()
	a.API.Note(info)
	return nil
}
func (a *jsapi) PrintSystem(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	info := call.GetArg(0).String()
	a.API.PrintSystem(info)
	return nil
}
func (a *jsapi) SendImmediate(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	info := call.GetArg(0).String()

	return call.Context().NewInt32(int32(a.API.SendImmediate(info))).Consume()

}
func (a *jsapi) Send(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	info := call.GetArg(0).String()

	res := a.API.Send(info)
	return call.Context().NewInt32(int32(res)).Consume()
}
func (a *jsapi) Execute(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	info := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.Execute(info))).Consume()
}
func (a *jsapi) SendNoEcho(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	info := call.GetArg(0).String()

	return call.Context().NewInt32(int32(a.API.SendNoEcho(info))).Consume()
}
func (a *jsapi) GetVariable(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	val := a.API.GetVariable(call.GetArg(0).String())
	return call.Context().NewString(val).Consume()
}
func (a *jsapi) DeleteVariable(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteVariable(name))).Consume()
}
func (a *jsapi) SetVariable(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	value := call.GetArg(1).String()
	return call.Context().NewInt32(int32(a.API.SetVariable(name, value))).Consume()
}
func (a *jsapi) GetVariableList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	list := a.API.GetVariableList()
	result := make([]string, len(list))
	for k := range list {
		result = append(result, k)
	}
	return call.Context().NewStringArray(result...).Consume()
}
func (a *jsapi) GetVariableComment(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	val := a.API.GetVariableComment(call.GetArg(0).String())
	return call.Context().NewString(val).Consume()
}
func (a *jsapi) SetVariableComment(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	value := call.GetArg(1).String()
	return call.Context().NewInt32(int32(a.API.SetVariableComment(name, value))).Consume()
}
func (a *jsapi) Version(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.Version()).Consume()
}
func (a *jsapi) Hash(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewString(a.API.Hash(name)).Consume()
}
func (a *jsapi) Base64Encode(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	src := call.GetArg(0).String()
	ok := call.GetArg(1).Boolean()
	return call.Context().NewString(a.API.Base64Encode(src, ok)).Consume()
}
func (a *jsapi) Base64Decode(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	src := call.GetArg(0).String()
	result := a.API.Base64Decode(src)
	if result == nil {
		return nil
	}
	return call.Context().NewString(*result).Consume()
}
func (a *jsapi) Connect(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.Connect())).Consume()
}
func (a *jsapi) IsConnected(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewBoolean(a.API.IsConnected()).Consume()

}
func (a *jsapi) Disconnect(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.Disconnect())).Consume()
}

func (a *jsapi) GetWorldById(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return nil
}

func (a *jsapi) GetWorld(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return nil
}

func (a *jsapi) GetWorldID(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.GetWorldID()).Consume()

}
func (a *jsapi) GetWorldIdList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewArray().Consume()
}
func (a *jsapi) GetWorldList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewArray().Consume()
}
func (a *jsapi) WorldName(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.WorldName()).Consume()
}
func (a *jsapi) WorldAddress(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.WorldAddress()).Consume()
}
func (a *jsapi) WorldPort(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.WorldPort())).Consume()
}
func (a *jsapi) WorldProxy(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.WorldProxy()).Consume()
}

func (a *jsapi) Trim(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	src := call.GetArg(0).String()
	return call.Context().NewString(a.API.Trim(src)).Consume()
}
func (a *jsapi) GetUniqueNumber(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt64(int64(a.API.GetUniqueNumber())).Consume()
}
func (a *jsapi) GetUniqueID(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.GetUniqueID()).Consume()
}
func (a *jsapi) CreateGUID(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.CreateGUID()).Consume()
}
func (a *jsapi) FlashIcon(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.FlashIcon()
	return nil
}
func (a *jsapi) SetStatus(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	text := call.GetArg(0).String()
	a.API.SetStatus(text)
	return nil
}
func (a *jsapi) DeleteCommandHistory(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.DeleteCommandHistory()
	return nil
}
func (a *jsapi) DiscardQueue(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.DiscardQueue(call.GetArg(0).Boolean()))).Consume()
}
func (a *jsapi) LockQueue(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	a.API.LockQueue()
	return nil
}
func (a *jsapi) GetQueue(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	cmds := a.API.GetQueue()
	return call.Context().NewStringArray(cmds...).Consume()
}
func (a *jsapi) Queue(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.Queue(call.GetArg(0).String(), call.GetArg(1).Boolean()))).Consume()
}
func (a *jsapi) DoAfter(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	return call.Context().NewInt32(int32(a.API.DoAfter(seconds, send))).Consume()
}
func (a *jsapi) DoAfterNote(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	return call.Context().NewInt32(int32(a.API.DoAfterNote(seconds, send))).Consume()

}
func (a *jsapi) DoAfterSpeedWalk(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	return call.Context().NewInt32(int32(a.API.DoAfterSpeedWalk(seconds, send))).Consume()
}
func (a *jsapi) DoAfterSpecial(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	seconds := call.GetArg(0).Number()
	send := call.GetArg(1).String()
	sendto := int(call.GetArg(2).Integer())
	return call.Context().NewInt32(int32(a.API.DoAfterSpecial(seconds, send, sendto))).Consume()

}

func (a *jsapi) DeleteGroup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteGroup(name))).Consume()
}

func (a *jsapi) AddTimer(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	hour := int(call.GetArg(1).Integer())
	min := int(call.GetArg(2).Integer())
	seconds := call.GetArg(3).Number()
	send := call.GetArg(4).String()
	flags := int(call.GetArg(5).Integer())
	script := call.GetArg(6).String()
	return call.Context().NewInt32(int32(a.API.AddTimer(name, hour, min, seconds, send, flags, script))).Consume()
}
func (a *jsapi) DeleteTimer(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteTimer(name))).Consume()

}
func (a *jsapi) DeleteTemporaryTimers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.DeleteTemporaryTimers())).Consume()

}
func (a *jsapi) DeleteTimerGroup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteTimerGroup(name))).Consume()
}

func (a *jsapi) EnableTimer(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Context().NewInt32(int32(a.API.EnableTimer(name, enabled))).Consume()
}
func (a *jsapi) EnableTimerGroup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	group := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Context().NewInt32(int32(a.API.EnableTimerGroup(group, enabled))).Consume()
}

func (a *jsapi) GetTimerList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	list := a.API.GetTimerList()
	result := []*v8js.Consumed{}
	for _, v := range list {
		result = append(result, call.Context().NewString(v).Consume())
	}
	return call.Context().NewArray(result...).Consume()
}
func (a *jsapi) IsTimer(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.IsTimer(name))).Consume()
}

func (a *jsapi) ResetTimer(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.ResetTimer(name))).Consume()
}

func (a *jsapi) ResetTimers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.ResetTimers()
	return nil
}

func (a *jsapi) GetTimerOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		return nil
	} else {
		switch option {
		case "active_closed", "at_time", "enabled", "omit_from_log", "omit_from_output", "one_shot":
			return call.Context().NewBoolean(result == world.StringYes).Consume()
		case "group", "name", "script", "send", "variable":
			return call.Context().NewString(result).Consume()
		case "hour", "minute", "offset_hour", "offset_minute", "offset_second", "send_to", "user":
			i, _ := strconv.Atoi(result)
			return call.Context().NewInt32(int32(i)).Consume()
		case "second":
			i, _ := strconv.ParseFloat(result, 64)
			return call.Context().NewNumber(i).Consume()
		}
	}
	return nil
}
func (a *jsapi) SetTimerOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
	return call.Context().NewInt32(int32(a.API.SetTimerOption(name, option, value))).Consume()
}

func (a *jsapi) AddAlias(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	match := call.GetArg(1).String()
	send := call.GetArg(2).String()
	flags := int(call.GetArg(3).Integer())
	script := call.GetArg(4).String()
	return call.Context().NewInt32(int32(a.API.AddAlias(name, match, send, flags, script))).Consume()
}
func (a *jsapi) DeleteAlias(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteAlias(name))).Consume()
}
func (a *jsapi) DeleteTemporaryAliases(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.DeleteTemporaryAliases())).Consume()

}
func (a *jsapi) DeleteAliasGroup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteAliasGroup(name))).Consume()
}

func (a *jsapi) EnableAlias(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Context().NewInt32(int32(a.API.EnableAlias(name, enabled))).Consume()
}
func (a *jsapi) EnableAliasGroup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	group := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Context().NewInt32(int32(a.API.EnableAliasGroup(group, enabled))).Consume()
}

func (a *jsapi) GetAliasList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	list := a.API.GetAliasList()
	result := []*v8js.Consumed{}
	for _, v := range list {
		result = append(result, call.Context().NewString(v).Consume())
	}
	return call.Context().NewArray(result...).Consume()
}
func (a *jsapi) IsAlias(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.IsAlias(name))).Consume()
}

func (a *jsapi) GetAliasOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	result, code := a.API.GetTimerOption(name, option)
	if code != api.EOK {
		return nil
	} else {
		switch option {
		case "echo_alias", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			return call.Context().NewBoolean(result == world.StringYes).Consume()
		case "group", "name", "match", "script", "send", "variable":
			return call.Context().NewString(result).Consume()
		case "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			return call.Context().NewInt32(int32(i)).Consume()
		}
	}
	return nil
}
func (a *jsapi) SetAliasOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
	return call.Context().NewInt32(int32(a.API.SetAliasOption(name, option, value))).Consume()
}

func (a *jsapi) AddTrigger(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	match := call.GetArg(1).String()
	send := call.GetArg(2).String()
	flags := int(call.GetArg(3).Integer())
	color := int(call.GetArg(4).Integer())
	wildcard := int(call.GetArg(5).Integer())
	sound := call.GetArg(6).String()
	script := call.GetArg(7).String()
	return call.Context().NewInt32(int32(a.API.AddTrigger(name, match, send, flags, color, wildcard, sound, script))).Consume()
}
func (a *jsapi) AddTriggerEx(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
	return call.Context().NewInt32(int32(a.API.AddTriggerEx(name, match, send, flags, color, wildcard, sound, script, sendto, sequence))).Consume()
}
func (a *jsapi) DeleteTrigger(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteTrigger(name))).Consume()
}
func (a *jsapi) DeleteTemporaryTriggers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.DeleteTemporaryTimers())).Consume()

}
func (a *jsapi) DeleteTriggerGroup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.DeleteTriggerGroup(name))).Consume()
}

func (a *jsapi) EnableTrigger(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Context().NewInt32(int32(a.API.EnableTrigger(name, enabled))).Consume()
}
func (a *jsapi) EnableTriggerGroup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	group := call.GetArg(0).String()
	enabled := call.GetArg(1).Boolean()
	return call.Context().NewInt32(int32(a.API.EnableTriggerGroup(group, enabled))).Consume()
}

func (a *jsapi) GetTriggerList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	list := a.API.GetTriggerList()
	result := []*v8js.Consumed{}
	for _, v := range list {
		result = append(result, call.Context().NewString(v).Consume())
	}
	return call.Context().NewArray(result...).Consume()
}
func (a *jsapi) IsTrigger(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.IsTrigger(name))).Consume()
}

func (a *jsapi) GetTriggerOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	option := call.GetArg(1).String()
	result, code := a.API.GetTriggerOption(name, option)
	if code != api.EOK {
		return nil
	} else {
		switch option {
		case "echo_trigger", "enabled", "expand_variables", "ignore_case", "keep_evaluating", "menu", "omit_from_command_history", "regexp", "omit_from_log", "omit_from_output", "one_shot":
			return call.Context().NewBoolean(result == world.StringYes).Consume()
		case "group", "name", "match", "script", "send", "variable":
			return call.Context().NewString(result).Consume()
		case "send_to", "user", "sequence":
			i, _ := strconv.Atoi(result)
			return call.Context().NewInt32(int32(i)).Consume()
		}
	}
	return nil
}
func (a *jsapi) SetTriggerOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
	return call.Context().NewInt32(int32(a.API.SetTriggerOption(name, option, value))).Consume()
}

func (a *jsapi) StopEvaluatingTriggers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.StopEvaluatingTriggers()
	return nil
}
func (a *jsapi) GetTriggerWildcard(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	result := a.API.GetTriggerWildcard(call.GetArg(0).String(), call.GetArg(1).String())
	if result == nil {
		return nil
	}
	return call.Context().NewString(*result).Consume()
}

func (a *jsapi) ColourNameToRGB(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	v := a.API.ColourNameToRGB(call.GetArg(0).String())
	return call.Context().NewInt32(int32(v)).Consume()
}
func (a *jsapi) SetSpeedWalkDelay(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.SetSpeedWalkDelay(int(call.GetArg(0).Integer()))
	return nil
}
func (a *jsapi) GetSpeedWalkDelay(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.SpeedWalkDelay())).Consume()
}

func (a *jsapi) NewGetModInfoAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	mod := a.API.GetModInfo(a.Plugin)
	result := call.Context().NewObject()
	result.Set("Enabled", call.Context().NewBoolean(mod.Enabled).Consume())
	result.Set("Exists", call.Context().NewBoolean(mod.Exists).Consume())
	result.Set("FileList", call.Context().NewStringArray(mod.FileList...).Consume())
	result.Set("FolderList", call.Context().NewStringArray(mod.FolderList...).Consume())
	return result.Consume()
}
func (a *jsapi) NewHasFileAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewBoolean(a.API.HasFile(a.Plugin, call.GetArg(0).String())).Consume()
}
func (a *jsapi) NewReadFileAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return call.Context().NewString(a.API.ReadFile(a.Plugin, call.GetArg(0).String())).Consume()
}
func (a *jsapi) NewHasModFileAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewBoolean(a.API.HasModFile(a.Plugin, call.GetArg(0).String())).Consume()
}
func (a *jsapi) NewMakeHomeFolderAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewBoolean(a.API.MakeHomeFolder(a.Plugin, call.GetArg(0).String())).Consume()
}
func (a *jsapi) NewHasHomeFileAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return call.Context().NewBoolean(a.API.HasHomeFile(a.Plugin, call.GetArg(0).String())).Consume()
}
func (a *jsapi) NewWriteHomeFileAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.WriteHomeFile(a.Plugin, call.GetArg(0).String(), []byte(call.GetArg(1).String()))
	return nil
}
func (a *jsapi) NewReadModFileAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return call.Context().NewString(a.API.ReadModFile(a.Plugin, call.GetArg(0).String())).Consume()
}
func (a *jsapi) NewReadHomeFileAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return call.Context().NewString(a.API.ReadHomeFile(a.Plugin, call.GetArg(0).String())).Consume()
}
func (a *jsapi) NewReadLinesAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	lines := a.API.ReadLines(a.Plugin, call.GetArg(0).String())
	t := []*v8js.Consumed{}
	for _, v := range lines {
		t = append(t, call.Context().NewString(v).Consume())
	}
	return call.Context().NewArray(t...).Consume()

}
func (a *jsapi) NewReadModLinesAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	lines := a.API.ReadModLines(a.Plugin, call.GetArg(0).String())
	t := []*v8js.Consumed{}
	for _, v := range lines {
		t = append(t, call.Context().NewString(v).Consume())
	}
	return call.Context().NewArray(t...).Consume()
}
func (a *jsapi) NewReadHomeLinesAPI(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	lines := a.API.ReadHomeLines(a.Plugin, call.GetArg(0).String())
	t := []*v8js.Consumed{}
	for _, v := range lines {
		t = append(t, call.Context().NewString(v).Consume())
	}
	return call.Context().NewArray(t...).Consume()
}
func (a *jsapi) SplitNfunc(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	text := call.GetArg(0).String()
	sep := call.GetArg(1).String()
	n := int(call.GetArg(2).Integer())
	s := a.API.SplitN(text, sep, n)
	return call.Context().NewStringArray(s...).Consume()
}

func (a *jsapi) UTF8Len(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	text := call.GetArg(0).String()
	return call.Context().NewInt32(int32(a.API.UTF8Len(text))).Consume()
}
func (a *jsapi) UTF8Index(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	text := call.GetArg(0).String()
	sub := call.GetArg(1).String()
	return call.Context().NewInt32(int32(a.API.UTF8Index(text, sub))).Consume()
}
func (a *jsapi) ToUTF8(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	code := call.GetArg(0).String()
	text := call.GetArg(1).String()
	result := a.API.ToUTF8(code, text)
	if result == nil {
		return nil
	}
	return call.Context().NewString(*result).Consume()
}
func (a *jsapi) FromUTF8(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	code := call.GetArg(0).String()
	text := call.GetArg(1).String()
	result := a.API.FromUTF8(code, text)
	if result == nil {
		return nil
	}
	return call.Context().NewString(*result).Consume()
}
func (a *jsapi) UTF8Sub(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	text := call.GetArg(0).String()
	start := int(call.GetArg(1).Integer())
	end := int(call.GetArg(2).Integer())
	return call.Context().NewString(a.API.UTF8Sub(text, start, end)).Consume()
}
func (a *jsapi) Info(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	text := call.GetArg(0).String()
	a.API.Info(text)
	return nil
}
func (a *jsapi) InfoClear(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.InfoClear()
	return nil
}

func (a *jsapi) GetAlphaOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.GetAlphaOption(call.GetArg(0).String())).Consume()
}

func (a *jsapi) SetAlphaOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.SetAlphaOption(call.GetArg(0).String(), call.GetArg(1).String()))).Consume()
}
func (a *jsapi) WriteLog(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.WriteLog(call.GetArg(0).String()))).Consume()
}

func (a *jsapi) CloseLog(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.CloseLog())).Consume()
}
func (a *jsapi) OpenLog(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.OpenLog())).Consume()
}
func (a *jsapi) FlushLog(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.FlushLog())).Consume()
}

func (a *jsapi) GetLinesInBufferCount(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.GetLinesInBufferCount())).Consume()
}
func (a *jsapi) DeleteOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.DeleteOutput()
	return nil
}
func (a *jsapi) DeleteLines(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.DeleteLines(int(call.GetArg(0).Integer()))
	return nil
}
func (a *jsapi) GetLineCount(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.GetLineCount())).Consume()
}
func (a *jsapi) GetRecentLines(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.GetRecentLines(int(call.GetArg(0).Integer()))).Consume()
}
func (a *jsapi) GetLineInfo(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	val, ok := a.API.GetLineInfo(int(call.GetArg(0).Integer()), int(call.GetArg(1).Integer()))
	if !ok {
		return nil
	}
	switch int(call.GetArg(1).Integer()) {
	case 1:
		return call.Context().NewString(val).Consume()
	case 2:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()
	case 3:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()
	case 4:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 5:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 6:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 7:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 8:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 9:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()
	case 11:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()
	}
	return nil
}
func (a *jsapi) BoldColour(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.BoldColour(int(call.GetArg(0).Integer())))).Consume()

}
func (a *jsapi) NormalColour(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.NormalColour(int(call.GetArg(0).Integer())))).Consume()
}

func (a *jsapi) GetStyleInfo(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	val, ok := a.API.GetStyleInfo(int(call.GetArg(0).Integer()), int(call.GetArg(1).Integer()), int(call.GetArg(2).Integer()))
	if !ok {
		return nil
	}
	switch int(call.GetArg(2).Integer()) {
	case 1:
		return call.Context().NewString(val).Consume()
	case 2:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()
	case 3:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()
	case 8:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 9:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 10:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 11:
		return call.Context().NewBoolean(world.FromStringBool(val)).Consume()
	case 14:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()
	case 15:
		return call.Context().NewInt32(int32(world.FromStringInt(val))).Consume()

	}
	return nil
}

func (a *jsapi) GetInfo(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.GetInfo(int(call.GetArg(0).Integer()))).Consume()
}
func (a *jsapi) GetTimerInfo(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	v, ok := a.API.GetTimerInfo(call.GetArg(0).String(), int(call.GetArg(1).Integer()))
	if ok != api.EOK {
		return nil
	}
	switch call.GetArg(1).Integer() {
	case 1:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 2:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 3:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 4:
		return call.Context().NewString(v).Consume()
	case 5:
		return call.Context().NewString(v).Consume()
	case 6:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 7:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 8:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 14:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 19:
		return call.Context().NewString(v).Consume()
	case 20:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 21:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 22:
		return call.Context().NewString(v).Consume()
	case 23:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 24:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	}
	return nil
}
func (a *jsapi) GetTriggerInfo(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	v, ok := a.API.GetTriggerInfo(call.GetArg(0).String(), int(call.GetArg(1).Integer()))
	if ok != api.EOK {
		return nil
	}
	switch call.GetArg(1).Integer() {
	case 1:
		return call.Context().NewString(v).Consume()
	case 2:
		return call.Context().NewString(v).Consume()
	case 3:
		return call.Context().NewString(v).Consume()
	case 4:
		return call.Context().NewString(v).Consume()
	case 5:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 6:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 7:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 8:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 9:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 10:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 11:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 13:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 15:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 16:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 23:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 25:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 26:
		return call.Context().NewString(v).Consume()
	case 27:
		return call.Context().NewString(v).Consume()
	case 28:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 31:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 36:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	}
	return nil
}

func (a *jsapi) GetAliasInfo(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	v, ok := a.API.GetAliasInfo(call.GetArg(0).String(), int(call.GetArg(1).Integer()))
	if ok != api.EOK {
		return nil
	}
	switch call.GetArg(1).Integer() {
	case 1:
		return call.Context().NewString(v).Consume()
	case 2:
		return call.Context().NewString(v).Consume()
	case 3:
		return call.Context().NewString(v).Consume()
	case 4:
		return call.Context().NewString(v).Consume()
	case 5:
		return call.Context().NewString(v).Consume()
	case 6:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 7:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 8:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 9:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 14:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 16:
		return call.Context().NewString(v).Consume()
	case 17:
		return call.Context().NewString(v).Consume()
	case 18:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 19:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 20:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 22:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()
	case 23:
		return call.Context().NewInt32(int32(world.FromStringInt(v))).Consume()
	case 29:
		return call.Context().NewBoolean(world.FromStringBool(v)).Consume()

	}
	return nil
}

func (a *jsapi) Broadcast(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	a.API.Broadcast(call.GetArg(0).String(), call.GetArg(1).Boolean())
	return nil
}
func (a *jsapi) Notify(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
func (a *jsapi) GetGlobalOption(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	result := a.API.GetGlobalOption(call.GetArg(0).String())
	switch call.GetArg(0).String() {
	default:
		switch result {
		case "0":
			return call.Context().NewInt32(int32(0)).Consume()
		case "1":
			return call.Context().NewInt32(int32(1)).Consume()
		default:
			return call.Context().NewString(result).Consume()
		}
	}
}

func (a *jsapi) CheckPermissions(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	items := call.GetArg(0).StringArrry()
	return call.Context().NewBoolean(a.API.CheckPermissions(items)).Consume()
}
func (a *jsapi) RequestPermissions(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
func (a *jsapi) CheckTrustedDomains(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	items := call.GetArg(0).StringArrry()
	return call.Context().NewBoolean(a.API.CheckTrustedDomains(items)).Consume()
}

func (a *jsapi) RequestTrustDomains(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
func (a *jsapi) Encrypt(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	data := call.GetArg(0).String()
	key := call.GetArg(1).String()
	result := a.API.Encrypt(data, key)
	if result == nil {
		return nil
	}
	return call.Context().NewString(*result).Consume()
}
func (a *jsapi) Decrypt(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	data := call.GetArg(0).String()
	key := call.GetArg(1).String()
	result := a.API.Decrypt(data, key)
	if result == nil {
		return nil
	}
	return call.Context().NewString(*result).Consume()
}

func (a *jsapi) DumpOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	length := int(call.GetArg(0).Integer())
	offset := int(call.GetArg(1).Integer())
	return call.Context().NewString(a.API.DumpOutput(length, offset)).Consume()
}

func (a *jsapi) ConcatOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	output1 := call.GetArg(0).String()
	output2 := call.GetArg(1).String()
	return call.Context().NewString(a.API.ConcatOutput(output1, output2)).Consume()
}
func (a *jsapi) SliceOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	output := call.GetArg(0).String()
	start := int(call.GetArg(1).Integer())
	end := int(call.GetArg(2).Integer())
	return call.Context().NewString(a.API.SliceOutput(output, start, end)).Consume()
}
func (a *jsapi) OutputToText(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	output := call.GetArg(0).String()
	return call.Context().NewString(a.API.OutputToText(output)).Consume()
}
func (a *jsapi) FormatOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	output := call.GetArg(0).String()
	return call.Context().NewString(a.API.FormatOutput(output)).Consume()
}
func (a *jsapi) PrintOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	output := call.GetArg(0).String()
	return call.Context().NewString(a.API.PrintOutput(output)).Consume()
}
func (a *jsapi) Simulate(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	text := call.GetArg(0).String()
	a.API.Simulate(text)
	return nil
}
func (a *jsapi) SimulateOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	output := call.GetArg(0).String()
	a.API.SimulateOutput(output)
	return nil
}

func (a *jsapi) DumpTriggers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	byUser := call.GetArg(0).Boolean()
	return call.Context().NewString(a.API.DumpTriggers(byUser)).Consume()
}
func (a *jsapi) RestoreTriggers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	data := call.GetArg(0).String()
	byUser := call.GetArg(1).Boolean()
	a.API.RestoreTriggers(data, byUser)
	return nil
}
func (a *jsapi) DumpTimers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	byUser := call.GetArg(0).Boolean()

	return call.Context().NewString(a.API.DumpTimers(byUser)).Consume()
}
func (a *jsapi) RestoreTimers(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	data := call.GetArg(0).String()
	byUser := call.GetArg(1).Boolean()
	a.API.RestoreTimers(data, byUser)
	return nil
}
func (a *jsapi) DumpAliases(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	byUser := call.GetArg(0).Boolean()
	return call.Context().NewString(a.API.DumpAliases(byUser)).Consume()
}
func (a *jsapi) RestoreAliases(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	data := call.GetArg(0).String()
	byUser := call.GetArg(1).Boolean()
	a.API.RestoreAliases(data, byUser)
	return nil
}
func (a *jsapi) SetHUDSize(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	size := call.GetArg(0).Integer()
	a.API.SetHUDSize(int(size))
	return nil
}
func (a *jsapi) GetHUDContent(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	content := a.API.GetHUDContent()
	return call.Context().NewString(content).Consume()
}
func (a *jsapi) GetHUDSize(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	size := a.API.GetHUDSize()
	return call.Context().NewInt32(int32(size)).Consume()
}
func (a *jsapi) UpdateHUD(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	start := call.GetArg(0).Integer()
	content := call.GetArg(1).String()
	result := a.API.UpdateHUD(int(start), content)
	return call.Context().NewBoolean(result).Consume()
}
func (a *jsapi) NewLine(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.NewLine()).Consume()
}
func (a *jsapi) NewWord(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	text := call.GetArg(0).String()
	return call.Context().NewString(a.API.NewWord(text)).Consume()
}

func (a *jsapi) SetPriority(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	value := int(call.GetArg(0).Integer())
	a.API.SetPriority(value)
	return nil
}
func (a *jsapi) GetPriority(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(a.API.GetPriority())).Consume()
}
func (a *jsapi) SetSummary(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	content := call.GetArg(0).String()
	a.API.SetSummary(content)
	return nil
}
func (a *jsapi) GetSummary(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(a.API.GetSummary()).Consume()
}
func (a *jsapi) Save(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewBoolean(a.API.Save()).Consume()
}
func (a *jsapi) Milliseconds(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt64(int64(a.API.Milliseconds())).Consume()
}

func (a *jsapi) OmitOutput(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
