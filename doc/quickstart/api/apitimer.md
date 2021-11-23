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