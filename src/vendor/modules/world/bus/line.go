package bus

import "time"

type Word struct {
	Text       string
	Color      string
	Background string
	Bold       bool
}

type Line struct {
	Words    []Word
	Time     int64
	IsPrint  bool
	IsSystem bool
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
		Time:     time.Now().Unix(),
		IsPrint:  false,
		IsSystem: false,
	}
}
