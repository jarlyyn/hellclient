package bus

import (
	"sync"

	"github.com/herb-go/misc/busevent"
)

type Bus struct {
	ID               string
	Locker           sync.RWMutex
	GetHostPort      func() string
	GetConnBuffer    func() []byte
	GetConnConnected func() bool
	GetCharset       func() string
	GetCurrentLines  func() []*Line
	DoSendToServer   func(cmd []byte) error
	DoSend           func(cmd []byte) error
	OnConnReceive    func(msg []byte)
	OnConnError      func(err error)
	OnConnPrompt     func(msg []byte)
	DoConnectServer  func() error
	DoCloseServer    func() error
	OnConverterError func(err error)
	lineEvent        *busevent.Event
	promptEvent      *busevent.Event
}

func (b *Bus) RaiseLineEvent(line *Line) {
	b.lineEvent.Raise(line)
}
func (b *Bus) BindLineEvent(id string, fn func(line *Line)) {
	b.lineEvent.Bind(busevent.CreateHandler(
		id,
		func(data interface{}) {
			fn(data.(*Line))
		},
	))

}
func (b *Bus) UnbindLineEvent(id string) {
	b.lineEvent.Unbind(id)
}
func (b *Bus) RaisePromptEvent(line *Line) {
	b.promptEvent.Raise(line)
}
func (b *Bus) BindPromptEvent(id string, fn func(line *Line)) {
	b.promptEvent.Bind(busevent.CreateHandler(id,
		func(data interface{}) {
			fn(data.(*Line))
		}),
	)
}
func (b *Bus) UnbindPromptEvent(id string) {
	b.promptEvent.Unbind(id)
}

func New() *Bus {
	return &Bus{
		lineEvent:   busevent.New(),
		promptEvent: busevent.New(),
	}
}
