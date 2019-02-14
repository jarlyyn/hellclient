package ui

import (
	"encoding/json"
	"fmt"
	"modules/services/client"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/jarlyyn/herb-go-experimental/connections"
	"github.com/jarlyyn/herb-go-experimental/connections/contexts"
	"github.com/jarlyyn/herb-go-experimental/connections/identifier"
	"github.com/jarlyyn/herb-go-experimental/connections/room"
	"github.com/jarlyyn/herb-go-experimental/connections/websocket"
)

func Send(conn connections.ConnectionOutput, msgtype string, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return conn.Send([]byte(msgtype + " " + string(bs)))
}

var users = identifier.NewMap()
var gateway = connections.NewGateway()
var rooms = room.NewRooms()
var current = &atomic.Value{}
var gamelock = sync.Mutex{}

func SendToUser(data []byte) error {
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

func (e *Engine) Location(conn connections.ConnectionOutput) *room.Location {
	ctx := e.Context(conn.ID())
	v, _ := ctx.Data.Load("rooms")
	return v.(*room.Location)
}
func (e *Engine) OnClose(conn connections.ConnectionOutput) {
	ctx := e.Context(conn.ID())
	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()
	v, _ := ctx.Data.Load("rooms")
	r := v.(*room.Location)
	r.LeaveAll()
	e.Contexts.OnClose(conn)
}
func (e *Engine) OnOpen(conn connections.ConnectionOutput) {
	e.Contexts.OnOpen(conn)
	ctx := e.Context(conn.ID())
	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()
	r := room.NewLocation(conn, rooms)
	var crid string
	v := current.Load()
	if v != nil {
		crid = v.(string)
		Send(conn, "current", crid)
	}
	if crid != "" {
		r.Join(crid)
	}
	client.DefaultManager.ExecClients()
	client.DefaultManager.ExecLines(crid)
	client.DefaultManager.ExecPrompt(crid)
	ctx.Data.Store("rooms", r)
}
func (e *Engine) OnMessage(msg *connections.Message) {
	go func() {
		_, _, cerr := handlers.Exec(msg)
		if cerr != nil {
			e.OnError(cerr)
		}
	}()
}
func (e *Engine) OnError(err *connections.Error) {
	fmt.Println(*err)
}

var CurretEngine = &Engine{}

var Change = func(roomid string) {
	var conn connections.ConnectionOutput
	gamelock.Lock()
	defer gamelock.Unlock()
	v, _ := users.Identities.Load("user")
	if v != nil {
		conn = v.(connections.ConnectionOutput)
		location := CurretEngine.Location(conn)
		v := current.Load()
		if v != nil {
			crid := v.(string)
			if crid != "" {
				location.Leave(crid)
			}
			location.Join(roomid)
		}
	}
	current.Store(roomid)
	go func() {
		client.DefaultManager.ExecLines(roomid)
		client.DefaultManager.ExecPrompt(roomid)
	}()
}

func init() {
	go connections.Consume(gateway, CurretEngine)
}
