package mapper

type Option struct {
	Blacklist []string
	Whitelist []string
}

func NewOption() *Option {
	return &Option{
		Blacklist: []string{},
		Whitelist: []string{},
	}
}
