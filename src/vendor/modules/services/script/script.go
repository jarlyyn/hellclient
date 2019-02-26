package script

type Script struct {
	Triggers  *Triggers
	Variables map[string]string
}

func New() *Script {
	return &Script{
		Triggers:  NewTriggers(),
		Variables: map[string]string{},
	}
}
