package hellswitch

import (
	"encoding/base64"
	"modules/app"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/herb-go/util"

	"github.com/gorilla/websocket"
)

const ReconnectDuration = 5 * time.Minute
const StatusDisabled = 0
const StatusDisconnected = 1
const StatusConnected = 2

var CmdBroadcast = []byte("broadcast ")
var CmdHello = []byte("hello")

type Hellswitch struct {
	ticker               *time.Ticker
	c                    chan struct{}
	locker               sync.Mutex
	conn                 *websocket.Conn
	OnGlobalMessage      func(msg []byte)
	OnSwitchStatusChange func(status int)
}

func (h *Hellswitch) Status() int {
	h.locker.Lock()
	defer h.locker.Unlock()
	if app.System.Switch == "" {
		return StatusDisabled
	}
	if h.conn == nil {
		return StatusDisconnected
	}
	return StatusConnected
}
func (h *Hellswitch) Broadcast(msg []byte) {
	h.locker.Lock()
	defer h.locker.Unlock()
	if h.conn == nil {
		return
	}
	cmd := append([]byte{}, CmdBroadcast...)
	cmd = append(cmd, msg...)
	h.conn.WriteMessage(websocket.TextMessage, cmd)
}
func (h *Hellswitch) listen(conn *websocket.Conn) {
	go func() {
		conn.WriteMessage(websocket.TextMessage, CmdHello)
	}()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if _, ok := err.(*websocket.CloseError); !ok {
				util.LogError(err)
				conn.Close()
			} else {
				h.locker.Lock()
				h.conn = nil
				h.locker.Unlock()
				go h.OnSwitchStatusChange(StatusDisconnected)
				return
			}
			go h.OnGlobalMessage(message)
		}
	}
}
func (h *Hellswitch) basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func (h *Hellswitch) Start() {
	h.locker.Lock()
	defer h.locker.Unlock()
	if h.conn != nil {
		return
	}
	s := app.System.Switch
	if s == "" {
		return
	}
	u, err := url.Parse(s)
	if err != nil {
		util.LogError(err)
		return
	}
	header := http.Header{}
	if u.User != nil {
		un := u.User.Username()
		up, ok := u.User.Password()
		if ok {
			header.Set("Authorization", "Basic "+h.basicAuth(un, up))
		}
		u.User = nil
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		go h.OnSwitchStatusChange(StatusDisconnected)
		util.LogError(err)
		return
	}
	h.conn = c
	go h.listen(c)
	go h.OnSwitchStatusChange(StatusConnected)
}

func (h *Hellswitch) Stop() {
	h.locker.Lock()
	defer h.locker.Unlock()
	if h.conn == nil {
		return
	}
	h.conn.Close()
}
func (h *Hellswitch) Close() {
	h.locker.Lock()
	defer h.locker.Unlock()
	close(h.c)
}
func (h *Hellswitch) reconnect(t *time.Ticker) {
	for {
		select {
		case <-t.C:
			h.Start()
		case <-h.c:
			t.Stop()
			return
		}
	}
}

func New() *Hellswitch {
	h := &Hellswitch{}
	h.ticker = time.NewTicker(ReconnectDuration)
	h.c = make(chan struct{})
	go h.reconnect(h.ticker)
	return h
}
