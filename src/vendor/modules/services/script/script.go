package script

import "modules/services/mapper"

type Script struct {
	Triggers  *Triggers
	Variables map[string]string
	Mapper    *mapper.Mapper
}

func New() *Script {
	return &Script{
		Triggers:  NewTriggers(),
		Variables: map[string]string{},
		Mapper:    mapper.New(),
	}
}
