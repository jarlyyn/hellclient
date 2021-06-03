package visual

import (
	"container/ring"
	"modules/world/bus"
	"sync"
)

type Visual struct {
	bus    *bus.Bus
	Lines  *ring.Ring
	Prompt *bus.Line
	Lock   sync.RWMutex
}

func (v *Visual) InstallTo(b *bus.Bus) {
	v.bus = b
	b.GetVisualCurrentLines = v.CurrentLines
	b.DoSend = v.Send
}

func (v *Visual) UninstallFrom(b *bus.Bus) {
	if v.bus != b {
		return
	}
	b.GetVisualCurrentLines = nil
	b.DoSend = nil
}

func (v *Visual) Init() {
	v.Lines = ring.New(1000)
}
func (v *Visual) onError(err error) {
	v.bus.OnVisualError(err)
}

func (v *Visual) Send(cmd []byte) error {
	v.Lock.Lock()
	defer v.Lock.Unlock()
	b, err := FromUTF8(v.bus.GetCharset(), []byte(cmd))
	if err != nil {
		return err
	}
	return v.bus.DoSendToServer(b)
}
func (v *Visual) CurrentLines() []*bus.Line {
	result := []*bus.Line{}
	v.Lock.RLock()
	defer v.Lock.RUnlock()
	v.Lines.Do(func(x interface{}) {
		line, ok := x.(*bus.Line)
		if ok && line != nil {
			result = append(result, line)
		}
	})
	return result
}

func (v *Visual) ConvertToLine(msg []byte) *bus.Line {
	charset := v.bus.GetCharset()
	return ConvertToLine(msg, charset, v.onError)
}

func New() *Visual {
	return &Visual{}
}
