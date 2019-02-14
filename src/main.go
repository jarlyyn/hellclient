package main

import (
	"modules/app"
	"modules/services/client"

	// "modules/services/ui"

	"github.com/herb-go/herbgo/util"
	"github.com/herb-go/herbgo/util/config"
)

//Must panic if any error rasied
var Must = util.Must

func loadConfigs() {
	//Uncomment next line to print config loading log .
	//config.Debug = true
	config.Lock.RLock()
	app.LoadConfigs()
	config.Lock.RUnlock()
}
func initModules() {
	util.InitModulesOrderByName()
	//Put Your own init code here.
}

//Main app run func.
var run = func() {
	//Run app as http server.
	RunHTTP()
}

func main() {
	// Set app root path.
	//Default rootpah is "src/../"
	//You can set os env  "HerbRoot" to overwrite this setting while starting app.
	util.RootPath = ""
	defer util.Recover()
	util.UpdatePaths()
	util.MustChRoot()
	loadConfigs()
	initModules()
	app.Development.MustNotInitializing()
	util.MustLoadRegisteredFolders()
	// ui.Change("hell")
	config := client.ClientConfig{}
	config.World.Host = "220.165.145.126"
	config.World.Port = "3001"
	config.World.Charset = "gbk"
	m := client.DefaultManager
	m.NewClient("hell", config)
	m.NewClient("hell2", config)
	m.NewClient("hell3", config)

	run()
}
