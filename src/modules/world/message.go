package world

import (
	"fmt"

	"github.com/herb-go/uniqueid"
)

const MessageCommandResponse = 0
const MessageCommandBatchCommand = 1

type Message struct {
	World   string
	ID      string
	Type    string
	Data    string
	Command int
}

func (m *Message) Desc() string {
	return fmt.Sprintf("%s %s %s", m.Type, m.ID, m.Data)
}
func NewMessage() *Message {
	return &Message{}
}

func CreateMessage(world string, msgtype string, data string) *Message {
	return &Message{
		World: world,
		ID:    uniqueid.MustGenerateID(),
		Type:  msgtype,
		Data:  data,
	}
}
