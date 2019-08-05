package appevents

import (
	"github.com/herb-go/herb/events"
	"github.com/herb-go/herb/ui/validator"
)

//FormValidated Event type of form validated.
var FormValidated = events.Type("formvalidated")

//EmitFormValidated emit form validated event
var EmitFormValidated = func(form validator.Fields) bool {
	return events.Emit(
		FormValidated.NewEvent().
			WithTarget(form.ComponentID()).
			WithData(form),
	)
}

//OnFormValidated register form validated event handler.
var OnFormValidated = events.WrapOn(FormValidated)
