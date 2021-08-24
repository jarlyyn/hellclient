define(function (require) {
    var app=require("/public/defaultui/js/app.js")
    var vm=require("/public/defaultui/js/vm.js")
    var Vue=require("vue")
    var _=require("lodash")
    var handlers=app.onScriptMessage;
    var send=app.send;
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
    
})