package world

import "github.com/herb-go/uniqueid"

const TriggerFlagEnabled = 1
const TriggerFlagOmitFromLog = 2
const TriggerFlagOmitFromOutput = 4
const TriggerFlagKeepEvaluating = 8
const TriggerFlagIgnoreCase = 16
const TriggerFlagRegularExpression = 32
const TriggerFlagExpandVariables = 512
const TriggerFlagReplace = 1024
const TriggerFlagTemporary = 16384
const TriggerFlagLowercaseWildcard = 2048
const TriggerFlagOneShot = 32768

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

func (t *Trigger) ByUser() bool {
	return t.byuser
}
func (t *Trigger) SetByUser(v bool) {
	t.byuser = v
}
func (t *Trigger) PrefixedName() string {
	if t.byuser {
		return PrefixByUser + t.Name
	}
	return PrefixByScript + t.Name
}

func NewTrigger() *Trigger {
	return &Trigger{}
}

func CreateTrigger() *Trigger {
	return &Trigger{
		ID:       uniqueid.MustGenerateID(),
		Sequence: 100,
	}
}

type Triggers []*Trigger

// Len is the number of elements in the collection.
func (t Triggers) Len() int {
	return len(t)
}

// Less reports whether the element with index i
func (t Triggers) Less(i, j int) bool {
	if t[i].Sequence != t[j].Sequence {
		return t[i].Sequence < t[j].Sequence
	}

	return t[i].ID < t[j].ID

}

// Swap swaps the elements with indexes i and j.
func (t Triggers) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
