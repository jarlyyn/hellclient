package world

import (
	"github.com/herb-go/util"
)

const Folder = "worlds"
const Ext = ".toml"

func GetWorldPath(id string) string {
	return util.AppData(Folder, id) + Ext
}
