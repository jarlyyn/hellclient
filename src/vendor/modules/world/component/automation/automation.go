package automation

import (
	"modules/world"
	"modules/world/bus"
)

type Automation struct {
}

func (a *Automation) trySendTo(b *bus.Bus, target int, message string, variable string) bool {
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
	case world.SendtoImmediate:
		b.DoSend(world.CreateCommand(message))
	default:
		return false
	}
	return true
}
