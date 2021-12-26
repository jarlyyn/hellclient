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
* [DeleteGroup 按组删除元素](apiworld.md#DeleteGroup)
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
* [Version 版本信息](apimisc.md#Version)
* [Hash 摘要](apimisc.md#Hash)
* [Base64Encode Base64编码](apimisc.md#Base64Encode)
* [Base64Decode Base64解码](apimisc.md#Base64Decode)
* [Trim 去空格](apimisc.md#Trim)
* [GetUniqueNumber 获取唯一数值](apimisc.md#GetUniqueNumber)
* [GetUniqueID 获取唯一ID](apimisc.md#GetUniqueID)
* [CreateGUID 获取GUID](apimisc.md#CreateGUID)
* [SplitN 分割字符串](apimisc.md#SplitN)
* [UTF8Len 获取UTF8长度](apimisc.md#UTF8Len)
* [UTF8Sub 获取UTF8子字符串](apimisc.md#UTF8Sub)
* [ToUTF8 转换为UTF8字符串](apimisc.md#ToUTF8)
* [FromUTF8 转换自UTF8字符串](apimisc.md#FromUTF8)
## 连接管理

* [Connect 连接](apiconnect.md#Connect)
* [IsConnected 判断接连状态](apiconnect.md#IsConnected)
* [Disconnect 断开连接](apiconnect.md#Disconnect)

## 界面
* [FlashIcon 废弃](apiui.md#FlashIcon)
* [SetStatus 设置状态文本](apiui.md#SetStatus)
* [DeleteCommandHistory 删除命令记录](apiui.md#DeleteCommandHistory)
* [Info 追加信息文本](apiui.md#Info)
* [InfoClear 清除信息文本](apiui.md#InfoClear)
* [GetAlphaOption 获取选项](apiui.md#GetAlphaOption)
* [SetAlphaOption 设置选项](apiui.md#SetAlphaOption)
* [GetInfo 获取信息](apiui.md#GetInfo)
* [GetGlobalOption 获取全局选项](apiui.md#GetGlobalOption)
 

## 色彩接口
* [ColourNameToRGB 色彩名转rgb](apicolor.md#)
* [NormalColour 获取普通色彩rgb](apicolor.md#)
* [BoldColour 获取高亮色彩rgb](apicolor.md#)

## 文件处理
* [ReadFile 读取脚本文件](apifile.md#ReadFile)
* [ReadLines 读取脚本文件并分行](apifile.md#ReadLines)
* [HasHomeFile 检查用户文件](apifile.md#HasHomeFile)
* [ReadHomeFile 读取用户文件](apifile.md#ReadHomeFile)
* [WriteHomeFile 写入用户文件](apifile.md#WriteHomeFile)
* [ReadHomeLines 读取用户文件并分行](apifile.md#ReadHomeLines)
* [WriteLog 写入日志](apifile.md#WriteLog)
* [CloseLog 废弃](apifile.md#CloseLog)
* [OpenLog 废弃](apifile.md#OpenLog)
* [FlushLog 废弃](apifile.md#FlushLog)

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