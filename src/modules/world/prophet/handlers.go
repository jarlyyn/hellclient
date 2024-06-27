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

// 消息处理函数
// 所有客户端指令在这里进行实际函数的分配

func initHandlers(p *Prophet, handlers *command.Handlers) {
	//切换游戏指令
	handlers.Register("change", p.onCmdChange)
	//连线指令
	handlers.Register("connect", p.onCmdConnect)
	//断线指令
	handlers.Register("disconnect", p.onCmdDisconnect)
	//发送用户输入指令
	handlers.Register("send", p.onCmdSend)
	//历史信息指令
	handlers.Register("allLines", p.onCmdAllLines)
	//创建游戏指令
	handlers.Register("createGame", p.onCmdCreateGame)
	//未打开游戏清单指令，由于显示打开界面
	handlers.Register("notopened", p.onCmdNotOpened)
	//打开游戏指令
	handlers.Register("open", p.onCmdOpen)
	//关闭游戏指令
	handlers.Register("close", p.onCmdClose)
	//保存游戏指令
	handlers.Register("save", p.onCmdSave)
	//请求脚本信息指令
	handlers.Register("scriptinfo", p.onCmdScriptInfo)
	//创建脚本指令
	handlers.Register("createScript", p.onCmdCreateScript)
	//列出全部脚本指令
	handlers.Register("listScriptinfo", p.onCmdListScriptinfo)
	//使用脚本指令
	handlers.Register("usescript", p.onCmdUseScript)
	//保存脚本指令
	handlers.Register("savescript", p.onCmdSaveScript)
	//重新加载脚本指令
	handlers.Register("reloadScript", p.onCmdReloadScript)
	//获取状态行内容指令
	handlers.Register("status", p.onCmdListStatus)
	//获取计时器清单指令
	handlers.Register("timers", p.onCmdTimers)
	//创建计时器指令
	handlers.Register("createTimer", p.onCmdCreateTimer)
	//删除计时器指令
	handlers.Register("deleteTimer", p.onCmdDeleteTimer)
	//获取单个计时器指令
	handlers.Register("loadTimer", p.onCmdLoadTimer)
	//更新计时器指令
	handlers.Register("updateTimer", p.onCmdUpdateTimer)
	//别名列表指令
	handlers.Register("aliases", p.onCmdAliases)
	//创建别名指令
	handlers.Register("createAlias", p.onCmdCreateAlias)
	//删除别名指令
	handlers.Register("deleteAlias", p.onCmdDeleteAlias)
	//获取单个别名指令
	handlers.Register("loadAlias", p.onCmdLoadAlias)
	//更新别名指令
	handlers.Register("updateAlias", p.onCmdUpdateAlias)
	//触发器列表指令
	handlers.Register("triggers", p.onCmdTriggers)
	//创建触发器指令
	handlers.Register("createTrigger", p.onCmdCreateTrigger)
	//删除触发器指令
	handlers.Register("deleteTrigger", p.onCmdDeleteTrigger)
	//加载单个触发器指令
	handlers.Register("loadTrigger", p.onCmdLoadTrigger)
	//更新触发器指令
	handlers.Register("updateTrigger", p.onCmdUpdateTrigger)
	//更新密码指令
	handlers.Register("updatepassword", p.onCmdUpdatePassword)
	//查找历史指令，已废弃
	handlers.Register("findhistory", p.onCmdFindHistory)
	//列出变量指令
	handlers.Register("params", p.onCmdParams)
	//更新变量指令
	handlers.Register("updateParam", p.onCmdUpdateParam)
	//删除变量指令
	handlers.Register("deleteParam", p.onCmdDeleteParam)
	//更新变量备注指令
	handlers.Register("updateParamComment", p.onCmdUpdateParamComment)
	//指定回调指令
	handlers.Register("callback", p.onCmdCallback)
	//调用助理按钮对应功能指令
	handlers.Register("assist", p.onCmdAssist)
	//显示服务器介绍信息指令
	handlers.Register("about", p.onCmdAbout)
	//请求游戏信息指令
	handlers.Register("worldSettings", p.onCmdWorldSettings)
	//请求脚本信息指令
	handlers.Register("scriptSettings", p.onCmdScriptSettings)
	//请求脚本参数列表指令
	handlers.Register("requiredParams", p.onCmdRequiredParams)
	//更新脚本参数指令
	handlers.Register("updateRequiredParams", p.onCmdUpdateRequiredParams)
	//更新游戏设置指令
	handlers.Register("updateWorldSettings", p.onCmdUpdateWorldSettings)
	//更新脚本信息指令
	handlers.Register("updateScriptSettings", p.onCmdUpdateScriptSettings)
	//请求默认服务器指令
	handlers.Register("defaultServer", p.onCmdDefaultServer)
	//请求默认编码指令
	handlers.Register("defaultCharset", p.onCmdDefaultCharset)
	//授权指令
	handlers.Register("requestPermissions", p.onCmdRequestPermissions)
	//授权信任域名指令
	handlers.Register("requestTrustDomains", p.onCmdRequestTrustDomains)
	//请求已授权信息指令
	handlers.Register("authorized", p.onCmdAuthorized)
	//注销已授权内容指令
	handlers.Register("revokeAuthorized", p.onCmdRevokeAuthorized)
	//批量发送文本指令
	handlers.Register("masssend", p.onCmdMasssend)
	//HUD被点击指令
	handlers.Register("hudclick", p.onCmdHUDClick)
	//对客户端进行排序指令
	handlers.Register("sortclients", p.onCmdSortClients)
	//用户按键指令
	handlers.Register("keyup", p.onCmdKeyUp)
	//批量发送指令
	handlers.Register("batchcommand", p.onCmdBatchCommand)
	//获取批量发送信息指令
	handlers.Register("batchcommandscripts", p.onCmdBatchCommandScripts)
}
