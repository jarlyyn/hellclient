package mapper

type WalkAllResult struct {
	Steps     []*Step
	Walked    []string
	NotWalked []string
}

var NewWalkAll = func() *WalkAll {
	return &WalkAll{}
}

type WalkAll struct {
	Targets     []string
	MaxDistance int
	rooms       *map[string]*Room
	tags        map[string]bool
	fly         []*Path
}

func (a *WalkAll) Walk(fr string, to []string) []*Step {
	w := a.newWalking()
	w.from = fr
	w.to = to
	w.maxdistance = a.MaxDistance
	return w.Walk()
}
func (a *WalkAll) newWalking() *Walking {
	walking := NewWalking()
	walking.rooms = a.rooms
	walking.tags = a.tags
	walking.fly = a.fly
	return walking

}
func (a *WalkAll) filter(in []string, filtered string) []string {
	result := make([]string, 0, len(in))
	for _, v := range in {
		if v != filtered {
			result = append(result, v)
		}
	}
	return result
}
func (a *WalkAll) Start() *WalkAllResult {
	result := &WalkAllResult{}
	if len(a.Targets) < 2 {
		return result
	}
	fr := a.Targets[0]
	left := append([]string{}, a.Targets[1:]...)
	result.Walked = []string{fr}
	for len(left) > 0 {

		steps := a.Walk(fr, left)
		if len(steps) == 0 {
			break
		}
		result.Steps = append(result.Steps, steps...)
		fr = steps[len(steps)-1].To
		left = a.filter(left, fr)
		result.Walked = append(result.Walked, fr)
	}
	result.NotWalked = left
	return result
}
