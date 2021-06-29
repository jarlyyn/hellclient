package automation

import (
	"modules/world"
	"sync"
)

type Alias struct {
	Locker  sync.RWMutex
	Data    *world.Alias
	Deleted bool
	Matcher world.Matcher
	ByUser  bool
}

func (a *Alias) Match(message string) (*world.MatchResult, error) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	if a.Deleted || !a.Data.Enabled {
		return nil, nil
	}
	if a.Matcher == nil {
		m, err := BuildMatcher(a.Data.Match, a.Data.Regexp, a.Data.IgnoreCase)
		if err != nil {
			return nil, err
		}
		a.Matcher = m
	}
	return a.Matcher.Match(message)
}
