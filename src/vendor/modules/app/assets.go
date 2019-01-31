package app

import (
	"github.com/herb-go/herb/file/store"
	"github.com/herb-go/herbgo/util"
	"github.com/herb-go/herbgo/util/config"
	"github.com/herb-go/herbgo/util/config/tomlconfig"
)

var Assets = store.Assets{}

func init() {
	config.RegisterLoader(util.Constants("/assets.toml"), func(configpath string) {
		util.Must(tomlconfig.Load(configpath, &Assets))
	})
}
