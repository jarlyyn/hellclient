define(function (require) {
    var app = require("/public/defaultui/js/app.js")
    var vm = require("/public/defaultui/js/vm.js")
    var onButton = app.onButton;
    var send = app.send;


    onButton.connect = function () {
        send("connect", vm.current)
    }
    onButton.unlock = function () {
        vm.$confirm('是否开启脚本编辑模式?在脚本编辑模式中可以对脚本的触发器，计时器和别名进行编辑', '提示', {
            confirmButtonText: '开启',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            localStorage.setItem("hellclient-advancemode", "on")
            vm.advancemode = true;
        }).catch(() => {

        })

    }
    onButton.lock = function () {
        localStorage.setItem("hellclient-advancemode", "")
        vm.advancemode = false;
    }
    onButton.disconnect = function () {
        send("disconnect", vm.current)
    }
    onButton.close = function () {
        vm.$confirm('是否要关闭本游戏?', '提示', {
            confirmButtonText: '关闭',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            send("close", vm.current)
        }).catch(() => {

        })
    }
    onButton.notopened = function () {
        vm.notopened = null;
        vm.notopenedVisible = true
        send("notopened")
    }
    onButton.script = function () {
        vm.script = null;
        vm.scriptVisible = true
        send("scriptinfo", vm.current)
    }

    onButton.save = function () {
        vm.$confirm('原游戏将被覆盖，是否要保存游戏?', '提示', {
            confirmButtonText: '覆盖',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            send("save", vm.current)
        }).catch(() => {

        })
    }
    onButton.savescript = function () {
        vm.$confirm('原脚本将被覆盖，是否要保存脚本?', '提示', {
            confirmButtonText: '覆盖',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            send("savescript", vm.current)
        }).catch(() => {

        })
    }

    onButton.reload = function () {
        vm.$confirm('脚本所有的修改将丢失，进行中的程序也将停止，是否要重新加载脚本?', '提示', {
            confirmButtonText: '重新加载',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            send("reloadScript", vm.current)
        }).catch(() => {

        })
    }
    onButton.clientquick = function () {
        let clients = [...vm.clients]
        if (clients.length) {
            clients.sort(function (a, b) {
                if (a.Priority != b.Priority) {
                    if (a.Priority > b.Priority) {
                        return -1
                    } else {
                        return 1
                    }
                }
                if (a.LastActive != b.LastActive) {
                    if (a.LastActive > b.LastActive) {
                        return 1
                    } else {
                        return -1
                    }
                }
                return 0
            })
            app.send("change", clients[0].ID)
        }
    }


    onButton.createGame = function () {
        send("defaultServer", "")
    }
    onButton.createGameSubmit = function () {
        send("createGame", vm.gameCreateForm)
    }

    onButton.allLines = function () {
        vm.allLines = null
        vm.allLinesVisible = true;
        send("allLines");
    }
    onButton.open = function (id) {
        send("open", id);
        vm.notopenedVisible = false;
    }
    onButton.createScript = function () {
        vm.createScriptFail = [];
        vm.scriptCreateForm = {};
        vm.scriptCreateFormVisible = true;
    }
    onButton.createScriptSubmit = function () {
        send("createScript", vm.scriptCreateForm);
    }
    onButton.listScriptinfo = function () {
        vm.scriptlist = null;
        vm.scriptlistVisible = true
        send("listScriptinfo")
    }
    onButton.cleanScript = function () {
        app.send("usescript", [vm.current, ""])
    }
    onButton.usertimers = function () {
        vm.usertimerlist = null;
        vm.timersVisible = true;
        vm.byuser = true;
        app.send("timers", [vm.current, "byuser"])
    }
    onButton.scripttimers = function () {
        vm.scripttimerlist = null;
        vm.timersVisible = true;
        vm.byuser = false;
        app.send("timers", [vm.current, ""])
    }
    onButton.createTimer = function () {
        vm.createFail = [];
        vm.timerCreateForm = {};
        vm.timerCreateFormVisible = true;
    }
    onButton.createTimerSubmit = function () {
        vm.timerCreateForm.World = vm.current
        vm.timerCreateForm.ByUser = vm.byuser
        vm.timerCreateForm.SendTo = vm.timerCreateForm.SendTo * 1
        send("createTimer", vm.timerCreateForm);
    }
    onButton.updateTimerSubmit = function () {
        vm.updatingTimer.Form.World = vm.current
        vm.updatingTimer.Form.ID = vm.updatingTimer.ID
        vm.updatingTimer.Form.SendTo = vm.updatingTimer.Form.SendTo * 1
        send("updateTimer", vm.updatingTimer.Form);
    }
    onButton.useraliases = function () {
        vm.useraliaslist = null;
        vm.aliasesVisible = true;
        vm.byuser = true;
        app.send("aliases", [vm.current, "byuser"])
    }
    onButton.scriptaliases = function () {
        vm.scriptaliaslist = null;
        vm.aliasesVisible = true;
        vm.byuser = false;
        app.send("aliases", [vm.current, ""])
    }
    onButton.about = function () {
        app.send("about", "")
    }
    onButton.createAlias = function () {
        vm.createFail = [];
        vm.aliasCreateForm = {
            SendTo: "0",
            Sequence: 100,
        };
        vm.aliasCreateFormVisible = true;
    }
    onButton.createAliasSubmit = function () {
        vm.aliasCreateForm.World = vm.current
        vm.aliasCreateForm.ByUser = vm.byuser
        vm.aliasCreateForm.SendTo = vm.aliasCreateForm.SendTo * 1
        send("createAlias", vm.aliasCreateForm);
    }
    onButton.updateAliasSubmit = function () {
        vm.updatingAlias.Form.World = vm.current
        vm.updatingAlias.Form.ID = vm.updatingAlias.ID
        vm.updatingAlias.Form.SendTo = vm.updatingAlias.Form.SendTo * 1
        send("updateAlias", vm.updatingAlias.Form);
    }
    onButton.usertriggers = function () {
        vm.usertriggerlist = null;
        vm.triggersVisible = true;
        vm.byuser = true;
        app.send("triggers", [vm.current, "byuser"])
    }
    onButton.scripttriggers = function () {
        vm.scripttriggerlist = null;
        vm.triggersVisible = true;
        vm.byuser = false;
        app.send("triggers", [vm.current, ""])
    }
    onButton.createTrigger = function () {
        vm.createFail = [];
        vm.triggerCreateForm = {
            SendTo: "0",
            Sequence: 100,
        };
        vm.triggerCreateFormVisible = true;
    }
    onButton.createTriggerSubmit = function () {
        vm.triggerCreateForm.World = vm.current
        vm.triggerCreateForm.ByUser = vm.byuser
        vm.triggerCreateForm.SendTo = vm.triggerCreateForm.SendTo * 1
        send("createTrigger", vm.triggerCreateForm);
    }
    onButton.updateTriggerSubmit = function () {
        vm.updatingTrigger.Form.World = vm.current
        vm.updatingTrigger.Form.ID = vm.updatingTrigger.ID
        vm.updatingTrigger.Form.SendTo = vm.updatingTrigger.Form.SendTo * 1
        send("updateTrigger", vm.updatingTrigger.Form);
    }
    onButton.variable = function () {
        vm.paramsinfo = null;
        vm.variablesVisible = true;
        vm.showRequiredParams = true;
        app.send("params", vm.current)
    }
    onButton.createVariable = function () {
        vm.$prompt('请输入变量名', '创建变量', {
            confirmButtonText: '添加',
            cancelButtonText: '取消',
            inputPattern: /.+/,
            inputErrorMessage: '格式不正确'
        }).then(({ value }) => {
            app.send("updateParam", [vm.current, value, ""])
        }).catch(() => {
        });

    }
    onButton.overview = function () {
        send("change", "")
    }
    onButton.worldSettings = function () {
        vm.worldSettings = null;
        app.send("worldSettings", vm.current)
    }
    onButton.updateWorldSettings = function () {
        vm.createFail = []
        vm.worldsettingsUpdateForm = {
            Charset: vm.worldSettings.Charset,
            CommandStackCharacter: vm.worldSettings.CommandStackCharacter,
            Host: vm.worldSettings.Host,
            Name: vm.worldSettings.Name,
            Port: vm.worldSettings.Port,
            Proxy: vm.worldSettings.Proxy,
            ID: vm.current,
            ScriptPrefix: vm.worldSettings.ScriptPrefix,
            ShowBroadcast: vm.worldSettings.ShowBroadcast,
            ShowSubneg: vm.worldSettings.ShowSubneg,
            ModEnabled: vm.worldSettings.ModEnabled
        }
        vm.worldsettingsUpdateFormVisible = true
    }
    onButton.worldsettingsUpdateSubmit = function () {
        vm.worldsettingsUpdateForm.ID = vm.current
        send("updateWorldSettings", vm.worldsettingsUpdateForm);
    }
    onButton.updateScriptSettings = function () {
        vm.createFail = []
        vm.scriptsettingsUpdateForm = {
            Channel: vm.scriptSettings.Channel,
            Desc: vm.scriptSettings.Desc,
            Intro: vm.scriptSettings.Intro,
            Name: vm.scriptSettings.Name,
            OnAssist: vm.scriptSettings.OnAssist,
            OnKeyUp: vm.scriptSettings.OnKeyUp,
            OnBroadcast: vm.scriptSettings.OnBroadcast,
            OnClose: vm.scriptSettings.OnClose,
            OnConnect: vm.scriptSettings.OnConnect,
            OnDisconnect: vm.scriptSettings.OnDisconnect,
            OnResponse: vm.scriptSettings.OnResponse,
            OnOpen: vm.scriptSettings.OnOpen,
            Type: vm.scriptSettings.Type,
            OnHUDClick: vm.scriptSettings.OnHUDClick,
            OnBuffer: vm.scriptSettings.OnBuffer,
            OnBufferMin: vm.scriptSettings.OnBufferMin,
            OnBufferMax: vm.scriptSettings.OnBufferMax,
            OnSubneg: vm.scriptSettings.OnSubneg,
            OnFocus:vm.scriptSettings.OnFocus,
            OnLoseFocus:vm.scriptSettings.OnLoseFocus,
        }
        vm.scriptsettingsUpdateFormVisible = true
    }
    onButton.scriptSettings = function () {
        vm.scriptSettings = null;
        app.send("scriptSettings", vm.current)
    }
    onButton.scriptsettingsUpdateSubmit = function () {
        vm.scriptsettingsUpdateForm.ID = vm.current
        send("updateScriptSettings", vm.scriptsettingsUpdateForm);
    }
    onButton.batchcommand=function(){
        send("batchcommandscripts");
    }
    onButton.batchcommandsend=function(){
        var result=[]
        for (var i=0;i<vm.BatchCommandForm.Scripts.length;i++){
            if (vm.BatchCommandForm.Scripts[i].value){
            result.push(vm.BatchCommandForm.Scripts[i].key)
            }
        }
        send('batchcommand',{Scripts:result,Command:vm.BatchCommandForm.Command});
        vm.BatchCommandFormVisible=false;
    }
    onButton.requiredParams = function () {
        vm.requiredParams = null;
        app.send("requiredParams", vm.current)
    }
    onButton.createRequireParam = function () {
        vm.createFail = []
        vm.requiredParamCreateForm = {
            Name: "",
            Desc: "",
            Intro: "",
        };
        vm.requiredParamCreateFormVisible = true
    }
    onButton.requiredParamCreateSubmit = function () {
        vm.createFail = []
        if (vm.requiredParamCreateForm.Name) {
            vm.requiredParamCreateForm.Name = vm.requiredParamCreateForm.Name.trim()
        }
        if (!vm.requiredParamCreateForm.Name) {
            vm.createFail.push({ Field: "Name", Label: "变量名", Msg: "变量名必填" })
            return
        }
        for (var key in vm.requiredParams) {
            if (vm.requiredParams[key].Name == vm.requiredParamCreateForm.Name) {
                vm.createFail.push({ Field: "Name", Label: "变量名", Msg: "变量名重复" })
                return
            }
        }
        vm.requiredParams.push(vm.requiredParamCreateForm)
        vm.requiredParamCreateFormVisible = false
        vm.updateRequiredParams()
    }
    onButton.requiredParamUpdateSubmit = function () {
        for (var key in vm.requiredParams) {
            if (vm.requiredParams[key].Name == vm.requiredParamUpdateForm.Name) {
                vm.requiredParams[key].Desc = vm.requiredParamUpdateForm.Desc
                vm.requiredParams[key].Intro = vm.requiredParamUpdateForm.Intro
                vm.requiredParamUpdateFormVisible = false
                vm.updateRequiredParams()
                return
            }
        }
        vm.requiredParamUpdateFormVisible = false
    }
    onButton.userinputsubmit = function () {
        vm.callback(vm.userinputList, 0, JSON.stringify(vm.userinputList.Data.Values || []))
        vm.userinputList = null
        vm.userinputListVisible = false
        vm.userinputListSearch = ""

    }
    onButton.requestpermissions = function () {
        data = vm.RequestPermissions
        vm.RequestPermissions = null
        vm.RequestPermissionsVisible = false
        app.send("requestPermissions", data)
    }
    onButton.requesttrustdomains = function () {
        data = vm.RequestTrustDomains
        vm.RequestTrustDomains = null
        vm.RequestTrustDomainsVisible = false
        app.send("requestTrustDomains", data)
    }
    onButton.authorized = function () {
        app.send("authorized", vm.current)
    }
    onButton.revokeauthorized = function () {
        vm.$confirm('是否要注销所有权限?', '注销权限', {
            confirmButtonText: '注销',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(() => {
            app.send("revokeAuthorized", vm.current)
        }).catch(() => {
        })
    }
    onButton.visualPromptRefresh = function () {
        if (vm.visualPrompt.Data.RefreshCallback) {
            var data = {
                ID: vm.visualPrompt.ID,
                Name: vm.visualPrompt.Name,
                Script: vm.visualPrompt.Data.RefreshCallback,
            }
            vm.callback(data, 0, "")
        }
    }
    onButton.masssend = function () {
        vm.domasssend()
    }
    onButton.visualPromptSubmit = function () {
        if (vm.visualPrompt) {
            var data = {
                ID: vm.visualPrompt.ID,
                Name: vm.visualPrompt.Name,
                Script: vm.visualPrompt.Script,
            }
            var val = vm.visualPrompt.Data.Value
            vm.visualPrompt = null
            vm.visualPromptVisible = false
            vm.callback(data, 0, val)
        }
    }
    onButton.password = function () {
        vm.usesrpasswordForm = {}
        vm.usesrpasswordFormVisible = true
    }
    onButton.userpasswordSubmit = function () {
        app.send("updatepassword", vm.usesrpasswordForm)
    }
})

