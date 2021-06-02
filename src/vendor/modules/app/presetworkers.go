package app

import (
	"sync/atomic"

	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
	"github.com/herb-go/worker"
)

//PresetWorkers workers config
var PresetWorkers = &worker.Config{}

var syncPresetWorkers atomic.Value

//StoreWorkers atomically store preset workers config
func (a *appSync) StorePresetWorkers(c *worker.Config) {
	syncPresetWorkers.Store(c)
}

//LoadWorkers atomically load preset workers config
func (a *appSync) LoadPresetWorkers() *worker.Config {
	v := syncPresetWorkers.Load()
	if v == nil {
		return nil
	}
	return v.(*worker.Config)
}

func init() {
	config.RegisterLoader(util.ConstantsFile("/presetworkers.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, PresetWorkers))
		Sync.StorePresetWorkers(PresetWorkers)
	})
}
