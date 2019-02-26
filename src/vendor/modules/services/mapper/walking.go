package mapper

import (
	"container/list"
)

type Step struct {
	To      string
	From    string
	Command string
	Delay   int
	Length  int
}

var EmptyStep = &Step{}

type Walking struct {
	Tags      []string
	Rooms     *map[string]*Room
	From      string
	To        []string
	Fly       []*Path
	Walked    map[string]*Step
	Forwading *list.List
}

func (w *Walking) FlyStep(p *Path) *Step {
	step := w.Step(p)
	step.From = w.From
	return step
}
func (w *Walking) Step(p *Path) *Step {
	length := p.Delay
	if length < 1 {
		length = 1
	}
	return &Step{
		To:      p.To,
		From:    p.From,
		Command: p.Command,
		Delay:   length,
		Length:  length,
	}
}

func (w *Walking) ValidateTags(p *Path) bool {
	var matched = false
	for _, v := range w.Tags {
		if p.ExcludeTags[v] {
			return false
		}
		if p.Tags[v] {
			matched = true
		}
	}
	return len(p.Tags) == 0 || matched
}
func (w *Walking) Walk() []*Step {
	rooms := *w.Rooms
	var tolist = map[string]bool{}
	if rooms[w.From] == nil {
		return nil
	}
	w.Walked[w.From] = EmptyStep

	for _, v := range rooms[w.From].Exits {
		if w.Walked[v.To] == nil && w.ValidateTags(v) {
			w.Forwading.PushBack(w.Step(v))
		}
	}
	for _, v := range w.Fly {
		if w.Walked[v.To] == nil && w.ValidateTags(v) {
			w.Forwading.PushBack(w.FlyStep(v))
		}
	}
	if w.Forwading.Len() == 0 {
		return nil
	}
	for _, v := range w.To {
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
		v := w.Forwading.Front()
		newExits := list.New()
		for {
			if v == nil {
				break
			}
			step := v.Value.(*Step)
			if w.Walked[step.To] == nil {
				if step.Delay > 0 {
					step.Delay--
				} else {
					w.Walked[step.To] = step
					if tolist[step.To] {
						matchedRoom = step.To
						break Matching
					}
					if rooms[step.To] == nil {
						break
					}
					for _, v := range rooms[step.To].Exits {
						if w.Walked[v.To] == nil && w.ValidateTags(v) {
							newExits.PushBack(w.Step(v))
						}
					}
				}
			} else {
				w.Forwading.Remove(v)
			}
			v = v.Next()
		}
		w.Forwading.PushBackList(newExits)
		if w.Forwading.Len() == 0 {
			break Matching
		}
	}
	if matchedRoom == "" {
		return nil
	}
	result := list.New()
	step := w.Walked[matchedRoom]
	for {
		if step == nil || step == EmptyStep {
			break
		}
		result.PushFront(step)
		step = w.Walked[step.From]
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
		Tags:      []string{},
		To:        []string{},
		Fly:       []*Path{},
		Walked:    map[string]*Step{},
		Forwading: list.New(),
	}
}
