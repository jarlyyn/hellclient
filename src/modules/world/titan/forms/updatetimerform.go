package forms

import (
	"encoding/json"
	"modules/world"
	"modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

// UpdateTimerFormFieldLabels form field labels map.
var UpdateTimerFormFieldLabels = map[string]string{}

// UpdateTimerForm form struct for create game
type UpdateTimerForm struct {
	formdata.Form
	World                 string
	ID                    string
	Hour                  int
	Minute                int
	Second                float64
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

// UpdateTimerFormID form id of form create timer
const UpdateTimerFormID = "formUpdateTimer"

// NewUpdateTimerForm create new create game form
func NewUpdateTimerForm() *UpdateTimerForm {
	form := &UpdateTimerForm{}
	form.SetComponentLabels(ui.MapLabels(UpdateTimerFormFieldLabels))
	return form
}

func (f *UpdateTimerForm) ComponentID() string {
	return UpdateTimerFormID
}

// Validate Validate form and return any error if raised.
func (f *UpdateTimerForm) Validate() error {
	f.ValidateFieldf(f.ID != "", "ID", "无效的ID")
	f.ValidateFieldf(f.SendTo >= world.SendToMin && f.SendTo <= world.SendToMax, "SendTo", "发送到无效")
	f.ValidateFieldf((f.Hour != 0 || f.Minute != 0 || f.Second != 0) || f.AtTime, "Second", "时间无效")
	if !f.HasError() {
		f.ValidateFieldf(f.Name == "" || world.IDRegexp.MatchString(f.Name), "Name", "名称不可用")
	}
	return nil
}

// InitWithRequest init  create timer form  with http request.
func (f *UpdateTimerForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func UpdateTimer(t *titan.Titan, data []byte) {
	form := NewUpdateTimerForm()
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
	timer.ID = form.ID
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
	result := t.DoUpdateTimer(form.World, timer)
	if result != world.UpdateOK {
		switch result {
		case world.UpdateFailDuplicateName:
			form.AddError("Name", "名称重复")
		case world.UpdateFailNotFound:
			form.AddError("ID", "未找到")
		}
		t.OnCreateFail(form.Errors())
		return

	}
	go func() {
		t.OnUpdateTimerSuccess(form.World, timer.ID)
		t.HandleCmdTimers(form.World, true)
		t.HandleCmdTimers(form.World, false)

	}()
}
