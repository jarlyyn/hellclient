package converter

import (
	"modules/world/bus"
	"sync"
)

type Converter struct {
	Lock sync.RWMutex
}

func (c *Converter) InstallTo(b *bus.Bus) {
	b.DoSend = c.Send
	b.HandleConnReceive = c.onMsg
	b.HandleConnPrompt = c.onPrompt
	b.DoPrint = c.DoPrint
	b.DoPrintSystem = c.DoPrintSystem
}

func (c *Converter) UninstallFrom(b *bus.Bus) {
}
func (c *Converter) onPrompt(bus *bus.Bus, msg []byte) {
	line := c.ConvertToLine(bus, msg)
	bus.RaisePromptEvent(line)
}
func (c *Converter) onMsg(bus *bus.Bus, msg []byte) {
	if len(msg) == 0 {
		return
	}
	line := c.ConvertToLine(bus, msg)
	bus.RaiseLineEvent(line)
}
func (c *Converter) onError(bus *bus.Bus, err error) {
	bus.HandleConverterError(bus, err)
}

func (c *Converter) Send(bus *bus.Bus, cmd []byte) error {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	b, err := FromUTF8(bus.GetCharset(bus), []byte(cmd))
	if err != nil {
		return err
	}
	return bus.DoSendToServer(bus, b)
}

func (c *Converter) DoPrintSystem(b *bus.Bus, msg string) {
	line := bus.NewLine()
	line.IsSystem = true
	w := bus.Word{
		Text: msg,
	}
	line.Append(w)
	b.RaiseLineEvent(line)
}

func (c *Converter) DoPrint(b *bus.Bus, msg string) {
	line := bus.NewLine()
	line.IsPrint = true
	w := bus.Word{
		Text: msg,
	}
	line.Append(w)
	b.RaiseLineEvent(line)
}

func (c *Converter) ConvertToLine(bus *bus.Bus, msg []byte) *bus.Line {
	charset := bus.GetCharset(bus)
	return ConvertToLine(msg, charset, func(err error) { c.onError(bus, err) })
}

func New() *Converter {
	return &Converter{}
}
