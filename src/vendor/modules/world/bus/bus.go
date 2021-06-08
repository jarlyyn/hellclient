package bus

import (
	"sync"

	"github.com/herb-go/misc/busevent"
)

type Bus struct {
	ID     string
	Locker sync.RWMutex

	GetConnBuffer        func() []byte
	GetConnConnected     func() bool
	GetHost              func() string
	GetPort              func() string
	GetCharset           func() string
	GetCurrentLines      func() []*Line
	GetPrompt            func() *Line
	DoSendToServer       func(cmd []byte) error
	DoSend               func(cmd []byte) error
	DoSave               func() error
	DoLoad               func() error
	DoPrint              func(string)
	DoPrintSystem        func(string)
	HandleConnReceive    func(msg []byte)
	HandleConnError      func(err error)
	HandleConnPrompt     func(msg []byte)
	DoConnectServer      func() error
	DoCloseServer        func() error
	HandleConverterError func(err error)
	HandleCmdError       func(err error)
	lineEvent            *busevent.Event
	promptEvent          *busevent.Event
	connectedEvent       *busevent.Event
	disconnectedEvent    *busevent.Event
}

func (b *Bus) RaiseLineEvent(line *Line) {
	b.lineEvent.Raise(line)
}
func (b *Bus) BindLineEvent(id string, fn func(b *Bus, line *Line)) {
	b.lineEvent.Bind(busevent.CreateHandler(
		id,
		func(data interface{}) {
			fn(b, data.(*Line))
		},
	))

}
func (b *Bus) UnbindLineEvent(id string) {
	b.lineEvent.Unbind(id)
}

func (b *Bus) RaisePromptEvent(line *Line) {
	b.promptEvent.Raise(line)
}
func (b *Bus) BindPromptEvent(id string, fn func(b *Bus, line *Line)) {
	b.promptEvent.Bind(busevent.CreateHandler(id,
		func(data interface{}) {
			fn(b, data.(*Line))
		}),
	)
}
func (b *Bus) UnbindPromptEvent(id string) {
	b.promptEvent.Unbind(id)
}

func (b *Bus) RaiseContectedEvent() {
	b.connectedEvent.Raise(nil)
}
func (b *Bus) BindContectedEvent(id string, fn func(b *Bus)) {
	b.connectedEvent.Bind(busevent.CreateHandler(id,
		func(data interface{}) {
			fn(b)
		}),
	)
}
func (b *Bus) UnbindContectedEvent(id string) {
	b.connectedEvent.Unbind(id)
}

func (b *Bus) RaiseDiscontectedEvent() {
	b.disconnectedEvent.Raise(nil)
}
func (b *Bus) BindDiscontectedEvent(id string, fn func(b *Bus)) {
	b.disconnectedEvent.Bind(busevent.CreateHandler(id,
		func(data interface{}) {
			fn(b)
		}),
	)
}
func (b *Bus) UnbindDiscontectedEvent(id string) {
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
