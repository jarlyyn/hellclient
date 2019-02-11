package ui

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/jarlyyn/herb-go-experimental/connections"
	"github.com/jarlyyn/herb-go-experimental/connections/contexts"
	"github.com/jarlyyn/herb-go-experimental/connections/identifier"
	"github.com/jarlyyn/herb-go-experimental/connections/room"
	"github.com/jarlyyn/herb-go-experimental/connections/websocket"
)

var users = identifier.NewMap()
var gateway = connections.NewGateway()
var rooms = room.NewRooms()
var current = &atomic.Value{}
var gamelock = sync.Mutex{}

func Send(data []byte) error {
	return users.SendByID("user", data)
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
	contexts.Contexts
}

func (e *Engine) OnOpen(conn connections.ConnectionOutput) {
	v := current.Load()
	if v != nil {
		crid := v.(string)
		conn.Send([]byte("current " + crid))
	}
}
func (e *Engine) OnMessage(msg *connections.Message) {
	cmd := ParseCmd(msg.Message)
	switch string(cmd[0]) {
	case CmdsChange:
		Change(string(cmd[1]))
		msg.Conn.Send([]byte("current " + string(cmd[1])))
	case CmdsConnect:

	}
}
func (e *Engine) OnError(err *connections.Error) {
	fmt.Println(*err)
}

var CurretEngine = &Engine{}

var Change = func(roomid string) {
	gamelock.Lock()
	defer gamelock.Unlock()
	v, _ := users.Identities.Load("user")
	if v != nil {
		conn := v.(connections.ConnectionOutput)
		v := current.Load()
		if v != nil {
			crid := v.(string)
			if crid != "" {
				rooms.Leave(crid, conn)
			}
			rooms.Join(roomid, conn)
		}
	}
	current.Store(roomid)
}
var cmdsep = []byte(" ")

func ParseCmd(data []byte) [2][]byte {
	var result [2][]byte
	cmds := bytes.SplitN(data, cmdsep, 2)
	if len(cmds) > 0 {
		result[0] = cmds[0]
	}
	if len(cmds) > 1 {
		result[1] = cmds[1]
	}
	return result
}

func init() {
	go connections.Consume(gateway, CurretEngine)
}
