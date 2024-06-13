define(["vue", "/public/defaultui/js/app.js", "lodash", "/public/defaultui/js/canvas.js", "public/defaultui/js/vue-2-boring-avatars.umd.js"], function (Vue, app, _, canvas, avatar) {
    var onButton = app.onButton;
    var data = {
        typeclass: {
            0: " print",
            1: " system",
            3: " echo",
            4: " prompt",
            5: " localbcout",
            6: " globalbcout",
            7: " localbcin",
            8: " globalbcin",
            9: " request",
            10: " response",
            11: "subneg"
        },
        avatorcolors: [],
        hud: [],
        current: "",
        currenttab: "",
        status: "",
        history: [],
        suggestion: [],
        createFail: [],
        cmd: "",
        info: {},
        clients: [],
        gameCreateFormVisible: false,
        gameCreateForm: null,
        allLines: [],
        allLinesVisible: false,
        notopened: [],
        notopenedVisible: false,
        advancemode: false,
        script: null,
        scriptVisible: false,
        scriptCreateFormVisible: false,
        createScriptFail: [],
        scriptCreateForm: null,
        scriptlistVisible: false,
        scriptlist: null,
        byuser: false,
        timerlist: null,
        usertimerlist: null,
        scripttimerlist: null,
        timersVisible: false,
        timerCreateForm: null,
        timerCreateFormVisible: false,
        updatingTimer: null,
        timerUpdateFormVisible: false,
        aliaslist: null,
        useraliaslist: null,
        scriptaliaslist: null,
        aliasesVisible: false,
        aliasCreateForm: null,
        aliasCreateFormVisible: false,
        updatingAlias: null,
        aliasUpdateFormVisible: false,
        triggerlist: null,
        usertriggerlist: null,
        scripttriggerlist: null,
        triggersVisible: false,
        triggerCreateForm: null,
        triggerCreateFormVisible: false,
        updatingTrigger: null,
        triggerUpdateFormVisible: false,
        variablesVisible: false,
        paramsinfo: null,
        aboutVisible: false,
        version: "",
        uiversion: "23.07.30",
        showRequiredParams: true,
        allgameVisible: false,
        userinputList: null,
        userinputListVisible: false,
        userinputListSearch: "",
        userinputDatagrid: null,
        userinputDatagridVisible: false,
        userinputNote: false,
        userinputNoteVisible: false,
        switchstatus: 0,
        statusVisible: false,
        worldSettings: null,
        worldSettingsVisible: false,
        scriptSettings: null,
        scriptSettingsVisible: false,
        requiredParams: [],
        requiredParamsVisible: false,
        requiredParamCreateForm: null,
        requiredParamCreateFormVisible: false,
        requiredParamUpdateForm: null,
        requiredParamUpdateFormVisible: false,
        worldsettingsUpdateForm: null,
        worldsettingsUpdateFormVisible: false,
        scriptsettingsUpdateForm: null,
        scriptsettingsUpdateFormVisible: false,
        RequestPermissions: null,
        RequestPermissionsVisible: false,
        RequestTrustDomains: null,
        RequestTrustDomainsVisible: false,
        MassSendForm: null,
        MassSendFormVisible: false,
        Authorized: null,
        AuthorizedVisible: false,
        visualPrompt: null,
        visualPromptVisible: false,
        historypos: -1,
        lastinput:"",
        usesrpasswordForm: null,
        usesrpasswordFormVisible: false,
        sendto: {
            0: "游戏",
            1: "命令",
            2: "输出",
            3: "状态栏",
            4: "记事本",
            5: "追加到记事本",
            6: "日志",
            7: "替换记事本",
            8: "命令队列",
            9: "变量",
            10: "执行",
            11: "快速行走",
            12: "脚本",
            13: "立刻发送",
            14: "脚本(过滤后)",
        },
        BatchCommandScripts: [],
        BatchCommandFormVisible: false,
        BatchCommandForm: {
            Scripts: [],
            Command: '',
        }
    }
    var vm = new Vue({
        el: "#app",
        data: data,
        mounted: function () {
            this.advancemode = (localStorage.getItem("hellclient-advancemode") != "");
            document.getElementsByTagName("body")[0].style.visibility = "visible";
        },
        components: {
            "avatar": avatar,
        },
        methods: {
            onEsc:function(){
                vm.suggestion=[]
            },
            onKeyUp: function (event) {
                app.KeyUp(event)
            },
            onInputKeyDown: function (event) {
                app.onInputKeyDown(event)
            },
            hideSuggestion:function(){
                vm.suggestion=[]
                vm.historypos=0
            },
            selectAll:function(){
                vm.$nextTick(function(){
                    var dom=document.getElementById("user-input").getElementsByTagName("input")[0]
                    dom.select();
                })
            },
            onFindHistory:function(){

                if (vm.historypos==0){
                    vm.cmd = vm.lastinput
                    vm.selectAll();
                    return
                }
                if (vm.historypos<=vm.suggestion.length){
                    vm.cmd=vm.suggestion[vm.suggestion.length-vm.historypos]
                    vm.selectAll();
                    return
                }
                var history=vm.history[vm.history.length-(vm.historypos-vm.suggestion.length)]
                if (history){
                    vm.cmd=history
                    vm.selectAll();
                }
            },
            onUp: function () {
                vm.historypos=vm.historypos+1
                if (vm.historypos>(vm.history.length+vm.suggestion.length)){
                    vm.historypos=0
                }
                vm.onFindHistory()
            },
            onDown: function () {
                vm.historypos=vm.historypos-1
                if (vm.historypos <= 0) {
                    vm.historypos = 0
                }
                vm.onFindHistory()
            },
            onSuggest:function(event){
                vm.cmd=event
                vm.selectAll();
                vm.suggestion=[]
            },
            onInputChange: function () {
                vm.historypos=0
                var val = document.getElementById("user-input").getElementsByTagName("input")[0].value
                vm.lastinput=val
                if (vm.history) {
                    vm.suggestion = []
                    var val = document.getElementById("user-input").getElementsByTagName("input")[0].value
                    if (val) {
                        vm.history.forEach(function (data) {
                            if (data.indexOf(val) > -1) {
                                vm.suggestion.push(data)
                            }
                        })
                        if (vm.suggestion.length > 5) {
                            vm.suggestion = vm.suggestion.slice(-10)
                        }
                    }
                }
            },
            send: function () {
                app.send("send", this.cmd)
                vm.historypos = 0
                vm.suggestion = []
                vm.selectAll();
                vm.lastinput=""
            },
            domasssend: function () {
                if (vm.MassSendForm) {
                    app.send("masssend", vm.MassSendForm.Value)
                }
                vm.MassSendForm = null
                vm.MassSendFormVisible = false
            },
            onChange: function (current) {
                if (vm.clients.length) {
                    app.send("change", current.name)
                }
                return false
            },
            onClient: function (client) {
                app.send("change", client.ID)
            },
            onGamelistClick: function (row, column, event) {
                if (vm.clients.length) {
                    app.send("change", row.ID)
                }
                this.allgameVisible = false
            },
            onButton: function (data) {
                onButton[data]()
            },
            onOpen: function (id) {
                onButton.open(id)
            },
            onUseScript: function (script) {
                this.scriptlistVisible = false
                app.send("usescript", [this.current, script])
            },
            onUpdateTrigger: function (data) {
                vm.saveTriggerFail = [];
                vm.triggerName = data.Name;
                vm.triggerSaveForm = data;
                vm.triggerSaveFormVisible = true;
                vm.triggersVisible = true;
            },
            onHistory: function (command) {
                app.send("send", command)
            },
            onDeleteTimer: function (id) {
                vm.$confirm('是否要删除该计时器?', '删除', {
                    confirmButtonText: '删除',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    app.send("deleteTimer", [vm.current, id])
                }).catch(() => {
                })
            },
            onUpdateTimer: function (id) {
                app.send("loadTimer", [vm.current, id])
                vm.updatingTimer = {
                    ID: id,
                    Form: {},
                }
                vm.timerUpdateFormVisible = true
            },
            onDeleteAlias: function (id) {
                vm.$confirm('是否要删除该别名?', '删除', {
                    confirmButtonText: '删除',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    app.send("deleteAlias", [vm.current, id])
                }).catch(() => {
                })
            },
            onUpdateAlias: function (id) {
                app.send("loadAlias", [vm.current, id])
                vm.updatingAlias = {
                    ID: id,
                    Form: {},
                }
                vm.aliasUpdateFormVisible = true
            },
            onDeleteTrigger: function (id) {
                vm.$confirm('是否要删除该触发器?', '删除', {
                    confirmButtonText: '删除',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    app.send("deleteTrigger", [vm.current, id])
                }).catch(() => {
                })
            },
            onUpdateTrigger: function (id) {
                app.send("loadTrigger", [vm.current, id])
                vm.updatingTrigger = {
                    ID: id,
                    Form: {},
                }
                vm.triggerUpdateFormVisible = true
            },
            onDeleteVariable: function (name) {
                vm.$confirm('是否要删除该变量?', '删除', {
                    confirmButtonText: '删除',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    app.send("deleteParam", [vm.current, name])
                }).catch(() => {
                })
            },
            onUpdateVariable: function (name, current) {
                vm.$prompt('请输入变量值', '编辑变量' + name, {
                    confirmButtonText: '编辑',
                    cancelButtonText: '取消',
                    inputValue: current,
                    inputType: "textarea",
                    customClass: "update-variable",
                }).then(({ value }) => {
                    app.send("updateParam", [vm.current, name, value])
                }).catch(() => {
                });

            },
            onUpdateRequiredParam: function (row) {
                vm.$prompt(row.Intro, '设置变量' + row.Name + "[" + row.Desc + "]", {
                    confirmButtonText: '设置',
                    cancelButtonText: '取消',
                    inputType: "textarea",
                    customClass: "update-required",
                    inputValue: vm.paramsinfo.Params[row.Name],
                }).then(({ value }) => {
                    app.send("updateParam", [vm.current, row.Name, value])
                }).catch(() => {
                });
            },
            onUpdateParamComment: function (row) {
                vm.$prompt("备注", '备注变量' + row.Name + "[" + row.Desc + "]", {
                    confirmButtonText: '备注',
                    cancelButtonText: '取消',
                    customClass: "update-comment",
                    inputType: "textarea",
                    inputValue: vm.paramsinfo.ParamComments[row.Name],
                }).then(({ value }) => {
                    app.send("updateParamComment", [vm.current, row.Name, value])
                }).catch(() => {
                });
            },
            callback: function (msg, code, data) {
                var cb = {
                    Name: msg.Name,
                    Script: msg.Script,
                    ID: msg.ID,
                    Code: code,
                    data: data,
                }
                if (msg.Script) {
                    app.send("callback", [this.current, JSON.stringify(cb)])
                }
            },
            onDrop: function () {
                vm.allLinesVisible = false
                vm.cmd = vm.cmd + document.getSelection().toString()
            },
            gamelistRowClassName: function (data) {
                return data.row.Running ? "game-list-running" : "game-list-not-running"
            },
            handleSelectVisualPromptList: function (index, row) {
                this.callback(this.visualPrompt, 0, row.Key)
                vm.visualPrompt = null
                vm.visualPromptVisible = false
            },
            handleSelectUserinputList: function (index, row) {
                this.callback(this.userinputList, 0, row.Key)
                vm.userinputList = null
                vm.userinputListVisible = false
                vm.userinputListSearch = ""
            },
            onUserinputListClose: function () {
                this.callback(this.userinputList, -1, "")
                vm.userinputList = null
                vm.userinputListVisible = false
                vm.userinputListSearch = ""
            },
            onUserinputDatagridClose: function () {
                this.callback(this.userinputDatagrid, -1, "")
                vm.userinputDatagrid = null
                vm.userinputDatagridVisible = false
            },
            onVisualPromptOpen: function () {
                vm.$refs.visualPromptValue.select()
            },
            assist: function () {
                app.send("assist", this.current)
            },
            masssend: function () {
                vm.MassSendForm = {},
                    vm.MassSendFormVisible = true
            },
            onMDClick: function (event) {
                if (event.target.localName == "a") {
                    this.callback(this.userinputNote, 0, event.target.attributes.href.value)
                    event.preventDefault()
                }

            },
            updateRequiredParams: function () {
                app.send("updateRequiredParams", { Current: vm.current, RequiredParams: this.requiredParams })
            },
            onUpdateScriptRequiredParam: function (data) {
                this.requiredParam = data
                this.createFail = []
                this.requiredParamUpdateForm = {
                    Name: data.Name,
                    Desc: data.Desc,
                    Intro: data.Intro,
                };
                this.requiredParamUpdateFormVisible = true

            },
            RequiredParamUp: function (index) {
                if (index <= 0) {
                    return
                }
                this.requiredParams.splice(index, 0, vm.requiredParams.splice(index - 1, 1)[0])
                this.updateRequiredParams()

            },
            RequiredParamDown: function (index) {
                if (index >= this.requiredParams.length - 1) {
                    return
                }
                this.requiredParams.splice(index, 0, vm.requiredParams.splice(index + 1, 1)[0])
                this.updateRequiredParams()
            },
            RequiredParamRemove: function (index) {
                vm.requiredParams.splice(index, 1)
                this.updateRequiredParams()
            },
            onUserinputListMutliChange: function (val) {
                this.userinputList.Data.Values = []
                var self = this
                val.forEach(function (data) {
                    self.userinputList.Data.Values.push(data.Key)
                })
            },
            onUserinputDatagridPage: function (page) {
                if (this.userinputDatagrid.Data.OnPage) {
                    var data = {
                        ID: this.userinputDatagrid.ID,
                        Name: this.userinputDatagrid.Name,
                        Script: this.userinputDatagrid.Data.OnPage,
                    }
                    vm.callback(data, 0, page + "")
                }
            },
            handleUserinputDatagridView: function (index, row) {
                if (this.userinputDatagrid.Data.OnView) {
                    var data = {
                        ID: this.userinputDatagrid.ID,
                        Name: this.userinputDatagrid.Name,
                        Script: this.userinputDatagrid.Data.OnView,
                    }
                    vm.callback(data, 0, row.Key)
                }
            },
            handleUserinputDatagridSelect: function (index, row) {
                if (this.userinputDatagrid.Data.OnSelect) {
                    var data = {
                        ID: this.userinputDatagrid.ID,
                        Name: this.userinputDatagrid.Name,
                        Script: this.userinputDatagrid.Data.OnSelect,
                    }
                    vm.callback(data, 0, row.Key)
                }
            },
            handleUserinputDatagridUpdate: function (index, row) {
                if (this.userinputDatagrid.Data.OnUpdate) {
                    var data = {
                        ID: this.userinputDatagrid.ID,
                        Name: this.userinputDatagrid.Name,
                        Script: this.userinputDatagrid.Data.OnUpdate,
                    }
                    vm.callback(data, 0, row.Key)
                }
            },
            handleUserinputDatagridDelete: function (index, row) {
                if (this.userinputDatagrid.Data.OnDelete) {
                    var data = {
                        ID: this.userinputDatagrid.ID,
                        Name: this.userinputDatagrid.Name,
                        Script: this.userinputDatagrid.Data.OnDelete,
                    }
                    vm.callback(data, 0, row.Key)
                    vm.$confirm('是否要删除该数据?', row.Value, {
                        confirmButtonText: '删除',
                        cancelButtonText: '取消',
                        type: 'warning'
                    }).then(() => {
                        vm.callback(data, 0, row.Key)
                    }).catch(() => {
                    })
                }
            },
            onUserinputDatagridFilter: function () {
                var vm = this
                if (vm.userinputDatagrid.Data.OnFilter) {
                    var data = {
                        ID: this.userinputDatagrid.ID,
                        Name: this.userinputDatagrid.Name,
                        Script: this.userinputDatagrid.Data.OnFilter,
                    }
                    vm.$prompt("设置搜索关键字", {
                        confirmButtonText: '提交',
                        cancelButtonText: '取消',
                        inputValue: vm.userinputDatagrid.Data.Filter,
                        onClose: function () {
                            vm.callback(data, -1, "")
                        }
                    }).then(({ value }) => {
                        vm.callback(data, 0, value)
                    }).catch(() => {
                    });
                }
            },
            onUserinputDatagridCreate: function () {
                var vm = this
                if (vm.userinputDatagrid.Data.OnCreate) {
                    var data = {
                        ID: this.userinputDatagrid.ID,
                        Name: this.userinputDatagrid.Name,
                        Script: this.userinputDatagrid.Data.OnCreate,
                    }
                    vm.callback(data, 0, "")
                }
            },
            onHUDClick: function (e) {
                app.send("hudclick", { X: e.offsetX / e.target.width, Y: e.offsetY / e.target.height })
            },
            doFocus: function () {
                let input = document.getElementById("mud-input")
                if (input) { input.focus() }
            },
            onClientDragStart: function (e) {
                e.dataTransfer.setData("clientindex", e.target.dataset.index);
            },
            onClientDrop: function (e) {
                if (!isNaN(e.dataTransfer.getData("clientindex"))) {
                    let from = e.dataTransfer.getData("clientindex") - 0
                    for (var i = 0; i < e.path.length; i++) {
                        if (e.path[i].dataset["index"] != undefined) {
                            e.preventDefault()
                            let to = e.path[i].dataset["index"] - 0
                            if (from != to) {
                                vm.doSwipClient(from, to)
                            }
                            return
                        }
                    }
                }
            },
            doSwipClient: function (from, to) {
                let clients = vm.clients
                if (from < 0 || to < 0 || from >= clients.length || to >= clients.length) {
                    return
                }
                let order = []
                for (var i = 0; i < clients.length; i++) {
                    order.push(clients[i].ID)
                }
                let fromid = order[from]
                order[from] = order[to]
                order[to] = fromid
                app.send("sortclients", order)
            },
            onClientDargOver: function (e) {
                if (e.dataTransfer.types.includes("clientindex")) {
                    e.preventDefault()
                }
            }
        }
    })
    return vm
})