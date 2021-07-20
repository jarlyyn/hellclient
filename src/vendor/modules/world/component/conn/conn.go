package conn

import (
	"container/list"
	"io"
	"modules/world/bus"
	"net"
	"sync"
	"time"

	"github.com/herb-go/misc/debounce"
	"github.com/ziutek/telnet"
)

const DefaultDebounceDuration = 200 * time.Millisecond

//Conn :mud conn
type Conn struct {
	telnet    *telnet.Conn
	c         chan int
	running   bool
	buffer    []byte
	Lock      sync.RWMutex
	SendLock  sync.RWMutex
	sendQueue *list.List
	Debounce  *debounce.Debounce
}

func isClosedError(err error) bool {
	if err == io.EOF {
		return true
	}
	if operr, ok := err.(*net.OpError); ok {
		if operr.Err.Error() == "use of closed network connection" {
			return true
		}
	}
	return false
}

func (conn *Conn) InstallTo(b *bus.Bus) {
	d := debounce.New(DefaultDebounceDuration, func() { conn.UpdatePrompt(b) })
	d.MaxDuration = 0
	conn.Debounce = d

	b.DoSendToConn = b.WrapHandleBytes(conn.Send)
	b.DoConnectServer = b.WrapDo(conn.Connect)
	b.DoCloseServer = b.WrapDo(conn.Close)
	// b.GetConnBuffer = b.WrapGetString(conn.Buffer)
	b.GetConnConnected = b.WrapGetBool(conn.Connected)
	b.BindCloseEvent(conn, conn.Stop)
}

func (conn *Conn) C() chan int {
	return conn.c
}
func (conn *Conn) UpdatePrompt(bus *bus.Bus) {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	if conn.running {
		go bus.HandleConnPrompt(conn.buffer)
	}
}
func (conn *Conn) Stop(b *bus.Bus) {
	conn.Close(b)
}

//Connect :connect to mud
func (conn *Conn) Connect(bus *bus.Bus) error {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	if conn.running == true {
		return nil
	}
	t, err := telnet.Dial("tcp", bus.GetHost()+":"+bus.GetPort())
	if err != nil {
		return err
	}
	conn.running = true
	conn.c = make(chan int)
	conn.buffer = make([]byte, 1024)
	conn.telnet = t
	go conn.Receiver(bus)
	go bus.RaiseConnectedEvent()
	return nil
}

//Close :close mud conn
func (conn *Conn) Close(bus *bus.Bus) error {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	if conn.running == false {
		return nil
	}
	conn.running = false
	conn.buffer = []byte{}
	go bus.HandleConnPrompt(conn.buffer)
	go conn.Debounce.Discard()
	close(conn.c)
	err := conn.telnet.Close()
	conn.telnet = nil
	go bus.RaiseDisconnectedEvent()

	return err
}
func (conn *Conn) Receiver(bus *bus.Bus) {
	del := byte(10)
	del2 := byte(13)
	for {
		running := conn.Connected(bus)
		if !running {
			return
		}
		s, err := conn.telnet.ReadByte()
		if err != nil {
			if !isClosedError(err) {
				bus.HandleConnError(err)
			}
			conn.Lock.Lock()
			if conn.running == true {
				go bus.RaiseServerCloseEvent()
			}
			conn.Lock.Unlock()
			conn.Close(bus)

			return
		}
		conn.Lock.Lock()
		if s == del2 {
			conn.Lock.Unlock()
			continue
		}
		if s == del {
			if err != nil {
				conn.Lock.Unlock()
				bus.HandleConnError(err)
				return
			}
			buf := conn.buffer
			conn.Debounce.Reset()
			conn.buffer = []byte{}
			conn.Lock.Unlock()
			bus.HandleConnReceive(buf)
			continue
		}
		conn.buffer = append(conn.buffer, s)
		conn.Lock.Unlock()
		conn.Debounce.Exec()
	}
}
func (conn *Conn) Connected(bus *bus.Bus) bool {
	if conn == nil {
		return false
	}
	conn.Lock.Lock()
	defer conn.Lock.Unlock()

	return conn.running
}
func (conn *Conn) Buffer(bus *bus.Bus) []byte {
	conn.Lock.RLock()
	defer conn.Lock.RUnlock()
	return conn.buffer
}
func (conn *Conn) Send(bus *bus.Bus, cmd []byte) {
	conn.SendLock.Lock()
	defer conn.SendLock.Unlock()
	len := conn.sendQueue.Len()
	conn.sendQueue.PushBack(cmd)
	if len == 0 {
		go conn.sendingLoop(bus)
	}

}
func (conn *Conn) sendingLoop(bus *bus.Bus) {
	for {
		conn.SendLock.RLock()
		len := conn.sendQueue.Len()
		if len == 0 {
			conn.SendLock.RUnlock()
			return
		}
		conn.SendLock.RUnlock()
		conn.sending(bus)
	}
}
func (conn *Conn) sending(bus *bus.Bus) {
	conn.SendLock.Lock()
	defer conn.SendLock.Unlock()
	if conn.telnet == nil {
		conn.sendQueue = list.New()
		return
	}
	if conn.sendQueue.Len() > 0 {
		el := conn.sendQueue.Front()
		conn.sendQueue.Remove(el)
		_, err := conn.telnet.Conn.Write(el.Value.([]byte))
		if err != nil {
			bus.HandleConnError(err)
		}
	}
}
func New() *Conn {
	c := &Conn{
		telnet:    nil,
		c:         make(chan int),
		running:   false,
		sendQueue: list.New(),
	}
	return c
}
