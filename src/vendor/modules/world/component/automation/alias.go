package automation

import (
	"modules/world"
	"strconv"
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

func (a *Alias) Option(name string) (string, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	switch name {
	case "echo_alias":
		return world.ToStringBool(!a.Data.OmitFromOutput), true
	case "enabled":
		return world.ToStringBool(a.Data.Enabled), true
	case "expand_variables":
		return world.ToStringBool(a.Data.ExpandVariables), true
	case "group":
		return a.Data.Group, true
	case "ignore_case":
		return world.ToStringBool(a.Data.IgnoreCase), true
	case "keep_evaluating":
		return world.ToStringBool(a.Data.KeepEvaluating), true
	case "match":
		return a.Data.Match, true
	case "menu":
		return world.ToStringBool(a.Data.Menu), true
	case "name":
		return a.Data.Name, true
	case "offset_hour":
		return "0", true
	case "offset_minute":
		return "0", true
	case "offset_second":
		return "0", true
	case "omit_from_command_history":
		return world.ToStringBool(a.Data.OmitFromCommandHistory), true
	case "omit_from_log":
		return world.ToStringBool(a.Data.OmitFromLog), true
	case "omit_from_output":
		return world.ToStringBool(a.Data.OmitFromOutput), true
	case "one_shot":
		return world.ToStringBool(a.Data.OneShot), true
	case "regexp":
		return world.ToStringBool(a.Data.Regexp), true
	case "script":
		return a.Data.Script, true
	case "send":
		return a.Data.Send, true
	case "send_to":
		return strconv.Itoa(a.Data.SendTo), true
	case "sequence":
		return strconv.Itoa(a.Data.Sequence), true
	case "user":
		return "0", true
	case "variable":
		return a.Data.Variable, true
	}
	return "", false
}
func (a *Alias) Info(infotype int) (string, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	switch infotype {
	case 1:
		return a.Data.Match, true
	case 2:
		return a.Data.Send, true
	case 3:
		return a.Data.Script, true
	case 4:
		return a.Data.Send, true
	case 5:
		return a.Data.Script, true
	case 6:
		return world.ToStringBool(a.Data.Enabled), true
	case 7:
		return world.ToStringBool(a.Data.Regexp), true
	case 8:
		return world.ToStringBool(a.Data.IgnoreCase), true
	case 9:
		return world.ToStringBool(a.Data.ExpandVariables), true
	case 14:
		return world.ToStringBool(a.Data.Temporary), true
	case 16:
		return a.Data.Group, true
	case 17:
		return a.Data.Variable, true
	case 18:
		return strconv.Itoa(a.Data.SendTo), true
	case 19:
		return world.ToStringBool(a.Data.KeepEvaluating), true
	case 20:
		return strconv.Itoa(a.Data.Sequence), true
	case 22:
		return world.ToStringBool(a.Data.OmitFromCommandHistory), true
	case 23:
		return strconv.Itoa(0), true
	case 29:
		return world.ToStringBool(a.Data.OneShot), true
	}
	return "", false
}
func (a *Alias) SetOption(name string, val string) (bool, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	switch name {
	case "echo_alias":
		a.Data.OmitFromOutput = !world.FromStringBool(val)
		return true, true
	case "enabled":
		a.Data.Enabled = world.FromStringBool(val)
		return true, true
	case "expand_variables":
		a.Data.ExpandVariables = world.FromStringBool(val)
		return true, true
	case "ignore_case":
		a.Data.IgnoreCase = world.FromStringBool(val)
		return true, true
	case "keep_evaluating":
		a.Data.KeepEvaluating = world.FromStringBool(val)
		return true, true
	case "match":
		a.Data.Match = val
		return true, true
	case "offset_hour":
		return false, false
	case "offset_minute":
		return false, false
	case "offset_second":
		return false, false
	case "menu":
		a.Data.Menu = world.FromStringBool(val)
		return true, true
	case "omit_from_command_history":
		a.Data.OmitFromCommandHistory = world.FromStringBool(val)
		return true, true
	case "omit_from_log":
		a.Data.OmitFromLog = world.FromStringBool(val)
		return true, true
	case "omit_from_output":
		a.Data.OmitFromOutput = world.FromStringBool(val)
		return true, true
	case "one_shot":
		a.Data.OneShot = world.FromStringBool(val)
		return true, true
	case "regexp":
		a.Data.Regexp = world.FromStringBool(val)
		return true, true

	case "script":
		a.Data.Script = val
		return true, true
	case "send":
		a.Data.Send = val
		return true, true
	case "send_to":
		a.Data.SendTo = world.FromStringInt(val)
		return true, true
	case "user":
		return false, false
	case "sequence":
		a.Data.Sequence = world.FromStringInt(val)
	case "variable":
		a.Data.Variable = val
		return true, true
	}

	return false, false
}
