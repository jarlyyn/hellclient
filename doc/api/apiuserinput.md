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
* 选择操作
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

## Userinput.newlist

新建列表

### 原型

```
NewList(title string, intro string, withfilter bool)*List
```

### 描述

创建一个列表

列表需要发布(publish)后才会显示
* title 列表标题
* intro 列表介绍
* withfilter 列表是否带过滤器
### 代码范例

Javascript
```
var list=Userinput.newlist("test list","A test list",true)
```

Lua
```
local list=Userinput:newlist("test list","A test list",true)
```

### 返回值

新列表对象

## Userinput.newdatagrid

新建数据表格

### 原型
```
NewDatagrid(title string, intro string) *Datagrid
```

### 描述

创建一个数据表格

表格需要发布(publish)后才会显示
* title 表格标题
* intro 表格介绍
### 代码范例

Javascript
```
var datagrid=Userinput.newdatagrid("test data grid","A test data grid")
```

Lua
```
local datagrid=Userinput:newdatagrid("test data grid","A test data grid")
```

### 返回值

新数据表格对象

## Userinput.newvisualprompt

新建可视化输入

### 原型
```
NewVisualPrompt(title string, intro string, source string) *VisualPrompt
```

### 描述

创建一个可视化输入

可视化输入需要发布(publish)后才会显示
* title 表格标题
* intro 表格介绍
* source 媒体来源 媒体来源应该在可信域名下
### 代码范例

Javascript
```
var vp=Userinput.newvisualprompt("test data visualprompt","A test data grid","http://127.0.0.1/test.png")
```

Lua
```
local vp=Userinput:newvisualprompt("test data visualprompt","A test data grid","http://127.0.0.1/test.png)
```

### 返回值

新可视化输入对象

## Userinput.hideall

隐藏界面UI

### 描述

隐藏界面的UI

由于confirm,Alert,promot属于弹框，所以调用该方法后不保证会关闭，这个取决于实际前端的实现

List,Datagrid,VirualPrompt的操作界面会进行隐藏

### 代码范例
Javascript
```
Userinput.hideall()
```

Lua
```
Userinput.hideall()
```

### 返回值

无
## List.append

列表追加

## 原型

```
Append(key string, value string)
```

### 描述

向列表里加入一对键值

* key 选择后返回给回调的key
* value 显示的值

### 范例代码

Javascript
```
list.append("sword","长剑")
```

Lua
```
list.append("sword","长剑")
```

### 返回值

无

## List.setmutli

列表设置多选

### 原型

```
SetMutli(mutli bool)
```

### 描述

设置列表是否为多选列表

多选列表，选中的值将以字符串数组序列化后发给回调函数

如
```
["key1","key2","key3"]
```

### 代码范例

Javascript
```
List.setmutli(true)
```

Lua
```
List.setmutli(true)
```

### 返回值

无

## List.setvalue

列表设置值

### 原型
```
SetValues(values []string)
```

### 描述

设置列表的初始值为给到的字符串列表

只对多选列表有效

### 代码范例

Javascript
```
List.setvalue(["key1","key2","key3"])
```

Lua
```
List.setvalue({"key1","key2","key3"})
```

### 返回值

无

## List.publish

列表发布

### 代码原型
```
Publish(script string) string
```

### 代码描述

发布列表

* script 接受数据的回调

### 回调函数

* code 取消为0,确认非0
* data 单选为用户选择的对象key,多选列表为JSON序列化后的用户选自的对象key的字符串列表

### 范例

Javascript
```
List.publish("handlelist")
```

Lua
```
List.publish("handlelist")
```

### 返回值

唯一id

## Datagrid.append

数据表格追加数据

### 原型
```
Append(key string, value string)
```

### 描述

向数据表格内追加一行数据

### 范例代码
Javascript
```
Datagrid.append("key","value")
```

Lua
```
Datagrid.append("key","value")
```

### 返回值

无

## Datagrid.resetitems

数据表格重置元素

### 原型
```
ResetItems() 
```

### 描述

重置清空表单的所有元素

### 代码范例

Javascript
```
Datagrid.resetitems()
```

Lua
```
Datagrid.resetitems()
```

### 返回值

空
## Datagrid.setfilter

数据表格设置过滤值

### 原型
```
SetFilter(filter string) 
```

### 描述
设置数据表格的初始过滤器信息

### 代码范例

Javascript
```
Datagrid.setfilter("filter)
```

Lua
```
Datagrid.setfilter("filter)
```

### 返回值

无

## Datagrid.getfilter

数据表格获取过滤值

## 原型
```
GetFilter() string
```

### 描述

返回数据表格当前的过滤器设置

### 代码范例

Javascript
```
world.Note(Datagrid.getfilter())
```

Lua
```
Note(Datagrid.getfilter())
```
### 返回值

字符串

## Datagrid.setmaxpage

数据表格设置最大页数

### 原型
```
SetMaxPage(page int)
```

### 描述
设置数据表格的最大分页数

### 范例代码
Javascript
```
Datagrid.setmaxpage(10)
```

Lua
```
Datagrid.setmaxpage(10)
```

### 返回值

无

## Datagrid.setpage

数据表格设置当前页

### 原型
```
SetPage(page int)
```

### 描述

设置数据表格的当前页

### 范例代码
Javascript
```
Datagrid.setpage(3)
```

Lua
```
Datagrid.setpage(3)
```

### 返回值

无

## Datagrid.getpage

数据表格获取当前页

### 原型
```
GetPage() int
```

### 描述

获取数据表格的当前页

### 代码范例

Javascript
```
world.Note(Datagrid.getpage)
```

Lua
```
Note(Datagrid.getpage)
```
### 返回值

整数的当前页

## Datagrid.setoncreate

数据表格设置创建回调

### 原型

```
SetOnCreate(oncreate string)
```

### 描述

设置数据表格的创建按钮回调

只有设置了回调，创建按钮才会出现

### 代码范例

Javascript
```
Datagrid.setonupdate("handleupdate")
```

Lua
```
Datagrid.setonupdate("handleupdate")
```

### 返回值

无

## Datagrid.setonupdate

数据表格设置更新回调

### 原型

```
SetOnUpdate(onupdate string)
```

### 描述

设置数据表格的更新按钮回调

只有设置了回调，更新按钮才会出现

### 回调函数

* data 需要操作的元素的key

### 代码范例

Javascript
```
Datagrid.setonupdate("handleupdate")
```

Lua
```
Datagrid.setonupdate("handleupdate")
```

### 返回值

无

## Datagrid.setonview

数据表格设置查看回调

### 原型

```
SetOnView(onview string)
```

### 描述

设置数据表格的查看按钮回调

只有设置了回调，查看按钮才会出现

### 回调函数

* data 需要操作的元素的key

### 代码范例

Javascript
```
Datagrid.setonview("handleview")
```

Lua
```
Datagrid.setonview("handleview")
```

### 返回值

无

## Datagrid.setonselect

数据表格设置选择回调

### 原型

```
SetOnSelect(onselect string)
```

### 描述

设置数据表格的选择按钮回调

只有设置了回调，查看按钮才会出现

### 回调函数

* data 需要操作的元素的key

### 代码范例

Javascript
```
Datagrid.setonselect("handleselect")
```

Lua
```
Datagrid.setonselect("handleselect")
```

### 返回值

无

## Datagrid.setondelete

数据表格设置删除回调

### 原型

```
SetOnDelete(ondelete string)
```

### 描述

设置数据表格的删除按钮回调

只有设置了回调，删除按钮才会出现

### 回调函数

* data 需要操作的元素的key

### 代码范例

Javascript
```
Datagrid.setondelete("handledelete")
```

Lua
```
Datagrid.setondelete("handledelete")
```

### 返回值

无

## Datagrid.setonfilter

数据表格设置过滤回调

### 原型

```
SetOnFilter(onfilter string)
```

### 描述

设置数据表格的过滤回调

只有设置了回调，过滤框才会出现

### 回调函数

* data 过滤值

### 代码范例

Javascript
```
Datagrid.setonfilter("handlefilter")
```

Lua
```
Datagrid.setonfilter("handlefilter")
```

### 返回值

无

## Datagrid.setonpage

数据表格设置分页回调

### 原型

```
SetOnPage(onpage string)
```

### 描述

设置数据表格的分页回调

只有设置了回调，分页按钮才会出现

### 回调函数

* data 分页值

### 代码范例

Javascript
```
Datagrid.setonpage("handlepage")
```

Lua
```
Datagrid.setonpage("handlepage")
```

### 返回值

无

## Datagrid.hide

数据表格隐藏

### 原型
```
Hide()
```

### 描述

隐藏数据表格

### 代码范例

Javascript
```
Datagrid.hide()
```

Lua
```
Datagrid.hide()
```

### 返回值

无

## Datagrid.publish

数据表格发布

### 原型
```
Publish(script string) string
```

### 描述
发布数据表格

* script 为回调，一般用于处理关闭事件

### 代码范例

Javascript
```
Datagrid.publish("onclose")
```

Lua
```
Datagrid.publish("onclose")
```

### 返回值

唯一id

## VisualPrompt.setmediatype

可视化输入设置媒体类型

### 原型
```
SetMediaType(t string)
```

### 描述

设置可视化输入的媒体类型，暂时只支持image，output和text

默认为image

output需要将url设置为DumpOutput的返回值(JSON格式的Line数组)

### 代码范例

Javascript
```
VisualPrompt.setmediatype("image")
```

Lua
```
VisualPrompt.setmediatype("image")
```

### 返回值

无

## VisualPrompt.setportrait

可视化输入设置垂直模式

### 原型
```
SetPortrait(v bool)
```

### 描述

设置可视化输入是否应该以垂直模式显示媒体

### 范例代码

Javascript
```
VisualPrompt.setportrait(true)
```

Lua
```
VisualPrompt.setportrait(true)
```

### 返回值

无

### VisualPrompt.setrefreshcallback

可视化输入设置刷新回调

### 原型

```
SetRefreshCallback(callback string)
```

### 描述

设置刷新回调

只有设置了回调，才会出现刷新按钮

### 代码范例

Javascript
```
VisualPrompt.setrefreshcallback("handlerefresh")
```

Lua
```
VisualPrompt.setrefreshcallback("handlerefresh")
```

### 返回值

无

## VisualPrompt.publish

可视化输入发布

### 原型

```
Publish(script string) string
```

### 描述

发布可视化输入

script为接受用户提交的回调

### 回调函数

* code 取消为0,成功非0
* data 用户输入数据

### 代码范例

Javascript
```
VisualPrompt.publish("onvisualprompt")
```

Lua
```
VisualPrompt.publish("onvisualprompt")
```

### 返回值

唯一ID

## VisualPrompt.append

列表追加

## 原型

```
Append(key string, value string)
```

### 描述

向列表里加入一对键值

* key 选择后返回给回调的key
* value 显示的值

当VisualPrompt的列表不为空时，将现实类似List的列表选择，代替原有的输入框

### 范例代码

Javascript
```
VisualPrompt.append("sword","长剑")
```

Lua
```
VisualPrompt.append("sword","长剑")
```

### 返回值

无