package trigger

import (
	"modules/world/bus"
	"regexp"
	"strings"
)

type MatchResult struct {
	List  []string
	Named map[string]string
}

func NewMatchResult() *MatchResult {
	return &MatchResult{
		List:  []string{},
		Named: map[string]string{},
	}
}

type Context struct {
	Bus      *bus.Bus
	Line     *bus.Line
	Expanded *strings.Replacer
}
type Matcher interface {
	Match(line *bus.Line) (*MatchResult, error)
}

type Trigger struct {
	Deleted  bool
	Data     *bus.Trigger
	Matcher  Matcher
	ByUser   bool
	RawMatch string
}

func (t *Trigger) OnSuccess(ctx *Context, result *MatchResult) {

}
func (t *Trigger) BuildMatcher(match string) error {
	if t.Data.Regexp {
		re, err := regexp.Compile(match)
		if err != nil {
			return err
		}
		t.Matcher = &RegexpMatcher{matcher: re}
		return nil
	}
	if t.Data.IgnoreCase {
		t.Matcher = &StarMatcher{
			matcher: starIC.New(match),
		}
	} else {
		t.Matcher = &StarMatcher{
			matcher: starNI.New(match),
		}
	}
	return nil
}
func (t *Trigger) Match(ctx *Context) (*MatchResult, error) {
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
	return t.Matcher.Match(ctx.Line)
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
