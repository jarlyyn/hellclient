# 输出内容接口

## GetLinesInBufferCount

获取历史行数


对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetLinesInBufferCount

### 原型
```
GetLinesInBufferCount() int
```

### 描述

获取当前游戏输出行数

正常情况下永远返回客户端支持的最大行数

### 代码范例

Javascript
```
world.note(world.GetLinesInBufferCount());
```

Lua
```
Note(GetLinesInBufferCount())
```

### 返回值

输出的行数

## DeleteOutput

废弃

### 返回值

空

## DeleteLines

删除历史内容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteLines


### 原型
```
DeleteLines(count int)
```

### 描述

删除游戏历史输出里指定行数的输出

* count 需要删除的行数

如果count大于行数，则按行数计算

如果count小于等于0,则直接返回

### 范例代码

Javascript
```
world.DeleteLines(10);
```

Lua
```
DeleteLines(10)
```

### 返回值

无

## GetLineCount

获取接受总行数

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetLineCount

### 原型
```
GetLineCount() int
```

### 描述
获取运行后游戏加载过的总行数

### 代码范例
Javascript
```
world.note(world.GetLineCount());
```

Lua
```
Note(GetLineCount())
```

### 运行结果

总行数

## GetRecentLines

获取近期内容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetRecentLines

### 原型
```
GetRecentLines(count int) string
```

### 描述

获取最近的指定行内容

* count 最近行数

如果count大于100,则取最后100行

返回的内容会以 \n 拼接

### 代码范例

Javascript
```
world.Note (world.GetRecentLines (10))
```

Lua
```
Note (GetRecentLines (10))
```

### 返回结果

字符串