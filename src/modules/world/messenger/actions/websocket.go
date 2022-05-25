package actions

import (
	"hellclient/modules/world/messenger"
	"net/http"
)

func WebsocketAction(w http.ResponseWriter, r *http.Request) {
	err := messenger.TaiBaiJinXing.Enter(w, r)
	if err != nil {
		panic(err)
	}
}
