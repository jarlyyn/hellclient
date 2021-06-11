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
	Locker sync.Mutex
	Data   *Data
}

func (c *Config) GetPort(bus *bus.Bus) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Port
}
func (c *Config) GetHost(bus *bus.Bus) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Host
}
func (c *Config) GetCharset(bus *bus.Bus) string {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	return c.Data.Charset
}
func (c *Config) Save(bus *bus.Bus) error {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	filename := world.GetWorldPath(bus.ID)
	data, err := toml.Marshal(c.Data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, util.DefaultFileMode)
}
func (c *Config) Load(bus *bus.Bus) error {
	c.Locker.Lock()
	defer c.Locker.Unlock()
	filename := world.GetWorldPath(bus.ID)
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
	b.GetHost = b.WrapGetString(c.GetHost)
	b.GetPort = b.WrapGetString(c.GetPort)
	b.GetCharset = b.WrapGetString(c.GetCharset)
}

func (c *Config) UninstallFrom(b *bus.Bus) {

}
