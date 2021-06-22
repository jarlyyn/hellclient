package converter

import (
	"modules/world"
	"modules/world/bus"
	"sync"
)

type Converter struct {
	Lock sync.RWMutex
}

func (c *Converter) InstallTo(b *bus.Bus) {
	b.DoSend = b.WrapHandleBytes(c.Send)
	b.HandleConnReceive = b.WrapHandleBytes(c.onMsg)
	b.HandleConnPrompt = b.WrapHandleBytes(c.onPrompt)
	b.DoPrint = b.WrapHandleString(c.DoPrint)
	b.DoPrintSystem = b.WrapHandleString(c.DoPrintSystem)
}

func (c *Converter) UninstallFrom(b *bus.Bus) {
}
func (c *Converter) onPrompt(bus *bus.Bus, msg []byte) {
	line := c.ConvertToLine(bus, msg)
	if line != nil {
		bus.RaisePromptEvent(line)
	}
}
func (c *Converter) onMsg(bus *bus.Bus, msg []byte) {
	if len(msg) == 0 {
		return
	}
	line := c.ConvertToLine(bus, msg)
	if line != nil {
		line.IsReal = true
		bus.RaiseLineEvent(line)
	}
}
func (c *Converter) onError(bus *bus.Bus, err error) {
	bus.HandleConverterError(err)
}

func (c *Converter) Send(bus *bus.Bus, cmd []byte) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	b, err := FromUTF8(bus.GetCharset(), cmd)
	if err != nil {
		bus.HandleConverterError(err)
	}

	bus.DoSendToConn(b)
	bus.DoSendToConn([]byte("\n"))
}

func (c *Converter) DoPrintSystem(b *bus.Bus, msg string) {
	line := world.NewLine()
	line.IsSystem = true
	w := world.Word{
		Text: msg,
	}
	line.Append(w)
	b.RaiseLineEvent(line)
}

func (c *Converter) DoPrint(b *bus.Bus, msg string) {
	line := world.NewLine()
	line.IsPrint = true
	w := world.Word{
		Text: msg,
	}
	line.Append(w)
	b.RaiseLineEvent(line)
}

func (c *Converter) ConvertToLine(bus *bus.Bus, msg []byte) *world.Line {
	charset := bus.GetCharset()
	return ConvertToLine(msg, charset, func(err error) { c.onError(bus, err) })
}

func New() *Converter {
	return &Converter{}
}
