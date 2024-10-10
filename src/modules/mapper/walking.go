package mapper

import (
	"container/list"
)

func ValidateTags(tags map[string]bool, p *Path) bool {
	var matched = false
	for k, v := range tags {
		if v {
			if p.ExcludeTags[k] {
				return false
			}
			if p.Tags[k] {
				matched = true
			}
		}
	}
	return len(p.Tags) == 0 || matched
}

type Step struct {
	To      string
	From    string
	Command string
	Delay   int
	remain  int
}

var EmptyStep = &Step{}

type Walking struct {
	tags        map[string]bool
	rooms       *Rooms
	from        string
	to          []string
	fly         []*Path
	walked      map[string]*Step
	forwading   *list.List
	blacklist   map[string]bool
	whitelist   map[string]bool
	blockedpath map[string]map[string]bool
	maxdistance int
}

func (w *Walking) FlyStep(p *Path) *Step {
	step := w.step(p)
	step.From = w.from
	return step
}
func (w *Walking) step(p *Path) *Step {
	length := p.Delay
	if length < 1 {
		length = 1
	}
	return &Step{
		To:      p.To,
		From:    p.From,
		Command: p.Command,
		Delay:   length,
		remain:  length,
	}
}

func (w *Walking) validateExit(p *Path) bool {
	if w.blacklist[p.To] {
		return false
	}
	if len(w.whitelist) > 0 && !w.whitelist[p.To] {
		return false
	}
	f := w.blockedpath[p.From]
	if f != nil {
		if f[p.To] {
			return false
		}
	}
	return ValidateTags(w.tags, p)
}
func (w *Walking) Walk() []*Step {
	distance := 0
	rooms := *w.rooms
	var tolist = map[string]bool{}
	room := rooms.GetRoom(w.from)
	texits := rooms.GetTemporaryPaths(w.from)
	if room == nil && texits == nil {
		return nil
	}
	w.walked[w.from] = EmptyStep
	if room != nil {
		for _, v := range room.Exits {
			if w.walked[v.To] == nil && w.validateExit(v) {
				w.forwading.PushBack(w.step(v))
			}
		}
	}
	if texits != nil {
		for _, v := range texits {
			if w.walked[v.To] == nil && w.validateExit(v) {
				w.forwading.PushBack(w.step(v))
			}
		}
	}
	for _, v := range w.fly {
		if w.walked[v.To] == nil && w.validateExit(v) {
			w.forwading.PushBack(w.FlyStep(v))
		}
	}
	if w.forwading.Len() == 0 {
		return nil
	}
	for _, v := range w.to {
		if rooms.GetRoom(v) == nil && rooms.GetTemporaryPaths(v) == nil {
			continue
		}
		if w.from == v {
			return []*Step{}
		}
		tolist[v] = true
	}
	if len(tolist) == 0 {
		return nil
	}
	var matchedRoom = ""
Matching:
	for {
		newExits := list.New()
		distance++
		if w.maxdistance > 0 && distance > w.maxdistance {
			break
		}
		for {
			v := w.forwading.Front()
			if v == nil {
				break
			}
			step := v.Value.(*Step)
			w.forwading.Remove(v)
			room := rooms.GetRoom(step.To)
			texits := rooms.GetTemporaryPaths(step.To)

			if w.walked[step.To] != nil || (room == nil && texits == nil) {
				continue
			}
			if w.maxdistance > 0 && step.Delay > w.maxdistance {
				continue
			}
			step.remain--
			if step.remain > 0 {
				newExits.PushBack(step)
				continue
			}
			w.walked[step.To] = step
			if tolist[step.To] {
				matchedRoom = step.To
				break Matching
			}
			if room != nil {
				for _, exit := range room.Exits {
					if w.walked[exit.To] == nil && w.validateExit(exit) {
						newExits.PushBack(w.step(exit))
					}
				}
			}
			if texits != nil {
				for _, exit := range texits {
					if w.walked[exit.To] == nil && w.validateExit(exit) {
						newExits.PushBack(w.step(exit))
					}
				}
			}

		}
		w.forwading.PushBackList(newExits)
		if w.forwading.Len() == 0 {
			break Matching
		}
	}
	if matchedRoom == "" {
		return nil
	}
	result := list.New()
	step := w.walked[matchedRoom]
	for {
		if step == nil || step == EmptyStep {
			break
		}
		result.PushFront(step)
		step = w.walked[step.From]
	}
	steps := make([]*Step, result.Len())
	var i = 0
	v := result.Front()
	for {
		if v == nil {
			break
		}
		steps[i] = v.Value.(*Step)
		i++
		v = v.Next()
	}
	return steps
}

func NewWalking(option *Option) *Walking {
	walking := &Walking{
		tags:        map[string]bool{},
		to:          []string{},
		fly:         []*Path{},
		walked:      map[string]*Step{},
		forwading:   list.New(),
		blacklist:   map[string]bool{},
		whitelist:   map[string]bool{},
		blockedpath: map[string]map[string]bool{},
	}
	if option != nil {
		for _, v := range option.Blacklist {
			walking.blacklist[v] = true
		}
		for _, v := range option.Whitelist {
			walking.whitelist[v] = true
		}
		for _, v := range option.BlockedPath {
			if len(v) == 2 {
				f := walking.blockedpath[v[0]]
				if f == nil {
					f = map[string]bool{}
					walking.blockedpath[v[0]] = f
				}
				f[v[1]] = true
			}
		}
	}
	return walking
}
