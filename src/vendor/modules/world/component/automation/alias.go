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
		if !a.Data.OmitFromOutput {
			return world.StringYes, true
		}
		return "", true
	case "enabled":
		if a.Data.Enabled {
			return world.StringYes, true
		}
		return "", true
	case "expand_variables":
		if a.Data.ExpandVariables {
			return world.StringYes, true
		}
		return "", true
	case "group":
		return a.Data.Group, true
	case "ignore_case":
		if a.Data.IgnoreCase {
			return world.StringYes, true
		}
		return "", true
	case "keep_evaluating":
		if a.Data.KeepEvaluating {
			return world.StringYes, true
		}
		return "", true
	case "match":
		return a.Data.Match, true
	case "menu":
		if a.Data.Menu {
			return world.StringYes, true
		}
		return "", true
	case "name":
		return a.Data.Name, true
	case "offset_hour":
		return "0", true
	case "offset_minute":
		return "0", true
	case "offset_second":
		return "0", true
	case "omit_from_command_history":
		if a.Data.OmitFromCommandHistory {
			return world.StringYes, true
		}
		return "", true

	case "omit_from_log":
		if a.Data.OmitFromLog {
			return world.StringYes, true
		}
		return "", true
	case "omit_from_output":
		if a.Data.OmitFromOutput {
			return world.StringYes, true
		}
		return "", true
	case "one_shot":
		if a.Data.OneShot {
			return world.StringYes, true
		}
		return "", true
	case "regexp":
		if a.Data.Regexp {
			return world.StringYes, true
		}
		return "", true

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

func (a *Alias) SetOption(name string, val string) (bool, bool) {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	switch name {
	case "echo_alias":
		a.Data.OmitFromOutput = !(val == world.StringYes)
		return true, true
	case "enabled":
		a.Data.Enabled = (val == world.StringYes)
		return true, true
	case "expand_variables":
		a.Data.ExpandVariables = (val == world.StringYes)
		return true, true
	case "ignore_case":
		a.Data.IgnoreCase = (val == world.StringYes)
		return true, true
	case "keep_evaluating":
		a.Data.KeepEvaluating = (val == world.StringYes)
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
		a.Data.Menu = (val == world.StringYes)
		return true, true
	case "omit_from_command_history":
		a.Data.OmitFromCommandHistory = (val == world.StringYes)
		return true, true
	case "omit_from_log":
		a.Data.OmitFromLog = (val == world.StringYes)
		return true, true
	case "omit_from_output":
		a.Data.OmitFromOutput = (val == world.StringYes)
		return true, true
	case "one_shot":
		a.Data.OneShot = (val == world.StringYes)
		return true, true
	case "regexp":
		a.Data.Regexp = (val == world.StringYes)
		return true, true

	case "script":
		a.Data.Script = val
		return true, true
	case "send":
		a.Data.Send = val
		return true, true
	case "send_to":
		i, _ := strconv.Atoi(val)
		a.Data.SendTo = i
		return true, true
	case "user":
		return false, false
	case "sequence":
		seq, err := strconv.Atoi(val)
		if err != nil {
			seq = 0
		}
		a.Data.Sequence = seq
	case "variable":
		a.Data.Variable = val
		return true, true
	}

	return false, false
}
