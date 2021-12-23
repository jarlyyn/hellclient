package automation

import (
	"hellclient/modules/world"
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
	data := *t.Data
	t.Timer = nil
	t.Locker.Unlock()
	t.OnFire(&data)
	if !t.Data.OneShot {
		t.Start()
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
		t.stop()
		t.start()
	}
}
func (t *Timer) Info(infotype int) (string, bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	switch infotype {
	case 1:
		return strconv.Itoa(t.Data.Hour), true
	case 2:
		return strconv.Itoa(t.Data.Minute), true
	case 3:
		return strconv.FormatFloat(t.Data.Second, 'f', 2, 64), true
	case 4:
		return t.Data.Send, true
	case 5:
		return t.Data.Script, true
	case 6:
		return world.ToStringBool(t.Data.Enabled), true
	case 7:
		return world.ToStringBool(t.Data.OneShot), true
	case 8:
		return world.ToStringBool(t.Data.AtTime), true
	case 14:
		return world.ToStringBool(t.Data.Temporary), true
	case 19:
		return t.Data.Group, true
	case 20:
		return strconv.Itoa(t.Data.SendTo), true
	case 21:
		return strconv.Itoa(0), true
	case 22:
		return t.Data.Name, true
	case 23:
		return world.ToStringBool(t.Data.OmitFromOutput), true
	case 24:
		return world.ToStringBool(t.Data.OmitFromLog), true
	}
	return "", false
}
func (t *Timer) Option(name string) (string, bool) {
	t.Locker.Lock()
	defer t.Locker.Unlock()
	switch name {
	case "active_closed":
		return world.ToStringBool(t.Data.ActionWhenDisconnectd), true
	case "at_time":
		return world.ToStringBool(t.Data.AtTime), true
	case "enabled":
		return world.ToStringBool(t.Data.Enabled), true
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
		return world.ToStringBool(t.Data.OmitFromLog), true
	case "omit_from_output":
		return world.ToStringBool(t.Data.OmitFromOutput), true
	case "one_shot":
		return world.ToStringBool(t.Data.OneShot), true
	case "script":
		return t.Data.Script, true
	case "second":
		return strconv.FormatFloat(t.Data.Second, 'f', 2, 64), true
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
		t.Data.ActionWhenDisconnectd = world.FromStringBool(val)
		return true, true
	case "at_time":
		t.Data.AtTime = world.FromStringBool(val)
		return true, true
	case "enabled":
		t.Data.Enabled = world.FromStringBool(val)
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
		t.Data.Minute = world.FromStringInt(val)
		return true, true
	case "offset_hour":
		return false, false
	case "offset_minute":
		return false, false
	case "offset_second":
		return false, false
	case "omit_from_log":
		t.Data.OmitFromLog = world.FromStringBool(val)
		return true, true
	case "omit_from_output":
		t.Data.OmitFromOutput = world.FromStringBool(val)
		return true, true
	case "one_shot":
		t.Data.OneShot = world.FromStringBool(val)
		return true, true
	case "script":
		t.Data.Script = val
		return true, true
	case "second":
		t.Data.Second = world.FromStringFloat(val)
		return true, true
	case "send":
		t.Data.Send = val
		return true, true
	case "send_to":
		t.Data.SendTo = world.FromStringInt(val)
		return true, true
	case "user":
		return false, false
	case "variable":
		t.Data.Variable = val
		return true, true
	}

	return false, false
}
