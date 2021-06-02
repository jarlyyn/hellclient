define(["vue","/public/defaultui/js/app.js"],function (Vue,app) {
    var onButton=app.onButton; 
    var send=app.send;
var data={
    current:"",
    currenttab:"",
    lines:[],
    prompt:{},
    createFail:[],
    cmd:"",
    running:{},
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
}
var vm = new Vue({
    el:"#app",
    data: data,
    mounted:function(){
        document.getElementsByTagName("body")[0].style.visibility="visible";
    },
    methods:{
        send:function(){
             send("send",this.cmd)
             document.getElementById("user-input").getElementsByTagName("input")[0].select()
        },
        onChange:function(current){
            if (vm.clients.length){
            send("change",current.name)
            }
            return false
        },
        onButton:function(data){
            onButton[data]()
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