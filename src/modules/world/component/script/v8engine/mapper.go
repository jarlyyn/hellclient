package v8engine

import (
	"context"
	"modules/mapper"
	"modules/world/bus"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/v8local"
	"github.com/herb-go/v8local/v8plugin"
)

type JsWalkAllResult struct {
	result *mapper.WalkAllResult
}

func (result *JsWalkAllResult) Convert(r *v8local.Local) *v8local.JsValue {
	t := r.NewObject()
	steps := make([]*v8local.JsValue, 0, len(result.result.Steps))

	for _, v := range result.result.Steps {
		s := &JsStep{step: v}
		steps = append(steps, s.Convert(r))
	}
	t.Set("steps", r.NewArray(steps...))
	t.Set("walked", r.NewStringArray(result.result.Walked...))
	t.Set("notwalked", r.NewStringArray(result.result.NotWalked...))
	return t
}
func ConvertJsPath(r *v8local.Local, v *v8local.JsValue) *JsPath {

	t := v
	p := &JsPath{
		path: mapper.NewPath(),
	}
	cmd := t.Get("command")
	p.path.Command = cmd.String()
	frm := t.Get("from")
	p.path.From = frm.String()
	to := t.Get("to")
	p.path.To = to.String()
	delay := t.Get("delay")
	p.path.Delay = int(delay.Integer())

	tags := []string{}
	tagsvalue := t.Get("tags")
	if tagsvalue != nil {
		tags = tagsvalue.StringArrry()
	}
	for _, v := range tags {
		p.path.Tags[v] = true
	}
	etags := []string{}
	etagsvalue := t.Get("excludetags")
	if etagsvalue != nil {
		etags = etagsvalue.StringArrry()
	}
	for _, v := range etags {
		p.path.ExcludeTags[v] = true
	}

	return p
}

type JsPath struct {
	path *mapper.Path
}

func (p *JsPath) Convert(r *v8local.Local) *v8local.JsValue {
	t := r.NewObject()
	t.Set("command", r.NewString(p.path.Command))
	t.Set("from", r.NewString(p.path.From))
	t.Set("to", r.NewString(p.path.To))
	t.Set("delay", r.NewInt32(int32(p.path.Delay)))
	tags := []string{}
	for k, v := range p.path.Tags {
		if v {
			tags = append(tags, k)
		}
	}
	t.Set("tags", r.NewStringArray(tags...))
	etags := []string{}
	for k, v := range p.path.ExcludeTags {
		if v {
			etags = append(etags, k)
		}
	}
	t.Set("excludetags", r.NewStringArray(etags...))
	return t
}

func ConvertJsStep(r *v8local.Context, v *v8local.JsValue) *JsStep {
	t := v
	s := &JsStep{
		step: &mapper.Step{},
	}
	cmd := t.Get("command")
	s.step.Command = cmd.String()
	frm := t.Get("from")
	s.step.From = frm.String()
	to := t.Get("to")
	s.step.To = to.String()
	delay := t.Get("delay")
	s.step.Delay = int(delay.Integer())
	return s
}

type JsStep struct {
	step *mapper.Step
}

func (s *JsStep) Convert(r *v8local.Local) *v8local.JsValue {
	t := r.NewObject()
	t.Set("command", r.NewString(s.step.Command))
	t.Set("from", r.NewString(s.step.From))
	t.Set("to", r.NewString(s.step.To))
	t.Set("delay", r.NewInt32(int32(s.step.Delay)))
	return t
}

type JsMapper struct {
	mapper *mapper.Mapper
}

func (m *JsMapper) Reset(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	m.mapper.Reset()
	return nil
}
func (m *JsMapper) ResetTemporary(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	m.mapper.ResetTemporary()
	return nil
}

func (m *JsMapper) AddTags(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	tags := []string{}
	for _, v := range call.Args() {
		tags = append(tags, v.String())
	}
	m.mapper.AddTags(tags)
	return nil
}
func (m *JsMapper) SetTag(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	m.mapper.SetTag(call.GetArg(0).String(), call.GetArg(1).Boolean())
	return nil
}
func (m *JsMapper) FlashTags(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	m.mapper.FlashTags()
	return nil
}

func (m *JsMapper) SetTags(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	tagsv := call.GetArg(0)
	if tagsv.IsNull() {
		return nil
	}
	tags := call.GetArg(0).StringArrry()
	m.mapper.AddTags(tags)
	return nil
}

func (m *JsMapper) Tags(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	return call.Local().NewStringArray(m.mapper.Tags()...)
}
func (m *JsMapper) WalkAll(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	targets := call.GetArg(0).StringArrry()
	fly := int(call.GetArg(1).Integer())
	maxdistance := int(call.GetArg(2).Integer())
	v := call.GetArg(3)
	var option *v8local.JsValue
	if !v.IsNull() {
		option = call.GetArg(3)
	}
	var opt *mapper.Option
	if option != nil {
		opt = mapper.NewOption()
		blacklist := option.Get("blacklist")
		if blacklist != nil {
			opt.Blacklist = blacklist.StringArrry()
		}
		whitelist := option.Get("whitelist")
		if whitelist != nil {
			opt.Whitelist = whitelist.StringArrry()
		}
		blockedpath := option.Get("blockedpath")
		if blockedpath != nil {
			blocked := blockedpath.Array()
			for _, v := range blocked {
				opt.BlockedPath = append(opt.BlockedPath, v.StringArrry())
			}
		}
	}
	result := m.mapper.WalkAll(targets, fly != 0, maxdistance, opt)
	jsresult := &JsWalkAllResult{result: result}
	return jsresult.Convert(call.Local())
}
func (m *JsMapper) GetPath(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	if len(call.Args()) < 3 {
		return nil
	}

	from := call.GetArg(0).String()
	fly := int(call.GetArg(1).Integer())
	to := call.GetArg(2).StringArrry()
	var opt *mapper.Option
	jso := call.GetArg(3)
	if !jso.IsNull() {
		option := jso

		if option != nil {
			opt = mapper.NewOption()
			jbl := option.Get("blacklist")
			if jbl != nil {
				opt.Blacklist = jbl.StringArrry()
			}
			jwl := option.Get("whitelist")
			if jwl != nil {
				opt.Whitelist = jwl.StringArrry()
			}
			blockedpath := option.Get("blockedpath")
			if blockedpath != nil {
				blocked := blockedpath.Array()
				for _, v := range blocked {
					opt.BlockedPath = append(opt.BlockedPath, v.StringArrry())
				}
			}
		}
	}
	steps := m.mapper.GetPath(from, fly != 0, to, opt)
	if steps == nil {
		return nil
	}
	t := []*v8local.JsValue{}
	for i := range steps {
		s := &JsStep{step: steps[i]}
		t = append(t, s.Convert(call.Local()))
	}
	return call.Local().NewArray(t...)
}
func (m *JsMapper) AddPath(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	if len(call.Args()) < 2 {
		return nil
	}

	id := call.GetArg(0).String()
	path := ConvertJsPath(call.Local(), call.GetArg(1))
	if path == nil {
		return nil
	}
	return call.Local().NewBoolean(m.mapper.AddPath(id, path.path))
}
func (m *JsMapper) AddTemporaryPath(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	if len(call.Args()) < 2 {
		return nil
	}

	id := call.GetArg(0).String()
	path := ConvertJsPath(call.Local(), call.GetArg(1))
	if path == nil {
		return nil
	}
	return call.Local().NewBoolean(m.mapper.AddTemporaryPath(id, path.path))
}

func (m *JsMapper) NewPath(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	p := &JsPath{
		path: mapper.NewPath(),
	}
	return p.Convert(call.Local())
}
func (m *JsMapper) GetRoomID(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	name := call.GetArg(0).String()
	ids := m.mapper.GetRoomID(name)
	return call.Local().NewStringArray(ids...)
}
func (m *JsMapper) GetRoomName(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	id := call.GetArg(0).String()
	return call.Local().NewString(m.mapper.GetRoomName(id))
}
func (m *JsMapper) SetRoomName(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	id := call.GetArg(0).String()
	name := call.GetArg(1).String()
	m.mapper.SetRoomName(id, name)
	return nil
}
func (m *JsMapper) ClearRoom(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	id := call.GetArg(0).String()
	m.mapper.ClearRoom(id)
	return nil
}
func (m *JsMapper) RemoveRoom(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	id := call.GetArg(0).String()
	return call.Local().NewBoolean(m.mapper.RemoveRoom(id))
}
func (m *JsMapper) NewArea(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	size := int(call.GetArg(0).Integer())
	ids := m.mapper.NewArea(size)
	return call.Local().NewStringArray(ids...)
}
func (m *JsMapper) GetExits(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	id := call.GetArg(0).String()
	exits := m.mapper.GetExits(id, call.GetArg(1).Boolean())
	t := []*v8local.JsValue{}
	for _, v := range exits {
		p := &JsPath{
			path: v,
		}
		t = append(t, p.Convert(call.Local()))
	}
	return call.Local().NewArray(t...)

}
func (m *JsMapper) SetFlyList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	fl := []*v8local.JsValue{}
	flv := call.GetArg(0)
	if flv == nil {
		return nil
	}

	fl = call.GetArg(0).Array()
	var result = []*mapper.Path{}
	for _, v := range fl {
		p := ConvertJsPath(call.Local(), v).path
		result = append(result, p)
	}
	m.mapper.SetFlyList(result)
	return nil
}
func (m *JsMapper) FlyList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	fl := []*v8local.JsValue{}
	result := m.mapper.FlyList()
	for _, v := range result {
		fl = append(fl, (&JsPath{path: v}).Convert(call.Local()))
	}
	return call.Local().NewArray(fl...)
}
func (m *JsMapper) Convert(r *v8local.Local) *v8local.JsValue {
	t := r.NewObject()
	t.Set("reset", r.NewFunction(m.Reset))
	t.Set("Reset", r.NewFunction(m.Reset))
	t.Set("resettemporary", r.NewFunction(m.ResetTemporary))
	t.Set("ResetTemporary", r.NewFunction(m.ResetTemporary))
	t.Set("addtags", r.NewFunction(m.AddTags))
	t.Set("AddTags", r.NewFunction(m.AddTags))
	t.Set("settag", r.NewFunction(m.SetTag))
	t.Set("SetTag", r.NewFunction(m.SetTag))
	t.Set("settags", r.NewFunction(m.SetTags))
	t.Set("SetTags", r.NewFunction(m.SetTags))
	t.Set("tags", r.NewFunction(m.Tags))
	t.Set("Tags", r.NewFunction(m.Tags))
	t.Set("getpath", r.NewFunction(m.GetPath))
	t.Set("GetPath", r.NewFunction(m.GetPath))
	t.Set("addpath", r.NewFunction(m.AddPath))
	t.Set("AddPath", r.NewFunction(m.AddPath))
	t.Set("addtemporarypath", r.NewFunction(m.AddTemporaryPath))
	t.Set("AddTemporaryPath", r.NewFunction(m.AddTemporaryPath))

	t.Set("newpath", r.NewFunction(m.NewPath))
	t.Set("NewPath", r.NewFunction(m.NewPath))
	t.Set("getroomid", r.NewFunction(m.GetRoomID))
	t.Set("GetRoomID", r.NewFunction(m.GetRoomID))
	t.Set("removeroom", r.NewFunction(m.RemoveRoom))
	t.Set("RemoveRoom", r.NewFunction(m.RemoveRoom))

	t.Set("getroomname", r.NewFunction(m.GetRoomName))
	t.Set("GetRoomName", r.NewFunction(m.GetRoomName))
	t.Set("setroomname", r.NewFunction(m.SetRoomName))
	t.Set("SetRoomName", r.NewFunction(m.SetRoomName))
	t.Set("clearroom", r.NewFunction(m.ClearRoom))
	t.Set("ClearRoom", r.NewFunction(m.ClearRoom))
	t.Set("newarea", r.NewFunction(m.NewArea))
	t.Set("NewArea", r.NewFunction(m.NewArea))
	t.Set("getexits", r.NewFunction(m.GetExits))
	t.Set("GetExits", r.NewFunction(m.GetExits))
	t.Set("flashtags", r.NewFunction(m.FlashTags))
	t.Set("FlashTags", r.NewFunction(m.FlashTags))
	t.Set("flylist", r.NewFunction(m.FlyList))
	t.Set("FlyList", r.NewFunction(m.FlyList))
	t.Set("setflylist", r.NewFunction(m.SetFlyList))
	t.Set("SetFlyList", r.NewFunction(m.SetFlyList))
	t.Set("WalkAll", r.NewFunction(m.WalkAll))
	t.Set("walkall", r.NewFunction(m.WalkAll))
	return t
}
func NewMapperModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("mapper",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Top
			m := &JsMapper{b.GetMapper()}
			global := r.Global()
			global.Set("Mapper", m.Convert(r))
			next(ctx, plugin)
		},
		nil,
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			b.GetMapper().Reset()
			next(ctx, plugin)
		})

}
