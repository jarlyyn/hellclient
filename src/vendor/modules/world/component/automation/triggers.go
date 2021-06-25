package automation

import (
	"modules/world"
	"modules/world/bus"
	"sort"
	"sync"
)

type Triggers struct {
	Locker    sync.RWMutex
	All       map[string]*Trigger
	Named     map[string]*Trigger
	Temporary map[string]*Trigger
	List      TriggerList
	Grouped   map[string]map[string]*Trigger
}

func (t *Triggers) sort() {
	sort.Sort(t.List)
	var enabled = 0
	for enabled = len(t.List); enabled > 0; enabled-- {
		if !t.List[enabled].Deleted {
			break
		}
	}
	t.List = t.List[:enabled]
}

func (t *Triggers) remove(id string) {
	tr := t.All[id]
	if tr != nil {
		delete(t.All, id)
		tr.Deleted = true
	}
}
func (t *Triggers) init() {
	t.All = map[string]*Trigger{}
	t.List = TriggerList{}
	t.Grouped = map[string]map[string]*Trigger{}

}
func (t *Triggers) Flush() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.init()
}
func (t *Triggers) Exec(bus *bus.Bus, line *world.Line) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	ctx := &Context{
		Bus:  bus,
		Line: line,
	}
	for _, v := range t.List {
		result, err := v.Match(ctx)
		if err != nil {
			bus.HandleTriggerError(err)
		}
		if result != nil {
			v.OnSuccess(ctx, result)
		}
		if v.Data.OmitFromLog {
			line.OmitFromLog = true
		}
		if v.Data.OmitFromOutput {
			line.OmitFromOutput = true
		}
		if !v.Data.KeepEvaluating {
			break
		}
	}
}
func NewTriggers() *Triggers {
	t := &Triggers{}
	t.init()
	return t
}
