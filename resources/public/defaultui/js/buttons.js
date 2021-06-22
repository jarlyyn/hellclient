define(function (require) {
var app=require("/public/defaultui/js/app.js")
var vm=require("/public/defaultui/js/vm.js")
var onButton=app.onButton; 
var send=app.send;


onButton.connect=function(){
    send("connect",vm.current)
}
onButton.unlock=function(){
    vm.$confirm('是否开启脚本编辑模式?在脚本编辑模式中可以对脚本的触发器，计时器和别名进行编辑', '提示', {
        confirmButtonText: '开启',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        localStorage.setItem("hellclient-advancemode","on")
        vm.advancemode=true;
      }).catch(()=>{

      })

}
onButton.lock=function(){
    localStorage.setItem("hellclient-advancemode","")
    vm.advancemode=false;
}
onButton.disconnect=function(){
    send("disconnect",vm.current)
}
onButton.close=function(){
    vm.$confirm('是否要关闭本游戏?', '提示', {
        confirmButtonText: '关闭',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        send("close",vm.current)
      }).catch(()=>{

      })
}
onButton.notopened=function(){
    vm.notopened=null;
    vm.notopenedVisible=true
    send("notopened")
}
onButton.script=function(){
    vm.script=null;
    vm.scriptVisible=true
    send("scriptinfo",vm.current)
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
onButton.createScript=function(){
    vm.createScriptFail=[];
    vm.scriptCreateForm={};
    vm.scriptCreateFormVisible=true;
}
onButton.createScriptSubmit=function(){
    send("createScript",vm.scriptCreateForm);
}
onButton.listScriptinfo=function(){
    vm.scriptlist=null;
    vm.scriptlistVisible=true
    send("listScriptinfo")
}
onButton.cleanScript=function(){
    app.send("usescript",[vm.current,""])
}
})