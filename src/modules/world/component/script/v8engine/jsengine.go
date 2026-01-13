package v8engine

import (
	"errors"
	"fmt"
	"modules/world"
	"modules/world/bus"
	"strconv"
	"sync"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/util"
	"github.com/herb-go/v8go"
	"github.com/herb-go/v8local"

	"github.com/herb-go/v8local/v8plugin"
)

func newJsInitializer(b *bus.Bus) *v8plugin.Initializer {
	i := v8plugin.NewInitializer()
	i.Entry = "main.js"
	i.DisableBuiltin = true
	i.Modules = []*herbplugin.Module{
		NewMapperModule(b),
		NewAPIModule(b),
		ModuleEval,
		NewHTTPModule(b),
		NewMetronomeModule(b),
		NewUserinputModule(b),
		NewBinaryModule(b),
	}
	return i
}

type JsEngine struct {
	Locker       sync.RWMutex
	Plugin       *v8plugin.Plugin
	onOpen       string
	onClose      string
	onDisconnect string
	onConnect    string
	onBroadCast  string
	onHUDClick   string
	onResponse   string
	onBuffer     string
	onSubneg     string
	onBufferMin  int
	onBufferMax  int
	onFocus      string
	onLoseFocus  string
	onKeyUp      string
}

func NewJsEngine() *JsEngine {
	return &JsEngine{
		Plugin: v8plugin.New(),
	}
}

type openScript struct {
	b *bus.Bus
	e *JsEngine
}

func (s *openScript) init() {
	newJsInitializer(s.b).MustApplyInitializer(s.e.Plugin)
}
func (s *openScript) lanuch() {
	opt := s.b.GetPluginOptions()
	herbplugin.Lanuch(s.e.Plugin, opt)
}
func (s *openScript) onOpen() {
	global := s.e.Plugin.Top.Global()
	fn := global.Get(s.e.onOpen)
	if !fn.IsNullOrUndefined() {
		fn.Call(s.e.Plugin.Top.NullValue())
	}
}
func (e *JsEngine) Open(b *bus.Bus) error {
	data := b.GetScriptData()
	e.onOpen = data.OnOpen
	e.onClose = data.OnClose
	e.onConnect = data.OnConnect
	e.onDisconnect = data.OnDisconnect
	e.onBroadCast = data.OnBroadcast
	e.onResponse = data.OnResponse
	e.onHUDClick = data.OnHUDClick
	e.onBuffer = data.OnBuffer
	e.onSubneg = data.OnSubneg
	e.onBufferMax = data.OnBufferMax
	e.onBufferMin = data.OnBufferMin
	e.onFocus = data.OnFocus
	e.onLoseFocus = data.OnLoseFocus
	e.onKeyUp = data.OnKeyUp
	s := &openScript{
		b: b,
		e: e,
	}
	err := formatError(util.Catch(s.init))
	if err != nil {
		return err
	}
	err = formatError(util.Catch(s.lanuch))
	if err != nil {
		return err
	}
	if data.OnOpen != "" {
		b.HandleScriptError(formatError(util.Catch(s.onOpen)))
	}
	return nil
}
func (e *JsEngine) Close(b *bus.Bus) {
	e.Locker.Lock()
	if e.onClose != "" {
		e.Locker.Unlock()
		local := e.Plugin.Runtime.NewLocal()
		e.Call(b, local, e.onClose)
		local.Close()
	} else {
		e.Locker.Unlock()
	}
	b.HandleScriptError(formatError(util.Catch(e.Plugin.MustClosePlugin)))
}
func (e *JsEngine) OnConnect(b *bus.Bus) {
	if e.onConnect != "" {
		local := e.Plugin.Runtime.NewLocal()
		defer local.Close()

		e.Call(b, local, e.onConnect)
	}
}
func (e *JsEngine) OnDisconnect(b *bus.Bus) {
	if e.onDisconnect != "" {
		local := e.Plugin.Runtime.NewLocal()
		defer local.Close()
		e.Call(b, local, e.onDisconnect)
	}
}

func (e *JsEngine) OnTrigger(b *bus.Bus, line *world.Line, trigger *world.Trigger, result *world.MatchResult) {
	if trigger.Script == "" {
		return
	}
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()

	model := local.NewObject()
	for k, v := range result.Named {
		model.Set(k, local.NewString(v))
	}
	for k, v := range result.List {
		if k == 0 {
			model.Set("10", local.NewString(v))
			continue
		} else if k > 9 {
			break
		}
		model.Set(strconv.Itoa(k-1), local.NewString(v))
	}

	e.Locker.Unlock()
	e.Call(b, local, trigger.Script, local.NewString(trigger.Name), local.NewString(line.Plain()), model)
}
func (e *JsEngine) OnAlias(b *bus.Bus, message string, alias *world.Alias, result *world.MatchResult) {
	if alias.Script == "" {
		return
	}
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()

	model := local.NewObject()
	for k, v := range result.Named {
		model.Set(k, local.NewString(v))
	}
	for k, v := range result.List {
		if k == 0 {
			model.Set("10", local.NewString(v))
			continue
		} else if k > 9 {
			break
		}
		model.Set(strconv.Itoa(k-1), local.NewString(v))
	}
	e.Locker.Unlock()
	e.Call(b, local, alias.Script, local.NewString(alias.Name), local.NewString(message), model)

}
func (e *JsEngine) OnTimer(b *bus.Bus, timer *world.Timer) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, timer.Script, local.NewString(timer.Name))
}
func (e *JsEngine) OnBroadCast(b *bus.Bus, bc *world.Broadcast) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, e.onBroadCast, local.NewString(bc.Message), local.NewBoolean(bc.Global), local.NewString(bc.Channel), local.NewString(bc.ID))
}
func (e *JsEngine) OnHUDClick(b *bus.Bus, c *world.Click) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	var local = e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, e.onHUDClick, local.NewNumber(c.X), local.NewNumber(c.Y))

}

func (e *JsEngine) OnResponse(b *bus.Bus, msg *world.Message) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, e.onResponse, local.NewString(msg.Type), local.NewString(msg.ID), local.NewString(msg.Data))
}

func (e *JsEngine) OnBuffer(b *bus.Bus, data []byte) bool {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onBuffer == "" {
		e.Locker.Unlock()
		return false
	}
	l := len(data)
	if data != nil && (l < e.onBufferMin || (e.onBufferMax > 0 && l > e.onBufferMax)) {
		e.Locker.Unlock()
		return false
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()

	var result bool
	if data != nil {
		result = e.Call(b, local, e.onBuffer, local.NewString(string(data)), local.NewArrayBuffer(data))
	} else {
		result = e.Call(b, local, e.onBuffer, local.NullValue(), local.NullValue())
	}
	return result
}
func (e *JsEngine) OnSubneg(b *bus.Bus, code byte, data []byte) bool {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onSubneg == "" {
		e.Locker.Unlock()
		return false
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	result := e.Call(b, local, e.onSubneg, local.NewInt32(int32(code)), local.NewString(string(data)))
	return result
}
func (e *JsEngine) OnFocus(b *bus.Bus) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onFocus == "" {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, e.onFocus)
}
func (e *JsEngine) OnLoseFocus(b *bus.Bus) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onLoseFocus == "" {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, e.onLoseFocus)
}
func (e *JsEngine) OnKeyUp(b *bus.Bus, key string) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onKeyUp == "" {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, e.onKeyUp, local.NewString(key))
}

func (e *JsEngine) OnCallback(b *bus.Bus, cb *world.Callback) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, cb.Script, local.NewString(cb.Name), local.NewString(cb.ID), local.NewInt32(int32(cb.Code)), local.NewString(cb.Data))

}
func (e *JsEngine) OnAssist(b *bus.Bus, script string) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()
	e.Call(b, local, script)
}

type jsRun struct {
	local *v8local.Local
	cmd   string
}

func (r *jsRun) Run() {
	r.local.RunScript(r.cmd, "run")
}
func (e *JsEngine) Run(b *bus.Bus, cmd string) {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	local := e.Plugin.Runtime.NewLocal()
	defer local.Close()

	r := &jsRun{
		local: local,
		cmd:   cmd,
	}
	b.HandleScriptError(formatError(util.Catch(r.Run)))
}

type runScript struct {
	bus    *bus.Bus
	output bool
	source string
	local  *v8local.Local
	args   []*v8local.JsValue
}

func (s *runScript) Exec() {
	fn := s.local.RunScript(s.source, "")
	if fn == nil || fn.IsNull() || !fn.IsFunction() {
		s.bus.HandleScriptError(errors.New(fmt.Sprintf("js function %s not found", s.source)))
		s.output = false
		return
	}
	var result *v8local.JsValue
	result = fn.Call(s.local.NullValue(), s.args...)
	if result.IsBoolean() {
		s.output = result.Boolean()
	} else {
		s.output = false
	}
	s.local.Context().RunIdleTasks(false, 0.005)
}

type ScriptError struct {
	*v8go.JSError
}

func (s *ScriptError) Error() string {
	return s.StackTrace
}

func formatError(err error) error {
	if err != nil {
		jserr, ok := err.(*v8go.JSError)
		if ok {
			return &ScriptError{JSError: jserr}
		}
	}
	return err
}
func (e *JsEngine) Call(b *bus.Bus, local *v8local.Local, source string, args ...*v8local.JsValue) bool {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	r := e.Plugin.Runtime
	if r == nil {
		return false
	}
	if e.Plugin.Runtime.Raw == nil {
		return false
	}
	if source == "" {
		return false
	}
	var output bool
	s := &runScript{
		bus:    b,
		source: source,
		local:  local,
		args:   args,
	}
	b.HandleScriptError(formatError(util.Catch(s.Exec)))
	return output
}

func init() {
	v8go.SetFlags("--expose-gc")
}
