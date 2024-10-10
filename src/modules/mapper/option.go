package mapper

type Option struct {
	Blacklist   []string
	Whitelist   []string
	BlockedPath [][]string
}

func NewOption() *Option {
	return &Option{
		Blacklist:   []string{},
		Whitelist:   []string{},
		BlockedPath: [][]string{},
	}
}
