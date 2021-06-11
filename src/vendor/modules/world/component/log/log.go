package log

import (
	"modules/world/bus"

	"github.com/herb-go/util"
)

type Log struct {
}

func (l *Log) InstallTo(b *bus.Bus) {
	dolog := b.WrapHandleError(l.DoLogError)
	b.HandleConverterError = dolog
	b.HandleCmdError = dolog
	b.HandleConnError = dolog
}
func (l *Log) DoLogError(b *bus.Bus, err error) {
	util.LogError(err)
}

func (l *Log) UninstallFrom(b *bus.Bus) {

}
