package client

import (
	"container/ring"
	"fmt"
	"modules/services/conn"
	"sync"
	"time"

	"github.com/jarlyyn/ansi"
)

type World struct {
	Host     string
	Port     string
	Charset  string
	MaxLines int
}

type Word struct {
	Text       string
	Color      string
	Background string
	Bold       bool
}

type Line struct {
	Words   []Word
	Time    int64
	IsPrint bool
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

func NewLine(isPrint bool) *Line {
	return &Line{
		Words:   []Word{},
		Time:    time.Now().Unix(),
		IsPrint: isPrint,
	}
}
func New() *Client {
	return &Client{}
}

type Client struct {
	ID      string
	Manager *Manager
	World   World
	Conn    *conn.Conn
	Lock    sync.RWMutex
	Lines   *ring.Ring
	Prompt  Line
	Exit    chan int
}

func (c *Client) Init() {
	c.Lines = ring.New(1000)
}
func (c *Client) ConvertToLine(msg []byte) *Line {
	w := Word{}
	line := NewLine(false)
	for {
		out, s, err := ansi.Decode(msg)
		msg = out
		if s != nil && s.Type == "" {
			w.Text = string(s.Code)
			line.Append(w)
		}
		if s != nil && s.Type == "CSI" {
			// fmt.Println("CSI", s.Params)
			for _, v := range s.Params {
				switch v {
				case "0":
					{
						w.Color = ""
						w.Background = ""
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
						w.Background = "Black"
					}
				case "41":
					{
						w.Background = "Red"
					}
				case "42":
					{
						w.Background = "Green"
					}
				case "43":
					{
						w.Background = "Yellow"
					}
				case "44":
					{
						w.Background = "Blue"
					}
				case "45":
					{
						w.Background = "Magenta"
					}
				case "46":
					{
						w.Background = "Cyan"
					}
				case "47":
					{
						w.Background = "White"
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
						w.Background = "Bright-Black"
					}
				case "101":
					{
						w.Background = "Bright-Red"
					}
				case "102":
					{
						w.Background = "Bright-Green"
					}
				case "103":
					{
						w.Background = "Bright-Yellow"
					}
				case "104":
					{
						w.Background = "Bright-Blue"
					}
				case "105":
					{
						w.Background = "Bright-Magenta"
					}
				case "106":
					{
						w.Background = "Bright-Cyan"
					}
				case "107":
					{
						w.Background = "Bright-White"
					}
				}
			}

			// fmt.Println(string(msg))
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
	c.Manager.OnPrompt(c.ID, line)
}

func (c *Client) onMsg(msg []byte) {
	if len(msg) == 0 {
		return
	}
	line := c.ConvertToLine(msg)
	c.NewLine(line)
	c.Manager.OnLine(c.ID, line)

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

func (c *Client) Connect() error {
	c.Lock.Lock()
	c.Lock.Unlock()
	if c.Conn == nil {
		c.Conn = conn.New(c.World.Host+":"+c.World.Port, c.World.Charset)
		c.Conn.OnReceive = c.onMsg
		c.Conn.OnError = c.onError
		c.Conn.OnPrompt = c.onPrompt
	}
	return c.Conn.Connect()
}

func init() {
	ansi.FlagIgnoreC1 = true
}
