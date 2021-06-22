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
	b.HandleScriptError = dolog
	b.HandleTriggerError = dolog
}
func (l *Log) DoLogError(b *bus.Bus, err error) {
	if err == nil {
		return
	}
	go b.DoPrintSystem("Error: " + err.Error())
	go util.LogError(err)
}

func New() *Log {
	return &Log{}
}
