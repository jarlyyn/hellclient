package titan

import (
	"modules/msg"
	"modules/world/bus"
	"sync"

	"github.com/herb-go/connections/room/message"
	"github.com/herb-go/misc/busevent"
)

type Titan struct {
	Locker   sync.RWMutex
	Worlds   map[string]*bus.Bus
	msgEvent *busevent.Event
}

func (t *Titan) find(id string) *bus.Bus {
	return t.Worlds[id]
}

func (t *Titan) World(id string) *bus.Bus {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	return t.find(id)
}

func (t *Titan) DoSendTo(id string, msg []byte) error {
	w := t.World(id)
	return w.DoSend(msg)
}
func (t *Titan) Publish(msg *message.Message) {
	go func() {
		t.msgEvent.Raise(msg)
	}()
}

func (t *Titan) onConnected(b *bus.Bus) {
	msg.PublishConnected(t, b.ID)
}
func (t *Titan) onDisonnected(b *bus.Bus) {
	msg.PublishDisconnected(t, b.ID)
}
func (t *Titan) onPrompt(b *bus.Bus, prompt *bus.Line) {
	msg.PublishPrompt(t, b.ID, prompt)
}
func (t *Titan) onLine(b *bus.Bus, line *bus.Line) {
	msg.PublishLine(t, b.ID, line)
}

func (t *Titan) HandleCmdConnect(id string) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoConnectServer())
	}
}
func (t *Titan) HandleCmdDisconnect(id string) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoCloseServer())
	}
}
func (t *Titan) HandleCmdSend(id string, msg []byte) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoSend(msg))
	}
}
func (t *Titan) HandleCmdAllLines(id string) {
	w := t.World(id)
	if w != nil {
		alllines := w.GetCurrentLines()
		msg.PublishAllLines(t, id, alllines)
	}
}
func (t *Titan) HandleCmdLines(id string) {
	w := t.World(id)
	if w != nil {
		alllines := w.GetCurrentLines()
		msg.PublishLines(t, id, alllines)
	}
}
func (t *Titan) HandleCmdPrompt(id string) {
	w := t.World(id)
	if w != nil {
		pormpt := w.GetPrompt()
		msg.PublishPrompt(t, id, pormpt)
	}
}
func (t *Titan) InstallTo(b *bus.Bus) {
	b.BindContectedEvent("titan.oncontected", t.onConnected)
	b.BindDiscontectedEvent("titan.ondiscontected", t.onConnected)
	b.BindLineEvent("titan.online", t.onLine)
	b.BindLineEvent("titan.onprompt", t.onPrompt)
}

func (t *Titan) RaiseMsgEvent(msg *message.Message) {
	t.msgEvent.Raise(msg)
}
func (t *Titan) BindMsgEvent(id interface{}, fn func(t *Titan, msg *message.Message)) {
	t.msgEvent.Bind(busevent.CreateHandler(
		id,
		func(data interface{}) {
			fn(t, data.(*message.Message))
		},
	))

}
func (t *Titan) UnbindMsgEvent(id interface{}) {
	t.msgEvent.Unbind(id)
}

func New() *Titan {
	return &Titan{
		msgEvent: busevent.New(),
	}
}

var Pangu = New()
