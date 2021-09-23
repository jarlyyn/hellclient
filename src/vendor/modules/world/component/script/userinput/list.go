package userinput

import (
	"modules/world/bus"
)

type List struct {
	Title      string
	Intro      string
	Items      []*Item
	Mutli      bool
	Values     []string
	WithFilter bool
}

func (l *List) SetValues(v []string) {
	l.Values = v
}
func (l *List) SetMutli(m bool) {
	l.Mutli = m
}
func (l *List) Append(key string, value string) {
	l.Items = append(l.Items, &Item{Key: key, Value: value})
}
func (l *List) Send(b *bus.Bus, script string) *Userinput {
	ui := CreateUserInput(NameList, script, l)
	b.RaiseScriptMessageEvent(ui)
	return ui
}
func CreateList(title string, intro string, withfilter bool) *List {
	return &List{
		Title:      title,
		Intro:      intro,
		WithFilter: withfilter,
		Values:     []string{},
	}
}
