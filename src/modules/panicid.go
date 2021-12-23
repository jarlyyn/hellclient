package modules

import (
	"github.com/herb-go/uniqueid"
	"github.com/herb-go/util"
)

func init() {
	util.IDGenerator = uniqueid.TryGenerateID
}
