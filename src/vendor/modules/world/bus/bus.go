package bus

import (
	"sync"
)

type Bus struct {
	ID              string
	Locker          sync.RWMutex
	GetHost         func() string
	GetConnBuffer   func() []byte
	DoSendToServer  func(cmd []byte) error
	OnConnReceive   func(msg []byte)
	OnConnError     func(err error)
	OnConnPrompt    func(msg []byte)
	DoConnectServer func() error
	DoCloseServer   func() error
}

func New() *Bus {
	return &Bus{}
}
