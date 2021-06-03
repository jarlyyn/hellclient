package bus

import (
	"sync"
)

type Bus struct {
	ID                    string
	Locker                sync.RWMutex
	GetHost               func() string
	GetConnBuffer         func() []byte
	GetConnConnected      func() bool
	GetCharset            func() string
	GetVisualCurrentLines func() []*Line
	DoSendToServer        func(cmd []byte) error
	DoSend                func(cmd []byte) error
	OnConnReceive         func(msg []byte)
	OnConnError           func(err error)
	OnConnPrompt          func(msg []byte)
	DoConnectServer       func() error
	DoCloseServer         func() error
	OnVisualError         func(err error)
}

func New() *Bus {
	return &Bus{}
}
