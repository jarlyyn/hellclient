package jsengine

import (
	"errors"
	"fmt"
	"modules/world"
	"modules/world/bus"
	"strconv"
	"sync"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/util"

	"github.com/herb-go/herbplugin/jsplugin"
)

func newJsInitializer(b *bus.Bus) *jsplugin.Initializer {
	i := jsplugin.NewInitializer()
	i.Entry = "main.js"
	i.DisableBuiltin = true
	i.Modules = []*herbplugin.Module{
		NewMapperModule(b),
		NewAPIModule(b),
		ModuleEval,
		ModuleJScript,
		NewFileSystemObjectModule(b),
		NewMetronomeModule(b),
		NewUserinputModule(b),
	}
	return i
}

type JsEngine struct {
	Locker       sync.RWMutex
	Plugin       *jsplugin.Plugin
	onClose      string
	onDisconnect string
	onConnect    string
	onBroadCast  string
}

func NewJsEngine() *JsEngine {
	return &JsEngine{
		Plugin: jsplugin.New(),
	}
}

func (e *JsEngine) Open(b *bus.Bus) error {
	opt := b.GetScriptPluginOptions()
	data := b.GetScriptData()
	e.onClose = data.OnClose
	e.onConnect = data.OnConnect
	e.onDisconnect = data.OnDisconnect
	e.onBroadCast = data.OnBroadcast
	err := util.Catch(func() {
		newJsInitializer(b).MustApplyInitializer(e.Plugin)
	})
	if err != nil {
		return err
	}
	err = util.Catch(func() {
		herbplugin.Lanuch(e.Plugin, opt)
	})
	if err != nil {
		return err
	}
	if data.OnOpen != "" {
		open, ok := goja.AssertFunction(e.Plugin.Runtime.Get(data.OnOpen))
		if ok {
			open(goja.Undefined())
		}
		b.HandleScriptError(err)
	}
	return nil
}
func (e *JsEngine) Close(b *bus.Bus) {
	if e.onClose != "" {
		e.Call(b, e.onClose)
	}
	b.HandleScriptError(util.Catch(func() {
		e.Plugin.MustClosePlugin()
	}))
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
	t := e.Plugin.Runtime.NewObject()
	for k, v := range result.Named {
		t.Set(k, v)
	}
	for k, v := range result.List {
		if k == 0 {
			t.Set("10", v)
			continue
		} else if k > 9 {
			break
		}
		t.Set(strconv.Itoa(k-1), v)
	}

	e.Locker.Unlock()
	e.Call(b, trigger.Script, trigger.Name, line.Plain(), t)

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
	t := e.Plugin.Runtime.NewObject()

	for k, v := range result.Named {
		t.Set(k, v)
	}
	for k, v := range result.List {
		if k == 0 {
			t.Set("10", v)
			continue
		} else if k > 9 {
			break
		}
		t.Set(strconv.Itoa(k-1), v)
	}
	e.Locker.Unlock()
	go e.Call(b, alias.Script, alias.Name, message, t)

}
func (e *JsEngine) OnTimer(b *bus.Bus, timer *world.Timer) {
	if timer.Script == "" {
		return
	}
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, timer.Script, timer.Name)
}
func (e *JsEngine) OnBroadCast(b *bus.Bus, bc *world.Broadcast) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, bc.Message, bc.Global, bc.Channel, bc.ID)
}

func (e *JsEngine) OnCallback(b *bus.Bus, cb *world.Callback) {
	e.Locker.Lock()
	if e.Plugin.Runtime == nil {
		e.Locker.Unlock()
		return
	}
	e.Locker.Unlock()
	go e.Call(b, cb.Script, cb.Name, cb.ID, cb.Code, cb.Data)

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
func (e *JsEngine) Run(b *bus.Bus, cmd string) {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	_, err := e.Plugin.Runtime.RunString(cmd)
	b.HandleScriptError(err)
}

func (e *JsEngine) Call(b *bus.Bus, source string, args ...interface{}) {
	e.Locker.Lock()
	defer e.Locker.Unlock()
	r := e.Plugin.Runtime
	if r == nil {
		return
	}
	s, err := r.RunString(source)
	if err != nil {
		b.HandleScriptError(err)
		return
	}
	fn, ok := goja.AssertFunction(s)
	if !ok {
		b.HandleScriptError(errors.New(fmt.Sprintf("js function %s not found", source)))
		return
	}
	jargs := []goja.Value{}
	for _, v := range args {
		jargs = append(jargs, r.ToValue(v))
	}
	b.HandleScriptError(util.Catch(func() {
		_, err := fn(goja.Undefined(), jargs...)
		b.HandleScriptError(err)
	}))
}
