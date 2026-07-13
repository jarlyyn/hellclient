package conn

import (
	"io"
	"modules/app"
	"modules/world/bus"
	"net"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/herb-go/misc/debounce"
	"github.com/herb-go/util"
	"github.com/jarlyyn/telnet"
	"golang.org/x/net/proxy"
)

const TTYPE = byte(24)
const TTYPESend = byte(1)
const TTYPEIs = byte(0)

const TerminalType = "VT100"
const MTTS = "MTTS 7"
const DefaultDebounceDuration = 200 * time.Millisecond
const DefaultWriteTimeout = 5 * time.Second

// Conn :mud conn
type Conn struct {
	telnet      *telnet.Conn
	c           chan int
	running     atomic.Bool
	connecting  atomic.Bool
	buffer      []byte
	ConnectLock sync.RWMutex
	BufferLock  sync.RWMutex
	SendLock    sync.RWMutex
	Debounce    *debounce.Debounce
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
	d := debounce.New(DefaultDebounceDuration, func() { go conn.UpdatePrompt(b) })
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
	conn.ConnectLock.Lock()
	defer conn.ConnectLock.Unlock()
	if conn.running.Load() {
		bus.HandleConnPrompt(conn.buffer)
	}
}
func (conn *Conn) Stop(b *bus.Bus) {
	conn.Close(b)
	conn.Dispose()
}

// Connect :connect to mud
func (conn *Conn) Connect(bus *bus.Bus) error {
	if conn.connecting.Load() {
		return nil
	}
	conn.ConnectLock.Lock()
	defer conn.ConnectLock.Unlock()
	conn.connecting.Store(true)
	defer conn.connecting.Store(false)
	if conn.running.Load() {
		return nil
	}
	timeout := app.System.ConnectTimeout
	if timeout <= 0 {
		timeout = 1
	}
	proxydata := strings.TrimSpace(bus.GetProxy())
	var netconn net.Conn
	var err error
	if proxydata == "" {
		netconn, err = net.DialTimeout("tcp", bus.GetHost()+":"+bus.GetPort(), time.Duration(timeout)*time.Second)
		if err != nil {
			go bus.RaiseServerCloseEvent()
			return err
		}
	} else {
		proxyurl, err := url.Parse(proxydata)
		if err != nil {
			return err
		}
		dialer, err := proxy.FromURL(proxyurl, &net.Dialer{Timeout: time.Duration(timeout) * time.Second})
		if err != nil {
			return err
		}
		netconn, err = dialer.Dial("tcp", bus.GetHost()+":"+bus.GetPort())
		if err != nil {
			go bus.RaiseServerCloseEvent()
			return err
		}

	}
	t, err := telnet.NewConn(netconn)
	if err != nil {
		go bus.RaiseServerCloseEvent()
		return err
	}
	t.GMCP = true
	var ttype []string
	if app.System.TerminalType != "" {
		t.TerminalType = true
		ttype = []string{app.System.TerminalType, TerminalType, MTTS, MTTS}
	}
	t.OnGA = func() {
		conn.BufferLock.Lock()
		conn.flushBuffer(bus)
		bus.HandleBuffer(nil)
	}
	t.OnSubneg = func(data []byte) {
		if len(data) > 1 {
			if data[0] == TTYPE && data[1] == TTYPESend {
				conn.BufferLock.Lock()
				if len(ttype) > 0 {
					data = []byte{TTYPEIs}
					data = append(data, []byte(ttype[0])...)
					t.Sub(TTYPE, data...)
					ttype = ttype[1:]
				}
				conn.BufferLock.Unlock()
			}
		}
		bus.HandleSubneg(data)
	}
	conn.running.Store(true)

	conn.BufferLock.Lock()
	conn.c = make(chan int)
	conn.buffer = make([]byte, 0, 1024)
	conn.telnet = t
	conn.BufferLock.Unlock()

	go conn.Receiver(bus)
	go bus.RaiseConnectedEvent()
	return nil
}

// Close :close mud conn
func (conn *Conn) Close(bus *bus.Bus) error {
	conn.ConnectLock.Lock()
	defer conn.ConnectLock.Unlock()
	if !conn.running.Load() {
		return nil
	}
	conn.running.Store(false)
	conn.BufferLock.Lock()
	defer conn.BufferLock.Unlock()
	conn.SendLock.Lock()
	defer conn.SendLock.Unlock()
	buffer := conn.buffer
	conn.buffer = []byte{}
	close(conn.c)
	err := conn.telnet.Close()
	conn.telnet = nil

	go bus.HandleConnPrompt(buffer)
	go conn.Debounce.Discard()
	go bus.RaiseDisconnectedEvent()
	go bus.RaiseServerCloseEvent()

	return err
}
func (conn *Conn) flushBuffer(bus *bus.Bus) {
	buf := conn.buffer
	conn.Debounce.Reset()
	conn.buffer = []byte{}
	conn.BufferLock.Unlock()
	bus.HandleConnReceive(buf)
}
func (conn *Conn) Receiver(bus *bus.Bus) {
	del := byte(10)
	del2 := byte(13)
	nop := byte(0)
	for {
		running := conn.Connected(bus)
		if !running {
			return
		}
		var err error
		var s byte
		err2 := util.Catch(func() {
			s, err = conn.telnet.ReadByte()
		})
		if err2 != nil {
			err = err2
		}
		if err != nil {
			if !isClosedError(err) {
				bus.HandleConnError(err)
			}
			// conn.RunningLock.Lock()
			// if conn.running == true {
			// 	go bus.RaiseServerCloseEvent()
			// }
			// conn.RunningLock.Unlock()
			conn.Close(bus)

			return
		}
		conn.BufferLock.Lock()
		if s == del2 || s == nop {
			conn.BufferLock.Unlock()
			continue
		}
		if s == del {
			if err != nil {
				conn.BufferLock.Unlock()
				bus.HandleConnError(err)
				return
			}
			conn.flushBuffer(bus)
			continue
		}
		conn.buffer = append(conn.buffer, s)
		if bus.HandleBuffer(conn.buffer) {
			conn.flushBuffer(bus)
			continue
		}
		conn.BufferLock.Unlock()
		conn.Debounce.Exec()
	}
}
func (conn *Conn) Connected(bus *bus.Bus) bool {
	if conn == nil {
		return false
	}
	return conn.running.Load()
}
func (conn *Conn) Buffer(bus *bus.Bus) []byte {
	conn.BufferLock.RLock()
	defer conn.BufferLock.RUnlock()
	return conn.buffer
}
func (conn *Conn) Send(bus *bus.Bus, cmd []byte) {
	conn.SendLock.Lock()
	defer conn.SendLock.Unlock()
	conn.send(bus, cmd)
}

func (conn *Conn) send(bus *bus.Bus, cmd []byte) {
	if conn.telnet == nil {
		return
	}
	conn.telnet.SetWriteDeadline(time.Now().Add(DefaultWriteTimeout))
	_, err := conn.telnet.Conn.Write(cmd)
	if err != nil {
		bus.HandleConnError(err)
	}

}
func (conn *Conn) Dispose() {
	conn.Debounce.Discard()
}
func New() *Conn {
	c := &Conn{
		telnet: nil,
		c:      make(chan int),
	}
	c.running.Store(false)
	c.connecting.Store(false)
	return c
}
