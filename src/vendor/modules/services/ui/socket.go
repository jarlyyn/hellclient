package ui

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Current *websocket.Conn
var Locker sync.RWMutex
var upgrader = websocket.Upgrader{} // use default options

var Enter = func(w http.ResponseWriter, r *http.Request) error {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	Locker.Lock()
	old := Current
	Current = c
	Locker.Unlock()
	go func() {
		for {

		}
		for {

		}
	}()
	if c != nil {
		c.Close()
	}
}
