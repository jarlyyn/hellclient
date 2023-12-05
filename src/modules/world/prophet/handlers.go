package prophet

import (
	"encoding/json"
	"modules/world"
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
	p.Titan.ExecClients()
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
	itemtype := p.Titan.GetTimerType(msg[0], msg[1])
	if itemtype != nil {
		p.Titan.HandleCmdDeleteTimer(msg[0], msg[1])
		p.Titan.HandleCmdTimers(msg[0], *itemtype)
		if *itemtype {
			go p.Titan.AutoSaveWorld(msg[0])
		}
	}
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
	itemtype := p.Titan.GetAliasType(msg[0], msg[1])
	if itemtype != nil {
		p.Titan.HandleCmdDeleteAlias(msg[0], msg[1])
		p.Titan.HandleCmdAliases(msg[0], *itemtype)
		if *itemtype {
			go p.Titan.AutoSaveWorld(msg[0])
		}
	}
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
	itemtype := p.Titan.GetTriggerType(msg[0], msg[1])
	if itemtype != nil {
		p.Titan.HandleCmdDeleteTrigger(msg[0], msg[1])
		p.Titan.HandleCmdTriggers(msg[0], *itemtype)
		if *itemtype {
			go p.Titan.AutoSaveWorld(msg[0])
		}
	}
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
func (p *Prophet) onCmdUpdateParamComment(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 3 {
		return nil
	}
	p.Titan.HandleCmdUpdateParamComment(msg[0], msg[1], msg[2])
	return nil
}

func (p *Prophet) onCmdUpdateWorldSettings(conn connections.OutputConnection, cmd command.Command) error {
	forms.UpdateGame(p.Titan, cmd.Data())
	return nil
}
func (p *Prophet) onCmdUpdateScriptSettings(conn connections.OutputConnection, cmd command.Command) error {
	forms.UpdateScript(p.Titan, cmd.Data())
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
func (p *Prophet) onCmdCallback(conn connections.OutputConnection, cmd command.Command) error {
	var msg = []string{}
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	if len(msg) < 2 {
		return nil
	}
	cb := &world.Callback{}
	if json.Unmarshal([]byte(msg[1]), &cb) != nil {
		return nil
	}
	p.Titan.HandleCmdCallback(msg[0], cb)
	return nil
}
func (p *Prophet) onCmdAssist(conn connections.OutputConnection, cmd command.Command) error {
	var msg = ""
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdAssist(msg)
	return nil
}
func (p *Prophet) onCmdAbout(conn connections.OutputConnection, cmd command.Command) error {
	p.Titan.HandleCmdAbout()
	return nil
}

func (p *Prophet) onCmdWorldSettings(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdWorldSettings(msg)
	return nil
}

func (p *Prophet) onCmdScriptSettings(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdScriptSettings(msg)
	return nil
}
func (p *Prophet) onCmdRequiredParams(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdRequiredParams(msg)
	return nil
}
func (p *Prophet) onCmdUpdateRequiredParams(conn connections.OutputConnection, cmd command.Command) error {
	var msg = forms.RequiredParamsForm{}
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdUpdateRequiredParams(msg.Current, msg.RequiredParams)
	return nil
}
func (p *Prophet) onCmdDefaultServer(conn connections.OutputConnection, cmd command.Command) error {
	p.Titan.HandleCmdDefaultServer()
	return nil
}
func (p *Prophet) onCmdDefaultCharset(conn connections.OutputConnection, cmd command.Command) error {
	p.Titan.HandleCmdDefaultCharset()
	return nil
}
func (p *Prophet) onCmdRequestPermissions(conn connections.OutputConnection, cmd command.Command) error {
	var msg = &world.Authorization{}
	if json.Unmarshal(cmd.Data(), msg) != nil {
		return nil
	}
	p.Titan.HandleCmdRequestPermissions(msg)
	return nil
}
func (p *Prophet) onCmdRequestTrustDomains(conn connections.OutputConnection, cmd command.Command) error {
	var msg = &world.Authorization{}
	if json.Unmarshal(cmd.Data(), msg) != nil {
		return nil
	}
	p.Titan.HandleCmdRequestTrustDomains(msg)
	return nil
}
func (p *Prophet) onCmdAuthorized(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdAuthorized(msg)
	return nil
}
func (p *Prophet) onCmdRevokeAuthorized(conn connections.OutputConnection, cmd command.Command) error {
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdRevokeAuthorized(msg)
	return nil
}
func (p *Prophet) onCmdMasssend(conn connections.OutputConnection, cmd command.Command) error {
	id := p.Current.Load().(string)
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdMasssend(id, msg)
	return nil
}
func (p *Prophet) onCmdFindHistory(conn connections.OutputConnection, cmd command.Command) error {
	id := p.Current.Load().(string)
	var msg int
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdFindHistory(id, msg)
	return nil
}
func (p *Prophet) onCmdHUDClick(conn connections.OutputConnection, cmd command.Command) error {
	id := p.Current.Load().(string)
	var msg = &world.Click{}
	if json.Unmarshal(cmd.Data(), msg) != nil {
		return nil
	}
	p.Titan.HandleCmdHUDClick(id, msg)
	return nil
}

func (p *Prophet) onCmdUpdatePassword(conn connections.OutputConnection, cmd command.Command) error {
	if forms.UpdatePassword(p.Titan, cmd.Data()) {
		go conn.Close()
	}
	return nil
}

func (p *Prophet) onCmdSortClients(conn connections.OutputConnection, cmd command.Command) error {
	var msg []string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.DoSortClients(msg)
	p.Titan.ExecClients()
	return nil
}
func (p *Prophet) onCmdKeyUp(conn connections.OutputConnection, cmd command.Command) error {
	id := p.Current.Load().(string)
	var msg string
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleCmdKeyUp(id, msg)
	return nil
}
func (p *Prophet) onCmdBatchCommand(conn connections.OutputConnection, cmd command.Command) error {
	var msg = world.NewBatchCommand()
	if json.Unmarshal(cmd.Data(), &msg) != nil {
		return nil
	}
	p.Titan.HandleBatchCommand(msg)
	return nil
}
func (p *Prophet) onCmdBatchCommandScripts(conn connections.OutputConnection, cmd command.Command) error {
	p.Titan.HandleCmdBatchCommandScripts()
	return nil
}

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
	handlers.Register("updatepassword", p.onCmdUpdatePassword)

	handlers.Register("findhistory", p.onCmdFindHistory)

	handlers.Register("params", p.onCmdParams)
	handlers.Register("updateParam", p.onCmdUpdateParam)
	handlers.Register("deleteParam", p.onCmdDeleteParam)

	handlers.Register("updateParamComment", p.onCmdUpdateParamComment)
	handlers.Register("callback", p.onCmdCallback)
	handlers.Register("assist", p.onCmdAssist)
	handlers.Register("about", p.onCmdAbout)

	handlers.Register("worldSettings", p.onCmdWorldSettings)
	handlers.Register("scriptSettings", p.onCmdScriptSettings)
	handlers.Register("requiredParams", p.onCmdRequiredParams)
	handlers.Register("updateRequiredParams", p.onCmdUpdateRequiredParams)
	handlers.Register("updateWorldSettings", p.onCmdUpdateWorldSettings)
	handlers.Register("updateScriptSettings", p.onCmdUpdateScriptSettings)
	handlers.Register("defaultServer", p.onCmdDefaultServer)
	handlers.Register("defaultCharset", p.onCmdDefaultCharset)
	handlers.Register("requestPermissions", p.onCmdRequestPermissions)
	handlers.Register("requestTrustDomains", p.onCmdRequestTrustDomains)
	handlers.Register("authorized", p.onCmdAuthorized)
	handlers.Register("revokeAuthorized", p.onCmdRevokeAuthorized)
	handlers.Register("masssend", p.onCmdMasssend)
	handlers.Register("findhistory", p.onCmdFindHistory)
	handlers.Register("hudclick", p.onCmdHUDClick)
	handlers.Register("sortclients", p.onCmdSortClients)
	handlers.Register("keyup", p.onCmdKeyUp)
	handlers.Register("batchcommand", p.onCmdBatchCommand)
	handlers.Register("batchcommandscripts", p.onCmdBatchCommandScripts)

}
