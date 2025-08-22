package v8engine

import (
	"context"
	"modules/mapper"
	"modules/world/bus"

	"github.com/jarlyyn/v8js"

	"github.com/herb-go/herbplugin"
	"github.com/jarlyyn/v8js/v8plugin"
)

type JsWalkAllResult struct {
	result *mapper.WalkAllResult
}

func (result *JsWalkAllResult) Convert(r *v8js.Context) *v8js.Consumed {
	t := r.NewObject()
	steps := make([]*v8js.Consumed, 0, len(result.result.Steps))

	for _, v := range result.result.Steps {
		s := &JsStep{step: v}
		steps = append(steps, s.Convert(r).Consume())
	}
	t.Set("steps", r.NewArray(steps...).Consume())
	t.Set("walked", r.NewStringArray(result.result.Walked...).Consume())
	t.Set("notwalked", r.NewStringArray(result.result.NotWalked...).Consume())
	return t.Consume()
}
func ConvertJsPath(r *v8js.Context, v *v8js.Consumed) *JsPath {

	t := v
	p := &JsPath{
		path: mapper.NewPath(),
	}
	cmd := t.Get("command")
	p.path.Command = cmd.String()
	cmd.Release()
	frm := t.Get("from")
	p.path.From = frm.String()
	frm.Release()
	to := t.Get("to")
	p.path.To = to.String()
	to.Release()
	delay := t.Get("delay")
	p.path.Delay = int(delay.Integer())
	delay.Release()

	tags := []string{}
	tagsvalue := t.Get("tags")
	if tagsvalue != nil {
		tags = tagsvalue.StringArrry()
	}
	tagsvalue.Release()
	for _, v := range tags {
		p.path.Tags[v] = true
	}
	etags := []string{}
	etagsvalue := t.Get("excludetags")
	if etagsvalue != nil {
		etags = etagsvalue.StringArrry()
	}
	etagsvalue.Release()
	for _, v := range etags {
		p.path.ExcludeTags[v] = true
	}

	return p
}

type JsPath struct {
	path *mapper.Path
}

func (p *JsPath) Convert(r *v8js.Context) *v8js.Consumed {
	t := r.NewObject()
	t.Set("command", r.NewString(p.path.Command).Consume())
	t.Set("from", r.NewString(p.path.From).Consume())
	t.Set("to", r.NewString(p.path.To).Consume())
	t.Set("delay", r.NewInt32(int32(p.path.Delay)).Consume())
	tags := []string{}
	for k, v := range p.path.Tags {
		if v {
			tags = append(tags, k)
		}
	}
	t.Set("tags", r.NewStringArray(tags...).Consume())
	etags := []string{}
	for k, v := range p.path.ExcludeTags {
		if v {
			etags = append(etags, k)
		}
	}
	t.Set("excludetags", r.NewStringArray(etags...).Consume())
	return t.Consume()
}

func ConvertJsStep(r *v8js.Context, v *v8js.JsValue) *JsStep {
	t := v
	s := &JsStep{
		step: &mapper.Step{},
	}
	cmd := t.Get("command")
	s.step.Command = cmd.String()
	cmd.Release()
	frm := t.Get("from")
	s.step.From = frm.String()
	frm.Release()
	to := t.Get("to")
	s.step.To = to.String()
	to.Release()
	delay := t.Get("delay")
	s.step.Delay = int(delay.Integer())
	delay.Release()
	return s
}

type JsStep struct {
	step *mapper.Step
}

func (s *JsStep) Convert(r *v8js.Context) *v8js.Consumed {
	t := r.NewObject()
	t.Set("command", r.NewString(s.step.Command).Consume())
	t.Set("from", r.NewString(s.step.From).Consume())
	t.Set("to", r.NewString(s.step.To).Consume())
	t.Set("delay", r.NewInt32(int32(s.step.Delay)).Consume())
	return t.Consume()
}

type JsMapper struct {
	mapper *mapper.Mapper
}

func (m *JsMapper) Reset(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	m.mapper.Reset()
	return nil
}
func (m *JsMapper) ResetTemporary(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	m.mapper.ResetTemporary()
	return nil
}

func (m *JsMapper) AddTags(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	tags := []string{}
	for _, v := range call.Args() {
		tags = append(tags, v.String())
	}
	m.mapper.AddTags(tags)
	return nil
}
func (m *JsMapper) SetTag(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	m.mapper.SetTag(call.GetArg(0).String(), call.GetArg(1).Boolean())
	return nil
}
func (m *JsMapper) FlashTags(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	m.mapper.FlashTags()
	return nil
}

func (m *JsMapper) SetTags(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	tagsv := call.GetArg(0)
	if tagsv.IsNull() {
		return nil
	}
	tags := call.GetArg(0).StringArrry()
	m.mapper.AddTags(tags)
	return nil
}

func (m *JsMapper) Tags(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewStringArray(m.mapper.Tags()...).Consume()
}
func (m *JsMapper) WalkAll(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	targets := call.GetArg(0).StringArrry()
	fly := int(call.GetArg(1).Integer())
	maxdistance := int(call.GetArg(2).Integer())
	v := call.GetArg(3)
	var option *v8js.Consumed
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
		blacklist.Release()
		whitelist := option.Get("whitelist")
		if whitelist != nil {
			opt.Whitelist = whitelist.StringArrry()
		}
		whitelist.Release()
		blockedpath := option.Get("blockedpath")
		if blockedpath != nil {
			blocked := blockedpath.Array()
			for _, v := range blocked {
				opt.BlockedPath = append(opt.BlockedPath, v.StringArrry())
				v.Release()
			}
		}
		blockedpath.Release()
	}
	result := m.mapper.WalkAll(targets, fly != 0, maxdistance, opt)
	jsresult := &JsWalkAllResult{result: result}
	return jsresult.Convert(call.Context())
}
func (m *JsMapper) GetPath(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

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
			jbl.Release()
			jwl := option.Get("whitelist")
			if jwl != nil {
				opt.Whitelist = jwl.StringArrry()
			}
			jwl.Release()
			blockedpath := option.Get("blockedpath")
			if blockedpath != nil {
				blocked := blockedpath.Array()
				for _, v := range blocked {
					opt.BlockedPath = append(opt.BlockedPath, v.StringArrry())
					v.Release()
				}
			}
			blockedpath.Release()
		}
	}
	steps := m.mapper.GetPath(from, fly != 0, to, opt)
	if steps == nil {
		return nil
	}
	t := []*v8js.Consumed{}
	for i := range steps {
		s := &JsStep{step: steps[i]}
		t = append(t, s.Convert(call.Context()).Consume())
	}
	return call.Context().NewArray(t...).Consume()
}
func (m *JsMapper) AddPath(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	if len(call.Args()) < 2 {
		return nil
	}

	id := call.GetArg(0).String()
	path := ConvertJsPath(call.Context(), call.GetArg(1))
	if path == nil {
		return nil
	}
	return call.Context().NewBoolean(m.mapper.AddPath(id, path.path)).Consume()
}
func (m *JsMapper) AddTemporaryPath(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	if len(call.Args()) < 2 {
		return nil
	}

	id := call.GetArg(0).String()
	path := ConvertJsPath(call.Context(), call.GetArg(1))
	if path == nil {
		return nil
	}
	return call.Context().NewBoolean(m.mapper.AddTemporaryPath(id, path.path)).Consume()
}

func (m *JsMapper) NewPath(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	p := &JsPath{
		path: mapper.NewPath(),
	}
	return p.Convert(call.Context())
}
func (m *JsMapper) GetRoomID(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	name := call.GetArg(0).String()
	ids := m.mapper.GetRoomID(name)
	return call.Context().NewStringArray(ids...).Consume()
}
func (m *JsMapper) GetRoomName(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	id := call.GetArg(0).String()
	return call.Context().NewString(m.mapper.GetRoomName(id)).Consume()
}
func (m *JsMapper) SetRoomName(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	id := call.GetArg(0).String()
	name := call.GetArg(1).String()
	m.mapper.SetRoomName(id, name)
	return nil
}
func (m *JsMapper) ClearRoom(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	id := call.GetArg(0).String()
	m.mapper.ClearRoom(id)
	return nil
}
func (m *JsMapper) RemoveRoom(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	id := call.GetArg(0).String()
	return call.Context().NewBoolean(m.mapper.RemoveRoom(id)).Consume()
}
func (m *JsMapper) NewArea(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	size := int(call.GetArg(0).Integer())
	ids := m.mapper.NewArea(size)
	return call.Context().NewStringArray(ids...).Consume()
}
func (m *JsMapper) GetExits(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	id := call.GetArg(0).String()
	exits := m.mapper.GetExits(id, call.GetArg(1).Boolean())
	t := []*v8js.Consumed{}
	for _, v := range exits {
		p := &JsPath{
			path: v,
		}
		t = append(t, p.Convert(call.Context()))
	}
	return call.Context().NewArray(t...).Consume()

}
func (m *JsMapper) SetFlyList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	fl := []*v8js.JsValue{}
	flv := call.GetArg(0)
	if flv == nil {
		return nil
	}

	fl = call.GetArg(0).Array()
	var result = []*mapper.Path{}
	for _, v := range fl {
		p := ConvertJsPath(call.Context(), v.Consume()).path
		v.Release()
		result = append(result, p)
	}
	m.mapper.SetFlyList(result)
	return nil
}
func (m *JsMapper) FlyList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	fl := []*v8js.Consumed{}
	result := m.mapper.FlyList()
	for _, v := range result {
		fl = append(fl, (&JsPath{path: v}).Convert(call.Context()).Consume())
	}
	return call.Context().NewArray(fl...).Consume()
}
func (m *JsMapper) Convert(r *v8js.Context) *v8js.JsValue {
	t := r.NewObject()
	t.Set("reset", r.NewFunction(m.Reset).Consume())
	t.Set("Reset", r.NewFunction(m.Reset).Consume())
	t.Set("resettemporary", r.NewFunction(m.ResetTemporary).Consume())
	t.Set("ResetTemporary", r.NewFunction(m.ResetTemporary).Consume())
	t.Set("addtags", r.NewFunction(m.AddTags).Consume())
	t.Set("AddTags", r.NewFunction(m.AddTags).Consume())
	t.Set("settag", r.NewFunction(m.SetTag).Consume())
	t.Set("SetTag", r.NewFunction(m.SetTag).Consume())
	t.Set("settags", r.NewFunction(m.SetTags).Consume())
	t.Set("SetTags", r.NewFunction(m.SetTags).Consume())
	t.Set("tags", r.NewFunction(m.Tags).Consume())
	t.Set("Tags", r.NewFunction(m.Tags).Consume())
	t.Set("getpath", r.NewFunction(m.GetPath).Consume())
	t.Set("GetPath", r.NewFunction(m.GetPath).Consume())
	t.Set("addpath", r.NewFunction(m.AddPath).Consume())
	t.Set("AddPath", r.NewFunction(m.AddPath).Consume())
	t.Set("addtemporarypath", r.NewFunction(m.AddTemporaryPath).Consume())
	t.Set("AddTemporaryPath", r.NewFunction(m.AddTemporaryPath).Consume())

	t.Set("newpath", r.NewFunction(m.NewPath).Consume())
	t.Set("NewPath", r.NewFunction(m.NewPath).Consume())
	t.Set("getroomid", r.NewFunction(m.GetRoomID).Consume())
	t.Set("GetRoomID", r.NewFunction(m.GetRoomID).Consume())
	t.Set("removeroom", r.NewFunction(m.RemoveRoom).Consume())
	t.Set("RemoveRoom", r.NewFunction(m.RemoveRoom).Consume())

	t.Set("getroomname", r.NewFunction(m.GetRoomName).Consume())
	t.Set("GetRoomName", r.NewFunction(m.GetRoomName).Consume())
	t.Set("setroomname", r.NewFunction(m.SetRoomName).Consume())
	t.Set("SetRoomName", r.NewFunction(m.SetRoomName).Consume())
	t.Set("clearroom", r.NewFunction(m.ClearRoom).Consume())
	t.Set("ClearRoom", r.NewFunction(m.ClearRoom).Consume())
	t.Set("newarea", r.NewFunction(m.NewArea).Consume())
	t.Set("NewArea", r.NewFunction(m.NewArea).Consume())
	t.Set("getexits", r.NewFunction(m.GetExits).Consume())
	t.Set("GetExits", r.NewFunction(m.GetExits).Consume())
	t.Set("flashtags", r.NewFunction(m.FlashTags).Consume())
	t.Set("FlashTags", r.NewFunction(m.FlashTags).Consume())
	t.Set("flylist", r.NewFunction(m.FlyList).Consume())
	t.Set("FlyList", r.NewFunction(m.FlyList).Consume())
	t.Set("setflylist", r.NewFunction(m.SetFlyList).Consume())
	t.Set("SetFlyList", r.NewFunction(m.SetFlyList).Consume())
	t.Set("WalkAll", r.NewFunction(m.WalkAll).Consume())
	t.Set("walkall", r.NewFunction(m.WalkAll).Consume())
	return t
}
func NewMapperModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("mapper",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			m := &JsMapper{b.GetMapper()}
			global := r.Global()
			global.Set("Mapper", m.Convert(r).Consume())
			global.Release()
			next(ctx, plugin)
		},
		nil,
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			b.GetMapper().Reset()
			next(ctx, plugin)
		})

}
