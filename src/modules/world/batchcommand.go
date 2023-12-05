package world

type BatchCommand struct {
	Scripts []string
	Command string
}

func NewBatchCommand() *BatchCommand {
	return &BatchCommand{
		Scripts: []string{},
		Command: "",
	}
}

type BatchCommandScripts struct {
	Scripts []string
}

func NewBatchCommandScripts() *BatchCommandScripts {
	return &BatchCommandScripts{
		Scripts: []string{},
	}
}
