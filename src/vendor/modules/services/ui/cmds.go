package ui

import (
	"encoding/json"
	"modules/services/client"
	"modules/services/ui/forms"

	"github.com/jarlyyn/herb-go-experimental/connections"
	"github.com/jarlyyn/herb-go-experimental/connections/room"

	"github.com/jarlyyn/herb-go-experimental/connections/command"
)

const CmdsChange = "change"

const CmdsConnect = "connect"

var handlers = command.NewHandlers()
var onCmdChange = func(conn connections.ConnectionOutput, cmd command.Command) error {
	ctx := CurretEngine.Context(conn.ID())
	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()
	v, ok := ctx.Data.Load("rooms")
	if ok == false {
		v, _ = ctx.Data.LoadOrStore("rooms", room.NewLocation(conn, rooms))
	}
	j := v.(*room.Location)
	j.Leave(current.Load().(string))
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	Change(msg)
	j.Join(msg)
	return Send(conn, "current", msg)
}
var onCmdConnect = func(conn connections.ConnectionOutput, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	client.DefaultManager.ExecConnect(msg)
	return nil
}
var onCmdDisconnect = func(conn connections.ConnectionOutput, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	client.DefaultManager.ExecDisconnect(msg)
	return nil
}
var onCmdSend = func(conn connections.ConnectionOutput, cmd command.Command) error {
	id := current.Load().(string)
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	client.DefaultManager.Send(id, []byte(msg))
	return nil
}
var onCmdTriggers = func(conn connections.ConnectionOutput, cmd command.Command) error {
	id := current.Load().(string)
	client.DefaultManager.ExecTriggers(id)
	return nil
}
var onCmdCreateGame = func(conn connections.ConnectionOutput, cmd command.Command) error {
	forms.CreateGame(cmd.Data())
	return nil
}

func init() {
	handlers.Add("change", onCmdChange)
	handlers.Add("connect", onCmdConnect)
	handlers.Add("disconnect", onCmdDisconnect)
	handlers.Add("triggers", onCmdTriggers)
	handlers.Add("send", onCmdSend)
	handlers.Add("createGame", onCmdCreateGame)
	current.Store("")
}
