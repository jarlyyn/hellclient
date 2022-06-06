# 节拍限流

## 术语

节拍限流类似于mushclient的speedwalk,但机制有区别

节拍限流机制是在一个指定时间周期内，不超过最大限制的发送命令。如果周期内发送的命令到达限制，则将命令加入队列，直到最近一个周期内发送的内容不到beats

比如，周期为10秒的话，最大发送10个命令的话

* 第1秒尝试发送5个命令，全部发送成功
* 第2秒尝试发送 8个命令，5个发送成功，3个进入队列
* 第5秒尝试发送3个命令，全部进入队列
* 第11秒时，第1秒发送的5个命令已经超过一个周期，从队列中发送前5个命令，队列中还有1个命令
* 第15秒时，第2秒发送的命令全都超过一个周期，发送队列内剩余的1个命令，队列为空

注:出于兼容性考虑,Metronome的方法都可以使用全小写的方式调用

### tick
发送限制的周期

### beats
每个周期最多发送的命令数

### sent

最近一个周期内已经发送的内容

### queue
待发送的队列

### interval
再次发送队列的检查间隔

无法直接发送的命令将进入队列，并以interval为间隔检查是否可以再次发送

## Metronome.GetBeats

获取发送限制

### 原型

```
Beats() int
```

### 描述

返回限流器当前的发送限制

注意，限流值可以被设置为负值，但计算时最小值为1

### 范例代码

Javascript
```
world.Note(Metronome.GetBeats())
```

Lua
```
Note(Metronome:GetBeats())
```

### 返回值

设置的限制值

## Metronome.SetBeats

设置发送限制

### 原型
```
SetBeats(beats int)
```

### 描述

设置限流器的限制值。

注意，限流值可以被设置为负值，但计算时最小值为1

### 代码范例

Javascript
```
Metronome.SetBeats(15)
```

Lua
```
Metronome:SetBeats(15)
```

### 返回值

无

## Metronome.Reset

重置限流器

### 原型
```
Reset()
```

### 功能描述

重置限流器，清除所有的已发送记录，重新开始计算发送限制并立刻发送指令

### 代码范例

Javascript
```
Metronome.Reset()
```

Lua
```
Metronome:Reset()
```

### 返回值

无

## Metronome.GetSpace

获取剩余空间

### 原型

```
Space() int
```

### 描述

返回当前周期内还能输出多少命令

### 范例代码
Javascript
```
world.Note(Metronome.GetSpace())
```

Lua
```
Note(Metronome:GetSpace())
```

### 返回值

整数

## Metronome.GetQueue

获取队列

### 原型
```
Queue() []string
```

### 描述

返回目前限流器中待发送的命令

### 代码范例

Javascript
```
var queue=Metronome.GetQueue()
queue.forEach(function (cmd) {
    world.Note(cmd)
})
```

Lua
```
local queue=Metronome:GetQueue()
for k, cmd in pairs(queue) do
    Note(cmd)
end
```
### 返回值

命令速组

## Metronome.Discard

放弃队列

### 原型
```
Discard(force bool) bool
```

### 描述

销毁队列中的命令

* force 是否强制销毁。为false则不销毁队列中被锁定的内容

### 代码范例

Javascript
```
world.Note(Metronome.Discard(true))
```

Lua
```
Note(Metronome:Discard(true))
```

### 返回值

销毁是否成功

## Metronome.LockQueue

锁定队列

### 原型
```
LockQueue()
```

### 描述

将当前队列中的命令都标记为锁定状态

被锁定的命令无法通过 Metronome:discard(false) 清除

### 代码范例

Javascript
```
Metronome.LockQueue()
```

Lua
```
Metronome:LockQueue
```

### 返回值

无

## Metronome.Full

填充完整周期

### 原型
```
Full()
```

### 描述

在限流器里填充一个完整周期，确保当前队列输出后一个周期内不会发送命令

应该与Metronome.Discard 配合使用

### 代码范例

Javascript
```
Metronome.Full()
```

Lua
```
Metronome:Full()
```

### 返回值

无

## Metronome.FullTick

填充当前周期

### 原型

```
FullTick()
```

### 描述

填充当前周期，队列为空时确保下一周期才会发送命令

应该与Metronome.discard 配合使用

### 代码范例

Javascript
```
Metronome.FullTick()
```

Lua
```
Metronome:FullTick()
```

### 返回值

无

## Metronome.GetInterval

获取发送间隔

### 原型

```
Interval() int
```

### 描述

获取限流器的发送间隔，以毫秒(1/1000秒)为单位

注意，发送间隔可以任意设置，但小于等于0时按50计算


### 代码范例

Javascript
```
world.Note(Metronome.GetInterval())
```

Lua
```
Note(Metronome:GetInterval())
```

### 返回值

以毫秒(1/1000秒)为单位的整数

## Metronome.SetInterval

设置发送间隔

### 原型
```
SetInterval(interval int)
```

### 描述

设置限流器的发送间隔，以毫秒(1/1000秒)为单位

注意，发送间隔可以任意设置，但小于等于0时按50计算

### 代码范例

Javascript
```
Metronome.SetInterval(50)
```

Lua
```
Metronome:SetInterval(50)
```

### 返回值

空

## Metronome.GetTick

获取限流周期

### 原型
```
Tick() int
```

### 描述

获取限流器的限流周期，以毫秒(1/1000秒)为单位

注意，限流周期可以任意设置，但小于等于0时按1000计算


### 代码范例

Javascript
```
world.Note(Metronome.GetTick())
```

Lua
```
Note(Metronome:GetTick())
```

### 返回值

以毫秒(1/1000秒)为单位的整数

## Metronome.settick

设置限流周期

### 原型
```
SetTick(tick int)
```

### 描述

设置限流器的限流周期，以毫秒(1/1000秒)为单位

注意，限流周期可以任意设置，但小于等于0时按1000计算

### 代码范例

Javascript
```
Metronome.SetTick(50)
```

Lua
```
Metronome:SetTick(50)
```

### 返回值

空

## Metronome.Push

推送命令

### 原型

```
Push(cmds []string,grouped bool,echo bool)
```

### 描述

将给到的命令推送到限流器，尝试发送

* cmds 命令列表
* grouped 按组发送(当剩余空间不足时，按组发送的命令会到下一周期才发送)
* echo 是否回显

### 代码范例

Javascript
```
Metronome.Push(["w","n","s"],true,true)
```

Lua
```
Metronome:Push({"w","n","s"},true,true)
```

### 返回值

无