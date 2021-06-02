package app

import (
	"sync/atomic"

	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/loggerconfig"
	"github.com/herb-go/util/config/tomlconfig"
)

//Loggers app loggers settings.
var Loggers = &loggerconfig.Config{}

var syncLoggers atomic.Value

//StoreDevelopment atomically store loggers config
func (a *appSync) StoreLoggers(c *loggerconfig.Config) {
	syncLoggers.Store(c)
}

//LoadDevelopment atomically load loggers config
func (a *appSync) LoadLoggers() *loggerconfig.Config {
	v := syncLoggers.Load()
	if v == nil {
		return nil
	}
	return v.(*loggerconfig.Config)
}

func init() {
	config.RegisterLoader(util.ConfigFile("/loggers.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, Loggers))
		Sync.StoreLoggers(Loggers)
	})
}
