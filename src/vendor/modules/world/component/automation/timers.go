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

func (t *Timers) OnFire(*world.Timer) {

}
func (t *Timers) createTimer(ti *world.Timer) *Timer {
	result := &Timer{
		Data:   ti,
		OnFire: t.OnFire,
	}
	return result
}
func (t *Timers) addTimer(ti *world.Timer) bool {
	if t.All[ti.ID] != nil {
		return false
	}
	timer := t.createTimer(ti)
	t.All[ti.ID] = timer
	if ti.Name != "" {
		t.Named[ti.Name] = timer
	}
	if ti.Group != "" {
		g, ok := t.Grouped[ti.Group]
		if !ok {
			g = map[string]*Timer{}
			t.Grouped[ti.Group] = g
		}
		g[ti.Group] = timer
	}
	if ti.Temporary {
		t.Temporary[ti.ID] = timer
	}
	return true
}
func (t *Timers) AddTimer(ti *world.Timer) bool {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	return t.addTimer(ti)
}

func (t *Timers) removeTimer(id string) bool {
	ti, ok := t.All[id]
	if !ok {
		return false
	}
	ti.Stop()
	delete(t.All, id)
	if ti.Data.Name != "" {
		delete(t.Named, ti.Data.ID)
	}
	if ti.Data.Group != "" {
		delete(t.Grouped[ti.Data.Group], ti.Data.ID)
		if len(t.Grouped[ti.Data.Group]) == 0 {
			delete(t.Grouped, ti.Data.Group)
		}
	}
	if ti.Data.Temporary {
		delete(t.Temporary, ti.Data.ID)
	}
	return true
}

func (t *Timers) RemoveTimer(id string) bool {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	return t.removeTimer(id)
}
