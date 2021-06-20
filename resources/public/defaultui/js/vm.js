define(["vue","/public/defaultui/js/app.js"],function (Vue,app) {
    var onButton=app.onButton; 
var data={
    current:"",
    currenttab:"",
    lines:[],
    prompt:{},
    createFail:[],
    cmd:"",
    info:{},
    clients:[],
    gameCreateFormVisible:false,
    gameCreateForm:{},
    triggers:null,
    triggerName:"",
    triggersVisible:false,
    triggerSaveFormVisible:false,
    triggerSaveForm:{},
    saveTriggerFail:[],
    allLines:[],
    allLinesVisible:false,
    notopened:[],
    notopenedVisible:false,
}
var vm = new Vue({
    el:"#app",
    data: data,
    mounted:function(){
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
        onUpdateTrigger:function(data){
            vm.saveTriggerFail=[];
            vm.triggerName=data.Name;
            vm.triggerSaveForm=data;
            vm.triggerSaveFormVisible=true;
            vm.triggersVisible=true;
        }
    }
})
    return vm
})