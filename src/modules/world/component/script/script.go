package script

import (
	"bytes"
	"fmt"
	"modules/mapper"
	"modules/world"
	"modules/world/bus"
	"modules/world/component/script/jsengine"
	"modules/world/component/script/luaengine"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
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
	Options       *ScriptOptions
}
type ScriptOptions struct {
	Home    string
	ModPath string
	*herbplugin.PlainOptions
}

func (o *ScriptOptions) MustAuthorizePath(path string) bool {
	path = filepath.Clean(path)
	if strings.HasPrefix(path, o.Home) || strings.HasPrefix(path, o.ModPath) {
		return true
	}
	return o.PlainOptions.MustAuthorizePath(path)
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
func (s *Script) ReloadPermissions(b *bus.Bus) {
	s.Locker.Lock()
	defer s.Locker.Unlock()
	s.reloadPermissions(b)
}
func (s *Script) reloadPermissions(b *bus.Bus) {
	if s.Options == nil {
		return
	}
	opt := herbplugin.NewOptions()
	opt.Location.Path = path.Join(b.GetScriptPath(), b.GetScriptID(), "script")
	opt.Trusted = b.GetTrusted()
	opt.Permissions = b.GetPermissions()
	s.Options.PlainOptions = opt
}
func (s *Script) PluginOptions(b *bus.Bus) herbplugin.Options {
	var modpath = ""
	home := b.GetScriptHome()
	modpath = filepath.Join(modpath, b.GetScriptID())
	s.Locker.Lock()
	defer s.Locker.Unlock()
	s.Options = &ScriptOptions{
		Home:    home,
		ModPath: modpath,
	}
	s.reloadPermissions(b)
	return s.Options
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
	b.SetSummary([]*world.Line{})
	b.SetStatus("")
	b.SetHUDSize(0)
	s.SetCreator("", "")
	b.DoDeleteTimerByType(false)
	b.DoDeleteAliasByType(false)
	b.DoDeleteTriggerByType(false)
	b.DoDeleteTemporaryTimers()
	b.DoDeleteTemporaryAliases()
	b.DoDeleteTemporaryTriggers()
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
	b.SetHUDSize(0)
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
func (s *Script) Assist(b *bus.Bus) {
	s.Locker.Lock()
	if s.Data == nil {
		s.Locker.Unlock()
		return
	}
	onAssist := s.Data.OnAssist
	s.Locker.Unlock()
	if onAssist == "" {
		return
	}
	e := s.getEngine()
	if e != nil {
		s.SetCreator("onAssist", "")
		e.OnAssist(b, onAssist)
	}
}
func (s *Script) KeyUp(b *bus.Bus, key string) {
	e := s.getEngine()
	if e != nil {
		go func() {
			s.SetCreator("KeyUp", key)
			e.OnKeyUp(b, key)
		}()
	}

}

func (s *Script) SendCallback(b *bus.Bus, cb *world.Callback) {
	e := s.getEngine()
	if e != nil {
		if cb.Script == "" {
			return
		}
		s.SetCreator("callback", cb.Script)
		e.OnCallback(b, cb)
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
func (s *Script) SendResponse(b *bus.Bus, msg *world.Message) {
	e := s.getEngine()
	if e != nil {
		s.SetCreator("response", msg.Type)
		e.OnResponse(b, msg)
		if b.GetShowBroadcast() {
			b.DoPrintResponse(msg.Desc())
		}
	}
}
func (s *Script) SendHUDClick(b *bus.Bus, c *world.Click) {
	e := s.getEngine()
	if e != nil {
		s.SetCreator("hudclick", "")
		e.OnHUDClick(b, c)
	}
}

func (s *Script) SendBroadcast(b *bus.Bus, bc *world.Broadcast) {
	if !s.verifyChannel(bc.Channel) {
		return
	}
	e := s.getEngine()
	if e != nil {
		s.SetCreator("broadcast", bc.Message)
		e.OnBroadCast(b, bc)
		if b.GetShowBroadcast() {
			if bc.Global {
				b.DoPrintGlobalBroadcastIn(bc.Message)
			} else {
				b.DoPrintLocalBroadcastIn(bc.Message)
			}
		}
	}
}
func (s *Script) HandleFocus(b *bus.Bus) {
	e := s.getEngine()
	if e == nil {
		return
	}
	s.SetCreator("focus", "")
	e.OnFocus(b)
}
func (s *Script) HandleLoseFocus(b *bus.Bus) {
	e := s.getEngine()
	if e == nil {
		return
	}
	s.SetCreator("losefocus", "")
	e.OnLoseFocus(b)
}
func (s *Script) HandleBuffer(b *bus.Bus, data []byte) bool {
	e := s.getEngine()
	if e == nil {
		return false
	}
	s.SetCreator("buffer", "")
	return e.OnBuffer(b, data)
}

func (s *Script) HandleSubneg(b *bus.Bus, data []byte) bool {
	if len(data) < 2 {
		return false
	}
	e := s.getEngine()
	if e == nil {
		return false
	}
	if b.GetShowSubneg() && len(data) > 1 {
		b.DoPrintSubneg(fmt.Sprintf("%d %s", data[0], string(data[1:])))
	}
	s.SetCreator("buffer", "")
	return e.OnSubneg(b, data[0], data[1:])
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
	b.GetPluginOptions = b.WrapGetPluginOptions(s.PluginOptions)
	b.DoSendTimerToScript = b.WrapHandleTimer(s.SendTimer)
	b.DoSendAliasToScript = b.WrapHandleAlias(s.SendAlias)
	b.DoSendTriggerToScript = b.WrapHandleTrigger(s.SendTrigger)
	b.DoSendBroadcastToScript = b.WrapHandleBroadcast(s.SendBroadcast)
	b.DoSendCallbackToScript = b.WrapHandleCallback(s.SendCallback)
	b.DoSendKeyUpToScript = b.WrapHandleString(s.KeyUp)
	b.DoSendResponseToScript = b.WrapHandleResponse(s.SendResponse)
	b.DoSendHUDClickToScript = b.WrapHandleClick(s.SendHUDClick)
	b.DoReloadPermissions = b.Wrap(s.ReloadPermissions)
	b.DoRunScript = b.WrapHandleString(s.Run)
	b.GetStatus = s.GetStatus
	b.SetStatus = b.WrapHandleString(s.SetStatus)

	b.GetMapper = s.GetMapper

	b.GetRequiredParams = s.GetRequiredParams
	b.GetScriptType = s.GetScriptType
	b.GetScriptCaller = s.CreatorAndType
	b.DoAssist = b.Wrap(s.Assist)
	b.HandleBuffer = b.WrapHandleBytesForBool(s.HandleBuffer)
	b.HandleSubneg = b.WrapHandleBytesForBool(s.HandleSubneg)
	b.HandleFocus = b.Wrap(s.HandleFocus)
	b.HandleLoseFocus = b.Wrap(s.HandleLoseFocus)
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
