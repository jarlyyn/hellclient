package messages

import "github.com/herb-go/herbgo/util"

//Modulename module name used in initing and debuging.
const Modulename = "100Message"

func init() {
	util.RegisterModule(Modulename, func() {

	})
}
