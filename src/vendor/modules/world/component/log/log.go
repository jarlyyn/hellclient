package log

import (
	"modules/world/bus"

	"github.com/herb-go/util"
)

type Log struct {
	bus *bus.Bus
}

func (l *Log) InstallTo(b *bus.Bus) {
	l.bus = b
	b.HandleConverterError = l.DoLogError
	b.HandleCmdError = l.DoLogError
	b.HandleConnError = l.DoLogError
}
func (l *Log) DoLogError(err error) {
	util.LogError(err)
}

func (l *Log) UninstallFrom(b *bus.Bus) {
	if l.bus != b {
		return
	}
}
