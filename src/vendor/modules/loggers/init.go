package loggers

import (
	"modules/app"
	"os"
	"os/signal"

	"github.com/herb-go/logger"

	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "100loggers"

//MyLogger my logger
// var MyLogger = logger.PrintLogger.SubLogger().SetPrefixs(logger.DefaultTimePrefix).SetID("mylogger")
// var MyFormatLogger = MyLogger.FormatLogger()

var reopenSignals []os.Signal

func reopenLoggers() {
	logger.ReopenBuiltinLoggers()
	// logger.Reopen(MyLogger)

}
func listenForReopenLoggers() {
	var signalChan = make(chan os.Signal)
	signal.Notify(signalChan, reopenSignals...)
	go func() {
		<-util.QuitChan()
		signal.Stop(signalChan)
	}()
	go func() {
		for {
			select {
			case <-signalChan:
				go reopenLoggers()
			case <-util.QuitChan():
				return
			}
		}
	}()
}

func init() {
	util.RegisterModule(ModuleName, func() {
		util.MakeLoggerFolderIfNotExist()
		util.Must(app.Loggers.ApplyToBuiltinLoggers())
		util.ErrorLogger = logger.Error
		// util.Must(app.Loggers.ApplyTo(MyLogger, "MyLogger"))
		// util.Must(app.Loggers.SetFormatter(MyFormatLogger, "MyLoggerFormat"))
		if app.Development.Debug {
			logger.EnableDevelopmengLoggers()
		}
		listenForReopenLoggers()
	})
}
