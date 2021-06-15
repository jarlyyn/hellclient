package titan

import (
	"modules/app"
	"modules/msg"
	"modules/world/bus"
	"modules/world/component"
	"modules/world/component/config"
	"modules/world/component/conn"
	"modules/world/component/converter"
	"modules/world/component/info"
	"modules/world/component/log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/herb-go/connections/room/message"
	"github.com/herb-go/herb/ui/validator"
	"github.com/herb-go/misc/busevent"
	"github.com/herb-go/util"
)

type Titan struct {
	Locker   sync.RWMutex
	Worlds   map[string]*bus.Bus
	Path     string
	msgEvent *busevent.Event
}

func (t *Titan) CreateBus() *bus.Bus {
	b := bus.New()
	component.InstallComponents(b,
		config.New(),
		conn.New(),
		converter.New(),
		info.New(),
		log.New(),
		t,
	)
	b.RaiseReadyEvent()
	return b
}
func (t *Titan) DestoryBus(b *bus.Bus) {
	b.RaiseCloseEvent()
	b.Reset()
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
	return b
}

func (t *Titan) DoSendTo(id string, msg []byte) error {
	w := t.World(id)
	return w.DoSend(msg)
}
func (t *Titan) Publish(msg *message.Message) {
	go func() {
		t.msgEvent.Raise(msg)
	}()
}

func (t *Titan) onConnected(b *bus.Bus) {
	msg.PublishConnected(t, b.ID)
}
func (t *Titan) onDisonnected(b *bus.Bus) {
	msg.PublishDisconnected(t, b.ID)
}
func (t *Titan) onPrompt(b *bus.Bus, prompt *bus.Line) {
	msg.PublishPrompt(t, b.ID, prompt)
}
func (t *Titan) onLine(b *bus.Bus, line *bus.Line) {
	msg.PublishLine(t, b.ID, line)
}
func (t *Titan) OnCreateFail(errors []*validator.FieldError) {
	msg.PublishCreateFail(t, errors)
}
func (t *Titan) OnCreateSuccess(id string) {
	msg.PublishCreateSuccess(t, id)

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
func (t *Titan) HandleCmdSend(id string, msg []byte) {
	w := t.World(id)
	if w != nil {
		w.HandleCmdError(w.DoSend(msg))
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
		msg.PublishLines(t, id, alllines)
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
	if err != nil {
		util.LogError(err)
		return false
	}
	return ok
}

func (t *Titan) ExecClients() {
	t.Locker.RLock()
	defer t.Locker.RUnlock()
	var result = make([]*bus.ClientInfo, len(t.Worlds))
	var i = 0
	for _, v := range t.Worlds {
		result[i] = v.GetClientInfo()
		i++
	}
	msg.PublishClients(t, result)
}
func (t *Titan) InstallTo(b *bus.Bus) {
	b.BindContectedEvent(t, t.onConnected)
	b.BindDiscontectedEvent(t, t.onConnected)
	b.BindLineEvent(t, t.onLine)
	b.BindPromptEvent(t, t.onPrompt)
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

func (t *Titan) ListNotOpened() ([]*bus.WorldFile, error) {
	t.Locker.RLock()
	defer t.Locker.RUnlock()
	var result = []*bus.WorldFile{}
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
			result = append(result, &bus.WorldFile{
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
	return true, nil
}
func New() *Titan {
	return &Titan{
		Worlds:   map[string]*bus.Bus{},
		msgEvent: busevent.New(),
	}
}
