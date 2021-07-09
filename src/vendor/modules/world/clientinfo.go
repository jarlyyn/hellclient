package world

import (
	"github.com/herb-go/herbplugin"
)

type ClientInfo struct {
	ID       string
	ReadyAt  int64
	HostPort string
	ScriptID string
	Running  bool
}

type ClientInfos []*ClientInfo

// Len is the number of elements in the collection.
func (info ClientInfos) Len() int {
	return len(info)
}

// Less reports whether the element with index i
func (info ClientInfos) Less(i, j int) bool {
	return info[i].ReadyAt > info[j].ReadyAt
}

// Swap swaps the elements with indexes i and j.
func (info ClientInfos) Swap(i, j int) {
	info[i], info[j] = info[j], info[i]
}

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

func NewWorldData() *WorldData {
	return &WorldData{
		Params: map[string]string{},
	}
}
