package conn

import (
	"io"
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
	Running   bool
	OnReceive func(msg []byte)
	OnError   func(err error)
	OnPrompt  func(msg []byte)
	buffer    []byte
	Lock      sync.RWMutex
	Debounce  *debounce.Debounce
}

func New(host string) *Conn {
	c := &Conn{
		host:    host,
		telnet:  nil,
		c:       make(chan int),
		Running: false,
	}
	d := debounce.New(DefaultDebounceDuration, c.UpdatePrompt)
	d.MaxDuration = 0
	c.Debounce = d
	return c
}
func (conn *Conn) UpdatePrompt() {
	conn.Lock.Lock()
	defer conn.Lock.Unlock()
	// b, err := ToUTF8(conn.charset, conn.buffer)
	conn.OnPrompt(conn.buffer)
}

//Connect :connect to mud
func (conn *Conn) Connect() error {
	t, err := telnet.Dial("tcp", conn.host)
	if err != nil {
		return err
	}
	conn.buffer = make([]byte, 1024)
	conn.telnet = t
	go conn.Receiver()
	return nil
}

//Close :close mud conn
func (conn *Conn) Close() error {
	close(conn.c)
	conn.Running = false
	err := conn.telnet.Close()
	return err
}
func (conn *Conn) Receiver() {
	del := byte(10)
	// del2 := byte(27)
	for {
		s, err := conn.telnet.ReadByte()
		conn.Lock.Lock()
		if err == io.EOF {
			return
		}
		if err != nil {
			conn.OnError(err)
			return
		}
		if s == del {
			if err != nil {
				conn.OnError(err)
				conn.Lock.Unlock()
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
func (conn *Conn) Buffer() []byte {
	conn.Lock.RLock()
	defer conn.Lock.RUnlock()
	return conn.buffer
}
func (conn *Conn) Send(cmd []byte) error {
	_, err := conn.telnet.Conn.Write(cmd)
	if err != nil {
		return err
	}
	_, err = conn.telnet.Conn.Write([]byte("\n"))
	return err
}
