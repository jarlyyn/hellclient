package automation

import (
	"modules/world"
	"regexp"
)

type RegexpMatcher struct {
	matcher *regexp.Regexp
}

func (m *RegexpMatcher) Match(line *world.Line) (*MatchResult, error) {
	result := m.matcher.FindAllString(line.Plain(), -1)
	if result == nil {
		return nil, nil
	}
	r := NewMatchResult()
	r.List = result
	names := m.matcher.SubexpNames()
	for k := range names {
		if names[k] != "" {
			r.Named[names[k]] = result[k]
		}
	}
	return r, nil
}