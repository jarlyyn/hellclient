package client

import (
	"github.com/herb-go/herbgo/util"
)

//ModuleName module name
const ModuleName = "900services.client"

var WorldsPath = util.RegisterDataFolder("worlds")

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func()error{})
		util.InitOrderByName(ModuleName)

	})
}
