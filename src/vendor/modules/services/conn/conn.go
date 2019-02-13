package conn

import (
	"io"
	"net"
	"sync"
	"time"

	"github.com/herb-go/misc/debounce"

	"github.com/ziutek/telnet"
)

const DefaultDebounceDuration = 200 * time.Millisecond

//Conn :mud conn
type Conn struct {
	host      string
	telnet    *telnet.Conn
	c         chan int
	running   bool
	OnReceive func(msg []byte)
	OnError   func(err error)
	OnPrompt  func(msg []byte)
	buffer    []byte
	Lock      sync.RWMutex
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

func New(host string) *Conn {
	c := &Conn{
		host:    host,
		telnet:  nil,
		c:       make(chan int),
		running: false,
	}
	d := debounce.New(DefaultDebounceDuration, c.UpdatePrompt)
	d.MaxDuration = 0
	c.Debounce = d
	return c
}
func (conn *Conn) C() chan int {
	return conn.c
}
func (conn *Conn) UpdatePrompt() {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	conn.OnPrompt(conn.buffer)
}

//Connect :connect to mud
func (conn *Conn) Connect() error {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	if conn.running == true {
		return nil
	}
	t, err := telnet.Dial("tcp", conn.host)
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
	// del2 := byte(27)
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
			conn.OnError(err)
			return
		}
		conn.Lock.Lock()
		if s == del {
			if err != nil {
				conn.Lock.Unlock()
				conn.OnError(err)
				return
			}
			conn.OnReceive(conn.buffer)
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
