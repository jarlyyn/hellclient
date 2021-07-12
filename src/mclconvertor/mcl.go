package main

import (
	"strconv"
	"strings"

	"github.com/herb-go/uniqueid"
)

func MustAtoi(v string) int {
	if v == "" {
		return 0
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}
func MustAtoFloat(v string) float64 {
	if v == "" {
		return 0
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return f
}
func (v *MclVariable) Convert() *RequiredParam {
	return &RequiredParam{
		Name: v.Name,
	}
}

type MclVariable struct {
	Name string `xml:"name,attr"`
}

func (vs *MclVariables) Convert() []*RequiredParam {
	var result = []*RequiredParam{}
	for _, v := range vs.Variables {
		result = append(result, v.Convert())
	}
	return result
}

type MclVariables struct {
	Variables []*MclVariable `xml:"variable"`
}
type MclTimer struct {
	Enabled        string `xml:"enabled,attr"`
	Group          string `xml:"group,attr"`
	Name           string `xml:"name,attr"`
	OmitFromLog    string `xml:"omit_from_log,attr"`
	OmitFromOutput string `xml:"omit_from_output,attr"`
	Second         string `xml:"second,attr"`
	Hour           string `xml:"hour,attr"`
	Minute         string `xml:"minute,attr"`
	AtTime         string `xml:"at_time,attr"`
	Script         string `xml:"script,attr"`
	ActiveClosed   string `xml:"active_closed,attr"`
	Send           string `xml:"send"`
}

func (t *MclTimer) Convert() *Timer {
	ti := &Timer{}
	ti.ID = uniqueid.MustGenerateID()
	ti.Enabled = t.Enabled == "y"
	ti.Group = t.Group
	ti.Name = t.Name
	ti.OmitFromLog = t.OmitFromLog == "y"
	ti.OmitFromOutput = t.OmitFromOutput == "y"
	ti.AtTime = t.AtTime == "y"
	ti.Script = t.Script
	ti.Second = int(MustAtoFloat(t.Second))
	ti.Hour = MustAtoi(t.Hour)
	ti.Minute = MustAtoi(t.Minute)
	ti.ActionWhenDisconnectd = t.ActiveClosed == "y"
	ti.Send = t.Send
	return ti
}

type MclTimers struct {
	Timer []*MclTimer `xml:"timer"`
}

func (a *MclTimers) Convert() []*Timer {
	result := make([]*Timer, 0, len(a.Timer))
	for _, v := range a.Timer {
		result = append(result, v.Convert())
	}
	return result
}

type MclAlias struct {
	Enabled        string `xml:"enabled,attr"`
	Group          string `xml:"group,attr"`
	KeepEvaluating string `xml:"keep_evaluating,attr"`
	Match          string `xml:"match,attr"`
	Name           string `xml:"name,attr"`
	OmitFromLog    string `xml:"omit_from_log,attr"`
	OmitFromOutput string `xml:"omit_from_output,attr"`
	Regexp         string `xml:"regexp,attr"`
	Sequence       string `xml:"sequence,attr"`
	Script         string `xml:"script,attr"`
	Send           string `xml:"send"`
}

func (a *MclAlias) Convert() *Alias {
	al := &Alias{}
	al.ID = uniqueid.MustGenerateID()
	al.Enabled = a.Enabled == "y"
	al.Group = a.Group
	al.KeepEvaluating = a.KeepEvaluating == "y"
	al.Match = a.Match
	al.Name = a.Name
	al.OmitFromLog = a.OmitFromLog == "y"
	al.OmitFromOutput = a.OmitFromOutput == "y"
	al.Regexp = a.Regexp == "y"
	al.Sequence = MustAtoi(a.Sequence)
	al.Script = a.Script
	al.Send = a.Send
	return al
}

type MclAliases struct {
	Alias []*MclAlias `xml:"alias"`
}

func (a *MclAliases) Convert() []*Alias {
	result := make([]*Alias, 0, len(a.Alias))
	for _, v := range a.Alias {
		result = append(result, v.Convert())
	}
	return result
}

type MclTrigger struct {
	Enabled        string `xml:"enabled,attr"`
	Group          string `xml:"group,attr"`
	KeepEvaluating string `xml:"keep_evaluating,attr"`
	Match          string `xml:"match,attr"`
	Name           string `xml:"name,attr"`
	OmitFromLog    string `xml:"omit_from_log,attr"`
	OmitFromOutput string `xml:"omit_from_output,attr"`
	Regexp         string `xml:"regexp,attr"`
	Sequence       string `xml:"sequence,attr"`
	Script         string `xml:"script,attr"`
	Send           string `xml:"send"`
}

func (t *MclTrigger) Convert() *Trigger {
	tr := &Trigger{}
	tr.ID = uniqueid.MustGenerateID()
	tr.Enabled = t.Enabled == "y"
	tr.Group = t.Group
	tr.KeepEvaluating = t.KeepEvaluating == "y"
	tr.Match = t.Match
	tr.Name = t.Name
	tr.OmitFromLog = t.OmitFromLog == "y"
	tr.OmitFromOutput = t.OmitFromOutput == "y"
	tr.Regexp = t.Regexp == "y"
	tr.Sequence = MustAtoi(t.Sequence)
	tr.Script = t.Script
	tr.Send = t.Send
	return tr
}

type MclTriggers struct {
	Trigger []MclTrigger `xml:"trigger"`
}

func (t *MclTriggers) Convert() []*Trigger {
	result := make([]*Trigger, 0, len(t.Trigger))
	for _, v := range t.Trigger {
		result = append(result, v.Convert())
	}
	return result
}

type MclWorld struct {
	OnWorldDisconnect string `xml:"on_world_disconnect,attr"`
	ScriptLanguage    string `xml:"script_language,attr"`
}
type Mcl struct {
	World     *MclWorld     `xml:"world"`
	Triggers  *MclTriggers  `xml:"triggers"`
	Aliases   *MclAliases   `xml:"aliases"`
	Timers    *MclTimers    `xml:"timers"`
	Variables *MclVariables `xml:"variables"`
}

func (m *Mcl) ToScriptData() *ScriptData {
	data := &ScriptData{}
	data.OnDisconnect = m.World.OnWorldDisconnect
	data.Type = strings.ToLower(m.World.ScriptLanguage)
	data.Triggers = m.Triggers.Convert()
	data.Aliases = m.Aliases.Convert()
	data.Timers = m.Timers.Convert()
	data.RequiredParams = m.Variables.Convert()
	return data
}
