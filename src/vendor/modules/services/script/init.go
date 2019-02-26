package script

import (
	"fmt"
	"modules/services/mapper"
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
		// err := l.Load(path.Join(Path, "hell", "index.lua"))
		// // fmt.Println(err.Error())
		m, err := mapper.CommonRoomAllIniLoader.Open(path.Join(Path, "hell", "rooms_all.h"))
		if err != nil {
			fmt.Println(err.Error())
		}
		walking := m.NewWalking()
		walking.From = "26"
		walking.To = []string{"100"}
		steps := walking.Walk()
		for _, v := range steps {
			fmt.Println(v.Command)
		}
		// fmt.Println(m)
	})
}
