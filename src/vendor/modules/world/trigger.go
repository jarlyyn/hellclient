package world

const TriggerFlagEnabled = 1
const TriggerFlagOmitFromLog = 2
const TriggerFlagOmitFromOutput = 4
const TriggerFlagKeepEvaluating = 8
const TriggerFlagIgnoreCase = 16
const TriggerFlagRegularExpression = 32
const TriggerFlagExpandVariables = 512
const TriggerFlagReplace = 1024
const TriggerFlagTemporary = 16384

type Trigger struct {
	Name               string
	Enabled            bool
	Match              string
	Send               string
	Colour             int32
	Wildcard           int32
	SoundFileName      string
	ScriptName         string
	SendTo             int
	Sequence           int
	ExpandVariables    bool
	Temporary          bool
	OneShot            bool
	Regexp             bool
	Repeat             bool
	MutliLine          bool
	WildcardsLowerCase bool
	Group              string
	IgnoreCase         bool
	KeepEvaluating     bool
	OmitFromLog        bool
	OmitFromOutput     bool
}
