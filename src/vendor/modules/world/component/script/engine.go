package script

import (
	"modules/world/bus"
)

type Engine interface {
	Open(*bus.Bus) error
	Close(*bus.Bus)
	OnConnect(*bus.Bus)
	OnDisconnect(*bus.Bus)
	OnTrigger(*bus.Bus)
	OnAlias(*bus.Bus)
	OnTimer(*bus.Bus)
}
