# 用户输入

## 术语

用户输入指脚本发起请求用户输入或选择数据的界面交互

大部分数据输入支持回调函数，通过回调函数来相应用户的操作

注:出于兼容性考虑,Userinput,List,VisualPrompt,Datagrid的方法都可以使用全小写的方式调用


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

## Userinput.Prompt

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
Userinput.Prompt("handleprompt","Input","what's your name","")
```

Lua
```
Userinput:Prompt("handleprompt","Input","what's your name","")
```

### 返回值

唯一id

## Userinput.Confirm

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
Userinput.Confirm("handleprompt","Input","what's your name")
```

Lua
```
Userinput:Confirm("handleprompt","Input","what's your name")
```

### 返回值

唯一id

## Userinput.Alert

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
Userinput.Alert("handleprompt","Input","what's your name")
```

Lua
```
Userinput:Alert("handleprompt","Input","what's your name")
```

### 返回值

唯一id

## Userinput.Popup

弹出提醒框

### 原型
```
Popup (script string, title string, intro string, popuptype string) string
```

### 描述

弹窗请求用户输入

* script 回调脚本
* title 标题
* intro 描述
* popuptype  弹窗类型，可选值为 success warning info error,默认无

### 回调函数

* code 成功非0

### 代码范例

Javascript
```
Userinput.Popup("handlepopup","Popuptitle","Popupcontent","success")
```

Lua
```
Userinput:Popup("handlepopup","Popuptitle","Popupcontent","success")
```

### 返回值

唯一id

## Userinput.Note

显示文本

### 原型
```
Note (script string, title string, body string, type string) string
```

### 描述

弹窗显示

* script 回调脚本
* title 标题
* body 正文，类型取决于type
* type  正文类型，可选值为text,md,output,分别为纯文本格式，Markdown格式，output格式

### 回调函数

* code 成功非0
* data 当MD格式的链接被点击时，链接href将作为data回调
### 代码范例

Javascript
```
Userinput.Note("handlenote","Title","# md content [a](href)","md")
```

Lua
```
Userinput:Note("handlenote","Title","# md content [a](href)"","md")
```

### 返回值

唯一id

## Userinput.Custom

显示文本

### 原型
```
Custom (script string, value type,value string) string
```

### 描述

预留自定义数据，由具体连接客户端处理

* script 回调脚本
* value 值

### 回调函数

* code 成功非0
* 其他由具体客户端决定

### 代码范例

Javascript
```
Userinput.Custom("MyScript","MyType","Mydata")
```

Lua
```
Userinput:Custom("MyScript","MyType","Mydata")
```

### 返回值

唯一id

## Userinput.NewList

新建列表

### 原型

```
NewList(title string, intro string, withfilter bool)*List
```

### 描述

创建一个列表

列表需要发布(Publish)后才会显示
* title 列表标题
* intro 列表介绍
* withfilter 列表是否带过滤器
### 代码范例

Javascript
```
var list=Userinput.NewList("test list","A test list",true)
```

Lua
```
local list=Userinput:NewList("test list","A test list",true)
```

### 返回值

新列表对象

## Userinput.NewDatagrid

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
var datagrid=Userinput.NewDatagrid("test data grid","A test data grid")
```

Lua
```
local datagrid=Userinput:NewDatagrid("test data grid","A test data grid")
```

### 返回值

新数据表格对象

## Userinput.NewVisualPrompt

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
var vp=Userinput.NewVisualPrompt("test data visualprompt","A test data grid","http://127.0.0.1/test.png")
```

Lua
```
local vp=Userinput:NewVisualPrompt("test data visualprompt","A test data grid","http://127.0.0.1/test.png)
```

### 返回值

新可视化输入对象

## Userinput.HideAll

隐藏界面UI

### 描述

隐藏界面的UI

由于confirm,Alert,promot属于弹框，所以调用该方法后不保证会关闭，这个取决于实际前端的实现

List,Datagrid,VirualPrompt的操作界面会进行隐藏

### 代码范例
Javascript
```
Userinput.HideAll()
```

Lua
```
Userinput.HideAll()
```

### 返回值

无
## List.Append

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
list.Append("sword","长剑")
```

Lua
```
list.Append("sword","长剑")
```

### 返回值

无

## List.SetMulti

列表设置多选

### 原型

```
SetMulti(mutli bool)
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
List.SetMulti(true)
```

Lua
```
List.SetMulti(true)
```

### 返回值

无

## List.SetValue

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
List.SetValue(["key1","key2","key3"])
```

Lua
```
List.SetValue({"key1","key2","key3"})
```

### 返回值

无

## List.Publish

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
List.Publish("handlelist")
```

Lua
```
List.Publish("handlelist")
```

### 返回值

唯一id

## Datagrid.Append

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
Datagrid.Append("key","value")
```

Lua
```
Datagrid.Append("key","value")
```

### 返回值

无

## Datagrid.ResetItems

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
Datagrid.ResetItems()
```

Lua
```
Datagrid.ResetItems()
```

### 返回值

空
## Datagrid.SetFilter

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
Datagrid.SetFilter("filter)
```

Lua
```
Datagrid.SetFilter("filter)
```

### 返回值

无

## Datagrid.GetFilter

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
world.Note(Datagrid.GetFilter())
```

Lua
```
Note(Datagrid.GetFilter())
```
### 返回值

字符串

## Datagrid.SetMaxPage

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
Datagrid.SetMaxPage(10)
```

Lua
```
Datagrid.SetMaxPage(10)
```

### 返回值

无

## Datagrid.SetPage

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
Datagrid.SetPage(3)
```

Lua
```
Datagrid.SetPage(3)
```

### 返回值

无

## Datagrid.GetPage

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
world.Note(Datagrid.GetPage())
```

Lua
```
Note(Datagrid.GetPage())
```
### 返回值

整数的当前页

## Datagrid.SetOnCreate

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
Datagrid.SetOnCreate("handleupdate")
```

Lua
```
Datagrid.SetOnCreate("handleupdate")
```

### 返回值

无

## Datagrid.SetOnUpdate

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
Datagrid.SetOnUpdate("handleupdate")
```

Lua
```
Datagrid.SetOnUpdate("handleupdate")
```

### 返回值

无

## Datagrid.SetOnView

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
Datagrid.SetOnView("handleview")
```

Lua
```
Datagrid.SetOnView("handleview")
```

### 返回值

无

## Datagrid.SetOnSelect

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
Datagrid.SetOnSelect("handleselect")
```

Lua
```
Datagrid.SetOnSelect("handleselect")
```

### 返回值

无

## Datagrid.SetOnDelete

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
Datagrid.SetOnDelete("handledelete")
```

Lua
```
Datagrid.SetOnDelete("handledelete")
```

### 返回值

无

## Datagrid.SetOnFilter

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
Datagrid.SetOnFilter("handlefilter")
```

Lua
```
Datagrid.SetOnFilter("handlefilter")
```

### 返回值

无

## Datagrid.SetOnPage

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
Datagrid.SetOnPage("handlepage")
```

Lua
```
Datagrid.SetOnPage("handlepage")
```

### 返回值

无

## Datagrid.Hide

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
Datagrid.Hide()
```

Lua
```
Datagrid.Hide()
```

### 返回值

无

## Datagrid.Publish

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
Datagrid.Publish("onclose")
```

Lua
```
Datagrid.Publish("onclose")
```

### 返回值

唯一id

## VisualPrompt.SetMediaType

可视化输入设置媒体类型

### 原型
```
SetMediaType(t string)
```

### 描述

设置可视化输入的媒体类型，暂时只支持image，base64slide,output和text

默认为image

output需要将url设置为DumpOutput的返回值(JSON格式的Line数组)

base64需要将url设置为 | 分割的，带格式头的 base64图片，如 "data:image/jpeg;base64,"

如

```Javascript
 let data = Binary.Base64Encode(App.Core.CaptchaReq.ResponseBodyArrayBuffer())
 App.Core.CaptchaImages.push("data:image/jpeg;base64, " + data)
 let vp = Userinput.newvisualprompt("验证码", intro, App.Core.CaptchaImages.join("|"))
 vp.SetMediaType("base64slide")
 vp.publish("App.Core.OnCaptchaSubmit")
```

### 代码范例

Javascript
```
VisualPrompt.SetMediaType("image")
```

Lua
```
VisualPrompt.SetMediaType("image")
```

### 返回值

无

## VisualPrompt.SetPortrait

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
VisualPrompt.SetPortrait(true)
```

Lua
```
VisualPrompt.SetPortrait(true)
```

### 返回值

无

### VisualPrompt.SetRefreshCallback

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
VisualPrompt.SetRefreshCallback("handlerefresh")
```

Lua
```
VisualPrompt.SetRefreshCallback("handlerefresh")
```

### 返回值

无

## VisualPrompt.Publish

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
VisualPrompt.Publish("onvisualprompt")
```

Lua
```
VisualPrompt.Publish("onvisualprompt")
```

### 返回值

唯一ID

## VisualPrompt.Append

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
VisualPrompt.Append("sword","长剑")
```

Lua
```
VisualPrompt.Append("sword","长剑")
```

### 返回值

无