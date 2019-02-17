package forms

import (
	"net/http"

	"github.com/herb-go/herb/model/formdata"
)

//SaveTriggerFormFieldLabels form field labels map.
var SaveTriggerFormFieldLabels = map[string]string{
	"Pattern":  "匹配项",
	"IsRegExp": "正则",
	"Enabled":  "有效",
	"Commnd":   "命令",
	"Priority": "优先度",
	"Finally":  "不再继续匹配",
}

//SaveTriggerForm form struct for save trigger
type SaveTriggerForm struct {
	formdata.Form
	Pattern  string
	IsRegExp bool
	Enabled  bool
	Commnd   string
	Priority int
	Finally  bool
}

//SaveTriggerFormID form id of form save trigger
const SaveTriggerFormID = "formservices.savetrigger"

//NewSaveTriggerForm create new save trigger form
func NewSaveTriggerForm() *SaveTriggerForm {
	form := &SaveTriggerForm{}
	form.SetModelID(SaveTriggerFormID)
	form.SetFieldLabels(SaveTriggerFormFieldLabels)
	return form
}

//Validate Validate form and return any error if raised.
func (f *SaveTriggerForm) Validate() error {
	f.ValidateFieldf(f.Pattern != "", "Pattern", "匹配项不能为空")

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
