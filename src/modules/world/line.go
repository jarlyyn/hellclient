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

func (w *Word) Inherit() *Word {
	if w == nil {
		return &Word{}
	}
	return &Word{
		Text:       "",
		Color:      w.Color,
		Background: w.Background,
		Bold:       w.Bold,
		Underlined: w.Underlined,
		Blinking:   w.Blinking,
		Inverse:    w.Inverse,
	}
}
func (w *Word) GetColorRGB() int {
	return Colours[w.Color]
}
func (w *Word) GetBGColorRGB() int {
	return Colours[w.Background]
}

func NewWord() *Word {
	return &Word{}
}

//通过Print打印
const LineTypePrint = 0

//系统信息
const LineTypeSystem = 1

//收到的真实信息
const LineTypeReal = 2

//输入回显
const LineTypeEcho = 3

//输入行类型
const LineTypePrompt = 4

//发出的本地广播
const LineTypeLocalBroadcastOut = 5

//发出的全局广播
const LineTypeGlobalBroadcastOut = 6

//收到的本地广播
const LineTypeLocalBroadcastIn = 7

//收到的全局广播
const LineTypeGlobalBroadcastIn = 8

//Websocket发出的请求的信息
const LineTypeRequest = 9

//Websocket收到的响应的信息
const LineTypeResponse = 10

type Line struct {
	Words          []*Word
	ID             string
	Time           int64
	Type           int
	OmitFromLog    bool
	OmitFromOutput bool
	Triggers       []string
	CreatorType    string
	Creator        string
}

func (l *Line) Append(w *Word) {
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
	var result = 0
	for k, v := range l.Words {
		if k < idx-1 {
			result = result + len([]rune(v.Text))
			continue
		}
		break
	}
	return result
}
func NewLine() *Line {
	return &Line{
		Words: []*Word{},
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
	"Bright-Black":    0x7f7f7f,
	"Bright-Red":      0xff0000,
	"Bright-Green":    0x00fc00,
	"Bright-Yellow":   0xffff00,
	"Bright-Blue":     0x0000fc,
	"Bright-Magenta":  0xff00ff,
	"Bright-Cyan":     0x00ffff,
	"Bright-White":    0xffffff,
}
var NamedColor = map[string]int{
	"black":   Colours["Black"],
	"red":     Colours["Red"],
	"green":   Colours["Green"],
	"yellow":  Colours["Yellow"],
	"blue":    Colours["Blue"],
	"magenta": Colours["Magenta"],
	"cyan":    Colours["Cyan"],
	"white":   Colours["White"],
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
	return -1
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
