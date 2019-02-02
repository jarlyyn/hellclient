package actions

import (
	"modules/services/ui"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func WebsocketAction(w http.ResponseWriter, r *http.Request) {
	err := ui.Enter(w, r)
	if err != nil {
		panic(err)
	}
}
