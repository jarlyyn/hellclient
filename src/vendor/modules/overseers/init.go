package overseers

import (
	"modules/app"
	_ "modules/hired" //hired workers

	"github.com/herb-go/util"
	"github.com/herb-go/worker"
)

//MustInitOverseers init overseers.
//Panic if any error raised
func MustInitOverseers() {
	worker.Debug = app.Development.Debug
	util.Must(app.PresetWorkers.Apply())
	util.Must(app.Workers.Apply())
	util.Must(worker.InitOverseers())
	worker.Start()
	util.OnQuit(worker.Stop)
}

//MustTrainWorkers with workers config
func MustTrainWorkers() {
	util.Must(worker.TrainWorkers())
}
