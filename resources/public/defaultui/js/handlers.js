define(function (require) {
    var app=require("/public/defaultui/js/app.js")
    var vm=require("/public/defaultui/js/vm.js")
    var _=require("lodash")
    var handlers=app.handlers;
    var send=app.send;
   
var render=_.debounce(vm.RenderLines,50,{
    leading:true,
    maxWait:200,
})
handlers.current=function(data){
    vm.current=data
    vm.currenttab=data
    vm.lines=[]
    document.getElementById("mud-input").focus()
}
handlers.line=function(data){
    var lines=app.linesbuffer
    lines.push(data)
    app.linesbuffer=lines
    render()
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
handlers.switchStatus=function(data){
    vm.switchstatus=data
}
handlers.lines=function(data){
    var lines=[]
    data.forEach(function(element) {
        lines.push(element)
    if (lines.length>50){
        lines.shift()
    }
    })
    app.linesbuffer=lines
    render()
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
    vm.notopenedVisible=false;
    send("change",data)
}
handlers.updateSuccess=function(){
    vm.worldsettingsUpdateFormVisible=false
}
handlers.updateScriptSuccess=function(){
    vm.scriptsettingsUpdateFormVisible=false
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
    var allliens=document.getElementById("alllines-wrapper").parentElement
    setTimeout(function(){
        allliens.scrollTo(0,9999999)
    },0)        

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

handlers.useraliases=function(data){
    vm.useraliaslist=data
}
handlers.scriptaliases=function(data){
    vm.scriptaliaslist=data
}
handlers.createAliasSuccess=function(data){
    vm.aliasCreateFormVisible=false
}
handlers.updateAliasSuccess=function(data){
    vm.aliasUpdateFormVisible=false
}
handlers.alias=function(data){
    if (vm.updatingAlias&&vm.updatingAlias.ID==data.ID){
        vm.updatingAlias.Form=data
    }
}

handlers.usertriggers=function(data){
    vm.usertriggerlist=data
}
handlers.scripttriggers=function(data){
    vm.scripttriggerlist=data
}
handlers.createTriggerSuccess=function(data){
    vm.triggerCreateFormVisible=false
}
handlers.updateTriggerSuccess=function(data){
    vm.triggerUpdateFormVisible=false
}
handlers.trigger=function(data){
    if (vm.updatingTrigger&&vm.updatingTrigger.ID==data.ID){
        vm.updatingTrigger.Form=data
    }
}
handlers.paramsinfo=function(data){
     if (!data.RequiredParams){
         vm.showRequiredParams=false
     }
     data.ParamList=[]
     for (const value of Object.keys(data.Params)) {
        data.ParamList.push({Name:value,Value:data.Params[value]})
      }
      data.ParamList.sort(function(a, b) {
        return a.Name>b.Name?1:-1;
    });      
     vm.paramsinfo=data
}
handlers.scriptMessage=function(data){
    var h=app.onScriptMessage[data.Name]
    if (h){
        h(data)
    }
}
handlers.version=function(data){
    vm.version=data
    vm.aboutVisible=true
}
handlers.worldSettings=function(data){
    vm.worldSettingsVisible=true;
    vm.worldSettings=data
}
handlers.scriptSettings=function(data){
    vm.scriptSettingsVisible=true;
    vm.scriptSettings=data
}
handlers.requiredParams=function(data){
    vm.requiredParamsVisible=true;
    vm.requiredParams=data||[]
}
handlers.defaultServer=function(data){
    var server=data.split(":")
    var host=server[0]
    var port=server[1]
    vm.createFail=[];
    vm.gameCreateForm={
        Host:host,
        Port:port,
    };
    vm.gameCreateFormVisible=true;
}
})