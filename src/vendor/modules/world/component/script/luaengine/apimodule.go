package luaengine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/api"

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

func (a *luaapi) InstallAPIs(l *lua.LState) {
	l.SetGlobal("Note", l.NewFunction(a.Note))
	l.SetGlobal("SendImmediate", l.NewFunction(a.SendImmediate))
	l.SetGlobal("Send", l.NewFunction(a.Send))
	l.SetGlobal("SendNoEcho", l.NewFunction(a.SendNoEcho))
	l.SetGlobal("GetVariable", l.NewFunction(a.GetVariable))
	l.SetGlobal("SetVariable", l.NewFunction(a.SetVariable))
	l.SetGlobal("DeleteVariable", l.NewFunction(a.DeleteVariable))
	l.SetGlobal("GetVariableList", l.NewFunction(a.GetVariableList))
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
	l.SetGlobal("Trim", l.NewFunction(a.Trim))
	l.SetGlobal("GetUniqueNumber", l.NewFunction(a.GetUniqueNumber))
	l.SetGlobal("GetUniqueID", l.NewFunction(a.GetUniqueID))
	l.SetGlobal("CreateGUID", l.NewFunction(a.CreateGUID))
	l.SetGlobal("FlashIcon", l.NewFunction(a.FlashIcon))
	l.SetGlobal("SetStatus", l.NewFunction(a.SetStatus))
	l.SetGlobal("Execute", l.NewFunction(a.Execute))
	l.SetGlobal("DeleteCommandHistory", l.NewFunction(a.DeleteCommandHistory))
	l.SetGlobal("DiscardQueue", l.NewFunction(a.DiscardQueue))
	l.SetGlobal("GetQueue", l.NewFunction(a.GetQueue))
}
func (a *luaapi) Note(L *lua.LState) int {
	info := L.ToString(1)
	a.API.Note(info)
	return 0
}
func (a *luaapi) SendImmediate(L *lua.LState) int {
	info := L.ToString(1)
	L.Push(lua.LNumber(a.API.SendImmediate(info)))
	return 1
}
func (a *luaapi) Send(L *lua.LState) int {
	info := L.ToString(1)
	L.Push(lua.LNumber(a.API.Send(info)))
	return 1
}
func (a *luaapi) Execute(L *lua.LState) int {
	info := L.ToString(1)
	L.Push(lua.LNumber(a.API.Execute(info)))
	return 1
}
func (a *luaapi) SendNoEcho(L *lua.LState) int {
	info := L.ToString(1)
	L.Push(lua.LNumber(a.API.SendNoEcho(info)))
	return 1
}
func (a *luaapi) GetVariable(L *lua.LState) int {
	name := L.ToString(1)
	val := a.API.GetVariable(name)
	if val == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(*val))
	}
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
	text := L.ToString(1)
	a.API.SetStatus(text)
	return 0
}
func (a *luaapi) DeleteCommandHistory(L *lua.LState) int {
	a.API.DeleteCommandHistory()
	return 0
}
func (a *luaapi) DiscardQueue(L *lua.LState) int {
	L.Push(lua.LNumber(a.API.DiscardQueue()))
	return 1
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
func NewAPIModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("worldapi",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			createApi(b).InstallAPIs(l)
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
