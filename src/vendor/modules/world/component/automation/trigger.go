package automation

import (
	"modules/world"
	"modules/world/bus"
	"regexp"
	"strings"
)

type Context struct {
	Bus      *bus.Bus
	Line     *world.Line
	Expanded *strings.Replacer
}

type Trigger struct {
	Deleted  bool
	Data     *world.Trigger
	Matcher  world.Matcher
	ByUser   bool
	RawMatch string
}

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

func (t *Trigger) OnSuccess(ctx *Context, result *world.MatchResult) {

}
func (t *Trigger) BuildMatcher(match string) error {
	matcher, err := BuildMatcher(match, t.Data.Regexp, t.Data.IgnoreCase)
	if err != nil {
		return err
	}
	t.Matcher = matcher
	return nil
}
func (t *Trigger) Match(ctx *Context) (*world.MatchResult, error) {
	if t.Deleted || !t.Data.Enabled {
		return nil, nil
	}
	if t.Data.ExpandVariables {
		if ctx.Expanded == nil {
			params := ctx.Bus.GetParams()
			r := make([]string, 0, len(params)*2)
			for k := range params {
				r = append(r, k, params[k])
			}
			ctx.Expanded = strings.NewReplacer(r...)
		}
		match := ctx.Expanded.Replace(t.Data.Match)
		if t.RawMatch == "" || t.RawMatch != match {
			err := t.BuildMatcher(match)
			if err != nil {
				return nil, err
			}
			t.RawMatch = match
		}
	} else {
		if t.Matcher == nil {
			err := t.BuildMatcher(t.Data.Match)
			if err != nil {
				return nil, err
			}
		}
	}
	return t.Matcher.Match(ctx.Line.Plain())
}

type TriggerList []*Trigger

// Len is the number of elements in the collection.
func (t TriggerList) Len() int {
	return len(t)
}

// Less reports whether the element with index i
func (t TriggerList) Less(i, j int) bool {
	if t[i].Deleted != t[j].Deleted {
		return t[j].Deleted
	}
	return t[i].Data.Sequence < t[j].Data.Sequence
}

// Swap swaps the elements with indexes i and j.
func (t TriggerList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
