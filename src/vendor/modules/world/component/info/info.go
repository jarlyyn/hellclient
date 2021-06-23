package info

import (
	"container/ring"
	"modules/world"
	"modules/world/bus"
	"sync"
)

type Info struct {
	Lines   *ring.Ring
	History *ring.Ring
	Prompt  *world.Line
	Lock    sync.RWMutex
}

const MaxHistory = 20
const MaxLines = 1000

func (i *Info) Init(b *bus.Bus) {
	i.Lines = ring.New(MaxLines)
	i.History = ring.New(MaxHistory)
}
func (i *Info) ClientInfo(b *bus.Bus) *world.ClientInfo {
	info := &world.ClientInfo{}
	info.ID = b.ID
	info.HostPort = b.GetHost() + ":" + b.GetPort()
	info.ReadyAt = b.GetReadyAt()
	info.Running = b.GetConnConnected()
	info.ScriptID = b.GetScriptID()
	return info
}
func (i *Info) CurrentLines(b *bus.Bus) []*world.Line {
	result := []*world.Line{}
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	i.Lines.Do(func(x interface{}) {
		line, ok := x.(*world.Line)
		if ok && line != nil {
			result = append(result, line)
		}
	})
	// result = append(result, i.Prompt)
	return result
}
func (i *Info) CurrentPrompt(b *bus.Bus) *world.Line {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	return i.Prompt
}
func (i *Info) onPrompt(b *bus.Bus, line *world.Line) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.Prompt = line
}
func (i *Info) onNewLine(b *bus.Bus, line *world.Line) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	if line.OmitFromOutput {
		return
	}
	i.Lines.Value = line
	i.Lines = i.Lines.Next()
}
func (i *Info) AddHistory(b *bus.Bus, cmd string) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.History.Value = cmd
	i.History = i.History.Next()
	i.CurrentHistories(b)
}

func (i *Info) getHistories() []string {
	var result = make([]string, i.History.Len())
	i.History.Do(func(x interface{}) {
		data, ok := x.(string)
		if ok {
			result = append(result, data)
		}
	})
	return result
}
func (i *Info) GetHistories(b *bus.Bus) []string {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.CurrentHistories(b)
	return i.getHistories()
}
func (i *Info) FlushHistories(b *bus.Bus) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.History = ring.New(MaxHistory)
	i.CurrentHistories(b)
}
func (i *Info) CurrentHistories(b *bus.Bus) {
	go func() {
		b.RaiseHistoriesEvent(i.getHistories())
	}()
}
func (i *Info) InstallTo(b *bus.Bus) {
	b.GetCurrentLines = b.WrapGetLines(i.CurrentLines)
	b.GetPrompt = b.WrapGetLine(i.CurrentPrompt)
	b.GetClientInfo = b.WrapGetClientInfo(i.ClientInfo)
	b.AddHistory = b.WrapHandleString(i.AddHistory)
	b.GetHistories = b.WrapGetStrings(i.GetHistories)
	b.FlushHistories = b.Wrap(i.FlushHistories)
	b.BindLineEvent(i, i.onNewLine)
	b.BindPromptEvent(i, i.onPrompt)
	b.BindInitEvent(i, i.Init)
}

func New() *Info {
	return &Info{}
}
