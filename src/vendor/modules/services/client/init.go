package client

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/herb-go/herbgo/util/config/tomlconfig"

	"github.com/herb-go/herbgo/util"
)

//ModuleName module name
const ModuleName = "900services.client"

var WorldsPath string

func MustLoadWorlds() {
	files, err := ioutil.ReadDir(WorldsPath)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		filename := f.Name()
		if strings.HasSuffix(filename, ".toml") {
			id := filename[:len(filename)-5]
			world := &World{}
			tomlconfig.MustLoad(filepath.Join(WorldsPath, filename), world)
			DefaultManager.NewClient(id, world)
		}
	}

}
func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func()error{})
		util.InitOrderByName(ModuleName)
		WorldsPath = util.RegisterDataFolder("worlds")
		MustLoadWorlds()
	})
}
