package messenger

import (
	"encoding/json"
	"hellclient/modules/world"
	"hellclient/modules/world/titan"
	"net/http"

	"github.com/herb-go/connections"
	"github.com/herb-go/connections/room"
	"github.com/herb-go/connections/websocket"
	"github.com/herb-go/logger"
	"github.com/herb-go/uniqueid"
)

var TaiBaiJinXing *Messenger

type Messenger struct {
	connections.EmptyConsumer
	G     *connections.Gateway
	Room  *room.Room
	Titan *titan.Titan
}

func New() *Messenger {
	return &Messenger{
		G:    connections.NewGateway(),
		Room: room.NewRoom(),
	}
}
func (m *Messenger) Debug(err error) {
	logger.Debug(err.Error())
}
func (m *Messenger) Publish(t *titan.Titan, msg *world.Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		m.Debug(err)
		return
	}
	m.Room.Broadcast(data, m.Debug)
}
func (m *Messenger) Init(t *titan.Titan) {
	m.G.IDGenerator = uniqueid.DefaultGenerator.GenerateID
	t.BindRequestEvent(m, m.Publish)
	m.Titan = t
}

func (m *Messenger) Enter(w http.ResponseWriter, r *http.Request) error {
	wc, err := websocket.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	_, err = m.G.Register(wc)
	if err != nil {
		return err
	}
	return nil
}

//OnMessage called when connection message received.
func (m *Messenger) OnMessage(msg *connections.Message) {
	data := world.NewMessage()
	err := json.Unmarshal(msg.Message, data)
	if err != nil {
		logger.Debug(err.Error())
		return
	}
	m.Titan.OnResponse(data)
}

//OnError called when onconnection error raised.
func (m *Messenger) OnError(err *connections.Error) {
	logger.Debug(err.Error.Error())
}

//OnClose called when connection closed.
func (m *Messenger) OnClose(conn connections.OutputConnection) {
	m.Room.Leave(conn)
}

//OnOpen called when connection open.
func (m *Messenger) OnOpen(conn connections.OutputConnection) {
	m.Room.Join(conn)
}

// Stop stop consumer
func (m *Messenger) Stop() {
	m.G.Stop()
}

func (m *Messenger) Start() {
	go connections.Consume(m.G, m)

}
