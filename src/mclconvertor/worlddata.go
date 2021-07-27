package main

import "github.com/herb-go/herbplugin"

type WorldData struct {
	Host                  string
	Port                  string
	Charset               string
	Name                  string
	CommandStackCharacter string
	ScriptPrefix          string
	QueueDelay            int
	Params                map[string]string
	Permissions           []string
	ScriptID              string
	Trusted               herbplugin.Trusted
	Triggers              []*Trigger
	Timers                []*Timer
	Aliases               []*Alias
}
