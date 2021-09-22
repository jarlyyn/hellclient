package titan

import (
	"net/http"
	"os"

	"github.com/herb-go/util"
	"golang.org/x/net/webdav"
)

var Pangu *Titan

const WorldsFolder = "game/worlds"
const Ext = ".toml"

const ScriptsFolder = "/game/scripts"
const LogsFolder = "/game/logs"

const GamePrefix = "/game/"
const GameFolder = "game"

const SkeletonsFolder = "/game/skeletons"

func CreatePangu() {
	Pangu = New()
	Pangu.Path = util.AppData(WorldsFolder)
	Pangu.Scriptpath = util.AppData(ScriptsFolder)
	Pangu.Logpath = util.AppData(LogsFolder)
	Pangu.Skeletonpath = util.AppData(SkeletonsFolder)
	os.MkdirAll(Pangu.Path, util.DefaultFolderMode)
	os.MkdirAll(Pangu.Scriptpath, util.DefaultFolderMode)
	os.MkdirAll(Pangu.Logpath, util.DefaultFolderMode)

}

func NewWebdavServer() http.Handler {
	webdavserver := &webdav.Handler{
		Prefix:     GamePrefix,
		FileSystem: webdav.Dir(util.AppData(GameFolder)),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				panic(err)
			}
		},
	}
	return webdavserver
}
