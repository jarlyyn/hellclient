package converter

import (
	"fmt"
	"modules/world"

	"github.com/jarlyyn/ansi"
)

func ConvertToLine(msg []byte, charset string, errhandler func(err error)) *world.Line {
	w := world.Word{}
	line := world.NewLine()
	if len(msg) == 0 {
		return line
	}
	for {
		out, s, err := ansi.Decode(msg)
		msg = out
		if s != nil {
			if s.Type == "" {
				b, err := ToUTF8(charset, []byte(s.Code))
				if err != nil {
					errhandler(err)
					continue
				}
				w.Text = string(b)
				line.Append(w)
			} else if s.Type == "CSI" {
				for _, v := range s.Params {
					switch v {
					case "0":
						{
							w.Color = ""
							w.Background = "BG-"
							w.Bold = false
						}
					case "1":
						{
							w.Bold = true
						}
					case "2":
						{
							w.Bold = false
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
							w.Background = "BG-Black"
						}
					case "41":
						{
							w.Background = "BG-Red"
						}
					case "42":
						{
							w.Background = "BG-Green"
						}
					case "43":
						{
							w.Background = "BG-Yellow"
						}
					case "44":
						{
							w.Background = "BG-Blue"
						}
					case "45":
						{
							w.Background = "BG-Magenta"
						}
					case "46":
						{
							w.Background = "BG-Cyan"
						}
					case "47":
						{
							w.Background = "BG-White"
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
							w.Background = "BG-Bright-Black"
						}
					case "101":
						{
							w.Background = "BG-Bright-Red"
						}
					case "102":
						{
							w.Background = "BG-Bright-Green"
						}
					case "103":
						{
							w.Background = "BG-Bright-Yellow"
						}
					case "104":
						{
							w.Background = "BG-Bright-Blue"
						}
					case "105":
						{
							w.Background = "BG-Bright-Magenta"
						}
					case "106":
						{
							w.Background = "BG-Bright-Cyan"
						}
					case "107":
						{
							w.Background = "BG-Bright-White"
						}
					case "256":
						line = world.NewLine()
					}

				}
			} else {
				fmt.Println(s, s.Code)
			}
		}
		if err != nil {
			errhandler(err)
			// return
		}
		if msg == nil {
			break
		}
		if (len(msg)) == 0 {
			break
		}

	}
	return line
}
