package hired

import (
	"github.com/herb-go/util/local/localpersist"
)

//WORKER(LocalPersist):local persist factory

//LocalPersist persist factory.store data in appdata/persistdata
var LocalPersist = localpersist.Factory

//MyPersistFactory put your own  factory code here
// var MyPersistFactory = func(loader func(v interface{}) error) (persist.Store, error) {
// 	d := NewStore()
// 	err := loader(d)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return d, nil
// }
