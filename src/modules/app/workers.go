package app

import (
	"sync/atomic"

	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
	"github.com/herb-go/worker"
)

//Workers workers cpmfog
var Workers = &worker.Config{}

var syncWorkers atomic.Value

//StoreWorkers atomically store workers config
func (a *appSync) StoreWorkers(c *worker.Config) {
	syncWorkers.Store(c)
}

//LoadWorkers atomically load workers config
func (a *appSync) LoadWorkers() *worker.Config {
	v := syncWorkers.Load()
	if v == nil {
		return nil
	}
	return v.(*worker.Config)
}

func init() {
	config.RegisterLoader(util.ConfigFile("/workers.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, Workers))
		Sync.StoreWorkers(Workers)
	})
}
