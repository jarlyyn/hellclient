package client

import (
	"bytes"
	"container/ring"
	"errors"
	"fmt"
	"io/ioutil"
	"modules/app"
	"modules/services/conn"
	"modules/services/script"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/herb-go/herbgo/util/config/tomlconfig"

	"github.com/jarlyyn/ansi"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

type World struct {
	Host    string
	Port    string
	Charset string
}

type Word struct {
	Text       string
	Color      string
	Background string
	Bold       bool
}

type Line struct {
	Words    []Word
	Time     int64
	IsPrint  bool
	IsSystem bool
}

func (l *Line) Append(w Word) {
	l.Words = append(l.Words, w)
}
func (l *Line) Plain() string {
	var output string
	for k := range l.Words {
		output = output + l.Words[k].Text
	}
	return output
}

func NewLine() *Line {
	return &Line{
		Words:    []Word{},
		Time:     time.Now().Unix(),
		IsPrint:  false,
		IsSystem: false,
	}
}
func New() *Client {
	return &Client{
		Exit:   make(chan int),
		Script: script.New(),
	}
}

type ClientInfo struct {
	ID      string
	Running bool
}
type Client struct {
	ID      string
	Manager *Manager
	World   World
	Conn    *conn.Conn
	Lock    sync.RWMutex
	Lines   *ring.Ring
	Prompt  *Line
	Script  *script.Script
	Exit    chan int
}

func (c *Client) Save() error {
	path := filepath.Join(WorldsPath, c.ID+".toml")
	return tomlconfig.Save(path, c.World)
}
func (c *Client) Info() *ClientInfo {
	return &ClientInfo{
		ID:      c.ID,
		Running: c.Conn.Running(),
	}
}
func (c *Client) Init() {
	c.Lines = ring.New(1000)
}
func (c *Client) ConvertToLine(msg []byte) *Line {
	w := Word{}
	line := NewLine()
	if len(msg) == 0 {
		return line
	}
	for {
		out, s, err := ansi.Decode(msg)
		msg = out
		if s != nil && s.Type == "" {
			b, err := ToUTF8(c.World.Charset, []byte(s.Code))
			if err != nil {
				c.onError(err)
				continue
			}
			w.Text = string(b)
			line.Append(w)
		}
		if s != nil && s.Type == "CSI" {
			// fmt.Println("CSI", s.Params)
			for _, v := range s.Params {
				switch v {
				case "0":
					{
						w.Color = ""
						w.Background = "BG-"
						w.Bold = false
					}
				case "1":
					{
						w.Bold = true
					}
				case "2":
					{
						w.Bold = false
					}
				case "30":
					{
						w.Color = "Black"
					}
				case "31":
					{
						w.Color = "Red"
					}
				case "32":
					{
						w.Color = "Green"
					}
				case "33":
					{
						w.Color = "Yellow"
					}
				case "34":
					{
						w.Color = "Blue"
					}
				case "35":
					{
						w.Color = "Magenta"
					}
				case "36":
					{
						w.Color = "Cyan"
					}
				case "37":
					{
						w.Color = "White"
					}
				case "40":
					{
						w.Background = "BG-Black"
					}
				case "41":
					{
						w.Background = "BG-Red"
					}
				case "42":
					{
						w.Background = "BG-Green"
					}
				case "43":
					{
						w.Background = "BG-Yellow"
					}
				case "44":
					{
						w.Background = "BG-Blue"
					}
				case "45":
					{
						w.Background = "BG-Magenta"
					}
				case "46":
					{
						w.Background = "BG-Cyan"
					}
				case "47":
					{
						w.Background = "BG-White"
					}
				case "90":
					{
						w.Color = "Bright-Black"
					}
				case "91":
					{
						w.Color = "Bright-Red"
					}
				case "92":
					{
						w.Color = "Bright-Green"
					}
				case "93":
					{
						w.Color = "Bright-Yellow"
					}
				case "94":
					{
						w.Color = "Bright-Blue"
					}
				case "95":
					{
						w.Color = "Bright-Magenta"
					}
				case "96":
					{
						w.Color = "Bright-Cyan"
					}
				case "97":
					{
						w.Color = "Bright-White"
					}
				case "100":
					{
						w.Background = "BG-Bright-Black"
					}
				case "101":
					{
						w.Background = "BG-Bright-Red"
					}
				case "102":
					{
						w.Background = "BG-Bright-Green"
					}
				case "103":
					{
						w.Background = "BG-Bright-Yellow"
					}
				case "104":
					{
						w.Background = "BG-Bright-Blue"
					}
				case "105":
					{
						w.Background = "BG-Bright-Magenta"
					}
				case "106":
					{
						w.Background = "BG-Bright-Cyan"
					}
				case "107":
					{
						w.Background = "BG-Bright-White"
					}
				case "256":
					line = NewLine()

				}
			}
		}

		if err != nil {
			c.onError(err)
			// return
		}
		if msg == nil {
			break
		}
		if (len(msg)) == 0 {
			break
		}

	}
	return line
}
func (c *Client) onPrompt(msg []byte) {
	line := c.ConvertToLine(msg)
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Prompt = line
	c.Manager.OnPrompt(c.ID, line)
}
func (c *Client) Match(line string) {
	c.Script.Triggers.Match(line)
}
func (c *Client) onMsg(msg []byte) {
	if len(msg) == 0 {
		return
	}
	line := c.ConvertToLine(msg)
	c.NewLine(line)
	c.Manager.OnLine(c.ID, line)
	c.Match(line.Plain())
}
func (c *Client) onError(err error) {
	fmt.Println(err.Error())
}
func (c *Client) NewLine(line *Line) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Lines.Value = line
	c.Lines = c.Lines.Next()
}
func (c *Client) ConvertLines() []*Line {
	result := []*Line{}
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	c.Lines.Do(func(x interface{}) {
		line, ok := x.(*Line)
		if ok && line != nil {
			result = append(result, line)
		}
	})
	return result
}

func (c *Client) SendSystem(msg string) {
	line := NewLine()
	line.IsSystem = true
	w := Word{
		Text: msg,
	}
	line.Append(w)
	c.NewLine(line)
	c.Manager.OnLine(c.ID, line)
}

func (c *Client) Print(msg string) {
	line := NewLine()
	line.IsPrint = true
	w := Word{
		Text: msg,
	}
	line.Append(w)
	c.NewLine(line)
	c.Manager.OnLine(c.ID, line)
}
func (c *Client) Disconnect() error {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if c.Conn == nil {
		return nil
	}
	return c.Conn.Close()
}
func (c *Client) Connect() error {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if c.Conn == nil {
		c.Conn = conn.New(c.World.Host + ":" + c.World.Port)
		c.Conn.OnReceive = c.onMsg
		c.Conn.OnError = c.onError
		c.Conn.OnPrompt = c.onPrompt
	}
	err := c.Conn.Connect()
	if err == nil {
		go func() {
			c.Manager.OnConnected(c.ID)
			c.SendSystem(app.Time.Datetime(time.Now()) + "  成功连接服务器")
		}()
	}
	go func() {
		<-c.Conn.C()
		c.Manager.OnDisconnected(c.ID)
		c.SendSystem(app.Time.Datetime(time.Now()) + "  与服务器断开连接接 ")
	}()
	return err
}
func (c *Client) Send(cmd []byte) error {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if c.Conn == nil {
		return nil
	}
	b, err := FromUTF8(c.World.Charset, []byte(cmd))
	if err != nil {
		return err
	}
	return c.Conn.Send(b)
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

func init() {
	ansi.FlagIgnoreC1 = true
}
