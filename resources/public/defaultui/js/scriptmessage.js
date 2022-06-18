define(function (require) {
    var app=require("/public/defaultui/js/app.js")
    var vm=require("/public/defaultui/js/vm.js")
    var Vue=require("vue")
    var _=require("lodash")
    var handlers=app.onScriptMessage;
    var send=app.send;
    handlers["userinput.popup"]=function(data){
        vm.$notify({
            title: data.Data.Title,
            message: data.Data.Intro,
            duration: 5000,
            type:data.Data.Type,
            onClick:function(){
                vm.callback(data,-1,"ok")
            }
          });  
    }
    handlers["userinput.hideall"]=function(data){
        vm.userinputListVisible=false
        vm.userinputDatagridVisible=false
        vm.visualPromptVisible=false
    }
    handlers["userinput.prompt"]=function(data){
        var msgbody=data.Data
        if (msgbody==undefined){
            msgbody={}
        }
        vm.$prompt(msgbody.Intro,msgbody.Title, {
            confirmButtonText: '提交',
            cancelButtonText: '取消',
            inputValue:msgbody.Value,
            onClose:function(){
                vm.callback(data,-1,"")
            }
          }).then(({ value }) => {
            vm.callback(data,0,value)
          }).catch(() => {
          });
    }
    handlers["userinput.alert"]=function(data){
        var msgbody=data.Data
        if (msgbody==undefined){
            msgbody={}
        }
        vm.$alert(msgbody.Intro,msgbody.Title, {
            confirmButtonText: '确定',
            onClose:function(){
                vm.callback(data,-1,"")
            }
          }).then(({ value }) => {
            vm.callback(data,0,value)
          }).catch(() => {
          });
    }
    handlers["userinput.confirm"]=function(data){
        var msgbody=data.Data
        if (msgbody==undefined){
            msgbody={}
        }
        vm.$confirm(msgbody.Intro,msgbody.Title, {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            onClose:function(){
                vm.callback(data,-1,"")
            }
          }).then(({ value }) => {
            vm.callback(data,0,value)
          }).catch(() => {
          });
    }

    handlers["userinput.list"]=function(data){
        vm.userinputList=data
        vm.userinputListVisible=true
        vm.userinputListSearch=""
        Vue.nextTick(function(){
            if (data.Data.Mutli){
                data.Data.Values.forEach(function(value){
                    data.Data.Items.forEach(function(item){
                        if (item.Key==value){
                            vm.$refs.userinputListTable.toggleRowSelection(item,true)
                        }
                    })
                })
            }
        })
    }
    handlers["userinput.datagrid"]=function(data){
        vm.userinputDatagrid=data
        vm.userinputDatagridVisible=true
    }
    handlers["userinput.note"]=function(data){
        vm.userinputNote=data
        if (data.Data.Type=="md"){
            data.Data.MD=MD.render(data.Data.Body)
        }
        vm.userinputNoteVisible=true
    }
    handlers["userinput.hidedatagrid"]=function(data){
        vm.userinputDatagrid=null
        vm.userinputDatagridVisible=false
    }
    handlers["userinput.visualprompt"]=function(data){
        vm.visualPrompt=data
        switch (vm.visualPrompt.Data.MediaType){
            case "output":
                vm.visualPrompt.Data.Output=JSON.parse(vm.visualPrompt.Data.Source)
                vm.visualPrompt.Data.Source=""
                break
        }
        vm.visualPromptVisible=true    
    }  
    
})