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

不论是否被屏蔽，都会被获取

只获取收到的数据，不包括Note和系统信息

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

## GetLineInfo

获取指定行信息

部分兼容mushclient

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetLineInfo

### 原型
```
GetLineInfo(linenumber int, infotype int)
```

### 描述

获取指定行信息

* linenumner 行号，当前行为1,可以通过GetLinesInBufferCount获取最大行数
* infotype 要获取的信息类型

可用的信息类型包括部分兼容mushclient

* 1: 行文字
* 2: 文字长度
* 3: 是否是新行(结尾是\n)
* 4: 是否是world.Note显示的内容
* 5: 是否是用户输入
* 6: 是否进入日志(废弃)
* 7: false
* 8: false
* 9: 接受到行的时间
* 10: 实际行id(字符串)
* 11: 行内样式数量

### Lua注意事项

infotype为空时，将会传回一个Table,包含以下值

* text:     行文字                                 
* length:   文字长度                               
* newline:  是否是新行(结尾是\n)                 
* note:     是否是world.Note显示的内容                          
* user:     是否是用户输入                           
* log:      是否进入日志(废弃)                           
* bookmark: false                         
* hr:       false              
* time:     时间戳                        
* timestr:  文字格式日期时间
* line:     实际行id(字符串)
* styles:   行内样式数量

### 代码范例

Javascript
```
for (line = total_lines - 10; line <= total_lines; line++)
  {
  world.note ("Line " + line + " = " + world.GetLineInfo (line, 1));
  world.tell ("Received " + world.GetLineInfo (line, 9));
  world.note (" - Style runs = " + world.GetLineInfo (line, 11));
  }
```

Lua
```
for line = total_lines - 10, total_lines do
  Note ("Line ", line, " = ", GetLineInfo (line, 1))
  Tell ("Received ", GetLineInfo (line, 9))
  Note (" - Style runs = ", GetLineInfo (line, 11))
  end

```

### 返回值

见描述