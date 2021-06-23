package script

import (
	"bytes"
	"modules/world"
	"modules/world/bus"
	"modules/world/component/script/luaengine"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/util"
)

type Script struct {
	Locker       sync.Mutex
	EngineLocker sync.Mutex
	Status       string
	Data         *world.ScriptData
	Engine       Engine
}

func (s *Script) encodeScript() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := toml.NewEncoder(buf).Encode(s.Data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (s *Script) decodeScript(data []byte) error {
	scriptdata := world.NewScriptData()
	err := toml.Unmarshal(data, scriptdata)
	if err != nil {
		return err
	}
	s.Data = scriptdata
	return nil
}

func (s *Script) PluginOptions(b *bus.Bus) herbplugin.Options {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	opt := herbplugin.NewOptions()
	opt.Location.Path = path.Join(b.GetScriptPath(), b.GetScriptID(), "script")
	opt.Trusted = b.GetTrusted()
	opt.Permissions = b.GetPermissions()
	return opt
}
func (s *Script) ScriptData(b *bus.Bus) *world.ScriptData {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	return s.Data
}
func (s *Script) SaveScript(b *bus.Bus) error {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	return s.save(b)
}
func (s *Script) save(b *bus.Bus) error {
	id := b.GetScriptID()
	if id == "" {
		return nil
	}
	data, err := s.encodeScript()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(b.GetScriptPath(), id, "script.toml"), data, util.DefaultFileMode)

}
func (s *Script) OpenScript(b *bus.Bus) error {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	return s.open(b)
}
func (s *Script) open(b *bus.Bus) error {
	id := b.GetScriptID()
	if id == "" {
		return nil
	}
	data, err := os.ReadFile(filepath.Join(b.GetScriptPath(), id, "script.toml"))
	if err != nil {
		return err
	}
	err = s.decodeScript(data)
	if err != nil {
		return err
	}
	return nil
}
func (s *Script) Unload(b *bus.Bus) {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	s.unload(b)

}
func (s *Script) unload(b *bus.Bus) {
	if s.Engine != nil {
		s.Engine.Close(b)
	}
	s.Data = nil
	s.Engine = nil
}

func (s *Script) Load(b *bus.Bus) error {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	return s.load(b)
}
func (s *Script) load(b *bus.Bus) error {
	err := s.open(b)
	if err != nil {
		return err
	}
	data := b.GetScriptData()
	if data != nil {
		switch data.Type {
		case "lua":
			s.Engine = luaengine.NewLuaEngeine()
		}
	}
	if s.Engine != nil {
		return s.Engine.Open(b)
	}
	return nil
}
func (s *Script) ready(b *bus.Bus) {
	s.Load(b)
	go func() {
		b.RaiseStatusEvent(b.GetStatus())
	}()
}
func (s *Script) beforeClose(b *bus.Bus) {
	s.Unload(b)
}
func (s *Script) connected(b *bus.Bus) {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	if s.Engine != nil {
		s.Engine.OnConnect(b)
	}
}
func (s *Script) disconnected(b *bus.Bus) {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	if s.Engine != nil {
		s.Engine.OnDisconnect(b)
	}
}
func (s *Script) UseScript(b *bus.Bus, id string) {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	s.unload(b)
	b.SetScriptID(id)
	err := s.load(b)
	if err != nil {
		b.HandleScriptError(err)
	}
}
func (s *Script) GetStatus() string {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	return s.Status
}
func (s *Script) SetStatus(b *bus.Bus, val string) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	s.Status = val
	go func() {
		b.RaiseStatusEvent(val)
	}()
}
func (s *Script) InstallTo(b *bus.Bus) {
	b.GetScriptData = b.WrapGetScriptData(s.ScriptData)
	b.DoLoadScript = b.WrapDo(s.Load)
	b.DoSaveScript = b.WrapDo(s.SaveScript)
	b.DoUseScript = b.WrapHandleString(s.UseScript)
	b.GetScriptPluginOptions = b.WrapGetScriptPluginOptions(s.PluginOptions)

	b.GetStatus = s.GetStatus
	b.SetStatus = b.WrapHandleString(s.SetStatus)

	b.BindReadyEvent(s, s.ready)
	b.BindBeforeCloseEvent(s, s.beforeClose)
	b.BindConnectedEvent(s, s.connected)
	b.BindDisconnectedEvent(s, s.disconnected)
}

func New() *Script {
	return &Script{}
}
