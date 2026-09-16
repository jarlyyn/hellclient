# 底层接口

提供更底层的信息给脚本引擎，一般用于风险引擎绕过客户端机制直接从处理。

## LastAnsi

获取最后一行的Ansi字符串信息，unicode编码

版本 2.2026.08.31 后加入


### 原型

```
LastAnsi() string
```

### 描述

获取最后一行的Ansi字符串信息，unicode编码


### 代码范例

Javascript
```
world.LastAnsi();
```

Lua
```
LastAnsi()
```

### 返回值

包含Ansi信息的纯文本


## AddAnsi

插入Ansi行信息

版本 2.2026.08.31 后加入

### 原型

```
AddAnsi(string) 
```

### 描述

向游戏插入Ansi行信息，会参与后续触发，可以与OnLine配合修正mud输出

### 代码范例

Javascript
```
world.AddAnsi("text");
```

Lua
```
AddAnsi("text")
```

### 返回值

无


## GetScriptPath

获取脚本文件根目录

版本 2.2026.08.31 后加入

### 原型

```
GetScriptPath() string
```

### 描述

获取脚本位置，方便风险脚本引擎直接读取银盘文件

### 代码范例

Javascript
```
world.GetScriptPath();
```

Lua
```
GetScriptPath("text")
```

### 返回值

字符串


