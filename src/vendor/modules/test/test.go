package test

import (
	_ "modules" //modules init
	"modules/app"
	_ "modules/drivers" //drivers
	"modules/overseers"
	"path/filepath"

	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/testingtools"
)

func loadConfigs() {
	//Uncomment next line to print config loading log .
	//config.Debug = true
	config.Lock.RLock()
	config.LoadAll()
	config.Lock.RUnlock()
}

func initModules() {
	util.InitModulesOrderByName()
	//Put Your own init code here.
}

//Init init app config and modules
func Init() {
	util.ApplieationLock.Lock()
	defer util.ApplieationLock.Unlock()
	testingtools.SetRootPathRelativeToModules("../")
	util.UpdatePaths()
	util.ConfigPath = filepath.Join(util.RootPath, "test", "testconfig")
	util.AppDataPath = filepath.Join(util.RootPath, "test", "testappdata")
	util.MustChRoot()
	loadConfigs()
	overseers.MustInitOverseers()
	initModules()
	app.Development.TestingOrPanic()
	//Auto created appdata folder if not exists
	util.RegisterDataFolder()
	util.MustLoadRegisteredFolders()
	app.Development.InitializeAndPanicIfNeeded()
	overseers.MustTrainWorkers()
}

//Run run app
func Run() {
	//Put your run code here
	util.WaitingQuit()
	//Delay util.QuitDelayDuration for modules quit.
	util.DelayAndQuit()

}
