package userinput

import (
	"modules/world/bus"
)

type ListItem struct {
	Key   string
	Value string
}
type List struct {
	Title      string
	Intro      string
	Items      []*ListItem
	WithFilter bool
}

func (l *List) Append(key string, value string) {
	l.Items = append(l.Items, &ListItem{Key: key, Value: value})
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
	}
}
