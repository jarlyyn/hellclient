define(function (require) {
var app=require("/public/defaultui/js/app.js")
var vm=require("/public/defaultui/js/vm.js")
var onButton=app.onButton; 
var send=app.send;


onButton.connect=function(){
    send("connect",vm.current)
}
onButton.unlock=function(){
    localStorage.setItem("hellclient-advancemode","on")
    vm.advancemode=true;
}
onButton.lock=function(){
    localStorage.setItem("hellclient-advancemode","")
    vm.advancemode=false;
}
onButton.disconnect=function(){
    send("disconnect",vm.current)
}
onButton.close=function(){
    send("close",vm.current)
}
onButton.notopened=function(){
    vm.notopened=null;
    vm.notopenedVisible=true
    send("notopened")
}
onButton.save=function(){
    send("save",vm.current)
}
onButton.createGame=function(){
    vm.createFail=[];
    vm.gameCreateForm={};
    vm.gameCreateFormVisible=true;
}
onButton.createGameSubmit=function(){
    send("createGame",vm.gameCreateForm)
}
onButton.triggers=function(){
    vm.triggers=null;
    vm.triggersVisible=true
    send("triggers",current)
}
onButton.createTrigger=function(){
    vm.triggerSaveFormVisible=true
    vm.triggerName="";
    vm.triggerSaveForm={};
}
onButton.saveTriggerSubmit=function(){
    vm.triggerSaveForm.Name=vm.triggerName
    vm.triggerSaveForm.Priority=vm.triggerSaveForm.Priority*1;
    send("saveTrigger",vm.triggerSaveForm);
}
onButton.allLines=function(){
    vm.alllines=[]
    vm.allLinesVisible=true;
    send("allLines");
}
onButton.open=function(id){
    send("open",id);
    vm.notopenedVisible=false;
}
})