package ui

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/jarlyyn/herb-go-experimental/connections"
	"github.com/jarlyyn/herb-go-experimental/connections/identifier"
	"github.com/jarlyyn/herb-go-experimental/connections/websocket"
)

var users = identifier.NewMap()
var Current connections.ConnectionOutput
var Locker sync.RWMutex
var gateway = connections.NewGateway()

func Send(data []byte) error {
	Locker.Lock()
	c := Current
	Locker.Unlock()
	if c == nil {
		return nil
	}
	return c.Send(data)
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
	users.Login("user", c)
	return nil
}

type Engine struct {
	connections.EmptyConsumer
}

func (e *Engine) OnMessage(msg *connections.Message) {
	fmt.Println(string(msg.Message))

}
func (e *Engine) OnError(err *connections.Error) {
	fmt.Println(*err)
}

var CurretEngine = &Engine{}

func init() {
	go connections.Consume(gateway, CurretEngine)
}
