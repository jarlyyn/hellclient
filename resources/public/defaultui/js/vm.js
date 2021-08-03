define(["vue","/public/defaultui/js/app.js","lodash"],function (Vue,app,_) {
    var onButton=app.onButton; 
var data={
    typeclass:{
        0:" print",
        1:" system",
        3:" echo",
        4:" prompt"
    },
    current:"",
    currenttab:"",
    lines:[],
    status:"",
    history:[],
    prompt:{},
    createFail:[],
    cmd:"",
    info:{},
    clients:[],
    gameCreateFormVisible:false,
    gameCreateForm:{},
    allLines:[],
    allLinesVisible:false,
    notopened:[],
    notopenedVisible:false,
    advancemode:false,
    script:null,
    scriptVisible:false,
    scriptCreateFormVisible:false,
    createScriptFail:[],
    scriptCreateForm:{},
    scriptlistVisible:false,
    scriptlist:null,
    byuser:false,
    timerlist:null,
    usertimerlist:null,
    scripttimerlist:null,
    timersVisible:false,
    timerCreateForm:{},
    timerCreateFormVisible:false,
    updatingTimer:null,
    timerUpdateFormVisible:false,
    aliaslist:null,
    useraliaslist:null,
    scriptaliaslist:null,
    aliasesVisible:false,
    aliasCreateForm:{},
    aliasCreateFormVisible:false,
    updatingAlias:null,
    aliasUpdateFormVisible:false,
    triggerlist:null,
    usertriggerlist:null,
    scripttriggerlist:null,
    triggersVisible:false,
    triggerCreateForm:{},
    triggerCreateFormVisible:false,
    updatingTrigger:null,
    triggerUpdateFormVisible:false,
    variablesVisible:false,
    paramsinfo:null,
    showRequiredParams:false,
    allgameVisible:false,
    sendto:{
        0:"游戏",
        1:"命令",
        2:"输出",
        3:"状态栏",
        4:"记事本",
        5:"追加到记事本",
        6:"日志",
        7:"替换记事本",
        8:"命令队列",
        9:"变量",
        10:"执行",
        11:"快速行走",
        12:"脚本",
        13:"立刻发送",
        14:"脚本(过滤后)",
    },
}
var vm = new Vue({
    el:"#app",
    data: data,
    mounted:function(){
        this.advancemode=(localStorage.getItem("hellclient-advancemode")!="");
        document.getElementsByTagName("body")[0].style.visibility="visible";
    },
    methods:{
        send:function(){
            app.send("send",this.cmd)
             document.getElementById("user-input").getElementsByTagName("input")[0].select()
        },
        onChange:function(current){
            if (vm.clients.length){
                app.send("change",current.name)
            }
            return false
        },
        onGamelistClick:function(row, column, event){
            if (vm.clients.length){
                app.send("change",row.ID)
            }
            this.allgameVisible=false
        },
        onButton:function(data){
            onButton[data]()
        },
        onOpen:function(id){
            onButton.open(id)
        },
        onUseScript:function(script){
            this.scriptlistVisible=false
            app.send("usescript",[this.current,script])
        },
        onUpdateTrigger:function(data){
            vm.saveTriggerFail=[];
            vm.triggerName=data.Name;
            vm.triggerSaveForm=data;
            vm.triggerSaveFormVisible=true;
            vm.triggersVisible=true;
        },
        onHistory:function(command){
            app.send("send",command)
        },
        onDeleteTimer:function (id){
            vm.$confirm('是否要删除该计时器?', '删除', {
                confirmButtonText: '删除',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                app.send("deleteTimer",[vm.current,id])
              }).catch(()=>{
              })
        },
        onUpdateTimer:function (id){
            app.send("loadTimer",[vm.current,id])
            vm.updatingTimer={
                ID:id,
                Form:{},
            }
            vm.timerUpdateFormVisible=true
        },
        onDeleteAlias:function (id){
            vm.$confirm('是否要删除该别名?', '删除', {
                confirmButtonText: '删除',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                app.send("deleteAlias",[vm.current,id])
              }).catch(()=>{
              })
        },
        onUpdateAlias:function (id){
            app.send("loadAlias",[vm.current,id])
            vm.updatingAlias={
                ID:id,
                Form:{},
            }
            vm.aliasUpdateFormVisible=true
        },
        onDeleteTrigger:function (id){
            vm.$confirm('是否要删除该触发器?', '删除', {
                confirmButtonText: '删除',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                app.send("deleteTrigger",[vm.current,id])
              }).catch(()=>{
              })
        },
        onUpdateTrigger:function (id){
            app.send("loadTrigger",[vm.current,id])
            vm.updatingTrigger={
                ID:id,
                Form:{},
            }
            vm.triggerUpdateFormVisible=true
        },
        RenderLines:function(){
                var lines=app.linesbuffer
                lines.sort(function(a, b) {
                    return a.ID>b.ID?1:-1;
                });      
                if (lines.length>60){
                    lines=_.clone(lines.slice(-60))
                }
                app.linesbuffer=lines
                vm.lines=_.clone(app.linesbuffer)
                setTimeout(function(){
                body.scrollTo(body.offsetLeft,body.offsetHeight)
                },0)        
        },
        onDeleteVariable:function (name){
            vm.$confirm('是否要删除该变量?', '删除', {
                confirmButtonText: '删除',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                app.send("deleteParam",[vm.current,name])
              }).catch(()=>{
              })
        },
        onUpdateVariable:function(name,current){
            vm.$prompt('请输入变量值', '编辑变量'+name, {
                confirmButtonText: '编辑',
                cancelButtonText: '取消',
                inputValue:current,
              }).then(({ value }) => {
                  app.send("updateParam",[vm.current ,name,value])
              }).catch(() => {
              });
        
        },
        onUpdateRequiredParam:function(row){
            vm.$prompt(row.Intro, '设置变量'+row.Name+"["+row.Desc+"]", {
                confirmButtonText: '设置',
                cancelButtonText: '取消',
                customClass:"update-required",
                inputValue:vm.paramsinfo.Params[row.Name],
              }).then(({ value }) => {
                  app.send("updateParam",[vm.current ,row.Name,value])
              }).catch(() => {
              });
        },
        onUpdateParamComment:function(row){
            vm.$prompt("备注", '备注变量'+row.Name+"["+row.Desc+"]", {
                confirmButtonText: '备注',
                cancelButtonText: '取消',
                customClass:"update-comment",
                inputType:"textarea",
                inputValue:vm.paramsinfo.ParamComments[row.Name],
              }).then(({ value }) => {
                  app.send("updateParamComment",[vm.current ,row.Name,value])
              }).catch(() => {
              });
        },

        onDrop:function(){
            vm.allLinesVisible=false
            vm.cmd=vm.cmd+document.getSelection().toString()
        },
        gamelistRowClassName:function(data){
            return data.row.Running?"game-list-running":"game-list-not-running"
        },
    }
})
    return vm
})