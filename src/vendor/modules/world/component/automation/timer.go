package automation

import (
	"modules/world"
	"strconv"
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
func (t *Timer) start() {
	if t.Timer == nil {
		t.Timer = time.AfterFunc(t.Data.GetDuration(), t.onFire)
	}
}
func (t *Timer) Start() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.start()
}
func (t *Timer) stop() {
	if t.Timer != nil {
		t.Timer.Stop()
		t.Timer = nil
	}
}
func (t *Timer) Stop() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.stop()
}

func (t *Timer) Reset() {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	if t.Timer != nil {
		t.Timer.Reset(t.Data.GetDuration())
	}
}
func (t *Timer) Option(name string) (string, bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	switch name {
	case "active_closed":
		if t.Data.ActionWhenDisconnectd {
			return world.StringYes, true
		}
		return "", true
	case "at_time":
		if t.Data.AtTime {
			return world.StringYes, true
		}
		return "", true
	case "enabled":
		if t.Data.Enabled {
			return world.StringYes, true
		}
		return "", true
	case "group":
		return t.Data.Group, true
	case "hour":
		return strconv.Itoa(t.Data.Hour), true
	case "minute":
		return strconv.Itoa(t.Data.Minute), true
	case "name":
		return t.Data.Name, true
	case "offset_hour":
		return "0", true
	case "offset_minute":
		return "0", true
	case "offset_second":
		return "0", true
	case "omit_from_log":
		return "", true
	case "omit_from_output":
		return "", true
	case "one_shot":
		if t.Data.OneShot {
			return world.StringYes, true
		}
		return "", true
	case "script":
		return t.Data.Script, true
	case "second":
		return strconv.Itoa(t.Data.Second), true
	case "send":
		return t.Data.Send, true
	case "send_to":
		return strconv.Itoa(t.Data.SendTo), true
	case "user":
		return "0", true
	case "variable":
		return t.Data.Variable, true
	}

	return "", false
}

func (t *Timer) SetOption(name string, val string) (bool, bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	switch name {
	case "active_closed":
		t.Data.ActionWhenDisconnectd = (val == world.StringYes)
		return true, true
	case "at_time":
		t.Data.AtTime = (val == world.StringYes)
		return true, true
	case "enabled":
		t.Data.Enabled = (val == world.StringYes)
		if t.Data.Enabled {
			t.start()
		} else {
			t.stop()
		}
		return true, true
	case "group":
		t.Data.Group = val
		return true, true
	case "hour":
		i, err := strconv.Atoi(val)
		if err != nil {
			return false, true
		}
		t.Data.Hour = i
		return true, true

	case "minute":
		i, err := strconv.Atoi(val)
		if err != nil {
			return false, true
		}
		t.Data.Minute = i
		return true, true
	case "name":
		t.Data.Name = val
		return true, true
	case "offset_hour":
		return false, false
	case "offset_minute":
		return false, false
	case "offset_second":
		return false, false
	case "omit_from_log":
		return false, false
	case "omit_from_output":
		return false, false
	case "one_shot":
		t.Data.OneShot = (val == world.StringYes)
		return true, true
	case "script":
		t.Data.Script = val
		return true, true
	case "second":
		i, err := strconv.Atoi(val)
		if err != nil {
			return false, true
		}
		t.Data.Second = i
		return true, true
	case "send":
		t.Data.Send = val
		return true, true
	case "send_to":
		i, _ := strconv.Atoi(val)

		t.Data.SendTo = i
		return true, true
	case "user":
		return false, false
	case "variable":
		t.Data.Variable = val
		return true, true
	}

	return false, false
}
