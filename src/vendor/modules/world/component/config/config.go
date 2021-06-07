package config

import (
	"io/ioutil"
	"modules/world"
	"modules/world/bus"
	"sync"

	"github.com/herb-go/util"

	"github.com/pelletier/go-toml"
)

type Data struct {
	Host    string
	Port    string
	Charset string
	Params  map[string]string
}

func NewData() *Data {
	return &Data{
		Params: map[string]string{},
	}
}

type Config struct {
	bus    *bus.Bus
	Locker sync.Mutex
	Data   *Data
}

func (c *Config) GetHostPort() string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Host + ":" + c.Data.Port
}
func (c *Config) GetCharset() string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Charset
}
func (c *Config) Save() error {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	filename := world.GetWorldPath(c.bus.ID)
	data, err := toml.Marshal(c.Data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, util.DefaultFileMode)
}
func (c *Config) Load() error {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	filename := world.GetWorldPath(c.bus.ID)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	configdata := NewData()
	err = toml.Unmarshal(data, configdata)
	if err != nil {
		return err
	}
	c.Data = configdata
	return nil
}
func (c *Config) InstallTo(b *bus.Bus) {
	c.bus = b
	b.GetHostPort = c.GetHostPort
	b.GetCharset = c.GetCharset
}

func (c *Config) UninstallFrom(b *bus.Bus) {
	if c.bus != b {
		return
	}
	b.GetHostPort = nil
}
