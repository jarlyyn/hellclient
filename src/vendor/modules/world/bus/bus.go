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
	SetHost              func(string)
	GetPort              func() string
	SetPort              func(string)
	GetCharset           func() string
	SetCharset           func(string)
	GetCurrentLines      func() []*Line
	GetPrompt            func() *Line
	GetClientInfo        func() *ClientInfo
	DoSendToServer       func(cmd []byte) error
	DoSend               func(cmd []byte) error
	DoEncode             func() ([]byte, error)
	DoDecode             func([]byte) error
	DoPrint              func(msg string)
	DoPrintSystem        func(msg string)
	HandleConnReceive    func(msg []byte)
	HandleConnError      func(err error)
	HandleConnPrompt     func(msg []byte)
	DoConnectServer      func() error
	DoCloseServer        func() error
	HandleConverterError func(err error)
	HandleCmdError       func(err error)
	LineEvent            busevent.Event
	PromptEvent          busevent.Event
	ConnectedEvent       busevent.Event
	DisconnectedEvent    busevent.Event
	ReadyEvent           busevent.Event
	CloseEvent           busevent.Event
}

func (b *Bus) WrapDo(f func(bus *Bus) error) func() error {
	return func() error {
		return f(b)
	}
}
func (b *Bus) WrapDoCmd(f func(bus *Bus, cmd []byte) error) func(cmd []byte) error {
	return func(cmd []byte) error {
		return f(b, cmd)
	}
}
func (b *Bus) WrapGetString(f func(bus *Bus) string) func() string {
	return func() string {
		return f(b)
	}
}
func (b *Bus) WrapHandleError(f func(bus *Bus, err error)) func(err error) {
	return func(err error) {
		f(b, err)
	}
}
func (b *Bus) WrapGetLine(f func(bus *Bus) *Line) func() *Line {
	return func() *Line {
		return f(b)
	}
}
func (b *Bus) WrapGetLines(f func(bus *Bus) []*Line) func() []*Line {
	return func() []*Line {
		return f(b)
	}
}
func (b *Bus) WrapGetBool(f func(bus *Bus) bool) func() bool {
	return func() bool {
		return f(b)
	}
}
func (b *Bus) WrapGetClientInfo(f func(bus *Bus) *ClientInfo) func() *ClientInfo {
	return func() *ClientInfo {
		return f(b)
	}
}
func (b *Bus) WrapHandleBytes(f func(bus *Bus, bs []byte)) func(bs []byte) {
	return func(bs []byte) {
		f(b, bs)
	}
}
func (b *Bus) WrapHandleString(f func(bus *Bus, s string)) func(s string) {
	return func(s string) {
		f(b, s)
	}
}
func (b *Bus) RaiseLineEvent(line *Line) {
	b.LineEvent.Raise(line)
}
func (b *Bus) BindLineEvent(id interface{}, fn func(b *Bus, line *Line)) {
	b.LineEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*Line))
		},
	)

}

func (b *Bus) RaisePromptEvent(line *Line) {
	b.PromptEvent.Raise(line)
}
func (b *Bus) BindPromptEvent(id interface{}, fn func(b *Bus, line *Line)) {
	b.PromptEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*Line))
		},
	)
}

func (b *Bus) RaiseContectedEvent() {
	b.ConnectedEvent.Raise(nil)
}
func (b *Bus) BindContectedEvent(id interface{}, fn func(b *Bus)) {
	b.ConnectedEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b)
		},
	)
}

func (b *Bus) RaiseDiscontectedEvent() {
	b.DisconnectedEvent.Raise(nil)
}
func (b *Bus) BindDiscontectedEvent(id interface{}, fn func(b *Bus)) {
	b.DisconnectedEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b)
		},
	)
}
func (b *Bus) RaiseCloseEvent() {
	b.CloseEvent.Raise(nil)
}
func (b *Bus) BindCloseEvent(id interface{}, fn func(b *Bus)) {
	b.CloseEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b)
		},
	)
}
func (b *Bus) RaiseReadyEvent() {
	b.ReadyEvent.Raise(nil)
}
func (b *Bus) BindReadyEvent(id interface{}, fn func(b *Bus)) {
	b.ReadyEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b)
		},
	)
}
func (b *Bus) Reset() {
	*b = *New()
}
func New() *Bus {
	return &Bus{}
}
