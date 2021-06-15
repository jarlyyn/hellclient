package world

import (
	"modules/world/prophet"
	"modules/world/titan"

	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "900world"

func Start() {
	titan.CreatePangu()
	prophet.Laozi = prophet.New()
	prophet.Laozi.Init(titan.Pangu)
	prophet.Laozi.Start()
}

func Stop() {
	prophet.Laozi.Stop()
}
func init() {
	util.RegisterModule(ModuleName, func() {

		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
	})
}
