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

// type VisualPrompt struct {
// 	VisualPrompt *userinput.VisualPrompt
// 	bus          *bus.Bus
// }

func VisualPromptSetMediaType(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("MediaType", call.GetArg(0))
	return nil
}
func VisualPromptSetPortrait(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("Portrait", call.GetArg(0))
	return nil
}
func VisualPromptSetValue(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("Value", call.GetArg(0))
	return nil
}
func VisualPromptSetRefreshCallback(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("RefreshCallback", call.GetArg(0))
	return nil
}
func VisualPromptAppend(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	items := call.This().Get("Items")
	items.Get("push").Call(items, call.Context().NewArray(call.GetArg(0), call.GetArg(1)).Consume())
	return nil
}
func VisualPromptPublish(b *bus.Bus) func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
		vp := userinput.CreateVisualPrompt(
			call.This().Get("Title").String(),
			call.This().Get("Intro").String(),
			call.This().Get("Source").String(),
		)
		mt := call.This().Get("MediaType")
		if !mt.IsNullOrUndefined() {
			vp.MediaType = mt.String()
		}
		p := call.This().Get("Portrait")
		if !p.IsNullOrUndefined() {
			vp.Portrait = p.Boolean()
		}
		cb := call.This().Get("RefreshCallback")
		if !cb.IsNullOrUndefined() {
			vp.RefreshCallback = cb.String()
		}
		items := call.This().Get("Items")
		if !items.IsNullOrUndefined() {
			for _, v := range items.Array() {
				vp.Append(v.GetIdx(0).String(), v.GetIdx(1).String())
			}
		}
		vp.Value = call.This().Get("Value").String()
		if vp.IsURL() {
			u, err := url.Parse(vp.Source)
			if err != nil {
				panic(err)
			}
			if !b.GetPluginOptions().MustAuthorizeDomain(u.Host) {
				panic(herbplugin.NewUnauthorizeDomainError(u.Host))
			}
		}
		ui := vp.Publish(b, call.GetArg(0).String())
		return call.Context().NewString(ui.ID).Consume()
	}
}
func VisualPromptConvert(r *v8js.Context, ui *Userinput, title string, intro string, source string) *v8js.JsValue {
	vp := userinput.CreateVisualPrompt(title, intro, source)
	obj := r.NewObject()
	obj.Set("Title", r.NewString(vp.Title).Consume())
	obj.Set("Intro", r.NewString(vp.Intro).Consume())
	obj.Set("Source", r.NewString(vp.Source).Consume())
	obj.Set("MediaType", r.NewString(vp.MediaType).Consume())
	obj.Set("Portrait", r.NewBoolean(vp.Portrait).Consume())
	obj.Set("Value", r.NewString(vp.Value).Consume())
	obj.Set("RefreshCallback", r.NewString(vp.RefreshCallback).Consume())
	obj.Set("Items", r.NewArray().Consume())

	obj.Set("setmediatype", ui.Functions["VisualPromptSetMediaType"].Consume())
	obj.Set("SetMediaType", ui.Functions["VisualPromptSetMediaType"].Consume())
	obj.Set("setportrait", ui.Functions["VisualPromptSetPortrait"].Consume())
	obj.Set("SetPortrait", ui.Functions["VisualPromptSetPortrait"].Consume())
	obj.Set("setvalue", ui.Functions["VisualPromptSetValue"].Consume())
	obj.Set("SetValue", ui.Functions["VisualPromptSetValue"].Consume())
	obj.Set("append", ui.Functions["VisualPromptAppend"].Consume())
	obj.Set("Append", ui.Functions["VisualPromptAppend"].Consume())
	obj.Set("setrefreshcallback", ui.Functions["VisualPromptSetRefreshCallback"].Consume())
	obj.Set("SetRefreshCallback", ui.Functions["VisualPromptSetRefreshCallback"].Consume())
	obj.Set("publish", ui.Functions["VisualPromptPublish"].Consume())
	obj.Set("Publish", ui.Functions["VisualPromptPublish"].Consume())
	return obj
}

type Datagrid struct {
	Datagrid *userinput.Datagrid
	bus      *bus.Bus
}

func DatagridSetPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("Page", call.GetArg(0))
	return nil
}
func DatagridGetPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return call.This().Get("Page").Consume()
}
func DatagridSetMaxPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("MaxPage", call.GetArg(0))
	return nil
}
func DatagridSetFilter(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("Filter", call.GetArg(0))
	return nil
}
func DatagridGetFilter(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return call.This().Get("Filter").Consume()
}
func DatagridSetOnPage(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("OnPage", call.GetArg(0))
	return nil
}
func DatagridSetOnFilter(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("OnFilter", call.GetArg(0))
	return nil
}
func DatagridSetOnDelete(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("OnDelete", call.GetArg(0))
	return nil
}
func DatagridSetOnView(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("OnView", call.GetArg(0))
	return nil
}
func DatagridSetOnSelect(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("OnSelect", call.GetArg(0))
	return nil
}
func DatagridSetOnCreate(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("OnCreate", call.GetArg(0))
	return nil
}
func DatagridSetOnUpdate(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("OnUpdate", call.GetArg(0))
	return nil
}
func DatagridResetItems(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("Items", call.Context().NewArray().Consume())
	return nil
}
func DatagridAppend(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	items := call.This().Get("Items")
	items.Get("push").Call(items, call.Context().NewArray(call.GetArg(0), call.GetArg(1)).Consume())
	return nil
}
func DatagridPublish(b *bus.Bus) func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
		g := userinput.CreateDatagrid(
			call.This().Get("Title").String(),
			call.This().Get("Intro").String(),
		)
		items := call.This().Get("Items")
		for _, v := range items.Array() {
			g.Append(v.GetIdx(0).String(), v.GetIdx(1).String())
		}
		g.Page = int(call.This().Get("Page").Number())
		g.MaxPage = int(call.This().Get("MaxPage").Number())
		g.OnCreate = call.This().Get("OnCreate").String()
		g.OnView = call.This().Get("OnView").String()
		g.OnSelect = call.This().Get("OnSelect").String()
		g.OnUpdate = call.This().Get("OnUpdate").String()
		g.OnPage = call.This().Get("OnPage").String()
		g.OnFilter = call.This().Get("OnFilter").String()
		g.OnDelete = call.This().Get("OnDelete").String()

		ui := g.Publish(b, call.GetArg(0).String())
		return call.Context().NewString(ui.ID).Consume()
	}
}
func DatagridHide(b *bus.Bus) func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
		(&userinput.Datagrid{}).Hide(b)
		return nil
	}
}
func DatagridConvert(r *v8js.Context, ui *Userinput, title string, intro string) *v8js.JsValue {
	dg := userinput.CreateDatagrid(title, intro)
	obj := r.NewObject()
	obj.Set("Title", r.NewString(dg.Title).Consume())
	obj.Set("Intro", r.NewString(dg.Intro).Consume())
	obj.Set("Items", r.NewArray().Consume())
	obj.Set("Page", r.NewNumber(float64(dg.Page)).Consume())
	obj.Set("MaxPage", r.NewNumber(float64(dg.MaxPage)).Consume())
	obj.Set("OnCreate", r.NewString(dg.OnCreate).Consume())
	obj.Set("OnView", r.NewString(dg.OnView).Consume())
	obj.Set("OnSelect", r.NewString(dg.OnSelect).Consume())
	obj.Set("OnUpdate", r.NewString(dg.OnUpdate).Consume())
	obj.Set("OnPage", r.NewString(dg.OnPage).Consume())
	obj.Set("OnFilter", r.NewString(dg.OnFilter).Consume())
	obj.Set("OnDelete", r.NewString(dg.OnDelete).Consume())

	obj.Set("append", ui.Functions["DatagridAppend"].Consume())
	obj.Set("Append", ui.Functions["DatagridAppend"].Consume())
	obj.Set("publish", ui.Functions["DatagridPublish"].Consume())
	obj.Set("Publish", ui.Functions["DatagridPublish"].Consume())
	obj.Set("resetitems", ui.Functions["DatagridResetItems"].Consume())
	obj.Set("ResetItems", ui.Functions["DatagridResetItems"].Consume())
	obj.Set("setoncreate", ui.Functions["DatagridSetOnCreate"].Consume())
	obj.Set("SetOnCreate", ui.Functions["DatagridSetOnCreate"].Consume())
	obj.Set("setonupdate", ui.Functions["DatagridSetOnUpdate"].Consume())
	obj.Set("SetOnUpdate", ui.Functions["DatagridSetOnUpdate"].Consume())
	obj.Set("setonview", ui.Functions["DatagridSetOnView"].Consume())
	obj.Set("SetOnView", ui.Functions["DatagridSetOnView"].Consume())
	obj.Set("setonselect", ui.Functions["DatagridSetOnSelect"].Consume())
	obj.Set("SetOnSelect", ui.Functions["DatagridSetOnSelect"].Consume())
	obj.Set("setondelete", ui.Functions["DatagridSetOnDelete"].Consume())
	obj.Set("SetOnDelete", ui.Functions["DatagridSetOnDelete"].Consume())
	obj.Set("setonfilter", ui.Functions["DatagridSetOnFilter"].Consume())
	obj.Set("SetOnFilter", ui.Functions["DatagridSetOnFilter"].Consume())
	obj.Set("setonpage", ui.Functions["DatagridSetOnPage"].Consume())
	obj.Set("SetOnPage", ui.Functions["DatagridSetOnPage"].Consume())
	obj.Set("setfilter", ui.Functions["DatagridSetFilter"].Consume())
	obj.Set("SetFilter", ui.Functions["DatagridSetFilter"].Consume())
	obj.Set("getfilter", ui.Functions["DatagridGetFilter"].Consume())
	obj.Set("GetFilter", ui.Functions["DatagridGetFilter"].Consume())
	obj.Set("setmaxpage", ui.Functions["DatagridSetMaxPage"].Consume())
	obj.Set("SetMaxPage", ui.Functions["DatagridSetMaxPage"].Consume())
	obj.Set("setpage", ui.Functions["DatagridSetPage"].Consume())
	obj.Set("SetPage", ui.Functions["DatagridSetPage"].Consume())
	obj.Set("getpage", ui.Functions["DatagridGetPage"].Consume())
	obj.Set("GetPage", ui.Functions["DatagridGetPage"].Consume())
	obj.Set("hide", ui.Functions["DatagridHide"].Consume())
	obj.Set("Hide", ui.Functions["DatagridHide"].Consume())
	obj.Set("Publish", ui.Functions["DatagridPublish"].Consume())
	return obj

}

type List struct {
	List *userinput.List
	bus  *bus.Bus
}

func ListPublish(b *bus.Bus) func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return func(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
		l := userinput.CreateList(
			call.This().Get("Title").String(),
			call.This().Get("Intro").String(),
			call.This().Get("WithFilter").Boolean(),
		)
		items := call.This().Get("Items")
		for _, v := range items.Array() {
			l.Append(v.GetIdx(0).String(), v.GetIdx(1).String())
		}
		l.SetMulti(call.This().Get("Multi").Boolean())
		ui := l.Publish(b, call.GetArg(0).String())
		return call.Context().NewString(ui.ID).Consume()
	}
}
func ListAppend(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	items := call.This().Get("Items")
	items.Get("push").Call(items, call.Context().NewArray(call.GetArg(0), call.GetArg(1)).Consume())
	return nil
}
func ListSetValues(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("Items", call.GetArg(0))
	return nil
}
func ListSetMulti(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	call.This().Set("Multi", call.GetArg(0))
	return nil
}
func ListConvert(r *v8js.Context, ui *Userinput, title string, intro string, withFilter bool) *v8js.JsValue {
	l := userinput.CreateList(title, intro, withFilter)
	obj := r.NewObject()
	obj.Set("Title", r.NewString(l.Title).Consume())
	obj.Set("Intro", r.NewString(l.Intro).Consume())
	obj.Set("Items", r.NewArray().Consume())
	obj.Set("Multi", r.NewBoolean(l.Mutli).Consume())
	obj.Set("WithFilter", r.NewBoolean(l.WithFilter).Consume())

	obj.Set("append", ui.Functions["ListAppend"].Consume())
	obj.Set("Append", ui.Functions["ListAppend"].Consume())
	obj.Set("publish", ui.Functions["ListPublish"].Consume())
	obj.Set("Publish", ui.Functions["ListPublish"].Consume())
	obj.Set("setmulti", ui.Functions["ListSetMulti"].Consume())
	obj.Set("SetMulti", ui.Functions["ListSetMulti"].Consume())
	obj.Set("setmutli", ui.Functions["ListSetMulti"].Consume()) //老api的typo,保持兼容性
	obj.Set("SetMutli", ui.Functions["ListSetMulti"].Consume()) //老api的typo,保持兼容性
	obj.Set("setvalues", ui.Functions["ListSetValues"].Consume())
	obj.Set("SetValues", ui.Functions["ListSetValues"].Consume())
	obj.Set("Publish", ui.Functions["ListPublish"].Consume())
	return obj
}

type Userinput struct {
	bus       *bus.Bus
	Functions map[string]*v8js.Reusable
}

func (u *Userinput) HideAll(call *v8js.FunctionCallbackInfo) *v8js.Consumed {

	userinput.HideAll(u.bus)
	return nil
}
func (u *Userinput) NewVisualPrompt(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return VisualPromptConvert(call.Context(), u, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String()).Consume()
}
func (u *Userinput) NewDatagrid(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return DatagridConvert(call.Context(), u, call.GetArg(0).String(), call.GetArg(1).String()).Consume()
}
func (u *Userinput) NewList(call *v8js.FunctionCallbackInfo) *v8js.Consumed {
	return ListConvert(call.Context(), u, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).Boolean()).Consume()
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
func (u *Userinput) Init(r *v8js.Context) {
	u.Functions = make(map[string]*v8js.Reusable)

	u.Functions["VisualPromptSetMediaType"] = r.NewFunction(VisualPromptSetMediaType).ConsumeReuseble()
	u.Functions["VisualPromptSetPortrait"] = r.NewFunction(VisualPromptSetPortrait).ConsumeReuseble()
	u.Functions["VisualPromptSetValue"] = r.NewFunction(VisualPromptSetValue).ConsumeReuseble()
	u.Functions["VisualPromptSetRefreshCallback"] = r.NewFunction(VisualPromptSetRefreshCallback).ConsumeReuseble()
	u.Functions["VisualPromptAppend"] = r.NewFunction(VisualPromptAppend).ConsumeReuseble()
	u.Functions["VisualPromptPublish"] = r.NewFunction(VisualPromptPublish(u.bus)).ConsumeReuseble()

	u.Functions["DatagridSetPage"] = r.NewFunction(DatagridSetPage).ConsumeReuseble()
	u.Functions["DatagridGetPage"] = r.NewFunction(DatagridGetPage).ConsumeReuseble()
	u.Functions["DatagridSetMaxPage"] = r.NewFunction(DatagridSetMaxPage).ConsumeReuseble()
	u.Functions["DatagridSetFilter"] = r.NewFunction(DatagridSetFilter).ConsumeReuseble()
	u.Functions["DatagridGetFilter"] = r.NewFunction(DatagridGetFilter).ConsumeReuseble()
	u.Functions["DatagridSetOnPage"] = r.NewFunction(DatagridSetOnPage).ConsumeReuseble()
	u.Functions["DatagridSetOnFilter"] = r.NewFunction(DatagridSetOnFilter).ConsumeReuseble()
	u.Functions["DatagridSetOnDelete"] = r.NewFunction(DatagridSetOnDelete).ConsumeReuseble()
	u.Functions["DatagridSetOnView"] = r.NewFunction(DatagridSetOnView).ConsumeReuseble()
	u.Functions["DatagridSetOnSelect"] = r.NewFunction(DatagridSetOnSelect).ConsumeReuseble()
	u.Functions["DatagridSetOnCreate"] = r.NewFunction(DatagridSetOnCreate).ConsumeReuseble()
	u.Functions["DatagridSetOnUpdate"] = r.NewFunction(DatagridSetOnUpdate).ConsumeReuseble()
	u.Functions["DatagridResetItems"] = r.NewFunction(DatagridResetItems).ConsumeReuseble()
	u.Functions["DatagridAppend"] = r.NewFunction(DatagridAppend).ConsumeReuseble()
	u.Functions["DatagridPublish"] = r.NewFunction(DatagridPublish(u.bus)).ConsumeReuseble()
	u.Functions["DatagridHide"] = r.NewFunction(DatagridHide(u.bus)).ConsumeReuseble()

	u.Functions["ListPublish"] = r.NewFunction(ListPublish(u.bus)).ConsumeReuseble()
	u.Functions["ListAppend"] = r.NewFunction(ListAppend).ConsumeReuseble()
	u.Functions["ListSetValues"] = r.NewFunction(ListSetValues).ConsumeReuseble()
	u.Functions["ListSetMulti"] = r.NewFunction(ListSetMulti).ConsumeReuseble()
}
func NewUserinputModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("userinput",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			u := &Userinput{bus: b}
			global := r.Global()
			u.Init(r)
			global.Set("Userinput", u.Convert(r).Consume())
			global.Release()
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
