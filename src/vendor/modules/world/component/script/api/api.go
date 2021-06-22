package api

import (
	"modules/world/bus"
)

type API struct {
	Bus *bus.Bus
}

func (a *API) Note(info string) {
	a.Bus.DoPrint(info)
}
func (a *API) SendImmediate(info string) {
	a.Bus.DoSend([]byte(info), true)
}
func (a *API) Send(info string) {
	a.Bus.DoSendToQueue([]byte(info))
}
