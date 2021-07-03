package world

type Command struct {
	Mesasge     string
	Echo        bool
	Queue       bool
	Log         bool
	History     bool
	Creator     string
	CreatorType string
}

func CreateCommand(message string) *Command {
	return &Command{
		Mesasge: message,
		Echo:    true,
	}
}
