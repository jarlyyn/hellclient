package userinput

import (
	"hellclient/modules/world/bus"
	"strings"
)

const MediaTypeImage = "image"

type VisualPrompt struct {
	Title           string
	Intro           string
	Source          string
	MediaType       string
	Portrait        bool
	RefreshCallback string
}

func (p *VisualPrompt) IsURL() bool {
	t := strings.ToLower(p.MediaType)
	return t == "" || t == "image"
}
func (p *VisualPrompt) SetMediaType(t string) {
	p.MediaType = t
}
func (p *VisualPrompt) SetPortrait(v bool) {
	p.Portrait = v
}
func (p *VisualPrompt) SetRefreshCallback(c string) {
	p.RefreshCallback = c
}
func (p *VisualPrompt) Publish(b *bus.Bus, script string) *Userinput {
	ui := CreateUserInput(NameVisualPrompt, script, p)
	b.RaiseScriptMessageEvent(ui)
	return ui
}

func CreateVisualPrompt(title string, intro string, source string) *VisualPrompt {
	return &VisualPrompt{
		Title:     title,
		Intro:     intro,
		Source:    source,
		Portrait:  false,
		MediaType: MediaTypeImage,
	}
}
