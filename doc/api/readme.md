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
* [DumpAliases 导出别名列表](apialias.md#DumpAliases)
* [RestoreAliases 导入别名列表](apialias.md#RestoreAliases)

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
* [DumpTimers 导出计时器列表](apitimer.md#DumpTimers)
* [RestoreTimers 导入计时器列表](apitimer.md#RestoreTimers)

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
* [DumpTriggers 导出触发器列表](apitrigger.md#DumpTriggers)
* [RestoreTriggers 导入触发器列表](apitrigger.md#RestoreTriggers)
  

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
* [Save 保存](apiworld.md#Save)


## 发送接口
* [print 打印](apisend.md#print)
* [Note 显示](apisend.md#Note)
* [PrintSystem 模拟系统提示](apisend.md#PrintSystem)
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
* [Encrypt 加密字符串](apimisc.md#Encrypt)
* [Decrypt 解密字符串](apimisc.md#Decrypt)
* [Milliseconds 毫秒时间戳](apimisc.md#Milliseconds)
  
## 连接管理
* [Connect 连接](apiconnect.md#Connect)
* [IsConnected 判断接连状态](apiconnect.md#IsConnected)
* [Disconnect 断开连接](apiconnect.md#Disconnect)

## 界面
* [GetHUDSize 获取HUD尺寸](apiui.md#GetHUDSize)
* [SetHUDSize 设置HUD尺寸](apiui.md#SetHUDSize)
* [GetHUDContent 获取HUD内容](apiui.md#GetHUDContent)
* [UpdateHUDContent 更新HUD内容](apiui.md#UpdateHUDContent)

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
* [HasFile 读取脚本文件](apifile.md#HasFile)
* [ReadFile 读取脚本文件](apifile.md#ReadFile)
* [ReadLines 读取脚本文件并分行](apifile.md#ReadLines)
* [HasHomeFile 检查用户文件](apifile.md#HasHomeFile)
* [ReadHomeFile 读取用户文件](apifile.md#ReadHomeFile)
* [WriteHomeFile 写入用户文件](apifile.md#WriteHomeFile)
* [ReadHomeLines 读取用户文件并分行](apifile.md#ReadHomeLines)
* [MakeHomeFolder 创建用户目录](apifile.md#MakeHomeFolder)


* [HasModFile 检查用户文件](apifile.md#HasModFile)
* [ReadModFile 读取用户文件](apifile.md#ReadModFile)
* [ReadModLines 读取用户文件并分行](apifile.md#ReadModLines)
* [WriteLog 写入日志](apifile.md#WriteLog)
* [CloseLog 废弃](apifile.md#CloseLog)
* [OpenLog 废弃](apifile.md#OpenLog)
* [FlushLog 废弃](apifile.md#FlushLog)

## 输出内容接口
* [GetLinesInBufferCount 获取历史行数](apioutput.md#GetLinesInBufferCount)
* [DeleteOutput 废弃](apioutput.md#DeleteOutput)
* [DeleteLines 删除历史内容](apioutput.md#DeleteLines)
* [GetLineCount 获取接受总行数](apioutput.md#GetLineCount)
* [GetRecentLines 获取近期内容](apioutput.md#GetRecentLines)
* [GetLineInfo 获取指定行信息](apioutput.md#GetLineInfo)
* [GetStyleInfo 获取样式信息](apioutput.md#GetStyleInfo)
* [DumpOutput 导出输出](apioutput.md#DumpOutput)
* [ConcatOutput 连接输出](apioutput.md#ConcatOutput)
* [SliceOutput 切割输出](apioutput.md#SliceOutput)
* [OutputToText 输出转文字](apioutput.md#OutputToText)
* [FormatOutput 格式化输出](apioutput.md#FormatOutput)
* [PrintOutput 格式化输出](apioutput.md#PrintOutput)
* [Simulate 模拟发送文字](apioutput.md#Simulate)
* [SimulateOutput 模拟发送输出](apioutput.md#SimulateOutput)
* [OmitOutput 屏蔽当前行显示](apioutput.md#OmitOutput)


## 通讯接口

* [Broadcast 广播](apicommunication.md#Broadcast)
* [Notify 通知](apicommunication.md#Notify)
* [Request 请求](apicommunication.md#Request)

## 授权接口
* [CheckPermissions 检查权限](apiauth.md#CheckPermissions)
* [RequestPermissions 请求权限](apiauth.md#RequestPermissions)
* [CheckTrustedDomains 检查信任域名](apiauth.md#CheckTrustedDomains)
* [RequestTrustDomains 请求信任域名](apiauth.md#RequestTrustDomains)

## 地图接口
* [术语](apimapper.md#术语)
* [Mapper.Reset 重置地图](apimapper.md#MapperReset)
* [Mapper.ResetTemporary 重置临时路径](apimapper.md#MapperResetTemporary)
* [Mapper.AddTags 添加标签](apimapper.md#MapperAddTags)
* [Mapper.SetTag 设置标签](apimapper.md#MapperSetTag)
* [Mapper.SetTags 设置标签列表](apimapper.md#MapperSetTags)
* [Mapper.FlashTags 清理标签](apimapper.md#MapperFlashTags)
* [Mapper.Tags 获取标签](apimapper.md#MapperTags)
* [Mapper.GetPath 获取路径](apimapper.md#MapperGetPath)
* [Mapper.AddPath 添加路径](apimapper.md#MapperAddPath)
* [Mapper.AddTemporaryPath 添加临时路径](apimapper.md#MapperAddTemporaryPath)
* [Mapper.NewPath 新建路径](apimapper.md#MapperNewPath)
* [Mapper.GetRoomID 获取房间id](apimapper.md#MapperGetRoomID)
* [Mapper.GetRoomName 获取房间名](apimapper.md#MapperGetRoomName)
* [Mapper.SetRoomName 设置房间](apimapper.md#MapperSetRoomName)
* [Mapper.ClearRoom 清理房间](apimapper.md#MapperClearRoom)
* [Mapper.RemoveRoom 清理房间](apimapper.md#MapperRemoveRoom)
* [Mapper.NewArea 新建区域](apimapper.md#MapperNewArea)
* [Mapper.GetExits 获取房间出口](apimapper.md#MapperGetExits)
* [Mapper.FlyList 获取飞行列表](apimapper.md#MapperFlyList)
* [Mapper.SetFlyList 设置飞行列表](apimapper.md#MapperSetFlyList)

## 节拍限流器
* [术语](apimetronome.md#术语)
* [Metronome.GetBeats 获取发送限制](apimetronome.md#MetronomeGetBeats)
* [Metronome.SetBeats 设置发送限制](apimetronome.md#MetronomeSetBeats)
* [Metronome.Reset 重置限流器](apimetronome.md#MetronomeReset)
* [Metronome.GetSpace 获取剩余空间](apimetronome.md#MetronomeGetSpace)
* [Metronome.GetQueue 获取队列](apimetronome.md#MetronomeGetQueue)
* [Metronome.Discard 放弃队列](apimetronome.md#MetronomeDiscard)
* [Metronome.LockQueue 锁定队列](apimetronome.md#MetronomeLockQueue)
* [Metronome.Full 填充完整周期](apimetronome.md#MetronomeFull)
* [Metronome.FullTick 填充当前周期](apimetronome.md#MetronomeFullTick)
* [Metronome.GetInterval 获取发送间隔](apimetronome.md#MetronomeGetInterval)
* [Metronome.SetInterval 设置发送间隔](apimetronome.md#MetronomeSetInterval)
* [Metronome.GetTick 获取限流周期](apimetronome.md#MetronomeGetTick)
* [Metronome.SetTick 设置限流周期](apimetronome.md#MetronomeSetTick)
* [Metronome.Push 推送命令](apimetronome.md#MetronomePush)

## 用户输入

* [术语](apiuserinput.md#术语)
* [Userinput.Prompt 用户输入框](apiuserinput.md#UserinputPrompt)
* [Userinput.Confirm 用户确认框](apiuserinput.md#UserinputConfirm)
* [Userinput.Alert 用户提示](apiuserinput.md#UserinputAlert)
* [Userinput.Popup 弹框提示](apiuserinput.md#UserinputPopup)
* [Userinput.Note 弹出文本](apiuserinput.md#UserinputNote)
* [Userinput.Custom 自定义事件](apiuserinput.md#Custom)
* [Userinput.NewList 新建列表](apiuserinput.md#UserinputNewList)
* [Userinput.NewDatagrid 新建数据表格](apiuserinput.md#Userinputnewdatagrid)
* [Userinput.NewVisualPrompt 新建可视化输入](apiuserinput.md#UserinputNewVisualPrompt)
* [Userinput.HideAll 隐藏界面UI](apiuserinput.md#HideAll)
* [List.Append 列表追加](apiuserinput.md#ListAppend)
* [List.SetMulti 列表设置多选](apiuserinput.md#ListSetMulti)  
* [List.SetValue 列表设置值](apiuserinput.md#ListSetValue)
* [List.Publish 列表发布](apiuserinput.md#ListPublish) 
* [Datagrid.Append 数据表格追加数据](apiuserinput.md#DatagridAppend)
* [Datagrid.ResetItems 数据表格重置元素](apiuserinput.md#DatagridResetItems)
* [Datagrid.SetFilter 数据表格设置过滤值](apiuserinput.md#DatagridSetFilter)
* [Datagrid.GetFilter 数据表格获取过滤值](apiuserinput.md#DatagridGetFilter)
* [Datagrid.SetMaxPage 数据表格设置最大页数](apiuserinput.md#DatagridSetMaxPage)
* [Datagrid.SetPage 数据表格设置当前页](apiuserinput.md#DatagridSetPage)
* [Datagrid.GetPage 数据表格获取当前页](apiuserinput.md#DatagridGetPage)
* [Datagrid.SetOnCreate 数据表格设置创建回调](apiuserinput.md#DatagridSetOnCreate)
* [Datagrid.SetOnUpdate 数据表格设置更新回调](apiuserinput.md#DatagridSetOnUpdate)
* [Datagrid.SetOnSelect 数据表格设置选择回调](apiuserinput.md#DatagridSetOnSelect) 
* [Datagrid.SetOnView 数据表格设置查看回调](apiuserinput.md#DatagridSetOnView)
* [Datagrid.SetOnDelete 数据表格设置删除回调](apiuserinput.md#DatagridSetOnDelete)
* [Datagrid.SetOnFilter 数据表格设置过滤回调](apiuserinput.md#DatagridSetOnFilter)
* [Datagrid.SetOnPage 数据表格设置分页回调](apiuserinput.md#Datagridsetonpage)
* [Datagrid.Hide 数据表格隐藏](apiuserinput.md#DatagridHide)
* [Datagrid.Publish 数据表格发布](apiuserinput.md#DatagridPublish) 
* [VisualPrompt.SetMediatype 可视化输入设置媒体类型](apiuserinput.md#VisualPromptSetMediaType)
* [VisualPrompt.SetPortrait 可视化输入设置垂直模式](apiuserinput.md#VisualPromptSetPortrait)
* [VisualPrompt.SetRefreshCallback 可视化输入设置刷新回调](apiuserinput.md#SetRefreshCallback)
* [VisualPrompt.Append 列表追加](apiuserinput.md#VisualPromptpublish#Append) 
* [VisualPrompt.Publish 可视化输入发布](apiuserinput.md#VisualPromptpublish#Publish)

## http请求
* [术语](apihttp.md#术语)
* [HTTP.PraseURL 解析URL地址](apihttp.md#HTTPPraseURL)
* [HTTP.New 创建新请求](apihttp.md#HTTPNew)
* [Request.GetID 获取唯一id](apihttp.md#RequestGetID)
* [Request.GetURL 获取请求URL](apihttp.md#RequestGetURL)
* [Request.SetURL 设置请求URL](apihttp.md#RequestSetURL)
* [Request.GetMethod 获取请求Method](apihttp.md#Request.GetMethod)
* [Request.SetMethod 设置请求URL](apihttp.md#RequestSetMethod)
* [Request.GetBody 获取请求正文](apihttp.md#RequestGetBody)
* [Request.SetBody 设置请求正文](apihttp.md#RequestSetBody)
* [Request.SetHeader 设置请求头](apihttp.md#RequestSetHeader)
* [Request.AddHeader 添加请求头](apihttp.md#RequestAddHeader)
* [Request.DelHeader 删除请求头](apihttp.md#RequestDelHeader)
* [Request.GetHeader 获取请求头](apihttp.md#RequestGetHeader)
* [Request.HeaderValues 获取请求头全部值](apihttp.md#RequestHeaderValues)
* [Request.HeaderFields 获取请求头全部字段](apihttp.md#RequestHeaderFields)
* [Request.ResetHeader 重置请求头](apihttp.md#RequestResetHeader)
* [Request.AsyncExecute 异步执行请求](apihttp.md#RequestAsyncExecute)
* [Request.ExecuteStatus 获取执行状态](apihttp.md#RequestExecuteStatus)
* [Request.FinishedAt 获取请求成功时间](apihttp.md#RequestFinishedAt)
* [Request.ResponseStatusCode 获取相应状态码](apihttp.md#RequestResponseStatusCode)
* [Request.ResponseBody 获取响应正文](apihttp.md#RequestResponseBody)
* [Request.ResponseHeader 获取响应头](apihttp.md#RequestResponseHeader)
* [Request.ResponseHeaderValues 获取响应头全部值](apihttp.md#RequestResponseHeaderValues)
* [Request.ResponseHeaderFields 获取响应头全部字段](apihttp.md#RequestResponseHeaderFields)
