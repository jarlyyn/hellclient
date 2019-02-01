package conn

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/herb-go/misc/debounce"

	"github.com/ziutek/telnet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

const DefaultDebounceDuration = 200 * time.Millisecond

//Conn :mud conn
type Conn struct {
	host      string
	charset   string
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

func New(host string, charset string) *Conn {
	c := &Conn{
		host:    host,
		charset: charset,
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
	b, err := ToUTF8(conn.charset, conn.buffer)
	if err != nil {
		conn.OnError(err)
		return
	}
	conn.OnPrompt(b)
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
		if err == io.EOF {
			close(conn.c)
			return
		}
		if err != nil {
			conn.OnError(err)
			return
		}
		if s == del {
			b, err := ToUTF8(conn.charset, conn.buffer)
			if err != nil {
				conn.OnError(err)
				return
			}

			conn.OnReceive(b)
			conn.Lock.Lock()
			conn.buffer = []byte{}
			conn.Lock.Unlock()
		}
		conn.Lock.Lock()
		conn.buffer = append(conn.buffer, s)
		conn.Lock.Unlock()
		conn.Debounce.Exec()
	}
}
func (conn *Conn) Buffer() ([]byte, error) {
	conn.Lock.RLock()
	defer conn.Lock.RUnlock()
	return ToUTF8(conn.charset, conn.buffer)
}
func (conn *Conn) Send(cmd string) error {
	b, err := FromUTF8(conn.charset, []byte(cmd))
	if err != nil {
		return err
	}
	_, err = conn.telnet.Conn.Write(b)
	if err != nil {
		return err
	}
	_, err = conn.telnet.Conn.Write([]byte("\n"))
	return err
}

//ToUTF8 : convert from CJK encoding to UTF-8
func ToUTF8(from string, s []byte) ([]byte, error) {
	var reader *transform.Reader
	switch strings.ToLower(from) {
	case "gbk", "cp936", "windows-936":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	case "gb18030":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
	case "gb2312":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewDecoder())
	case "big5", "big-5", "cp950":
		reader = transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewDecoder())
	case "euc-kr", "euckr", "cp949":
		reader = transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewDecoder())
	case "euc-jp", "eucjp":
		reader = transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewDecoder())
	case "shift-jis":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewDecoder())
	case "iso-2022-jp", "cp932", "windows-31j":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewDecoder())
	default:
		return s, errors.New("Unsupported encoding " + from)
	}
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// FromUTF8 convert from UTF-8 encoding to CJK encoding
func FromUTF8(to string, s []byte) ([]byte, error) {
	var reader *transform.Reader
	switch strings.ToLower(to) {
	case "gbk", "cp936", "windows-936":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	case "gb18030":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder())
	case "gb2312":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	case "big5", "big-5", "cp950":
		reader = transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewEncoder())
	case "euc-kr", "euckr", "cp949":
		reader = transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewEncoder())
	case "euc-jp", "eucjp":
		reader = transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewEncoder())
	case "shift-jis":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewEncoder())
	case "iso-2022-jp", "cp932", "windows-31j":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewEncoder())
	default:
		return s, errors.New("Unsupported encoding " + to)
	}
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
