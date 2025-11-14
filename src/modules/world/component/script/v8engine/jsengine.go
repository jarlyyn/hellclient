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
	"github.com/jarlyyn/v8js"

	"github.com/jarlyyn/v8js/v8plugin"
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
	global := s.e.Plugin.Runtime.Global()
	fn := global.Get(s.e.onOpen)
	global.Release()
	if !fn.IsNullOrUndefined() {
		fn.Call(s.e.Plugin.Runtime.NullValue())
	}
	fn.Release()
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
	err := util.Catch(s.init)
	if err != nil {
		return err
	}
	err = util.Catch(s.lanuch)
	if err != nil {
		return err
	}
	if data.OnOpen != "" {
		b.HandleScriptError(util.Catch(s.onOpen))
	}
	return nil
}
func (e *JsEngine) Close(b *bus.Bus) {
	e.Locker.Lock()
	if e.onClose != "" {
		e.Locker.Unlock()
		e.Call(b, e.onClose)
	} else {
		e.Locker.Unlock()
	}
	b.HandleScriptError(util.Catch(e.Plugin.MustClosePlugin))
}
func (e *JsEngine) OnConnect(b *bus.Bus) {
	if e.onConnect != "" {
		e.Call(b, e.onConnect)
	}
}
func (e *JsEngine) OnDisconnect(b *bus.Bus) {
	if e.onDisconnect != "" {
		e.Call(b, e.onDisconnect)
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
	model := e.Plugin.Runtime.NewObject()
	for k, v := range result.Named {
		model.Set(k, e.Plugin.Runtime.NewString(v).Consume())
	}
	for k, v := range result.List {
		if k == 0 {
			model.Set("10", e.Plugin.Runtime.NewString(v).Consume())
			continue
		} else if k > 9 {
			break
		}
		model.Set(strconv.Itoa(k-1), e.Plugin.Runtime.NewString(v).Consume())
	}

	e.Locker.Unlock()
	e.Call(b, trigger.Script, e.Plugin.Runtime.NewString(trigger.Name).Consume(), e.Plugin.Runtime.NewString(line.Plain()).Consume(), model.Consume())

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
	model := e.Plugin.Runtime.NewObject()
	for k, v := range result.Named {
		model.Set(k, e.Plugin.Runtime.NewString(v).Consume())
	}
	for k, v := range result.List {
		if k == 0 {
			model.Set("10", e.Plugin.Runtime.NewString(v).Consume())
			continue
		} else if k > 9 {
			break
		}
		model.Set(strconv.Itoa(k-1), e.Plugin.Runtime.NewString(v).Consume())
	}
	e.Locker.Unlock()
	go e.Call(b, alias.Script, e.Plugin.Runtime.NewString(alias.Name).Consume(), e.Plugin.Runtime.NewString(message).Consume(), model.Consume())

}
func (e *JsEngine) OnTimer(b *bus.Bus, timer *world.Timer) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, timer.Script, e.Plugin.Runtime.NewString(timer.Name).Consume())
}
func (e *JsEngine) OnBroadCast(b *bus.Bus, bc *world.Broadcast) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, e.onBroadCast, e.Plugin.Runtime.NewString(bc.Message).Consume(), e.Plugin.Runtime.NewBoolean(bc.Global).Consume(), e.Plugin.Runtime.NewString(bc.Channel).Consume(), e.Plugin.Runtime.NewString(bc.ID).Consume())
}
func (e *JsEngine) OnHUDClick(b *bus.Bus, c *world.Click) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, e.onHUDClick, e.Plugin.Runtime.NewNumber(c.X).Consume(), e.Plugin.Runtime.NewNumber(c.Y).Consume())

}

func (e *JsEngine) OnResponse(b *bus.Bus, msg *world.Message) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, e.onResponse, e.Plugin.Runtime.NewString(msg.Type).Consume(), e.Plugin.Runtime.NewString(msg.ID).Consume(), e.Plugin.Runtime.NewString(msg.Data).Consume())
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
	var result bool
	if data != nil {
		result = e.Call(b, e.onBuffer, e.Plugin.Runtime.NewString(string(data)).Consume(), e.Plugin.Runtime.NewArrayBuffer(data).Consume())
	} else {
		result = e.Call(b, e.onBuffer, e.Plugin.Runtime.NullValue().Consume(), e.Plugin.Runtime.NullValue().Consume())
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
	result := e.Call(b, e.onSubneg, e.Plugin.Runtime.NewInt32(int32(code)).Consume(), e.Plugin.Runtime.NewString(string(data)).Consume())
	return result
}
func (e *JsEngine) OnFocus(b *bus.Bus) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onFocus == "" {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	e.Call(b, e.onFocus)
}
func (e *JsEngine) OnLoseFocus(b *bus.Bus) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onLoseFocus == "" {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	e.Call(b, e.onLoseFocus)
}
func (e *JsEngine) OnKeyUp(b *bus.Bus, key string) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil || e.onKeyUp == "" {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	e.Call(b, e.onKeyUp, e.Plugin.Runtime.NewString(key).Consume())
}

func (e *JsEngine) OnCallback(b *bus.Bus, cb *world.Callback) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, cb.Script, e.Plugin.Runtime.NewString(cb.Name).Consume(), e.Plugin.Runtime.NewString(cb.ID).Consume(), e.Plugin.Runtime.NewInt32(int32(cb.Code)).Consume(), e.Plugin.Runtime.NewString(cb.Data).Consume())

}
func (e *JsEngine) OnAssist(b *bus.Bus, script string) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, script)
}

type jsRun struct {
	ctx *v8js.Context
	cmd string
}

func (r *jsRun) Run() {
	r.ctx.RunScript(r.cmd, "run").Release()
}
func (e *JsEngine) Run(b *bus.Bus, cmd string) {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	r := &jsRun{
		ctx: e.Plugin.Runtime,
		cmd: cmd,
	}
	b.HandleScriptError(util.Catch(r.Run))
}

type runScript struct {
	bus    *bus.Bus
	output bool
	source string
	ctx    *v8js.Context
	args   []*v8js.Consumed
}

func (s *runScript) Exec() {
	fn := s.ctx.RunScript(s.source, "")
	if fn == nil || fn.IsNull() || !fn.IsFunction() {
		s.bus.HandleScriptError(errors.New(fmt.Sprintf("js function %s not found", s.source)))
		s.output = false
		return
	}
	defer fn.Release()
	var result *v8js.JsValue
	result = fn.Call(s.ctx.NullValue(), s.args...)
	defer result.Release()
	if result.IsBoolean() {
		s.output = result.Boolean()
	} else {
		s.output = false
	}
}
func (e *JsEngine) Call(b *bus.Bus, source string, args ...*v8js.Consumed) bool {
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
		ctx:    r,
		args:   args,
	}
	b.HandleScriptError(util.Catch(s.Exec))
	return output
}

func init() {
	v8go.SetFlags("--expose-gc")
}
