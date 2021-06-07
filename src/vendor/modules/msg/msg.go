package msg

import (
	"modules/world/bus"

	"github.com/herb-go/connections/room/message"
	"gopkg.in/go-playground/validator.v8"
)

const MsgTypeConnected = "connected"
const MsgTypeDisconnected = "disconnected"
const MsgTypeCreateFail = "createFail"
const MsgTypeCreateSuccess = "createSuccess"
const MsgTypeLine = "line"
const MsgTypePrompt = "prompt"

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
func PublishLine(p Publisher, id string, line *bus.Line) {
	p.Publish(New(MsgTypeLine, id, line))
}
func PublishPrompt(p Publisher, id string, prompt *bus.Line) {
	p.Publish(New(MsgTypePrompt, id, prompt))
}
