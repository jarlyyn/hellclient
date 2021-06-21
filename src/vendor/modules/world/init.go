package world

import (
	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "900world"

func init() {
	util.RegisterModule(ModuleName, func() {
		initTemplates()
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
	})
}
