package automation

import (
	"modules/world"
	"regexp"
)

type RegexpMatcher struct {
	matcher *regexp.Regexp
}

func (m *RegexpMatcher) Match(message string) (*world.MatchResult, error) {
	result := m.matcher.FindStringSubmatch(message)
	if result == nil {
		return nil, nil
	}
	r := world.NewMatchResult()
	r.List = result
	names := m.matcher.SubexpNames()
	for k := range names {
		if names[k] != "" {
			r.Named[names[k]] = result[k]
		}
	}
	return r, nil
}
