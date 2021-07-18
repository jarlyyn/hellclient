package jsengine

import (
	"errors"

	"github.com/dop251/goja"
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
