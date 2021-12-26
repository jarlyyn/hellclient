# 文件处理

## ReadFile

读取脚本文件

### 原型
```
ReadFile(name string) string
```

### 描述

读取给定文件的文件内容

* name 相对于脚本文件目录的的文件名


如果文件位置在授权范围外，将会报错

其他的文件操作错误也会报错

### 代码范例

Javascript
```
var data=world.ReadFile("mydata.txt")
```

Lua
```
local data=ReadFile("mydata.txt")
```

### 返回值

文件内容


## ReadLines

读取脚本文件并分行

### 原型
```
ReadLines(name string) []string
```

### 描述

读取给定文件的文件内容，并以\n为分割符分割为字符串列表

* name 相对于脚本文件目录的的文件名


如果文件位置在授权范围外，将会报错

其他的文件操作错误也会报错

### 代码范例
Javascript
```
var data=world.ReadLines("mydata.txt")
```

Lua
```
local data=ReadLines("mydata.txt")
```

### 返回值

字符串列表

## HasHomeFile

检查用户文件

### 原型
```
HasHomeFile(name string) bool
```

### 描述

判断指定的用户文件是否存在

* name 以用户脚本根目录为基准的文件位置

如果文件在家目录范围外，会报告错误
  
如果有任何其他错误，认为文件不存在

### 代码范例

Javascript
```
var exists=world.HasHomeFile("myfile")
```

Lua
```
local exists=HasHomeFile("myfile")
```

### 返回值

布尔值

## ReadHomeFile

读取用户文件

### 原型
```
ReadHomeFile(name string) string
```

### 描述

读取给定用户文件的文件内容

* name 以用户脚本根目录为基准的文件位置


如果文件位置在授权范围外，将会报错

其他的文件操作错误也会报错

### 代码范例

Javascript
```
var data=world.ReadHomeFile("mydata.txt")
```

Lua
```
local data=ReadHomeFile("mydata.txt")
```

### 返回值

文件内容

## WriteHomeFile

写入用户文件

### 原型
```
WriteHomeFile(name string, body []byte)
```

### 描述

将数据写入给定的用户文件内

* name 以用户脚本根目录为基准的文件位置
* body 原始数据

如果文件位置在授权范围外，将会报错

其他的文件操作错误也会报错

### 代码范例

Javascript
```
world.WriteHomeFile("testfile","testdata")
```

Lua
```
WriteHomeFile("testfile","testdata")
```

### 返回值

无