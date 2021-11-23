package automation

import (
	"modules/world"
	"sort"
	"sync"
)

type AliasQueue []*Alias

// Len is the number of elements in the collection.
func (q AliasQueue) Len() int {
	return len(q)
}

// Less reports whether the element with index i
func (q AliasQueue) Less(i, j int) bool {
	if q[i].Deleted != q[j].Deleted {
		return !q[i].Deleted
	}
	if q[i].Data.Sequence != q[j].Data.Sequence {
		return q[i].Data.Sequence < q[j].Data.Sequence
	}
	if q[i].Data.Enabled != q[j].Data.Enabled {
		return q[i].Data.Enabled
	}
	if q[i].Data.ByUser() != q[j].Data.ByUser() {
		return q[i].Data.ByUser()
	}
	return q[i].Data.ID < q[j].Data.ID
}

// Swap swaps the elements with indexes i and j.
func (q AliasQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

type Aliases struct {
	Locker      sync.RWMutex
	All         map[string]*Alias
	ByUser      map[string]*Alias
	ByScript    map[string]*Alias
	Named       map[string]*Alias
	Temporary   map[string]*Alias
	Grouped     map[string]map[string]*Alias
	Disorder    bool
	cachedqueue []*Alias
}

func (a *Aliases) unloadAlias(id string) *Alias {
	al, ok := a.All[id]
	if !ok {
		return nil
	}
	al.Locker.Lock()
	defer al.Locker.Unlock()
	delete(a.All, id)
	if !al.Data.Temporary {
		if al.Data.ByUser() {
			delete(a.ByUser, al.Data.ID)
		} else {
			delete(a.ByScript, al.Data.ID)
		}
	}

	if al.Data.Name != "" {
		delete(a.Named, al.Data.PrefixedName())
	}
	if al.Data.Group != "" {
		delete(a.Grouped[al.Data.Group], al.Data.ID)
		if len(a.Grouped[al.Data.Group]) == 0 {
			delete(a.Grouped, al.Data.Group)
		}
	}
	if al.Data.Temporary {
		delete(a.Temporary, al.Data.ID)
	}
	a.Disorder = true
	al.Matcher = nil
	return al
}
func (a *Aliases) loadAlias(alias *Alias) {
	al := alias.Data
	a.All[al.ID] = alias
	if !al.Temporary {
		if al.ByUser() {
			a.ByUser[al.ID] = alias
		} else {
			a.ByScript[al.ID] = alias
		}
	}
	if al.Name != "" {
		a.Named[al.PrefixedName()] = alias
	}
	if al.Group != "" {
		g, ok := a.Grouped[al.Group]
		if !ok {
			g = map[string]*Alias{}
			a.Grouped[al.Group] = g
		}
		g[al.ID] = alias
	}
	if al.Temporary {
		a.Temporary[al.ID] = alias
	}
	a.Disorder = true
}
func (a *Aliases) createAlias(al *world.Alias) *Alias {
	result := &Alias{
		Data: al,
	}
	return result
}
func (a *Aliases) addAlias(al *world.Alias) bool {
	if a.All[al.ID] != nil {
		return false
	}
	alias := a.createAlias(al)
	a.loadAlias(alias)

	return true
}
func (a *Aliases) removeAlias(id string) bool {
	al := a.unloadAlias(id)
	if al == nil {
		return false

	}
	al.Locker.Lock()
	al.Deleted = true
	al.Locker.Unlock()
	return true
}

func (a *Aliases) Queue() AliasQueue {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	if !a.Disorder {
		return a.cachedqueue
	}
	q := make(AliasQueue, 0, len(a.All))
	for _, v := range a.All {
		q = append(q, v)
	}
	sort.Sort(q)
	a.cachedqueue = q
	a.Disorder = false
	return q
}

func (a *Aliases) AddAlias(al *world.Alias, replace bool) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	name := al.PrefixedName()
	if name != "" {
		named := a.Named[name]
		if named != nil {
			if !replace {
				return false
			}
			a.removeAlias(named.Data.ID)
		}
	}
	return a.addAlias(al)
}
func (a *Aliases) RemoveAlias(id string) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	return a.removeAlias(id)
}
func (a *Aliases) AddAliases(ts []*world.Alias) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	for _, v := range ts {
		a.addAlias(v)
	}
}
func (a *Aliases) DoDeleteAliasByType(byuser bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	var list map[string]*Alias
	if byuser {
		list = a.ByUser
	} else {
		list = a.ByScript
	}
	for _, v := range list {
		a.removeAlias(v.Data.ID)
	}
}
func (a *Aliases) GetAlias(id string) *world.Alias {
	a.Locker.Lock()
	al := a.All[id]
	a.Locker.Unlock()
	if al == nil {
		return nil
	}
	al.Locker.Lock()
	defer al.Locker.Unlock()
	return al.Data
}
func (a *Aliases) DoUpdateAlias(al *world.Alias) int {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	old := a.All[al.ID]
	if old == nil {
		return world.UpdateFailNotFound
	}
	al.SetByUser(old.Data.ByUser())
	if al.Name != "" && al.Name != old.Data.Name && a.Named[al.PrefixedName()] != nil {
		return world.UpdateFailDuplicateName
	}
	a.unloadAlias(old.Data.ID)
	old.Locker.Lock()
	old.Data = al
	old.Locker.Unlock()
	a.loadAlias(old)
	return world.UpdateOK
}

func (a *Aliases) DoDeleteAliasByName(name string) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	al := a.Named[name]
	if al == nil {
		return false
	}
	return a.removeAlias(al.Data.ID)

}
func (a *Aliases) DoDeleteTemporaryAliases() int {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	count := len(a.Temporary)
	for _, v := range a.Temporary {
		a.removeAlias(v.Data.ID)
	}
	return count

}
func (a *Aliases) DoDeleteAliasGroup(group string, byUser bool) int {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	count := 0
	for _, v := range a.Grouped[group] {
		if v.ByUser == byUser {
			count++
			a.removeAlias(v.Data.ID)
		}
	}
	return count
}
func (a *Aliases) DoEnableAliasByName(name string, enabled bool) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	al := a.Named[name]
	if al == nil {
		return false
	}
	al.Locker.Lock()
	defer al.Locker.Unlock()
	al.Data.Enabled = enabled
	a.Disorder = true
	return true

}
func (a *Aliases) DoEnableAliasGroup(group string, enabled bool) int {
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

func (a *Aliases) GetAliasesByType(byuser bool) []*world.Alias {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	var all map[string]*Alias
	if byuser {
		all = a.ByUser
	} else {
		all = a.ByScript
	}
	result := make([]*world.Alias, 0, len(all))
	for _, v := range all {
		v.Locker.Lock()
		result = append(result, v.Data)
		v.Locker.Unlock()
	}
	return result
}

func (a *Aliases) GetAliasOption(name string, option string) (string, bool, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	al := a.Named[name]
	if al == nil {
		return "", false, false
	}
	result, ok := al.Option(option)
	return result, ok, true
}
func (a *Aliases) GetAliasInfo(name string, infotype int) (string, bool, bool) {
	a.Locker.Lock()
	al := a.Named[name]
	a.Locker.Unlock()
	if al == nil {
		return "", false, false
	}
	result, ok := al.Info(infotype)
	return result, ok, true
}

func (a *Aliases) SetAliasOption(name string, option string, value string) (bool, bool, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	al := a.Named[name]
	if al == nil {
		return false, false, false
	}
	result, ok := al.SetOption(option, value)
	if result && ok {
		a.unloadAlias(al.Data.ID)
		a.loadAlias(al)
	}
	return result, ok, true
}

func (a *Aliases) HasNamedAlias(name string) bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	_, ok := a.Named[name]
	return ok

}
func (a *Aliases) DoListAliasNames(byUser bool) []string {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	result := make([]string, 0, len(a.Named))
	for _, v := range a.Named {
		if v.ByUser == byUser {
			result = append(result, v.Data.Name)
		}
	}
	return result
}

func NewAliases() *Aliases {
	return &Aliases{
		All:       map[string]*Alias{},
		ByUser:    map[string]*Alias{},
		ByScript:  map[string]*Alias{},
		Named:     map[string]*Alias{},
		Temporary: map[string]*Alias{},
		Grouped:   map[string]map[string]*Alias{},
	}
}
