define(function (require) {
    var app=require("/public/defaultui/js/app.js")
    var vm=require("/public/defaultui/js/vm.js")
    var handlers=app.handlers;
    var send=app.send;
   

handlers.current=function(data){
    vm.current=data
    vm.currenttab=data
    vm.lines=[]
}
handlers.line=function(data){
    var lines=vm.lines
    lines.push(data)
    lines.sort(function(a, b) {
        return a.ID>b.ID?1:-1;
    });      
    if (lines.length>500){
        lines.shift()
    }else{
        setTimeout(function(){
        body.scrollTo(body.offsetLeft,body.offsetHeight)
        },0)
    }
    vm.lines=lines
}
handlers.prompt=function(data){
    vm.prompt=data
}
handlers.clients=function(data){
    vm.info=[]
    vm.clients=data
    data.forEach(function(client) {
        vm.info[client.ID]=client
    })
}
handlers.lines=function(data){
    var lines=[]
    data.forEach(function(element) {
        lines.push(element)
    if (lines.length>500){
        lines.shift()
    }
    })
    vm.lines=lines
    setTimeout(function(){
        body.scrollTo(body.offsetLeft,body.offsetHeight)
        },0)
}
handlers.connected=function(data){
    if (vm.info[data]){
    vm.info[data].Running=true
    }
}
handlers.disconnected=function(data){
    if (vm.info[data]){
    vm.info[data].Running=false
    }
}
handlers.createFail=function(data){
    vm.createFail=data
}
handlers.createSuccess=function(data){
    vm.gameCreateFormVisible=false;
    send("change",data)
}
handlers.triggers=function(data){
    vm.triggers=data
}
handlers.triggerFail=function(data){
    vm.saveTriggerFail=data
}
handlers.triggerSuccess=function(data){
    vm.triggerSaveFormVisible=false;
}
handlers.allLines=function(data){
    vm.allLines=data
}
handlers.notopened=function(data){
    vm.notopened=data
}
handlers.scriptinfo=function(data){
    vm.script=data
}
handlers.createScriptFail=function(data){
    vm.createScriptFail=data
}
handlers.createScriptSuccess=function(data){
    vm.scriptCreateFormVisible=false;
}
handlers.scriptinfoList=function(data){
    vm.scriptlist=data;
}
handlers.status=function(data){
    vm.status=data;
}
handlers.history=function(data){
    vm.history=data;
}
handlers.usertimers=function(data){
    vm.usertimerlist=data
}
handlers.scripttimers=function(data){
    vm.scripttimerlist=data
}
handlers.createTimerSuccess=function(data){
    vm.timerCreateFormVisible=false
}
handlers.updateTimerSuccess=function(data){
    vm.timerUpdateFormVisible=false
}
handlers.timer=function(data){
    if (vm.updatingTimer&&vm.updatingTimer.ID==data.ID){
        vm.updatingTimer.Form=data
    }
}

})