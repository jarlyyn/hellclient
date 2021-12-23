package appevents

import "github.com/herb-go/events"

//ExitApp Event type of exit app.
var ExitApp = events.Type("exitapp")

//EmitExitApp emit exit app event
var EmitExitApp = func() bool {
	return events.Emit(ExitApp.NewEvent())
}

//OnExitApp register exit app event handler.
var OnExitApp = events.WrapOn(ExitApp)
