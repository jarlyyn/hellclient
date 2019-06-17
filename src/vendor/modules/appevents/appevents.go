package appevents

import (
	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "100AppEvents"

func init() {
	util.RegisterModule(ModuleName, func() {
		util.InitOrderByName(ModuleName)
	})
}
