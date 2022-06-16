package prophet

import (
	"encoding/json"
	"hellclient/modules/world/titan"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/herb-go/logger"
	"github.com/herb-go/uniqueid"

	"github.com/herb-go/connections"

	"github.com/herb-go/connections/contexts"
	"github.com/herb-go/connections/room"
	"github.com/herb-go/connections/websocket"

	"github.com/herb-go/connections/command"
	"github.com/herb-go/connections/identifier"
	"github.com/herb-go/connections/room/message"
)

type Prophet struct {
	Current  atomic.Value
	Users    *identifier.Map
	Locker   sync.RWMutex
	Titan    *titan.Titan
	Adapter  *message.Adapter
	Handlers *command.Handlers
	Rooms    *room.Rooms
	Gateway  *connections.Gateway
	*contexts.Contexts
}

func (p *Prophet) Init(t *titan.Titan) {
	p.Titan = t
	t.BindMsgEvent(p, p.Publish)
	p.Current.Store("")
	p.Gateway.IDGenerator = uniqueid.DefaultGenerator.GenerateID
	initAdapter(p, p.Adapter)
	initHandlers(p, p.Handlers)
}
func (p *Prophet) newRoomAdapter(cmdtype string) func(m *message.Message) error {
	return func(m *message.Message) error {
		var err error
		if m.Room != "" {
			data := command.New()
			data.CommandType = cmdtype
			data.CommandData, err = json.Marshal(m.Data)
			if err != nil {
				return err
			}
			msg, err := data.Encode()
			if err != nil {
				return err
			}
			p.Rooms.Broadcast(m.Room, msg, nil)
		}
		return nil
	}
}
func (p *Prophet) sendToUser(data []byte) error {
	return p.Users.SendByID("user", data)
}

func (p *Prophet) newUserAdapter(cmdtype string) func(m *message.Message) error {
	return func(m *message.Message) error {
		var err error
		if m.Room == "" {
			data := command.New()
			data.CommandType = cmdtype
			data.CommandData, err = json.Marshal(m.Data)
			if err != nil {
				return err
			}
			msg, err := data.Encode()
			if err != nil {
				return err
			}
			return p.sendToUser(msg)
		}
		return nil
	}
}
func (p *Prophet) Location(conn connections.OutputConnection) *room.Location {
	ctx := p.Context(conn.ID())
	if ctx == nil {
		return nil
	}
	v, _ := ctx.Data.Load("rooms")
	return v.(*room.Location)
}

func (p *Prophet) Change(roomid string) {
	var conn connections.OutputConnection
	p.Locker.Lock()
	defer p.Locker.Unlock()
	v, _ := p.Users.Identities.Load("user")
	if v != nil {
		conn = v.(connections.OutputConnection)
		location := p.Location(conn)
		if location == nil {
			return
		}
		v := p.Current.Load()
		if v != nil {
			crid := v.(string)
			if crid != "" {
				location.Leave(crid)
			}
			location.Join(roomid)
		}
	}
	p.Current.Store(roomid)
	go func() {
		p.onCurrent(roomid)
	}()
}
func (p *Prophet) Enter(w http.ResponseWriter, r *http.Request) error {
	wc, err := websocket.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	c, err := p.Gateway.Register(wc)
	if err != nil {
		return err
	}
	p.Users.Login("user", c)
	return nil

}

//OnMessage called when connection message received.
func (p *Prophet) OnMessage(msg *connections.Message) {
	_, _, cerr := p.Handlers.Exec(msg)
	if cerr != nil {
		p.OnError(cerr)
	}

}

//OnError called when onconnection error raised.
func (p *Prophet) OnError(err *connections.Error) {
	logger.Debug(err.Error.Error())
}

//OnClose called when connection closed.
func (p *Prophet) OnClose(conn connections.OutputConnection) {
	ctx := p.Context(conn.ID())
	if ctx != nil {
		ctx.Lock.Lock()
		defer ctx.Lock.Unlock()
		v, _ := ctx.Data.Load("rooms")
		r := v.(*room.Location)
		r.LeaveAll()
	}
	p.Contexts.OnClose(conn)
}

//OnOpen called when connection open.
func (p *Prophet) OnOpen(conn connections.OutputConnection) {
	p.Contexts.OnOpen(conn)
	ctx := p.Context(conn.ID())
	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()
	r := room.NewLocation(conn, p.Rooms)
	var crid string
	v := p.Current.Load()
	if v != nil {
		crid = v.(string)
		Send(conn, "current", crid)
	}
	if crid != "" {
		r.Join(crid)
	}
	p.Titan.ExecClients()
	p.Titan.ExecSwitchStatus()
	p.onCurrent(crid)
	ctx.Data.Store("rooms", r)

}

func (p *Prophet) onCurrent(roomid string) {
	p.Titan.HandleCmdLines(roomid)
	p.Titan.HandleCmdPrompt(roomid)
	p.Titan.HandleCmdStatus(roomid)
	p.Titan.HandleCmdHistory(roomid)
	p.Titan.HandleCmdHUDContent(roomid)

}
func (p *Prophet) Start() {
	go connections.Consume(p.Gateway, p)
}

// Stop stop consumer
func (p *Prophet) Stop() {
	p.Gateway.Stop()
}
func (p *Prophet) Publish(t *titan.Titan, msg *message.Message) {
	p.Adapter.Exec(msg)
}
func New() *Prophet {
	return &Prophet{
		Adapter:  message.NewAdapter(),
		Handlers: command.NewHandlers(),
		Users:    identifier.New(),
		Rooms:    room.NewRooms(),
		Gateway:  connections.NewGateway(),
		Contexts: contexts.New(),
	}
}

var Laozi *Prophet
