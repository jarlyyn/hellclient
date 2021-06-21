package world

import "github.com/herb-go/util"

const ScriptTypeNone = ""
const ScriptTypeLua = "lua"

var AvailableScriptTypes = map[string]bool{
	ScriptTypeNone: true,
	ScriptTypeLua:  true,
}

var ScriptTomlTemplates = map[string]string{}
var ScriptTemplates = map[string]string{}
var ScriptTargets = map[string]string{}

func initTemplates() {
	ScriptTomlTemplates[ScriptTypeLua] = util.System("template", "script", "lua.toml")
	ScriptTemplates[ScriptTypeLua] = util.System("template", "script", "main.lua")
	ScriptTargets[ScriptTypeLua] = "main.lua"
}

type ScriptInfo struct {
	ID           string
	Type         string
	OnOpen       string
	OnClose      string
	OnConnect    string
	OnDisconnect string
}
