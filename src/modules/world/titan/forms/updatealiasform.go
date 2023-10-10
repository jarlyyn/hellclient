package forms

import (
	"encoding/json"
	"modules/world"
	"modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

// UpdateAliasFormFieldLabels form field labels map.
var UpdateAliasFormFieldLabels = map[string]string{}

// UpdateAliasForm form struct for update game
type UpdateAliasForm struct {
	formdata.Form
	World            string
	ByUser           bool
	ID               string
	Name             string
	Enabled          bool
	Match            string
	Send             string
	Script           string
	SendTo           int
	Sequence         int
	ExpandVariables  bool
	Temporary        bool
	OneShot          bool
	Regexp           bool
	Group            string
	Variable         string
	IgnoreCase       bool
	KeepEvaluating   bool
	Menu             bool
	OmitFromLog      bool
	ReverseSpeedwalk bool
	OmitFromOutput   bool
}

// UpdateAliasFormID form id of form update alias
const UpdateAliasFormID = "formupdatealias"

// NewUpdateAliasForm update new update game form
func NewUpdateAliasForm() *UpdateAliasForm {
	form := &UpdateAliasForm{}
	form.SetComponentLabels(ui.MapLabels(UpdateAliasFormFieldLabels))
	return form
}

func (f *UpdateAliasForm) ComponentID() string {
	return UpdateAliasFormID
}

// Validate Validate form and return any error if raised.
func (f *UpdateAliasForm) Validate() error {
	f.ValidateFieldf(f.SendTo >= world.SendToMin && f.SendTo <= world.SendToMax, "SendTo", "发送到无效")
	f.ValidateFieldf(f.Match != "", "Match", "别名无效")
	if !f.HasError() {
		f.ValidateFieldf(f.Name == "" || world.IDRegexp.MatchString(f.Name), "Name", "名称不可用")
	}
	return nil
}

// InitWithRequest init  update alias form  with http request.
func (f *UpdateAliasForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func UpdateAlias(t *titan.Titan, data []byte) {
	form := NewUpdateAliasForm()
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
	alias := world.CreateAlias()
	alias.ID = form.ID
	alias.Name = form.Name
	alias.Enabled = form.Enabled
	alias.Match = form.Match
	alias.Send = form.Send
	alias.Script = form.Script
	alias.SendTo = form.SendTo
	alias.Sequence = form.Sequence
	alias.Variable = form.Variable
	alias.ExpandVariables = form.ExpandVariables
	alias.Temporary = form.Temporary
	alias.OneShot = form.OneShot
	alias.Regexp = form.Regexp
	alias.Group = form.Group
	alias.IgnoreCase = form.IgnoreCase
	alias.KeepEvaluating = form.KeepEvaluating
	alias.Menu = form.Menu
	alias.OmitFromLog = form.OmitFromLog
	alias.ReverseSpeedwalk = form.ReverseSpeedwalk
	alias.OmitFromOutput = form.OmitFromOutput
	result := t.DoUpdateAlias(form.World, alias)
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
		t.OnUpdateAliasSuccess(form.World, alias.ID)
		t.HandleCmdAliases(form.World, true)
		t.HandleCmdAliases(form.World, false)

	}()

}
