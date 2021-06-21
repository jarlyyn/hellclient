package script

import (
	"bytes"
	"modules/world"
	"modules/world/bus"
	"path"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/herb-go/herbplugin"
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

func NewData() *Data {
	return &Data{}
}
func EmptyScriptData(scripttype string) *Data {
	return &Data{
		Type:         scripttype,
		OnOpen:       "onOpen",
		OnClose:      "onClose",
		OnConnect:    "onConnect",
		OnDisconnect: "onDisconnect",
	}
}

type Script struct {
	Locker sync.Mutex
	Data   *Data
}

func (s *Script) EncodeScript() ([]byte, error) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	buf := bytes.NewBuffer(nil)
	err := toml.NewEncoder(buf).Encode(s.Data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (s *Script) DecodeScript(data []byte) error {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	scriptdata := NewData()
	err := toml.Unmarshal(data, scriptdata)
	if err != nil {
		return err
	}
	s.Data = scriptdata
	return nil
}

func (s *Script) PluginOptions(b *bus.Bus) *herbplugin.PlainOptions {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	opt := herbplugin.NewOptions()
	opt.Location.Path = path.Join(b.GetScriptPath(), b.GetScriptID(), "scripts")
	opt.Trusted = b.GetTrusted()
	opt.Permissions = b.GetPermissions()
	return opt
}
func (s *Script) ScriptInfo(b *bus.Bus) *world.ScriptInfo {
	id := b.GetScriptID()
	s.Locker.Lock()
	defer s.Locker.Unlock()
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
	b.DoEncodeScript = s.EncodeScript
	b.DoDecodeScript = s.DecodeScript
}

func New() *Script {
	return &Script{}
}
