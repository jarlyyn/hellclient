package ui

import (
	"encoding/json"
	"fmt"
	"modules/services/client"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/herb-go/connections"
	"github.com/herb-go/connections/contexts"
	"github.com/herb-go/connections/identifier"
	"github.com/herb-go/connections/room"
	"github.com/herb-go/connections/websocket"
)

func Send(conn connections.OutputConnection, msgtype string, data interface{}) error {
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

func CurrentGameID() string {
	return current.Load().(string)
}
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

func (e *Engine) Location(conn connections.OutputConnection) *room.Location {
	ctx := e.Context(conn.ID())
	v, _ := ctx.Data.Load("rooms")
	return v.(*room.Location)
}
func (e *Engine) OnClose(conn connections.OutputConnection) {
	ctx := e.Context(conn.ID())
	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()
	v, _ := ctx.Data.Load("rooms")
	r := v.(*room.Location)
	r.LeaveAll()
	e.Contexts.OnClose(conn)
}
func (e *Engine) OnOpen(conn connections.OutputConnection) {
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
	var conn connections.OutputConnection
	gamelock.Lock()
	defer gamelock.Unlock()
	v, _ := users.Identities.Load("user")
	if v != nil {
		conn = v.(connections.OutputConnection)
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
