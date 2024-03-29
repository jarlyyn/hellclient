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
    var createLine = function (id, index,bcolor) {
        var c = document.createElement("canvas")
        c.width = settings.linewidth
        c.height = settings.lineheight
        c.lineheight = settings.lineheight + "px"
        var ctx = c.getContext('2d')
        ctx.textBaseline = "middle"
        ctx.font = settings.font
        ctx.fillStyle = bcolor
        ctx.fillRect(0, 0, settings.linewidth, settings.lineheight)
        return {
            ID: id,
            Canvas: c,
            Position: 0,
            Index: index,
        }
    }
    var RenderLine = function (line, withouticon, nocr,bcolor) {
        result = []
        var index = 0
        if (!bcolor){
            bcolor=settings.background
        }
        var l = createLine(line.ID, index,bcolor)
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
                case 11:
                    color=settings.bccolor
                    icon=settings.subnegicon
                    iconcolor=settings.bccolor
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
                let colorname
                let backgroundname
                if (word.Inverse){
                    colorname=word.Background
                    backgroundname=word.Color
                }else{
                    colorname=word.Color
                    backgroundname=word.Background

                }
                let fontcolor
                let bgcolor=backgroundname?settings[backgroundname]:settings.bcolor
                if (colorname){
                    fontcolor=word.Bold?settings["Bold"+colorname]:settings[colorname]    
                }else{
                    fontcolor=color
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
                if (word.Underlined){
                    ctx.fillRect(l.Position, settings.underline, width, settings.underlineheight);
                }
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
    var RenderHUD=function(content){
        var hudwrapper=document.getElementById("hudwrapper")
        let hud=document.getElementById("hud")
        if (content.length==0){
            hud.style.height=0  
            hud.height=0
            hudwrapper.className="hide"
            return 
        }
        hudwrapper.className=""
        hud.style.height=settings.lineheight*content.length/devicePixelRatio+"px"
        hud.height=settings.lineheight*content.length
        hud.width=settings.linewidth
        let top=0
        content.forEach(function(data){
            if (data){
                let result=RenderLine(data,true, true,settings.hudbackground)
                var ctx=hud.getContext('2d');
                ctx.drawImage(result[0].Canvas, 0, top, settings.linewidth, settings.lineheight)
            }
            top+=settings.lineheight
        })
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
        RenderHUD:RenderHUD,
        Drawline: Drawline,
        Render: Render,
        Clean: Clean,
        DrawPrompt: DrawPrompt,
    }
})