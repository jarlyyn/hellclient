package titan

import (
	"bytes"
	"errors"
	"hellclient/modules/app"
	"hellclient/modules/msg"
	"hellclient/modules/version"
	"hellclient/modules/world"
	"hellclient/modules/world/bus"
	"hellclient/modules/world/component"
	"hellclient/modules/world/component/automation"
	"hellclient/modules/world/component/config"
	"hellclient/modules/world/component/conn"
	"hellclient/modules/world/component/converter"
	"hellclient/modules/world/component/info"
	"hellclient/modules/world/component/log"
	"hellclient/modules/world/component/metronome"
	"hellclient/modules/world/component/queue"
	"hellclient/modules/world/component/script"
	"hellclient/modules/world/hellswitch"

	"path"
	"sort"

	"github.com/BurntSushi/toml"

	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/herb-go/connections/room/message"
	"github.com/herb-go/herb/ui/validator"
	"github.com/herb-go/misc/busevent"
	"github.com/herb-go/util"
)

type Titan struct {
	Locker         sync.RWMutex
	Worlds         map[string]*bus.Bus
	Path           string
	hellswitch     *hellswitch.Hellswitch
	Scriptpath     string
	Skeletonpath   string
	Logpath        string
	MaxHistory     int
	MaxLines       int
	MaxRecent      int
	LinesPerScreen int
	msgEvent       *busevent.Event
}

func (t *Titan) CreateBus() *bus.Bus {
	b := bus.New()
	component.InstallComponents(b,
		config.New(),
		conn.New(),
		converter.New(),
		info.New(),
		automation.New(),
		log.New(),
		queue.New(),
		script.New(),
		metronome.New(),
		t,
	)
	b.RaiseInitEvent()
	return b
}
func (t *Titan) DestoryBus(b *bus.Bus) {
	b.RaiseBeforeCloseEvent()
	b.RaiseCloseEvent()
}
func (t *Titan) find(id string) *bus.Bus {
	return t.Worlds[id]
}

func (t *Titan) World(id string) *bus.Bus {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	return t.find(id)
}

func (t *Titan) NewWorld(id string) *bus.Bus {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] != nil {
		return nil
	}
	b := t.CreateBus()
	b.ID = id
	t.Worlds[id] = b
	b.RaiseReadyEvent()
	return b
}

func (t *Titan) Publish(msg *message.Message) {
	go func() {
		t.msgEvent.Raise(msg)
	}()
}

func (t *Titan) onConnected(b *bus.Bus) {
	b.DoPrintSystem(app.Time.Datetime(time.Now()) + "  成功连接服务器")
	msg.PublishConnected(t, b.ID)
}
func (t *Titan) onDisconnected(b *bus.Bus) {
	b.DoPrintSystem(app.Time.Datetime(time.Now()) + "  与服务器断开连接接 ")
	msg.PublishDisconnected(t, b.ID)
}
func (t *Titan) onPrompt(b *bus.Bus, prompt *world.Line) {
	msg.PublishPrompt(t, b.ID, prompt)
}

func (t *Titan) onStatus(b *bus.Bus, status string) {
	msg.PublishStatus(t, b.ID, status)
}

func (t *Titan) onHistory(b *bus.Bus, h []string) {
	msg.PublishHistory(t, b.ID, h)
}
func (t *Titan) onScriptMessage(b *bus.Bus, data interface{}) {
	msg.PublishScriptMessage(t, b.ID, data)
}

func (t *Titan) onLines(b *bus.Bus, lines []*world.Line) {
	msg.PublishLines(t, b.ID, lines)
}
func (t *Titan) onLine(b *bus.Bus, line *world.Line) {
	if line.OmitFromOutput {
		return
	}
	msg.PublishLine(t, b.ID, line)
}
func (t *Titan) onBroadcast(b *bus.Bus, bc *world.Broadcast) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	for _, v := range t.Worlds {
		go v.DoSendBroadcastToScript(bc)
	}
	if bc.Global {
		go t.hellswitch.Broadcast(bytes.Join([][]byte{[]byte(bc.Channel), []byte(bc.Message)}, GlobalMessageSep))
	}
}
func (t *Titan) OnCreateFail(errors []*validator.FieldError) {
	msg.PublishCreateFail(t, errors)
}
func (t *Titan) OnCreateSuccess(id string) {
	msg.PublishCreateSuccess(t, id)
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoConnectServer())
	}
}
func (t *Titan) OnUpdateSuccess(id string) {
	msg.PublishUpdateSuccess(t, id)
}
func (t *Titan) OnUpdateScriptSuccess(id string) {
	msg.PublishUpdateScriptSuccess(t, id)
}

func (t *Titan) OnCreateScriptFail(errors []*validator.FieldError) {
	msg.PublishCreateScriptFail(t, errors)
}
func (t *Titan) OnCreateScriptSuccess(id string) {
	msg.PublishCreateScriptSuccess(t, id)
}

func (t *Titan) HandleCmdConnect(id string) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoConnectServer())
	}
}
func (t *Titan) HandleCmdDisconnect(id string) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoCloseServer())
	}
}
func (t *Titan) HandleCmdStatus(id string) {
	w := t.World(id)
	if w != nil {
		status := w.GetStatus()
		msg.PublishStatus(t, id, status)
	}
}
func (t *Titan) ExecSwitchStatus() {
	go func() {
		msg.PublishSwitchStatusMessage(t, t.hellswitch.Status())
	}()
}
func (t *Titan) HandleCmdHistory(id string) {
	w := t.World(id)
	if w != nil {
		h := w.GetHistories()
		msg.PublishHistory(t, id, h)
	}
}
func (t *Titan) HandleCmdAllLines(id string) {
	w := t.World(id)
	if w != nil {
		alllines := w.GetCurrentLines()
		msg.PublishAllLines(t, id, alllines)
	}
}
func (t *Titan) HandleCmdLines(id string) {
	w := t.World(id)
	if w != nil {
		alllines := w.GetCurrentLines()
		start := len(alllines) - t.LinesPerScreen
		if start < 0 {
			start = 0
		}
		msg.PublishLines(t, id, alllines[start:])
	}
}
func (t *Titan) HandleCmdPrompt(id string) {
	w := t.World(id)
	if w != nil {
		pormpt := w.GetPrompt()
		msg.PublishPrompt(t, id, pormpt)
	}
}
func (t *Titan) HandleCmdNotOpened() {
	list, err := t.ListNotOpened()
	if err != nil {
		return
	}
	msg.PublishNotOpened(t, list)
}
func (t *Titan) HandleCmdOpen(id string) bool {
	ok, err := t.OpenWorld(id)
	if err != nil && !os.IsNotExist(err) {
		util.LogError(err)
		return false
	}
	return ok
}
func (t *Titan) HandleCmdSave(id string) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(t.SaveWorld(id))
	}
}
func (t *Titan) HandleCmdSaveScript(id string) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoSaveScript())
	}
}
func (t *Titan) HandleCmdScriptInfo(id string) {
	w := t.World(id)
	if w != nil {
		info := w.GetScriptData().ConvertInfo(w.GetScriptID())
		msg.PublishScriptInfo(t, id, info)
	}
}
func (t *Titan) HandleCmdListScriptInfo() {
	info, err := t.ListScripts()
	if err != nil {
		util.LogError(err)
		return
	}
	msg.PublishScriptInfoList(t, info)
}
func (t *Titan) HandleCmdUseScript(id string, script string) {
	w := t.World(id)
	if w != nil {
		w.DoUseScript(script)
	}
}

func (t *Titan) HandleCmdReloadScript(id string) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoReloadScript())
	}
}
func (t *Titan) HandleCmdCallback(id string, cb *world.Callback) {
	w := t.World(id)
	if w != nil {
		go w.DoSendCallbackToScript(cb)
	}
}
func (t *Titan) HandleCmdAssist(id string) {
	w := t.World(id)
	if w != nil {
		go w.DoAssist()
	}
}
func (t *Titan) HandleCmdAbout() {
	go msg.PublishVersionMessage(t, version.Version.FullVersionCode())
}
func (t *Titan) HandleCmdDefaultServer() {
	go msg.PublishDefaultServerMessage(t, app.System.DefaultServer)
}

func (t *Titan) ExecClients() {
	t.Locker.RLock()
	defer t.Locker.RUnlock()
	var result = make(world.ClientInfos, len(t.Worlds))
	var i = 0
	for _, v := range t.Worlds {
		result[i] = v.GetClientInfo()
		i++
	}
	sort.Sort(result)
	msg.PublishClients(t, result)
}
func (t *Titan) GetMaxHistory() int {
	return t.MaxHistory
}
func (t *Titan) GetMaxLines() int {
	return t.MaxLines
}
func (t *Titan) GetMaxRecent() int {
	return t.MaxRecent
}

func (t *Titan) InstallTo(b *bus.Bus) {
	b.BindConnectedEvent(t, t.onConnected)
	b.BindDisconnectedEvent(t, t.onDisconnected)
	b.BindLineEvent(t, t.onLine)
	b.BindPromptEvent(t, t.onPrompt)
	b.BindStatusEvent(t, t.onStatus)
	b.BindHistoriesEvent(t, t.onHistory)
	b.BindLinesEvent(t, t.onLines)
	b.BindBroadcastEvent(t, t.onBroadcast)
	b.BindScriptMessageEvent(t, t.onScriptMessage)
	b.GetScriptPath = t.GetScriptPath
	b.GetLogsPath = t.GetLogsPath
	b.GetSkeletonPath = t.GetSkeletonPath
	b.GetScriptHome = b.WrapGetString(t.GetScriptHome)
	b.RequestPermissions = b.WrapHandleAuthorization(t.RequestPermissions)
	b.RequestTrustDomains = b.WrapHandleAuthorization(t.RequestTrustDomains)
	b.GetMaxHistory = t.GetMaxHistory
	b.GetMaxLines = t.GetMaxLines
	b.GetMaxRecent = t.GetMaxRecent
}
func (t *Titan) RequestPermissions(b *bus.Bus, a *world.Authorization) {
	w := t.World(b.ID)
	if w != nil {
		msg.PublishRequestPermissions(t, b.ID, a)
	}
}
func (t *Titan) RequestTrustDomains(b *bus.Bus, a *world.Authorization) {
	w := t.World(b.ID)
	if w != nil {
		msg.PublishRequestTrustDomains(t, b.ID, a)
	}
}
func (t *Titan) RaiseMsgEvent(msg *message.Message) {
	t.msgEvent.Raise(msg)
}
func (t *Titan) BindMsgEvent(id interface{}, fn func(t *Titan, msg *message.Message)) {
	t.msgEvent.BindAs(
		id,
		func(data interface{}) {
			fn(t, data.(*message.Message))
		},
	)
}

func (t *Titan) GetWorldPath(id string) string {
	return filepath.Join(t.Path, id) + Ext
}

func (t *Titan) IsWorldExist(id string) (bool, error) {
	_, err := os.Stat(t.GetWorldPath(id))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func (t *Titan) IsAliasNameAvaliable(id string, name string, byuser bool) bool {
	w := t.World(id)
	if w != nil {
		name = world.PrefixedName(name, byuser)
		return w.HasNamedAlias(name)
	}
	return false
}

func (t *Titan) DoCreateAlias(id string, alias *world.Alias) bool {
	w := t.World(id)
	if w != nil {
		return w.AddAlias(alias, false)
	}
	return false
}
func (t *Titan) DoUpdateAlias(id string, alias *world.Alias) int {
	w := t.World(id)
	if w != nil {
		return w.DoUpdateAlias(alias)
	}
	return world.UpdateFailNotFound
}
func (t *Titan) OnCreateAliasSuccess(world string, id string) {
	msg.PublishCreateAliasSuccess(t, world, id)
}
func (t *Titan) OnUpdateAliasSuccess(world string, id string) {
	msg.PublishUpdateAliasSuccess(t, world, id)
}
func (t *Titan) HandleCmdAliases(id string, byuser bool) {
	w := t.World(id)
	if w != nil {
		aliases := w.GetAliasesByType(byuser)
		if byuser {
			msg.PublishUserAliases(t, id, aliases)
		} else {
			msg.PublishScriptAliases(t, id, aliases)
		}
	}
}
func (t *Titan) HandleCmdDeleteAlias(world string, id string) {
	w := t.World(world)
	if w != nil {
		w.DoDeleteAlias(id)
	}
}
func (t *Titan) HandleCmdLoadAlias(world string, id string) {
	w := t.World(world)
	if w != nil {
		alias := w.GetAlias(id)
		if alias != nil {
			msg.PublishAlias(t, world, alias)
		}
	}
}

func (t *Titan) IsTriggerNameAvaliable(id string, name string, byuser bool) bool {
	w := t.World(id)
	if w != nil {
		name = world.PrefixedName(name, byuser)
		return w.HasNamedTrigger(name)
	}
	return false
}
func (t *Titan) DoCreateTrigger(id string, trigger *world.Trigger) bool {
	w := t.World(id)
	if w != nil {
		return w.AddTrigger(trigger, false)
	}
	return false
}
func (t *Titan) DoUpdateTrigger(id string, trigger *world.Trigger) int {
	w := t.World(id)
	if w != nil {
		return w.DoUpdateTrigger(trigger)
	}
	return world.UpdateFailNotFound
}
func (t *Titan) OnCreateTriggerSuccess(world string, id string) {
	msg.PublishCreateTriggerSuccess(t, world, id)
}
func (t *Titan) OnUpdateTriggerSuccess(world string, id string) {
	msg.PublishUpdateTriggerSuccess(t, world, id)
}
func (t *Titan) HandleCmdTriggers(id string, byuser bool) {
	w := t.World(id)
	if w != nil {
		triggers := w.GetTriggersByType(byuser)
		if byuser {
			msg.PublishUserTriggers(t, id, triggers)
		} else {
			msg.PublishScriptTriggers(t, id, triggers)
		}
	}
}
func (t *Titan) HandleCmdDeleteTrigger(world string, id string) {
	w := t.World(world)
	if w != nil {
		w.DoDeleteTrigger(id)
	}
}
func (t *Titan) HandleCmdLoadTrigger(world string, id string) {
	w := t.World(world)
	if w != nil {
		trigger := w.GetTrigger(id)
		if trigger != nil {
			msg.PublishTrigger(t, world, trigger)
		}
	}
}

func (t *Titan) IsTimerNameAvaliable(id string, name string, byuser bool) bool {
	w := t.World(id)
	if w != nil {
		name = world.PrefixedName(name, byuser)
		return w.HasNamedTimer(name)
	}
	return false
}

func (t *Titan) DoCreateTimer(id string, timer *world.Timer) bool {
	w := t.World(id)
	if w != nil {
		return w.AddTimer(timer, false)
	}
	return false
}
func (t *Titan) DoUpdateTimer(id string, timer *world.Timer) int {
	w := t.World(id)
	if w != nil {
		return w.DoUpdateTimer(timer)
	}
	return world.UpdateFailNotFound
}
func (t *Titan) OnCreateTimerSuccess(world string, id string) {
	msg.PublishCreateTimerSuccess(t, world, id)
}
func (t *Titan) OnUpdateTimerSuccess(world string, id string) {
	msg.PublishUpdateTimerSuccess(t, world, id)
}

func (t *Titan) HandleCmdTimers(id string, byuser bool) {
	w := t.World(id)
	if w != nil {
		timers := w.GetTimersByType(byuser)
		if byuser {
			msg.PublishUserTimers(t, id, timers)
		} else {
			msg.PublishScriptTimers(t, id, timers)
		}
	}
}
func (t *Titan) HandleCmdDeleteTimer(world string, id string) {
	w := t.World(world)
	if w != nil {
		w.DoDeleteTimer(id)
	}
}
func (t *Titan) HandleCmdLoadTimer(world string, id string) {
	w := t.World(world)
	if w != nil {
		timer := w.GetTimer(id)
		if timer != nil {
			msg.PublishTimer(t, world, timer)
		}
	}
}

func (t *Titan) HandleCmdSend(id string, msg string) {
	w := t.World(id)
	if w != nil && msg != "" {
		w.AddHistory(msg)
		w.DoExecute(msg)
	}
}
func (t *Titan) HandleCmdMasssend(id string, msg string) {
	w := t.World(id)
	if w != nil {
		m := world.CreateCommand(msg)
		m.History = false
		w.DoSend(m)
	}
}
func (t *Titan) HandleCmdFindHistory(id string, position int) {
	if position < 0 {
		return
	}
	w := t.World(id)
	if w != nil {
		h := w.GetHistories()
		if position >= len(h) {
			return
		}
		msg.PublishFoundHistory(t, id, world.CreateFoundHistory(position, h[len(h)-1-position]))
	}
}

func (t *Titan) GetSkeletonPath() string {
	return t.Skeletonpath
}
func (t *Titan) GetScriptPath() string {
	return t.Scriptpath
}
func (t *Titan) GetLogsPath() string {
	return t.Logpath
}
func (t *Titan) IsScriptExist(id string) (bool, error) {
	_, err := os.Stat(path.Join(t.Scriptpath, id))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (t *Titan) ListNotOpened() ([]*world.WorldFile, error) {
	t.Locker.RLock()
	defer t.Locker.RUnlock()
	var result = []*world.WorldFile{}
	files, err := os.ReadDir(t.Path)
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		if !v.IsDir() {
			name := v.Name()
			if !strings.HasSuffix(name, Ext) {
				continue

			}
			id := strings.TrimSuffix(name, Ext)
			if t.Worlds[id] != nil {
				continue
			}
			i, err := v.Info()
			if err != nil {
				return nil, err
			}
			ut := app.Time.Datetime(i.ModTime())
			result = append(result, &world.WorldFile{
				ID:          id,
				LastUpdated: ut,
			})
		}

	}
	return result, nil
}
func (t *Titan) listWorlds() ([]string, error) {
	result := []string{}
	files, err := os.ReadDir(t.Path)
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		if !v.IsDir() {
			name := v.Name()
			if strings.HasSuffix(name, Ext) {
				result = append(result, strings.TrimSuffix(name, Ext))
			}
		}
	}
	return result, nil
}
func (t *Titan) ListScripts() ([]*world.ScriptInfo, error) {
	t.Locker.RLock()
	defer t.Locker.RUnlock()
	result := []*world.ScriptInfo{}
	files, err := os.ReadDir(t.Scriptpath)
	if err != nil {
		return nil, err
	}
	for _, v := range files {
		if v.IsDir() {
			id := v.Name()
			if world.IDRegexp.Match([]byte(id)) {
				data, err := os.ReadFile(filepath.Join(t.Scriptpath, id, "script.toml"))
				if err != nil {
					continue
				}
				d := world.NewScriptData()
				err = toml.Unmarshal(data, d)
				if err != nil {
					continue
				}
				result = append(result, d.ConvertInfo(id))
			}
		}
	}
	return result, nil
}
func (t *Titan) CloseWorld(id string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	w := t.Worlds[id]
	if w == nil {
		return
	}
	delete(t.Worlds, id)
	t.DestoryBus(w)
}
func (t *Titan) SaveWorld(id string) error {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	w := t.Worlds[id]
	if w == nil {
		return nil
	}
	data, err := w.DoEncode()
	if err != nil {
		return err
	}
	return os.WriteFile(t.GetWorldPath(id), data, util.DefaultFileMode)
}
func (t *Titan) OpenWorld(id string) (bool, error) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] != nil {
		return false, nil
	}
	b := t.CreateBus()
	b.ID = id
	data, err := os.ReadFile(t.GetWorldPath(id))
	if err != nil {
		return false, err
	}
	err = b.DoDecode(data)
	if err != nil {
		return false, err
	}
	t.Worlds[id] = b
	b.RaiseReadyEvent()
	go b.HandleCmdError(b.DoConnectServer())

	return true, nil
}
func (t *Titan) HandleCmdParams(id string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	i := &world.ParamsInfo{}
	i.Params = t.Worlds[id].GetParams()
	i.ParamComments = t.Worlds[id].GetParamComments()
	i.RequiredParams = t.Worlds[id].GetRequiredParams()
	msg.PublishParamsinfo(t, id, i)

}
func (t *Titan) HandleCmdDeleteParam(id string, name string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	t.Worlds[id].DeleteParam(name)
	msg.PublishParamDeleted(t, id, name)
	go t.HandleCmdParams(id)

}
func (t *Titan) HandleCmdUpdateParam(id string, name string, value string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	t.Worlds[id].SetParam(name, value)
	msg.PublishParamUpdated(t, id, name)
	go t.HandleCmdParams(id)

}
func (t *Titan) HandleCmdUpdateParamComment(id string, name string, value string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	t.Worlds[id].SetParamComment(name, value)
	msg.PublishParamUpdated(t, id, name)
	go t.HandleCmdParams(id)

}
func (t *Titan) HandleCmdWorldSettings(id string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	s := t.Worlds[id].GetWorldData().ConvertSettings(id)
	msg.PublishWorldSettingsMessage(t, id, s)
}
func (t *Titan) HandleCmdScriptSettings(id string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	w := t.Worlds[id]
	s := w.GetScriptData().ConvertSettings(w.GetScriptID())
	msg.PublishScriptSettingsMessage(t, id, s)
}
func (t *Titan) HandleCmdRequiredParams(id string) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	var p []*world.RequiredParam
	s := t.Worlds[id].GetScriptData()
	if s != nil {
		p = s.RequiredParams
	}
	msg.PublishRequiredParamsMessage(t, id, p)
}

func (t *Titan) HandleCmdRequestPermissions(a *world.Authorization) {
	w := t.World(a.World)
	if w != nil {
		items := w.GetPermissions()
	Next:
		for _, authed := range a.Items {
			for _, owned := range items {
				if owned == authed {
					continue Next
				}
			}
			items = append(items, authed)
		}
		w.SetPermissions(items)
		if a.Script != "" {
			go w.DoRunScript(a.Script)
		}
	}
}

func (t *Titan) HandleCmdRequestTrustDomains(a *world.Authorization) {
	w := t.World(a.World)
	if w != nil {
		trusted := w.GetTrusted()
	Next:
		for _, authed := range a.Items {
			for _, owned := range trusted.Domains {
				if owned == authed {
					continue Next
				}
			}
			trusted.Domains = append(trusted.Domains, authed)
		}
		w.SetTrusted(trusted)
		if a.Script != "" {
			go w.DoRunScript(a.Script)
		}
	}
}
func (t *Titan) HandleCmdAuthorized(id string) {
	w := t.World(id)
	if w != nil {
		a := world.NewAuthorized()
		p := w.GetPermissions()
		trusted := w.GetTrusted()
		a.Permissions = append([]string{}, p...)
		a.Domains = append([]string{}, trusted.Domains...)
		msg.PublishAuthorized(t, id, a)
	}
}
func (t *Titan) HandleCmdRevokeAuthorized(id string) {
	w := t.World(id)
	if w != nil {
		w.SetPermissions([]string{})
		trusted := w.GetTrusted()
		trusted.Domains = []string{}
		w.SetTrusted(trusted)
		msg.PublishAuthorized(t, id, world.NewAuthorized())
	}
}

func (t *Titan) HandleCmdUpdateRequiredParams(id string, p []*world.RequiredParam) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Worlds[id] == nil {
		return
	}
	data := t.Worlds[id].GetScriptData()
	if data != nil {
		data.SetRequiredParams(p)
	}
	msg.PublishRequiredParamsMessage(t, id, data.RequiredParams)
}
func (t *Titan) NewScript(id string, scripttype string) error {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	ok, err := t.IsScriptExist(id)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("script exists")
	}
	err = os.MkdirAll(filepath.Join(t.Scriptpath, id, "script"), util.DefaultFolderMode)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(world.ScriptTomlTemplates[scripttype])
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(t.Scriptpath, id, "script.toml"), data, util.DefaultFileMode)
	if err != nil {
		return err
	}
	data, err = os.ReadFile(world.ScriptTemplates[scripttype])
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(t.Scriptpath, id, "script", world.ScriptTargets[scripttype]), data, util.DefaultFileMode)
	if err != nil {
		return err
	}
	return nil

}
func (t *Titan) GetScriptHome(b *bus.Bus) string {
	sid := b.GetScriptID()
	if sid == "" {
		return ""
	}
	return filepath.Join(t.Path, b.ID, sid)
}

var GlobalMessageSep = []byte(" ")

func (t *Titan) OnGlobalMessage(msg []byte) {
	var data = bytes.SplitN(msg, GlobalMessageSep, 3)
	switch string(data[0]) {
	case "broadcast":
		if len(data) == 3 {
			t.Locker.Lock()
			bc := world.CreateBroadcast(string(data[1]), string(data[2]), true)
			for _, v := range t.Worlds {
				go v.DoSendBroadcastToScript(bc)
			}
			t.Locker.Unlock()
		}
	}

}

func (t *Titan) OnSwitchStatusChange(status int) {

	go func() {
		msg.PublishSwitchStatusMessage(t, status)
	}()
}
func (t *Titan) Start() {
	go t.hellswitch.Start()
}
func (t *Titan) Stop() {
	t.hellswitch.Stop()
	t.hellswitch.Close()
}
func New() *Titan {
	t := &Titan{
		Worlds:   map[string]*bus.Bus{},
		msgEvent: busevent.New(),
	}
	t.hellswitch = hellswitch.New()
	t.hellswitch.OnGlobalMessage = t.OnGlobalMessage
	t.hellswitch.OnSwitchStatusChange = t.OnSwitchStatusChange
	return t
}
