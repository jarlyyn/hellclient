package log

import (
	"hellclient/modules/app"
	"hellclient/modules/world/bus"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/herb-go/logger"
	"github.com/herb-go/util"
)

type Log struct {
	Locker sync.RWMutex
}

func (l *Log) InstallTo(b *bus.Bus) {
	dolog := b.WrapHandleError(l.DoLogError)
	debug := b.WrapHandleError(l.DoDebugError)
	b.HandleConverterError = debug
	b.HandleCmdError = dolog
	b.HandleConnError = dolog
	b.HandleScriptError = dolog
	b.HandleTriggerError = dolog
	b.DoLog = b.WrapHandleString(l.Log)
}
func (l *Log) DoDebugError(b *bus.Bus, err error) {
	if err == nil {
		return
	}
	go func() {
		logger.Debug(err.Error())
	}()
}
func (l *Log) DoLogError(b *bus.Bus, err error) {
	if err == nil {
		return
	}
	go b.DoPrintSystem("Error: " + err.Error())
	go util.LogError(err)
	go b.DoLog(err.Error())
}
func (l *Log) Log(b *bus.Bus, message string) {
	l.Locker.Lock()
	defer l.Locker.Unlock()
	path := filepath.Join(b.GetLogsPath(), b.ID+".log")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, util.DefaultFileMode)
	if err != nil {
		util.LogError(err)
		return
	}
	defer f.Close()
	line := app.Time.Datetime(time.Now()) + " " + message + "\n"
	if _, err := f.Write([]byte(line)); err != nil {
		log.Fatal(err)
	}
}
func New() *Log {
	return &Log{}
}
