package bus

import (
	"modules/mapper"
	"modules/world"
	"time"

	"github.com/herb-go/herbplugin"

	"github.com/herb-go/misc/busevent"
)

type Bus struct {
	ID                        string
	GetMaxHistory             func() int
	GetMaxLines               func() int
	GetMaxRecent              func() int
	GetConnBuffer             func() []byte
	GetConnConnected          func() bool
	GetHost                   func() string
	SetHost                   func(string)
	GetPort                   func() string
	SetPort                   func(string)
	GetProxy                  func() string
	SetProxy                  func(string)
	GetName                   func() string
	SetName                   func(string)
	GetShowBroadcast          func() bool
	SetShowBroadcast          func(bool)
	GetShowSubneg             func() bool
	SetShowSubneg             func(bool)
	GetModEnabled             func() bool
	SetModEnabled             func(bool)
	GetAutoSave               func() bool
	SetAutoSave               func(bool)
	GetIgnoreBatchCommand     func() bool
	SetIgnoreBatchCommand     func(bool)
	GetCommandStackCharacter  func() string
	SetCommandStackCharacter  func(string)
	GetScriptPrefix           func() string
	SetScriptPrefix           func(string)
	GetStatus                 func() string
	SetStatus                 func(string)
	GetQueueDelay             func() int
	SetQueueDelay             func(int)
	GetQueue                  func() []*world.Command
	GetParam                  func(string) string
	GetParams                 func() map[string]string
	SetParam                  func(string, string)
	DeleteParam               func(string)
	GetParamComment           func(string) string
	GetParamComments          func() map[string]string
	SetParamComment           func(string, string)
	GetCharset                func() string
	SetCharset                func(string)
	GetReadyAt                func() int64
	GetPosition               func() int
	SetPosition               func(int)
	GetCurrentLines           func() []*world.Line
	GetPrompt                 func() *world.Line
	GetClientInfo             func() *world.ClientInfo
	GetWorldData              func() *world.WorldData
	GetScriptData             func() *world.ScriptData
	SetPermissions            func([]string)
	GetPermissions            func() []string
	RequestPermissions        func(*world.Authorization)
	GetScriptID               func() string
	SetScriptID               func(string)
	GetScriptType             func() string
	GetScriptPath             func() string
	GetModPath                func() string
	GetScriptModPath          func() string
	GetLogsPath               func() string
	GetScriptHome             func() string
	DoLog                     func(string)
	SetTrusted                func(*herbplugin.Trusted)
	GetTrusted                func() *herbplugin.Trusted
	RequestTrustDomains       func(*world.Authorization)
	GetPluginOptions          func() herbplugin.Options
	DoReloadPermissions       func()
	DoSendToConn              func(cmd []byte)
	DoSend                    func(*world.Command)
	DoSendToQueue             func(*world.Command)
	DoExecute                 func(message string)
	DoEncode                  func() ([]byte, error)
	DoDecode                  func([]byte) error
	DoReloadScript            func() error
	DoSaveScript              func() error
	DoUseScript               func(string)
	GetRequiredParams         func() []*world.RequiredParam
	DoRunScript               func(string)
	DoPrint                   func(msg string)
	DoPrintSystem             func(msg string)
	DoPrintLocalBroadcastIn   func(msg string)
	DoPrintGlobalBroadcastIn  func(msg string)
	DoPrintLocalBroadcastOut  func(msg string)
	DoPrintGlobalBroadcastOut func(msg string)
	DoPrintSubneg             func(msg string)
	DoPrintRequest            func(msg string)
	DoPrintResponse           func(msg string)

	DoDiscardQueue          func(force bool) int
	DoLockQueue             func()
	DoOmitOutput            func()
	DoDeleteLines           func(int)
	GetLineCount            func() int
	DoSendHUDClickToScript  func(*world.Click)
	DoSendBroadcastToScript func(*world.Broadcast)
	HandleBuffer            func([]byte) bool
	HandleSubneg            func([]byte) bool
	HandleFocus             func()
	HandleLoseFocus         func()
	DoSendTimerToScript     func(*world.Timer)
	DoDeleteTimer           func(string) bool
	DoDeleteTimerByName     func(string) bool
	DoDeleteTemporaryTimers func() int
	DoDeleteTimerGroup      func(string, bool) int
	DoEnableTimerByName     func(string, bool) bool
	DoEnableTimerGroup      func(string, bool) int
	DoResetNamedTimer       func(string) bool
	GetTimer                func(string) *world.Timer
	GetTimersByType         func(bool) []*world.Timer
	DoDeleteTimerByType     func(bool)
	AddTimers               func(ts []*world.Timer)
	DoResetTimers           func()
	GetTimerOption          func(name string, option string) (string, bool, bool)
	GetTimerInfo            func(name string, infotype int) (string, bool, bool)
	SetTimerOption          func(name string, option string, value string) (bool, bool, bool)
	HasNamedTimer           func(string) bool
	DoListTimerNames        func(byUser bool) []string
	AddTimer                func(*world.Timer, bool) bool
	DoUpdateTimer           func(*world.Timer) int
	DoSendAliasToScript     func(message string, alias *world.Alias, result *world.MatchResult)

	DoDeleteAlias            func(string) bool
	DoDeleteAliasByName      func(string) bool
	DoDeleteTemporaryAliases func() int
	DoDeleteAliasGroup       func(string, bool) int
	DoEnableAliasByName      func(string, bool) bool
	DoEnableAliasGroup       func(string, bool) int
	GetAlias                 func(string) *world.Alias
	GetAliasesByType         func(bool) []*world.Alias
	DoDeleteAliasByType      func(bool)
	AddAliases               func([]*world.Alias)
	GetAliasOption           func(name string, option string) (string, bool, bool)
	GetAliasInfo             func(name string, infotype int) (string, bool, bool)
	SetAliasOption           func(name string, option string, value string) (bool, bool, bool)
	HasNamedAlias            func(string) bool
	DoListAliasNames         func(bool) []string
	AddAlias                 func(*world.Alias, bool) bool
	DoUpdateAlias            func(*world.Alias) int

	DoDeleteTrigger           func(string) bool
	DoDeleteTriggerByName     func(string) bool
	DoDeleteTemporaryTriggers func() int
	DoDeleteTriggerGroup      func(string, bool) int
	DoEnableTriggerByName     func(string, bool) bool
	DoEnableTriggerGroup      func(string, bool) int
	GetTrigger                func(string) *world.Trigger
	GetTriggersByType         func(bool) []*world.Trigger
	DoDeleteTriggerByType     func(bool)
	AddTriggers               func([]*world.Trigger)
	GetTriggerOption          func(name string, option string) (string, bool, bool)
	GetTriggerInfo            func(name string, infotype int) (string, bool, bool)
	SetTriggerOption          func(name string, option string, value string) (bool, bool, bool)
	HasNamedTrigger           func(string) bool
	DoListTriggerNames        func(byUser bool) []string
	AddTrigger                func(*world.Trigger, bool) bool
	DoUpdateTrigger           func(*world.Trigger) int
	DoSendTriggerToScript     func(line *world.Line, trigger *world.Trigger, result *world.MatchResult)
	DoGetTriggerWildcard      func(name string) *world.MatchResult
	DoSendCallbackToScript    func(cb *world.Callback)
	DoSendKeyUpToScript       func(string)
	DoAssist                  func()
	DoMultiLinesAppend        func(string)
	DoMultiLinesFlush         func()
	DoMultiLinesLast          func(int) []string
	GetLinesInBufferCount     func() int
	GetRecentLines            func(count int) []*world.Line
	GetLine                   func(idx int) *world.Line
	GetMapper                 func() *mapper.Mapper
	GetPriority               func() int
	SetPriority               func(int)
	GetSummary                func() []*world.Line
	SetSummary                func([]*world.Line)
	UpdateLastActive          func()
	GetLastActive             func() int64
	AddHistory                func(string)
	GetHistories              func() []string
	FlushHistories            func()
	HandleConnReceive         func(msg []byte)
	HandleConnError           func(err error)
	HandleConnPrompt          func(msg []byte)
	DoConnectServer           func() error
	DoCloseServer             func() error
	HandleConverterError      func(err error)
	HandleCmdError            func(err error)
	HandleTriggerError        func(err error)
	HandleScriptError         func(err error)
	GetScriptCaller           func() (string, string)
	DoStopEvaluatingTriggers  func()

	GetMetronomeBeats      func() int
	SetMetronomeBeats      func(int)
	DoResetMetronome       func()
	GetMetronomeSpace      func() int
	GetMetronomeQueue      func() []string
	DoDiscardMetronome     func(force bool) bool
	DoLockMetronomeQueue   func()
	DoFullMetronome        func()
	DoFullTickMetronome    func()
	SetMetronomeInterval   func(time.Duration)
	GetMetronomeInterval   func() time.Duration
	SetMetronomeTick       func(time.Duration)
	GetMetronomeTick       func() time.Duration
	DoPushMetronome        func(cmds []*world.Command, grouped bool)
	DoMetronomeSend        func(cmds *world.Command)
	DoMetronomeLock        func()
	DoSendResponseToScript func(msg *world.Message)

	GetHUDSize       func() int
	SetHUDSize       func(int)
	GetHUDContent    func() []*world.Line
	UpdateHUDContent func(start int, content []*world.Line) bool

	RequestEvent           busevent.Event
	BroadcastEvent         busevent.Event
	LineEvent              busevent.Event
	PromptEvent            busevent.Event
	ConnectedEvent         busevent.Event
	DisconnectedEvent      busevent.Event
	ServerCloseEvent       busevent.Event
	InitEvent              busevent.Event
	ReadyEvent             busevent.Event
	BeforeCloseEvent       busevent.Event
	CloseEvent             busevent.Event
	HistoriesEvent         busevent.Event
	StatusEvent            busevent.Event
	LinesEvent             busevent.Event
	QueueDelayUpdatedEvent busevent.Event
	ScriptMessageEvent     busevent.Event
	HUDUpdateEvent         busevent.Event
	HUDContentEvent        busevent.Event
	ClientInfoEvent        busevent.Event
	SaveEvent              busevent.Event
}

func (b *Bus) Dispose() {
	b.GetMaxHistory = nil
	b.GetMaxLines = nil
	b.GetMaxRecent = nil
	b.GetConnBuffer = nil
	b.GetConnConnected = nil
	b.GetHost = nil
	b.SetHost = nil
	b.GetPort = nil
	b.SetPort = nil
	b.GetProxy = nil
	b.SetProxy = nil
	b.GetName = nil
	b.SetName = nil
	b.GetShowBroadcast = nil
	b.SetShowBroadcast = nil
	b.GetShowSubneg = nil
	b.SetShowSubneg = nil
	b.GetModEnabled = nil
	b.SetModEnabled = nil
	b.GetAutoSave = nil
	b.SetAutoSave = nil
	b.GetIgnoreBatchCommand = nil
	b.SetIgnoreBatchCommand = nil
	b.GetCommandStackCharacter = nil
	b.SetCommandStackCharacter = nil
	b.GetScriptPrefix = nil
	b.SetScriptPrefix = nil
	b.GetStatus = nil
	b.SetStatus = nil
	b.GetQueueDelay = nil
	b.SetQueueDelay = nil
	b.GetQueue = nil
	b.GetParam = nil
	b.GetParams = nil
	b.SetParam = nil
	b.DeleteParam = nil
	b.GetParamComment = nil
	b.GetParamComments = nil
	b.SetParamComment = nil
	b.GetCharset = nil
	b.SetCharset = nil
	b.GetReadyAt = nil
	b.GetPosition = nil
	b.SetPosition = nil
	b.GetCurrentLines = nil
	b.GetPrompt = nil
	b.GetClientInfo = nil
	b.GetWorldData = nil
	b.GetScriptData = nil
	b.SetPermissions = nil
	b.GetPermissions = nil
	b.RequestPermissions = nil
	b.GetScriptID = nil
	b.SetScriptID = nil
	b.GetScriptType = nil
	b.GetScriptPath = nil
	b.GetModPath = nil
	b.GetScriptModPath = nil
	b.GetLogsPath = nil
	b.GetScriptHome = nil
	b.DoLog = nil
	b.SetTrusted = nil
	b.GetTrusted = nil
	b.RequestTrustDomains = nil
	b.GetPluginOptions = nil
	b.DoReloadPermissions = nil
	b.DoSendToConn = nil
	b.DoSend = nil
	b.DoSendToQueue = nil
	b.DoExecute = nil
	b.DoEncode = nil
	b.DoDecode = nil
	b.DoReloadScript = nil
	b.DoSaveScript = nil
	b.DoUseScript = nil
	b.GetRequiredParams = nil
	b.DoRunScript = nil
	b.DoPrint = nil
	b.DoPrintSystem = nil
	b.DoPrintLocalBroadcastIn = nil
	b.DoPrintGlobalBroadcastIn = nil
	b.DoPrintLocalBroadcastOut = nil
	b.DoPrintGlobalBroadcastOut = nil
	b.DoPrintSubneg = nil
	b.DoPrintRequest = nil
	b.DoPrintResponse = nil
	b.DoDiscardQueue = nil
	b.DoLockQueue = nil
	b.DoOmitOutput = nil
	b.DoDeleteLines = nil
	b.GetLineCount = nil
	b.DoSendHUDClickToScript = nil
	b.DoSendBroadcastToScript = nil
	b.HandleBuffer = nil
	b.HandleSubneg = nil
	b.HandleFocus = nil
	b.HandleLoseFocus = nil
	b.DoSendTimerToScript = nil
	b.DoDeleteTimer = nil
	b.DoDeleteTimerByName = nil
	b.DoDeleteTemporaryTimers = nil
	b.DoDeleteTimerGroup = nil
	b.DoEnableTimerByName = nil
	b.DoEnableTimerGroup = nil
	b.DoResetNamedTimer = nil
	b.GetTimer = nil
	b.GetTimersByType = nil
	b.DoDeleteTimerByType = nil
	b.AddTimers = nil
	b.DoResetTimers = nil
	b.GetTimerOption = nil
	b.GetTimerInfo = nil
	b.SetTimerOption = nil
	b.HasNamedTimer = nil
	b.DoListTimerNames = nil
	b.AddTimer = nil
	b.DoUpdateTimer = nil
	b.DoSendAliasToScript = nil
	b.DoDeleteAlias = nil
	b.DoDeleteAliasByName = nil
	b.DoDeleteTemporaryAliases = nil
	b.DoDeleteAliasGroup = nil
	b.DoEnableAliasByName = nil
	b.DoEnableAliasGroup = nil
	b.GetAlias = nil
	b.GetAliasesByType = nil
	b.DoDeleteAliasByType = nil
	b.AddAliases = nil
	b.GetAliasOption = nil
	b.GetAliasInfo = nil
	b.SetAliasOption = nil
	b.HasNamedAlias = nil
	b.DoListAliasNames = nil
	b.AddAlias = nil
	b.DoUpdateAlias = nil
	b.DoDeleteTrigger = nil
	b.DoDeleteTriggerByName = nil
	b.DoDeleteTemporaryTriggers = nil
	b.DoDeleteTriggerGroup = nil
	b.DoEnableTriggerByName = nil
	b.DoEnableTriggerGroup = nil
	b.GetTrigger = nil
	b.GetTriggersByType = nil
	b.DoDeleteTriggerByType = nil
	b.AddTriggers = nil
	b.GetTriggerOption = nil
	b.GetTriggerInfo = nil
	b.SetTriggerOption = nil
	b.HasNamedTrigger = nil
	b.DoListTriggerNames = nil
	b.AddTrigger = nil
	b.DoUpdateTrigger = nil
	b.DoSendTriggerToScript = nil
	b.DoGetTriggerWildcard = nil
	b.DoSendCallbackToScript = nil
	b.DoSendKeyUpToScript = nil
	b.DoAssist = nil
	b.DoMultiLinesAppend = nil
	b.DoMultiLinesFlush = nil
	b.DoMultiLinesLast = nil
	b.GetLinesInBufferCount = nil
	b.GetRecentLines = nil
	b.GetLine = nil
	b.GetMapper = nil
	b.GetPriority = nil
	b.SetPriority = nil
	b.GetSummary = nil
	b.SetSummary = nil
	b.UpdateLastActive = nil
	b.GetLastActive = nil
	b.AddHistory = nil
	b.GetHistories = nil
	b.FlushHistories = nil
	b.HandleConnReceive = nil
	b.HandleConnError = nil
	b.HandleConnPrompt = nil
	b.DoConnectServer = nil
	b.DoCloseServer = nil
	b.HandleConverterError = nil
	b.HandleCmdError = nil
	b.HandleTriggerError = nil
	b.HandleScriptError = nil
	b.GetScriptCaller = nil
	b.DoStopEvaluatingTriggers = nil
	b.GetMetronomeBeats = nil
	b.SetMetronomeBeats = nil
	b.DoResetMetronome = nil
	b.GetMetronomeSpace = nil
	b.GetMetronomeQueue = nil
	b.DoDiscardMetronome = nil
	b.DoLockMetronomeQueue = nil
	b.DoFullMetronome = nil
	b.DoFullTickMetronome = nil
	b.SetMetronomeInterval = nil
	b.GetMetronomeInterval = nil
	b.SetMetronomeTick = nil
	b.GetMetronomeTick = nil
	b.DoPushMetronome = nil
	b.DoMetronomeSend = nil
	b.DoMetronomeLock = nil
	b.DoSendResponseToScript = nil
	b.GetHUDSize = nil
	b.SetHUDSize = nil
	b.GetHUDContent = nil
	b.UpdateHUDContent = nil
	b.RequestEvent.Flush()
	b.BroadcastEvent.Flush()
	b.LineEvent.Flush()
	b.PromptEvent.Flush()
	b.ConnectedEvent.Flush()
	b.DisconnectedEvent.Flush()
	b.ServerCloseEvent.Flush()
	b.InitEvent.Flush()
	b.ReadyEvent.Flush()
	b.BeforeCloseEvent.Flush()
	b.CloseEvent.Flush()
	b.HistoriesEvent.Flush()
	b.StatusEvent.Flush()
	b.LinesEvent.Flush()
	b.QueueDelayUpdatedEvent.Flush()
	b.ScriptMessageEvent.Flush()
	b.HUDUpdateEvent.Flush()
	b.HUDContentEvent.Flush()
	b.ClientInfoEvent.Flush()
	b.SaveEvent.Flush()
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
func (b *Bus) WrapDiscard(f func(bus *Bus, force bool) int) func(bool) int {
	return func(force bool) int {
		return f(b, force)
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
func (b *Bus) WrapSetLines(f func(bus *Bus, lines []*world.Line)) func([]*world.Line) {
	return func(lines []*world.Line) {
		f(b, lines)
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

func (b *Bus) WrapSetInt(f func(bus *Bus, value int)) func(int) {
	return func(value int) {
		f(b, value)
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
func (b *Bus) WrapGetPluginOptions(f func(bus *Bus) herbplugin.Options) func() herbplugin.Options {
	return func() herbplugin.Options {
		return f(b)
	}
}
func (b *Bus) WrapSetDuration(f func(bus *Bus, d time.Duration)) func(time.Duration) {
	return func(d time.Duration) {
		f(b, d)
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
func (b *Bus) WrapHandleBroadcast(f func(b *Bus, bc *world.Broadcast)) func(bc *world.Broadcast) {
	return func(bc *world.Broadcast) {
		f(b, bc)
	}
}
func (b *Bus) WrapHandleClick(f func(b *Bus, c *world.Click)) func(c *world.Click) {
	return func(c *world.Click) {
		f(b, c)
	}
}
func (b *Bus) WrapHandleResponse(f func(b *Bus, bc *world.Message)) func(msg *world.Message) {
	return func(msg *world.Message) {
		f(b, msg)
	}
}
func (b *Bus) WrapHandleCallback(f func(b *Bus, bc *world.Callback)) func(bc *world.Callback) {
	return func(cb *world.Callback) {
		f(b, cb)
	}
}
func (b *Bus) WrapHandleStringForBool(f func(bus *Bus, str string) bool) func(str string) bool {
	return func(str string) bool {
		return f(b, str)
	}
}
func (b *Bus) WrapHandleBytesForBool(f func(bus *Bus, data []byte) bool) func(data []byte) bool {
	return func(data []byte) bool {
		return f(b, data)
	}
}
func (b *Bus) WrapHandlePushGroupedCommands(f func(b *Bus, cmds []*world.Command, grouped bool)) func(cmds []*world.Command, grouped bool) {
	return func(cmds []*world.Command, grouped bool) {
		f(b, cmds, grouped)
	}
}
func (b *Bus) WrapHandleAuthorization(f func(b *Bus, a *world.Authorization)) func(a *world.Authorization) {
	return func(a *world.Authorization) {
		f(b, a)
	}
}

type lineBinder struct {
	fn  func(b *Bus, line *world.Line)
	bus *Bus
}

func newLineBinder(bus *Bus, fn func(b *Bus, line *world.Line)) *lineBinder {
	return &lineBinder{
		bus: bus,
		fn:  fn,
	}
}
func (b *lineBinder) Exec(data interface{}) {
	b.fn(b.bus, data.(*world.Line))
}
func (b *Bus) RaiseLineEvent(line *world.Line) {
	b.LineEvent.Raise(line)
}
func (b *Bus) BindLineEvent(id interface{}, fn func(b *Bus, line *world.Line)) {
	b.LineEvent.BindAs(
		id,
		newLineBinder(b, fn).Exec,
	)

}

func (b *Bus) RaisePromptEvent(line *world.Line) {
	b.PromptEvent.Raise(line)
}
func (b *Bus) BindPromptEvent(id interface{}, fn func(b *Bus, line *world.Line)) {
	b.PromptEvent.BindAs(
		id,
		newLineBinder(b, fn).Exec,
	)
}

type binder struct {
	fn  func(b *Bus)
	bus *Bus
}

func newBinder(bus *Bus, fn func(b *Bus)) *binder {
	return &binder{
		bus: bus,
		fn:  fn,
	}
}
func (b *binder) Exec(data interface{}) {
	b.fn(b.bus)
}
func (b *Bus) RaiseConnectedEvent() {
	b.ConnectedEvent.Raise(nil)
}
func (b *Bus) BindConnectedEvent(id interface{}, fn func(b *Bus)) {
	b.ConnectedEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}

func (b *Bus) RaiseDisconnectedEvent() {
	b.DisconnectedEvent.Raise(nil)
}
func (b *Bus) BindDisconnectedEvent(id interface{}, fn func(b *Bus)) {
	b.DisconnectedEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}
func (b *Bus) RaiseCloseEvent() {
	b.CloseEvent.Raise(nil)
}
func (b *Bus) BindCloseEvent(id interface{}, fn func(b *Bus)) {
	b.CloseEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}
func (b *Bus) RaiseServerCloseEvent() {
	b.ServerCloseEvent.Raise(nil)
}
func (b *Bus) BindServerCloseEvent(id interface{}, fn func(b *Bus)) {
	b.ServerCloseEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}

func (b *Bus) RaiseBeforeCloseEvent() {
	b.BeforeCloseEvent.Raise(nil)
}
func (b *Bus) BindBeforeCloseEvent(id interface{}, fn func(b *Bus)) {
	b.BeforeCloseEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}
func (b *Bus) RaiseInitEvent() {
	b.InitEvent.Raise(nil)
}
func (b *Bus) BindInitEvent(id interface{}, fn func(b *Bus)) {
	b.InitEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}

func (b *Bus) RaiseReadyEvent() {
	b.ReadyEvent.Raise(nil)
}
func (b *Bus) BindReadyEvent(id interface{}, fn func(b *Bus)) {
	b.ReadyEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
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
func (b *Bus) RaiseLinesEvent(lines []*world.Line) {
	b.LinesEvent.Raise(lines)
}
func (b *Bus) BindLinesEvent(id interface{}, fn func(b *Bus, lines []*world.Line)) {
	b.LinesEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.([]*world.Line))
		},
	)
}
func (b *Bus) RaiseRequestEvent(msg *world.Message) {
	b.RequestEvent.Raise(msg)
}
func (b *Bus) BindRequestEvent(id interface{}, fn func(b *Bus, msg *world.Message)) {
	b.RequestEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*world.Message))
		},
	)
}
func (b *Bus) RaiseBroadcastEvent(bc *world.Broadcast) {
	b.BroadcastEvent.Raise(bc)
}
func (b *Bus) BindBroadcastEvent(id interface{}, fn func(b *Bus, bc *world.Broadcast)) {
	b.BroadcastEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*world.Broadcast))
		},
	)
}
func (b *Bus) RaiseQueueDelayUpdatedEvent() {
	b.QueueDelayUpdatedEvent.Raise(nil)
}
func (b *Bus) BindQueueDelayUpdatedEvent(id interface{}, fn func(b *Bus)) {
	b.QueueDelayUpdatedEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}

func (b *Bus) RaiseScriptMessageEvent(msg interface{}) {
	b.ScriptMessageEvent.Raise(msg)
}
func (b *Bus) BindScriptMessageEvent(id interface{}, fn func(b *Bus, msg interface{})) {
	b.ScriptMessageEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data)
		},
	)
}
func (b *Bus) RaiseHUDUpdateEvent(diff *world.DiffLines) {
	b.HUDUpdateEvent.Raise(diff)
}
func (b *Bus) BindHUDUpdateEvent(id interface{}, fn func(b *Bus, diff *world.DiffLines)) {
	b.HUDUpdateEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*world.DiffLines))
		},
	)
}
func (b *Bus) RaiseHUDContentEvent(lines []*world.Line) {
	b.HUDContentEvent.Raise(lines)
}
func (b *Bus) BindHUDContentEvent(id interface{}, fn func(b *Bus, lines []*world.Line)) {
	b.HUDContentEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.([]*world.Line))
		},
	)
}
func (b *Bus) RaiseClientInfoEvent(info *world.ClientInfo) {
	b.ClientInfoEvent.Raise(info)
}
func (b *Bus) BindClientInfoEvent(id interface{}, fn func(b *Bus, info *world.ClientInfo)) {
	b.ClientInfoEvent.BindAs(
		id,
		func(data interface{}) {
			fn(b, data.(*world.ClientInfo))
		},
	)
}
func (b *Bus) RaiseSaveEvent() {
	b.SaveEvent.Raise(nil)
}
func (b *Bus) BindSaveEvent(id interface{}, fn func(b *Bus)) {
	b.SaveEvent.BindAs(
		id,
		newBinder(b, fn).Exec,
	)
}

func (b *Bus) Reset() {
	*b = Bus{}
}
func New() *Bus {
	return &Bus{}
}
