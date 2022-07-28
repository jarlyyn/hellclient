package converter

import (
	"hellclient/modules/world"

	"github.com/jarlyyn/ansi"
)

func ConvertToLine(last *world.Word, msg []byte, charset string, errhandler func(err error) bool) (*world.Line, *world.Word) {
	line := world.NewLine()
	if len(msg) == 0 {
		return line, last
	}
	w := last.Inherit()
	var s *ansi.S
	var err error
	var b []byte
	for len(msg) > 0 {
		msg, s, err = ansi.Decode(msg)
		if s != nil {
			if s.Type == "" {
				b, err = world.ToUTF8(charset, []byte(s.Code))
				if err != nil {
					errhandler(err)
					continue
				}
				w.Text = string(b)
				line.Append(w)
				w = w.Inherit()
			} else if s.Type == "CSI" {
				for _, v := range s.Params {
					switch v {
					case "0":
						{
							w.Color = ""
							w.Background = ""
							w.Bold = false
							w.Inverse = false
							w.Underlined = false
							w.Blinking = false
						}
					case "1":
						{
							w.Bold = true
						}
					case "2":
						{
							w.Bold = false
						}
					case "4":
						{
							w.Underlined = true
						}
					case "5":
						{
							w.Blinking = true
						}
					case "7":
						{
							w.Inverse = true
						}
					case "24":
						{
							w.Underlined = false
						}
					case "25":
						{
							w.Blinking = false
						}
					case "27":
						{
							w.Inverse = false
						}
					case "30":
						{
							w.Color = "Black"
						}
					case "31":
						{
							w.Color = "Red"
						}
					case "32":
						{
							w.Color = "Green"
						}
					case "33":
						{
							w.Color = "Yellow"
						}
					case "34":
						{
							w.Color = "Blue"
						}
					case "35":
						{
							w.Color = "Magenta"
						}
					case "36":
						{
							w.Color = "Cyan"
						}
					case "37":
						{
							w.Color = "White"
						}
					case "40":
						{
							w.Background = "Black"
						}
					case "41":
						{
							w.Background = "Red"
						}
					case "42":
						{
							w.Background = "Green"
						}
					case "43":
						{
							w.Background = "Yellow"
						}
					case "44":
						{
							w.Background = "Blue"
						}
					case "45":
						{
							w.Background = "Magenta"
						}
					case "46":
						{
							w.Background = "Cyan"
						}
					case "47":
						{
							w.Background = "White"
						}
					case "90":
						{
							w.Color = "Bright-Black"
						}
					case "91":
						{
							w.Color = "Bright-Red"
						}
					case "92":
						{
							w.Color = "Bright-Green"
						}
					case "93":
						{
							w.Color = "Bright-Yellow"
						}
					case "94":
						{
							w.Color = "Bright-Blue"
						}
					case "95":
						{
							w.Color = "Bright-Magenta"
						}
					case "96":
						{
							w.Color = "Bright-Cyan"
						}
					case "97":
						{
							w.Color = "Bright-White"
						}
					case "100":
						{
							w.Background = "Bright-Black"
						}
					case "101":
						{
							w.Background = "Bright-Red"
						}
					case "102":
						{
							w.Background = "Bright-Green"
						}
					case "103":
						{
							w.Background = "Bright-Yellow"
						}
					case "104":
						{
							w.Background = "Bright-Blue"
						}
					case "105":
						{
							w.Background = "Bright-Magenta"
						}
					case "106":
						{
							w.Background = "Bright-Cyan"
						}
					case "107":
						{
							w.Background = "Bright-White"
						}
					case "256":
						line = world.NewLine()
					}

				}
			} else {
				// fmt.Println(s, s.Code)
			}
		}
		if err != nil {
			if !errhandler(err) {
				return nil, w
			}
		}
	}
	return line, w
}
