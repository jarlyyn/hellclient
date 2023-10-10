package persistdata

import (
	"encoding/json"

	"github.com/herb-go/herb/persist"
	"github.com/herb-go/util"

	"modules/app"
)

// ModuleName module name
const ModuleName = "800persistdata"

// Store persist store
var Store persist.Store

// LoadBytes load data from Store with given key.
func LoadBytes(key string) ([]byte, error) {
	return Store.LoadBytes(key)
}

// LoadString load string from Store with given key.
func LoadString(key string) (string, error) {
	bs, err := Store.LoadBytes(key)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// Load load data from Store with given key.
func Load(key string, v interface{}) error {
	bs, err := Store.LoadBytes(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}

// SaveBytes save bytes to Store with given key.
func SaveBytes(key string, data []byte) error {
	return Store.SaveBytes(key, data)
}

// SaveString save string to Store with given key.
func SaveString(key string, data string) error {
	return Store.SaveBytes(key, []byte(data))
}

// Save save data to Store with given key.
func Save(key string, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return Store.SaveBytes(key, bs)
}

func init() {
	util.RegisterModule(ModuleName, func() {
		var err error
		Store, err = app.Persistdata.CreateStore()
		util.Must(err)
		util.Must(Store.Start())
		util.OnQuitAndLogError(Store.Stop)
	})
}
