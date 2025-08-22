package v8engine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/userinput"
	"net/url"

	"github.com/herb-go/herbplugin"
	"github.com/jarlyyn/v8js"
	"github.com/jarlyyn/v8js/v8plugin"
)

type VisualPrompt struct {
	VisualPrompt *userinput.VisualPrompt
	bus          *bus.Bus
}

func (p *VisualPrompt) SetMediaType(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	p.VisualPrompt.SetMediaType(call.GetArg(0).String())
	return nil
}
func (p *VisualPrompt) SetPortrait(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	p.VisualPrompt.SetPortrait(call.GetArg(0).Boolean())
	return nil
}
func (p *VisualPrompt) SetValue(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	p.VisualPrompt.SetValue(call.GetArg(0).String())
	return nil
}
func (p *VisualPrompt) SetRefreshCallback(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	p.VisualPrompt.SetRefreshCallback(call.GetArg(0).String())
	return nil
}
func (p *VisualPrompt) Append(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	p.VisualPrompt.Append(call.GetArg(0).String(), call.GetArg(1).String())
	return nil
}
func (p *VisualPrompt) Publish(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	if p.VisualPrompt.IsURL() {
		u, err := url.Parse(p.VisualPrompt.Source)
		if err != nil {
			panic(err)
		}
		if !p.bus.GetPluginOptions().MustAuthorizeDomain(u.Host) {
			panic(herbplugin.NewUnauthorizeDomainError(u.Host))
		}
	}
	ui := p.VisualPrompt.Publish(p.bus, call.GetArg(0).String())
	return call.Context().NewString(ui.ID).Consume()

}
func (p *VisualPrompt) Convert(r *v8js.Context) *v8js.JsValue {
	obj := r.NewObject()
	obj.Set("setmediatype", r.NewFunction(p.SetMediaType).Consume())
	obj.Set("SetMediaType", r.NewFunction(p.SetMediaType).Consume())
	obj.Set("setportrait", r.NewFunction(p.SetPortrait).Consume())
	obj.Set("SetPortrait", r.NewFunction(p.SetPortrait).Consume())
	obj.Set("setvalue", r.NewFunction(p.SetValue).Consume())
	obj.Set("SetValue", r.NewFunction(p.SetValue).Consume())
	obj.Set("append", r.NewFunction(p.Append).Consume())
	obj.Set("Append", r.NewFunction(p.Append).Consume())
	obj.Set("setrefreshcallback", r.NewFunction(p.SetRefreshCallback).Consume())
	obj.Set("SetRefreshCallback", r.NewFunction(p.SetRefreshCallback).Consume())
	obj.Set("publish", r.NewFunction(p.Publish).Consume())
	obj.Set("Publish", r.NewFunction(p.Publish).Consume())
	return obj

}

type Datagrid struct {
	Datagrid *userinput.Datagrid
	bus      *bus.Bus
}

func (g *Datagrid) SetPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetPage(int(call.GetArg(0).Integer()))
	return nil
}
func (g *Datagrid) GetPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewInt32(int32(g.Datagrid.GetPage())).Consume()
}
func (g *Datagrid) SetMaxPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetMaxPage(int(call.GetArg(0).Integer()))
	return nil
}
func (g *Datagrid) SetFilter(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetFilter(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) GetFilter(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	return call.Context().NewString(g.Datagrid.GetFilter()).Consume()
}
func (g *Datagrid) SetOnPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetOnPage(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) SetOnFilter(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetOnFilter(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) SetOnDelete(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetOnDelete(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) SetOnView(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetOnView(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) SetOnSelect(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetOnSelect(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) SetOnCreate(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetOnCreate(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) SetOnUpdate(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.SetOnUpdate(call.GetArg(0).String())
	return nil
}
func (g *Datagrid) ResetItems(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.ResetItems()
	return nil
}
func (g *Datagrid) Append(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.Append(call.GetArg(0).String(), call.GetArg(1).String())
	return nil
}
func (g *Datagrid) Publish(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := g.Datagrid.Publish(g.bus, call.GetArg(0).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (g *Datagrid) Hide(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	g.Datagrid.Hide(g.bus)
	return nil
}
func (g *Datagrid) Convert(r *v8js.Context) *v8js.JsValue {
	obj := r.NewObject()
	obj.Set("append", r.NewFunction(g.Append).Consume())
	obj.Set("Append", r.NewFunction(g.Append).Consume())
	obj.Set("publish", r.NewFunction(g.Publish).Consume())
	obj.Set("Publish", r.NewFunction(g.Publish).Consume())
	obj.Set("resetitems", r.NewFunction(g.ResetItems).Consume())
	obj.Set("ResetItems", r.NewFunction(g.ResetItems).Consume())
	obj.Set("setoncreate", r.NewFunction(g.SetOnCreate).Consume())
	obj.Set("SetOnCreate", r.NewFunction(g.SetOnCreate).Consume())
	obj.Set("setonupdate", r.NewFunction(g.SetOnUpdate).Consume())
	obj.Set("SetOnUpdate", r.NewFunction(g.SetOnUpdate).Consume())
	obj.Set("setonview", r.NewFunction(g.SetOnView).Consume())
	obj.Set("SetOnView", r.NewFunction(g.SetOnView).Consume())
	obj.Set("setonselect", r.NewFunction(g.SetOnSelect).Consume())
	obj.Set("SetOnSelect", r.NewFunction(g.SetOnSelect).Consume())
	obj.Set("setondelete", r.NewFunction(g.SetOnDelete).Consume())
	obj.Set("SetOnDelete", r.NewFunction(g.SetOnDelete).Consume())
	obj.Set("setonfilter", r.NewFunction(g.SetOnFilter).Consume())
	obj.Set("SetOnFilter", r.NewFunction(g.SetOnFilter).Consume())
	obj.Set("setonpage", r.NewFunction(g.SetOnPage).Consume())
	obj.Set("SetOnPage", r.NewFunction(g.SetOnPage).Consume())
	obj.Set("setfilter", r.NewFunction(g.SetFilter).Consume())
	obj.Set("SetFilter", r.NewFunction(g.SetFilter).Consume())
	obj.Set("getfilter", r.NewFunction(g.GetFilter).Consume())
	obj.Set("GetFilter", r.NewFunction(g.GetFilter).Consume())
	obj.Set("setmaxpage", r.NewFunction(g.SetMaxPage).Consume())
	obj.Set("SetMaxPage", r.NewFunction(g.SetMaxPage).Consume())
	obj.Set("setpage", r.NewFunction(g.SetPage).Consume())
	obj.Set("SetPage", r.NewFunction(g.SetPage).Consume())
	obj.Set("getpage", r.NewFunction(g.GetPage).Consume())
	obj.Set("GetPage", r.NewFunction(g.GetPage).Consume())
	obj.Set("hide", r.NewFunction(g.Hide).Consume())
	obj.Set("Hide", r.NewFunction(g.Hide).Consume())
	return obj

}

type List struct {
	List *userinput.List
	bus  *bus.Bus
}

func (l *List) Publish(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := l.List.Publish(l.bus, call.GetArg(0).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (l *List) Append(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	l.List.Append(call.GetArg(0).String(), call.GetArg(1).String())
	return nil
}
func (l *List) SetValues(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	v := call.GetArg(0).StringArrry()
	l.List.SetValues(v)
	return nil
}
func (l *List) SetMulti(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	l.List.SetMulti(call.GetArg(0).Boolean())
	return nil
}
func (l *List) Convert(r *v8js.Context) *v8js.JsValue {
	obj := r.NewObject()
	obj.Set("append", r.NewFunction(l.Append).Consume())
	obj.Set("Append", r.NewFunction(l.Append).Consume())
	obj.Set("publish", r.NewFunction(l.Publish).Consume())
	obj.Set("Publish", r.NewFunction(l.Publish).Consume())
	obj.Set("setmulti", r.NewFunction(l.SetMulti).Consume())
	obj.Set("SetMulti", r.NewFunction(l.SetMulti).Consume())
	obj.Set("setmutli", r.NewFunction(l.SetMulti).Consume()) //老api的typo,保持兼容性
	obj.Set("SetMutli", r.NewFunction(l.SetMulti).Consume()) //老api的typo,保持兼容性
	obj.Set("setvalues", r.NewFunction(l.SetValues).Consume())
	obj.Set("SetValues", r.NewFunction(l.SetValues).Consume())
	return obj
}

type Userinput struct {
	bus *bus.Bus
}

func (u *Userinput) HideAll(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	userinput.HideAll(u.bus)
	return nil
}
func (u *Userinput) NewVisualPrompt(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	vp := &VisualPrompt{
		VisualPrompt: userinput.CreateVisualPrompt(call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String()),
		bus:          u.bus,
	}
	return vp.Convert(call.Context()).Consume()
}
func (u *Userinput) NewDatagrid(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	datagrid := &Datagrid{
		Datagrid: userinput.CreateDatagrid(call.GetArg(0).String(), call.GetArg(1).String()),
		bus:      u.bus,
	}
	return datagrid.Convert(call.Context()).Consume()
}
func (u *Userinput) NewList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	list := &List{
		List: userinput.CreateList(call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).Boolean()),
		bus:  u.bus,
	}
	return list.Convert(call.Context()).Consume()
}
func (u *Userinput) Prompt(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := userinput.SendPrompt(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String(), call.GetArg(3).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (u *Userinput) Confirm(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := userinput.SendConfirm(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (u *Userinput) Alert(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := userinput.SendAlert(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (u *Userinput) Popup(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := userinput.SendPopup(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String(), call.GetArg(3).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (u *Userinput) Note(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := userinput.SendNote(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String(), call.GetArg(3).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (u *Userinput) Custom(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	ui := userinput.SendCustom(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(1).String())
	return call.Context().NewString(ui.ID).Consume()
}
func (u *Userinput) Convert(r *v8js.Context) *v8js.JsValue {
	obj := r.NewObject()
	obj.Set("prompt", r.NewFunction(u.Prompt).Consume())
	obj.Set("Prompt", r.NewFunction(u.Prompt).Consume())
	obj.Set("confirm", r.NewFunction(u.Confirm).Consume())
	obj.Set("Confirm", r.NewFunction(u.Confirm).Consume())
	obj.Set("alert", r.NewFunction(u.Alert).Consume())
	obj.Set("Alert", r.NewFunction(u.Alert).Consume())
	obj.Set("popup", r.NewFunction(u.Popup).Consume())
	obj.Set("Popup", r.NewFunction(u.Popup).Consume())
	obj.Set("note", r.NewFunction(u.Note).Consume())
	obj.Set("Note", r.NewFunction(u.Note).Consume())
	obj.Set("newlist", r.NewFunction(u.NewList).Consume())
	obj.Set("Newlist", r.NewFunction(u.NewList).Consume())
	obj.Set("newdatagrid", r.NewFunction(u.NewDatagrid).Consume())
	obj.Set("NewDatagrid", r.NewFunction(u.NewDatagrid).Consume())
	obj.Set("newvisualprompt", r.NewFunction(u.NewVisualPrompt).Consume())
	obj.Set("NewVisualPrompt", r.NewFunction(u.NewVisualPrompt).Consume())
	obj.Set("hideall", r.NewFunction(u.HideAll).Consume())
	obj.Set("HideAll", r.NewFunction(u.HideAll).Consume())
	obj.Set("Custom", r.NewFunction(u.Custom).Consume())
	obj.Set("custom", r.NewFunction(u.Custom).Consume())
	return obj
}
func NewUserinputModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("userinput",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			u := &Userinput{bus: b}
			global := r.Global()
			global.Set("Userinput", u.Convert(r).Consume())
			global.Release()
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
