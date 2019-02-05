package ui

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/jarlyyn/herb-go-experimental/websocketmanager"
	"github.com/jarlyyn/herb-go-experimental/websocketmanager/websocket"
)

var Current *websocketmanager.RegisteredConn
var Locker sync.RWMutex
var socketmamager = websocketmanager.New()

func OnMsg(msg *websocketmanager.ConnMessage) {
	fmt.Println(string(msg.Message))
}
func Send(data []byte) error {
	Locker.Lock()
	c := Current
	Locker.Unlock()
	if c == nil {
		return nil
	}
	return c.Send(data)
}

var OnErr func(err *websocketmanager.ConnError)

func init() {
	OnErr = func(err *websocketmanager.ConnError) {
		fmt.Println(*err)
	}
}

var Enter = func(w http.ResponseWriter, r *http.Request) error {

	wc, err := websocket.Upgrade(w, r, websocket.MsgTypeText)
	if err != nil {
		return err
	}
	c, err := socketmamager.Register(wc)
	if err != nil {
		return err
	}
	Locker.Lock()
	old := Current
	Current = c
	Locker.Unlock()
	go func() {
		<-c.C()
	}()
	if old != nil {
		old.Close()
	}
	return nil
}

func Listen() {
	for {
		select {
		case m := <-socketmamager.ConnMessages:
			OnMsg(m)
		case err := <-socketmamager.ConnErrors:
			OnErr(err)
		}
	}
}
func init() {
	go Listen()
}
