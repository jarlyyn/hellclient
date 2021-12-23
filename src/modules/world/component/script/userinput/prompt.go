package userinput

import (
	"hellclient/modules/world/bus"

	"github.com/herb-go/uniqueid"
)

type Userinput struct {
	Name   string
	Script string
	ID     string
	Data   interface{}
}

func CreateUserInput(name string, script string, data interface{}) *Userinput {
	return &Userinput{
		Name:   name,
		Script: script,
		ID:     uniqueid.MustGenerateID(),
		Data:   data,
	}
}
func SendConfirm(b *bus.Bus, script string, title string, intro string) *Userinput {
	data := map[string]interface{}{
		"Title": title,
		"Intro": intro,
	}
	ui := CreateUserInput(NameConfirm, script, data)
	b.RaiseScriptMessageEvent(ui)
	return ui
}

func SendAlert(b *bus.Bus, script string, title string, intro string) *Userinput {
	data := map[string]interface{}{
		"Title": title,
		"Intro": intro,
	}
	ui := CreateUserInput(NameAlert, script, data)
	b.RaiseScriptMessageEvent(ui)
	return ui

}

func SendPrompt(b *bus.Bus, script string, title string, intro string, value string) *Userinput {
	data := map[string]interface{}{
		"Title": title,
		"Intro": intro,
		"Value": value,
	}
	ui := CreateUserInput(NamePrompt, script, data)
	b.RaiseScriptMessageEvent(ui)
	return ui

}
