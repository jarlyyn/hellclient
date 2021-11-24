# 计时器接口

[返回](api.md)

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