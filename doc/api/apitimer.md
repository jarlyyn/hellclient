# 计时器接口

[返回](readme.md)
## AddTimer

添加计时器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=AddTimer

### 原型
```
AddTimer(timerName string, hour int, minute int, second float64, responseText string, flags int, scriptName string) int
```

### 描述

添加一个脚本计时器

* timerName 计时器名
* hour 小时数
* minute 分钟数
* second 秒数
* responseText 发送文本
* flags 标识位
* scriptName 脚本名称

flags的可选值为

* eEnabled = 1; // 计时器是否有效
* eAtTime = 2; // 计时器是否为指定时间模式
* eOneShot = 4; // 是否是一次性及私企
* eTimerSpeedWalk = 8; //废弃
* eTimerNote = 16; //废弃
* eActiveWhenClosed = 32; // 断线时是否有效
* eReplace = 1024; // 是否替代同名计时器
* eTemporary = 16384; // 是否是临时计时器

Lua可使用timer_flag表

* timer_flag.Enabled = 1
* timer_flag.AtTime = 2
* timer_flag.OneShot = 4
* timer_flag.TimerSpeedWalk = 8
* timer_flag.TimerNote = 16
* timer_flag.ActiveWhenClosed = 32
* timer_flag.Replace = 1024
* timer_flag.Temporary = 16384

### 代码范例

Javascript:
```
world.addtimer("my_timer", 0, 0, 1, "go north", 5, "");
```

Lua:
```
AddTimer ("my_timer", 0, 0, 1.5, "go north", 
          timer_flag.Enabled + timer_flag.OneShot, "")
```

### 返回值

* eTimerAlreadyExists Timer已经存在
* eOK 添加成功

## DeleteTimer

删除计时器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteTimer

### 原型
```
DeleteTimer(name string) int
```
### 描述

删除指定名称的脚本计时器

* name 计时器名

### 代码范例

Javascript:

```
world.DeleteTimer("mytimer");
```

Lua范例

```
DeleteTimer("mytimer")
```

### 返回值

* eTimerNotFound 计时器未找到
* eOK 执行成功

## DeleteTimerGroup

删除计时器组

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteTimerGroup


### 原型
```
DeleteTimerGroup(group string) int
```

### 描述

按给到的分组删除计时器

只删除分组内的脚本计时器

### 代码范例

Javascript:
```
world.DeleteTimerGroup ("groupname");
```
Lua:
```
DeleteTimerGroup ("groupname")
```

### 返回值

删除的计时器的数量

## DeleteTemporaryTimers

删除临时计时器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteTemporaryTimers

### 原型

```
DeleteTemporaryTimers() int
```
### 描述

删除所有临时计数器

### 代码范例

Javascript:
```
world.DeleteTemporaryTimers();
```

Lua:
```
DeleteTemporaryTimers()
```
### 返回值

删除的计时器数量

## EnableTimer

激活计时器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=EnableTimer

### 原型
```
EnableTimer(name string, enabled bool) int
```

### 描述

激活脚本计时器

* name 计时器名
* enabled 是否激活计时器

### Lua注意事项

enabled为空时，值为True

### 代码范例

Javascript:
```
world.EnableTimer("heartbeat", true);
```

Lua:
```
EnableTimer("heartbeat", true)
```

### 返回值

* eTimerNotFound 计时器未找到
* eOK 执行成功

## EnableTimerGroup

激活计时器组

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=EnableTimerGroup

### 原型

```
EnableTimerGroup(group string, enabled bool) int
```

### 描述

按组激活计时器

* group 组名
* enabled 激活状态

### Lua注意事项

enabled为空时，值为True

### 代码范例

Javascript:

```
world.EnableTimerGroup ("groupname", 1);  // enable the group
world.EnableTimerGroup ("groupname", 0);  // disable the group
```

Lua:

```
EnableTimerGroup ("groupname", true)  -- enable the group
EnableTimerGroup ("groupname", false)  -- disable the group
```
### 返回值

组内计时器的数量

## GetTimerInfo

获取计时器信息

部分兼容mushclient

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetTimerInfo

### 原型
```
GetTimerInfo(name string, infotype int) (string, int)
```

### 描述

获取指定的脚本计时器的信息

* name 计时器名称
* infotype 信息类型

可用的infotype值

1: 小时值 (short)
2: 分钟值 (short)
3: 秒值 (short)
4: 发送内容 (string)
5: 脚本名 (string)
6: 是否激活 (boolean)
7: 是否是一次性计时器 (boolean)
8: 是否是指定时间点模式计时器，为False则为固定间隔(Every)模式 (boolean)
9: 废弃
10: 废弃
11: 废弃
12: 废弃
13: 废弃
14: 是否是临时计时器 (boolean)
15: 废弃
16: 废弃
17: 断线后是否激活 (boolean)
18: 废弃 (boolean)
19: 计时器组名 (string)
20: 发送到位置 (long)
21: 废弃
22: 名称 (string)
23: 屏蔽输出标识 (boolean)
24: 废弃
25: 废弃
26: 废弃

### 代码范例

Javascript:
```
world.note(world.GetTimerInfo ("my_timer", 2));
```

Lua:
```
Note(GetTimerInfo ("my_timer", 2))
```

### 返回值

* 成功获取的返回值
* 别名没找到返回空
* Infotype无效返回空

## GetTimerList

获取脚本计时器列表

不兼容mushclient

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetTimerList

### 原型
```
GetTimerList() []string
```
### 描述

返回脚本计时器名称列表

### 代码范例

Javascript:
```
timerlist = new VBArray(world.GetTimerList()).toArray();

if (timerlist)  // if not empty
 for (i = 0; i < timerlist.length; i++)
   world.note(timerlist [i]);
```

Lua:
```
tl = GetTimerList ()
if tl then
  for k, v in ipairs (tl) do 
    Note (v) 
  end  -- for
end -- if we have any timers
```

### 返回值

存有脚本计时器名的字符串列表

## GetTimerOption

获取计时器选项

部分兼容mushclisnt

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetTimerOption


### 原型
```
GetTimerOption(name string, option string) (string, int)
```

### 描述

返回给定名称的计时器的指定选项

* name 计时器名
* option 选项名

可用的option值

* "active_closed": y/n - 关闭时计时器是否有效
* "at_time": y/n - 固定时间点模式
* "enabled": y/n - 计时器是否有效
* "group": (string - 计时器分组)
* "hour": 小时数
* "minute": 分钟数
* "name": (string - name/label of alias)
* "offset_hour": 废弃
* "offset_minute": 废弃
* "offset_second": 废弃
* "omit_from_log": y/n - 废弃
* "omit_from_output": y/n - 是否屏蔽输出
* "one_shot": y/n - 是否是一次性计时器
* "script": (string - 调用的脚本名)
* "second": 秒数
* "send": (multi-line string - 发送内容)
* "send_to": 0 - 13 - 发送位置
* "user": 废弃
* "variable": (string - 发送到变量值)

布尔值会返回 0\(false\) 或 1\(true\)

### 代码范例

Javascript:
```
Note (world.GetTimerOption ("mytimer", "group"));
```

Lua:
```
Note (GetTimerOption ("mytimer", "group"))
```

### 返回值

如描述所列

## IsTimer

判断计时器是否存在

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=IsTimer

### 原型
```
IsTimer(name string) int
```

### 描述

检查指定的脚本计时器是否存在

### 代码范例

Javascript:
```
world.note(world.IsTimer("mytimer"));
```

Lua:
```
Note(IsTimer("mytimer"))
```

### 返回值

* eTimerNotFound 脚本计时器未找到
* eOK 脚本计时器存在

## ResetTimer

重置计时器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=ResetTimer

### 原型
```
ResetTimer(name string) int
```

### 描述

重置指定的计时器

未激活的计时器不会被重置

* name 计时器名

### 代码范例

Javascript:
```
world.ResetTimer("mytimer");
```

Lua:
```
ResetTimer("mytimer")
```

### 返回值

* eTimerNotFound 计时器未找到
* eOK 重置成功

## ResetTimers

重置全部计时器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=ResetTimers

### 原型
```
ResetTimers() {
```

### 描述

重置所有用户和脚本计时器

未激活的计时器不会被重置


### 代码范例

Javascript:
```
world.ResetTimers();
```

Lua:
```
ResetTimers()
```

### 返回值

无

## SetTimerOption

设置计时器选项

部分兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SetTimerOption


### 原型

```
SetTimerOption(name string, option string, value string) int
```

### 描述

设置指定的脚本计时器值

* name 计时器名称
* option 选项名称
* value 选项值

可用的option值：

* "active_closed": y/n - 关闭时计时器是否有效
* "at_time": y/n - 固定时间点模式
* "enabled": y/n - 计时器是否有效
* "group": (string - 计时器分组)
* "hour": 小时数
* "minute": 分钟数
* "name": (string - name/label of alias)
* "offset_hour": 废弃
* "offset_minute": 废弃
* "offset_second": 废弃
* "omit_from_log": y/n - 废弃
* "omit_from_output": y/n - 是否屏蔽输出
* "one_shot": y/n - 是否是一次性计时器
* "script": (string - 调用的脚本名)
* "second": 秒数
* "send": (multi-line string - 发送内容)
* "send_to": 0 - 13 - 发送位置
* "user": 废弃
* "variable": (string - 发送到变量值)

对于数字型的值，传入的字符串应该能转换为数字

对于布尔型的值，传入的值必须是
* "y", "Y", or "1" 为真值
* "n", "N", or "0" 为假值

### 代码范例

Javascript:
```
world.SetTimerOption ("mytimer", "minute", "5");
```

Lua:
```
SetTimerOption ("mytimer", "minute", "5")
```
### 返回值

* eTimerNotFound 计时器未找到
* eTimeInvalid 时间无效
* eOK 设置成功