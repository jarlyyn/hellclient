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
	Underlined bool
	Blinking   bool
	Inverse    bool
}

func (w *Word) GetColorCode() int {
	return Colours[w.Color]
}
func (w *Word) GetBGColorCode() int {
	return Colours[w.Background]
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
	Triggers       []string
	CreatorType    string
	Creator        string
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
func (l *Line) IsNewline() bool {
	t := l.Plain()
	newline := len(t) > 1 && t[len(t)-1] == '\n'
	return newline
}
func (l *Line) IsEmpty() bool {
	return l == nil || len(l.Words) == 0
}
func (l *Line) GetWordStartColumn(idx int) int {
	if idx < 1 || idx > len(l.Words) {
		return -1
	}
	var result = 1
	for k, v := range l.Words {
		if k < idx {
			result = result + len(v.Text)
			continue
		}
		break
	}
	return result
}
func NewLine() *Line {
	return &Line{
		Words: []Word{},
		ID:    uniqueid.MustGenerateID(),
		Time:  time.Now().Unix(),
	}
}

var Colours = map[string]int{
	"Black":           0x000000,
	"Red":             0x7f0000,
	"Green":           0x009300,
	"Yellow":          0xfc7f00,
	"Blue":            0x00007f,
	"Magenta":         0x9c009c,
	"Cyan":            0x009393,
	"White":           0xd2d2d2,
	"BrightBlack":     0x7f7f7f,
	"BrightRed":       0xff0000,
	"BrightGreen":     0x00fc00,
	"BrightYellow":    0xffff00,
	"BrightBlue":      0x0000fc,
	"BrightMagenta":   0xff00ff,
	"BrightCyan":      0x00ffff,
	"BrightWhite":     0xffffff,
	"BGBlack":         0x000000,
	"BGRed":           0x7f0000,
	"BGGreen":         0x009300,
	"BGYellow":        0xfc7f00,
	"BGBlue":          0x00007f,
	"BGMagenta":       0x9c009c,
	"BGCyan":          0x009393,
	"BGWhite":         0xd2d2d2,
	"BGBrightBlack":   0x7f7f7f,
	"BGBrightRed":     0xff0000,
	"BGBrightGreen":   0x00fc00,
	"BGBrightYellow":  0xffff00,
	"BGBrightBlue":    0x0000fc,
	"BGBrightMagenta": 0xff00ff,
	"BGBrightCyan":    0x00ffff,
	"BGBrightWhite":   0xffffff,
}

func GetNormalColour(code int) int {
	switch code {
	case 1:
		return Colours["Black"]
	case 2:
		return Colours["Red"]
	case 3:
		return Colours["Green"]
	case 4:
		return Colours["Yellow"]
	case 5:
		return Colours["Blue"]
	case 6:
		return Colours["Magenta"]
	case 7:
		return Colours["Cyan"]
	case 8:
		return Colours["White"]
	}
	return 0
}

func GetBoldColour(code int) int {
	switch code {
	case 1:
		return Colours["BrightBlack"]
	case 2:
		return Colours["BrightRed"]
	case 3:
		return Colours["BrightGreen"]
	case 4:
		return Colours["BrightYellow"]
	case 5:
		return Colours["BrightBlue"]
	case 6:
		return Colours["BrightMagenta"]
	case 7:
		return Colours["BrightCyan"]
	case 8:
		return Colours["BrightWhite"]
	}
	return 0
}
