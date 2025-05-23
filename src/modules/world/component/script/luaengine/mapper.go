package luaengine

import (
	"context"
	"errors"
	"modules/mapper"
	"modules/world/bus"

	lua "github.com/yuin/gopher-lua"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
)

type LuaWalkAllResult struct {
	result *mapper.WalkAllResult
}

func (result *LuaWalkAllResult) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	steps := L.NewTable()
	for _, v := range result.result.Steps {
		s := &LuaStep{step: v}
		steps.Append(s.Convert(L))
	}
	walked := L.NewTable()
	for _, v := range result.result.Walked {
		walked.Append(lua.LString(v))
	}
	notwalked := L.NewTable()
	for _, v := range result.result.NotWalked {
		notwalked.Append(lua.LString(v))
	}
	t.RawSetString("steps", steps)
	t.RawSetString("walked", walked)
	t.RawSetString("notwalked", notwalked)
	return t
}

func ConvertLuaPath(v lua.LValue) *LuaPath {
	if v.Type() != lua.LTTable {
		return nil
	}
	t := v.(*lua.LTable)
	p := &LuaPath{
		path: mapper.NewPath(),
	}
	p.path.Command = t.RawGetString("command").String()
	p.path.From = t.RawGetString("from").String()
	p.path.To = t.RawGetString("to").String()
	p.path.Delay = int(lua.LVAsNumber(t.RawGetString("delay")))
	tags := t.RawGetString("tags")
	switch tags.Type() {
	case lua.LTNil:
	case lua.LTTable:
		t := tags.(*lua.LTable)
		max := t.MaxN()
		for i := 1; i <= max; i++ {
			p.path.Tags[lua.LVAsString(t.RawGetInt(i))] = true
		}
	default:
		panic(errors.New("tags must be table"))
	}
	etags := t.RawGetString("excludetags")
	switch etags.Type() {
	case lua.LTNil:
	case lua.LTTable:
		t := etags.(*lua.LTable)
		max := t.MaxN()
		for i := 1; i <= max; i++ {
			p.path.ExcludeTags[lua.LVAsString(t.RawGetInt(i))] = true
		}
	default:
		panic(errors.New("excludetags must be table"))
	}
	return p
}

type LuaPath struct {
	path *mapper.Path
}

func (p *LuaPath) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("command", lua.LString(p.path.Command))
	t.RawSetString("from", lua.LString(p.path.From))
	t.RawSetString("to", lua.LString(p.path.To))
	t.RawSetString("delay", lua.LNumber(p.path.Delay))
	tags := L.NewTable()
	for k, v := range p.path.Tags {
		if v {
			tags.Append(lua.LString(k))
		}
	}
	t.RawSetString("tags", tags)
	etags := L.NewTable()
	for k, v := range p.path.ExcludeTags {
		if v {
			etags.Append(lua.LString(k))
		}
	}
	t.RawSetString("excludetags", etags)
	return t

}

func ConvertLuaStep(v lua.LValue) *LuaStep {
	if v.Type() != lua.LTTable {
		return nil
	}
	t := v.(*lua.LTable)
	s := &LuaStep{
		step: &mapper.Step{},
	}
	s.step.Command = t.RawGetString("command").String()
	s.step.From = t.RawGetString("from").String()
	s.step.To = t.RawGetString("to").String()
	s.step.Delay = int(lua.LVAsNumber(t.RawGetString("delay")))

	return s
}

type LuaStep struct {
	step *mapper.Step
}

func (s *LuaStep) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("command", lua.LString(s.step.Command))
	t.RawSetString("from", lua.LString(s.step.From))
	t.RawSetString("to", lua.LString(s.step.To))
	t.RawSetString("delay", lua.LNumber(s.step.Delay))
	return t
}

type LuaMapper struct {
	mapper *mapper.Mapper
}

func (m *LuaMapper) Reset(L *lua.LState) int {
	m.mapper.Reset()
	return 0
}
func (m *LuaMapper) ResetTemporary(L *lua.LState) int {
	m.mapper.Reset()
	return 0
}

func (m *LuaMapper) AddTags(L *lua.LState) int {
	_ = L.Get(1) //self
	tags := []string{}
	count := L.GetTop()
	for i := 1; i < count; i++ {
		tags = append(tags, L.ToString(i+1))
	}
	m.mapper.AddTags(tags)
	return 0
}
func (m *LuaMapper) SetTag(L *lua.LState) int {
	_ = L.Get(1) //self
	tag := L.ToString(2)
	enabled := L.ToBool(3)
	m.mapper.SetTag(tag, enabled)
	return 0
}
func (m *LuaMapper) FlashTags(L *lua.LState) int {
	_ = L.Get(1) //self
	m.mapper.FlashTags()
	return 0
}

func (m *LuaMapper) SetTags(L *lua.LState) int {
	_ = L.Get(1) //self
	tags := L.Get(2)
	if tags.Type() != lua.LTTable {
		panic(errors.New("tags must be table"))
	}
	t := tags.(*lua.LTable)
	result := []string{}
	t.ForEach(func(k lua.LValue, v lua.LValue) {
		result = append(result, v.String())
	})
	m.mapper.AddTags(result)
	return 0
}

func (m *LuaMapper) Tags(L *lua.LState) int {
	result := m.mapper.Tags()
	t := L.NewTable()
	for k := range result {
		t.Append(lua.LString(result[k]))
	}
	L.Push(t)
	return 1
}
func (m *LuaMapper) toOption(L *lua.LState, option *lua.LTable) *mapper.Option {
	var opt *mapper.Option
	if option != nil {
		opt = mapper.NewOption()
		lblacklist, ok := option.RawGetString("blacklist").(*lua.LTable)
		if ok {
			lblacklist.ForEach(func(l1, l2 lua.LValue) {
				opt.Blacklist = append(opt.Blacklist, l2.String())
			})
		}
		lwhitelist, ok := option.RawGetString("whitelist").(*lua.LTable)
		if ok {
			lwhitelist.ForEach(func(l1, l2 lua.LValue) {
				opt.Whitelist = append(opt.Whitelist, l2.String())
			})
		}
		lblockedpath, ok := option.RawGetString("blockedpath").(*lua.LTable)
		if ok {
			lblockedpath.ForEach(func(l1, l2 lua.LValue) {
				path := []string{}
				lpath, ok2 := l2.(*lua.LTable)
				if ok2 {
					path = append(path, lpath.String())
				}
				if len(path) > 1 {
					opt.BlockedPath = append(opt.BlockedPath, path)
				}
			})
		}
	}
	return opt
}
func (m *LuaMapper) WalkAll(L *lua.LState) int {
	_ = L.Get(1) //self
	targets := []string{}
	ltargets := L.ToTable(2)
	ltargets.ForEach(func(l1, l2 lua.LValue) {
		targets = append(targets, l2.String())
	})
	fly := L.ToInt(3)
	maxdistance := L.ToInt(4)
	option := L.ToTable(5)
	opt := m.toOption(L, option)
	result := m.mapper.WalkAll(targets, fly != 0, maxdistance, opt)
	luaresult := &LuaWalkAllResult{result: result}
	L.Push(luaresult.Convert(L))
	return 1
}

func (m *LuaMapper) GetPath(L *lua.LState) int {
	_ = L.Get(1) //self
	from := L.ToString(2)
	fly := L.ToInt(3)
	count := L.GetTop()
	to := []string{}
	for i := 3; i < count; i++ {
		to = append(to, L.ToString(i+1))
	}
	option := L.ToTable(5)
	opt := m.toOption(L, option)
	steps := m.mapper.GetPath(from, fly != 0, to, opt)
	if steps == nil {
		L.Push(lua.LNil)
		return 1
	}
	t := L.NewTable()
	for i := range steps {
		s := &LuaStep{step: steps[i]}
		t.Append(s.Convert(L))
	}
	L.Push(t)
	return 1
}
func (m *LuaMapper) AddPath(L *lua.LState) int {
	_ = L.Get(1) //self
	id := L.ToString(2)
	path := ConvertLuaPath(L.Get(3))
	if path == nil {
		L.Push(lua.LBool(false))
		return 1
	}
	L.Push(lua.LBool(m.mapper.AddPath(id, path.path)))
	return 1
}
func (m *LuaMapper) AddTemporaryPath(L *lua.LState) int {
	_ = L.Get(1) //self
	id := L.ToString(2)
	path := ConvertLuaPath(L.Get(3))
	if path == nil {
		L.Push(lua.LBool(false))
		return 1
	}
	L.Push(lua.LBool(m.mapper.AddTemporaryPath(id, path.path)))
	return 1
}

func (m *LuaMapper) NewPath(L *lua.LState) int {
	p := &LuaPath{
		path: mapper.NewPath(),
	}
	L.Push(p.Convert(L))
	return 1
}
func (m *LuaMapper) GetRoomID(L *lua.LState) int {
	_ = L.Get(1) //self
	name := L.ToString(2)
	ids := m.mapper.GetRoomID(name)
	t := L.NewTable()
	for _, v := range ids {
		t.Append(lua.LString(v))
	}
	L.Push(t)
	return 1
}
func (m *LuaMapper) GetRoomName(L *lua.LState) int {
	_ = L.Get(1) //self
	id := L.ToString(2)
	L.Push(lua.LString(m.mapper.GetRoomName(id)))
	return 1
}
func (m *LuaMapper) RemoveRoom(L *lua.LState) int {
	_ = L.Get(1) //self
	id := L.ToString(2)
	L.Push(lua.LBool((m.mapper.RemoveRoom(id))))
	return 1
}
func (m *LuaMapper) SetRoomName(L *lua.LState) int {
	_ = L.Get(1) //self
	id := L.ToString(2)
	name := L.ToString(3)
	m.mapper.SetRoomName(id, name)
	return 0
}
func (m *LuaMapper) ClearRoom(L *lua.LState) int {
	_ = L.Get(1) //self
	id := L.ToString(2)
	m.mapper.ClearRoom(id)
	return 0
}
func (m *LuaMapper) NewArea(L *lua.LState) int {
	_ = L.Get(1) //self
	size := L.ToInt(2)
	ids := m.mapper.NewArea(size)
	t := L.NewTable()
	for _, v := range ids {
		t.Append(lua.LString(v))
	}
	L.Push(t)
	return 1
}
func (m *LuaMapper) GetExits(L *lua.LState) int {
	_ = L.Get(1) //self
	id := L.ToString(2)
	all := L.ToBool(3)
	exits := m.mapper.GetExits(id, all)
	t := L.NewTable()
	for _, v := range exits {
		p := &LuaPath{
			path: v,
		}
		t.Append(p.Convert(L))
	}
	L.Push(t)
	return 1

}

func (m *LuaMapper) SetFlyList(L *lua.LState) int {
	_ = L.Get(1) //self
	flv := L.ToTable(2)
	if flv == nil {
		panic(errors.New("flylist must be array"))
	}
	var result = []*mapper.Path{}
	flv.ForEach(func(key lua.LValue, value lua.LValue) {
		p := ConvertLuaPath(value).path
		result = append(result, p)
	})
	m.mapper.SetFlyList(result)
	return 0
}
func (m *LuaMapper) FlyList(L *lua.LState) int {
	fl := L.NewTable()
	result := m.mapper.FlyList()
	for _, v := range result {
		p := &LuaPath{path: v}
		fl.Append(p.Convert(L))
	}
	L.Push(fl)
	return 1
}
func (m *LuaMapper) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("reset", L.NewFunction(m.Reset))
	t.RawSetString("Reset", L.NewFunction(m.Reset))
	t.RawSetString("resettemporary", L.NewFunction(m.ResetTemporary))
	t.RawSetString("ResetTemporary", L.NewFunction(m.ResetTemporary))
	t.RawSetString("addtags", L.NewFunction(m.AddTags))
	t.RawSetString("Addtags", L.NewFunction(m.AddTags))
	t.RawSetString("settag", L.NewFunction(m.SetTag))
	t.RawSetString("SetTag", L.NewFunction(m.SetTag))
	t.RawSetString("settags", L.NewFunction(m.SetTags))
	t.RawSetString("SetTags", L.NewFunction(m.SetTags))
	t.RawSetString("tags", L.NewFunction(m.Tags))
	t.RawSetString("Tags", L.NewFunction(m.Tags))
	t.RawSetString("getpath", L.NewFunction(m.GetPath))
	t.RawSetString("GetPath", L.NewFunction(m.GetPath))
	t.RawSetString("addpath", L.NewFunction(m.AddPath))
	t.RawSetString("AddPath", L.NewFunction(m.AddPath))
	t.RawSetString("addtemporarypath", L.NewFunction(m.AddTemporaryPath))
	t.RawSetString("AddTemporaryPath", L.NewFunction(m.AddTemporaryPath))

	t.RawSetString("newpath", L.NewFunction(m.NewPath))
	t.RawSetString("NewPath", L.NewFunction(m.NewPath))
	t.RawSetString("getroomid", L.NewFunction(m.GetRoomID))
	t.RawSetString("GetRoomID", L.NewFunction(m.GetRoomID))
	t.RawSetString("getroomname", L.NewFunction(m.GetRoomName))
	t.RawSetString("GetRoomName", L.NewFunction(m.GetRoomName))
	t.RawSetString("removeroom", L.NewFunction(m.RemoveRoom))
	t.RawSetString("RemoveRoom", L.NewFunction(m.RemoveRoom))
	t.RawSetString("setroomname", L.NewFunction(m.SetRoomName))
	t.RawSetString("SetRoomName", L.NewFunction(m.SetRoomName))
	t.RawSetString("clearroom", L.NewFunction(m.ClearRoom))
	t.RawSetString("ClearRoom", L.NewFunction(m.ClearRoom))
	t.RawSetString("newarea", L.NewFunction(m.NewArea))
	t.RawSetString("NewArea", L.NewFunction(m.NewArea))
	t.RawSetString("getexits", L.NewFunction(m.GetExits))
	t.RawSetString("GetExits", L.NewFunction(m.GetExits))
	t.RawSetString("flashtags", L.NewFunction(m.FlashTags))
	t.RawSetString("FlashTags", L.NewFunction(m.FlashTags))
	t.RawSetString("flylist", L.NewFunction(m.FlyList))
	t.RawSetString("FlyList", L.NewFunction(m.FlyList))
	t.RawSetString("setflylist", L.NewFunction(m.SetFlyList))
	t.RawSetString("SetFlyList", L.NewFunction(m.SetFlyList))
	t.RawSetString("walkall", L.NewFunction(m.WalkAll))
	t.RawSetString("WalkAll", L.NewFunction(m.WalkAll))
	return t
}
func NewMapperModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("mapper",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			m := &LuaMapper{b.GetMapper()}
			l.SetGlobal("Mapper", m.Convert(l))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
