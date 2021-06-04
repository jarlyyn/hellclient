package config

import (
	"modules/world/bus"
)

type Config struct {
	bus     *bus.Bus
	Host    string
	Port    string
	Charset string
}

func (c *Config) HostPort() string {
	return c.Host + ":" + c.Port
}

func (c *Config) InstallTo(b *bus.Bus) {
	c.bus = b
	b.GetHostPort = c.HostPort
}

func (c *Config) UninstallFrom(b *bus.Bus) {
	if c.bus != b {
		return
	}
	b.GetHostPort = nil
}
