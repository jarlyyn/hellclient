package app

import (
	"path/filepath"
	"testing"

	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/testingtools"
)

func TestConfig(t *testing.T) {
	util.ForceDebug = true
	testingtools.SetRootPathRelativeToModules("../")
	util.UpdatePaths()
	util.ConfigPath = filepath.Join(util.RootPath, "test", "testconfig")
	util.AppDataPath = filepath.Join(util.RootPath, "test", "testappdata")
	config.LoadAll()
}
