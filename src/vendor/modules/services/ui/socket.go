package ui

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/jarlyyn/herb-go-experimental/connections"
	"github.com/jarlyyn/herb-go-experimental/connections/websocket"
)

var Current *connections.Conn
var Locker sync.RWMutex
var gateway = connections.NewGateway()

func OnMsg(msg *connections.Message) {
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

var OnErr func(err *connections.Error)

func init() {
	OnErr = func(err *connections.Error) {
		fmt.Println(*err)
	}
}

var Enter = func(w http.ResponseWriter, r *http.Request) error {

	wc, err := websocket.Upgrade(w, r, websocket.MsgTypeText)
	if err != nil {
		return err
	}
	c, err := gateway.Register(wc)
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
		case m := <-gateway.Messages:
			OnMsg(m)
		case err := <-gateway.Errors:
			OnErr(err)
		}
	}
}
func init() {
	go Listen()
}
