package world

const ScriptTypeNone = ""
const ScriptTypeLua = "lua"

var AvailableScriptTypes = map[string]bool{
	ScriptTypeNone: true,
	ScriptTypeLua:  true,
}

type ScriptInfo struct {
	ID           string
	Type         string
	OnOpen       string
	OnClose      string
	OnConnect    string
	OnDisconnect string
}
