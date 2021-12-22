# 界面接口

## FlashIcon

废弃

## 返回值
无

## SetStatus

设置状态文本

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SetStatus

### 原型

```
SetStatus(text string)
```

### 描述

将状态栏设置为指定文本

* text 字符串

### Lua注意事项

多个参数会以空格拼接发送

### 代码范例

Javascript
```
world.SetStatus("Current HP = " + world.GetVariable("hp"));
```

Lua
```
SetStatus ("Current HP = ", GetVariable("hp"))
```

### 返回值

空

## DeleteCommandHistory

删除命令记录

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteCommandHistory

### 原型
```
DeleteCommandHistory()
```

### 描述

清楚当前游戏的命令历史

### 代码范例

Javascript
```
world.DeleteCommandHistory();
```

Lua
```
DeleteCommandHistory()
```

### 返回值

空

## Info

追加信息文本

不完全兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Info

### 原型

```
Info(text string)
```

### 描述

将给定文本追加到信息栏

和Mushclient不同，Info命令没有独立的显示区域，使用和SetStatus同一块显示区域,每个游戏独立

### Lua注意事项

多个参数会以空格拼接发送

### 代码范例

Javascript
```
world.Info ("You are now connected");
```

Lua
```
Info ("You are now connected")
```

### 返回值

无

## InfoClear

清除信息文本

不完全兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=InfoClear

### 原型

```
InfoClear()
```

### 描述

清楚信息文本状态栏

和Mushclient不同，InfoClear命令没有独立的显示区域，使用和SetStatus同一块显示区域,每个游戏独立

### 代码范例

Javascript
```
world.Info ("You are now connected");
```

Lua
```
Info ("You are now connected")
```

### 返回值

无

## GetAlphaOption

获取选项

不完全兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetAlphaOption

### 原型
```
GetAlphaOption(name string) string
```

### 描述

返回当前游戏指定的选项

* name 选项名

选项名的可选值为

* "name" 游戏名
* "id" 游戏 id
* "command_stack_character" 命令分割符
* "script_prefix" 脚本命令前缀

如果获取其他选项会报一个错误

### 代码范例

Javascript
```
/world.note(world.getalphaoption ("name"));
```

Lua
```
world.note(world.getalphaoption ("name"))
```

### 返回值

如描述

## SetAlphaOption

设置选项

不完全兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SetAlphaOption

### 原型
```
SetAlphaOption(name string, value string) int
```

### 描述

设置当前游戏指定的选项

* name 选项名
* value 选项值

选项名的可选值为

* "name" 游戏名

如果设置其他选项会报一个错误

### 代码范例

Javascript
```
world.setalphaoption ("name", "Gandalf");
```

Lua
```
SetAlphaOption ("name", "Gandalf")
```

### 返回值

eOK

## GetInfo

获取信息

不完全兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SetAlphaOption

### 原型
```
GetInfo(infotype int) string
```

### 描述

获取当前游戏的部分信息

可选infotype为

* 1 服务器地址
* 2 游戏名
* 8 空字符串
* 28 脚本类型
* 35 脚本ID
* 36 执行脚本前缀
* 40 游戏ID.log
* 51 游戏ID.log
* 53 当前游戏信息栏
* 54 游戏ID.toml
* 55 游戏id
* 56 固定字符串 "hellclient"
* 57 固定字符串"./"
* 58 固定字符串"./"
* 59 固定字符串"./"
* 64 固定字符串"./"
* 66 固定字符串"./"
* 67 固定字符串"./"
* 68 固定字符串"./"

如果获取其他选项，会报错误

### 代码范例

Javascript
```
info = world.GetInfo (i);
```

Lua
```
info = GetInfo(i)
```

### 返回值

如描述所示

## GetGlobalOption

获取全局选项

不完全兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetGlobalOption

### 原型
```
GetGlobalOption(optionname string) string
```

### 描述

已废弃，获取全局选项

*optionname 选项名，选项名为TimerInterval 返回0,否者返回空字符串

### 范例代码

Javascript
```
/world.Note (world.GetGlobalOption ("TimerInterval"));
```

Lua
```
Note (GetGlobalOption ("TimerInterval"))
```

### 返回值

如描述