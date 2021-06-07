package prophet

import (
	"github.com/herb-go/connections/room/message"
)

func initAdapter(p *Prophet, adapter *message.Adapter) {
	adapter.Register("line", p.newRoomAdapter("line"))
	adapter.Register("lines", p.newRoomAdapter("lines"))
	adapter.Register("prompt", p.newRoomAdapter("prompt"))
	adapter.Register("triggers", p.newRoomAdapter("triggers"))
	adapter.Register("clients", p.newUserAdapter("clients"))
	adapter.Register("connected", p.newUserAdapter("connected"))
	adapter.Register("disconnected", p.newUserAdapter("disconnected"))
	adapter.Register("createFail", p.newUserAdapter("createFail"))
	adapter.Register("createSuccess", p.newUserAdapter("createSuccess"))
	adapter.Register("triggerFail", p.newRoomAdapter("triggerFail"))
	adapter.Register("triggerSuccess", p.newRoomAdapter("triggerSuccess"))
	adapter.Register("allLines", p.newRoomAdapter("allLines"))
}