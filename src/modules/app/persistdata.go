package app

import (
	"sync/atomic"

	"github.com/herb-go/herb-drivers/persist/hiredpersist"
	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
)

//Persistdata persist config
var Persistdata = &hiredpersist.Config{}

var syncPersistdata atomic.Value

//StorePersistdata atomically store persist config
func (a *appSync) StorePersistdata(m *hiredpersist.Config) {
	syncPersistdata.Store(m)
}

//LoadPersistdata atomically load persist config
func (a *appSync) LoadPersistdata() *hiredpersist.Config {
	v := syncPersistdata.Load()
	if v == nil {
		return nil
	}
	return v.(*hiredpersist.Config)
}

func init() {
	config.RegisterLoader(util.ConfigFile("/persistdata.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, Persistdata))
		Sync.StorePersistdata(Persistdata)
	})
}
