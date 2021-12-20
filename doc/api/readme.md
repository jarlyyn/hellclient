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
* [GetWorldById 废弃接口](apiworld.md#GetWorldById)
* [GetWorld 废弃接口](apiworld.md#GetWorld)
* [GetWorldID 获取当前游戏ID](apiworld.md#GetWorldID)
* [GetWorldIdList 废弃接口](apiworld.md#GetWorldIdList)
* [GetWorldList 废弃接口](apiworld.md#GetWorldList)
* [WorldName 获取游戏名](apiworld.md#WorldName)
* [WorldAddress 获取游戏网络地址](apiworld.md#WorldAddress)
* [WorldPort 获取游戏网络端口](apiworld.md#WorldPort)
## 发送接口
* [print 打印](apisend.md#print)
* [Note 显示](apisend.md#Note)
* [SendImmediate 立即发送](apisend.md#SendImmediate)
* [Send 发送](apisend.md#Send)
* [SendNoEcho 静默发送](apisend.md#SendNoEcho)
* [SendSpecial 高级发送](apisend.md#SendSpecial)
* [Execute 执行](apisend.md#Execute)
* [Queue 队列发送](apisend.md#Queue) 
* [DiscardQueue 取消队列](apisend.md#DiscardQueue)
* [LockQueue 锁定队列](apisend.md#LockQueue)
* [GetQueue 获取队列内容](apisend.md#GetQueue)
* [DoAfter 延迟发送](apisend.md#DoAfter)
* [DoAfterNote 延迟显示](apisend.md#DoAfterNote)
* [DoAfterSpeedWalk 延迟加入队列](apisend.md#DoAfterSpeedWalk)
* [DoAfterSpecial 高级延迟执行](apisend.md#DoAfterSpecial)
* [SetSpeedWalkDelay 设置队列延迟](apisend.md#SetSpeedWalkDelay/GetSpeedWalkDelay)
* [GetSpeedWalkDelay 获取队列延迟](apisend.md#SetSpeedWalkDelay/GetSpeedWalkDelay)

## 变量接口

* [GetVariable 获取变量值](apivariable.md#GetVariable)
* [SetVariable 设置变量](apivariable.md#SetVariable)
* [DeleteVariable 删除变量](apivariable.md#DeleteVariable)
* [GetVariableList 获取变量列表](apivariable.md#GetVariableList)
* [GetVariableComment 获取变量备注](apivariable.md#GetVariableComment)
* [SetVariableComment 设置变量备注](apivariable.md#SetVariableComment)

## 杂项
* [Version](apimisc.md#)
* [Hash](apimisc.md#)
* [Base64Encode](apimisc.md#)
* [Base64Decode](apimisc.md#)
* [Trim](apimisc.md#)
* [GetUniqueNumber](apimisc.md#)
* [GetUniqueID](apimisc.md#)
* [CreateGUID](apimisc.md#)
* [SplitN](apimisc.md#)
* [UTF8Len](apimisc.md#)
* [UTF8Sub](apimisc.md#)
## 连接管理

* [Connect](apiconnect.md#)
* [IsConnected](apiconnect.md#)
* [Disconnect](apiconnect.md#)

## 界面
* [FlashIcon](apiui.md#)
* [SetStatus](apiui.md#)
* [DeleteCommandHistory](apiui.md#)
* [Info](apiui.md#)
* [InfoClear](apiui.md#)
* [GetAlphaOption](apiui.md#)
* [SetAlphaOption](apiui.md#)
* [GetInfo](apiui.md#)
* [GetGlobalOption](apiui.md#)
## 色彩接口
* [ColourNameToRGB](apicolor.md#)
* [BoldColour](apicolor.md#)
* [NormalColour](apicolor.md#)

## 文件处理
* [ReadFile](apifile.md#)
* [ReadLines](apifile.md#)
* [HasHomeFile](apifile.md#)
* [ReadHomeFile](apifile.md#)
* [WriteHomeFile](apifile.md#)
* [ReadHomeLines](apifile.md#)
* [WriteLog](apifile.md#)
* [CloseLog](apifile.md#)
* [OpenLog](apifile.md#)
* [FlushLog](apifile.md#)

## 输出内容接口
* [GetLinesInBufferCount](apioutput.md#)
* [DeleteOutput](apioutput.md#)
* [DeleteLines](apioutput.md#)
* [GetLineCount](apioutput.md#)
* [GetRecentLines](apioutput.md#)
* [GetLineInfo](apioutput.md#)
* [GetStyleInfo](apioutput.md#)

## 通讯接口

* [Broadcast](apicommunication.md#)
* [Notify](apicommunication.md#)

## 授权接口
* [CheckPermissions](apiauth.md#)
* [RequestPermissions](apiauth.md#)
* [CheckTrustedDomains](apiauth.md#)
* [RequestTrustDomains](apiauth.md#)