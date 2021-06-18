package trigger

import (
	"sort"
	"sync"
)

type Triggers struct {
	Locker sync.RWMutex
	Map    map[string]*Trigger
	List   TriggerList
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

func (t *Triggers) Flush() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.Map = map[string]*Trigger{}
	t.List = TriggerList{}
}
