package script

import (
	"container/list"
	"regexp"
)

type MatchResult struct {
	Plain   string
	Matched []string
}

var failMatcher = func(data string) *MatchResult {
	return nil
}

func makePlainMatchResult(plain string) *MatchResult {
	return &MatchResult{
		Plain:   plain,
		Matched: []string{},
	}
}
func makeRegexpMatchResult(plain string, matched []string) *MatchResult {
	return &MatchResult{
		Plain:   plain,
		Matched: matched,
	}
}

type Trigger struct {
	Matcher     func(string) *MatchResult
	Enabled     bool
	Group       string
	Name        string
	IsVMTrigger bool
	Priority    int
	Commnd      string
	Callback    string
	Finally     bool
}

type basicTrigger struct {
	Pattern  string
	IsRegExp bool
	Enabled  bool
	Commnd   string
	Name     string
	Priority int
	Finally  bool
}

func (t *basicTrigger) newTrigger() *Trigger {
	trigger := &Trigger{}
	if t.IsRegExp {
		r, err := regexp.Compile(t.Pattern)
		if err != nil {
			trigger.Matcher = failMatcher
		} else {
			trigger.Matcher = func(data string) *MatchResult {
				result := r.FindStringSubmatch(data)
				if len(result) == 0 {
					return nil
				}
				return makeRegexpMatchResult(r.FindString(data), result)
			}
		}
	} else {
		trigger.Matcher = func(data string) *MatchResult {
			if data == t.Pattern {
				return makePlainMatchResult(data)
			}
			return nil
		}
	}
	trigger.Enabled = t.Enabled
	trigger.Priority = t.Priority
	trigger.Commnd = t.Commnd
	trigger.Name = t.Name
	trigger.Finally = t.Finally
	return trigger
}

type WorldTrigger struct {
	basicTrigger
	CreatedTime int64
}

func (t *WorldTrigger) Trigger() *Trigger {
	trigger := t.newTrigger()
	trigger.IsVMTrigger = false
	return trigger
}

type VMTrigger struct {
	basicTrigger
	Group    string
	Callback string
}

func (t *VMTrigger) Trigger() *Trigger {
	trigger := t.newTrigger()
	trigger.IsVMTrigger = true
	trigger.Callback = t.Callback
	return trigger
}

type Triggers struct {
	WorldTriggers   map[string]*Trigger
	VMTriggers      map[string]*Trigger
	VMTriggerGroups map[string]*list.List
	Queue           *list.List
	Enabled         bool
}

func (t *Triggers) EnableWorldTrigger(name string, enabled bool) {
	trigger := t.WorldTriggers[name]
	if trigger != nil {
		trigger.Enabled = enabled
	}
}
func (t *Triggers) EnableVMTrigger(name string, enabled bool) {
	trigger := t.VMTriggers[name]
	if trigger != nil {
		trigger.Enabled = enabled
	}
}
func (t *Triggers) EnableVMGroup(group string, enabled bool) {
	g := t.VMTriggerGroups[group]
	if g != nil {
		e := g.Front()
		for {
			if e == nil {
				break
			}
			e.Value.(*Trigger).Enabled = enabled
			e = e.Next()
		}
	}
}
func (t *Triggers) Remove(trigger *Trigger) {
	if trigger.IsVMTrigger {
		delete(t.VMTriggers, trigger.Name)
		if trigger.Group != "" {
			group, ok := t.VMTriggerGroups[trigger.Group]
			if ok != false {
				e := group.Front()
				for {
					if e == nil {
						break
					}
					if e.Value.(*Trigger) == trigger {
						group.Remove(e)
						break
					}
					e = e.Next()
				}
				if group.Front() == nil {
					delete(t.VMTriggerGroups, trigger.Group)
				} else {
					t.VMTriggerGroups[trigger.Group] = group
				}
			}
		}
	} else {
		delete(t.WorldTriggers, trigger.Name)
	}
	e := t.Queue.Front()
	for {
		if e == nil {
			break
		}
		if e.Value.(*Trigger) == trigger {
			t.Queue.Remove(e)
			break
		}
		e = e.Next()
	}

}
func (t *Triggers) Add(trigger *Trigger) {
	if trigger.IsVMTrigger {
		old, ok := t.VMTriggers[trigger.Name]
		if ok {
			t.Remove(old)
		}
		if trigger.Group != "" {
			group, ok := t.VMTriggerGroups[trigger.Group]
			if ok == false {
				group = list.New()
				t.VMTriggerGroups[trigger.Group] = group
			}
			group.PushBack(trigger)
			t.VMTriggerGroups[trigger.Group] = group
		}
	} else {
		old, ok := t.WorldTriggers[trigger.Name]
		if ok {
			t.Remove(old)
		}
	}
	e := t.Queue.Front()
	for {
		if e == nil {
			t.Queue.PushBack(trigger)
			break
		}
		v := e.Value.(*Trigger)
		if trigger.Priority >= v.Priority {
			t.Queue.PushBack(trigger)
			break
		}
		e = e.Next()
	}
}
func (t *Triggers) Match(data string) []*TriggerMatchResult {
	results := []*TriggerMatchResult{}
	if t.Enabled == false {
		return results
	}
	e := t.Queue.Front()
	for {
		if e == nil {
			break
		}
		trigger := e.Value.(*Trigger)
		if trigger.Enabled {
			r := trigger.Matcher(data)
			if r != nil {
				result := NewTriggerMatchResult()
				result.Trigger = trigger
				result.Result = r
				results = append(results, result)
				if trigger.Finally {
					break
				}
			}
		}
		e = e.Next()
	}

	return results
}
func NewTriggers() *Triggers {
	return &Triggers{
		WorldTriggers:   map[string]*Trigger{},
		VMTriggers:      map[string]*Trigger{},
		VMTriggerGroups: map[string]*list.List{},
		Queue:           list.New(),
		Enabled:         true,
	}
}

type TriggerMatchResult struct {
	Trigger *Trigger
	Result  *MatchResult
}

func NewTriggerMatchResult() *TriggerMatchResult {
	return &TriggerMatchResult{}
}
