package automation

import (
	"modules/world"
	"modules/world/bus"
)

type Automation struct {
	Timers *Timers
}

func (a *Automation) InstallTo(b *bus.Bus) {
	a.Timers = NewTimers()
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
