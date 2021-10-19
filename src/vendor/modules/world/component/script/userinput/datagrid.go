package userinput

import "modules/world/bus"

type Datagrid struct {
	Title    string
	Intro    string
	Items    []*Item
	MaxPage  int
	Page     int
	Filter   string
	OnPage   string
	OnFilter string
	OnDelete string
	OnUpdate string
	OnView   string
	OnCreate string
}

func (g *Datagrid) Hide(b *bus.Bus) {
	b.RaiseScriptMessageEvent(CreateUserInput(NameHideDatagrid, "", nil))
}
func (g *Datagrid) SetPage(page int) {
	g.Page = page
}
func (g *Datagrid) GetPage() int {
	return g.Page
}
func (g *Datagrid) SetMaxPage(page int) {
	g.MaxPage = page
}
func (g *Datagrid) SetFilter(filter string) {
	g.Filter = filter
}
func (g *Datagrid) GetFilter() string {
	return g.Filter
}
func (g *Datagrid) SetOnPage(onpage string) {
	g.OnPage = onpage
}
func (g *Datagrid) SetOnFilter(onfilter string) {
	g.OnFilter = onfilter
}
func (g *Datagrid) SetOnDelete(ondelete string) {
	g.OnDelete = ondelete
}
func (g *Datagrid) SetOnView(onview string) {
	g.OnView = onview
}
func (g *Datagrid) SetOnCreate(oncreate string) {
	g.OnCreate = oncreate
}
func (g *Datagrid) SetOnUpdate(onupdate string) {
	g.OnUpdate = onupdate
}
func (g *Datagrid) ResetItems() {
	g.Items = []*Item{}
}
func (g *Datagrid) Append(key string, value string) {
	g.Items = append(g.Items, &Item{Key: key, Value: value})
}
func (g *Datagrid) Publish(b *bus.Bus, script string) *Userinput {
	ui := CreateUserInput(NameDatagrid, script, g)
	b.RaiseScriptMessageEvent(ui)
	return ui
}
func CreateDatagrid(title string, intro string) *Datagrid {
	return &Datagrid{
		Title: title,
		Intro: intro,
		Page:  1,
	}
}
