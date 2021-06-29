package script

import (
	"modules/world"
	"modules/world/bus"
)

type Engine interface {
	Open(*bus.Bus) error
	Close(*bus.Bus)
	OnConnect(*bus.Bus)
	OnDisconnect(*bus.Bus)
	OnTrigger(*bus.Bus)
	OnAlias(b *bus.Bus, message string, alias *world.Alias, result *world.MatchResult)
	OnTimer(b *bus.Bus, timer *world.Timer)
	Run(*bus.Bus, string)
}
