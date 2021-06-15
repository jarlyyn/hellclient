package component

import "modules/world/bus"

type Component interface {
	InstallTo(b *bus.Bus)
}

func InstallComponents(b *bus.Bus, c ...Component) {
	for _, v := range c {
		v.InstallTo(b)
	}
}
