package bus

import (
	"time"

	"github.com/herb-go/uniqueid"
)

type Word struct {
	Text       string
	Color      string
	Background string
	Bold       bool
}

type Line struct {
	Words          []Word
	ID             string
	Time           int64
	IsReal         bool
	IsPrint        bool
	IsSystem       bool
	OmitFromLog    bool
	OmitFromOutput bool
}

func (l *Line) Append(w Word) {
	l.Words = append(l.Words, w)
}
func (l *Line) Plain() string {
	var output string
	for k := range l.Words {
		output = output + l.Words[k].Text
	}
	return output
}

func NewLine() *Line {
	return &Line{
		Words:    []Word{},
		ID:       uniqueid.MustGenerateID(),
		Time:     time.Now().Unix(),
		IsPrint:  false,
		IsSystem: false,
	}
}
