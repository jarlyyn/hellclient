package world

import (
	"time"

	"github.com/herb-go/uniqueid"
)

const TimerFlagEnabled = 1
const TimerFlagAtTime = 2
const TimerFlagOneShot = 4
const TimerFlagTimerSpeedWalk = 8
const TimerFlagTimerNote = 16
const TimerFlagActiveWhenClosed = 32
const TimerFlagReplace = 1024
const TimerFlagTemporary = 16384

type Timer struct {
	ID                    string
	Name                  string
	Enabled               bool
	Hour                  int
	Minute                int
	Second                float64
	Send                  string
	Script                string
	AtTime                bool
	SendTo                int
	ActionWhenDisconnectd bool
	Temporary             bool
	OneShot               bool
	Group                 string
	Variable              string
	OmitFromLog           bool
	OmitFromOutput        bool
	byuser                bool
}

func (t *Timer) ByUser() bool {
	return t.byuser
}
func (t *Timer) SetByUser(v bool) {
	t.byuser = v
}
func (t *Timer) PrefixedName() string {
	if t.byuser {
		return PrefixByUser + t.Name
	}
	return PrefixByScript + t.Name
}
func (t *Timer) GetDuration() time.Duration {
	if t.AtTime {
		now := time.Now()
		at := time.Date(now.Year(), now.Month(), now.Day(), t.Hour, t.Minute, int(t.Second), 0, now.Location())
		if at.Before(now) {
			at = at.Add(24 * time.Hour)
		}
		return at.Sub(now)
	}
	d := time.Duration(t.Hour)*time.Hour + time.Duration(t.Minute)*time.Minute + time.Duration(t.Second*float64(time.Second))
	if d <= 0 {
		d = time.Second
	}
	return d
}
func NewTimer() *Timer {
	return &Timer{}
}

func CreateTimer() *Timer {
	return &Timer{
		ID: uniqueid.MustGenerateID(),
	}
}

type Timers []*Timer

// Len is the number of elements in the collection.
func (t Timers) Len() int {
	return len(t)
}

// Less reports whether the element with index i
func (t Timers) Less(i, j int) bool {
	return t[i].ID < t[j].ID
}

// Swap swaps the elements with indexes i and j.
func (t Timers) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
