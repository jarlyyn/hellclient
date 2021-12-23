package world

type FoundHistory struct {
	Position int
	Command  string
}

func CreateFoundHistory(Position int, Command string) *FoundHistory {
	return &FoundHistory{
		Position: Position,
		Command:  Command,
	}
}
