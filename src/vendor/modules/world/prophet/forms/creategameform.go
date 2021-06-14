package forms

import (
	"encoding/json"
	"modules/world/titan"
	"net/http"
	"regexp"

	"modules/world/bus"

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
	world   *bus.Bus
}

//CreateGameFormID form id of form create game
const CreateGameFormID = "formcreategame"

//NewCreateGameForm create new create game form
func NewCreateGameForm() *CreateGameForm {
	form := &CreateGameForm{}
	form.SetComponentLabels(ui.MapLabels(CreateGameFormFieldLabels))
	return form
}

var createFormIDReg = regexp.MustCompile(`^[0-9a-zA-Z\-\_\@\.\[\]\(\)\+]*$`)

func (f *CreateGameForm) ComponentID() string {
	return CreateGameFormID
}

//Validate Validate form and return any error if raised.
func (f *CreateGameForm) Validate() error {
	f.ValidateFieldf(len(f.ID) > 2, "ID", "名称至少需要2个字符")
	f.ValidateFieldf(len(f.ID) < 64, "ID", "名称不能超过64个字符")

	f.ValidateFieldf(createFormIDReg.MatchString(f.ID), "ID", "名称只能包含数字，字母，- _ @ .()[]+")
	f.ValidateFieldf(f.Host != "", "Host", "网址不能为空")
	f.ValidateFieldf(f.Port != "", "Port", "端口不能为空")
	f.ValidateFieldf(f.Charset != "", "Charset", "字符编码不能为空")
	if !f.HasError() {
		f.world = titan.Pangu.NewWorld(f.ID)
		f.ValidateFieldf(f.world == nil, "ID", "名称已经存在")
	}
	return nil
}

//InitWithRequest init  create game form  with http request.
func (f *CreateGameForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func CreateGame(data []byte) {
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
		titan.Pangu.OnCreateFail(errors)
		return
	}
	form.world.SetHost(form.Host)
	form.world.SetPort(form.Port)
	form.world.SetCharset(form.Charset)
	err = form.world.DoSave()
	if err != nil {
		return
	}
	go func() {
		titan.Pangu.OnCreateSuccess(form.ID)
		titan.Pangu.ExecClients()
	}()
}
