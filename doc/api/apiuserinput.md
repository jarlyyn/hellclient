# 用户输入

## 术语

用户输入指脚本发起请求用户输入或选择数据的界面交互

大部分数据输入支持回调函数，通过回调函数来相应用户的操作

### List

列表

列表指现实一个表格，让用户选择一个或多个元素。

列表可以支持过滤器，帮助用户快速定位元素

### Datagrid

数据表格

数据表格可以维护一个表格，并对表格提供基本的维护功能，包括、

* 查看操作
* 编辑操作
* 删除操作
* 创建操作
* 分页操作
* 筛选操作
* 关闭操作

### VisualPrompt

可视化输入

可视化输入一般用于识别图片或验证码，并输入内容

## Userinput.prompt

用户输入框

### 原型
```
Prompt (script string, title string, intro string, value string) string
```

### 描述

弹窗请求用户输入

* script 回调脚本
* title 输入框标题
* intro 输入框描述
* value 默认值

### 回调函数

* code 取消为0,成功非0
* data 用户输入数据

### 代码范例

Javascript
```
Userinput.prompt("handleprompt","Input","what's your name","")
```

Lua
```
Userinput:prompt("handleprompt","Input","what's your name","")
```

### 返回值

唯一id

## Userinput.confirm

用户确认框

### 原型
```
Confirm(script string, title string, intro string) string
```
### 回调函数

* code 取消为0,确认非0

### 描述

弹窗请求用户输入是或否

* script 回调脚本
* title 输入框标题
* intro 输入框描述

### 代码范例

Javascript
```
Userinput.confirm("handleprompt","Input","what's your name")
```

Lua
```
Userinput:confirm("handleprompt","Input","what's your name")
```

### 返回值

唯一id

## Userinput.alert

用户提示框

### 原型
```
Alert(script string, title string, intro string) string
```
### 回调函数

* code 取消为0,确认非0

### 描述

弹窗请求用户输入是或否

* script 回调脚本
* title 输入框标题
* intro 输入框描述

### 代码范例

Javascript
```
Userinput.alert("handleprompt","Input","what's your name")
```

Lua
```
Userinput:alert("handleprompt","Input","what's your name")
```

### 返回值

唯一id