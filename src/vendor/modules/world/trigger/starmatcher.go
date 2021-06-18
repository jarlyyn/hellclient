package trigger

import (
	"modules/world/bus"

	"github.com/herb-go/misc/starpattern"
)

var starNI = &starpattern.Options{
	Wildcard:   '*',
	IgnoreCase: false,
}
var starIC = &starpattern.Options{
	Wildcard:   '*',
	IgnoreCase: true,
}

type StarMatcher struct {
	matcher *starpattern.Pattern
}

func (m *StarMatcher) Match(line *bus.Line) (*MatchResult, error) {
	ok, found := m.matcher.Find(line.Plain())
	if ok {
		r := NewMatchResult()
		r.List = found
		return r, nil
	}
	return nil, nil
}
