package titan

import (
	"modules/app"
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
const ModsFolder = "/game/mods"

const GamePrefix = "/game/"
const GameFolder = "game"

const SkeletonsFolder = "/game/skeletons"

func CreatePangu() {
	Pangu = New()
	Pangu.Path = util.AppData(WorldsFolder)
	Pangu.Scriptpath = util.AppData(ScriptsFolder)
	Pangu.Logpath = util.AppData(LogsFolder)
	Pangu.Modpath = util.AppData(ModsFolder)
	Pangu.MaxHistory = app.System.MaxHistory
	if Pangu.MaxHistory <= 0 {
		Pangu.MaxHistory = DefaultMaxHistory
	}
	Pangu.MaxLines = app.System.MaxLines
	if Pangu.MaxLines <= 0 {
		Pangu.MaxLines = DefaultMaxLines
	}
	Pangu.MaxRecent = app.System.MaxRecent
	if Pangu.MaxRecent == 0 {
		Pangu.MaxRecent = DefaultMaxRecent
	} else if Pangu.MaxRecent < 0 {
		Pangu.MaxRecent = 0
	}
	Pangu.LinesPerScreen = app.System.LinesPerScreen
	if Pangu.LinesPerScreen <= 0 {
		Pangu.LinesPerScreen = DefaultLinesPerScreen
	}
	os.MkdirAll(Pangu.Path, util.DefaultFolderMode)
	os.MkdirAll(Pangu.Scriptpath, util.DefaultFolderMode)
	os.MkdirAll(Pangu.Logpath, util.DefaultFolderMode)
	os.MkdirAll(Pangu.Modpath, util.DefaultFolderMode)
}

func NewWebdavServer() http.Handler {
	webdavserver := &webdav.Handler{
		Prefix:     GamePrefix,
		FileSystem: webdav.Dir(util.AppData(GameFolder)),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				if os.IsNotExist(err) {
					return
				}
				util.LogError(err)
			}
		},
	}
	return webdavserver
}
