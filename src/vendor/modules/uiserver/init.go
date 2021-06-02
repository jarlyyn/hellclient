package uiserver

import (
	"modules/app"
	"modules/routers"
	"net/http"

	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "900uiserver"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		s := http.Server{}
		s.Addr = app.System.Addr
		s.Handler = routers.RouterUIFactory.CreateRouter()
		go s.ListenAndServe()
		util.OnQuitAndLogError(s.Close)
		util.InitOrderByName(ModuleName)
	})
}
