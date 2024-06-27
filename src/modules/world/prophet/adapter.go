package prophet

import (
	"github.com/herb-go/connections/room/message"
)

// 只有在这里注册过的信息才会发送给客户端
// 注册的适配器与 /src/world/msg/msg,go中的msgtype对应,由msg.go负责序列化为标准信息
// 适配器负责个根据当前连接状态决定是否将信息发送给客户端。
// 具体与客户端的交互可以参考 /resources/public/defaultui/js/handlers.js下的handler.[适配器名] 函数
// newRoomAdapter 为发送给当前游戏
// newUserAdapter 为发送给客户端，无视游戏

func initAdapter(p *Prophet, adapter *message.Adapter) {
	//单行信息，预期客户端在主输出新增一行
	adapter.Register("line", p.newRoomAdapter("line"))
	//多行信息，预期客户端清空主输出并更新内容
	adapter.Register("lines", p.newRoomAdapter("lines"))
	//提示行更新，预期更新客户端提示行部分
	adapter.Register("prompt", p.newRoomAdapter("prompt"))
	//触发列表，预期客户端弹出触发器列表
	adapter.Register("triggers", p.newRoomAdapter("triggers"))
	//客户端信息，预期客户端更新所有已经打开的客户端的信息
	adapter.Register("clients", p.newUserAdapter("clients"))
	//连线信息，通知客户端更新已经连线状态
	adapter.Register("connected", p.newUserAdapter("connected"))
	//断线信息，通知客户端更新已经连线状态
	adapter.Register("disconnected", p.newUserAdapter("disconnected"))
	//创建失败的错误信息，预期客户端提示错误
	adapter.Register("createFail", p.newUserAdapter("createFail"))
	//创建成功，预期客户端关闭创建窗口
	adapter.Register("createSuccess", p.newUserAdapter("createSuccess"))
	//更新成功，预期客户端关闭更新窗口
	adapter.Register("updateSuccess", p.newUserAdapter("updateSuccess"))
	//触发维护失败，预期客户端提示错误
	adapter.Register("triggerFail", p.newRoomAdapter("triggerFail"))
	//触发维护成功，预期客户端关闭更新窗口
	adapter.Register("triggerSuccess", p.newRoomAdapter("triggerSuccess"))
	//历史信息，预期客户端弹出历史信息界面
	adapter.Register("allLines", p.newRoomAdapter("allLines"))
	//未打开游戏列表，预期客户端弹出打开游戏界面
	adapter.Register("notopened", p.newUserAdapter("notopened"))
	//脚本信息，预期客户端弹出脚本信息详情
	adapter.Register("scriptinfo", p.newRoomAdapter("scriptinfo"))
	//创建脚本失败，预期客户端提示用户错误信息
	adapter.Register("createScriptFail", p.newUserAdapter("createScriptFail"))
	//创建脚本成功，预期客户端关闭创建界面
	adapter.Register("createScriptSuccess", p.newUserAdapter("createScriptSuccess"))
	//更新脚本成功，预期客户端关闭更新界面
	adapter.Register("updateScriptSuccess", p.newUserAdapter("updateScriptSuccess"))
	//脚本列表，预期客户端弹出脚本列表，并选择游戏脚本
	adapter.Register("scriptinfoList", p.newUserAdapter("scriptinfoList"))
	//更新状态行，一般在脚本SetStatus后更新，预期客户端更新状态行
	adapter.Register("status", p.newRoomAdapter("status"))
	//输入历史更新，预期客户端更新历史信息/补全信息。
	adapter.Register("history", p.newRoomAdapter("history"))
	//用户计时器信息列表，预期客户端弹出用户计时器列表
	adapter.Register("usertimers", p.newRoomAdapter("usertimers"))
	//脚本时器信息列表，预期客户端弹出脚本计时器列表
	adapter.Register("scripttimers", p.newRoomAdapter("scripttimers"))
	//创建计时器成功，预期客户端关闭计时器创建界面
	adapter.Register("createTimerSuccess", p.newRoomAdapter("createTimerSuccess"))
	//计时器详情，预期客户端弹出计时器详情
	adapter.Register("timer", p.newRoomAdapter("timer"))
	//更新计时器成功，预期客户端关闭计时器更新界面
	adapter.Register("updateTimerSuccess", p.newRoomAdapter("updateTimerSuccess"))
	//用户别名列表，预期客户端弹出用户别名列表界面
	adapter.Register("useraliases", p.newRoomAdapter("useraliases"))
	//脚本别名列表，预期客户端弹出脚本别名列表界面
	adapter.Register("scriptaliases", p.newRoomAdapter("scriptaliases"))
	//创建别名成功，预期客户端会关闭创建别名窗口
	adapter.Register("createAliasSuccess", p.newRoomAdapter("createAliasSuccess"))
	//别名详情，预期客户端会弹出别名详情
	adapter.Register("alias", p.newRoomAdapter("alias"))
	//更新别名成功，预期客户端会关闭更新别名差窗口
	adapter.Register("updateAliasSuccess", p.newRoomAdapter("updateAliasSuccess"))
	//用户触发器列表，预期客户端弹出用户触发器列表界面
	adapter.Register("usertriggers", p.newRoomAdapter("usertriggers"))
	//脚本触发器列表，预期客户端弹出脚本触发器列表界面
	adapter.Register("scripttriggers", p.newRoomAdapter("scripttriggers"))
	//创建触发器成功，预期客户端关闭创建触发器界面
	adapter.Register("createTriggerSuccess", p.newRoomAdapter("createTriggerSuccess"))
	//触发器详情，预期客户端弹出触发器详情
	adapter.Register("trigger", p.newRoomAdapter("trigger"))
	//更新触发器成功，预期客户端关闭更新触发器界面
	adapter.Register("updateTriggerSuccess", p.newRoomAdapter("updateTriggerSuccess"))
	//变量列表，预期客户端弹出变量界面
	adapter.Register("paramsinfo", p.newRoomAdapter("paramsinfo"))
	//变量更新成功，预期客户端关闭更新变量界面
	adapter.Register("paramupdated", p.newRoomAdapter("paramupdated"))
	//变量删除成功，预期客户端维护变量界面
	adapter.Register("paramdeleted", p.newRoomAdapter("paramdeleted"))
	//变量备注更新成功，预期客户端关闭变量备注界面
	adapter.Register("paramcommentupdated", p.newRoomAdapter("paramcommentupdated"))
	//脚本信息(userinput),预期客户端根据信息内容与用户进行交互
	adapter.Register("scriptMessage", p.newRoomAdapter("scriptMessage"))
	//交换机状态，预期客户端更新交换机信息
	adapter.Register("switchStatus", p.newUserAdapter("switchStatus"))
	//版本信息，预期客户端弹出版本信息界面
	adapter.Register("version", p.newUserAdapter("version"))
	//API信息，连接时发送，预期客户端根据API版本显示不同界面
	adapter.Register("apiversion", p.newUserAdapter("apiversion"))
	//游戏信息界面，预期客户端弹出游戏信息
	adapter.Register("worldSettings", p.newRoomAdapter("worldSettings"))
	//脚本信息界面，预期客户端弹出脚本信息
	adapter.Register("scriptSettings", p.newRoomAdapter("scriptSettings"))
	//游戏变量界面，预期客户端显示游戏变量维护界面
	adapter.Register("requiredParams", p.newRoomAdapter("requiredParams"))
	//默认服务器信息，预期客户端在创建游戏时填入默认服务器信息
	adapter.Register("defaultServer", p.newUserAdapter("defaultServer"))
	//默认编码信息，预期客户端在创建游戏时填入默认编码信息
	adapter.Register("defaultCharset", p.newUserAdapter("defaultCharset"))
	//授权请求，预期客户端弹授权界面
	adapter.Register("requestPermissions", p.newRoomAdapter("requestPermissions"))
	//授权域名信息，预期客户端弹出授权域名界面
	adapter.Register("requestTrustDomains", p.newRoomAdapter("requestTrustDomains"))
	//已授权界面，预期客户端弹出已授权内容界面
	adapter.Register("authorized", p.newRoomAdapter("authorized"))
	//查找历史信息，已废弃
	adapter.Register("foundhistory", p.newRoomAdapter("foundhistory"))
	//HUD全量更新，预期客户端更新完整的HUD信息
	adapter.Register("hudcontent", p.newRoomAdapter("hudcontent"))
	//HUD部分更新，预期客户端更新HUD中的指定行
	adapter.Register("hudupdate", p.newRoomAdapter("hudupdate"))
	//客户端信息，预期客户端更新客户端界面中的单个客户端信息
	adapter.Register("clientinfo", p.newConsoleAdapter("clientinfo"))
	//批量执行脚本信息，预期客户端弹出批量执行脚本界面
	adapter.Register("batchcommandscripts", p.newUserAdapter("batchcommandscripts"))

}
