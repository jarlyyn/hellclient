package scripts

import (
	"net/http"

	"github.com/herb-go/util"
	"golang.org/x/net/webdav"
)

//ModuleName module name
const ModuleName = "900scripts"

const ScriptFolder = "scripts"

const ScriptPrefix = "/scripts/"

func NewWebdavServer() http.Handler {
	webdavserver := &webdav.Handler{
		Prefix:     ScriptPrefix,
		FileSystem: webdav.Dir(util.AppData(ScriptFolder)),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				panic(err)
			}
		},
	}
	return webdavserver
}

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.RegisterDataFolder(ScriptFolder)
		util.InitOrderByName(ModuleName)
	})
}
