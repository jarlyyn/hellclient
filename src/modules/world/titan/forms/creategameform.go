package forms

import (
	"encoding/json"
	"hellclient/modules/world"
	"hellclient/modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//CreateGameFormFieldLabels form field labels map.
var CreateGameFormFieldLabels = map[string]string{
	"ID":      "名称",
	"Host":    "网址",
	"Port":    "端口",
	"Charset": "字符编码",
}

//CreateGameForm form struct for create game
type CreateGameForm struct {
	formdata.Form
	ID      string
	Host    string
	Port    string
	Charset string
}

//CreateGameFormID form id of form create game
const CreateGameFormID = "formcreategame"

//NewCreateGameForm create new create game form
func NewCreateGameForm() *CreateGameForm {
	form := &CreateGameForm{}
	form.SetComponentLabels(ui.MapLabels(CreateGameFormFieldLabels))
	return form
}

func (f *CreateGameForm) ComponentID() string {
	return CreateGameFormID
}

//Validate Validate form and return any error if raised.
func (f *CreateGameForm) Validate() error {
	f.ValidateFieldf(len(f.ID) > 2, "ID", "名称至少需要2个字符")
	f.ValidateFieldf(len(f.ID) < 64, "ID", "名称不能超过64个字符")

	f.ValidateFieldf(world.IDRegexp.MatchString(f.ID), "ID", "名称只能包含数字，字母，- _ @ .()[]+")
	f.ValidateFieldf(f.Host != "", "Host", "网址不能为空")
	f.ValidateFieldf(f.Port != "", "Port", "端口不能为空")
	f.ValidateFieldf(f.Charset != "", "Charset", "字符编码不能为空")
	if !f.HasError() {
		ok, err := titan.Pangu.IsWorldExist(f.ID)
		if err != nil {
			return err
		}
		f.ValidateFieldf(ok == false, "ID", "名称已经存在")
	}
	return nil
}

//InitWithRequest init  create game form  with http request.
func (f *CreateGameForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func CreateGame(t *titan.Titan, data []byte) {
	form := NewCreateGameForm()
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
	w := t.NewWorld(form.ID)
	if w == nil {
		return
	}
	w.SetHost(form.Host)
	w.SetPort(form.Port)
	w.SetCharset(form.Charset)
	go func() {
		t.OnCreateSuccess(form.ID)
		t.ExecClients()
	}()
	err = t.SaveWorld(form.ID)
	if err != nil {
		return
	}
}
