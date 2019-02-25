package script

import (
	"fmt"
	"path"

	"github.com/herb-go/herbgo/util"
)

//ModuleName module name
const ModuleName = "900services.script"

var Path string

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func()error{})
		util.InitOrderByName(ModuleName)
		Path = util.RegisterDataFolder("scripts")
		l := NewLua()
		err := l.Load(path.Join(Path, "helllua-master", "hell.lua"))
		fmt.Println(err.Error())
	})
}
