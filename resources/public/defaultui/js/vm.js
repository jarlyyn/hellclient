define(["vue","/public/defaultui/js/app.js"],function (Vue,app) {
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
        }
    }
})
    return vm
})