package automation

import (
	"modules/world"
	"modules/world/bus"
	"strings"
)

type Automation struct {
	Timers  *Timers
	Aliases *Aliases
}

func (a *Automation) InstallTo(b *bus.Bus) {
	a.Timers = NewTimers()
	a.Aliases = NewAliases()
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
	b.SetAliasOption = a.SetAliasOption
	b.HasNamedAlias = a.HasNamedAlias
	b.DoListAliasNames = a.DoListAliasNames
	b.AddAlias = a.AddAlias
	b.DoUpdateAlias = a.DoUpdateAlias

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
		var target int
		var variable string
		var omitlog bool
		var omitoutput bool
		var keep bool
		v.Locker.Lock()

		if v.Data.Send != "" {
			target = v.Data.SendTo
			variable = v.Data.Variable
			omitlog = v.Data.OmitFromLog
			omitoutput = v.Data.OmitFromOutput
			keep = v.Data.KeepEvaluating
			rl := r.ReplaceList(v.Data.Name)
			if v.Data.ExpandVariables {
				rl = append(rl, BuildParamsReplacer(b)...)
			}
			send = strings.NewReplacer(rl...).Replace(v.Data.Send)
		}
		v.Locker.Unlock()
		if send != "" {
			a.trySendTo(b, target, send, variable, omitlog, omitoutput)
		}
		if !keep {
			return true
		}
	}
	return matched
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
		b.DoSendTimerToScript(&ti)
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
		b.DoSend(cmd)
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
		b.DoSend(cmd)
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
