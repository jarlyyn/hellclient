package bus

import (
	"sync"

	"github.com/herb-go/misc/busevent"
)

type Bus struct {
	ID     string
	Locker sync.RWMutex

	GetConnBuffer        func(bus *Bus) []byte
	GetConnConnected     func(bus *Bus) bool
	GetHost              func(bus *Bus) string
	GetPort              func(bus *Bus) string
	GetCharset           func(bus *Bus) string
	GetCurrentLines      func(bus *Bus) []*Line
	GetPrompt            func(bus *Bus) *Line
	DoSendToServer       func(bus *Bus, cmd []byte) error
	DoSend               func(bus *Bus, cmd []byte) error
	DoSave               func(bus *Bus) error
	DoLoad               func(bus *Bus) error
	DoPrint              func(bus *Bus, msg string)
	DoPrintSystem        func(bus *Bus, msg string)
	HandleConnReceive    func(bus *Bus, msg []byte)
	HandleConnError      func(bus *Bus, err error)
	HandleConnPrompt     func(bus *Bus, msg []byte)
	DoConnectServer      func(bus *Bus) error
	DoCloseServer        func(bus *Bus) error
	HandleConverterError func(bus *Bus, err error)
	HandleCmdError       func(bus *Bus, err error)
	lineEvent            *busevent.Event
	promptEvent          *busevent.Event
	connectedEvent       *busevent.Event
	disconnectedEvent    *busevent.Event
}

func (b *Bus) RaiseLineEvent(line *Line) {
	b.lineEvent.Raise(line)
}
func (b *Bus) BindLineEvent(id interface{}, fn func(b *Bus, line *Line)) {
	b.lineEvent.Bind(busevent.CreateHandler(
		id,
		func(data interface{}) {
			fn(b, data.(*Line))
		},
	))

}
func (b *Bus) UnbindLineEvent(id interface{}) {
	b.lineEvent.Unbind(id)
}

func (b *Bus) RaisePromptEvent(line *Line) {
	b.promptEvent.Raise(line)
}
func (b *Bus) BindPromptEvent(id interface{}, fn func(b *Bus, line *Line)) {
	b.promptEvent.Bind(busevent.CreateHandler(id,
		func(data interface{}) {
			fn(b, data.(*Line))
		}),
	)
}
func (b *Bus) UnbindPromptEvent(id interface{}) {
	b.promptEvent.Unbind(id)
}

func (b *Bus) RaiseContectedEvent() {
	b.connectedEvent.Raise(nil)
}
func (b *Bus) BindContectedEvent(id interface{}, fn func(b *Bus)) {
	b.connectedEvent.Bind(busevent.CreateHandler(id,
		func(data interface{}) {
			fn(b)
		}),
	)
}
func (b *Bus) UnbindContectedEvent(id interface{}) {
	b.connectedEvent.Unbind(id)
}

func (b *Bus) RaiseDiscontectedEvent() {
	b.disconnectedEvent.Raise(nil)
}
func (b *Bus) BindDiscontectedEvent(id interface{}, fn func(b *Bus)) {
	b.disconnectedEvent.Bind(busevent.CreateHandler(id,
		func(data interface{}) {
			fn(b)
		}),
	)
}
func (b *Bus) UnbindDiscontectedEvent(id interface{}) {
	b.disconnectedEvent.Unbind(id)
}

func New() *Bus {
	return &Bus{
		lineEvent:         busevent.New(),
		promptEvent:       busevent.New(),
		connectedEvent:    busevent.New(),
		disconnectedEvent: busevent.New(),
	}
}
