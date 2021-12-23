package msg

import (
	"github.com/herb-go/util"
)
//ModuleName module name
const ModuleName="900msg"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
	})
}
