package titan

import (
	"net/http"
	"os"

	"github.com/herb-go/util"
	"golang.org/x/net/webdav"
)

var Pangu *Titan

const WorldsFolder = "worlds"
const Ext = ".toml"
const WorldsPrefix = "/worlds/"
const ScriptsFolder = "worlds"
const ScriptsPrefix = "scripts"

func CreatePangu() {
	Pangu = New()
	Pangu.Path = util.AppData(WorldsFolder)
	Pangu.Scriptpath = util.AppData(ScriptsFolder)
	os.MkdirAll(Pangu.Path, util.DefaultFolderMode)
	os.MkdirAll(Pangu.Scriptpath, util.DefaultFolderMode)

}

func NewWebdavServer() http.Handler {
	webdavserver := &webdav.Handler{
		Prefix:     WorldsPrefix,
		FileSystem: webdav.Dir(util.AppData(WorldsFolder)),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				panic(err)
			}
		},
	}
	return webdavserver
}
