package automation

import (
	"modules/world"
	"sort"
	"sync"
)

type Triggers struct {
	Locker      sync.RWMutex
	All         map[string]*Trigger
	ByUser      map[string]*Trigger
	ByScript    map[string]*Trigger
	Named       map[string]*Trigger
	Temporary   map[string]*Trigger
	Grouped     map[string]map[string]*Trigger
	Disorder    bool
	cachedqueue []*Trigger
}

func (a *Triggers) unloadTrigger(id string) *Trigger {
	tr, ok := a.All[id]
	if !ok {
		return nil
	}
	tr.Locker.Lock()
	defer tr.Locker.Unlock()
	delete(a.All, id)
	if !tr.Data.Temporary {
		if tr.Data.ByUser() {
			delete(a.ByUser, tr.Data.ID)
		} else {
			delete(a.ByScript, tr.Data.ID)
		}
	}

	if tr.Data.Name != "" {
		delete(a.Named, tr.Data.PrefixedName())
	}
	if tr.Data.Group != "" {
		delete(a.Grouped[tr.Data.Group], tr.Data.ID)
		if len(a.Grouped[tr.Data.Group]) == 0 {
			delete(a.Grouped, tr.Data.Group)
		}
	}
	if tr.Data.Temporary {
		delete(a.Temporary, tr.Data.ID)
	}
	a.Disorder = true
	tr.Matcher = nil
	return tr
}
func (a *Triggers) loadTrigger(trigger *Trigger) {
	tr := trigger.Data
	a.All[tr.ID] = trigger
	if !tr.Temporary {
		if tr.ByUser() {
			a.ByUser[tr.ID] = trigger
		} else {
			a.ByScript[tr.ID] = trigger
		}
	}
	if tr.Name != "" {
		a.Named[tr.PrefixedName()] = trigger
	}
	if tr.Group != "" {
		g, ok := a.Grouped[tr.Group]
		if !ok {
			g = map[string]*Trigger{}
			a.Grouped[tr.Group] = g
		}
		g[tr.Group] = trigger
	}
	if tr.Temporary {
		a.Temporary[tr.ID] = trigger
	}
	a.Disorder = true
}
func (a *Triggers) createTrigger(tr *world.Trigger) *Trigger {
	result := &Trigger{
		Data: tr,
	}
	return result
}
func (a *Triggers) addTrigger(tr *world.Trigger) bool {
	if a.All[tr.ID] != nil {
		return false
	}
	trigger := a.createTrigger(tr)
	a.loadTrigger(trigger)

	return true
}
func (a *Triggers) removeTrigger(id string) bool {
	tr := a.unloadTrigger(id)
	if tr == nil {
		return false

	}
	tr.Locker.Lock()
	tr.Deleted = true
	tr.Locker.Unlock()
	return true
}

func (a *Triggers) Queue() TriggerQueue {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	if !a.Disorder {
		return a.cachedqueue
	}
	q := make(TriggerQueue, 0, len(a.All))
	for _, v := range a.All {
		q = append(q, v)
	}
	sort.Sort(q)
	a.cachedqueue = q
	a.Disorder = false
	return q
}

func (a *Triggers) AddTrigger(tr *world.Trigger, replace bool) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	name := tr.PrefixedName()
	if name != "" {
		named := a.Named[name]
		if named != nil {
			if !replace {
				return false
			}
			a.removeTrigger(named.Data.ID)
		}
	}
	return a.addTrigger(tr)
}
func (a *Triggers) RemoveTrigger(id string) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	return a.removeTrigger(id)
}
func (a *Triggers) AddTriggers(ts []*world.Trigger) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	for _, v := range ts {
		a.addTrigger(v)
	}
}
func (a *Triggers) DoDeleteTriggerByType(byuser bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	var list map[string]*Trigger
	if byuser {
		list = a.ByUser
	} else {
		list = a.ByScript
	}
	for _, v := range list {
		a.removeTrigger(v.Data.ID)
	}
}
func (a *Triggers) GetTrigger(id string) *world.Trigger {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	tr := a.All[id]
	if tr == nil {
		return nil
	}
	tr.Locker.Lock()
	defer tr.Locker.Unlock()
	return tr.Data
}
func (a *Triggers) DoUpdateTrigger(tr *world.Trigger) int {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	old := a.All[tr.ID]
	if old == nil {
		return world.UpdateFailNotFound
	}
	tr.SetByUser(old.Data.ByUser())
	if tr.Name != "" && tr.Name != old.Data.Name && a.Named[tr.PrefixedName()] != nil {
		return world.UpdateFailDuplicateName
	}
	a.unloadTrigger(old.Data.ID)
	old.Locker.Lock()
	old.Data = tr
	old.Locker.Unlock()
	a.loadTrigger(old)
	return world.UpdateOK
}

func (a *Triggers) DoDeleteTriggerByName(name string) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	tr := a.Named[name]
	if tr == nil {
		return false
	}
	return a.removeTrigger(tr.Data.ID)

}
func (a *Triggers) DoDeleteTemporaryTriggers() int {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	count := len(a.Temporary)
	for _, v := range a.Temporary {
		a.removeTrigger(v.Data.ID)
	}
	return count

}
func (a *Triggers) DoDeleteTriggerGroup(group string) int {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	count := len(a.Grouped[group])
	for _, v := range a.Grouped[group] {
		a.removeTrigger(v.Data.ID)
	}
	return count
}
func (a *Triggers) DoEnableTriggerByName(name string, enabled bool) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	tr := a.Named[name]
	if tr == nil {
		return false
	}
	tr.Locker.Lock()
	defer tr.Locker.Unlock()
	tr.Data.Enabled = enabled
	a.Disorder = true
	return true

}
func (a *Triggers) DoEnableTriggerGroup(group string, enabled bool) int {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	count := len(a.Grouped[group])
	for _, v := range a.Grouped[group] {
		v.Locker.Lock()
		v.Data.Enabled = enabled
		v.Locker.Unlock()
	}
	a.Disorder = true
	return count
}

func (a *Triggers) GetTriggersByType(byuser bool) []*world.Trigger {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	var trl map[string]*Trigger
	if byuser {
		trl = a.ByUser
	} else {
		trl = a.ByScript
	}
	result := make([]*world.Trigger, 0, len(trl))
	for _, v := range trl {
		v.Locker.Lock()
		result = append(result, v.Data)
		v.Locker.Unlock()
	}
	return result
}

func (a *Triggers) GetTriggerOption(name string, option string) (string, bool, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	tr := a.Named[name]
	if tr == nil {
		return "", false, false
	}
	result, ok := tr.Option(option)
	return result, ok, true
}
func (a *Triggers) SetTriggerOption(name string, option string, vtrue string) (bool, bool, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	tr := a.Named[name]
	if tr == nil {
		return false, false, false
	}
	result, ok := tr.SetOption(option, vtrue)
	if result && ok {
		a.unloadTrigger(tr.Data.ID)
		a.loadTrigger(tr)
	}
	return result, ok, true
}
func (a *Triggers) HasNamedTrigger(name string) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	_, ok := a.Named[name]
	return ok

}
func (a *Triggers) DoListTriggerNames() []string {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	result := make([]string, 0, len(a.Named))
	for _, v := range a.Named {
		result = append(result, v.Data.Name)
	}
	return result
}

func NewTriggers() *Triggers {
	return &Triggers{
		All:       map[string]*Trigger{},
		ByUser:    map[string]*Trigger{},
		ByScript:  map[string]*Trigger{},
		Named:     map[string]*Trigger{},
		Temporary: map[string]*Trigger{},
		Grouped:   map[string]map[string]*Trigger{},
	}
}
