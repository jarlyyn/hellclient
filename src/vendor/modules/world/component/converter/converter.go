package converter

import (
	"modules/world/bus"
	"sync"
)

type Converter struct {
	bus  *bus.Bus
	Lock sync.RWMutex
}

func (c *Converter) InstallTo(b *bus.Bus) {
	c.bus = b
	b.DoSend = c.Send
	b.HandleConnReceive = c.onMsg
	b.HandleConnPrompt = c.onPrompt
	b.DoPrint = c.DoPrint
	b.DoPrintSystem = c.DoPrintSystem
}

func (c *Converter) UninstallFrom(b *bus.Bus) {
	if c.bus != b {
		return
	}
}
func (c *Converter) onPrompt(msg []byte) {
	line := c.ConvertToLine(msg)
	c.bus.RaisePromptEvent(line)
}
func (c *Converter) onMsg(msg []byte) {
	if len(msg) == 0 {
		return
	}
	line := c.ConvertToLine(msg)
	c.bus.RaiseLineEvent(line)
}
func (c *Converter) onError(err error) {
	c.bus.HandleConverterError(err)
}

func (c *Converter) Send(cmd []byte) error {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	b, err := FromUTF8(c.bus.GetCharset(), []byte(cmd))
	if err != nil {
		return err
	}
	return c.bus.DoSendToServer(b)
}

func (c *Converter) DoPrintSystem(msg string) {
	line := bus.NewLine()
	line.IsSystem = true
	w := bus.Word{
		Text: msg,
	}
	line.Append(w)
	c.bus.RaiseLineEvent(line)
}

func (c *Converter) DoPrint(msg string) {
	line := bus.NewLine()
	line.IsPrint = true
	w := bus.Word{
		Text: msg,
	}
	line.Append(w)
	c.bus.RaiseLineEvent(line)
}

func (c *Converter) ConvertToLine(msg []byte) *bus.Line {
	charset := c.bus.GetCharset()
	return ConvertToLine(msg, charset, c.onError)
}

func New() *Converter {
	return &Converter{}
}
