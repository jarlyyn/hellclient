package overseers

import (
	overseer "github.com/herb-go/herb-drivers/overseers/persistoverseer"
	"github.com/herb-go/herb/persist"
	worker "github.com/herb-go/worker"
)

//PersistFactoryWorker empty persist worker.
var PersistFactoryWorker func(loader func(v interface{}) error) (persist.Store, error)

//PersistFactoryOverseer cache overseer
var PersistFactoryOverseer = worker.NewOrverseer("persist", &PersistFactoryWorker)

func init() {
	PersistFactoryOverseer.WithInitFunc(func(t *worker.OverseerTranning) error {
		return overseer.New().ApplyTo(PersistFactoryOverseer)
	})
	worker.Appoint(PersistFactoryOverseer)
}
