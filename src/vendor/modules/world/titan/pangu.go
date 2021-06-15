package titan

import (
	"os"

	"github.com/herb-go/util"
)

var Pangu *Titan

const WorldsFolder = "worlds"
const Ext = ".toml"

func CreatePangu() {
	Pangu = New()
	Pangu.Path = util.AppData(WorldsFolder)
	os.MkdirAll(Pangu.Path, util.DefaultFolderMode)
}
