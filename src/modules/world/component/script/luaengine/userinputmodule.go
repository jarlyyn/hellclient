package luaengine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/userinput"
	"net/url"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/lua51plugin"
	lua "github.com/yuin/gopher-lua"
)

type VisualPrompt struct {
	VisualPrompt *userinput.VisualPrompt
	bus          *bus.Bus
}

func (p *VisualPrompt) SetMediaType(L *lua.LState) int {
	p.VisualPrompt.SetMediaType(L.ToString((1)))
	return 0
}
func (p *VisualPrompt) SetPortrait(L *lua.LState) int {
	p.VisualPrompt.SetPortrait(L.ToBool((1)))
	return 0
}
func (p *VisualPrompt) SetValue(L *lua.LState) int {
	p.VisualPrompt.SetValue(L.ToString((1)))
	return 0
}
func (p *VisualPrompt) SetRefreshCallback(L *lua.LState) int {
	p.VisualPrompt.SetRefreshCallback(L.ToString((1)))
	return 0
}
func (p *VisualPrompt) Append(L *lua.LState) int {
	p.VisualPrompt.Append(L.ToString(1), L.ToString(2))
	return 0
}
func (p *VisualPrompt) Publish(L *lua.LState) int {
	if p.VisualPrompt.IsURL() {
		u, err := url.Parse(p.VisualPrompt.Source)
		if err != nil {
			panic(err)
		}
		if !p.bus.GetPluginOptions().MustAuthorizeDomain(u.Host) {
			panic(herbplugin.NewUnauthorizeDomainError(u.Host))
		}
	}
	ui := p.VisualPrompt.Publish(p.bus, L.ToString((1)))
	L.Push(lua.LString(ui.ID))
	return 1

}
func (p *VisualPrompt) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("setmediatype", L.NewFunction(p.SetMediaType))
	t.RawSetString("SetMediaType", L.NewFunction(p.SetMediaType))
	t.RawSetString("setportrait", L.NewFunction(p.SetPortrait))
	t.RawSetString("SetPortrait", L.NewFunction(p.SetPortrait))
	t.RawSetString("setvalue", L.NewFunction(p.SetValue))
	t.RawSetString("SetValue", L.NewFunction(p.SetValue))
	t.RawSetString("Append", L.NewFunction(p.Append))
	t.RawSetString("Append", L.NewFunction(p.Append))
	t.RawSetString("SetRefreshCallback", L.NewFunction(p.SetRefreshCallback))
	t.RawSetString("SetRefreshCallback", L.NewFunction(p.SetRefreshCallback))
	t.RawSetString("Publish", L.NewFunction(p.Publish))
	t.RawSetString("Publish", L.NewFunction(p.Publish))
	return t
}

type Datagrid struct {
	Datagrid *userinput.Datagrid
	bus      *bus.Bus
}

func (g *Datagrid) SetPage(L *lua.LState) int {
	g.Datagrid.SetPage(L.ToInt(1))
	return 0
}
func (g *Datagrid) GetPage(L *lua.LState) int {
	L.Push(lua.LNumber(g.Datagrid.GetPage()))
	return 1
}
func (g *Datagrid) SetMaxPage(L *lua.LState) int {
	g.Datagrid.SetMaxPage(L.ToInt(1))
	return 0
}
func (g *Datagrid) SetFilter(L *lua.LState) int {
	g.Datagrid.SetFilter(L.ToString(1))
	return 0
}
func (g *Datagrid) GetFilter(L *lua.LState) int {
	L.Push(lua.LString(g.Datagrid.GetFilter()))
	return 1
}
func (g *Datagrid) SetOnPage(L *lua.LState) int {
	g.Datagrid.SetOnPage(L.ToString(1))
	return 0
}
func (g *Datagrid) SetOnFilter(L *lua.LState) int {
	g.Datagrid.SetOnFilter(L.ToString(1))
	return 0
}
func (g *Datagrid) SetOnDelete(L *lua.LState) int {
	g.Datagrid.SetOnDelete(L.ToString(1))
	return 0
}
func (g *Datagrid) SetOnView(L *lua.LState) int {
	g.Datagrid.SetOnView(L.ToString(1))
	return 0
}
func (g *Datagrid) SetOnSelect(L *lua.LState) int {
	g.Datagrid.SetOnSelect(L.ToString(1))
	return 0
}
func (g *Datagrid) SetOnCreate(L *lua.LState) int {
	g.Datagrid.SetOnCreate(L.ToString(1))
	return 0
}
func (g *Datagrid) SetOnUpdate(L *lua.LState) int {
	g.Datagrid.SetOnUpdate(L.ToString(1))
	return 0
}
func (g *Datagrid) ResetItems(L *lua.LState) int {
	g.Datagrid.ResetItems()
	return 0
}
func (g *Datagrid) Append(L *lua.LState) int {
	g.Datagrid.Append(L.ToString(1), L.ToString(2))
	return 0
}
func (g *Datagrid) Publish(L *lua.LState) int {
	ui := g.Datagrid.Publish(g.bus, L.ToString(1))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (g *Datagrid) Hide(L *lua.LState) int {
	g.Datagrid.Hide(g.bus)
	return 0
}
func (g *Datagrid) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("append", L.NewFunction(g.Append))
	t.RawSetString("publish", L.NewFunction(g.Publish))
	t.RawSetString("resetitems", L.NewFunction(g.ResetItems))
	t.RawSetString("setoncreate", L.NewFunction(g.SetOnCreate))
	t.RawSetString("setonupdate", L.NewFunction(g.SetOnUpdate))
	t.RawSetString("setonview", L.NewFunction(g.SetOnView))
	t.RawSetString("setonselect", L.NewFunction(g.SetOnSelect))
	t.RawSetString("setondelete", L.NewFunction(g.SetOnDelete))
	t.RawSetString("setonfilter", L.NewFunction(g.SetOnFilter))
	t.RawSetString("setonpage", L.NewFunction(g.SetOnPage))
	t.RawSetString("setfilter", L.NewFunction(g.SetFilter))
	t.RawSetString("getfilter", L.NewFunction(g.GetFilter))
	t.RawSetString("setmaxpage", L.NewFunction(g.SetMaxPage))
	t.RawSetString("setpage", L.NewFunction(g.SetPage))
	t.RawSetString("getpage", L.NewFunction(g.GetPage))
	t.RawSetString("hide", L.NewFunction(g.Hide))

	t.RawSetString("Append", L.NewFunction(g.Append))
	t.RawSetString("Publish", L.NewFunction(g.Publish))
	t.RawSetString("ResetItems", L.NewFunction(g.ResetItems))
	t.RawSetString("SetOnCreate", L.NewFunction(g.SetOnCreate))
	t.RawSetString("SetOnUpdate", L.NewFunction(g.SetOnUpdate))
	t.RawSetString("SetOnView", L.NewFunction(g.SetOnView))
	t.RawSetString("SetOnSelect", L.NewFunction(g.SetOnSelect))
	t.RawSetString("SetOnDelete", L.NewFunction(g.SetOnDelete))
	t.RawSetString("SetOnFilter", L.NewFunction(g.SetOnFilter))
	t.RawSetString("SetOnPage", L.NewFunction(g.SetOnPage))
	t.RawSetString("SetFilter", L.NewFunction(g.SetFilter))
	t.RawSetString("GetFilter", L.NewFunction(g.GetFilter))
	t.RawSetString("SetMaxPage", L.NewFunction(g.SetMaxPage))
	t.RawSetString("SetPage", L.NewFunction(g.SetPage))
	t.RawSetString("GetPage", L.NewFunction(g.GetPage))
	t.RawSetString("Hide", L.NewFunction(g.Hide))
	return t
}

type List struct {
	List *userinput.List
	bus  *bus.Bus
}

func (l *List) Publish(L *lua.LState) int {
	ui := l.List.Publish(l.bus, L.ToString(1))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (l *List) Append(L *lua.LState) int {
	l.List.Append(L.ToString(1), L.ToString(2))
	return 0
}
func (l *List) SetValues(L *lua.LState) int {
	v := []string{}
	t := L.ToTable(1)
	m := t.MaxN()
	for i := 1; i <= m; i++ {
		v = append(v, t.RawGetInt(i).String())
	}
	l.List.SetValues(v)
	return 0
}
func (l *List) SetMutli(L *lua.LState) int {
	l.List.SetMutli(L.ToBool(1))
	return 0
}

func (l *List) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("append", L.NewFunction(l.Append))
	t.RawSetString("publish", L.NewFunction(l.Publish))
	t.RawSetString("setvalues", L.NewFunction(l.SetValues))
	t.RawSetString("setmutli", L.NewFunction(l.SetMutli))

	t.RawSetString("Append", L.NewFunction(l.Append))
	t.RawSetString("Publish", L.NewFunction(l.Publish))
	t.RawSetString("SetValues", L.NewFunction(l.SetValues))
	t.RawSetString("SetMutli", L.NewFunction(l.SetMutli))

	return t
}

type Userinput struct {
	bus *bus.Bus
}

func (u *Userinput) HideAll(L *lua.LState) int {
	userinput.HideAll(u.bus)
	return 0
}
func (u *Userinput) NewVisualPrompt(L *lua.LState) int {
	_ = L.Get(1) //this
	vp := &VisualPrompt{
		VisualPrompt: userinput.CreateVisualPrompt(L.ToString(2), L.ToString(3), L.ToString(4)),
		bus:          u.bus,
	}
	L.Push(vp.Convert(L))
	return 1
}
func (u *Userinput) NewDatagrid(L *lua.LState) int {
	_ = L.Get(1) //this
	datagrid := &Datagrid{
		Datagrid: userinput.CreateDatagrid(L.ToString(2), L.ToString(3)),
		bus:      u.bus,
	}
	L.Push(datagrid.Convert(L))
	return 1
}

func (u *Userinput) NewList(L *lua.LState) int {
	_ = L.Get(1) //this
	list := &List{
		List: userinput.CreateList(L.ToString(2), L.ToString(3), L.ToBool(4)),
		bus:  u.bus,
	}
	L.Push(list.Convert(L))
	return 1
}
func (u *Userinput) Prompt(L *lua.LState) int {
	_ = L.Get(1) //this
	ui := userinput.SendPrompt(u.bus, L.ToString(2), L.ToString(3), L.ToString(4), L.ToString(5))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Confirm(L *lua.LState) int {
	_ = L.Get(1) //this
	ui := userinput.SendConfirm(u.bus, L.ToString(2), L.ToString(3), L.ToString(4))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Alert(L *lua.LState) int {
	_ = L.Get(1) //this
	ui := userinput.SendAlert(u.bus, L.ToString(2), L.ToString(3), L.ToString(4))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Popup(L *lua.LState) int {
	_ = L.Get(1) //this
	ui := userinput.SendPopup(u.bus, L.ToString(2), L.ToString(3), L.ToString(4), L.ToString(5))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Note(L *lua.LState) int {
	_ = L.Get(1) //this
	ui := userinput.SendNote(u.bus, L.ToString(2), L.ToString(3), L.ToString(4), L.ToString(5))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Custom(L *lua.LState) int {
	_ = L.Get(1) //this
	ui := userinput.SendCustom(u.bus, L.ToString(2), L.ToString(3), L.ToString(4))
	L.Push(lua.LString(ui.ID))
	return 1
}
func (u *Userinput) Convert(L *lua.LState) lua.LValue {
	t := L.NewTable()
	t.RawSetString("prompt", L.NewFunction(u.Prompt))
	t.RawSetString("confirm", L.NewFunction(u.Confirm))
	t.RawSetString("alert", L.NewFunction(u.Alert))
	t.RawSetString("popup", L.NewFunction(u.Popup))
	t.RawSetString("note", L.NewFunction(u.Note))

	t.RawSetString("newlist", L.NewFunction(u.NewList))
	t.RawSetString("newdatagrid", L.NewFunction(u.NewDatagrid))
	t.RawSetString("newvisualprompt", L.NewFunction(u.NewVisualPrompt))
	t.RawSetString("hideall", L.NewFunction(u.HideAll))
	t.RawSetString("custom", L.NewFunction(u.Custom))

	t.RawSetString("Prompt", L.NewFunction(u.Prompt))
	t.RawSetString("Confirm", L.NewFunction(u.Confirm))
	t.RawSetString("Alert", L.NewFunction(u.Alert))
	t.RawSetString("Popup", L.NewFunction(u.Popup))
	t.RawSetString("Note", L.NewFunction(u.Note))

	t.RawSetString("NewList", L.NewFunction(u.NewList))
	t.RawSetString("NewDatagrid", L.NewFunction(u.NewDatagrid))
	t.RawSetString("NewVisualPrompt", L.NewFunction(u.NewVisualPrompt))
	t.RawSetString("HideAll", L.NewFunction(u.HideAll))
	t.RawSetString("Custom", L.NewFunction(u.Custom))

	return t
}
func NewUserinputModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("userinput",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			luapluing := plugin.(lua51plugin.LuaPluginLoader).LoadLuaPlugin()
			l := luapluing.LState
			u := &Userinput{bus: b}
			l.SetGlobal("Userinput", u.Convert(l))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
