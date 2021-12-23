package genesis

import (
	"hellclient/modules/world/prophet"
	"hellclient/modules/world/titan"
)

func Start() {
	titan.CreatePangu()
	prophet.Laozi = prophet.New()
	prophet.Laozi.Init(titan.Pangu)
	prophet.Laozi.Start()
	titan.Pangu.Start()
}

func Stop() {
	prophet.Laozi.Stop()
	titan.Pangu.Stop()
}
