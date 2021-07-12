package automation

import (
	"modules/world"
	"modules/world/bus"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Context struct {
	Bus      *bus.Bus
	Line     *world.Line
	Expanded *strings.Replacer
}

type Trigger struct {
	Locker   sync.RWMutex
	Deleted  bool
	Data     *world.Trigger
	Matcher  world.Matcher
	ByUser   bool
	RawMatch string
}

func (t *Trigger) Option(name string) (string, bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	switch name {
	case "clipboard_arg":
		return "0", true
	case "colour_change_type":
		return strconv.Itoa(t.Data.ColourChangeType), true
	case "custom_colour":
		return strconv.Itoa(t.Data.Colour), true
	case "enabled":
		return world.ToStringBool(t.Data.Enabled), true
	case "expand_variables":
		return world.ToStringBool(t.Data.ExpandVariables), true
	case "group":
		return t.Data.Group, true
	case "ignore_case":
		return world.ToStringBool(t.Data.IgnoreCase), true
	case "inverse":
		return world.ToStringBool(t.Data.Inverse), true
	case "italic":
		return world.ToStringBool(t.Data.Italic), true
	case "keep_evaluating":
		return world.ToStringBool(t.Data.KeepEvaluating), true
	case "lines_to_match":
		return strconv.Itoa(t.Data.LinesToMatch), true
	case "lowercase_wildcard":
		return world.ToStringBool(t.Data.WildcardLowerCase), true
	case "match":
		return t.Data.Match, true
	case "match_style":
		return "0", true
	case "multi_line":
		return world.ToStringBool(t.Data.MultiLine), true
	case "name":
		return t.Data.Name, true
	case "new_style":
		return "0", true
	case "omit_from_log":
		return world.ToStringBool(t.Data.OmitFromLog), true
	case "omit_from_output":
		return world.ToStringBool(t.Data.OmitFromOutput), true
	case "one_shot":
		return world.ToStringBool(t.Data.OneShot), true
	case "other_back_colour":
		return "0", true
	case "other_text_colour":
		return "0", true
	case "regexp":
		return world.ToStringBool(t.Data.Regexp), true
	case "repeat":
		return world.ToStringBool(t.Data.Repeat), true
	case "script":
		return t.Data.Script, true
	case "send":
		return t.Data.Send, true
	case "send_to":
		return strconv.Itoa(t.Data.SendTo), true
	case "sequence":
		return strconv.Itoa(t.Data.Sequence), true
	case "user":
		return "0", true
	case "variable":
		return t.Data.Variable, true
	}

	return "", false
}
func (t *Trigger) Info(infotype int) (string, bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	switch infotype {
	case 1:
		return t.Data.Match, true
	case 2:
		return t.Data.Send, true
	case 3:
		return t.Data.SoundFileName, true
	case 4:
		return t.Data.Script, true
	case 5:
		return world.ToStringBool(t.Data.OmitFromLog), true
	case 6:
		return world.ToStringBool(t.Data.OmitFromOutput), true
	case 7:
		return world.ToStringBool(t.Data.KeepEvaluating), true
	case 8:
		return world.ToStringBool(t.Data.Enabled), true
	case 9:
		return world.ToStringBool(t.Data.Regexp), true
	case 10:
		return world.ToStringBool(t.Data.IgnoreCase), true
	case 11:
		return world.ToStringBool(t.Data.Repeat), true
	case 13:
		return world.ToStringBool(t.Data.ExpandVariables), true
	case 15:
		return strconv.Itoa(t.Data.SendTo), true
	case 16:
		return strconv.Itoa(t.Data.Sequence), true
	case 23:
		return world.ToStringBool(t.Data.Temporary), true
	case 25:
		return world.ToStringBool(t.Data.WildcardLowerCase), true
	case 26:
		return t.Data.Group, true
	case 27:
		return t.Data.Variable, true
	case 28:
		return strconv.Itoa(0), true
	case 36:
		return world.ToStringBool(t.Data.OneShot), true
	}
	return "", false
}
func (t *Trigger) SetOption(name string, val string) (bool, bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	switch name {
	case "clipboard_arg":
		return true, true
	case "colour_change_type":
		return true, true
	case "custom_colour":
		return true, true
	case "enabled":
		t.Data.Enabled = world.FromStringBool(val)
		return true, true
	case "expand_variables":
		t.Data.ExpandVariables = world.FromStringBool(val)
		return true, true
	case "group":
		t.Data.Group = val
		return true, true
	case "ignore_case":
		t.Data.IgnoreCase = world.FromStringBool(val)
		return true, true
	case "inverse":
		t.Data.Inverse = world.FromStringBool(val)
		return true, true
	case "italic":
		t.Data.Italic = world.FromStringBool(val)
		return true, true
	case "lines_to_match":
		t.Data.LinesToMatch = world.FromStringInt(val)
		return true, true
	case "lowercase_wildcard":
		t.Data.WildcardLowerCase = world.FromStringBool(val)
		return true, true
	case "match":
		t.Data.Match = val
		return true, true
	case "match_style":
		return true, true
	case "multi_line":
		t.Data.MultiLine = world.FromStringBool(val)
		return true, true
	case "name":
		t.Data.Name = val
		return true, true
	case "new_style":
		return true, true
	case "omit_from_log":
		t.Data.OmitFromLog = world.FromStringBool(val)
		return true, true
	case "omit_from_output":
		t.Data.OmitFromOutput = world.FromStringBool(val)
		return true, true
	case "one_shot":
		t.Data.OneShot = world.FromStringBool(val)
		return true, true
	case "other_back_colour":
		return true, true
	case "other_text_colour":
		return true, true
	case "send":
		t.Data.Send = val
		return true, true
	case "regexp":
		t.Data.Regexp = world.FromStringBool(val)
		return true, true
	case "repeat":
		t.Data.Repeat = world.FromStringBool(val)
		return true, true
	case "script":
		t.Data.Script = val
		return true, true
	case "send_to":
		t.Data.SendTo = world.FromStringInt(val)
		return true, true
	case "sequence":
		t.Data.Sequence = world.FromStringInt(val)
	case "sound":
		return true, true
	case "sound_if_inactive":
		return true, true
	case "user":
		return false, false
	case "variable":
		t.Data.Variable = val
		return true, true
	}

	return false, false
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
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Deleted || !t.Data.Enabled {
		return nil, nil
	}
	if t.Data.ExpandVariables {
		if ctx.Expanded == nil {
			params := ctx.Bus.GetParams()
			r := make([]string, 0, len(params)*4)
			for k := range params {
				r = append(r, "@"+k, params[k], "@!"+k, regexp.QuoteMeta(params[k]))
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
	var line string
	if t.Data.MultiLine && t.Data.Regexp {
		line = strings.Join(ctx.Bus.DoMultiLinesLast(t.Data.LinesToMatch), "\n")
	} else {
		line = ctx.Line.Plain()
	}
	return t.Matcher.Match(line)
}

type TriggerQueue []*Trigger

// Len is the number of elements in the collection.
func (q TriggerQueue) Len() int {
	return len(q)
}

// Less reports whether the element with index i
func (q TriggerQueue) Less(i, j int) bool {
	if q[i].Deleted != q[j].Deleted {
		return q[j].Deleted
	}
	return q[i].Data.Sequence < q[j].Data.Sequence
}

// Swap swaps the elements with indexes i and j.
func (q TriggerQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
