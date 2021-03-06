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
	Type           string
	Intro          string
	Desc           string
	OnOpen         string
	OnClose        string
	OnConnect      string
	OnDisconnect   string
	OnBroadcast    string
	Channel        string
	Triggers       []*Trigger
	Timers         []*Timer
	Aliases        []*Alias
	RequiredParams []*RequiredParam
}

func (d *ScriptData) ConvertInfo(id string) *ScriptInfo {
	info := &ScriptInfo{
		ID: id,
	}
	if d != nil {
		info.Type = d.Type
		info.Intro = d.Intro
		info.Desc = d.Desc
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
	Intro        string
	Type         string
	OnOpen       string
	OnClose      string
	OnConnect    string
	OnDisconnect string
}
