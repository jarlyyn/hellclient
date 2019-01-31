package app

import (
	"github.com/herb-go/herbgo/util"
	"github.com/herb-go/herbgo/util/config"
	"github.com/herb-go/herbgo/util/config/commonconfig"
	"github.com/herb-go/herbgo/util/config/tomlconfig"
)

//Time app time settigs
var Time = commonconfig.TimeConfig{}

func init() {
	config.RegisterLoader(util.Constants("/time.toml"), func(configpath string) {
		util.Must(tomlconfig.Load(configpath, &Time))
	})
}
