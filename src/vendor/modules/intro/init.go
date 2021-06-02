package intro

import (
	"fmt"
	"modules/app"

	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "z99intro"

const Version = "0.0.1"

func init() {
	util.StageFinish.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		fmt.Printf("Hellclient v%s\n", Version)
		fmt.Printf("Listening http on %s\n", app.System.Addr)
		util.InitOrderByName(ModuleName)
	})
}
