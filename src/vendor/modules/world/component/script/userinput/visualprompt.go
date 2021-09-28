package userinput

import "modules/world/bus"

const DefaultRatio = 0.75
const MediaTypeImage = "image"

type VisualPrompt struct {
	Title           string
	Intro           string
	Source          string
	MediaType       string
	Ratio           float32
	RefreshCallback string
}

func (p *VisualPrompt) SetMediaType(t string) {
	p.MediaType = t
}
func (p *VisualPrompt) SetRatio(r float32) {
	p.Ratio = r
}
func (p *VisualPrompt) SetRefreshCallback(c string) {
	p.RefreshCallback = c
}
func (p *VisualPrompt) Send(b *bus.Bus, script string) *Userinput {
	ui := CreateUserInput(NameVisualPrompt, script, p)
	b.RaiseScriptMessageEvent(ui)
	return ui
}

func CreateVisualPrompt(title string, intro string, source string) *VisualPrompt {
	return &VisualPrompt{
		Title:     title,
		Intro:     intro,
		Source:    source,
		Ratio:     DefaultRatio,
		MediaType: MediaTypeImage,
	}
}
