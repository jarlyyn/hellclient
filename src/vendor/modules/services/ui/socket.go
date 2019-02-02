package ui

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Current *Conn
var Locker sync.RWMutex
var upgrader = websocket.Upgrader{} // use default options

type Conn struct {
	*websocket.Conn
	closed bool
	Lock   sync.Mutex
}

func (c *Conn) Close() error {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.closed = true
	return c.Conn.Close()
}

func (c *Conn) isClosed() bool {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	return c.closed

}

type Msg struct {
	Conn *Conn
	Type int
	Data []byte
}

func NewMsg() *Msg {
	return &Msg{}
}
func NewConn() *Conn {
	return &Conn{}
}

func OnMsg(msg *Msg) {
	fmt.Println(msg)
}
func Send(msgType int, data []byte) error {
	Locker.Lock()
	c := Current
	Locker.Unlock()
	if c == nil {
		return nil
	}
	return c.WriteMessage(msgType, data)
}

var OnErr func(err error)

func init() {
	OnErr = func(err error) {
		fmt.Println(err)
	}
}

var Enter = func(w http.ResponseWriter, r *http.Request) error {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	conn := NewConn()
	conn.Conn = c
	Locker.Lock()
	old := Current
	Current = conn
	Locker.Unlock()
	go func() {
		defer func() {

		}()
		defer func() {
			recover()
		}()
		for {
			mt, msg, err := c.ReadMessage()
			if err == io.EOF {
				break
			}
			if err != nil {
				if conn.isClosed() {
					break
				}
				if websocket.IsUnexpectedCloseError(err) {
					conn.Close()
					break
				}
				OnErr(err)
				continue
			}
			m := NewMsg()
			m.Conn = conn
			m.Type = mt
			m.Data = msg
			OnMsg(m)
		}
	}()
	if old != nil {
		old.Close()
	}
	return nil
}
