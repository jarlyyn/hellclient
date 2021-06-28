package forms

import (
	"encoding/json"
	"modules/world"
	"modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//CreateTimerFormFieldLabels form field labels map.
var CreateTimerFormFieldLabels = map[string]string{}

//CreateTimerForm form struct for create game
type CreateTimerForm struct {
	formdata.Form
	World                 string
	ByUser                bool
	Hour                  int
	Minute                int
	Second                int
	Name                  string
	SendTo                int
	Send                  string
	Script                string
	Group                 string
	Variable              string
	AtTime                bool
	Enabled               bool
	ActionWhenDisconnectd bool
	OneShot               bool
	Temporary             bool
	OmitFromOutput        bool
	OmitFromLog           bool
}

//CreateTimerFormID form id of form create timer
const CreateTimerFormID = "formcreatetimer"

//NewCreateTimerForm create new create game form
func NewCreateTimerForm() *CreateTimerForm {
	form := &CreateTimerForm{}
	form.SetComponentLabels(ui.MapLabels(CreateTimerFormFieldLabels))
	return form
}

func (f *CreateTimerForm) ComponentID() string {
	return CreateTimerFormID
}

//Validate Validate form and return any error if raised.
func (f *CreateTimerForm) Validate() error {
	f.ValidateFieldf(f.SendTo >= world.SendToMin && f.SendTo <= world.SendToMin, "SendTo", "发送到无效")
	f.ValidateFieldf((f.Hour != 0 || f.Minute != 0 || f.Second != 0) || f.AtTime, "Second", "时间无效")
	if !f.HasError() {
		f.ValidateFieldf(f.Name == "" || titan.Pangu.IsNameAvaliable(f.World, f.Name, f.ByUser), "Name", "名称不可用")
	}
	return nil
}

//InitWithRequest init  create timer form  with http request.
func (f *CreateTimerForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func CreateTimer(t *titan.Titan, data []byte) {
	form := NewCreateTimerForm()
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
	timer := world.CreateTimer()
	timer.SetByUser(form.ByUser)
	timer.Name = form.Name
	timer.Hour = form.Hour
	timer.Minute = form.Minute
	timer.Second = form.Second
	timer.SendTo = form.SendTo
	timer.Send = form.Send
	timer.Script = form.Script
	timer.Group = form.Group
	timer.Variable = form.Variable
	timer.AtTime = form.AtTime
	timer.Enabled = form.Enabled
	timer.ActionWhenDisconnectd = form.ActionWhenDisconnectd
	timer.OneShot = form.OneShot
	timer.Temporary = form.Temporary
	timer.OmitFromOutput = form.OmitFromOutput
	timer.OmitFromLog = form.OmitFromLog
	ok := t.DoCreateTimer(form.World, timer)
	if !ok {
		form.AddError("Name", "添加失败")
		t.OnCreateFail(form.Errors())
		return

	}
	go func() {
		t.OnCreateTimerSuccess(form.World, timer.ID)
		t.HandleCmdTimers(form.World, form.ByUser)
	}()
}
