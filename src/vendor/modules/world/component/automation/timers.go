package automation

import (
	"modules/world"
	"sync"
)

type Timers struct {
	Locker    sync.RWMutex
	All       map[string]*Timer
	Named     map[string]*Timer
	Temporary map[string]*Timer
	Grouped   map[string]map[string]*Timer
}

func (t *Timers) createTimer(ti *world.Timer) *Timer {
	result := &Timer{
		Data: ti,
	}
	return result
}

// func (t *Timers) AddTimer(ti *world.Timer) bool {
// 	t.Locker.Lock()
// 	defer t.Locker.Unlock()
// 	if t.All[ti.ID] != nil {
// 		return false
// 	}
// 	t.All[ti.ID] = ti
// }
