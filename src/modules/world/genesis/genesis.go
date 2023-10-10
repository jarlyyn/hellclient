package genesis

import (
	"modules/world/messenger"
	"modules/world/prophet"
	"modules/world/titan"
)

func Start() {
	titan.CreatePangu()
	prophet.Laozi = prophet.New()
	prophet.Laozi.Init(titan.Pangu)
	prophet.Laozi.Start()
	messenger.TaiBaiJinXing = messenger.New()
	messenger.TaiBaiJinXing.Init(titan.Pangu)
	messenger.TaiBaiJinXing.Start()
	titan.Pangu.Start()
}

func Stop() {
	prophet.Laozi.Stop()
	titan.Pangu.Stop()
	messenger.TaiBaiJinXing.Stop()
}
