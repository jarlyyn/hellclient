package forms

import (
	"encoding/json"
	"hellclient/modules/world"
	"hellclient/modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//UpdateScriptFormFieldLabels form field labels map.
var UpdateScriptFormFieldLabels = map[string]string{}

//UpdateScriptForm form struct for update script
type UpdateScriptForm struct {
	formdata.Form
	*world.ScriptSettings
	ID string
}

//UpdateScriptFormID form id of form update script
const UpdateScriptFormID = "formupdatescript"

//NewUpdateScriptForm create new update script form
func NewUpdateScriptForm() *UpdateScriptForm {
	form := &UpdateScriptForm{}
	form.SetComponentLabels(ui.MapLabels(UpdateScriptFormFieldLabels))
	return form
}

func (f *UpdateScriptForm) ComponentID() string {
	return UpdateScriptFormID
}

//Validate Validate form and return any error if raised.
func (f *UpdateScriptForm) Validate() error {

	if !f.HasError() {
		ok, err := titan.Pangu.IsWorldExist(f.ID)
		if err != nil {
			return err
		}
		f.ValidateFieldf(ok == true, "ID", "游戏未找到")
	}
	return nil
}

//InitWithRequest init  update script form  with http request.
func (f *UpdateScriptForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func UpdateScript(t *titan.Titan, data []byte) {
	form := NewUpdateScriptForm()
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
	sd := w.GetScriptData()
	if sd == nil {
		return
	}
	sd.Intro = form.Intro
	sd.Desc = form.Desc
	sd.OnOpen = form.OnOpen
	sd.OnClose = form.OnClose
	sd.OnConnect = form.OnConnect
	sd.OnDisconnect = form.OnDisconnect
	sd.OnBroadcast = form.OnBroadcast
	sd.OnResponse = form.OnResponse
	sd.OnAssist = form.OnAssist
	sd.Channel = form.Channel
	t.Locker.Unlock()
	go func() {
		t.OnUpdateScriptSuccess(form.ID)
		t.HandleCmdScriptSettings(form.ID)
		t.ExecClients()
	}()

}
