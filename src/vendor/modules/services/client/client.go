package client

import (
	"container/ring"
	"modules/services/conn"
	"sync"
	"time"
)

type World struct {
	Host     string
	Port     string
	Charset  string
	MaxLines int
}

type Word struct {
	Text string
	Cmd  string
}

type Line struct {
	Words   []Word
	Time    int64
	IsPrint bool
}

func NewLine(isPrint bool) *Line {
	return &Line{
		Words:   []Word{},
		Time:    time.Now().Unix(),
		IsPrint: isPrint,
	}
}

type Client struct {
	ID     string
	World  World
	Conn   *conn.Conn
	Lock   sync.RWMutex
	Lines  *ring.Ring
	Prompt Line
}

func (c *Client) NewLine(line *Line) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Lines.Value = line
	c.Lines = c.Lines.Next()
}
func (c *Client) ConvertLines() []*Line {
	result := make([]*Line, c.Lines.Len())
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	var i = 0
	c.Lines.Do(func(x interface{}) {
		result[i] = x.(*Line)
	})
	return result
}
