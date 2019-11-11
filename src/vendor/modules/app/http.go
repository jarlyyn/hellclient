package app

import (
	forwarded "github.com/herb-go/herb/middleware/forwarded"
	"github.com/herb-go/herb/middleware/misc"
	"github.com/herb-go/herb/service/httpservice"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
)

type HTTPConfig struct {
	Forwarded forwarded.Middleware
	Config    httpservice.Config
	Headers   misc.Headers
}

var HTTP = &HTTPConfig{
	Forwarded: forwarded.Middleware{},
	Config:    httpservice.Config{},
	Headers:   misc.Headers{},
}

func init() {
	config.RegisterLoaderAndWatch(util.ConfigFile("/http.toml"), func(configpath util.FileObject) {
		util.Must(tomlconfig.Load(configpath, HTTP))
		util.SetWarning("Forwarded", HTTP.Forwarded.Warnings()...)
	})
}
