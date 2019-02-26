package mapper

import (
	"github.com/herb-go/herbgo/util"
)
//ModuleName module name
const ModuleName="900services.mapper"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
	})
}
