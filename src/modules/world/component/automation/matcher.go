package automation

import (
	"modules/world"
	"regexp"
)

func BuildMatcher(match string, isregexp bool, ignore_case bool) (world.Matcher, error) {
	if isregexp {
		re, err := regexp.Compile(match)
		if err != nil {
			return nil, err
		}
		return &RegexpMatcher{matcher: re}, nil
	}
	if ignore_case {
		return &StarMatcher{matcher: starIC.New(match)}, nil
	}
	return &StarMatcher{matcher: starNI.New(match)}, nil
}
