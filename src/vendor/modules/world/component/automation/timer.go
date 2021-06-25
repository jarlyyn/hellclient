package automation

import (
	"modules/world"
	"sync"
	"time"
)

type Timer struct {
	Locker sync.RWMutex
	Data   *world.Timer
	Timer  *time.Timer
	OnFire func(*world.Timer)
}

func (t *Timer) onFire() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.Timer = nil
	go func() {
		t.OnFire(t.Data)
	}()
	if !t.Data.OneShot {
		go func() {
			t.Start()
		}()
	}
}
func (t *Timer) Start() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.Timer = time.AfterFunc(t.Data.GetDuration(), t.onFire)
}
func (t *Timer) Stop() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Timer != nil {
		t.Timer.Stop()
		t.Timer = nil
	}
}
