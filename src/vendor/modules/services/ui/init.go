package ui

import (
	"github.com/herb-go/herbgo/util"
)
//ModuleName module name
const ModuleName="900services.ui"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func()error{})
		util.InitOrderByName(ModuleName)
	})
}
