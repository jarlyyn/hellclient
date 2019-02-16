package  forms

import (
	"github.com/herb-go/herb/model/formdata"
    "net/http"
    "modules/messages"
)


//SaveTriggerFormFieldLabels form field labels map.
var SaveTriggerFormFieldLabels = map[string]string{
	"Field1":  "Field 1",
	"Field2":    "Field 2",
}

//SaveTriggerForm form struct for save trigger
type SaveTriggerForm struct {
	formdata.Form
	Field1  *string
	Field2    *string
}

//SaveTriggerFormID form id of form save trigger
const SaveTriggerFormID = "formservices.savetrigger"

//NewSaveTriggerForm create new save trigger form
func NewSaveTriggerForm() *SaveTriggerForm{
	form:=&SaveTriggerForm{}
	form.SetModelID(SaveTriggerFormID)
	form.SetFieldLabels(SaveTriggerFormFieldLabels)
	return form
}
//Validate Validate form and return any error if raised.
func (f *SaveTriggerForm) Validate() error {
    f.ValidateFieldf(f.Field1 != nil, "Field1", messages.MsgFormFieldRequired) 
    f.ValidateFieldf(f.Field2 != nil, "Field2", messages.MsgFormFieldRequired) 
	if !f.HasError() {
	}
	return nil
}

//InitWithRequest init  save trigger form  with http request.
func (f *SaveTriggerForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}