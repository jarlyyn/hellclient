package app

import (
	"sync/atomic"

	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
)

//SystemConfig system config data struct.
//Struct must  unmarshaleable by Toml lib.
//You should comment this struct if you use third party config struct.
type SystemConfig struct {
	Addr     string
	Username string
	Password string
	Switch   string
}

//System config instance of system.
var System = &SystemConfig{}

var syncSystem atomic.Value

//StoreSystem atomically store system config
func (a *appSync) StoreSystem(c *SystemConfig) {
	syncSystem.Store(c)
}

//LoadSystem atomically load system config
func (a *appSync) LoadSystem() *SystemConfig {
	v := syncSystem.Load()
	if v == nil {
		return nil
	}
	return v.(*SystemConfig)
}

func init() {
	//Register loader which will be execute when Config.LoadAll func be called.
	//You can put your init code after load.
	//You must panic if any error rasied when init.
	config.RegisterLoader(util.ConfigFile("/system.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, System))
		Sync.StoreSystem(System)
	})
}
