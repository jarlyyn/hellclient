define(function (require) {
    var lineheight = 20
    var linewidth = 80 * 14
    var middleline = lineheight / 2
    var maxlines = 60
    var font = "14px monospace"
    var fontbold= "bold 14px monospace"
    var fontblinking= "italic 14px monospace"
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
    var requesticon="☎请求 "
    var responseicon="☎响应 "
    var bccolor="rgb(127,127,127)"
    return {
        lineheight: lineheight,
        linewidth: linewidth,
        middleline: middleline,
        maxlines: maxlines,
        font: font,
        fontbold:fontbold,
        fontblinking:fontblinking,
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
        requesticon:requesticon,
        responseicon:responseicon,
    }
})