package script

import (
	"hellclient/modules/world"
	"hellclient/modules/world/bus"
)

type Engine interface {
	Open(*bus.Bus) error
	Close(*bus.Bus)
	OnConnect(*bus.Bus)
	OnDisconnect(*bus.Bus)
	OnTrigger(b *bus.Bus, line *world.Line, trigger *world.Trigger, result *world.MatchResult)
	OnAlias(b *bus.Bus, message string, alias *world.Alias, result *world.MatchResult)
	OnTimer(b *bus.Bus, timer *world.Timer)
	OnCallback(b *bus.Bus, cb *world.Callback)
	OnBroadCast(b *bus.Bus, bc *world.Broadcast)
	OnHUDClick(b *bus.Bus, c *world.Click)
	OnResponse(b *bus.Bus, msg *world.Message)
	OnAssist(b *bus.Bus, script string)
	OnBuffer(b *bus.Bus, data []byte) bool
	OnFocus(b *bus.Bus)
	OnSubneg(b *bus.Bus, code byte, data []byte) bool
	Run(*bus.Bus, string)
}
