define(function (require) {
    var app=require("/public/defaultui/js/app.js")
    var vm=require("/public/defaultui/js/vm.js")
    var handlers=app.handlers;
    var send=app.send;
   

handlers.current=function(data){
    vm.current=data
    vm.currenttab=data
}
handlers.line=function(data){
    var lines=vm.lines
    lines.push(data)
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
    vm.clients=data
    data.forEach(function(client) {
        vm.running[client.ID]=client.Running
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
    vm.running[data]=true
}
handlers.disconnected=function(data){
    vm.running[data]=false
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

})