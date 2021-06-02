package app

import (
	"sync/atomic"

	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/commonconfig"
	"github.com/herb-go/util/config/tomlconfig"
)

//Development app development settings.
var Development = &commonconfig.DevelopmentConfig{}

var syncDevelopment atomic.Value

//StoreDevelopment atomically store development config
func (a *appSync) StoreDevelopment(c *commonconfig.DevelopmentConfig) {
	syncDevelopment.Store(c)
}

//LoadDevelopment atomically load development config
func (a *appSync) LoadDevelopment() *commonconfig.DevelopmentConfig {
	v := syncDevelopment.Load()
	if v == nil {
		return nil
	}
	return v.(*commonconfig.DevelopmentConfig)
}

func init() {
	config.RegisterLoader(util.ConfigFile("/development.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, Development))
		Sync.StoreDevelopment(Development)
		if util.ForceDebug {
			Development.Debug = true
		}
		if Development.Debug {
			util.Debug = true
			util.SetWarning("development.toml", "Debug value is true.")
		}
	})
}
