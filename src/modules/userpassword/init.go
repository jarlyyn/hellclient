package userpassword

import (
	"hellclient/modules/persistdata"

	"github.com/herb-go/herb/persist"
	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "900userpassword"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
		u := &UserPassword{}
		err := persistdata.Load("userpassword", u)
		if err == nil {
			current.Store(u)
		} else if err != persist.ErrNotFound {
			panic(err)
		}

	})
}
