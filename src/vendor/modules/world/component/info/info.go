package info

import (
	"container/ring"
	"modules/world"
	"modules/world/bus"
	"sync"
)

type Info struct {
	Lines     *ring.Ring
	History   *ring.Ring
	Prompt    *world.Line
	Lock      sync.RWMutex
	LineCount int
}

const MaxHistory = 20
const MaxLines = 1000

func (i *Info) OmitOutput(b *bus.Bus) {
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	l := i.Lines.Value
	if l != nil {
		l.(*world.Line).OmitFromOutput = true
	}
	// i.linesUpdated(b)

}
func (i *Info) DeleteLines(b *bus.Bus, count int) {
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	l := i.Lines.Len()
	if count > l {
		count = l
	}
	var removed = 0
	for removed < count {
		i.Lines.Value = nil
		i.Lines = i.Lines.Prev()
		removed++
	}
	i.linesUpdated(b)
}
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

func (i *Info) linesUpdated(b *bus.Bus) {
	b.RaiseLinesEvent(i.lines())
}
func (i *Info) lines() []*world.Line {
	result := []*world.Line{}
	i.Lines.Next().Do(func(x interface{}) {
		line, ok := x.(*world.Line)
		if ok && line != nil && !line.OmitFromOutput {
			result = append(result, line)
		}
	})
	return result

}
func (i *Info) CurrentLines(b *bus.Bus) []*world.Line {
	i.Lock.RLock()
	defer i.Lock.RUnlock()
	return i.lines()
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
	i.Lines = i.Lines.Next()
	i.Lines.Value = line
	i.LineCount++
}
func (i *Info) GetLineCount() int {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	return i.LineCount
}
func (i *Info) AddHistory(b *bus.Bus, cmd string) {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	i.History.Value = cmd
	i.History = i.History.Next()
	i.CurrentHistories(b)
}
func (i *Info) GetRecentLines(count int) []*world.Line {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	if count < 0 {
		count = 0
	}
	if count > i.Lines.Len() {
		count = i.Lines.Len()
	}
	result := make([]*world.Line, 0, count)

	r := i.Lines.Move(1 - count)
	var current = 0
	for current < count {
		current = current + 1
		v := r.Value
		if v != nil {
			result = append(result, v.(*world.Line))
		}
		r = r.Next()
	}
	return result

}
func (i *Info) getHistories() []string {
	var result = make([]string, 0, i.History.Len())
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

func (i *Info) GetLinesInBufferCount() int {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	return i.Lines.Len()
}

func (i *Info) GetLine(idx int) *world.Line {
	i.Lock.Lock()
	defer i.Lock.Unlock()
	if idx < 0 || idx > i.Lines.Len() {
		return nil
	}
	r := i.Lines.Move(idx)
	v := r.Value
	if v == nil {
		return nil
	}
	return v.(*world.Line)
}
func (i *Info) InstallTo(b *bus.Bus) {
	b.GetCurrentLines = b.WrapGetLines(i.CurrentLines)
	b.GetPrompt = b.WrapGetLine(i.CurrentPrompt)
	b.GetClientInfo = b.WrapGetClientInfo(i.ClientInfo)
	b.AddHistory = b.WrapHandleString(i.AddHistory)
	b.GetHistories = b.WrapGetStrings(i.GetHistories)
	b.FlushHistories = b.Wrap(i.FlushHistories)
	b.DoOmitOutput = b.Wrap(i.OmitOutput)
	b.DoDeleteLines = b.WrapHandleInt(i.DeleteLines)
	b.GetLineCount = i.GetLineCount
	b.GetRecentLines = i.GetRecentLines
	b.GetLinesInBufferCount = i.GetLinesInBufferCount
	b.GetLine = i.GetLine
	b.BindLineEvent(i, i.onNewLine)
	b.BindPromptEvent(i, i.onPrompt)
	b.BindInitEvent(i, i.Init)
}

func New() *Info {
	return &Info{}
}
