package userpassword

import "sync/atomic"

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
	current.Store((&UserPassword{
		Username: username,
		Password: password,
	}))
}
