package userpassword

import (
	"hellclient/modules/persistdata"
	"sync/atomic"

	"github.com/herb-go/util"
)

type UserPassword struct {
	Username string
	Password string
}

var current atomic.Value

func Load() (string, string) {
	u, ok := current.Load().(*UserPassword)
	if !ok || u == nil {
		return "", ""
	}
	return u.Username, u.Password
}

func Set(username string, password string) {
	u := &UserPassword{
		Username: username,
		Password: password,
	}
	current.Store(u)
	go func() {
		defer util.Recover()
		persistdata.Save("userpassword", u)
	}()
}
