package bus

import (
	"modules/world"
	"sync"

	"github.com/herb-go/herbplugin"

	"github.com/herb-go/misc/busevent"
)

type Bus struct {
	ID     string
	Locker sync.RWMutex

	GetConnBuffer          func() []byte
	GetConnConnected       func() bool
	GetHost                func() string
	SetHost                func(string)
	GetPort                func() string
	SetPort                func(string)
	GetQueueDelay          func() int
	SetQueueDelay          func(int)
	GetParam               func(string) string
	GetParams              func() map[string]string
	SetParam               func(string, string)
	DeleteParam            func(string)
	GetCharset             func() string
	SetCharset             func(string)
	GetReadyAt             func() int64
	GetCurrentLines        func() []*world.Line
	GetPrompt              func() *world.Line
	GetClientInfo          func() *world.ClientInfo
	GetScriptData          func() *world.ScriptData
	SetPermissions         func([]string)
	GetPermissions         func() []string
	GetScriptID            func() string
	SetScriptID            func(string)
	GetScriptPath          func() string
	SetTrusted             func(*herbplugin.Trusted)
	GetTrusted             func() *herbplugin.Trusted
	GetScriptPluginOptions func() herbplugin.Options
	DoSendToConn           func(cmd []byte)
	DoSend                 func(cmd []byte)
	DoSendToQueue          func(cmd []byte)
	DoEncode               func() ([]byte, error)
	DoDecode               func([]byte) error
	DoUnloadScript         func()
	DoLoadScript           func() error
	DoSaveScript           func() error
	DoUseScript            func(string)
	DoPrint                func(msg string)
	DoPrintSystem          func(msg string)
	DoDiscardQueue         func()
	HandleConnReceive      func(msg []byte)
	HandleConnError        func(err error)
	HandleConnPrompt       func(msg []byte)
	DoConnectServer        func() error
	DoCloseServer          func() error
	HandleConverterError   func(err error)
	HandleCmdError         func(err error)
	HandleTriggerError     func(err error)
	HandleScriptError      func(err error)

	LineEvent         busevent.Event
	PromptEvent       busevent.Event
	ConnectedEvent    busevent.Event
	DisconnectedEvent busevent.Event
	InitEvent         busevent.Event
	ReadyEvent        busevent.Event
	BeforeCloseEvent  busevent.Event
	CloseEvent        busevent.Event
}

func (b *Bus) Wrap(f func(bus *Bus)) func() {
	return func() {
		f(b)
	}
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
func (b *Bus) WrapGetLine(f func(bus *Bus) *world.Line) func() *world.Line {
	return func() *world.Line {
		return f(b)
	}
}
func (b *Bus) WrapGetLines(f func(bus *Bus) []*world.Line) func() []*world.Line {
	return func() []*world.Line {
		return f(b)
	}
}
func (b *Bus) WrapGetBool(f func(bus *Bus) bool) func() bool {
	return func() bool {
		return f(b)
	}
}
func (b *Bus) WrapGetInt(f func(bus *Bus) int) func() int {
	return func() int {
		return f(b)
	}
}
func (b *Bus) WrapGetClientInfo(f func(bus *Bus) *world.ClientInfo) func() *world.ClientInfo {
	return func() *world.ClientInfo {
		return f(b)
	}
}
func (b *Bus) WrapGetScriptData(f func(bus *Bus) *world.ScriptData) func() *world.ScriptData {
	return func() *world.ScriptData {
		return f(b)
	}
}
func (b *Bus) WrapGetScriptPluginOptions(f func(bus *Bus) herbplugin.Options) func() herbplugin.Options {
	return func() herbplugin.Options {
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
func (b *Bus) WrapHandleInt(f func(bus *Bus, i int)) func(i int) {
	return func(i int) {
		f(b, i)
	}
}
func (b *Bus) RaiseLineEvent(line *world.Line) {
	b.LineEvent.Raise(line)
}
func (b *Bus) BindLineEvent(id interface{}, fn func(b *Bus, line *world.Line)) {
	b.LineEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*world.Line))
		},
	)

}

func (b *Bus) RaisePromptEvent(line *world.Line) {
	b.PromptEvent.Raise(line)
}
func (b *Bus) BindPromptEvent(id interface{}, fn func(b *Bus, line *world.Line)) {
	b.PromptEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*world.Line))
		},
	)
}

func (b *Bus) RaiseConnectedEvent() {
	b.ConnectedEvent.Raise(nil)
}
func (b *Bus) BindConnectedEvent(id interface{}, fn func(b *Bus)) {
	b.ConnectedEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b)
		},
	)
}

func (b *Bus) RaiseDisconnectedEvent() {
	b.DisconnectedEvent.Raise(nil)
}
func (b *Bus) BindDisconnectedEvent(id interface{}, fn func(b *Bus)) {
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
func (b *Bus) RaiseBeforeCloseEvent() {
	b.BeforeCloseEvent.Raise(nil)
}
func (b *Bus) BindBeforeCloseEvent(id interface{}, fn func(b *Bus)) {
	b.BeforeCloseEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b)
		},
	)
}
func (b *Bus) RaiseInitEvent() {
	b.InitEvent.Raise(nil)
}
func (b *Bus) BindInitEvent(id interface{}, fn func(b *Bus)) {
	b.InitEvent.BindAs(
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
