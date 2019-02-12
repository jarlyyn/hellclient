package client

import (
	"sync"
)

type ClientConfig struct {
	World World
}
type Manager struct {
	Clients       map[string]*Client
	Lock          sync.RWMutex
	CommandOutput chan *Command
}

func NewManger() *Manager {
	return &Manager{
		Clients:       map[string]*Client{},
		CommandOutput: make(chan *Command),
	}
}
func (m *Manager) NewClient(id string, config ClientConfig) *Client {
	client := New()
	client.ID = id
	client.World = config.World
	client.Exit = make(chan int)
	client.Manager = m
	client.Init()
	m.Lock.Lock()
	m.Clients[id] = client
	m.Lock.Unlock()
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
	cmd := &Command{
		Type: "line",
		Room: id,
		Data: line,
	}
	m.CommandOutput <- cmd
}

func (m *Manager) OnPrompt(id string, line *Line) {
	cmd := &Command{
		Type: "prompt",
		Room: id,
		Data: line,
	}
	m.CommandOutput <- cmd

}
func (m *Manager) ExecConnect(id string) {
	c := m.Client(id)
	if c != nil {
		c.Connect()
	}
}
func (m *Manager) Send(id string, msg []byte) {
	c := m.Client(id)
	if c != nil {
		c.Send(msg)
	}
}

var DefaultManager = NewManger()
