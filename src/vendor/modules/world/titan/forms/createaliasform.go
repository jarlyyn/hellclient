package forms

import (
	"encoding/json"
	"modules/world"
	"modules/world/titan"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//CreateAliasFormFieldLabels form field labels map.
var CreateAliasFormFieldLabels = map[string]string{}

//CreateAliasForm form struct for create game
type CreateAliasForm struct {
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

//CreateAliasFormID form id of form create alias
const CreateAliasFormID = "formcreatealias"

//NewCreateAliasForm create new create game form
func NewCreateAliasForm() *CreateAliasForm {
	form := &CreateAliasForm{}
	form.SetComponentLabels(ui.MapLabels(CreateAliasFormFieldLabels))
	return form
}

func (f *CreateAliasForm) ComponentID() string {
	return CreateAliasFormID
}

//Validate Validate form and return any error if raised.
func (f *CreateAliasForm) Validate() error {
	f.ValidateFieldf(f.SendTo >= world.SendToMin && f.SendTo <= world.SendToMax, "SendTo", "发送到无效")
	f.ValidateFieldf(f.Match != "", "Match", "别名无效")
	if !f.HasError() {
		f.ValidateFieldf(f.Name == "" || world.IDRegexp.MatchString(f.Name) || titan.Pangu.IsAliasNameAvaliable(f.World, f.Name, f.ByUser), "Name", "名称不可用")
	}
	return nil
}

//InitWithRequest init  create alias form  with http request.
func (f *CreateAliasForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.
	return nil
}

func CreateAlias(t *titan.Titan, data []byte) {
	form := NewCreateAliasForm()
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
	alias.SetByUser(form.ByUser)
	alias.Name = form.Name
	alias.Enabled = form.Enabled
	alias.Variable = form.Variable
	alias.Match = form.Match
	alias.Send = form.Send
	alias.Script = form.Script
	alias.SendTo = form.SendTo
	alias.Sequence = form.Sequence
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
	ok := t.DoCreateAlias(form.World, alias)
	if !ok {
		form.AddError("Name", "添加失败")
		t.OnCreateFail(form.Errors())
		return

	}
	go func() {
		t.OnCreateAliasSuccess(form.World, alias.ID)
		t.HandleCmdAliases(form.World, form.ByUser)
	}()
}
