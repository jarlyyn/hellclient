package forms

import (
	"encoding/json"
	"hellclient/modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//UpdateGameFormFieldLabels form field labels map.
var UpdateGameFormFieldLabels = map[string]string{
	"ID":      "名称",
	"Host":    "网址",
	"Port":    "端口",
	"Charset": "字符编码",
}

//UpdateGameForm form struct for update game
type UpdateGameForm struct {
	formdata.Form
	Name                  string
	ID                    string
	Host                  string
	Port                  string
	Charset               string
	ScriptPrefix          string
	CommandStackCharacter string
}

//UpdateGameFormID form id of form update game
const UpdateGameFormID = "formupdategame"

//NewUpdateGameForm create new update game form
func NewUpdateGameForm() *UpdateGameForm {
	form := &UpdateGameForm{}
	form.SetComponentLabels(ui.MapLabels(UpdateGameFormFieldLabels))
	return form
}

func (f *UpdateGameForm) ComponentID() string {
	return UpdateGameFormID
}

//Validate Validate form and return any error if raised.
func (f *UpdateGameForm) Validate() error {

	f.ValidateFieldf(f.Host != "", "Host", "网址不能为空")
	f.ValidateFieldf(f.Port != "", "Port", "端口不能为空")
	f.ValidateFieldf(f.Charset != "", "Charset", "字符编码不能为空")
	if !f.HasError() {
		ok, err := titan.Pangu.IsWorldExist(f.ID)
		if err != nil {
			return err
		}
		f.ValidateFieldf(ok == true, "ID", "游戏未找到")
	}
	return nil
}

//InitWithRequest init  update game form  with http request.
func (f *UpdateGameForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func UpdateGame(t *titan.Titan, data []byte) {
	form := NewUpdateGameForm()
	err := json.Unmarshal(data, form)
	if err != nil {
		return
	}
	err = form.Validate()
	if err != nil {
		return
	}
	errors := form.Errors()
	if len(errors) != 0 {
		t.OnCreateFail(errors)
		return
	}
	t.Locker.Lock()
	w := t.Worlds[form.ID]
	if w == nil {
		t.Locker.Unlock()
		return
	}
	w.SetHost(form.Host)
	w.SetPort(form.Port)
	w.SetCharset(form.Charset)
	w.SetScriptPrefix(form.ScriptPrefix)
	w.SetCommandStackCharacter(form.CommandStackCharacter)
	w.SetName(form.Name)
	t.Locker.Unlock()
	go func() {
		t.OnUpdateSuccess(form.ID)
		t.HandleCmdWorldSettings(form.ID)
		t.ExecClients()
	}()

}
