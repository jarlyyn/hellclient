package automation

import (
	"modules/world"
	"modules/world/bus"
	"strings"
	"sync"
)

type Automation struct {
	Timers                 *Timers
	Aliases                *Aliases
	Triggers               *Triggers
	MultiLines             *MultiLines
	Locker                 sync.RWMutex
	evaluatingTriggersStop bool
}

func (a *Automation) InstallTo(b *bus.Bus) {
	a.Timers = NewTimers()
	a.Aliases = NewAliases()
	a.Triggers = NewTriggers()
	a.MultiLines = NewMultiLines()
	a.Timers.OnFire = b.WrapHandleTimer(a.OnFire)
	b.AddTimer = a.AddTimer
	b.DoDeleteTimer = a.RemoveTimer
	b.DoDeleteTimerByName = a.RemoveTimerByName
	b.DoDeleteTemporaryTimers = a.DeleteTemporaryTimers
	b.DoDeleteTimerGroup = a.DeleteTimerGroup
	b.DoEnableTimerByName = a.EnableTimerByName
	b.DoEnableTimerGroup = a.EnableTimerGroup
	b.DoListTimerNames = a.ListTimerNames
	b.HasNamedTimer = a.HasNamedTimer
	b.DoResetNamedTimer = a.ResetNamedTimer
	b.DoResetTimers = a.ResetTimers
	b.GetTimerOption = a.GetTimerOption
	b.GetTimerInfo = a.GetTimerInfo
	b.SetTimerOption = a.SetTimerOption
	b.GetTimersByType = a.GetTimersByType
	b.AddTimers = a.AddTimers
	b.DoDeleteTimerByType = a.DoDeleteTimerByType
	b.GetTimer = a.GetTimer
	b.DoUpdateTimer = a.DoUpdateTimer

	b.DoDeleteAlias = a.DoDeleteAlias
	b.DoDeleteAliasByName = a.DoDeleteAliasByName
	b.DoDeleteTemporaryAliases = a.DoDeleteTemporaryAliases
	b.DoDeleteAliasGroup = a.DoDeleteAliasGroup
	b.DoEnableAliasByName = a.DoEnableAliasByName
	b.DoEnableAliasGroup = a.DoEnableAliasGroup
	b.GetAlias = a.GetAlias
	b.GetAliasesByType = a.GetAliasesByType
	b.DoDeleteAliasByType = a.DoDeleteAliasByType
	b.AddAliases = a.AddAliases
	b.GetAliasOption = a.GetAliasOption
	b.GetAliasInfo = a.GetAliasInfo
	b.SetAliasOption = a.SetAliasOption
	b.HasNamedAlias = a.HasNamedAlias
	b.DoListAliasNames = a.DoListAliasNames
	b.AddAlias = a.AddAlias
	b.DoUpdateAlias = a.DoUpdateAlias

	b.DoDeleteAlias = a.DoDeleteAlias
	b.DoDeleteAliasByName = a.DoDeleteAliasByName
	b.DoDeleteTemporaryAliases = a.DoDeleteTemporaryAliases
	b.DoDeleteAliasGroup = a.DoDeleteAliasGroup
	b.DoEnableAliasByName = a.DoEnableAliasByName
	b.DoEnableAliasGroup = a.DoEnableAliasGroup
	b.GetAlias = a.GetAlias
	b.GetAliasesByType = a.GetAliasesByType
	b.DoDeleteAliasByType = a.DoDeleteAliasByType
	b.AddAliases = a.AddAliases
	b.GetAliasOption = a.GetAliasOption
	b.SetAliasOption = a.SetAliasOption
	b.HasNamedAlias = a.HasNamedAlias
	b.DoListAliasNames = a.DoListAliasNames
	b.AddAlias = a.AddAlias
	b.DoUpdateAlias = a.DoUpdateAlias

	b.DoDeleteTrigger = a.DoDeleteTrigger
	b.DoDeleteTriggerByName = a.DoDeleteTriggerByName
	b.DoDeleteTemporaryTriggers = a.DoDeleteTemporaryTriggers
	b.DoDeleteTriggerGroup = a.DoDeleteTriggerGroup
	b.DoEnableTriggerByName = a.DoEnableTriggerByName
	b.DoEnableTriggerGroup = a.DoEnableTriggerGroup
	b.GetTrigger = a.GetTrigger
	b.GetTriggersByType = a.GetTriggersByType
	b.DoDeleteTriggerByType = a.DoDeleteTriggerByType
	b.AddTriggers = a.AddTriggers
	b.GetTriggerOption = a.GetTriggerOption
	b.GetTriggerInfo = a.GetTriggerInfo
	b.SetTriggerOption = a.SetTriggerOption
	b.HasNamedTrigger = a.HasNamedTrigger
	b.DoListTriggerNames = a.DoListTriggerNames
	b.AddTrigger = a.AddTrigger
	b.DoUpdateTrigger = a.DoUpdateTrigger
	b.DoGetTriggerWildcard = a.GetTriggerWildcard

	b.DoMultiLinesAppend = a.DoMultiLinesAppend
	b.DoMultiLinesFlush = a.DoMultiLinesFlush
	b.DoMultiLinesLast = a.DoMultiLinesLast

	b.DoExecute = b.WrapHandleString(a.DoExecute)

	b.BindLineEvent(b, a.OnLine)
	b.BindBeforeCloseEvent(a, a.OnClose)
}
func (a *Automation) OnClose(b *bus.Bus) {
	a.Timers.Flush()
}
func (a *Automation) DoMultiLinesAppend(message string) {
	a.MultiLines.Append(message)
}
func (a *Automation) DoMultiLinesFlush() {
	a.MultiLines.Flush()

}
func (a *Automation) DoMultiLinesLast(count int) []string {
	return a.MultiLines.Last(count)
}

func (a *Automation) DoStopEvaluatingTriggers() {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	a.evaluatingTriggersStop = true
}
func (a *Automation) EvaluatingTriggersStop() bool {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	return a.evaluatingTriggersStop
}
func (a *Automation) ReadyForLine() {
	a.Locker.Lock()
	defer a.Locker.Unlock()
	a.evaluatingTriggersStop = false
}
func (a *Automation) GetTriggerWildcard(name string) *world.MatchResult {
	a.Triggers.Locker.Lock()
	t := a.Triggers.Named[name]
	a.Triggers.Locker.Unlock()
	if t == nil {
		return nil
	}
	return t.Wildcards()
}
func (a *Automation) OnLine(b *bus.Bus, line *world.Line) {
	if line == nil || line.Type != world.LineTypeReal {
		return
	}
	a.ReadyForLine()
	b.DoMultiLinesAppend(line.Plain())
	queue := a.Triggers.Queue()
	ctx := &Context{
		Line: line,
		Bus:  b,
	}
	for _, v := range queue {
		r, err := v.Match(ctx)
		if err != nil {
			b.HandleTriggerError(err)
			continue
		}
		if r == nil {
			continue
		}
		a.Locker.Lock()
		a.Locker.Unlock()
		a.Triggers.Locker.Lock()
		rawtrigger := a.Triggers.All[v.Data.ID]
		if rawtrigger != nil {
			rawtrigger.Locker.Lock()
			rawtrigger.wildcards = r
			rawtrigger.Locker.Unlock()
		}
		a.Triggers.Locker.Unlock()
		var send string
		var data world.Trigger
		v.Locker.Lock()
		data = *v.Data
		if data.Script != "" {
			line.Triggers = append(line.Triggers, data.Script)
		} else {
			line.Triggers = append(line.Triggers, "#"+data.ID)
		}
		if v.Data.Send != "" {
			rl := r.ReplaceList(v.Data.Name)
			if v.Data.ExpandVariables {
				rl = append(rl, BuildParamsReplacer(b)...)
			}
			send = strings.NewReplacer(rl...).Replace(data.Send)
		}
		v.Locker.Unlock()
		if data.OneShot {
			a.Triggers.RemoveTrigger(data.ID)
		}
		if data.OmitFromOutput {
			b.DoOmitOutput()
		}
		if data.OmitFromLog {
			line.OmitFromLog = true
		}
		if send != "" {
			go a.trySendTo(b, data.SendTo, send, data.Variable, data.OmitFromLog, data.OmitFromOutput)
		}
		if data.Script != "" {
			b.DoSendTriggerToScript(line, &data, r)
		}
		if !data.KeepEvaluating || a.EvaluatingTriggersStop() {
			return
		}
	}
}

func (a *Automation) MatchAlias(b *bus.Bus, message string) bool {
	var matched bool
	queue := a.Aliases.Queue()
	for _, v := range queue {
		r, err := v.Match(message)
		if err != nil {
			b.HandleTriggerError(err)
			continue
		}
		if r == nil {
			continue
		}
		var send string
		var data world.Alias
		v.Locker.Lock()
		data = *v.Data
		if v.Data.Send != "" {
			rl := r.ReplaceList(v.Data.Name)
			if v.Data.ExpandVariables {
				rl = append(rl, BuildParamsReplacer(b)...)
			}
			send = strings.NewReplacer(rl...).Replace(data.Send)
		}
		v.Locker.Unlock()
		if send != "" {
			a.trySendTo(b, data.SendTo, send, data.Variable, data.OmitFromLog, data.OmitFromOutput)
		}
		if data.Script != "" {
			b.DoSendAliasToScript(message, &data, r)
		}
		if data.OneShot {
			a.Aliases.RemoveAlias(data.ID)
		}
		if !data.KeepEvaluating {
			return true
		}
	}
	return matched
}

func (a *Automation) DoExecute(b *bus.Bus, message string) {
	if message == "" {
		return
	}
	replacers := []string{"\\\\", "\\"}
	sep := b.GetCommandStackCharacter()
	if sep != "" {
		replacers = append(replacers, "\\"+sep, sep, sep, "\n")
	}
	m := strings.NewReplacer(replacers...).Replace(message)
	cmds := strings.Split(m, "\n")
	for _, v := range cmds {
		a.executecmd(b, v)
	}
}
func (a *Automation) executecmd(b *bus.Bus, cmd string) {
	p := b.GetScriptPrefix()
	if p != "" && strings.HasPrefix(cmd, p) {
		b.DoRunScript(strings.TrimPrefix(cmd, p))
		return
	}
	if !a.MatchAlias(b, cmd) {
		cmd := world.CreateCommand(cmd)
		cmd.History = true
		b.DoSend(cmd)
	}

}
func (a *Automation) AddTimer(timer *world.Timer, replace bool) bool {
	return a.Timers.AddTimer(timer, replace)
}
func (a *Automation) AddTimers(ts []*world.Timer) {
	a.Timers.AddTimers(ts)
}
func (a *Automation) RemoveTimer(id string) bool {
	return a.Timers.RemoveTimer(id)
}
func (a *Automation) RemoveTimerByName(name string) bool {
	return a.Timers.RemoveTimerByName(name)
}
func (a *Automation) DeleteTemporaryTimers() int {
	return a.Timers.DeleteTemporaryTimers()
}
func (a *Automation) DeleteTimerGroup(group string) int {
	return a.Timers.DeleteTimerGroup(group)
}
func (a *Automation) EnableTimerByName(name string, enabled bool) bool {
	return a.Timers.EnableTimerByName(name, enabled)
}
func (a *Automation) EnableTimerGroup(group string, enabled bool) int {
	return a.Timers.EnableTimerGroup(group, enabled)
}
func (a *Automation) ListTimerNames() []string {
	return a.Timers.ListTimerNames()
}
func (a *Automation) HasNamedTimer(name string) bool {
	return a.Timers.HasNamedTimer(name)
}
func (a *Automation) ResetNamedTimer(name string) bool {
	return a.Timers.ResetNamedTimer(name)
}
func (a *Automation) ResetTimers() {
	a.Timers.ResetTimers()
}
func (a *Automation) GetTimerOption(name string, option string) (string, bool, bool) {
	return a.Timers.GetTimerOption(name, option)
}
func (a *Automation) GetTimerInfo(name string, infotype int) (string, bool, bool) {
	return a.Timers.GetTimerInfo(name, infotype)
}
func (a *Automation) SetTimerOption(name string, option string, value string) (bool, bool, bool) {
	return a.Timers.SetTimerOption(name, option, value)
}
func (a *Automation) GetTimersByType(byuser bool) []*world.Timer {
	return a.Timers.GetTimersByType(byuser)
}
func (a *Automation) GetTimer(id string) *world.Timer {
	return a.Timers.GetTimer(id)
}
func (a *Automation) OnFire(b *bus.Bus, timer *world.Timer) {
	connceted := b.GetConnConnected()
	if !connceted && !timer.ActionWhenDisconnectd {
		return
	}
	a.trySendTo(b, timer.SendTo, timer.Send, timer.Variable, timer.OmitFromLog, timer.OmitFromOutput)
	if timer.Script != "" {
		ti := *timer
		go b.DoSendTimerToScript(&ti)
	}
}
func (a *Automation) DoDeleteTimerByType(byuser bool) {
	a.Timers.DoDeleteTimerByType(byuser)
}
func (a *Automation) DoUpdateTimer(ti *world.Timer) int {
	return a.Timers.DoUpdateTimer(ti)
}

func (a *Automation) DoDeleteAlias(id string) bool {
	return a.Aliases.RemoveAlias(id)
}
func (a *Automation) DoDeleteAliasByName(name string) bool {
	return a.Aliases.DoDeleteAliasByName(name)
}
func (a *Automation) DoDeleteTemporaryAliases() int {
	return a.Aliases.DoDeleteTemporaryAliases()
}
func (a *Automation) DoDeleteAliasGroup(group string) int {
	return a.Aliases.DoDeleteAliasGroup(group)
}
func (a *Automation) DoEnableAliasByName(name string, enabled bool) bool {
	return a.Aliases.DoEnableAliasByName(name, enabled)
}
func (a *Automation) DoEnableAliasGroup(group string, enabled bool) int {
	return a.Aliases.DoEnableAliasGroup(group, enabled)
}
func (a *Automation) GetAlias(id string) *world.Alias {
	return a.Aliases.GetAlias(id)
}
func (a *Automation) GetAliasesByType(byuser bool) []*world.Alias {
	return a.Aliases.GetAliasesByType(byuser)
}
func (a *Automation) DoDeleteAliasByType(byuser bool) {
	a.Aliases.DoDeleteAliasByType(byuser)
}
func (a *Automation) AddAliases(al []*world.Alias) {
	a.Aliases.AddAliases(al)
}
func (a *Automation) GetAliasOption(name string, option string) (string, bool, bool) {
	return a.Aliases.GetAliasOption(name, option)
}
func (a *Automation) GetAliasInfo(name string, infotype int) (string, bool, bool) {
	return a.Aliases.GetAliasInfo(name, infotype)
}
func (a *Automation) SetAliasOption(name string, option string, value string) (bool, bool, bool) {
	return a.Aliases.SetAliasOption(name, option, value)
}
func (a *Automation) HasNamedAlias(name string) bool {
	return a.Aliases.HasNamedAlias(name)
}
func (a *Automation) DoListAliasNames() []string {
	return a.Aliases.DoListAliasNames()
}
func (a *Automation) AddAlias(al *world.Alias, replace bool) bool {
	return a.Aliases.AddAlias(al, replace)
}
func (a *Automation) DoUpdateAlias(al *world.Alias) int {
	return a.Aliases.DoUpdateAlias(al)
}

func (a *Automation) DoDeleteTrigger(id string) bool {
	return a.Triggers.RemoveTrigger(id)
}
func (a *Automation) DoDeleteTriggerByName(name string) bool {
	return a.Triggers.DoDeleteTriggerByName(name)
}
func (a *Automation) DoDeleteTemporaryTriggers() int {
	return a.Triggers.DoDeleteTemporaryTriggers()
}
func (a *Automation) DoDeleteTriggerGroup(group string) int {
	return a.Triggers.DoDeleteTriggerGroup(group)
}
func (a *Automation) DoEnableTriggerByName(name string, enabled bool) bool {
	return a.Triggers.DoEnableTriggerByName(name, enabled)
}
func (a *Automation) DoEnableTriggerGroup(group string, enabled bool) int {
	return a.Triggers.DoEnableTriggerGroup(group, enabled)
}
func (a *Automation) GetTrigger(id string) *world.Trigger {
	return a.Triggers.GetTrigger(id)
}
func (a *Automation) GetTriggersByType(byuser bool) []*world.Trigger {
	return a.Triggers.GetTriggersByType(byuser)
}
func (a *Automation) DoDeleteTriggerByType(byuser bool) {
	a.Triggers.DoDeleteTriggerByType(byuser)
}
func (a *Automation) AddTriggers(al []*world.Trigger) {
	a.Triggers.AddTriggers(al)
}
func (a *Automation) GetTriggerOption(name string, option string) (string, bool, bool) {
	return a.Triggers.GetTriggerOption(name, option)
}
func (a *Automation) GetTriggerInfo(name string, infotype int) (string, bool, bool) {
	return a.Triggers.GetTriggerInfo(name, infotype)
}
func (a *Automation) SetTriggerOption(name string, option string, value string) (bool, bool, bool) {
	return a.Triggers.SetTriggerOption(name, option, value)
}
func (a *Automation) HasNamedTrigger(name string) bool {
	return a.Triggers.HasNamedTrigger(name)
}
func (a *Automation) DoListTriggerNames() []string {
	return a.Triggers.DoListTriggerNames()
}
func (a *Automation) AddTrigger(al *world.Trigger, replace bool) bool {
	return a.Triggers.AddTrigger(al, replace)
}
func (a *Automation) DoUpdateTrigger(al *world.Trigger) int {
	return a.Triggers.DoUpdateTrigger(al)
}
func (a *Automation) trySendTo(b *bus.Bus, target int, message string, variable string, omit_from_log bool, omit_from_output bool) bool {
	if message == "" {
		return false
	}
	switch target {
	case world.SendtoWorld:
		cmd := world.CreateCommand(message)
		if omit_from_output {
			cmd.Echo = false
		}
		if omit_from_log {
			cmd.Log = false
		}
		b.DoMetronomeSend(cmd)
	case world.SendtoCommand:
	case world.SendtoOutput:
		b.DoPrint(message)
	case world.SendtoStatus:
		b.SetStatus(message)
	case world.SendtoNotepad:
	case world.SendtoNotepadAppend:
	case world.SendtoLogfile:
	case world.SendtoNotepadReplace:
	case world.SendtoCommandqueue:
		cmd := world.CreateCommand(message)
		if omit_from_output {
			cmd.Echo = false
		}
		if omit_from_log {
			cmd.Log = false
		}
		b.DoSendToQueue(cmd)
	case world.SendtoVariable:
		b.SetParam(variable, message)
	case world.SendtoExecute:
		b.DoExecute(message)
	case world.SendtoSpeedwalk:
		b.DoExecute(message)
	case world.SendtoScript:
		b.DoRunScript(message)
	case world.SendtoImmediate:
		cmd := world.CreateCommand(message)
		if omit_from_output {
			cmd.Echo = false
		}
		if omit_from_log {
			cmd.Log = false
		}
		b.DoMetronomeSend(cmd)
	case world.SendtoScriptAfterOmit:
		b.DoRunScript(message)
	default:
		return false
	}
	return true
}

func New() *Automation {
	return &Automation{}
}
