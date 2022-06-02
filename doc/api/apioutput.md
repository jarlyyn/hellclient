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

如果行号或者infotype无效，返回空

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

## GetStyleInfo

获取样式信息

部分兼容mushclient

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetStyleInfo

### 原型
```
GetStyleInfo(linenumber int, style int, infotype int) (string, bool)
```

### 描述

获取指定行样式的信息

* linenumber 行号，从1开始
* style 样式编号 从1开始
* infotype 信息类型

可选的信息类型包括

* 1: 文字
* 2: 样式长度
* 3: 开始信息
* 4: 废弃
* 5: 废弃
* 6: 废弃
* 7: 废弃
* 8: 是否加粗
* 9: 似乎否有下划线
* 10: 是否闪烁
* 11: 是否反转
* 12: 废弃
* 13: 废弃
* 14: 前景色的RGB
* 15: 背景色的RGB

如果行号，样式编号，信息类型任何一个无效，返回空

### Lua注意事项

如果infotype为0,以table形式返回所有信息

如果样式编号为0,以table形式返回所有样式信息

为保持兼容性的设定，不是很建议这样使用

### 范例代码

Javascript
```
world.note (world.GetStyleInfo (100, 2, 14));
```

Lua
```
Note (GetStyleInfo (100, 2, 14))

```

### 返回值

见描述

## DumpOutput

导出输出
### 原型

```
DumpOutput(length int, offset int) string
```

### 描述

将屏幕指定行数的输出导出到字符串变量内。

与GetRecentLines的区别是取得的是JSON序列化后的数据，可以用于直接分析，也可以在Userinput的VisualPrompt中以output的Mediatype直接显示。

* length 整数,总计返回的行数,小于0当作0处理
* offset 整数,跳开多少行开始导出，小于0当作0处理

能获取的文字上限为100行

### 代码范例

Javascript
```
world.Note(world.DumpOutput(10,2))
```

Lua
```
Note(DumpOutput(10,2))
```

### 返回值

字符串化的Line数组

[Line结构](../struct/line.md)

范例\(格式化后\)为

```json
[
 {
  "Words": [
   {
    "Text": "目前的字符集是简体，请输入GB/BIG5改变字符集，或直接登录用户。",
    "Color": "",
    "Background": "",
    "Bold": false,
    "Underlined": false,
    "Blinking": false,
    "Inverse": false
   }
  ],
  "ID": "d1dsufeg15qrdu71tgt3kp",
  "Time": 1653893585,
  "Type": 2,
  "OmitFromLog": false,
  "OmitFromOutput": false,
  "Triggers": [
   "on_global"
  ],
  "CreatorType": "",
  "Creator": ""
 },
 {
  "Words": [
   {
    "Text": "请输入您的英文名字(",
    "Color": "",
    "Background": "",
    "Bold": false,
    "Underlined": false,
    "Blinking": false,
    "Inverse": false
   },
   {
    "Text": "忘记密码请输入「pass」",
    "Color": "Cyan",
    "Background": "",
    "Bold": false,
    "Underlined": false,
    "Blinking": false,
    "Inverse": false
   },
   {
    "Text": ")：",
    "Color": "",
    "Background": "",
    "Bold": false,
    "Underlined": false,
    "Blinking": false,
    "Inverse": false
   }
  ],
  "ID": "d1dsufeg16eibl71tgt3kq",
  "Time": 1653893585,
  "Type": 2,
  "OmitFromLog": false,
  "OmitFromOutput": false,
  "Triggers": [
   "on_global"
  ],
  "CreatorType": "",
  "Creator": ""
 }
]
```

## ConcatOutput

连接输出

### 原型

```go
ConcatOutput(output1 string, output2 string) string
```

### 描述

解析两个输出数据，并拼接起来

* output1 开头部分的输出
* output2 结束部分的输出


### 代码范例

Javascript
```
world.Note(world.ConcatOutput(world.DumpOutput(10,2),world.DumpOutput(10,20)))
```

Lua
```
Note(ConcatOutput(DumpOutput(10,2),DumpOutput(10,20)))
```

### 返回值

字符串化的Line数组

## SliceOutput

切割输出

### 原型

```go
SliceOutput(output string, start int, end int) string
```

### 描述

从原输出的指定位置切割出新输出

* output 原输出
* start 从0开始序号，超过原输出序号最大值则默认为原序号最大值
* end 结束位置，小于等于0或者大于输出长度，则切割所有剩下的输出。如果比start小，则视为和start一样大


### 代码范例

Javascript
```
world.Note(world.SliceOutput(output,1,2))
```

Lua
```
Note(SliceOutput(output,1,2))
```

### 返回值

字符串化的Line数组

## OutputToText

输出转文字

### 原型

```go
OutputToText(output string) string
```

### 描述

将原输出的正文部分去掉格式后转成文字

* output 原输出


### 代码范例

Javascript
```
world.Note(world.OutputToText(output))
```

Lua
```
Note(OutputToText(output))
```

### 返回值

回车\(\\n\)分割的纯文本

## FormatOutput

格式化输出

### 原型

```go
FormatOutput(output string) string
```

### 描述

将输出格式化为易于阅读的形式

### 代码范例

Javascript
```
world.Note(world.FormatOutput(output))
```

Lua
```
Note(FormatOutput(output))
```

### 返回值

字符串化的Line数组


## Simulate

模拟发送文字

### 原型

```go
Simulate(lines string)
```

### 描述

将文字模拟为收到的无格式Mud信息

* lines 回车分割的文字

### 代码范例

Javascript
```
world.Simulate(output)
```

Lua
```
Simulate(output)
```

### 返回值

无

## SimulateOutput

模拟发送输出

### 原型

```go
Simulate(output string)
```

### 描述

将输出模拟为收到Mud信息

* output 输出

可以与 DumpOutput,WriteHomeFile,ReadHomeFile配合，进行输出的录制和重放，便于开发和维护机器

### 代码范例

Javascript
```
world.SimulateOutput(output)
```

Lua
```
SimulateOutput(output)
```

### 返回值

无