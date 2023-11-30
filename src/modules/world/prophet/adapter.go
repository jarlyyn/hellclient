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
	adapter.Register("updateSuccess", p.newUserAdapter("updateSuccess"))
	adapter.Register("triggerFail", p.newRoomAdapter("triggerFail"))
	adapter.Register("triggerSuccess", p.newRoomAdapter("triggerSuccess"))
	adapter.Register("allLines", p.newRoomAdapter("allLines"))
	adapter.Register("notopened", p.newUserAdapter("notopened"))
	adapter.Register("scriptinfo", p.newRoomAdapter("scriptinfo"))
	adapter.Register("createScriptFail", p.newUserAdapter("createScriptFail"))
	adapter.Register("createScriptSuccess", p.newUserAdapter("createScriptSuccess"))
	adapter.Register("updateScriptSuccess", p.newUserAdapter("updateScriptSuccess"))
	adapter.Register("scriptinfoList", p.newUserAdapter("scriptinfoList"))
	adapter.Register("status", p.newRoomAdapter("status"))
	adapter.Register("history", p.newRoomAdapter("history"))
	adapter.Register("usertimers", p.newRoomAdapter("usertimers"))
	adapter.Register("scripttimers", p.newRoomAdapter("scripttimers"))
	adapter.Register("createTimerSuccess", p.newRoomAdapter("createTimerSuccess"))
	adapter.Register("timer", p.newRoomAdapter("timer"))
	adapter.Register("updateTimerSuccess", p.newRoomAdapter("updateTimerSuccess"))
	adapter.Register("useraliases", p.newRoomAdapter("useraliases"))
	adapter.Register("scriptaliases", p.newRoomAdapter("scriptaliases"))
	adapter.Register("createAliasSuccess", p.newRoomAdapter("createAliasSuccess"))
	adapter.Register("alias", p.newRoomAdapter("alias"))
	adapter.Register("updateAliasSuccess", p.newRoomAdapter("updateAliasSuccess"))
	adapter.Register("usertriggers", p.newRoomAdapter("usertriggers"))
	adapter.Register("scripttriggers", p.newRoomAdapter("scripttriggers"))
	adapter.Register("createTriggerSuccess", p.newRoomAdapter("createTriggerSuccess"))
	adapter.Register("trigger", p.newRoomAdapter("trigger"))
	adapter.Register("updateTriggerSuccess", p.newRoomAdapter("updateTriggerSuccess"))
	adapter.Register("paramsinfo", p.newRoomAdapter("paramsinfo"))
	adapter.Register("paramupdated", p.newRoomAdapter("paramupdated"))
	adapter.Register("paramdeleted", p.newRoomAdapter("paramdeleted"))
	adapter.Register("paramcommentupdated", p.newRoomAdapter("paramcommentupdated"))
	adapter.Register("scriptMessage", p.newRoomAdapter("scriptMessage"))
	adapter.Register("switchStatus", p.newUserAdapter("switchStatus"))
	adapter.Register("version", p.newUserAdapter("version"))
	adapter.Register("apiversion", p.newUserAdapter("apiversion"))
	adapter.Register("worldSettings", p.newRoomAdapter("worldSettings"))
	adapter.Register("scriptSettings", p.newRoomAdapter("scriptSettings"))
	adapter.Register("requiredParams", p.newRoomAdapter("requiredParams"))
	adapter.Register("defaultServer", p.newUserAdapter("defaultServer"))
	adapter.Register("defaultCharset", p.newUserAdapter("defaultCharset"))
	adapter.Register("requestPermissions", p.newRoomAdapter("requestPermissions"))
	adapter.Register("requestTrustDomains", p.newRoomAdapter("requestTrustDomains"))
	adapter.Register("authorized", p.newRoomAdapter("authorized"))
	adapter.Register("foundhistory", p.newRoomAdapter("foundhistory"))
	adapter.Register("hudcontent", p.newRoomAdapter("hudcontent"))
	adapter.Register("hudupdate", p.newRoomAdapter("hudupdate"))
	adapter.Register("clientinfo", p.newConsoleAdapter("clientinfo"))

}
