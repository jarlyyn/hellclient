package genesis

import (
	"modules/world/prophet"
	"modules/world/titan"
)

func Start() {
	titan.CreatePangu()
	prophet.Laozi = prophet.New()
	prophet.Laozi.Init(titan.Pangu)
	prophet.Laozi.Start()
}

func Stop() {
	prophet.Laozi.Stop()
}
