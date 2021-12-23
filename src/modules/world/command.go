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
	var reuslt = []*Command{}
	if c == nil || c.Mesasge == "" {
		return reuslt
	}
	msgs := strings.Split(c.Mesasge, sep)
	for _, v := range msgs {
		cmd := c.Clone()
		cmd.Mesasge = v
		reuslt = append(reuslt, cmd)
	}
	return reuslt
}
func CreateCommand(message string) *Command {
	return &Command{
		Mesasge: message,
		Echo:    true,
	}
}
