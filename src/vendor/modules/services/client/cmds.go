package client

const CommandTypeNewLine = "newline"
const CommandTypeAllLines = "alllines"
const CommandTypeUpdatePrompt = "updatePrompt"

type Command struct {
	Type string
	Room string
	Data interface{}
}
