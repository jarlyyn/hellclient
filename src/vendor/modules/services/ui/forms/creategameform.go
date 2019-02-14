package forms

import (
	"encoding/json"
	"net/http"
	"regexp"

	"modules/services/client"

	"github.com/herb-go/herb/model/formdata"
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
	form.SetModelID(CreateGameFormID)
	form.SetFieldLabels(CreateGameFormFieldLabels)
	return form
}

var createFormIDReg = regexp.MustCompile(`^[0-9a-zA-Z\-\_\@\.]*$`)

//Validate Validate form and return any error if raised.
func (f *CreateGameForm) Validate() error {
	f.ValidateFieldf(len(f.ID) > 2, "ID", "名称至少需要2个字符")
	f.ValidateFieldf(createFormIDReg.MatchString(f.ID), "ID", "名称只能包含数字，字母，- _ @ .")
	f.ValidateFieldf(f.Host != "", "Host", "网址不能为空")
	f.ValidateFieldf(f.Port != "", "Port", "端口不能为空")
	f.ValidateFieldf(f.Charset != "", "Charset", "字符编码不能为空")
	if !f.HasError() {

		f.ValidateFieldf(client.DefaultManager.Clients[f.ID] == nil, "ID", "名称已经存在")
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
	client.DefaultManager.Lock.Lock()
	defer client.DefaultManager.Lock.Unlock()
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
		client.DefaultManager.OnCreateFail(errors)
		return
	}
	c := client.ClientConfig{}
	c.World.Host = form.Host
	c.World.Port = form.Port
	c.World.Charset = form.Charset
	client.DefaultManager.NewClient(form.ID, c)
	go func() {
		client.DefaultManager.ExecClients()
	}()
}
