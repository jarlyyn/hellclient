package forms

import (
	"hellclient/modules/world"

	"github.com/herb-go/herb/ui/validator/formdata"
)

//RequiredParamsForm form struct for update game
type RequiredParamsForm struct {
	formdata.Form
	Current        string
	RequiredParams []*world.RequiredParam
}
