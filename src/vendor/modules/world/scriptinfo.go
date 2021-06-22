package world

import (
	"github.com/herb-go/util"
)

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

type ScriptData struct {
	Type         string
	Desc         string
	OnOpen       string
	OnClose      string
	OnConnect    string
	OnDisconnect string
	Triggers     []*Trigger
	Timers       []*Timer
	Alias        []*Alias
}

func (d *ScriptData) ConvertInfo(id string) *ScriptInfo {
	info := &ScriptInfo{
		ID: id,
	}
	if d != nil {
		info.Type = d.Type
		info.OnOpen = d.OnOpen
		info.OnClose = d.OnClose
		info.OnConnect = d.OnConnect
		info.OnDisconnect = d.OnDisconnect
	}
	return info
}
func NewScriptData() *ScriptData {
	return &ScriptData{}
}

type ScriptInfo struct {
	ID           string
	Desc         string
	Type         string
	OnOpen       string
	OnClose      string
	OnConnect    string
	OnDisconnect string
}
