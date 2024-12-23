package world

import (
	"strings"
)

type Command struct {
	Mesasge     string
	Echo        bool
	Queue       bool
	Log         bool
	History     bool
	Locked      bool
	Creator     string
	CreatorType string
}

func (c *Command) Clone() *Command {
	cmd := *c
	return &cmd
}

func (c *Command) Split(sep string) []*Command {
	var result = []*Command{}
	if c == nil || c.Mesasge == "" {
		return result
	}
	msgs := strings.Split(c.Mesasge, sep)
	for _, v := range msgs {
		if v != "" {
			cmd := c.Clone()
			cmd.Mesasge = v
			result = append(result, cmd)
		}
	}
	return result
}
func CreateCommand(message string) *Command {
	return &Command{
		Mesasge: message,
		Echo:    true,
	}
}
