package app

import (
	"github.com/herb-go/herb/middleware/csrf"
	"github.com/herb-go/herbgo/util"
	"github.com/herb-go/herbgo/util/config"
	"github.com/herb-go/herbgo/util/config/tomlconfig"
)

var Csrf = csrf.Config{}

func init() {
	config.RegisterLoader(util.Config("/csrf.toml"), func(configpath string) {
		util.Must(tomlconfig.Load(configpath, &Csrf))
	})
}
