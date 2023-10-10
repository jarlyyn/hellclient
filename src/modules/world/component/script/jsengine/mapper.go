package jsengine

import (
	"context"
	"errors"
	"modules/mapper"
	"modules/world/bus"

	"github.com/dop251/goja"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
)

type JsWalkAllResult struct {
	result *mapper.WalkAllResult
}

func (result *JsWalkAllResult) Convert(r *goja.Runtime) goja.Value {
	t := r.NewObject()
	steps := make([]interface{}, 0, len(result.result.Steps))
	for _, v := range result.result.Steps {
		s := &JsStep{step: v}
		steps = append(steps, s.Convert(r))
	}
	t.Set("steps", r.NewArray(steps...))
	t.Set("walked", result.result.Walked)
	t.Set("notwalked", result.result.NotWalked)
	return t
}
func ConvertJsPath(r *goja.Runtime, v goja.Value) *JsPath {
	t := v.ToObject(r)
	p := &JsPath{
		path: mapper.NewPath(),
	}
	p.path.Command = t.Get("command").String()
	p.path.From = t.Get("from").String()
	p.path.To = t.Get("to").String()
	p.path.Delay = int(t.Get("delay").ToInteger())
	tags := []string{}
	tagsvalue := t.Get("tags")
	if tagsvalue != nil {
		err := r.ExportTo(tagsvalue, &tags)
		if err != nil {
			panic(errors.New("tags must be array"))
		}
	}
	for _, v := range tags {
		p.path.Tags[v] = true
	}
	etags := []string{}
	etagsvalue := t.Get("excludetags")
	if etagsvalue != nil {
		err := r.ExportTo(etagsvalue, &etags)
		if err != nil {
			panic(errors.New("excludetags must be array"))
		}
	}
	for _, v := range etags {
		p.path.ExcludeTags[v] = true
	}
	return p
}

type JsPath struct {
	path *mapper.Path
}

func (p *JsPath) Convert(r *goja.Runtime) goja.Value {
	t := r.NewObject()
	t.Set("command", p.path.Command)
	t.Set("from", p.path.From)
	t.Set("to", p.path.To)
	t.Set("delay", p.path.Delay)
	tags := []string{}
	for k, v := range p.path.Tags {
		if v {
			tags = append(tags, k)
		}
	}
	t.Set("tags", tags)
	etags := []string{}
	for k, v := range p.path.ExcludeTags {
		if v {
			etags = append(etags, k)
		}
	}
	t.Set("excludetags", etags)
	return t
}

func ConvertJsStep(r *goja.Runtime, v goja.Value) *JsStep {
	t := v.ToObject(r)
	s := &JsStep{
		step: &mapper.Step{},
	}
	s.step.Command = t.Get("command").String()
	s.step.From = t.Get("from").String()
	s.step.To = t.Get("to").String()
	s.step.Delay = int(t.Get("delay").ToInteger())
	return s
}

type JsStep struct {
	step *mapper.Step
}

func (s *JsStep) Convert(r *goja.Runtime) goja.Value {
	t := r.NewObject()
	t.Set("command", s.step.Command)
	t.Set("from", s.step.From)
	t.Set("to", s.step.To)
	t.Set("delay", s.step.Delay)
	return t
}

type JsMapper struct {
	mapper *mapper.Mapper
}

func (m *JsMapper) Reset(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	m.mapper.Reset()
	return nil
}
func (m *JsMapper) ResetTemporary(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	m.mapper.ResetTemporary()
	return nil
}

func (m *JsMapper) AddTags(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	tags := []string{}
	for _, v := range call.Arguments {
		tags = append(tags, v.String())
	}
	m.mapper.AddTags(tags)
	return nil
}
func (m *JsMapper) SetTag(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	m.mapper.SetTag(call.Argument(0).String(), call.Argument(1).ToBoolean())
	return nil
}
func (m *JsMapper) FlashTags(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	m.mapper.FlashTags()
	return nil
}

func (m *JsMapper) SetTags(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	tags := []string{}
	tagsv := call.Argument(0)
	if tagsv == nil {
		return nil
	}
	err := r.ExportTo(call.Argument(0), &tags)
	if err != nil {
		panic(errors.New("tags must be array"))
	}
	m.mapper.AddTags(tags)
	return nil
}

func (m *JsMapper) Tags(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(m.mapper.Tags())
}
func (m *JsMapper) WalkAll(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	targets := []string{}
	err := r.ExportTo(call.Argument(0), &targets)
	if err != nil {
		return nil
	}
	fly := int(call.Argument(1).ToInteger())
	maxdistance := int(call.Argument(2).ToInteger())
	result := m.mapper.WalkAll(targets, fly != 0, maxdistance)
	jsresult := &JsWalkAllResult{result: result}
	return jsresult.Convert(r)
}
func (m *JsMapper) GetPath(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if len(call.Arguments) < 3 {
		return nil
	}
	from := call.Argument(0).String()
	fly := int(call.Argument(1).ToInteger())
	to := []string{}
	err := r.ExportTo(call.Argument(2), &to)
	if err != nil {
		return nil
	}
	steps := m.mapper.GetPath(from, fly != 0, to)
	if steps == nil {
		return nil
	}
	t := []goja.Value{}
	for i := range steps {
		s := &JsStep{step: steps[i]}
		t = append(t, s.Convert(r))
	}
	return r.ToValue(t)
}
func (m *JsMapper) AddPath(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if len(call.Arguments) < 2 {
		return nil
	}

	id := call.Argument(0).String()
	path := ConvertJsPath(r, call.Argument(1))
	if path == nil {
		return nil
	}
	return r.ToValue(m.mapper.AddPath(id, path.path))
}
func (m *JsMapper) AddTemporaryPath(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if len(call.Arguments) < 2 {
		return nil
	}

	id := call.Argument(0).String()
	path := ConvertJsPath(r, call.Argument(1))
	if path == nil {
		return nil
	}
	return r.ToValue(m.mapper.AddTemporaryPath(id, path.path))
}

func (m *JsMapper) NewPath(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	p := &JsPath{
		path: mapper.NewPath(),
	}
	return r.ToValue(p.Convert(r))
}
func (m *JsMapper) GetRoomID(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	name := call.Argument(0).String()
	ids := m.mapper.GetRoomID(name)
	return r.ToValue(ids)
}
func (m *JsMapper) GetRoomName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	id := call.Argument(0).String()
	return r.ToValue(m.mapper.GetRoomName(id))
}
func (m *JsMapper) SetRoomName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	id := call.Argument(0).String()
	name := call.Argument(1).String()
	m.mapper.SetRoomName(id, name)
	return nil
}
func (m *JsMapper) ClearRoom(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	id := call.Argument(0).String()
	m.mapper.ClearRoom(id)
	return nil
}
func (m *JsMapper) NewArea(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	size := int(call.Argument(0).ToInteger())
	ids := m.mapper.NewArea(size)
	return r.ToValue(ids)
}
func (m *JsMapper) GetExits(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	id := call.Argument(0).String()
	exits := m.mapper.GetExits(id, call.Argument(1).ToBoolean())
	t := []goja.Value{}
	for _, v := range exits {
		p := &JsPath{
			path: v,
		}
		t = append(t, p.Convert(r))
	}
	return r.ToValue(t)

}
func (m *JsMapper) SetFlyList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	fl := []*goja.Object{}
	flv := call.Argument(0)
	if flv == nil {
		return nil
	}
	err := r.ExportTo(call.Argument(0), &fl)
	if err != nil {
		panic(errors.New("flylist must be array"))
	}
	var result = []*mapper.Path{}
	for _, v := range fl {
		p := ConvertJsPath(r, v).path
		result = append(result, p)
	}
	m.mapper.SetFlyList(result)
	return nil
}
func (m *JsMapper) FlyList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	fl := []goja.Value{}
	result := m.mapper.FlyList()
	for _, v := range result {
		fl = append(fl, (&JsPath{path: v}).Convert(r))
	}
	return r.ToValue(fl)
}
func (m *JsMapper) Convert(r *goja.Runtime) goja.Value {
	t := r.NewObject()
	t.Set("reset", m.Reset)
	t.Set("Reset", m.Reset)
	t.Set("resettemporary", m.ResetTemporary)
	t.Set("ResetTemporary", m.ResetTemporary)
	t.Set("addtags", m.AddTags)
	t.Set("AddTags", m.AddTags)
	t.Set("settag", m.SetTag)
	t.Set("SetTag", m.SetTag)
	t.Set("settags", m.SetTags)
	t.Set("SetTags", m.SetTags)
	t.Set("tags", m.Tags)
	t.Set("Tags", m.Tags)
	t.Set("getpath", m.GetPath)
	t.Set("GetPath", m.GetPath)
	t.Set("addpath", m.AddPath)
	t.Set("AddPath", m.AddPath)
	t.Set("addtemporarypath", m.AddTemporaryPath)
	t.Set("AddTemporaryPath", m.AddTemporaryPath)

	t.Set("newpath", m.NewPath)
	t.Set("NewPath", m.NewPath)
	t.Set("getroomid", m.GetRoomID)
	t.Set("GetRoomID", m.GetRoomID)
	t.Set("getroomname", m.GetRoomName)
	t.Set("GetRoomName", m.GetRoomName)
	t.Set("setroomname", m.SetRoomName)
	t.Set("SetRoomName", m.SetRoomName)
	t.Set("clearroom", m.ClearRoom)
	t.Set("ClearRoom", m.ClearRoom)
	t.Set("newarea", m.NewArea)
	t.Set("NewArea", m.NewArea)
	t.Set("getexits", m.GetExits)
	t.Set("GetExits", m.GetExits)
	t.Set("flashtags", m.FlashTags)
	t.Set("FlashTags", m.FlashTags)
	t.Set("flylist", m.FlyList)
	t.Set("FlyList", m.FlyList)
	t.Set("setflylist", m.SetFlyList)
	t.Set("SetFlyList", m.SetFlyList)
	t.Set("WalkAll", m.WalkAll)
	t.Set("walkall", m.WalkAll)

	return t
}
func NewMapperModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("mapper",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			m := &JsMapper{b.GetMapper()}
			r.Set("Mapper", m.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
