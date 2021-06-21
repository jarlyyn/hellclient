package world

type Alias struct {
	Name            string
	Enabled         bool
	Match           string
	Send            string
	ScriptName      string
	SendTo          int
	Sequence        int
	ExpandVariables bool
	Temporary       bool
	OneShot         bool
	Regexp          bool
	Repeat          bool
	Group           string
	IgnoreCase      bool
	KeepEvaluating  bool
	OmitFromLog     bool
	OmitFromOutput  bool
}
