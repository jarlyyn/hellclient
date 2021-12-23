package jsengine

import (
	"context"
	"hellclient/modules/world/bus"
	"hellclient/modules/world/component/script/userinput"
	"net/url"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
)

type VisualPrompt struct {
	VisualPrompt *userinput.VisualPrompt
	bus          *bus.Bus
}

func (p *VisualPrompt) SetMediaType(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	p.VisualPrompt.SetMediaType(call.Argument(0).String())
	return nil
}
func (p *VisualPrompt) SetPortrait(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	p.VisualPrompt.SetPortrait(call.Argument(0).ToBoolean())
	return nil

}
func (p *VisualPrompt) SetRefreshCallback(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	p.VisualPrompt.SetRefreshCallback(call.Argument(0).String())
	return nil
}
func (p *VisualPrompt) Publish(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	u, err := url.Parse(p.VisualPrompt.Source)
	if err != nil {
		panic(err)
	}
	if !p.bus.GetPluginOptions().MustAuthorizeDomain(u.Host) {
		panic(herbplugin.NewUnauthorizeDomainError(u.Host))
	}
	ui := p.VisualPrompt.Publish(p.bus, call.Argument(0).String())
	return r.ToValue(ui.ID)

}
func (p *VisualPrompt) Convert(r *goja.Runtime) goja.Value {
	obj := r.NewObject()
	obj.Set("setmediatype", p.SetMediaType)
	obj.Set("setportrait", p.SetPortrait)
	obj.Set("setrefreshcallback", p.SetRefreshCallback)
	obj.Set("publish", p.Publish)
	return obj
}

type Datagrid struct {
	Datagrid *userinput.Datagrid
	bus      *bus.Bus
}

func (g *Datagrid) SetPage(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetPage(int(call.Argument(0).ToInteger()))
	return nil
}
func (g *Datagrid) GetPage(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(g.Datagrid.GetPage())
}
func (g *Datagrid) SetMaxPage(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetMaxPage(int(call.Argument(0).ToInteger()))
	return nil
}
func (g *Datagrid) SetFilter(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetFilter(call.Argument(0).String())
	return nil
}
func (g *Datagrid) GetFilter(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(g.Datagrid.GetFilter())
}
func (g *Datagrid) SetOnPage(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetOnPage(call.Argument(0).String())
	return nil
}
func (g *Datagrid) SetOnFilter(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetOnFilter(call.Argument(0).String())
	return nil
}
func (g *Datagrid) SetOnDelete(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetOnDelete(call.Argument(0).String())
	return nil
}
func (g *Datagrid) SetOnView(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetOnView(call.Argument(0).String())
	return nil
}
func (g *Datagrid) SetOnCreate(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetOnCreate(call.Argument(0).String())
	return nil
}
func (g *Datagrid) SetOnUpdate(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.SetOnUpdate(call.Argument(0).String())
	return nil
}
func (g *Datagrid) ResetItems(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.ResetItems()
	return nil
}
func (g *Datagrid) Append(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.Append(call.Argument(0).String(), call.Argument(1).String())
	return nil
}
func (g *Datagrid) Publish(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := g.Datagrid.Publish(g.bus, call.Argument(0).String())
	return r.ToValue(ui.ID)
}
func (g *Datagrid) Hide(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	g.Datagrid.Hide(g.bus)
	return nil
}
func (g *Datagrid) Convert(r *goja.Runtime) goja.Value {
	obj := r.NewObject()
	obj.Set("append", g.Append)
	obj.Set("publish", g.Publish)
	obj.Set("resetitems", g.ResetItems)
	obj.Set("setoncreate", g.SetOnCreate)
	obj.Set("setonupdate", g.SetOnUpdate)
	obj.Set("setonview", g.SetOnView)
	obj.Set("setondelete", g.SetOnDelete)
	obj.Set("setonfilter", g.SetOnFilter)
	obj.Set("setonpage", g.SetOnPage)
	obj.Set("setfilter", g.SetFilter)
	obj.Set("getfilter", g.GetFilter)
	obj.Set("setmaxpage", g.SetMaxPage)
	obj.Set("setpage", g.SetPage)
	obj.Set("getpage", g.GetPage)
	obj.Set("hide", g.Hide)
	return obj
}

type List struct {
	List *userinput.List
	bus  *bus.Bus
}

func (l *List) Send(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := l.List.Send(l.bus, call.Argument(0).String())
	return r.ToValue(ui.ID)
}
func (l *List) Append(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	l.List.Append(call.Argument(0).String(), call.Argument(1).String())
	return nil
}
func (l *List) SetValues(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	v := []string{}
	err := r.ExportTo(call.Argument(0), &v)
	if err != nil {
		panic(err)
	}
	l.List.SetValues(v)
	return nil
}
func (l *List) SetMutli(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	l.List.SetMutli(call.Argument(0).ToBoolean())
	return nil
}
func (l *List) Convert(r *goja.Runtime) goja.Value {
	obj := r.NewObject()
	obj.Set("append", l.Append)
	obj.Set("send", l.Send)
	obj.Set("setmutli", l.SetMutli)
	obj.Set("setvalues", l.SetValues)
	return obj
}

type Userinput struct {
	bus *bus.Bus
}

func (u *Userinput) NewVisualPrompt(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	datagrid := &VisualPrompt{
		VisualPrompt: userinput.CreateVisualPrompt(call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String()),
		bus:          u.bus,
	}
	return datagrid.Convert(r)
}
func (u *Userinput) NewDatagrid(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	datagrid := &Datagrid{
		Datagrid: userinput.CreateDatagrid(call.Argument(0).String(), call.Argument(1).String()),
		bus:      u.bus,
	}
	return datagrid.Convert(r)
}
func (u *Userinput) NewList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	list := &List{
		List: userinput.CreateList(call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).ToBoolean()),
		bus:  u.bus,
	}
	return list.Convert(r)
}
func (u *Userinput) Prompt(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := userinput.SendPrompt(u.bus, call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String(), call.Argument(3).String())
	return r.ToValue(ui.ID)
}
func (u *Userinput) Confirm(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := userinput.SendConfirm(u.bus, call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String())
	return r.ToValue(ui.ID)
}
func (u *Userinput) Alert(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := userinput.SendAlert(u.bus, call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String())
	return r.ToValue(ui.ID)
}

func (u *Userinput) Convert(r *goja.Runtime) goja.Value {
	obj := r.NewObject()
	obj.Set("prompt", u.Prompt)
	obj.Set("confirm", u.Confirm)
	obj.Set("alert", u.Alert)
	obj.Set("newlist", u.NewList)
	obj.Set("newdatagrid", u.NewDatagrid)
	obj.Set("newvisualprompt", u.NewVisualPrompt)
	return obj
}
func NewUserinputModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("userinput",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			u := &Userinput{bus: b}
			r.Set("Userinput", u.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
