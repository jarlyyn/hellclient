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
}
var vm = new Vue({
    el:"#app",
    data: data,
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