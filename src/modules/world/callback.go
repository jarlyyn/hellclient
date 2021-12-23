package world

type Callback struct {
	ID     string
	Name   string
	Script string
	Code   int
	Data   string
}

func NewCallback() *Callback {
	return &Callback{}
}

func CreateCallback(name string, script string, id string) *Callback {
	return &Callback{
		Name:   name,
		Script: script,
		ID:     id,
	}
}
