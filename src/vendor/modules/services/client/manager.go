package client

import (
	"sync"

	"github.com/herb-go/herb/model"

	"github.com/jarlyyn/herb-go-experimental/connections/room/message"
)

type ClientConfig struct {
	World World
}
type Manager struct {
	Clients       map[string]*Client
	Lock          sync.RWMutex
	CommandOutput chan *message.Message
}

func NewManger() *Manager {
	return &Manager{
		Clients:       map[string]*Client{},
		CommandOutput: make(chan *message.Message),
	}
}
func (m *Manager) NewClient(id string, config ClientConfig) *Client {
	client := New()
	client.ID = id
	client.World = config.World
	client.Exit = make(chan int)
	client.Manager = m
	client.Init()
	m.Clients[id] = client
	go func() {
		<-client.Exit
	}()
	return client
}
func (m *Manager) Client(id string) *Client {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	return m.Clients[id]
}
func (m *Manager) Connect(id string) error {
	c := m.Client(id)
	if c == nil {
		return nil
	}
	return c.Connect()
}
func (m *Manager) OnLine(id string, line *Line) {
	msg := message.New()
	msg.Type = "line"
	msg.Room = id
	msg.Data = line
	m.CommandOutput <- msg
}
func newMsg(msgtype string, room string, data interface{}) *message.Message {
	msg := message.New()
	msg.Type = msgtype
	msg.Room = room
	msg.Data = data
	return msg
}
func (m *Manager) OnConnected(id string) {
	msg := newMsg("connected", "", id)
	go func() {
		m.CommandOutput <- msg
	}()
}
func (m *Manager) OnCreateFail(errors []model.FieldError) {
	msg := newMsg("createFail", "", errors)
	go func() {
		m.CommandOutput <- msg
	}()
}

func (m *Manager) OnDisconnected(id string) {
	msg := newMsg("disconnected", "", id)
	go func() {
		m.CommandOutput <- msg
	}()
}
func (m *Manager) OnPrompt(id string, line *Line) {
	msg := newMsg("prompt", id, line)
	go func() {
		m.CommandOutput <- msg
	}()

}
func (m *Manager) ExecPrompt(id string) {
	c := m.Client(id)
	var prompt = &Line{}
	if c != nil {
		prompt = c.Prompt
	}
	msg := newMsg("prompt", id, prompt)
	go func() {
		m.CommandOutput <- msg
	}()
}
func (m *Manager) ExecConnect(id string) {
	c := m.Client(id)
	if c != nil {
		c.Connect()
	}
}
func (m *Manager) ExecDisconnect(id string) {
	c := m.Client(id)
	if c != nil {
		c.Disconnect()
	}
}
func (m *Manager) ExecLines(id string) {
	c := m.Client(id)
	var lines = []*Line{}
	if c != nil {
		lines = c.ConvertLines()
	}
	msg := newMsg("lines", id, lines)
	go func() {
		m.CommandOutput <- msg
	}()
}
func (m *Manager) ExecClients() {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	var result = make([]*ClientInfo, len(m.Clients))
	var i = 0
	for _, v := range m.Clients {
		result[i] = v.Info()
		i++
	}
	msg := newMsg("clients", "", result)
	go func() {
		m.CommandOutput <- msg
	}()
}
func (m *Manager) Send(id string, msg []byte) {
	c := m.Client(id)
	if c != nil {
		c.Send(msg)
	}
}

var DefaultManager = NewManger()
