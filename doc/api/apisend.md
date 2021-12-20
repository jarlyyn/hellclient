# 发送接口

## print

打印

## 原型

```
print(cmd string...)
```

## 描述

以空格风格显示所有传入的参数

### 代码范例

Javascript
```
world.print("a","b","c")
```

Lua
```
print("a","b","c")
```

### 返回值

无


## Note

显示

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Note


### 原型

```
Note(cmd string)
```
### 描述

打印显示传入的参数

### Lua注意事项

多个参数会以空格拼接发送

### 代码范例

Javascript
```
world.note("This monster is worth avoiding!");
```

Lua
```
Note ("This monster is worth avoiding!")
```

## SendImmediate

立即发送

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SendImmediate

### 原型
```
SendImmediate(message string) int
```
### 描述

立即发送，不进入队列

### Lua注意事项

多个参数会以空格拼接发送
### 范例代码
Javascript
```
world.SendImmediate("go north");
```
Lua
```
SendImmediate ("go north")
```

### 返回值

eOK:发送成功

## Send

发送

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Send

### 原型
```
Send(message string) int
```
### 描述

立即发送，不进入队列(同SendImmediate)

### Lua注意事项

多个参数会以空格拼接发送
### 范例代码
Javascript
```
world.Send("go north");
```
Lua
```
Send ("go north")
```

### 返回值

eOK:发送成功

## SendNoEcho

静默发送

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SendNoEcho

### 原型
```
SendNoEcho(message string) int
```
### 描述

静默发送，不显示在输出里

### Lua注意事项

多个参数会以空格拼接发送
### 范例代码
Javascript
```
world.SendNoEcho("go north");
```
Lua
```
SendNoEcho ("go north")
```

### 返回值

eOK:发送成功

## SendSpecial

高级发送

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SendSpecial

### 原型
```
SendSpecial(message string, echo bool, queue bool, log bool, history bool)
```
### 描述

高级发送

* message 发送的内容
* echo 是否回显
* queue 是否进入队列
* log 废弃
* history 是否进入历史记录

### 范例代码
Javascript
```
world.SendSpecial ("go north", true, false, false, true);
```
Lua
```
SendSpecial ("go north", true, false, true, false)
```

### 返回值

eOK:发送成功

## Execute

执行

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Execute

### 原型
```
Execute(message string) int
```

### 描述

执行给道的命令，效果同在输入框输入

* message 需要执行的命令

### 代码范例

Javascript
```
world.Execute ("north");  // normal command
world.Execute ("/world.Debug ("colours ") ");  // execute a script
```
Lua
```
Execute ("north")  -- normal command
Execute ("/Debug ('colours') ")  -- execute a script
```

### 返回值

* eOK: 返回成功

## Queue

队列发送

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Queue

### 原型
```
Queue(message string, echo bool) int
```

### 描述

执行给道的命令，效果同在输入框输入

* message 需要执行的命令
* echo 是否回显

### Lua注意事项

echo为空时，值为True
### 代码范例

Javascript
```
world.queue("n", true);
```
Lua
```
Queue ("4n", true)
```
### 返回值

* eOK: 返回成功

## DiscardQueue

取消队列

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DiscardQueue

### 原型
```
DiscardQueue(force bool) int
```
### 描述
放弃队列中所有未发送的队列

*force 是否强制将已经锁定的命令也清空

### 代码范例

Javascript
```
world.discardqueue(false);
```
Lua
```
DiscardQueue(false)
```
### 返回值

被清除的命令数量

## LockQueue

锁定队列

### 原型
```
LockQueue()
```

### 描述
锁定队列，防止队列中的当前命令被清除。

被锁定的命令只有通过force参数才能清除掉。

### 代码范例

Javascript
```
world.lockqueue()
```

Lua
```
LockQueue()
```

### 返回值

无

## GetQueue

获取队列内容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetQueue

### 原型

```
GetQueue() []string
```

### 描述
获取当前队列里的未发送命令

### 范例代码

Javascript
```
commandList = world.GetQueue();

 if (commandList) 
   for (i = 0; i < commandList.length; i++)
       world.note (commandList [i]);

```

Lua
```
for k, v in pairs (GetQueue()) do 
  Note (v) 
end

```

### 返回值

字符串列表