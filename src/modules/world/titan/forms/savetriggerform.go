package forms

// import (
// 	"encoding/json"
// 	"hellclient/modules/services/client"
// 	"hellclient/modules/services/script"
// 	"net/http"
// 	"time"

// 	uuid "github.com/satori/go.uuid"

// 	"github.com/herb-go/herb/ui"
// 	"github.com/herb-go/herb/ui/validator/formdata"
// )

// //SaveTriggerFormFieldLabels form field labels map.
// var SaveTriggerFormFieldLabels = map[string]string{
// 	"Pattern":  "匹配项",
// 	"IsRegExp": "正则",
// 	"Enabled":  "有效",
// 	"Commnd":   "命令",
// 	"Priority": "优先度",
// 	"Finally":  "不再继续匹配",
// }

// //SaveTriggerForm form struct for save trigger
// type SaveTriggerForm struct {
// 	Name string
// 	formdata.Form
// 	Pattern  string
// 	IsRegExp bool
// 	Enabled  bool
// 	Command  string
// 	Priority int
// 	Finally  bool
// }

// //SaveTriggerFormID form id of form save trigger
// const SaveTriggerFormID = "formservices.savetrigger"

// //NewSaveTriggerForm create new save trigger form
// func NewSaveTriggerForm() *SaveTriggerForm {
// 	form := &SaveTriggerForm{}
// 	form.SetComponentLabels(ui.MapLabels(SaveTriggerFormFieldLabels))
// 	return form
// }

// func (f *SaveTriggerForm) ComponentID() string {
// 	return SaveTriggerFormID
// }

// //Validate Validate form and return any error if raised.
// func (f *SaveTriggerForm) Validate() error {
// 	f.ValidateFieldf(f.Pattern != "", "Pattern", "匹配项不能为空")

// 	if !f.HasError() {
// 	}
// 	return nil
// }

// //InitWithRequest init  save trigger form  with http request.
// func (f *SaveTriggerForm) InitWithRequest(r *http.Request) error {
// 	//Put your request form code here.
// 	//such as get current user id or ip address.
// 	return nil
// }

// func SaveTrigger(current string, data []byte) {

// 	form := NewSaveTriggerForm()
// 	err := json.Unmarshal(data, form)
// 	if err != nil {
// 		return
// 	}
// 	err = form.Validate()
// 	if err != nil {
// 		return
// 	}
// 	errors := form.Errors()
// 	if len(errors) != 0 {
// 		client.DefaultManager.OnTriggerFail(current, errors)
// 		return
// 	}
// 	t := &script.WorldTrigger{}
// 	t.Command = form.Command
// 	t.Enabled = form.Enabled
// 	t.CreatedTime = time.Now().Unix()
// 	t.Finally = form.Finally
// 	t.IsRegExp = form.IsRegExp
// 	t.Pattern = form.Pattern
// 	t.Priority = form.Priority
// 	t.Name = form.Name
// 	if t.Name == "" {
// 		id, err := uuid.NewV1()
// 		if err != nil {
// 			return
// 		}
// 		t.Name = id.String()
// 	}
// 	cli := client.DefaultManager.Client(current)
// 	if cli == nil {
// 		return
// 	}
// 	var found bool
// 	cli.Lock.Lock()
// 	defer cli.Lock.Unlock()
// 	for k := range cli.World.Triggers {
// 		if cli.World.Triggers[k].Name == t.Name {
// 			found = true
// 			cli.World.Triggers[k] = t
// 			break
// 		}
// 	}
// 	if !found {
// 		cli.World.Triggers = append(cli.World.Triggers, t)
// 	}
// 	cli.Script.Triggers.Add(t.Trigger())
// 	err = cli.Save()
// 	if err != nil {
// 		return
// 	}
// 	go func() {
// 		client.DefaultManager.OnTriggerSuccess(current)
// 		client.DefaultManager.ExecTriggers(current)
// 	}()
// }
