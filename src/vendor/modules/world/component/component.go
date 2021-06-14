package component

import "modules/world/bus"

type Component interface {
	InstallTo(b *bus.Bus)
}
