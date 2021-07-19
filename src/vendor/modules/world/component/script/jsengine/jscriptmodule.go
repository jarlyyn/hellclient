package jsengine

import (
	"context"
	"errors"
	"strconv"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
)

type VBArray struct {
	data []goja.Value
}

func (a *VBArray) Dimensions(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(1)
}

func (a *VBArray) ToArray(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(a.data)
}
func (a *VBArray) LBound(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if a.data == nil {
		return goja.Undefined()
	}
	return r.ToValue(0)
}
func (a *VBArray) UBound(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if a.data == nil {
		return goja.Undefined()
	}
	idx := call.Argument(0).ToInteger()
	if idx != 1 {
		panic(errors.New("Subscript out of range"))
	}
	return r.ToValue(len(a.data))
}
func (a *VBArray) GetItem(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if a.data == nil {
		return goja.Undefined()
	}
	idx := int(call.Argument(0).ToInteger())
	if idx < 0 || idx >= len(a.data) {
		return goja.Undefined()
	}
	return a.data[idx]
}

func (a *VBArray) ConvertTo(obj *goja.Object) {
	obj.Set("dimensions", a.Dimensions)
	obj.Set("getitem", a.GetItem)
	obj.Set("lbound", a.LBound)
	obj.Set("ubound", a.UBound)
	obj.Set("toArray", a.ToArray)
}

func CreateVBArray(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	data := call.Argument(0).ToObject(r)
	result := &VBArray{}
	l := []goja.Value{}
	var i = 0
	for {
		v := data.Get(strconv.Itoa(i))
		if v == nil {
			break
		}
		l = append(l, v)
		i++
	}
	result.data = l
	result.ConvertTo(call.This)
	return call.This
}

type Enumerator struct {
	data []goja.Value
	iter int
}

func (e *Enumerator) AtEnd(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(e.iter >= len(e.data)-1)
}
func (e *Enumerator) MoveFirst(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	e.iter = 0
	return nil
}
func (e *Enumerator) MoveNext(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	e.iter++
	return nil
}
func (e *Enumerator) Item(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if e.iter < 0 || e.iter > len(e.data)-1 {
		return goja.Undefined()
	}
	return e.data[e.iter]
}

func (e *Enumerator) ConvertTo(obj *goja.Object) {
	obj.Set("atEnd", e.AtEnd)
	obj.Set("moveFirst", e.MoveFirst)
	obj.Set("moveNext", e.MoveNext)
	obj.Set("item", e.Item)
}

func CreateEnumerator(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	data := call.Argument(0).ToObject(r)
	result := &VBArray{}
	if data != nil {
		l := []goja.Value{}
		var i = 0
		for {
			v := data.Get(strconv.Itoa(i))
			if v == nil {
				break
			}
			l = append(l, v)
			i++
		}
		result.data = l
	}
	result.ConvertTo(call.This)
	return call.This
}

var ModuleJScript = herbplugin.CreateModule("jscript",
	func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
		jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
		r := jsp.Runtime
		r.Set("VBArray", CreateVBArray)
		r.Set("Enumerator", CreateEnumerator)
		next(ctx, plugin)
	},
	nil,
	nil,
)
