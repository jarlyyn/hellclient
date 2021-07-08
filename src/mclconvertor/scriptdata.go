package main

type RequiredParam struct {
	Name  string
	Desc  string
	Intro string
}

type Timer struct {
	ID                    string
	Name                  string
	Enabled               bool
	Hour                  int
	Minute                int
	Second                int
	Send                  string
	Script                string
	AtTime                bool
	SendTo                int
	ActionWhenDisconnectd bool
	Temporary             bool
	OneShot               bool
	Group                 string
	Variable              string
	OmitFromLog           bool
	OmitFromOutput        bool
}

type Alias struct {
	ID                     string
	Name                   string
	Enabled                bool
	Match                  string
	Send                   string
	Script                 string
	SendTo                 int
	Sequence               int
	ExpandVariables        bool
	Temporary              bool
	OneShot                bool
	Regexp                 bool
	Group                  string
	IgnoreCase             bool
	KeepEvaluating         bool
	Menu                   bool
	OmitFromLog            bool
	Variable               string
	ReverseSpeedwalk       bool
	OmitFromOutput         bool
	OmitFromCommandHistory bool
}

type Trigger struct {
	ID                string
	Name              string
	Enabled           bool
	Match             string
	Send              string
	ColourChangeType  int
	Colour            int
	Wildcard          int
	SoundFileName     string
	SoundIfInactive   bool
	Script            string
	SendTo            int
	Sequence          int
	ExpandVariables   bool
	Temporary         bool
	OneShot           bool
	Regexp            bool
	Repeat            bool
	MultiLine         bool
	LinesToMatch      int
	WildcardLowerCase bool
	Group             string
	IgnoreCase        bool
	KeepEvaluating    bool
	OmitFromLog       bool
	OmitFromOutput    bool
	Inverse           bool
	Italic            bool
	Variable          string
	byuser            bool
}

type ScriptData struct {
	Type           string
	Desc           string
	OnOpen         string
	OnClose        string
	OnConnect      string
	OnDisconnect   string
	Triggers       []*Trigger
	Timers         []*Timer
	Aliases        []*Alias
	RequiredParams []*RequiredParam
}
