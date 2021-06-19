package config

import (
	"bytes"
	"modules/world/bus"
	"sync"

	"github.com/BurntSushi/toml"
)

type Data struct {
	Host       string
	Port       string
	Charset    string
	QueueDelay int
	Params     map[string]string
}

func NewData() *Data {
	return &Data{
		Params: map[string]string{},
	}
}

type Config struct {
	Locker  sync.Mutex
	Changed bool
	Data    *Data
}

func (c *Config) GetPort(bus *bus.Bus) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Port
}
func (c *Config) SetPort(bus *bus.Bus, port string) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Changed = true
	c.Data.Port = port
}
func (c *Config) GetQueueDelay(bus *bus.Bus) int {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.QueueDelay
}
func (c *Config) SetQueueDelay(bus *bus.Bus, delay int) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	c.Data.QueueDelay = delay
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
	return c.Data.Params
}
func (c *Config) Encode() ([]byte, error) {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	buf := bytes.NewBuffer(nil)
	err := toml.NewEncoder(buf).Encode(c.Data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (c *Config) Decode(data []byte) error {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	configdata := NewData()
	err := toml.Unmarshal(data, configdata)
	if err != nil {
		return err
	}
	c.Data = configdata
	return nil
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
	b.DoEncode = c.Encode
	b.DoDecode = c.Decode
	b.SetParam = c.SetParam
	b.GetParam = c.GetParam
	b.GetParams = c.GetParams
	b.DeleteParam = c.DeleteParam
}

func New() *Config {
	return &Config{
		Data: NewData(),
	}
}
