package world

import (
	"github.com/herb-go/util"
)

const ScriptTypeNone = ""
const ScriptTypeLua = "lua"
const ScriptTypeJavascript = "jscript"

var AvailableScriptTypes = map[string]bool{
	ScriptTypeJavascript: true,
	ScriptTypeLua:        true,
}

var ScriptTomlTemplates = map[string]string{}
var ScriptTemplates = map[string]string{}
var ScriptTargets = map[string]string{}

func initTemplates() {
	ScriptTomlTemplates[ScriptTypeLua] = util.System("template", "script", "lua.toml")
	ScriptTemplates[ScriptTypeLua] = util.System("template", "script", "main.lua")
	ScriptTargets[ScriptTypeLua] = "main.lua"
	ScriptTomlTemplates[ScriptTypeJavascript] = util.System("template", "script", "jscript.toml")
	ScriptTemplates[ScriptTypeJavascript] = util.System("template", "script", "main.js")
	ScriptTargets[ScriptTypeJavascript] = "main.js"
}

type ScriptSettings struct {
	Name         string
	Type         string
	Intro        string
	Desc         string
	OnOpen       string
	OnClose      string
	OnConnect    string
	OnDisconnect string
	OnBroadcast  string
	OnResponse   string
	OnHUDClick   string
	OnAssist     string
	OnBuffer     string
	OnBufferMin  int
	OnBufferMax  int
	OnSubneg     string
	OnFocus      string
	Channel      string
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
	OnFocus        string
	OnHUDClick     string
	OnResponse     string
	OnAssist       string
	OnBuffer       string
	OnSubneg       string
	OnBufferMin    int
	OnBufferMax    int
	Channel        string
	Triggers       []*Trigger
	Timers         []*Timer
	Aliases        []*Alias
	RequiredParams []*RequiredParam
}

func (d *ScriptData) SetRequiredParams(p []*RequiredParam) {
	d.RequiredParams = make([]*RequiredParam, 0, len(p))
	for _, v := range p {
	LOOP:
		for _, param := range d.RequiredParams {
			if v.Name == param.Name {
				continue LOOP
			}
		}
		d.RequiredParams = append(d.RequiredParams, v)
	}
}
func (d *ScriptData) ConvertSettings(name string) *ScriptSettings {
	settings := &ScriptSettings{}
	if d != nil {
		settings.Name = name
		settings.Type = d.Type
		settings.Intro = d.Intro
		settings.Desc = d.Desc
		settings.OnOpen = d.OnOpen
		settings.OnClose = d.OnClose
		settings.OnConnect = d.OnConnect
		settings.OnDisconnect = d.OnDisconnect
		settings.OnBroadcast = d.OnBroadcast
		settings.OnResponse = d.OnResponse
		settings.OnAssist = d.OnAssist
		settings.Channel = d.Channel
		settings.OnHUDClick = d.OnHUDClick
		settings.OnBuffer = d.OnBuffer
		settings.OnBufferMin = d.OnBufferMin
		settings.OnBufferMax = d.OnBufferMax
		settings.OnFocus = d.OnFocus
		settings.OnSubneg = d.OnSubneg

	}
	return settings
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
		info.OnAssist = d.OnAssist
		info.OnBroadCast = d.OnResponse
		info.OnResponse = d.OnResponse
		info.OnBuffer = d.OnBuffer
		info.OnBufferMin = d.OnBufferMin
		info.OnBufferMax = d.OnBufferMax
		info.OnSubneg = d.OnSubneg
		info.OnHUDClick = d.OnHUDClick
		info.OnFocus = d.OnFocus
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
	OnAssist     string
	OnBroadCast  string
	OnResponse   string
	OnHUDClick   string
	OnBuffer     string
	OnBufferMin  int
	OnBufferMax  int
	OnFocus      string
	OnSubneg     string
}
