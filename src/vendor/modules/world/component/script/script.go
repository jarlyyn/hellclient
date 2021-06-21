package script

import (
	"modules/world"
	"modules/world/bus"
	"sync"
)

type Data struct {
	Type         string
	OnOpen       string
	OnClose      string
	OnConnect    string
	OnDisconnect string
	Triggers     []*world.Trigger
	Timers       []*world.Timer
	Alias        []*world.Alias
}

type Script struct {
	Locker sync.Mutex
	Data   *Data
}

func (s *Script) ScriptInfo(b *bus.Bus) *world.ScriptInfo {
	id := b.GetScriptID()
	info := &world.ScriptInfo{
		ID: id,
	}
	if s.Data != nil {
		info.Type = s.Data.Type
		info.OnOpen = s.Data.OnOpen
		info.OnClose = s.Data.OnClose
		info.OnConnect = s.Data.OnConnect
		info.OnDisconnect = s.Data.OnDisconnect
	}
	return info
}
func (s *Script) InstallTo(b *bus.Bus) {

	b.GetScriptInfo = b.WrapGetScriptInfo(s.ScriptInfo)
}

func New() *Script {
	return &Script{}
}
