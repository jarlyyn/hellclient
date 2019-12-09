package app

import (
	"github.com/herb-go/herb/middleware/csrf"
	"github.com/herb-go/herbconfig/configuration"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
)

var Csrf = csrf.Config{}

func init() {
	config.RegisterLoader(util.ConfigFile("/csrf.toml"), func(configpath configuration.Configuration) {
		util.Must(tomlconfig.Load(configpath, &Csrf))
	})
}
