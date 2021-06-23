package prophet

import (
	"encoding/json"
	"modules/world/titan/forms"

	"github.com/herb-go/connections"
	"github.com/herb-go/connections/command"
	"github.com/herb-go/connections/room"
)

func Send(conn connections.OutputConnection, msgtype string, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return conn.Send([]byte(msgtype + " " + string(bs)))
}
func (p *Prophet) change(conn connections.OutputConnection, id string) error {
	ctx := p.Context(conn.ID())
	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()
	v, ok := ctx.Data.Load("rooms")
	if ok == false {
		v, _ = ctx.Data.LoadOrStore("rooms", room.NewLocation(conn, p.Rooms))
	}
	j := v.(*room.Location)
	j.Leave(p.Current.Load().(string))
	p.Change(id)
	j.Join(id)
	return Send(conn, "current", id)
}
func (p *Prophet) onCmdChange(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	return p.change(conn, msg)
}

func (p *Prophet) onCmdConnect(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdConnect(msg)
	return nil
}

func (p *Prophet) onCmdDisconnect(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdDisconnect(msg)
	return nil
}

func (p *Prophet) onCmdSend(conn connections.OutputConnection, cmd command.Command) error {
	id := p.Current.Load().(string)
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdSend(id, msg)
	return nil
}

func (p *Prophet) onCmdAllLines(conn connections.OutputConnection, cmd command.Command) error {
	id := p.Current.Load().(string)
	p.Titan.HandleCmdAllLines(id)
	return nil
}

func (p *Prophet) onCmdCreateGame(conn connections.OutputConnection, cmd command.Command) error {
	forms.CreateGame(p.Titan, cmd.Data())
	return nil
}
func (p *Prophet) onCmdNotOpened(conn connections.OutputConnection, cmd command.Command) error {
	p.Titan.HandleCmdNotOpened()
	return nil
}
func (p *Prophet) onCmdOpen(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	ok := p.Titan.HandleCmdOpen(msg)
	if ok {
		p.change(conn, msg)
		p.Titan.ExecClients()
	}
	return nil
}
func (p *Prophet) onCmdClose(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.CloseWorld(msg)
	p.change(conn, "")
	p.Titan.ExecClients()
	return nil
}
func (p *Prophet) onCmdSave(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdSave(msg)
	return nil
}
func (p *Prophet) onCmdScriptInfo(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdScriptInfo(msg)
	return nil
}
func (p *Prophet) onCmdCreateScript(conn connections.OutputConnection, cmd command.Command) error {
	forms.CreateScript(p.Titan, cmd.Data())
	return nil
}
func (p *Prophet) onCmdListScriptinfo(conn connections.OutputConnection, cmd command.Command) error {
	p.Titan.HandleCmdListScriptInfo()
	return nil
}
func (p *Prophet) onCmdListStatus(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}

	p.Titan.HandleCmdStatus(msg)
	return nil
}
func (p *Prophet) onCmdUseScript(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdUseScript(msg[0], msg[1])
	p.Titan.HandleCmdScriptInfo(msg[0])
	return nil
}

// func (p *Prophet) onCmdSaveTrigger(conn connections.OutputConnection, cmd command.Command) error {
// 	forms.SaveTrigger(CurrentGameID(), cmd.Data())
// 	return nil
// }
// func (p *Prophet) onCmdTriggers(conn connections.OutputConnection, cmd command.Command) error {
// 	id := p.Current.Load().(string)
// 	client.DefaultManager.ExecTriggers(id)
// 	return nil
// }

func initHandlers(p *Prophet, handlers *command.Handlers) {
	handlers.Register("change", p.onCmdChange)
	handlers.Register("connect", p.onCmdConnect)
	handlers.Register("disconnect", p.onCmdDisconnect)
	handlers.Register("send", p.onCmdSend)
	handlers.Register("allLines", p.onCmdAllLines)
	handlers.Register("createGame", p.onCmdCreateGame)
	handlers.Register("notopened", p.onCmdNotOpened)
	handlers.Register("open", p.onCmdOpen)
	handlers.Register("close", p.onCmdClose)
	handlers.Register("save", p.onCmdSave)
	handlers.Register("scriptinfo", p.onCmdScriptInfo)
	handlers.Register("createScript", p.onCmdCreateScript)
	handlers.Register("listScriptinfo", p.onCmdListScriptinfo)
	handlers.Register("usescript", p.onCmdUseScript)
	handlers.Register("status", p.onCmdListStatus)

	// handlers.Register("saveTrigger", p.onCmdSaveTrigger)
	// handlers.Register("triggers", p.onCmdTriggers)

}
