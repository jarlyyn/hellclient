package automation

import (
	"modules/world"
	"sync"
)

type Timers struct {
	Locker    sync.RWMutex
	All       map[string]*Timer
	ByUser    map[string]*Timer
	ByScript  map[string]*Timer
	Named     map[string]*Timer
	Temporary map[string]*Timer
	Grouped   map[string]map[string]*Timer
	OnFire    func(*world.Timer)
}

func (t *Timers) TimerCallback(timer *world.Timer) {
	t.Locker.Lock()
	onfire := t.OnFire
	ti := *timer
	if ti.OneShot {
		t.removeTimer(ti.ID)
	}
	t.Locker.Unlock()
	onfire(&ti)

}
func (t *Timers) createTimer(ti *world.Timer) *Timer {
	result := &Timer{
		Data:   ti,
		OnFire: t.TimerCallback,
	}
	return result
}
func (t *Timers) loadTimer(timer *Timer) {
	ti := timer.Data
	t.All[ti.ID] = timer
	if !ti.Temporary {
		if ti.ByUser() {
			t.ByUser[ti.ID] = timer
		} else {
			t.ByScript[ti.ID] = timer
		}
	}
	if ti.Name != "" {
		t.Named[ti.PrefixedName()] = timer
	}
	if ti.Group != "" {
		g, ok := t.Grouped[ti.Group]
		if !ok {
			g = map[string]*Timer{}
			t.Grouped[ti.Group] = g
		}
		g[ti.ID] = timer
	}
	if ti.Temporary {
		t.Temporary[ti.ID] = timer
	}
	if ti.Enabled {
		go timer.Start()
	}
}
func (t *Timers) addTimer(ti *world.Timer) bool {
	if t.All[ti.ID] != nil {
		return false
	}
	timer := t.createTimer(ti)
	t.loadTimer(timer)

	return true
}
func (t *Timers) AddTimer(ti *world.Timer, replace bool) bool {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	name := ti.PrefixedName()
	if name != "" {
		named := t.Named[name]
		if named != nil {
			if !replace {
				return false
			}
			t.removeTimer(named.Data.ID)
		}
	}
	return t.addTimer(ti)
}

func (t *Timers) removeTimer(id string) bool {
	ti := t.unloadTimer(id)
	if ti == nil {
		return false
	}
	go ti.Stop()
	return true
}

func (t *Timers) unloadTimer(id string) *Timer {
	ti, ok := t.All[id]
	if !ok {
		return nil
	}
	delete(t.All, id)
	if !ti.Data.Temporary {
		if ti.Data.ByUser() {
			delete(t.ByUser, ti.Data.ID)
		} else {
			delete(t.ByScript, ti.Data.ID)
		}
	}

	if ti.Data.Name != "" {
		delete(t.Named, ti.Data.PrefixedName())
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
	go ti.Stop()
	return ti
}
func (t *Timers) RemoveTimer(id string) bool {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	return t.removeTimer(id)
}
func (t *Timers) RemoveTimerByName(name string) bool {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	ti := t.Named[name]
	if ti == nil {
		return false
	}
	return t.removeTimer(ti.Data.ID)
}
func (t *Timers) DeleteTemporaryTimers() int {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	count := len(t.Temporary)
	for _, v := range t.Temporary {
		t.removeTimer(v.Data.ID)
	}
	return count
}
func (t *Timers) DeleteTimerGroup(group string) int {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	count := len(t.Grouped[group])
	for _, v := range t.Grouped[group] {
		t.removeTimer(v.Data.ID)
	}
	return count
}
func (t *Timers) EnableTimerByName(name string, enabled bool) bool {
	t.Locker.Lock()
	ti := t.Named[name]
	t.Locker.Unlock()
	if ti == nil {
		return false
	}
	ti.Locker.Lock()
	ti.Data.Enabled = enabled
	if enabled {
		go ti.Start()
	} else {
		go ti.Stop()
	}
	ti.Locker.Unlock()
	return true
}
func (t *Timers) EnableTimerGroup(group string, enabled bool) int {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	count := len(t.Grouped[group])
	for _, v := range t.Grouped[group] {
		v.Locker.Lock()
		v.Data.Enabled = true
		v.Locker.Unlock()
		v.Start()
	}
	return count
}
func (t *Timers) ListTimerNames() []string {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	result := make([]string, 0, len(t.Named))
	for _, v := range t.Named {
		result = append(result, v.Data.Name)
	}
	return result
}
func (t *Timers) HasNamedTimer(name string) bool {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	_, ok := t.Named[name]
	return ok
}

func (t *Timers) ResetNamedTimer(name string) bool {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	ti := t.Named[name]
	if ti == nil {
		return false
	}
	ti.Reset()
	return true
}
func (t *Timers) ResetTimers() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	for _, v := range t.All {
		v.Reset()
	}
}
func (t *Timers) GetTimerOption(name string, option string) (string, bool, bool) {
	t.Locker.Lock()
	ti := t.Named[name]
	t.Locker.Unlock()
	if ti == nil {
		return "", false, false
	}
	result, ok := ti.Option(option)
	return result, ok, true
}
func (t *Timers) GetTimerInfo(name string, infotype int) (string, bool, bool) {
	t.Locker.Lock()
	ti := t.Named[name]
	t.Locker.Unlock()
	if ti == nil {
		return "", false, false
	}
	result, ok := ti.Info(infotype)
	return result, ok, true
}

func (t *Timers) SetTimerOption(name string, option string, value string) (bool, bool, bool) {
	t.Locker.Lock()
	ti := t.Named[name]
	t.Locker.Unlock()

	if ti == nil {
		return false, false, false
	}
	result, ok := ti.SetOption(option, value)
	if result && ok {
		t.Locker.Lock()
		t.unloadTimer(ti.Data.ID)
		t.loadTimer(ti)
		t.Locker.Unlock()
	}
	return result, ok, true
}

func (t *Timers) GetTimersByType(byuser bool) []*world.Timer {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	var all map[string]*Timer
	if byuser {
		all = t.ByUser
	} else {
		all = t.ByScript
	}
	result := make([]*world.Timer, 0, len(all))
	for _, v := range all {
		result = append(result, v.Data)
	}
	return result
}
func (t *Timers) AddTimers(ts []*world.Timer) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	for _, v := range ts {
		t.addTimer(v)
	}
}
func (t *Timers) DoDeleteTimerByType(byuser bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	var list map[string]*Timer
	if byuser {
		list = t.ByUser
	} else {
		list = t.ByScript
	}
	for _, v := range list {
		t.removeTimer(v.Data.ID)
	}
}
func (t *Timers) GetTimer(id string) *world.Timer {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	ti := t.All[id]
	if ti == nil {
		return nil
	}
	return ti.Data
}
func (t *Timers) DoUpdateTimer(ti *world.Timer) int {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	old := t.All[ti.ID]
	if old == nil {
		return world.UpdateFailNotFound
	}
	ti.SetByUser(old.Data.ByUser())
	if ti.Name != "" && ti.Name != old.Data.Name && t.Named[ti.PrefixedName()] != nil {
		return world.UpdateFailDuplicateName
	}
	t.unloadTimer(old.Data.ID)
	old.Data = ti
	t.loadTimer(old)
	return world.UpdateOK
}
func (t *Timers) Flush() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	for _, v := range t.All {
		t.removeTimer(v.Data.ID)
	}
}
func NewTimers() *Timers {
	timers := &Timers{}
	timers.All = map[string]*Timer{}
	timers.Named = map[string]*Timer{}
	timers.Temporary = map[string]*Timer{}
	timers.Grouped = map[string]map[string]*Timer{}
	timers.ByUser = map[string]*Timer{}
	timers.ByScript = map[string]*Timer{}
	return timers
}
