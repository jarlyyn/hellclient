define(function (require) {
    var lineheight = 20
    var linewidth = 80 * 14
    var middleline = lineheight / 2
    var maxlines = 60
    var font = "14px monospace"
    var fontbold= "bold 14px monospace"
    var background = "#000000"
    var color = "#ffffff"
    var Black = "rgb(0,0,0)"
    var Red = "rgb(127,0,0)"
    var Green = "rgb(0,147,0)"
    var Yellow = "rgb(252,127,0)"
    var Blue = "rgb(0,0,127)"
    var Magenta = "rgb(156,0,156)"
    var Cyan = "rgb(0,147,147)"
    var White = "rgb(210,210,210)"
    var BrightBlack = "rgb(127,127,127)"
    var BrightRed = "rgb(255,0,0)"
    var BrightGreen = "rgb(0,252,0)"
    var BrightYellow = "rgb(255,255,0)"
    var BrightBlue = "rgb(0,0,252)"
    var BrightMagenta = "rgb(255,0,255)"
    var BrightCyan = "rgb(0,255,255)"
    var BrightWhite = "rgb(255,255,255)"
    var BGBlack = "rgb(0,0,0)"
    var BGRed = "rgb(127,0,0)"
    var BGGreen = "rgb(0,147,0)"
    var BGYellow = "rgb(252,127,0)"
    var BGBlue = "rgb(0,0,127)"
    var BGMagenta = "rgb(156,0,156)"
    var BGCyan = "rgb(0,147,147)"
    var BGWhite = "rgb(210,210,210)"
    var BGBrightBlack = "rgb(127,127,127)"
    var BGBrightRed = "rgb(255,0,0)"
    var BGBrightGreen = "rgb(0,252,0)"
    var BGBrightYellow = "rgb(255,255,0)"
    var BGBrightBlue = "rgb(0,0,252)"
    var BGBrightMagenta = "rgb(255,0,255)"
    var BGBrightCyan = "rgb(0,255,255)"
    var BGBrightWhite = "rgb(255,255,255)"
    var echocolor="cyan"
    var echoicon="↣"
    var echoiconcolor="teal"
    var systemcolor="red"
    var systemicon="⯳"
    var systemiconcolor="purple"
    var printcolor="mediumspringgreen"
    var printicon="↢"
    var printiconcolor="green"
    var localbcinicon="☎本地广播 "
    var globalbcinicon="☎全局广播 "
    var localbcouticon="☎本地广播出 "
    var globalbcouticon="☎全集广播出 "
    var bccolor="rgb(127,127,127)"
    return {
        lineheight: lineheight,
        linewidth: linewidth,
        middleline: middleline,
        maxlines: maxlines,
        font: font,
        fontbold:fontbold,
        background: background,
        color: color,
        Black: Black,
        Red: Red,
        Green: Green,
        Yellow: Yellow,
        Blue: Blue,
        Magenta: Magenta,
        Cyan: Cyan,
        White: White,
        "Bright-Black": BrightBlack,
        "Bright-Red": BrightRed,
        "Bright-Green": BrightGreen,
        "Bright-Yellow": BrightYellow,
        "Bright-Blue": BrightBlue,
        "Bright-Magenta": BrightMagenta,
        "Bright-Cyan": BrightCyan,
        "Bright-White": BrightWhite,
        "BG-Black": BGBlack,
        "BG-Red": BGRed,
        "BG-Green": BGGreen,
        "BG-Yellow": BGYellow,
        "BG-Blue": BGBlue,
        "BG-Magenta": BGMagenta,
        "BG-Cyan": BGCyan,
        "BG-White": BGWhite,
        "BG-Bright-Black": BGBrightBlack,
        "BG-Bright-Red": BGBrightRed,
        "BG-Bright-Green": BGBrightGreen,
        "BG-Bright-Yellow": BGBrightYellow,
        "BG-Bright-Blue": BGBrightBlue,
        "BG-Bright-Magenta": BGBrightMagenta,
        "BG-Bright-Cyan": BGBrightCyan,
        "BG-Bright-White": BGBrightWhite,
        echocolor:echocolor,
        echoicon:echoicon,
        echoiconcolor:echoiconcolor,
        systemcolor:systemcolor,
        systemicon:systemicon,
        systemiconcolor:systemiconcolor,
        printcolor:printcolor,
        printicon:printicon,
        printiconcolor:printiconcolor,
        bccolor:bccolor,
        localbcinicon:localbcinicon,
        globalbcinicon:globalbcinicon,
        localbcouticon:localbcouticon,
        globalbcouticon:globalbcouticon,
    }
})