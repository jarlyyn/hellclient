package forms

import (
	"encoding/json"
	"hellclient/modules/userpassword"
	"hellclient/modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//UpdatePasswordFormFieldLabels form field labels map.
var UpdatePasswordFormFieldLabels = map[string]string{}

//UpdatePasswordForm form struct for create game
type UpdatePasswordForm struct {
	formdata.Form
	Username       string
	Password       string
	RepeatPassword string
}

//UpdatePasswordFormID form id of form create timer
const UpdatePasswordFormID = "formUpdatePassword"

//NewUpdatePasswordForm create new create game form
func NewUpdatePasswordForm() *UpdatePasswordForm {
	form := &UpdatePasswordForm{}
	form.SetComponentLabels(ui.MapLabels(UpdatePasswordFormFieldLabels))
	return form
}

func (f *UpdatePasswordForm) ComponentID() string {
	return UpdatePasswordFormID
}

//Validate Validate form and return any error if raised.
func (f *UpdatePasswordForm) Validate() error {
	f.ValidateFieldf(f.Username != "", "Username", "用户名为空")
	f.ValidateFieldf(f.Password != "", "Password", "密码为空")
	if !f.HasError() {
		f.ValidateFieldf(f.Password == f.RepeatPassword, "RepeatPassword", "密码不匹配")

	}
	return nil
}

//InitWithRequest init  create timer form  with http request.
func (f *UpdatePasswordForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func UpdatePassword(t *titan.Titan, data []byte) bool {
	form := NewUpdatePasswordForm()
	err := json.Unmarshal(data, form)
	if err != nil {
		return false
	}
	err = form.Validate()
	if err != nil {
		return false
	}
	errors := form.Errors()
	if len(errors) != 0 {
		t.OnCreateFail(errors)
		return false
	}
	userpassword.Set(form.Username, form.Password)
	return true
}
