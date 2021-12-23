package forms

import (
	"encoding/json"
	"hellclient/modules/world"
	"hellclient/modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//CreateScriptFormFieldLabels form field labels map.
var CreateScriptFormFieldLabels = map[string]string{
	"ID":   "名称",
	"Type": "类型",
}

//CreateScriptForm form struct for create game
type CreateScriptForm struct {
	formdata.Form
	ID   string
	Type string
}

//CreateScriptFormID form id of form create game
const CreateScriptFormID = "formcreatescript"

//NewCreateScriptForm create new create game form
func NewCreateScriptForm() *CreateScriptForm {
	form := &CreateScriptForm{}
	form.SetComponentLabels(ui.MapLabels(CreateScriptFormFieldLabels))
	return form
}

func (f *CreateScriptForm) ComponentID() string {
	return CreateScriptFormID
}

//Validate Validate form and return any error if raised.
func (f *CreateScriptForm) Validate() error {
	f.ValidateFieldf(len(f.ID) > 2, "ID", "名称至少需要2个字符")
	f.ValidateFieldf(len(f.ID) < 64, "ID", "名称不能超过64个字符")

	f.ValidateFieldf(world.IDRegexp.MatchString(f.ID), "ID", "名称只能包含数字，字母，- _ @ .()[]+")
	f.ValidateFieldf(world.AvailableScriptTypes[f.Type], "Type", "脚本类型不可用")
	if !f.HasError() {
		ok, err := titan.Pangu.IsScriptExist(f.ID)
		if err != nil {
			return err
		}
		f.ValidateFieldf(ok == false, "ID", "名称已经存在")
	}
	return nil
}

//InitWithRequest init  create game form  with http request.
func (f *CreateScriptForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func CreateScript(t *titan.Titan, data []byte) error {
	form := NewCreateScriptForm()
	err := json.Unmarshal(data, form)
	if err != nil {
		return err
	}
	err = form.Validate()
	if err != nil {
		return err
	}
	errors := form.Errors()
	if len(errors) != 0 {
		t.OnCreateScriptFail(errors)
		return nil
	}
	err = t.NewScript(form.ID, form.Type)
	if err != nil {
		return err
	}
	go func() {
		t.OnCreateScriptSuccess(form.ID)
		t.HandleCmdListScriptInfo()
	}()
	return nil
}
