package mapper

import (
	"container/list"
)

type Step struct {
	To      string
	From    string
	Command string
	Delay   int
	remain  int
}

var EmptyStep = &Step{}

type Walking struct {
	tags      map[string]bool
	rooms     *map[string]*Room
	from      string
	to        []string
	fly       []*Path
	walked    map[string]*Step
	forwading *list.List
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
func (w *Walking) validateTags(p *Path) bool {
	return ValidateTags(w.tags, p)
}
func (w *Walking) Walk() []*Step {
	rooms := *w.rooms
	var tolist = map[string]bool{}
	if rooms[w.from] == nil {
		return nil
	}
	w.walked[w.from] = EmptyStep

	for _, v := range rooms[w.from].Exits {
		if w.walked[v.To] == nil && w.validateTags(v) {
			w.forwading.PushBack(w.step(v))
		}
	}
	for _, v := range w.fly {
		if w.walked[v.To] == nil && w.validateTags(v) {
			w.forwading.PushBack(w.FlyStep(v))
		}
	}
	if w.forwading.Len() == 0 {
		return nil
	}
	for _, v := range w.to {
		if rooms[v] == nil {
			continue
		}
		tolist[v] = true
	}
	if len(tolist) == 0 {
		return nil
	}
	var matchedRoom = ""
Matching:
	for {
		v := w.forwading.Front()
		newExits := list.New()
		for {
			if v == nil {
				break
			}
			step := v.Value.(*Step)
			if w.walked[step.To] == nil {
				if step.remain > 0 {
					step.remain--
				} else {
					w.walked[step.To] = step
					if tolist[step.To] {
						matchedRoom = step.To
						break Matching
					}
					if rooms[step.To] == nil {
						break
					}
					for _, v := range rooms[step.To].Exits {
						if w.walked[v.To] == nil && w.validateTags(v) {
							newExits.PushBack(w.step(v))
						}
					}
				}
			} else {
				w.forwading.Remove(v)
			}
			v = v.Next()
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

func NewWalking() *Walking {
	return &Walking{
		tags:      map[string]bool{},
		to:        []string{},
		fly:       []*Path{},
		walked:    map[string]*Step{},
		forwading: list.New(),
	}
}
