package appevents

import (
	"github.com/herb-go/herb/events"
	"github.com/herb-go/herb/model"
)

//FormValidated Event type of form validated.
var FormValidated = events.Type("formvalidated")

//EmitFormValidated emit form validated event
var EmitFormValidated = func(form model.Validator) bool {
	return events.Emit(
		FormValidated.NewEvent().
			WithTarget(form.ModelID()).
			WithData(form),
	)
}

//OnFormValidated register form validated event handler.
var OnFormValidated = events.WrapOn(FormValidated)
