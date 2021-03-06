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
	if ctx == nil {
		return nil
	}
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
func (p *Prophet) onCmdSaveScript(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdSaveScript(msg)
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

func (p *Prophet) onCmdReloadScript(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdReloadScript(msg)
	return nil
}
func (p *Prophet) onCmdTimers(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdTimers(msg[0], msg[1] == "byuser")
	return nil
}
func (p *Prophet) onCmdCreateTimer(conn connections.OutputConnection, cmd command.Command) error {
	forms.CreateTimer(p.Titan, cmd.Data())
	return nil

}
func (p *Prophet) onCmdDeleteTimer(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdDeleteTimer(msg[0], msg[1])
	p.Titan.HandleCmdTimers(msg[0], false)
	p.Titan.HandleCmdTimers(msg[0], true)
	return nil

}
func (p *Prophet) onCmdLoadTimer(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdLoadTimer(msg[0], msg[1])
	return nil
}
func (p *Prophet) onCmdUpdateTimer(conn connections.OutputConnection, cmd command.Command) error {
	forms.UpdateTimer(p.Titan, cmd.Data())
	return nil

}
func (p *Prophet) onCmdAliases(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdAliases(msg[0], msg[1] == "byuser")
	return nil
}
func (p *Prophet) onCmdCreateAlias(conn connections.OutputConnection, cmd command.Command) error {
	forms.CreateAlias(p.Titan, cmd.Data())
	return nil

}
func (p *Prophet) onCmdDeleteAlias(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdDeleteAlias(msg[0], msg[1])
	p.Titan.HandleCmdAliases(msg[0], false)
	p.Titan.HandleCmdAliases(msg[0], true)
	return nil

}
func (p *Prophet) onCmdLoadAlias(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdLoadAlias(msg[0], msg[1])
	return nil
}
func (p *Prophet) onCmdUpdateAlias(conn connections.OutputConnection, cmd command.Command) error {
	forms.UpdateAlias(p.Titan, cmd.Data())
	return nil
}
func (p *Prophet) onCmdTriggers(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdTriggers(msg[0], msg[1] == "byuser")
	return nil
}
func (p *Prophet) onCmdCreateTrigger(conn connections.OutputConnection, cmd command.Command) error {
	forms.CreateTrigger(p.Titan, cmd.Data())
	return nil

}
func (p *Prophet) onCmdDeleteTrigger(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdDeleteTrigger(msg[0], msg[1])
	p.Titan.HandleCmdTriggers(msg[0], false)
	p.Titan.HandleCmdTriggers(msg[0], true)
	return nil

}
func (p *Prophet) onCmdLoadTrigger(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdLoadTrigger(msg[0], msg[1])
	return nil
}
func (p *Prophet) onCmdUpdateTrigger(conn connections.OutputConnection, cmd command.Command) error {
	forms.UpdateTrigger(p.Titan, cmd.Data())
	return nil
}
func (p *Prophet) onCmdParams(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdParams(msg)
	return nil
}
func (p *Prophet) onCmdUpdateParam(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 3 {
		return nil
	}
	p.Titan.HandleCmdUpdateParam(msg[0], msg[1], msg[2])
	return nil
}
func (p *Prophet) onCmdDeleteParam(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	p.Titan.HandleCmdDeleteParam(msg[0], msg[1])
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
	handlers.Register("savescript", p.onCmdSaveScript)
	handlers.Register("reloadScript", p.onCmdReloadScript)
	handlers.Register("status", p.onCmdListStatus)
	handlers.Register("timers", p.onCmdTimers)
	handlers.Register("createTimer", p.onCmdCreateTimer)
	handlers.Register("deleteTimer", p.onCmdDeleteTimer)
	handlers.Register("loadTimer", p.onCmdLoadTimer)
	handlers.Register("updateTimer", p.onCmdUpdateTimer)
	handlers.Register("aliases", p.onCmdAliases)
	handlers.Register("createAlias", p.onCmdCreateAlias)
	handlers.Register("deleteAlias", p.onCmdDeleteAlias)
	handlers.Register("loadAlias", p.onCmdLoadAlias)
	handlers.Register("updateAlias", p.onCmdUpdateAlias)
	handlers.Register("triggers", p.onCmdTriggers)
	handlers.Register("createTrigger", p.onCmdCreateTrigger)
	handlers.Register("deleteTrigger", p.onCmdDeleteTrigger)
	handlers.Register("loadTrigger", p.onCmdLoadTrigger)
	handlers.Register("updateTrigger", p.onCmdUpdateTrigger)

	handlers.Register("params", p.onCmdParams)
	handlers.Register("updateParam", p.onCmdUpdateParam)
	handlers.Register("deleteParam", p.onCmdDeleteParam)

	// handlers.Register("saveTrigger", p.onCmdSaveTrigger)
	// handlers.Register("triggers", p.onCmdTriggers)

}
