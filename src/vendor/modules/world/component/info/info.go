package info

import (
	"container/ring"
	"modules/world/bus"
	"sync"
)

type Info struct {
	bus    *bus.Bus
	Lines  *ring.Ring
	Prompt *bus.Line
	Lock   sync.RWMutex
}

func (i *Info) Init() {
	i.Lines = ring.New(1000)
}
func (i *Info) CurrentLines() []*bus.Line {
	result := []*bus.Line{}
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	i.Lines.Do(func(x interface{}) {
		line, ok := x.(*bus.Line)
		if ok && line != nil {
			result = append(result, line)
		}
	})
	result = append(result, i.Prompt)
	return result
}
func (i *Info) onPrompt(b *bus.Bus, line *bus.Line) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.Prompt = line
}
func (i *Info) onNewLine(b *bus.Bus, line *bus.Line) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.Lines.Value = line
	i.Lines = i.Lines.Next()
}
func (i *Info) InstallTo(b *bus.Bus) {
	i.bus = b
	b.GetCurrentLines = i.CurrentLines
	b.BindLineEvent("info.onnewline", i.onNewLine)
	b.BindPromptEvent("info.onprompt", i.onPrompt)
}

func (i *Info) UninstallFrom(b *bus.Bus) {
	if i.bus != b {
		return
	}
	b.GetCurrentLines = i.CurrentLines
	b.UnbindLineEvent("info.onnewline")
	b.UnbindLineEvent("info.onprompt")
}
