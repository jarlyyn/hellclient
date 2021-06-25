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
}
func (a *Automation) AddTimer(timer *world.Timer, replace bool) {
	a.Timers.AddTimer(timer, replace)
}
func (a *Automation) OnFire(b *bus.Bus, timer *world.Timer) {
	connceted := b.GetConnConnected()
	if !connceted && !timer.ActionWhenDisconnectd {
		return
	}
	a.trySendTo(b, timer.SendTo, timer.Send, timer.Variable)
	if timer.Script != "" {
		b.DoSendTimerToScript(timer)
	}
}
func (a *Automation) trySendTo(b *bus.Bus, target int, message string, variable string) bool {
	if message == "" {
		return false
	}
	switch target {
	case world.SendtoWorld:
		b.DoSend(world.CreateCommand(message))
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
		b.DoSendToQueue(world.CreateCommand(message))
	case world.SendtoVariable:
		b.SetParam(variable, message)
	case world.SendtoExecute:
		b.DoExecute(message)
	case world.SendtoSpeedwalk:
		b.DoExecute(message)
	case world.SendtoScript:
		b.DoRunScript(message)
	case world.SendtoImmediate:
		b.DoSend(world.CreateCommand(message))
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
