package ui

import (
	"encoding/json"
	"modules/services/client"

	"github.com/herb-go/connections/command"
	"github.com/herb-go/connections/room/message"
)

func newRoomAdapter(cmdtype string) func(m *message.Message) error {
	return func(m *message.Message) error {
		var err error
		if m.Room != "" {
			data := command.New()
			data.CommandType = cmdtype
			data.CommandData, err = json.Marshal(m.Data)
			if err != nil {
				return err
			}
			msg, err := data.Encode()
			if err != nil {
				return err
			}
			rooms.Broadcast(m.Room, msg)
		}
		return nil
	}
}

func newUserAdapter(cmdtype string) func(m *message.Message) error {
	return func(m *message.Message) error {
		var err error
		if m.Room == "" {
			data := command.New()
			data.CommandType = cmdtype
			data.CommandData, err = json.Marshal(m.Data)
			if err != nil {
				return err
			}
			msg, err := data.Encode()
			if err != nil {
				return err
			}
			return SendToUser(msg)
		}
		return nil
	}
}

var adapter = message.NewAdapter()

func init() {
	adapter["line"] = newRoomAdapter("line")
	adapter["lines"] = newRoomAdapter("lines")
	adapter["prompt"] = newRoomAdapter("prompt")
	adapter["triggers"] = newRoomAdapter("triggers")
	adapter["clients"] = newUserAdapter("clients")
	adapter["connected"] = newUserAdapter("connected")
	adapter["disconnected"] = newUserAdapter("disconnected")
	adapter["createFail"] = newUserAdapter("createFail")
	adapter["createSuccess"] = newUserAdapter("createSuccess")
	adapter["triggerFail"] = newRoomAdapter("triggerFail")
	adapter["triggerSuccess"] = newRoomAdapter("triggerSuccess")
	adapter["allLines"] = newRoomAdapter("allLines")
	go func() {
		for {
			select {
			case m, ok := <-client.DefaultManager.CommandOutput:
				if ok == false {
					return
				}
				adapter.Exec(m)
			}
		}
	}()
}
