package log

import (
	"modules/world/bus"

	"github.com/herb-go/util"
)

type Log struct {
}

func (l *Log) InstallTo(b *bus.Bus) {
	b.HandleConverterError = l.DoLogError
	b.HandleCmdError = l.DoLogError
	b.HandleConnError = l.DoLogError
}
func (l *Log) DoLogError(b *bus.Bus, err error) {
	util.LogError(err)
}

func (l *Log) UninstallFrom(b *bus.Bus) {

}
