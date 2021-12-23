package app

import (
	"sync/atomic"
	_ "time/tzdata" //embed timezone data

	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/commonconfig"
	"github.com/herb-go/util/config/tomlconfig"
)

//Time app time settings
var Time = &commonconfig.TimeConfig{}

var syncTime atomic.Value

//StoreTime atomically store time config
func (a *appSync) StoreTime(c *commonconfig.TimeConfig) {
	syncTime.Store(c)
}

//LoadTime atomically load time config
func (a *appSync) LoadTime() *commonconfig.TimeConfig {
	v := syncTime.Load()
	if v == nil {
		return nil
	}
	return v.(*commonconfig.TimeConfig)
}

func init() {
	config.RegisterLoader(util.ConfigFile("/time.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, Time))
		Sync.StoreTime(Time)
	})
}
