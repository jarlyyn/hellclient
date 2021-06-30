package main

import (
	"strconv"
	"strings"
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
}

func (t *MclTimer) Convert() *Timer {
	ti := &Timer{}
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
}

func (a *MclAlias) Convert() *Alias {
	al := &Alias{}
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
}

func (t *MclTrigger) Convert() *Trigger {
	tr := &Trigger{}
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
	World    *MclWorld    `xml:"world"`
	Triggers *MclTriggers `xml:"triggers"`
	Aliases  *MclAliases  `xml:"aliases"`
	Timers   *MclTimers   `xml:"timers"`
}

func (m *Mcl) ToScriptData() *ScriptData {
	data := &ScriptData{}
	data.OnDisconnect = m.World.OnWorldDisconnect
	data.Type = strings.ToLower(m.World.ScriptLanguage)
	data.Triggers = m.Triggers.Convert()
	data.Aliases = m.Aliases.Convert()
	data.Timers = m.Timers.Convert()
	return data
}
