package world

import (
	"time"
)

type Timer struct {
	ID                    string
	Name                  string
	Enabled               bool
	Hour                  int
	Minute                int
	Second                int
	Send                  string
	ScriptName            string
	AtTime                bool
	SendTo                int
	ActionWhenDisconnectd bool
	Temporary             bool
	OneShot               bool
	SpeedWalk             bool
	Group                 string
}

func (t *Timer) GetDuration() time.Duration {
	if t.AtTime {
		now := time.Now()
		at := time.Date(now.Year(), now.Month(), now.Day(), t.Hour, t.Minute, t.Second, 0, now.Location())
		if at.Before(now) {
			at = at.Add(24 * time.Hour)
		}
		return at.Sub(now)
	}
	return time.Duration(t.Hour)*time.Hour + time.Duration(t.Minute)*time.Minute + time.Duration(t.Second)*time.Second
}
func NewTimer() *Timer {
	return &Timer{}
}
