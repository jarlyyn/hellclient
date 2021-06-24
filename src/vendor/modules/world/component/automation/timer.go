package automation

import (
	"modules/world"
	"sync"
	"time"
)

type Timer struct {
	Locker sync.Locker
	Data   *world.Timer
	Timer  *time.Timer
	OnFire func()
}

func (t *Timer) onFire() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.Timer = nil
	go func() {
		t.OnFire()
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
