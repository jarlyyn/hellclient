package config

import (
	"bytes"
	"hellclient/modules/world"
	"hellclient/modules/world/bus"
	"sort"
	"sync"
	"time"

	"github.com/herb-go/herbplugin"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Locker  sync.Mutex
	Data    *world.WorldData
	ReadyAt int64
}

func (c *Config) GetWorldData() *world.WorldData {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data
}
func (c *Config) GetReadyAt() int64 {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.ReadyAt
}
func (c *Config) GetPort(bus *bus.Bus) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Port
}
func (c *Config) SetPort(bus *bus.Bus, port string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.Port = port
}
func (c *Config) GetQueueDelay(bus *bus.Bus) int {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.QueueDelay
}
func (c *Config) SetQueueDelay(bus *bus.Bus, delay int) {
	c.Locker.Lock()
	c.Data.QueueDelay = delay
	c.Locker.Unlock()
	bus.RaiseQueueDelayUpdatedEvent()
}
func (c *Config) GetPermissions() []string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Permissions
}
func (c *Config) SetPermissions(v []string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.Permissions = v
}
func (c *Config) GetTrusted() *herbplugin.Trusted {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return &c.Data.Trusted
}
func (c *Config) SetTrusted(v *herbplugin.Trusted) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.Trusted = *v
}
func (c *Config) GetHost(bus *bus.Bus) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Host
}
func (c *Config) SetHost(bus *bus.Bus, host string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.Host = host
}
func (c *Config) GetCharset(bus *bus.Bus) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Charset
}
func (c *Config) SetCharset(bus *bus.Bus, charset string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.Charset = charset
}
func (c *Config) GetParam(key string) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Params[key]
}
func (c *Config) SetParam(key string, value string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.Params[key] = value
}
func (c *Config) DeleteParam(key string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	delete(c.Data.Params, key)
}
func (c *Config) GetParams() map[string]string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	result := make(map[string]string, len(c.Data.Params))
	for k := range c.Data.Params {
		result[k] = c.Data.Params[k]
	}
	return result
}

func (c *Config) GetParamComment(key string) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.ParamComments[key]
}
func (c *Config) SetParamComment(key string, value string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.ParamComments[key] = value
}
func (c *Config) GetParamComments() map[string]string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	result := make(map[string]string, len(c.Data.ParamComments))
	for k := range c.Data.ParamComments {
		result[k] = c.Data.ParamComments[k]
	}
	return c.Data.ParamComments
}

func (c *Config) GetScriptID() string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.ScriptID
}
func (c *Config) SetScriptID(id string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.ScriptID = id
}
func (c *Config) Encode(bus *bus.Bus) ([]byte, error) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	timers := bus.GetTimersByType(true)
	sort.Sort(world.Timers(timers))
	c.Data.Timers = timers
	aliases := bus.GetAliasesByType(true)
	sort.Sort(world.Aliases(aliases))
	c.Data.Aliases = aliases
	triggers := bus.GetTriggersByType(true)
	sort.Sort(world.Triggers(triggers))
	c.Data.Triggers = triggers
	buf := bytes.NewBuffer(nil)
	err := toml.NewEncoder(buf).Encode(c.Data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (c *Config) Decode(bus *bus.Bus, data []byte) error {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	configdata := world.NewWorldData()
	err := toml.Unmarshal(data, configdata)
	if err != nil {
		return err
	}
	for k := range configdata.Timers {
		configdata.Timers[k].SetByUser(true)
	}
	for k := range configdata.Aliases {
		configdata.Aliases[k].SetByUser(true)
	}
	for k := range configdata.Triggers {
		configdata.Triggers[k].SetByUser(true)
	}
	c.Data = configdata
	bus.AddTimers(c.Data.Timers)
	bus.AddAliases(c.Data.Aliases)
	bus.AddTriggers(c.Data.Triggers)
	return nil
}
func (c *Config) GetName() string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Name
}
func (c *Config) SetName(n string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.Name = n
}
func (c *Config) GetCommandStackCharacter() string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.CommandStackCharacter
}
func (c *Config) SetCommandStackCharacter(ch string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.CommandStackCharacter = ch
}
func (c *Config) GetScriptPrefix() string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.ScriptPrefix
}
func (c *Config) SetScriptPrefix(p string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.ScriptPrefix = p
}
func (c *Config) SetShowBroadcast(s bool) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.ShowBroadcast = s
}
func (c *Config) GetShowBroadcast() bool {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.ShowBroadcast
}
func (c *Config) SetModEnabled(s bool) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.ModEnabled = s
}
func (c *Config) GetModEnabled() bool {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.ModEnabled
}
func (c *Config) OnReady(b *bus.Bus) {
	c.ReadyAt = time.Now().Unix()
}
func (c *Config) InstallTo(b *bus.Bus) {
	b.GetHost = b.WrapGetString(c.GetHost)
	b.SetHost = b.WrapHandleString(c.SetHost)
	b.GetPort = b.WrapGetString(c.GetPort)
	b.SetPort = b.WrapHandleString(c.SetPort)
	b.GetCharset = b.WrapGetString(c.GetCharset)
	b.SetCharset = b.WrapHandleString(c.SetCharset)
	b.GetQueueDelay = b.WrapGetInt(c.GetQueueDelay)
	b.SetQueueDelay = b.WrapHandleInt(c.SetQueueDelay)
	b.GetWorldData = c.GetWorldData
	b.DoEncode = b.WrapDoEncode(c.Encode)
	b.DoDecode = b.WrapDoDecode(c.Decode)
	b.SetParam = c.SetParam
	b.GetParam = c.GetParam
	b.GetParams = c.GetParams
	b.SetParamComment = c.SetParamComment
	b.GetParamComment = c.GetParamComment
	b.GetParamComments = c.GetParamComments
	b.DeleteParam = c.DeleteParam
	b.GetReadyAt = c.GetReadyAt
	b.GetTrusted = c.GetTrusted
	b.SetTrusted = c.SetTrusted
	b.GetPermissions = c.GetPermissions
	b.SetPermissions = c.SetPermissions
	b.GetScriptID = c.GetScriptID
	b.SetScriptID = c.SetScriptID
	b.GetName = c.GetName
	b.SetName = c.SetName
	b.GetCommandStackCharacter = c.GetCommandStackCharacter
	b.SetCommandStackCharacter = c.SetCommandStackCharacter
	b.GetScriptPrefix = c.GetScriptPrefix
	b.SetScriptPrefix = c.SetScriptPrefix
	b.GetShowBroadcast = c.GetShowBroadcast
	b.SetShowBroadcast = c.SetShowBroadcast
	b.GetModEnabled = c.GetModEnabled
	b.SetModEnabled = c.SetModEnabled
	b.BindInitEvent(b, c.OnReady)
}

func New() *Config {
	return &Config{
		Data: world.NewWorldData(),
	}
}
