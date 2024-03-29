package world

import (
	"github.com/herb-go/herbplugin"
)

type ClientInfo struct {
	ID         string
	ReadyAt    int64
	Position   int
	HostPort   string
	ScriptID   string
	Running    bool
	Priority   int
	LastActive int64
	Summary    []*Line
}

type ClientInfos []*ClientInfo

// Len is the number of elements in the collection.
func (info ClientInfos) Len() int {
	return len(info)
}

// Less reports whether the element with index i
func (info ClientInfos) Less(i, j int) bool {
	if info[i].Position != info[j].Position {
		return info[i].Position < info[j].Position
	}
	return info[i].ReadyAt < info[j].ReadyAt
}

// Swap swaps the elements with indexes i and j.
func (info ClientInfos) Swap(i, j int) {
	info[i], info[j] = info[j], info[i]
}

const DefaultCommandStackCharacter = ";"
const DefaultScriptPrefix = "/"

type WorldData struct {
	Host                  string
	Port                  string
	Charset               string
	Proxy                 string
	Name                  string
	CommandStackCharacter string
	ScriptPrefix          string
	QueueDelay            int
	Params                map[string]string
	ParamComments         map[string]string
	Permissions           []string
	ScriptID              string
	ShowBroadcast         bool
	ShowSubneg            bool
	ModEnabled            bool
	Trusted               herbplugin.Trusted
	Triggers              []*Trigger
	Timers                []*Timer
	Aliases               []*Alias
	//api 23.11.30
	AutoSave           bool
	IgnoreBatchCommand bool
}

func (d *WorldData) ConvertSettings(id string) *WorldSettings {
	settings := &WorldSettings{
		ID: id,
	}
	if d != nil {
		settings.Host = d.Host
		settings.Port = d.Port
		settings.Proxy = d.Proxy
		settings.Charset = d.Charset
		settings.Name = d.Name
		settings.CommandStackCharacter = d.CommandStackCharacter
		settings.ScriptPrefix = d.ScriptPrefix
		settings.ShowBroadcast = d.ShowBroadcast
		settings.ShowSubneg = d.ShowSubneg
		settings.ModEnabled = d.ModEnabled
		settings.AutoSave = d.AutoSave
		settings.IgnoreBatchCommand = d.IgnoreBatchCommand
	}
	return settings
}

type WorldSettings struct {
	ID                    string
	Host                  string
	Port                  string
	Proxy                 string
	Charset               string
	Name                  string
	CommandStackCharacter string
	ScriptPrefix          string
	ShowBroadcast         bool
	ShowSubneg            bool
	ModEnabled            bool
	//api 23.11.30
	AutoSave           bool
	IgnoreBatchCommand bool
}

func NewWorldData() *WorldData {
	return &WorldData{
		Params:                map[string]string{},
		ParamComments:         map[string]string{},
		CommandStackCharacter: DefaultCommandStackCharacter,
		ScriptPrefix:          DefaultScriptPrefix,
	}
}
