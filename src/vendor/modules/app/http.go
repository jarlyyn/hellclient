package app

import (
	forwarded "github.com/herb-go/herb/middleware/forwarded"
	"github.com/herb-go/herb/middleware/misc"
	"github.com/herb-go/herbgo/util"
	"github.com/herb-go/herbgo/util/config"
	"github.com/herb-go/herbgo/util/config/tomlconfig"
	"github.com/herb-go/herbgo/util/httpserver"
)

type HTTPConfig struct {
	Forwarded forwarded.Middleware
	Config    httpserver.Config
	Headers   misc.Headers
}

var HTTP = &HTTPConfig{
	Forwarded: forwarded.Middleware{},
	Config:    httpserver.Config{},
	Headers:   misc.Headers{},
}

func init() {
	config.RegisterLoaderAndWatch(util.Config("/http.toml"), func(configpath string) {
		util.Must(tomlconfig.Load(configpath, HTTP))
		util.SetWarning("Forwarded", HTTP.Forwarded.Warnings()...)
	})
}
