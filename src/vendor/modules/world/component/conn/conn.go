package conn

import (
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
	telnet   *telnet.Conn
	c        chan int
	running  bool
	buffer   []byte
	Lock     sync.RWMutex
	Debounce *debounce.Debounce
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

	b.DoSendToServer = conn.Send
	b.DoConnectServer = conn.Connect
	b.DoCloseServer = conn.Close
	b.GetConnBuffer = conn.Buffer
	b.GetConnConnected = conn.Connected
}

func (conn *Conn) UninstallFrom(b *bus.Bus) {
}

func (conn *Conn) C() chan int {
	return conn.c
}
func (conn *Conn) UpdatePrompt(bus *bus.Bus) {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	bus.HandleConnPrompt(bus, conn.buffer)
}

//Connect :connect to mud
func (conn *Conn) Connect(bus *bus.Bus) error {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	if conn.running == true {
		return nil
	}
	t, err := telnet.Dial("tcp", bus.GetHost(bus)+":"+bus.GetPort(bus))
	if err != nil {
		return err
	}
	conn.running = true
	conn.c = make(chan int)
	conn.buffer = make([]byte, 1024)
	conn.telnet = t
	go conn.Receiver(bus)
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
	close(conn.c)
	go bus.RaiseDiscontectedEvent()
	err := conn.telnet.Close()
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
			if isClosedError(err) {
				conn.Close(bus)
				return
			}
			bus.HandleConnError(bus, err)
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
				bus.HandleConnError(bus, err)
				return
			}
			bus.HandleConnReceive(bus, conn.buffer)
			conn.buffer = []byte{}
			conn.Lock.Unlock()
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
func (conn *Conn) Send(bus *bus.Bus, cmd []byte) error {
	if conn.telnet == nil {
		return nil
	}
	_, err := conn.telnet.Conn.Write(cmd)
	if err != nil {
		return err
	}
	_, err = conn.telnet.Conn.Write([]byte("\n"))
	return err
}

func New(host string) *Conn {
	c := &Conn{
		telnet:  nil,
		c:       make(chan int),
		running: false,
	}
	return c
}
