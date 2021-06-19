package info

import (
	"container/ring"
	"modules/world/bus"
	"sync"
)

type Info struct {
	Lines  *ring.Ring
	Prompt *bus.Line
	Lock   sync.RWMutex
}

func (i *Info) Init(b *bus.Bus) {
	i.Lines = ring.New(1000)
}
func (i *Info) ClientInfo(b *bus.Bus) *bus.ClientInfo {
	info := &bus.ClientInfo{}
	info.ID = b.ID
	info.Running = b.GetConnConnected()
	return info
}
func (i *Info) CurrentLines(b *bus.Bus) []*bus.Line {
	result := []*bus.Line{}
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	i.Lines.Do(func(x interface{}) {
		line, ok := x.(*bus.Line)
		if ok && line != nil {
			result = append(result, line)
		}
	})
	// result = append(result, i.Prompt)
	return result
}
func (i *Info) CurrentPrompt(b *bus.Bus) *bus.Line {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	return i.Prompt
}
func (i *Info) onPrompt(b *bus.Bus, line *bus.Line) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.Prompt = line
}
func (i *Info) onNewLine(b *bus.Bus, line *bus.Line) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	if line.OmitFromOutput {
		return
	}
	i.Lines.Value = line
	i.Lines = i.Lines.Next()
}
func (i *Info) InstallTo(b *bus.Bus) {
	b.GetCurrentLines = b.WrapGetLines(i.CurrentLines)
	b.GetPrompt = b.WrapGetLine(i.CurrentPrompt)
	b.GetClientInfo = b.WrapGetClientInfo(i.ClientInfo)
	b.BindLineEvent(i, i.onNewLine)
	b.BindPromptEvent(i, i.onPrompt)
	b.BindReadyEvent(i, i.Init)
}

func New() *Info {
	return &Info{}
}
