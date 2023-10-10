package forms

import (
	"encoding/json"
	"modules/world"
	"modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

// UpdateTriggerFormFieldLabels form field labels map.
var UpdateTriggerFormFieldLabels = map[string]string{}

// UpdateTriggerForm form struct for update game
type UpdateTriggerForm struct {
	formdata.Form
	World             string
	ByUser            bool
	ID                string
	Name              string
	Enabled           bool
	Match             string
	Send              string
	Script            string
	SendTo            int
	Sequence          int
	ExpandVariables   bool
	Temporary         bool
	OneShot           bool
	Regexp            bool
	Group             string
	Variable          string
	IgnoreCase        bool
	KeepEvaluating    bool
	Menu              bool
	OmitFromLog       bool
	ReverseSpeedwalk  bool
	OmitFromOutput    bool
	MultiLine         bool
	Repeat            bool
	LinesToMatch      int
	WildcardLowerCase bool
}

// UpdateTriggerFormID form id of form update trigger
const UpdateTriggerFormID = "formupdatetrigger"

// NewUpdateTriggerForm update new update game form
func NewUpdateTriggerForm() *UpdateTriggerForm {
	form := &UpdateTriggerForm{}
	form.SetComponentLabels(ui.MapLabels(UpdateTriggerFormFieldLabels))
	return form
}

func (f *UpdateTriggerForm) ComponentID() string {
	return UpdateTriggerFormID
}

// Validate Validate form and return any error if raised.
func (f *UpdateTriggerForm) Validate() error {
	f.ValidateFieldf(f.SendTo >= world.SendToMin && f.SendTo <= world.SendToMax, "SendTo", "发送到无效")
	f.ValidateFieldf(f.Match != "", "Match", "匹配无效")
	if !f.HasError() {
		f.ValidateFieldf(f.Name == "" || world.IDRegexp.MatchString(f.Name), "Name", "名称不可用")
	}
	return nil
}

// InitWithRequest init  update trigger form  with http request.
func (f *UpdateTriggerForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func UpdateTrigger(t *titan.Titan, data []byte) {
	form := NewUpdateTriggerForm()
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
	trigger := world.CreateTrigger()
	trigger.ID = form.ID
	trigger.Name = form.Name
	trigger.Enabled = form.Enabled
	trigger.Match = form.Match
	trigger.Send = form.Send
	trigger.Script = form.Script
	trigger.SendTo = form.SendTo
	trigger.Sequence = form.Sequence
	trigger.Variable = form.Variable
	trigger.ExpandVariables = form.ExpandVariables
	trigger.Temporary = form.Temporary
	trigger.OneShot = form.OneShot
	trigger.Regexp = form.Regexp
	trigger.Group = form.Group
	trigger.IgnoreCase = form.IgnoreCase
	trigger.KeepEvaluating = form.KeepEvaluating
	trigger.OmitFromLog = form.OmitFromLog
	trigger.OmitFromOutput = form.OmitFromOutput
	trigger.MultiLine = form.MultiLine
	trigger.LinesToMatch = form.LinesToMatch
	trigger.WildcardLowerCase = form.WildcardLowerCase
	trigger.Repeat = form.Repeat

	result := t.DoUpdateTrigger(form.World, trigger)
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
		t.OnUpdateTriggerSuccess(form.World, trigger.ID)
		t.HandleCmdTriggers(form.World, true)
		t.HandleCmdTriggers(form.World, false)

	}()

}
