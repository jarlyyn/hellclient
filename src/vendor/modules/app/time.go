package app

import (
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/commonconfig"
	"github.com/herb-go/util/config/tomlconfig"
)

//Time app time settigs
var Time = commonconfig.TimeConfig{}

func init() {
	config.RegisterLoader(util.ConstantsFile("/time.toml"), func(configpath util.FileObject) {
		util.Must(tomlconfig.Load(configpath, &Time))
	})
}
