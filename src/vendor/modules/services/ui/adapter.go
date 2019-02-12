package ui

import (
	"encoding/json"
	"modules/services/client"

	"github.com/jarlyyn/herb-go-experimental/connections/command"
)

type adapter map[string]func(*client.Command) error

func (a adapter) Exec(cmd *client.Command) (bool, error) {
	handler, ok := a[cmd.Type]
	if ok == false {
		return false, nil
	}
	err := handler(cmd)
	if err != nil {
		return true, err
	}
	return true, nil

}

var PromptAdapter = func(cmd *client.Command) error {
	var err error
	if cmd.Room != "" {
		data := command.New()
		data.CommandType = "prompt"
		data.CommandData, err = json.Marshal(cmd.Data)
		if err != nil {
			return err
		}
		msg, err := data.Encode()
		if err != nil {
			return err
		}
		rooms.Broadcast(cmd.Room, msg)
	}
	return nil
}
var LineAdapter = func(cmd *client.Command) error {
	var err error
	if cmd.Room != "" {
		data := command.New()
		data.CommandType = "line"
		data.CommandData, err = json.Marshal(cmd.Data)
		if err != nil {
			return err
		}
		msg, err := data.Encode()
		if err != nil {
			return err
		}
		rooms.Broadcast(cmd.Room, msg)
	}
	return nil
}

var DefaultAdapter = adapter{}

func init() {
	DefaultAdapter["line"] = LineAdapter
	DefaultAdapter["prompt"] = PromptAdapter

	go func() {
		for {
			select {
			case m := <-client.DefaultManager.CommandOutput:
				go func() {
					DefaultAdapter.Exec(m)
				}()
			}
		}
	}()
}
