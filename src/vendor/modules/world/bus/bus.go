package bus

import (
	"modules/mapper"
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
	GetStatus              func() string
	SetStatus              func(string)
	GetQueueDelay          func() int
	SetQueueDelay          func(int)
	GetQueue               func() []*world.Command
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
	DoSend                 func(*world.Command)
	DoSendToQueue          func(*world.Command)
	DoExecute              func(message string)
	DoEncode               func() ([]byte, error)
	DoDecode               func([]byte) error
	DoReloadScript         func() error
	DoSaveScript           func() error
	DoUseScript            func(string)
	DoRunScript            func(string)
	DoPrint                func(msg string)
	DoPrintSystem          func(msg string)
	DoDiscardQueue         func() int

	DoSendTimerToScript     func(*world.Timer)
	DoDeleteTimer           func(string) bool
	DoDeleteTimerByName     func(string) bool
	DoDeleteTemporaryTimers func() int
	DoDeleteTimerGroup      func(string) int
	DoEnableTimerByName     func(string, bool) bool
	DoEnableTimerGroup      func(string, bool) int
	DoResetNamedTimer       func(string) bool
	GetTimer                func(string) *world.Timer
	GetTimersByType         func(bool) []*world.Timer
	DoDeleteTimerByType     func(bool)
	AddTimers               func(ts []*world.Timer)
	DoResetTimers           func()
	GetTimerOption          func(name string, option string) (string, bool, bool)
	SetTimerOption          func(name string, option string, value string) (bool, bool, bool)
	HasNamedTimer           func(string) bool
	DoListTimerNames        func() []string
	AddTimer                func(*world.Timer, bool) bool
	DoUpdateTimer           func(*world.Timer) int
	DoSendAliasToScript     func(message string, alias *world.Alias, result *world.MatchResult)

	DoDeleteAlias            func(string) bool
	DoDeleteAliasByName      func(string) bool
	DoDeleteTemporaryAliases func() int
	DoDeleteAliasGroup       func(string) int
	DoEnableAliasByName      func(string, bool) bool
	DoEnableAliasGroup       func(string, bool) int
	GetAlias                 func(string) *world.Alias
	GetAliasesByType         func(bool) []*world.Alias
	DoDeleteAliasByType      func(bool)
	AddAliases               func([]*world.Alias)
	GetAliasOption           func(name string, option string) (string, bool, bool)
	SetAliasOption           func(name string, option string, value string) (bool, bool, bool)
	HasNamedAlias            func(string) bool
	DoListAliasNames         func() []string
	AddAlias                 func(*world.Alias, bool) bool
	DoUpdateAlias            func(*world.Alias) int

	DoDeleteTrigger           func(string) bool
	DoDeleteTriggerByName     func(string) bool
	DoDeleteTemporaryTriggers func() int
	DoDeleteTriggerGroup      func(string) int
	DoEnableTriggerByName     func(string, bool) bool
	DoEnableTriggerGroup      func(string, bool) int
	GetTrigger                func(string) *world.Trigger
	GetTriggersByType         func(bool) []*world.Trigger
	DoDeleteTriggerByType     func(bool)
	AddTriggers               func([]*world.Trigger)
	GetTriggerOption          func(name string, option string) (string, bool, bool)
	SetTriggerOption          func(name string, option string, value string) (bool, bool, bool)
	HasNamedTrigger           func(string) bool
	DoListTriggerNames        func() []string
	AddTrigger                func(*world.Trigger, bool) bool
	DoUpdateTrigger           func(*world.Trigger) int
	DoSendTriggerToScript     func(line *world.Line, trigger *world.Trigger, result *world.MatchResult)

	DoMultiLinesAppend func(string)
	DoMultiLinesFlush  func()
	DoMultiLinesLast   func(int) []string

	GetMapper func() *mapper.Mapper

	AddHistory           func(string)
	GetHistories         func() []string
	FlushHistories       func()
	HandleConnReceive    func(msg []byte)
	HandleConnError      func(err error)
	HandleConnPrompt     func(msg []byte)
	DoConnectServer      func() error
	DoCloseServer        func() error
	HandleConverterError func(err error)
	HandleCmdError       func(err error)
	HandleTriggerError   func(err error)
	HandleScriptError    func(err error)

	DoStopEvaluatingTriggers func()
	LineEvent                busevent.Event
	PromptEvent              busevent.Event
	ConnectedEvent           busevent.Event
	DisconnectedEvent        busevent.Event
	InitEvent                busevent.Event
	ReadyEvent               busevent.Event
	BeforeCloseEvent         busevent.Event
	CloseEvent               busevent.Event
	HistoriesEvent           busevent.Event
	StatusEvent              busevent.Event
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
func (b *Bus) WrapGetStrings(f func(bus *Bus) []string) func() []string {
	return func() []string {
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
func (b *Bus) WrapHandleSend(f func(bus *Bus, cmd *world.Command)) func(cmd *world.Command) {
	return func(cmd *world.Command) {
		f(b, cmd)
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
func (b *Bus) WrapDoEncode(f func(bus *Bus) ([]byte, error)) func() ([]byte, error) {
	return func() ([]byte, error) {
		return f(b)
	}
}
func (b *Bus) WrapDoDecode(f func(bus *Bus, bs []byte) error) func([]byte) error {
	return func(bs []byte) error {
		return f(b, bs)
	}
}
func (b *Bus) WrapAddTimer(f func(bus *Bus, timer *world.Timer, replace bool) bool) func(*world.Timer, bool) bool {
	return func(timer *world.Timer, replace bool) bool {
		return f(b, timer, replace)
	}
}
func (b *Bus) WrapHandleTimer(f func(bus *Bus, timer *world.Timer)) func(*world.Timer) {
	return func(timer *world.Timer) {
		f(b, timer)
	}
}
func (b *Bus) WrapHandleAlias(f func(b *Bus, message string, alias *world.Alias, result *world.MatchResult)) func(message string, alias *world.Alias, result *world.MatchResult) {
	return func(message string, alias *world.Alias, result *world.MatchResult) {
		f(b, message, alias, result)
	}
}
func (b *Bus) WrapHandleTrigger(f func(b *Bus, line *world.Line, trigger *world.Trigger, result *world.MatchResult)) func(line *world.Line, trigger *world.Trigger, result *world.MatchResult) {
	return func(line *world.Line, trigger *world.Trigger, result *world.MatchResult) {
		f(b, line, trigger, result)
	}
}

func (b *Bus) WrapHandleStringForBool(f func(bus *Bus, str string) bool) func(str string) bool {
	return func(str string) bool {
		return f(b, str)
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
func (b *Bus) RaiseHistoriesEvent(histories []string) {
	b.HistoriesEvent.Raise(histories)
}
func (b *Bus) BindHistoriesEvent(id interface{}, fn func(b *Bus, histories []string)) {
	b.HistoriesEvent.BindAs(
		id,
		func(data interface{}) {
			h, _ := data.([]string)
			fn(b, h)
		},
	)
}
func (b *Bus) RaiseStatusEvent(status string) {
	b.StatusEvent.Raise(status)
}
func (b *Bus) BindStatusEvent(id interface{}, fn func(b *Bus, status string)) {
	b.StatusEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(string))
		},
	)
}

func New() *Bus {
	return &Bus{}
}
