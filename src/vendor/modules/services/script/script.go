package script

type Script struct {
	Triggers *Triggers
}

func New() *Script {
	return &Script{
		Triggers: NewTriggers(),
	}
}
