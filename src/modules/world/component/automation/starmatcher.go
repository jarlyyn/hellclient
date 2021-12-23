package automation

import (
	"hellclient/modules/world"

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

func (m *StarMatcher) Match(message string) (*world.MatchResult, error) {
	ok, found := m.matcher.Find(message)
	if ok {
		r := world.NewMatchResult()
		r.List = found
		return r, nil
	}
	return nil, nil
}
