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
	bus      *bus.Bus
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
	conn.bus = b
	b.DoSendToServer = conn.Send
	b.DoConnectServer = conn.Connect
	b.DoCloseServer = conn.Close
	b.GetConnBuffer = conn.Buffer
}

func (conn *Conn) UninstallFrom(b *bus.Bus) {
	if conn.bus != b {
		return
	}
	b.DoSendToServer = nil
	b.DoConnectServer = nil
	b.DoCloseServer = nil
	b.GetConnBuffer = nil
}

func (conn *Conn) C() chan int {
	return conn.c
}
func (conn *Conn) UpdatePrompt() {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	conn.bus.OnConnPrompt(conn.buffer)
}

//Connect :connect to mud
func (conn *Conn) Connect() error {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	if conn.running == true {
		return nil
	}
	t, err := telnet.Dial("tcp", conn.bus.GetHost())
	if err != nil {
		return err
	}
	conn.running = true
	conn.c = make(chan int)
	conn.buffer = make([]byte, 1024)
	conn.telnet = t
	go conn.Receiver()
	return nil
}

//Close :close mud conn
func (conn *Conn) Close() error {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	if conn.running == false {
		return nil
	}
	conn.running = false
	conn.buffer = []byte{}
	close(conn.c)

	err := conn.telnet.Close()
	return err
}
func (conn *Conn) Receiver() {
	del := byte(10)
	del2 := byte(13)
	for {
		running := conn.Running()
		if !running {
			return
		}
		s, err := conn.telnet.ReadByte()
		if err != nil {
			if isClosedError(err) {
				conn.Close()
				return
			}
			conn.bus.OnConnError(err)
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
				conn.bus.OnConnError(err)
				return
			}
			conn.bus.OnConnReceive(conn.buffer)
			conn.buffer = []byte{}
			conn.Lock.Unlock()
			continue
		}
		conn.buffer = append(conn.buffer, s)
		conn.Lock.Unlock()
		conn.Debounce.Exec()
	}
}
func (conn *Conn) Running() bool {
	if conn == nil {
		return false
	}
	conn.Lock.Lock()
	defer conn.Lock.Unlock()

	return conn.running
}
func (conn *Conn) Buffer() []byte {
	conn.Lock.RLock()
	defer conn.Lock.RUnlock()
	return conn.buffer
}
func (conn *Conn) Send(cmd []byte) error {
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
	d := debounce.New(DefaultDebounceDuration, c.UpdatePrompt)
	d.MaxDuration = 0
	c.Debounce = d
	return c
}
