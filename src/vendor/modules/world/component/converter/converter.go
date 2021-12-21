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
	b.DoSend = b.WrapHandleSend(c.Send)
	b.HandleConnReceive = b.WrapHandleBytes(c.onMsg)
	b.HandleConnPrompt = b.WrapHandleBytes(c.onPrompt)
	b.DoPrint = b.WrapHandleString(c.DoPrint)
	b.DoPrintSystem = b.WrapHandleString(c.DoPrintSystem)
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
		line.Type = world.LineTypeReal
		bus.RaiseLineEvent(line)
	}
}
func (c *Converter) onError(bus *bus.Bus, err error) {
	bus.HandleConverterError(err)
}

func (c *Converter) Send(bus *bus.Bus, cmd *world.Command) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	b, err := world.FromUTF8(bus.GetCharset(), []byte(cmd.Mesasge))
	if err != nil {
		bus.HandleConverterError(err)
		return
	}
	if cmd.Echo {
		c.DoPrintEcho(bus, cmd)
	}
	if cmd.History {
		bus.AddHistory(cmd.Mesasge)
	}
	bus.DoSendToConn(b)
	bus.DoSendToConn([]byte("\n"))
}
func (c *Converter) DoPrintEcho(b *bus.Bus, cmd *world.Command) {
	line := world.NewLine()
	line.Creator = cmd.Creator
	line.CreatorType = cmd.CreatorType
	line.Type = world.LineTypeEcho
	w := world.Word{
		Text: cmd.Mesasge,
	}
	line.Append(w)
	b.RaiseLineEvent(line)
}
func (c *Converter) DoPrintSystem(b *bus.Bus, msg string) {
	line := world.NewLine()
	line.Type = world.LineTypeSystem
	w := world.Word{
		Text: msg,
	}
	line.Append(w)
	b.RaiseLineEvent(line)
}

func (c *Converter) DoPrint(b *bus.Bus, msg string) {
	line := world.NewLine()
	line.Type = world.LineTypePrint
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
