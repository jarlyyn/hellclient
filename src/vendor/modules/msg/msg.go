package msg

import (
	"modules/world"

	"github.com/herb-go/connections/room/message"
	"github.com/herb-go/herb/ui/validator"
)

const MsgTypeConnected = "connected"
const MsgTypeDisconnected = "disconnected"
const MsgTypeCreateFail = "createFail"
const MsgTypeCreateSuccess = "createSuccess"
const MsgTypeCreateScriptFail = "createScriptFail"
const MsgTypeCreateScriptSuccess = "createScriptSuccess"

const MsgTypeLine = "line"
const MsgTypePrompt = "prompt"
const MsgTypeAllLines = "allLines"
const MsgTypeLines = "lines"
const MsgTypeClients = "clients"
const MsgTypeNotOpened = "notopened"
const MsgTypeScriptInfo = "scriptinfo"
const MsgTypeScriptInfoList = "scriptinfoList"
const MsgTypeStatus = "status"
const MsgTypeHistory = "history"
const MsgTypeUserTimers = "usertimers"
const MsgTypeScriptTimers = "scripttimers"
const MsgTypeCreateTimerSuccess = "createTimerSuccess"
const MsgTypeTimer = "timer"
const MsgTypeUpdateTimerSuccess = "updateTimerSuccess"

type Publisher interface {
	Publish(msg *message.Message)
}

func New(msgtype string, room string, data interface{}) *message.Message {
	msg := message.New()
	msg.Type = msgtype
	msg.Room = room
	msg.Data = data
	return msg
}

func PublishConnected(p Publisher, id string) {
	p.Publish(New(MsgTypeConnected, "", id))
}
func PublishDisconnected(p Publisher, id string) {
	p.Publish(New(MsgTypeDisconnected, "", id))
}

func PublishCreateFail(p Publisher, errors []*validator.FieldError) {
	p.Publish(New(MsgTypeCreateFail, "", errors))
}

func PublishCreateSuccess(p Publisher, id string) {
	p.Publish(New(MsgTypeCreateSuccess, "", id))
}

func PublishCreateScriptFail(p Publisher, errors []*validator.FieldError) {
	p.Publish(New(MsgTypeCreateScriptFail, "", errors))
}

func PublishCreateScriptSuccess(p Publisher, id string) {
	p.Publish(New(MsgTypeCreateScriptSuccess, "", id))
}
func PublishLine(p Publisher, id string, line *world.Line) {
	p.Publish(New(MsgTypeLine, id, line))
}
func PublishPrompt(p Publisher, id string, prompt *world.Line) {
	p.Publish(New(MsgTypePrompt, id, prompt))
}

func PublishAllLines(p Publisher, id string, lines []*world.Line) {
	p.Publish(New(MsgTypeAllLines, id, lines))
}
func PublishLines(p Publisher, id string, lines []*world.Line) {
	p.Publish(New(MsgTypeLines, id, lines))
}

func PublishClients(p Publisher, infos []*world.ClientInfo) {
	p.Publish(New(MsgTypeClients, "", infos))
}
func PublishNotOpened(p Publisher, list []*world.WorldFile) {
	p.Publish(New(MsgTypeNotOpened, "", list))
}
func PublishScriptInfo(p Publisher, id string, info *world.ScriptInfo) {
	p.Publish(New(MsgTypeScriptInfo, id, info))
}
func PublishScriptInfoList(p Publisher, info []*world.ScriptInfo) {
	p.Publish(New(MsgTypeScriptInfoList, "", info))
}

func PublishStatus(p Publisher, id string, status string) {
	p.Publish(New(MsgTypeStatus, id, status))
}

func PublishHistory(p Publisher, id string, history []string) {
	p.Publish(New(MsgTypeHistory, id, history))
}
func PublishUserTimers(p Publisher, id string, timers []*world.Timer) {
	p.Publish(New(MsgTypeUserTimers, id, timers))
}
func PublishScriptTimers(p Publisher, id string, timers []*world.Timer) {
	p.Publish(New(MsgTypeScriptTimers, id, timers))
}
func PublishCreateTimerSuccess(p Publisher, world string, id string) {
	p.Publish(New(MsgTypeCreateTimerSuccess, world, id))
}
func PublishTimer(p Publisher, world string, timer *world.Timer) {
	p.Publish(New(MsgTypeTimer, world, timer))
}
func PublishUpdateTimerSuccess(p Publisher, world string, id string) {
	p.Publish(New(MsgTypeUpdateTimerSuccess, world, id))
}
