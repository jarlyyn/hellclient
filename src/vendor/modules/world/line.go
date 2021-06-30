package world

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

const LineTypePrint = 0
const LineTypeSystem = 1
const LineTypeReal = 2
const LineTypeEcho = 3
const LineTypePrompt = 4

type Line struct {
	Words          []Word
	ID             string
	Time           int64
	Type           int
	OmitFromLog    bool
	OmitFromOutput bool
}

func (l *Line) Append(w Word) {
	l.Words = append(l.Words, w)
}
func (l *Line) Plain() string {
	if l == nil {
		return ""
	}
	var output string
	for k := range l.Words {
		output = output + l.Words[k].Text
	}
	return output
}
func (l *Line) IsEmpty() bool {
	return l == nil || len(l.Words) == 0
}
func NewLine() *Line {
	return &Line{
		Words: []Word{},
		ID:    uniqueid.MustGenerateID(),
		Time:  time.Now().Unix(),
	}
}
