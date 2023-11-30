package forms

import (
	"encoding/json"
	"modules/world"
	"modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

// CreateTriggerFormFieldLabels form field labels map.
var CreateTriggerFormFieldLabels = map[string]string{}

// CreateTriggerForm form struct for create game
type CreateTriggerForm struct {
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

// CreateTriggerFormID form id of form create trigger
const CreateTriggerFormID = "formcreatetrigger"

// NewCreateTriggerForm create new create game form
func NewCreateTriggerForm() *CreateTriggerForm {
	form := &CreateTriggerForm{}
	form.SetComponentLabels(ui.MapLabels(CreateTriggerFormFieldLabels))
	return form
}

func (f *CreateTriggerForm) ComponentID() string {
	return CreateTriggerFormID
}

// Validate Validate form and return any error if raised.
func (f *CreateTriggerForm) Validate() error {
	f.ValidateFieldf(f.SendTo >= world.SendToMin && f.SendTo <= world.SendToMax, "SendTo", "发送到无效")
	f.ValidateFieldf(f.Match != "", "Match", "匹配无效")
	if !f.HasError() {
		f.ValidateFieldf(f.Name == "" || world.IDRegexp.MatchString(f.Name) || titan.Pangu.IsTriggerNameAvaliable(f.World, f.Name, f.ByUser), "Name", "名称不可用")
	}
	return nil
}

// InitWithRequest init  create trigger form  with http request.
func (f *CreateTriggerForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func CreateTrigger(t *titan.Titan, data []byte) {
	form := NewCreateTriggerForm()
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
	trigger.SetByUser(form.ByUser)
	trigger.Name = form.Name
	trigger.Enabled = form.Enabled
	trigger.Variable = form.Variable
	trigger.Match = form.Match
	trigger.Send = form.Send
	trigger.Script = form.Script
	trigger.SendTo = form.SendTo
	trigger.Sequence = form.Sequence
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
	ok := t.DoCreateTrigger(form.World, trigger)
	if !ok {
		form.AddError("Name", "添加失败")
		t.OnCreateFail(form.Errors())
		return

	}
	go func() {
		t.OnCreateTriggerSuccess(form.World, trigger.ID)
		t.HandleCmdTriggers(form.World, form.ByUser)
		if form.ByUser {
			go t.AutoSaveWorld(form.World)
		}
	}()
}
