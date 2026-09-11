package converter

import (
	"modules/world"
	"modules/world/bus"
	"sync"
)

type Converter struct {
	SendLock     sync.RWMutex
	InputLock    sync.RWMutex
	Last         *world.Word
	LastAnsi     string
	PendingLines []*world.AnsiLine
}

func nopOnError(err error) bool {
	return false
}

func (c *Converter) InstallTo(b *bus.Bus) {
	b.DoSend = b.WrapHandleSend(c.Send)
	b.HandleConnReceive = b.WrapHandleBytes(c.onMsg)
	b.HandleConnPrompt = b.WrapHandleBytes(c.onPrompt)
	b.DoPrint = b.WrapHandleString(c.DoPrint)
	b.DoPrintSystem = b.WrapHandleString(c.DoPrintSystem)
	b.DoPrintLocalBroadcastIn = b.WrapHandleString(c.DoPrintLocalBroadcastIn)
	b.DoPrintGlobalBroadcastIn = b.WrapHandleString(c.DoPrintGlobalBroadcastIn)
	b.DoPrintLocalBroadcastOut = b.WrapHandleString(c.DoPrintLocalBroadcastOut)
	b.DoPrintGlobalBroadcastOut = b.WrapHandleString(c.DoPrintGlobalBroadcastOut)
	b.DoPrintSubneg = b.WrapHandleString(c.DoPrintSubneg)
	b.DoPrintRequest = b.WrapHandleString(c.DoPrintRequest)
	b.DoPrintResponse = b.WrapHandleString(c.DoPrintResponse)
	b.InsertAnsi = b.WrapHandleString(c.InsertAnsi)
	b.LastAnsi = c.GetLastAnsi
	b.ResetConverter = c.Reset
}
func (c *Converter) ExecLines(bus *bus.Bus) {
	for len(c.PendingLines) > 0 {
		line := c.PendingLines[0]
		newSlice := make([]*world.AnsiLine, len(c.PendingLines)-1)
		copy(newSlice, c.PendingLines[1:])
		c.PendingLines = newSlice
		c.LastAnsi = line.Ansi
		bus.RaiseLineEvent(line.Line)
	}
}
func (c *Converter) GetLastAnsi() string {
	return c.LastAnsi
}
func (c *Converter) InsertAnsi(bus *bus.Bus, msg string) {
	var needStart = len(c.PendingLines) == 0
	line := c.ConvertToLine(bus, msg, func(err error) bool { return c.onError(bus, err) })
	if line != nil {
		line.Type = world.LineTypeReal
		c.PendingLines = append(c.PendingLines, world.NewAnsiLine(line, msg))
		if needStart {
			c.ExecLines(bus)
		}
	}
}
func (c *Converter) onPrompt(bus *bus.Bus, data []byte) {
	c.InputLock.Lock()
	defer c.InputLock.Unlock()
	charset := bus.GetCharset()
	var msg, err = world.ToUTF8(charset, data)
	if err != nil {
		c.onError(bus, err)
		return
	}
	line := c.ConvertToLine(bus, string(msg), nopOnError)
	if line != nil {
		bus.RaisePromptEvent(line)
	}
}
func (c *Converter) onMsg(bus *bus.Bus, data []byte) {
	c.InputLock.Lock()
	defer c.InputLock.Unlock()
	if len(data) == 0 {
		return
	}
	charset := bus.GetCharset()
	var msg, err = world.ToUTF8(charset, data)
	if err != nil {
		c.onError(bus, err)
		return
	}
	c.InsertAnsi(bus, string(msg))
}
func (c *Converter) onError(bus *bus.Bus, err error) bool {
	bus.HandleConverterError(err)
	return true
}

func (c *Converter) Send(bus *bus.Bus, cmd *world.Command) {
	c.SendLock.Lock()
	defer c.SendLock.Unlock()
	if cmd.Message == "\x0f" {
		return
	}
	if bus.HandleSend(cmd.Message) {
		return
	}
	b, err := world.FromUTF8(bus.GetCharset(), []byte(cmd.Message))
	if err != nil {
		bus.HandleConverterError(err)
		return
	}
	if cmd.Echo {
		c.DoPrintEcho(bus, cmd)
	}
	if cmd.History {
		bus.AddHistory(cmd.Message)
	}
	bus.DoSendToConn(b)
	bus.DoSendToConn([]byte("\n"))
}
func (c *Converter) DoPrintEcho(b *bus.Bus, cmd *world.Command) {
	line := world.NewLine()
	line.Creator = cmd.Creator
	line.CreatorType = cmd.CreatorType
	line.Type = world.LineTypeEcho
	w := &world.Word{
		Text: cmd.Message,
	}
	line.Append(w)
	b.RaiseLineEvent(line)
}
func (c *Converter) DoPrintRequest(b *bus.Bus, msg string) {
	c.print(b, world.LineTypeRequest, msg)
}
func (c *Converter) DoPrintResponse(b *bus.Bus, msg string) {
	c.print(b, world.LineTypeResponse, msg)
}
func (c *Converter) DoPrintLocalBroadcastIn(b *bus.Bus, msg string) {
	c.print(b, world.LineTypeLocalBroadcastIn, msg)
}
func (c *Converter) DoPrintGlobalBroadcastIn(b *bus.Bus, msg string) {
	c.print(b, world.LineTypeGlobalBroadcastIn, msg)
}
func (c *Converter) DoPrintLocalBroadcastOut(b *bus.Bus, msg string) {
	c.print(b, world.LineTypeLocalBroadcastOut, msg)
}
func (c *Converter) DoPrintGlobalBroadcastOut(b *bus.Bus, msg string) {
	c.print(b, world.LineTypeGlobalBroadcastOut, msg)
}
func (c *Converter) DoPrintSubneg(b *bus.Bus, msg string) {
	c.print(b, world.LineTypeSubneg, msg)
}

func (c *Converter) DoPrintSystem(b *bus.Bus, msg string) {
	if b != nil {
		c.print(b, world.LineTypeSystem, msg)
	}
}

func (c *Converter) DoPrint(b *bus.Bus, msg string) {
	if b != nil {
		c.print(b, world.LineTypePrint, msg)
	}
}
func (c *Converter) print(b *bus.Bus, linetype int, msg string) {
	line := world.NewLine()
	line.Type = linetype
	w := &world.Word{
		Text: msg,
	}
	line.Append(w)
	b.RaiseLineEvent(line)

}
func (c *Converter) ConvertToLine(bus *bus.Bus, msg string, onError func(err error) bool) *world.Line {
	l, last := ConvertToLine(c.Last, []byte(msg), onError)
	c.Last = last
	return l
}
func (c *Converter) Reset() {
	c.PendingLines = []*world.AnsiLine{}
	c.LastAnsi = ""
	c.Last = nil
}
func New() *Converter {
	return &Converter{
		PendingLines: []*world.AnsiLine{},
	}
}
