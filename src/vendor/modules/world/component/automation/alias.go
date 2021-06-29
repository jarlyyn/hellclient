package automation

import (
	"modules/world"
	"sync"
)

type Alias struct {
	Locker  sync.RWMutex
	Data    *world.Alias
	Deleted bool
	Matcher world.Matcher
	ByUser  bool
}
