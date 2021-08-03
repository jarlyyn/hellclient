package script

import (
	"bytes"
	"modules/mapper"
	"modules/world"
	"modules/world/bus"
	"modules/world/component/script/jsengine"
	"modules/world/component/script/luaengine"
	"os"
	"path"
	"path/filepath"
	"sort"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/util"
)

type Script struct {
	CreatorLocker sync.Mutex
	creator       string
	creatorType   string
	Locker        sync.Mutex
	EngineLocker  sync.Mutex
	Status        string
	Data          *world.ScriptData
	Mapper        *mapper.Mapper
	engine        Engine
}

func (s *Script) GetMapper() *mapper.Mapper {
	return s.Mapper
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
	timers := b.GetTimersByType(false)
	sort.Sort(world.Timers(timers))
	s.Data.Timers = timers

	aliases := b.GetAliasesByType(false)
	sort.Sort(world.Aliases(aliases))
	s.Data.Aliases = aliases

	triggers := b.GetTriggersByType(false)
	sort.Sort(world.Triggers(triggers))
	s.Data.Triggers = triggers

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
	for k := range s.Data.Timers {
		s.Data.Timers[k].SetByUser(false)
	}
	for k := range s.Data.Aliases {
		s.Data.Aliases[k].SetByUser(false)
	}
	for k := range s.Data.Triggers {
		s.Data.Triggers[k].SetByUser(false)
	}
	b.AddTimers(s.Data.Timers)
	b.AddAliases(s.Data.Aliases)
	b.AddTriggers(s.Data.Triggers)
	return nil
}
func (s *Script) Unload(b *bus.Bus) {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	s.unload(b)

}
func (s *Script) SetCreator(creatortype string, creator string) {
	s.CreatorLocker.Lock()
	s.creator = creator
	s.creatorType = creatortype
	s.CreatorLocker.Unlock()
}
func (s *Script) CreatorAndType() (string, string) {
	s.CreatorLocker.Lock()
	defer s.CreatorLocker.Unlock()
	return s.creator, s.creatorType
}
func (s *Script) unload(b *bus.Bus) {
	if s.engine != nil {
		s.engine.Close(b)
	}
	s.Mapper.Reset()
	s.SetCreator("", "")
	b.DoDeleteTimerByType(false)
	b.DoDeleteAliasByType(false)
	b.DoDeleteTriggerByType(false)
	s.Data = nil
	s.engine = nil
}

func (s *Script) Load(b *bus.Bus) error {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	return s.load(b)
}
func (s *Script) load(b *bus.Bus) error {
	s.Mapper.Reset()
	err := s.open(b)
	if err != nil {
		return err
	}
	data := b.GetScriptData()
	if data != nil {
		switch data.Type {
		case "lua":
			s.engine = luaengine.NewLuaEngine()
		case "jscript":
			s.engine = jsengine.NewJsEngine()
		}
	}
	if s.engine != nil {
		return s.engine.Open(b)
	}
	return nil
}
func (s *Script) Reload(b *bus.Bus) error {
	b.DoMultiLinesFlush()
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	s.unload(b)
	return s.load(b)
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
	b.DoMultiLinesFlush()
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	if s.engine != nil {
		s.SetCreator("system", "connected")
		s.engine.OnConnect(b)
	}
}
func (s *Script) disconnected(b *bus.Bus) {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	if s.engine != nil {
		s.SetCreator("system", "disconnected")
		s.engine.OnDisconnect(b)
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
func (s *Script) SendTimer(b *bus.Bus, timer *world.Timer) {
	e := s.getEngine()
	if e != nil {
		go func() {
			if timer.Script != "" {
				s.SetCreator("timer", timer.Script)
			} else {
				s.SetCreator("timer", "#"+timer.ID)
			}
			e.OnTimer(b, timer)
		}()
	}
}
func (s *Script) getEngine() Engine {
	s.EngineLocker.Lock()
	defer s.EngineLocker.Unlock()
	return s.engine
}
func (s *Script) SendAlias(b *bus.Bus, message string, alias *world.Alias, result *world.MatchResult) {
	e := s.getEngine()
	if e != nil {
		go func() {
			if alias.Script != "" {
				s.SetCreator("alias", alias.Script)
			} else {
				s.SetCreator("alias", "#"+alias.ID)
			}
			e.OnAlias(b, message, alias, result)
		}()
	}
}
func (s *Script) SendTrigger(b *bus.Bus, line *world.Line, trigger *world.Trigger, result *world.MatchResult) {
	e := s.getEngine()
	if e != nil {
		if trigger.Script != "" {
			s.SetCreator("trigger", trigger.Script)
		} else {
			s.SetCreator("trigger", "#"+trigger.ID)
		}
		e.OnTrigger(b, line, trigger, result)
	}
}
func (s *Script) verifyChannel(channel string) bool {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	if s.Data == nil || channel == "" || s.Data.OnBroadcast == "" {
		return false
	}
	return channel == s.Data.Channel
}
func (s *Script) SendBroadcast(b *bus.Bus, bc *world.Broadcast) {
	if !s.verifyChannel(bc.Channel) {
		return
	}
	e := s.getEngine()
	if e != nil {
		s.SetCreator("broadcast", bc.Message)
		e.OnBroadCast(b, bc)
	}
}
func (s *Script) Run(b *bus.Bus, cmd string) {
	e := s.getEngine()
	if e != nil {
		s.SetCreator("run", "")
		go func() {
			e.Run(b, cmd)
		}()
	}
}
func (s *Script) GetRequiredParams() []*world.RequiredParam {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	if s.Data == nil {
		return nil
	}
	return append([]*world.RequiredParam{}, s.Data.RequiredParams...)
}

func (s *Script) GetScriptType() string {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	if s.Data == nil {
		return ""
	}
	return s.Data.Type
}
func (s *Script) InstallTo(b *bus.Bus) {
	b.GetScriptData = b.WrapGetScriptData(s.ScriptData)
	b.DoReloadScript = b.WrapDo(s.Reload)
	b.DoSaveScript = b.WrapDo(s.SaveScript)
	b.DoUseScript = b.WrapHandleString(s.UseScript)
	b.GetScriptPluginOptions = b.WrapGetScriptPluginOptions(s.PluginOptions)
	b.DoSendTimerToScript = b.WrapHandleTimer(s.SendTimer)
	b.DoSendAliasToScript = b.WrapHandleAlias(s.SendAlias)
	b.DoSendTriggerToScript = b.WrapHandleTrigger(s.SendTrigger)
	b.DoSendBroadcastToScript = b.WrapHandleBroadcast(s.SendBroadcast)
	b.DoRunScript = b.WrapHandleString(s.Run)
	b.GetStatus = s.GetStatus
	b.SetStatus = b.WrapHandleString(s.SetStatus)

	b.GetMapper = s.GetMapper

	b.GetRequiredParams = s.GetRequiredParams
	b.GetScriptType = s.GetScriptType
	b.GetScriptCaller = s.CreatorAndType
	b.BindReadyEvent(s, s.ready)
	b.BindBeforeCloseEvent(s, s.beforeClose)
	b.BindConnectedEvent(s, s.connected)
	b.BindServerCloseEvent(s, s.disconnected)

}

func New() *Script {
	return &Script{
		Mapper: mapper.New(),
	}
}
