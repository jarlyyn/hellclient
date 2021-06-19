package trigger

import (
	"modules/world/bus"
	"sort"
	"sync"
)

type Triggers struct {
	Locker sync.RWMutex
	Map    map[string]*Trigger
	List   TriggerList
	Groups map[string]map[string]*Trigger
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
	tr := t.Map[id]
	if tr != nil {
		delete(t.Map, id)
		tr.Deleted = true
	}
}
func (t *Triggers) init() {
	t.Map = map[string]*Trigger{}
	t.List = TriggerList{}
	t.Groups = map[string]map[string]*Trigger{}

}
func (t *Triggers) Flush() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.init()
}
func (t *Triggers) Exec(bus *bus.Bus, line *bus.Line) {
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
func New() *Triggers {
	t := &Triggers{}
	t.init()
	return t
}
