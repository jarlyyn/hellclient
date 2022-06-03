define(function (require) {
    var settings = require("./canvassettings.js")
    var Canvas = document.getElementById("output")
    Canvas.width = settings.linewidth
    Canvas.lineheight = settings.lineheight
    Canvas.height = settings.maxlines * settings.lineheight
    var PromptCanvas = document.getElementById("prompt-output")
    PromptCanvas.width = settings.linewidth
    PromptCanvas.lineheight = settings.lineheight
    PromptCanvas.height = settings.lineheight

    var Lines = []
    var createLine = function (id, index) {
        var c = document.createElement("canvas")
        c.width = settings.linewidth
        c.height = settings.lineheight
        c.lineheight = settings.lineheight + "px"
        var ctx = c.getContext('2d')
        ctx.textBaseline = "middle"
        ctx.font = settings.font
        ctx.fillStyle = settings.background
        ctx.fillRect(0, 0, settings.linewidth, settings.lineheight)
        return {
            ID: id,
            Canvas: c,
            Position: 0,
            Index: index,
        }
    }
    var RenderLine = function (line, withouticon, nocr) {
        result = []
        var index = 0
        var l = createLine(line.ID, index)
        var color = settings.color
        var icon = ""
        var iconcolor = ""
        if (!withouticon) {
            switch (line.Type) {
                case 0:
                    color = settings.printcolor
                    icon = settings.printicon
                    iconcolor = settings.printiconcolor
                    break;
                case 1:
                    color = settings.systemcolor
                    icon = settings.systemicon
                    iconcolor = settings.systemiconcolor
                    break;
                case 3:
                    color = settings.echocolor
                    icon = settings.echoicon
                    iconcolor = settings.echoiconcolor
                    break;
                case 5:
                    color=settings.bccolor
                    icon=settings.localbcouticon
                    iconcolor=settings.bccolor
                    break;
                case 6:
                    color=settings.bccolor
                    icon=settings.globalbcouticon
                    iconcolor=settings.bccolor
                    break;
                case 7:
                    color=settings.bccolor
                    icon=settings.localbcinicon
                    iconcolor=settings.bccolor
                    break;
                case 8:
                    color=settings.bccolor
                    icon=settings.globalbcinicon
                    iconcolor=settings.bccolor
                    break;
                case 9:
                    color=settings.bccolor
                    icon=settings.requesticon
                    iconcolor=settings.bccolor
                    break;
                case 10:
                    color=settings.bccolor
                    icon=settings.responseicon
                    iconcolor=settings.bccolor
                    break;
                default:
                    break;
            }
        }
        if (icon) {
            var ctx = l.Canvas.getContext('2d')
            var width = ctx.measureText(icon).width
            ctx.fillStyle = iconcolor
            ctx.fillText(icon, l.Position, settings.middleline)
            l.Position += width
        }
        line.Words.forEach(function (word) {
            let texts=[...word.Text]
            for (let i = 0; i < texts.length; i++) {
                var char = texts[i]
                if (char == "\n" && nocr) {
                    continue
                }
                var ctx = l.Canvas.getContext('2d')
                var width = ctx.measureText(char).width
                if (char == "\n" || l.Position + width >= settings.linewidth) {
                    result.push(l)
                    index++
                    l = createLine(line.ID, index)
                    ctx = l.Canvas.getContext('2d')
                }
                let fontcolor
                let bgcolor=word.Background?settings[word.Background]:settings.background
                if (word.Color){
                    fontcolor=(word.Bold && !word.Inverse)?settings["Bold"+word.Color]:settings[word.Color]    
                }else{
                    fontcolor=color
                }
                
                if (word.Inverse){
                    let c=fontcolor
                    fontcolor=bgcolor
                    bgcolor=c
                }
                if (bgcolor) {
                    ctx.fillStyle = bgcolor
                    ctx.fillRect(l.Position, 0, width, settings.lineheight)
                }
                ctx.fillStyle = fontcolor
                ctx.font = word.Bold ? settings.fontbold : settings.font
                if (word.Blinking){
                    ctx.font=settings.fontblinking
                }
                ctx.fillText(char, l.Position, settings.middleline)
                l.Position += width
                if (l.Position >= settings.linewidth && nocr) {
                    break
                }
            }
        })
        result.push(l)
        return result
    }
    var Drawline = function (line) {
        var result = RenderLine(line,false,false)
        result.forEach(function (line) {
            Lines.push(line)
        })
        Render()
    }
    var DrawPrompt = function (line) {
        var ctx = document.getElementById("prompt-output").getContext('2d');
        if (line) {
            var promptline = RenderLine(line, true, true)[0]
            ctx.drawImage(promptline.Canvas, 0, 0, settings.linewidth, settings.lineheight)
        } else {
            ctx.fillStyle = settings.background
            ctx.fillRect(0, 0, settings.linewidth, settings.lineheight)
        }
    }
    var Render = function () {
        Lines.sort(function (a, b) {
            if (a.ID != b.ID) {
                return a.ID > b.ID ? 1 : -1;
            }
            return a.Index > b.Index ? 1 : -1;

        });
        Lines = _.clone(Lines.slice(-settings.maxlines))
        Lines.forEach(function (line, index) {
            var top = (settings.maxlines - Lines.length + index) * settings.lineheight
            var ctx = document.getElementById("output").getContext('2d');
            ctx.drawImage(line.Canvas, 0, top, settings.linewidth, settings.lineheight)
        })
    }
    var Clean = function () {
        Lines = []
    }
    return {
        Drawline: Drawline,
        Render: Render,
        Clean: Clean,
        DrawPrompt: DrawPrompt,
    }
})