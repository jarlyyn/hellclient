package v8engine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/userinput"
	"net/url"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/v8local"
	"github.com/herb-go/v8local/v8plugin"
)

// type VisualPrompt struct {
// 	VisualPrompt *userinput.VisualPrompt
// 	bus          *bus.Bus
// }

func VisualPromptSetMediaType(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("MediaType", call.GetArg(0))
	return nil
}
func VisualPromptSetPortrait(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("Portrait", call.GetArg(0))
	return nil
}
func VisualPromptSetValue(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("Value", call.GetArg(0))
	return nil
}
func VisualPromptSetRefreshCallback(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("RefreshCallback", call.GetArg(0))
	return nil
}
func VisualPromptAppend(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	items := call.This().Get("Items")
	fn := items.Get("push")
	fn.Call(items, call.Local().NewArray(call.GetArg(0), call.GetArg(1)))
	return nil
}
func VisualPromptPublish(b *bus.Bus) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		var title = call.This().Get("Title")
		var intro = call.This().Get("Intro")
		var source = call.This().Get("Source")
		vp := userinput.CreateVisualPrompt(
			title.String(),
			intro.String(),
			source.String(),
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
				v0 := v.GetIdx(0)
				v1 := v.GetIdx(1)
				vp.Append(v0.String(), v1.String())
			}
		}
		v := call.This().Get("Value")
		vp.Value = v.String()
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
		return call.Local().NewString(ui.ID)
	}
}
func VisualPromptConvert(r *v8local.Local, ui *Userinput, title string, intro string, source string) *v8local.JsValue {
	vp := userinput.CreateVisualPrompt(title, intro, source)
	obj := r.NewObject()
	obj.Set("Title", r.NewString(vp.Title))
	obj.Set("Intro", r.NewString(vp.Intro))
	obj.Set("Source", r.NewString(vp.Source))
	obj.Set("MediaType", r.NewString(vp.MediaType))
	obj.Set("Portrait", r.NewBoolean(vp.Portrait))
	obj.Set("Value", r.NewString(vp.Value))
	obj.Set("RefreshCallback", r.NewString(vp.RefreshCallback))
	obj.Set("Items", r.NewArray())

	obj.Set("setmediatype", ui.Functions["VisualPromptSetMediaType"])
	obj.Set("SetMediaType", ui.Functions["VisualPromptSetMediaType"])
	obj.Set("setportrait", ui.Functions["VisualPromptSetPortrait"])
	obj.Set("SetPortrait", ui.Functions["VisualPromptSetPortrait"])
	obj.Set("setvalue", ui.Functions["VisualPromptSetValue"])
	obj.Set("SetValue", ui.Functions["VisualPromptSetValue"])
	obj.Set("append", ui.Functions["VisualPromptAppend"])
	obj.Set("Append", ui.Functions["VisualPromptAppend"])
	obj.Set("setrefreshcallback", ui.Functions["VisualPromptSetRefreshCallback"])
	obj.Set("SetRefreshCallback", ui.Functions["VisualPromptSetRefreshCallback"])
	obj.Set("publish", ui.Functions["VisualPromptPublish"])
	obj.Set("Publish", ui.Functions["VisualPromptPublish"])
	return obj
}

type Datagrid struct {
	Datagrid *userinput.Datagrid
	bus      *bus.Bus
}

func DatagridSetPage(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("Page", call.GetArg(0))
	return nil
}
func DatagridGetPage(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	page := call.This().Get("Page")
	return page
}
func DatagridSetMaxPage(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("MaxPage", call.GetArg(0))
	return nil
}
func DatagridSetFilter(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("Filter", call.GetArg(0))
	return nil
}
func DatagridGetFilter(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	filter := call.This().Get("Filter")
	return filter
}
func DatagridSetOnPage(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("OnPage", call.GetArg(0))
	return nil
}
func DatagridSetOnFilter(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("OnFilter", call.GetArg(0))
	return nil
}
func DatagridSetOnDelete(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("OnDelete", call.GetArg(0))
	return nil
}
func DatagridSetOnView(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("OnView", call.GetArg(0))
	return nil
}
func DatagridSetOnSelect(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("OnSelect", call.GetArg(0))
	return nil
}
func DatagridSetOnCreate(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("OnCreate", call.GetArg(0))
	return nil
}
func DatagridSetOnUpdate(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("OnUpdate", call.GetArg(0))
	return nil
}
func DatagridResetItems(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("Items", call.Local().NewArray())
	return nil
}
func DatagridAppend(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	items := call.This().Get("Items")
	fn := items.Get("push")
	fn.Call(items, call.Local().NewArray(call.GetArg(0), call.GetArg(1)))
	return nil
}
func DatagridPublish(b *bus.Bus) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		var title = call.This().Get("Title")
		var intro = call.This().Get("Intro")
		g := userinput.CreateDatagrid(
			title.String(),
			intro.String(),
		)
		items := call.This().Get("Items")
		for _, v := range items.Array() {
			v0 := v.GetIdx(0)
			v1 := v.GetIdx(1)
			g.Append(v0.String(), v1.String())
		}
		page := call.This().Get("Page")
		g.Page = int(page.Number())
		maxPage := call.This().Get("MaxPage")
		g.MaxPage = int(maxPage.Number())
		onCreate := call.This().Get("OnCreate")
		g.OnCreate = onCreate.String()
		onView := call.This().Get("OnView")
		g.OnView = onView.String()
		onSelect := call.This().Get("OnSelect")
		g.OnSelect = onSelect.String()
		onUpdate := call.This().Get("OnUpdate")
		g.OnUpdate = onUpdate.String()
		onPage := call.This().Get("OnPage")
		g.OnPage = onPage.String()
		onFilter := call.This().Get("OnFilter")
		g.OnFilter = onFilter.String()
		onDelete := call.This().Get("OnDelete")
		g.OnDelete = onDelete.String()
		ui := g.Publish(b, call.GetArg(0).String())
		return call.Local().NewString(ui.ID)
	}
}
func DatagridHide(b *bus.Bus) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		(&userinput.Datagrid{}).Hide(b)
		return nil
	}
}
func DatagridConvert(r *v8local.Local, ui *Userinput, title string, intro string) *v8local.JsValue {
	dg := userinput.CreateDatagrid(title, intro)
	obj := r.NewObject()
	obj.Set("Title", r.NewString(dg.Title))
	obj.Set("Intro", r.NewString(dg.Intro))
	obj.Set("Items", r.NewArray())
	obj.Set("Page", r.NewNumber(float64(dg.Page)))
	obj.Set("MaxPage", r.NewNumber(float64(dg.MaxPage)))
	obj.Set("OnCreate", r.NewString(dg.OnCreate))
	obj.Set("OnView", r.NewString(dg.OnView))
	obj.Set("OnSelect", r.NewString(dg.OnSelect))
	obj.Set("OnUpdate", r.NewString(dg.OnUpdate))
	obj.Set("OnPage", r.NewString(dg.OnPage))
	obj.Set("OnFilter", r.NewString(dg.OnFilter))
	obj.Set("OnDelete", r.NewString(dg.OnDelete))

	obj.Set("append", ui.Functions["DatagridAppend"])
	obj.Set("Append", ui.Functions["DatagridAppend"])
	obj.Set("publish", ui.Functions["DatagridPublish"])
	obj.Set("Publish", ui.Functions["DatagridPublish"])
	obj.Set("resetitems", ui.Functions["DatagridResetItems"])
	obj.Set("ResetItems", ui.Functions["DatagridResetItems"])
	obj.Set("setoncreate", ui.Functions["DatagridSetOnCreate"])
	obj.Set("SetOnCreate", ui.Functions["DatagridSetOnCreate"])
	obj.Set("setonupdate", ui.Functions["DatagridSetOnUpdate"])
	obj.Set("SetOnUpdate", ui.Functions["DatagridSetOnUpdate"])
	obj.Set("setonview", ui.Functions["DatagridSetOnView"])
	obj.Set("SetOnView", ui.Functions["DatagridSetOnView"])
	obj.Set("setonselect", ui.Functions["DatagridSetOnSelect"])
	obj.Set("SetOnSelect", ui.Functions["DatagridSetOnSelect"])
	obj.Set("setondelete", ui.Functions["DatagridSetOnDelete"])
	obj.Set("SetOnDelete", ui.Functions["DatagridSetOnDelete"])
	obj.Set("setonfilter", ui.Functions["DatagridSetOnFilter"])
	obj.Set("SetOnFilter", ui.Functions["DatagridSetOnFilter"])
	obj.Set("setonpage", ui.Functions["DatagridSetOnPage"])
	obj.Set("SetOnPage", ui.Functions["DatagridSetOnPage"])
	obj.Set("setfilter", ui.Functions["DatagridSetFilter"])
	obj.Set("SetFilter", ui.Functions["DatagridSetFilter"])
	obj.Set("getfilter", ui.Functions["DatagridGetFilter"])
	obj.Set("GetFilter", ui.Functions["DatagridGetFilter"])
	obj.Set("setmaxpage", ui.Functions["DatagridSetMaxPage"])
	obj.Set("SetMaxPage", ui.Functions["DatagridSetMaxPage"])
	obj.Set("setpage", ui.Functions["DatagridSetPage"])
	obj.Set("SetPage", ui.Functions["DatagridSetPage"])
	obj.Set("getpage", ui.Functions["DatagridGetPage"])
	obj.Set("GetPage", ui.Functions["DatagridGetPage"])
	obj.Set("hide", ui.Functions["DatagridHide"])
	obj.Set("Hide", ui.Functions["DatagridHide"])
	obj.Set("Publish", ui.Functions["DatagridPublish"])
	return obj

}

type List struct {
	List *userinput.List
	bus  *bus.Bus
}

func ListPublish(b *bus.Bus) func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return func(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
		title := call.This().Get("Title")
		intro := call.This().Get("Intro")
		withFilter := call.This().Get("WithFilter")
		l := userinput.CreateList(
			title.String(),
			intro.String(),
			withFilter.Boolean(),
		)
		items := call.This().Get("Items")
		for _, v := range items.Array() {
			v0 := v.GetIdx(0)
			v1 := v.GetIdx(1)
			l.Append(v0.String(), v1.String())
		}
		mutli := call.This().Get("Multi")
		l.SetMulti(mutli.Boolean())
		ui := l.Publish(b, call.GetArg(0).String())
		return call.Local().NewString(ui.ID)
	}
}
func ListAppend(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	items := call.This().Get("Items")
	fn := items.Get("push")
	fn.Call(items, call.Local().NewArray(call.GetArg(0), call.GetArg(1)))
	return nil
}
func ListSetValues(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("Items", call.GetArg(0))
	return nil
}
func ListSetMulti(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	call.This().Set("Multi", call.GetArg(0))
	return nil
}
func ListConvert(r *v8local.Local, ui *Userinput, title string, intro string, withFilter bool) *v8local.JsValue {
	l := userinput.CreateList(title, intro, withFilter)
	obj := r.NewObject()
	obj.Set("Title", r.NewString(l.Title))
	obj.Set("Intro", r.NewString(l.Intro))
	obj.Set("Items", r.NewArray())
	obj.Set("Multi", r.NewBoolean(l.Mutli))
	obj.Set("WithFilter", r.NewBoolean(l.WithFilter))

	obj.Set("append", ui.Functions["ListAppend"])
	obj.Set("Append", ui.Functions["ListAppend"])
	obj.Set("publish", ui.Functions["ListPublish"])
	obj.Set("Publish", ui.Functions["ListPublish"])
	obj.Set("setmulti", ui.Functions["ListSetMulti"])
	obj.Set("SetMulti", ui.Functions["ListSetMulti"])
	obj.Set("setmutli", ui.Functions["ListSetMulti"]) //老api的typo,保持兼容性
	obj.Set("SetMutli", ui.Functions["ListSetMulti"]) //老api的typo,保持兼容性
	obj.Set("setvalues", ui.Functions["ListSetValues"])
	obj.Set("SetValues", ui.Functions["ListSetValues"])
	obj.Set("Publish", ui.Functions["ListPublish"])
	return obj
}

type Userinput struct {
	bus       *bus.Bus
	Functions map[string]*v8local.JsValue
}

func (u *Userinput) HideAll(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	userinput.HideAll(u.bus)
	return nil
}
func (u *Userinput) NewVisualPrompt(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return VisualPromptConvert(call.Local(), u, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String())
}
func (u *Userinput) NewDatagrid(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return DatagridConvert(call.Local(), u, call.GetArg(0).String(), call.GetArg(1).String())
}
func (u *Userinput) NewList(call *v8local.FunctionCallbackInfo) *v8local.JsValue {
	return ListConvert(call.Local(), u, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).Boolean())
}
func (u *Userinput) Prompt(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	ui := userinput.SendPrompt(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String(), call.GetArg(3).String())
	return call.Local().NewString(ui.ID)
}
func (u *Userinput) Confirm(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	ui := userinput.SendConfirm(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String())
	return call.Local().NewString(ui.ID)
}
func (u *Userinput) Alert(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	ui := userinput.SendAlert(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String())
	return call.Local().NewString(ui.ID)
}
func (u *Userinput) Popup(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	ui := userinput.SendPopup(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String(), call.GetArg(3).String())
	return call.Local().NewString(ui.ID)
}
func (u *Userinput) Note(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	ui := userinput.SendNote(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(2).String(), call.GetArg(3).String())
	return call.Local().NewString(ui.ID)
}
func (u *Userinput) Custom(call *v8local.FunctionCallbackInfo) *v8local.JsValue {

	ui := userinput.SendCustom(u.bus, call.GetArg(0).String(), call.GetArg(1).String(), call.GetArg(1).String())
	return call.Local().NewString(ui.ID)
}
func (u *Userinput) Convert(r *v8local.Local) *v8local.JsValue {
	obj := r.NewObject()
	obj.Set("prompt", r.NewFunction(u.Prompt))
	obj.Set("Prompt", r.NewFunction(u.Prompt))
	obj.Set("confirm", r.NewFunction(u.Confirm))
	obj.Set("Confirm", r.NewFunction(u.Confirm))
	obj.Set("alert", r.NewFunction(u.Alert))
	obj.Set("Alert", r.NewFunction(u.Alert))
	obj.Set("popup", r.NewFunction(u.Popup))
	obj.Set("Popup", r.NewFunction(u.Popup))
	obj.Set("note", r.NewFunction(u.Note))
	obj.Set("Note", r.NewFunction(u.Note))
	obj.Set("newlist", r.NewFunction(u.NewList))
	obj.Set("Newlist", r.NewFunction(u.NewList))
	obj.Set("newdatagrid", r.NewFunction(u.NewDatagrid))
	obj.Set("NewDatagrid", r.NewFunction(u.NewDatagrid))
	obj.Set("newvisualprompt", r.NewFunction(u.NewVisualPrompt))
	obj.Set("NewVisualPrompt", r.NewFunction(u.NewVisualPrompt))
	obj.Set("hideall", r.NewFunction(u.HideAll))
	obj.Set("HideAll", r.NewFunction(u.HideAll))
	obj.Set("Custom", r.NewFunction(u.Custom))
	obj.Set("custom", r.NewFunction(u.Custom))
	return obj
}
func (u *Userinput) Init(r *v8local.Local) {
	u.Functions = make(map[string]*v8local.JsValue)

	u.Functions["VisualPromptSetMediaType"] = r.NewFunction(VisualPromptSetMediaType).AsExported()
	u.Functions["VisualPromptSetPortrait"] = r.NewFunction(VisualPromptSetPortrait).AsExported()
	u.Functions["VisualPromptSetValue"] = r.NewFunction(VisualPromptSetValue).AsExported()
	u.Functions["VisualPromptSetRefreshCallback"] = r.NewFunction(VisualPromptSetRefreshCallback).AsExported()
	u.Functions["VisualPromptAppend"] = r.NewFunction(VisualPromptAppend).AsExported()
	u.Functions["VisualPromptPublish"] = r.NewFunction(VisualPromptPublish(u.bus)).AsExported()

	u.Functions["DatagridSetPage"] = r.NewFunction(DatagridSetPage).AsExported()
	u.Functions["DatagridGetPage"] = r.NewFunction(DatagridGetPage).AsExported()
	u.Functions["DatagridSetMaxPage"] = r.NewFunction(DatagridSetMaxPage).AsExported()
	u.Functions["DatagridSetFilter"] = r.NewFunction(DatagridSetFilter).AsExported()
	u.Functions["DatagridGetFilter"] = r.NewFunction(DatagridGetFilter).AsExported()
	u.Functions["DatagridSetOnPage"] = r.NewFunction(DatagridSetOnPage).AsExported()
	u.Functions["DatagridSetOnFilter"] = r.NewFunction(DatagridSetOnFilter).AsExported()
	u.Functions["DatagridSetOnDelete"] = r.NewFunction(DatagridSetOnDelete).AsExported()
	u.Functions["DatagridSetOnView"] = r.NewFunction(DatagridSetOnView).AsExported()
	u.Functions["DatagridSetOnSelect"] = r.NewFunction(DatagridSetOnSelect).AsExported()
	u.Functions["DatagridSetOnCreate"] = r.NewFunction(DatagridSetOnCreate).AsExported()
	u.Functions["DatagridSetOnUpdate"] = r.NewFunction(DatagridSetOnUpdate).AsExported()
	u.Functions["DatagridResetItems"] = r.NewFunction(DatagridResetItems).AsExported()
	u.Functions["DatagridAppend"] = r.NewFunction(DatagridAppend).AsExported()
	u.Functions["DatagridPublish"] = r.NewFunction(DatagridPublish(u.bus)).AsExported()
	u.Functions["DatagridHide"] = r.NewFunction(DatagridHide(u.bus)).AsExported()

	u.Functions["ListPublish"] = r.NewFunction(ListPublish(u.bus)).AsExported()
	u.Functions["ListAppend"] = r.NewFunction(ListAppend).AsExported()
	u.Functions["ListSetValues"] = r.NewFunction(ListSetValues).AsExported()
	u.Functions["ListSetMulti"] = r.NewFunction(ListSetMulti).AsExported()
}
func NewUserinputModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("userinput",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*v8plugin.Plugin).LoadJsPlugin()
			r := jsp.Top
			u := &Userinput{bus: b}
			global := r.Global()
			u.Init(r)
			global.Set("Userinput", u.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
