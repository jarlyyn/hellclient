package actions

import (
	"hellclient/modules/world/prophet"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func WebsocketAction(w http.ResponseWriter, r *http.Request) {
	err := prophet.Laozi.Enter(w, r)
	if err != nil {
		panic(err)
	}
}
