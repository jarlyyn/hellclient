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
		s := New()
		l := NewLua()
		l.SetScript(s)
		err := l.Load(path.Join(Path, "hell", "index.lua"))
		fmt.Println(err.Error())
	})
}
