package app

import (
	"github.com/herb-go/herb/file/store"
	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
)

var Assets = store.Assets{}

func init() {
	config.RegisterLoader(util.ConstantsFile("/assets.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, &Assets))
	})
}
