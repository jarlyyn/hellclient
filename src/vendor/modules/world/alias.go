package world

import "github.com/herb-go/uniqueid"

const AliasFlagEnabled = 1
const AliasFlagKeepEvaluating = 8
const AliasFlagIgnoreAliasCase = 32
const AliasFlagOmitFromLogFile = 64
const AliasFlagRegularExpression = 128
const AliasFlagExpandVariables = 512
const AliasFlagReplace = 1024
const AliasFlagAliasSpeedWalk = 2048
const AliasFlagAliasQueue = 4096
const AliasFlagAliasMenu = 8192
const AliasFlagTemporary = 16384

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
	byuser                 bool
}

func (a *Alias) ByUser() bool {
	return a.byuser
}
func (a *Alias) SetByUser(v bool) {
	a.byuser = v
}
func (a *Alias) PrefixedName() string {
	if a.byuser {
		return PrefixByUser + a.Name
	}
	return PrefixByScript + a.Name
}

func NewAlias() *Alias {
	return &Alias{}
}

func CreateAlias() *Alias {
	return &Alias{
		ID: uniqueid.MustGenerateID(),
	}
}

type Aliases []*Alias

// Len is the number of elements in the collection.
func (a Aliases) Len() int {
	return len(a)
}

// Less reports whether the element with index i
func (a Aliases) Less(i, j int) bool {
	if a[i].Sequence != a[j].Sequence {
		return a[i].Sequence < a[j].Sequence
	}

	return a[i].ID < a[j].ID

}

// Swap swaps the elements with indexes i and j.
func (a Aliases) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
