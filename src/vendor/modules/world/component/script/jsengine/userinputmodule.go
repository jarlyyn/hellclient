package jsengine

import (
	"context"
	"modules/world/bus"
	"modules/world/component/script/userinput"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
)

type List struct {
	List *userinput.List
	bus  *bus.Bus
}

func (l *List) Send(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := l.List.Send(l.bus, call.Argument(0).String())
	return r.ToValue(ui.ID)
}
func (l *List) Append(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	l.List.Append(call.Argument(0).String(), call.Argument(1).String())
	return nil
}
func (l *List) Convert(r *goja.Runtime) goja.Value {
	obj := r.NewObject()
	obj.Set("append", l.Append)
	obj.Set("send", l.Send)
	return obj
}

type Userinput struct {
	bus *bus.Bus
}

func (u *Userinput) NewList(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	list := &List{
		List: userinput.CreateList(call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).ToBoolean()),
		bus:  u.bus,
	}
	return list.Convert(r)
}
func (u *Userinput) Prompt(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := userinput.SendPrompt(u.bus, call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String(), call.Argument(3).String())
	return r.ToValue(ui.ID)
}
func (u *Userinput) Confirm(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := userinput.SendConfirm(u.bus, call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String())
	return r.ToValue(ui.ID)
}
func (u *Userinput) Alert(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	ui := userinput.SendAlert(u.bus, call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String())
	return r.ToValue(ui.ID)
}

func (u *Userinput) Convert(r *goja.Runtime) goja.Value {
	obj := r.NewObject()
	obj.Set("prompt", u.Prompt)
	obj.Set("confirm", u.Confirm)
	obj.Set("alert", u.Alert)
	obj.Set("newlist", u.NewList)
	return obj
}
func NewUserinputModule(b *bus.Bus) *herbplugin.Module {
	return herbplugin.CreateModule("userinput",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
			r := jsp.Runtime
			u := &Userinput{bus: b}
			r.Set("Userinput", u.Convert(r))
			next(ctx, plugin)
		},
		nil,
		nil,
	)
}
