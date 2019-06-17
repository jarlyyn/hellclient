package app

import (
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/commonconfig"
	"github.com/herb-go/util/config/tomlconfig"
)

var Development = commonconfig.DevelopmentConfig{}

func init() {
	config.RegisterLoader(util.Config("/development.toml"), func(configpath string) {
		util.Must(tomlconfig.Load(configpath, &Development))
		if util.ForceDebug {
			Development.Debug = true
		}
		if Development.Debug {
			util.Debug = true
			util.SetWarning("development.toml", "Debug value is true.")
		}
	})
}
