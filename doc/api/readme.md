# Hellclient API列表

[与Mushclient接口兼容情况](mush.md)
## 别名接口

* [AddAlias 添加别名](apialias.md#AddAlias)
* [DeleteAlias 删除别名](apialias.md#DeleteAlias)
* [DeleteAliasGroup 删除别名组](apialias.md#DeleteAliasGroup)
* [DeleteTemporaryAliases 删除临时别名](apialias.md#DeleteTemporaryAliases)
* [EnableAlias 激活别名](apialias.md#EnableAlias)
* [EnableAliasGroup 激活别名组](apialias.md#EnableAliasGroup)
* [GetAliasInfo 获取别名信息](apialias.md#GetAliasInfo)
* [GetAliasList 获取脚本别名名称列表](apialias.md#GetAliasList)
* [GetAliasOption 获取别名选项](apialias.md#GetAliasOption)
* [IsAlias 判断别名是否存在](apialias.md#IsAlias)
* [SetAliasOption 设置别名选项](apialias.md#SetAliasOption)

## 计时器接口

* [AddTimer 添加计时器](apitimer.md#AddTimer)
* [DeleteTimer 删除计时器](apitimer.md#DeleteTimer)
* [DeleteTimerGroup 删除计时器组](apitimer.md#DeleteTimerGroup)
* [DeleteTemporaryTimers 删除临时计时器](apitimer.md#DeleteTemporaryTimers)
* [EnableTimer 激活计时器](apitimer.md#EnableTimer)
* [EnableTimerGroup 激活计时器组](apitimer.md#EnableTimerGroup)
* [GetTimerInfo 获取计时器信息](apitimer.md#GetTimerInfo)
* [GetTimerList 获取脚本计时器列表](apitimer.md#GetTimerList)
* [GetTimerOption 获取计时器选项](apitimer.md#GetTimerOption)
* [IsTimer 判断计时器是否存在](apitimer.md#IsTimer)
* [ResetTimer 重置计时器](apitimer.md#ResetTimer)
* [ResetTimers 重置全部计时器](apitimer.md#ResetTimers)
* [SetTimerOption 设置计时器选项](apitimer.md#SetTimerOption)

## 触发器接口

* [AddTrigger 添加触发器](apitrigger.md#AddTrigger)
* [AddTriggerEx 高阶添加触发器](apitrigger.md#AddTriggerEx)
* [DeleteTrigger 删除触发器](apitrigger.md#DeleteTrigger)
* [DeleteTriggerGroup 删除触发器组](apitrigger.md#DeleteTriggerGroup)
* [EnableTrigger 激活触发器](apitrigger.md#EnableTrigger)
* [EnableTriggerGroup 激活触发器组](apitrigger.md#EnableTriggerGroup)
* [GetTriggerInfo 获取触发器信息](apitrigger.md#GetTriggerInfo)
* [GetTriggerList 获取脚本触发器列表](apitrigger.md#GetTriggerList)
* [GetTriggerOption 获取触发器选项](apitrigger.md#GetTriggerOption)
* [GetTriggerWildcard 获取触发器匹配值](apitrigger.md#GetTriggerWildcard)
* [IsTrigger 判断触发器是否存在](apitrigger.md#IsTrigger)
* [SetTriggerOption 设置触发器选项](apitrigger.md#SetTriggerOption)
* [StopEvaluatingTriggers 停止执行触发器](apitrigger.md#StopEvaluatingTriggers)

## 游戏管理接口
* GetWorldById
* GetWorld
* GetWorldID
* GetWorldIdList
* GetWorldList
* WorldName
* WorldAddress
* WorldPort
## 发送接口
* print
* Note
* SendImmediate
* Send
* SendNoEcho
* GetVariable
* Execute
* DiscardQueue
* LockQueue
* GetQueue
* Queue
* DoAfter
* DoAfterNote
* DoAfterSpeedWalk
* DoAfterSpecial
* SetSpeedWalkDelay
* GetSpeedWalkDelay

## 变量接口

* SetVariable
* DeleteVariable
* GetVariableList
* GetVariableComment
* SetVariableComment

## 杂项
* Version
* Hash
* Base64Encode
* Base64Decode
* Trim
* GetUniqueNumber
* GetUniqueID
* CreateGUID
* SplitN
* UTF8Len
* UTF8Sub
## 连接管理

* Connect
* IsConnected
* Disconnect

## 界面
* FlashIcon
* SetStatus
* DeleteCommandHistory
* Info
* InfoClear
* GetAlphaOption
* SetAlphaOption
* GetInfo
* GetGlobalOption
## 色彩接口
* ColourNameToRGB
* BoldColour
* NormalColour

## 文件处理
* ReadFile
* ReadLines
* HasHomeFile
* ReadHomeFile
* WriteHomeFile
* ReadHomeLines
* WriteLog
* CloseLog
* OpenLog
* FlushLog

## 输出内容接口
* GetLinesInBufferCount
* DeleteOutput
* DeleteLines
* GetLineCount
* GetRecentLines
* GetLineInfo
* GetStyleInfo

## 通讯接口

* Broadcast
* Notify

## 授权接口
* CheckPermissions
* RequestPermissions
* CheckTrustedDomains
* RequestTrustDomains