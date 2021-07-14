package world

import "github.com/herb-go/uniqueid"

type Broadcast struct {
	ID      string
	Channel string
	Message string
	Global  bool
}

func CreateBroadcast(channel string, msg string, global bool) *Broadcast {
	return &Broadcast{
		ID:      uniqueid.MustGenerateID(),
		Channel: channel,
		Message: msg,
		Global:  global,
	}
}
