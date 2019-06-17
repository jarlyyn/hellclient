package conn

import (
	"github.com/herb-go/util"
)
//ModuleName module name
const ModuleName="900services.conn"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func()error{})
		util.InitOrderByName(ModuleName)
	})
}
