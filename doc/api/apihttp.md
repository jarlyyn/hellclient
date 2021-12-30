# http请求

## 属于
http请求是一个用于向外部发送http请求的组件

### Request

请求

http请求的主体，可以对Request进行维护，执行后可以获取相应的相应值

## HTTP.New

创建新请求

### 原型

```
New(method string, url string) *Request
```

### 描述

创建一个新请求代码

* method http请求method,一般为GET POST PUT DELETE
* url 请求的地址

### 代码范例

Javascript
```
var Request=HTTP.New("GET","http://www.baidu.com)
```

Lua
```
var Request=HTTP:New("GET","http://www.baidu.com)
```

### 返回值

Request对象

## Request.GetID
获取唯一id
### 原型
```
GetID() string
```

### 描述
获取请求的唯一ID

### 代码范例

Javascript
```
world.Note(Request.GetID())
```

Lua
```
Note(Request.GetID())
```

### 返回值

请求的唯一ID,字符串

## Request.GetURL

获取请求URL

### 原型
```
GetURL() string
```

### 描述

获取请求的地址

### 代码范例
Javascript
```
world.Note(Request.GetURL())
```

Lua
```
Note(Request.GetURL())
```

### 返回值

无

## Request.SetURL

设置请求URL

### 原型
```
SetURL(url string)
```

### 描述

设置请求的URL

### 代码范例
Javascript
```
Request.SetURL("https://github.com")
```

Lua
```
Request.SetURL("https://github.com")
```

### 返回值

无

## Request.GetMethod

获取请求Method

### 原型
```
GetMethod() string
```

### 描述

返回请求的method

### 代码范例

Javascript
```
world.Note(Request.GetMethod)
```

Lua
```
Note(Request.GetMethod)
```

### 返回值

字符串

## Request.SetMethod

设置请求URL

### 原型
```
SetMethod(method string)
```

### 描述

设置请求的URL

### 代码范例

Javascript
```
Request.SetMethod("POST")
```

Lua
```
Request.SetMethod("POST")
```

### 返回值

无

## Request.GetBody

获取请求正文

### 原型
```
GetBody() string
```

### 描述

获取请求的正文

### 代码范例
Javascript
```
world.Note(Request.GetBody())
```
Lua
```
Note(Request.GetBody())
```

### 返回值

字符串

## Request.SetBody

设置请求正文

### 原型
```
SetBody(body string)
```

### 描述

设置请求正文

### 原型
Javascript
```
Request.SetBody("text body")
```

Lua
```
Request.SetBody("text body")
```

### 返回值

无

## Request.SetHeader

设置请求头

### 原型
```
SetHeader(name string, value string)
```

### 描述
设置指定的请求头
* name 请求头名
* value 请求头值

### 代码范例
Javascript
```
Request.SetHeader("host","mydomain")
```

Lua
```
Request.SetHeader("host","mydomain")
```

### 返回值

无

## Request.AddHeader

添加请求头

### 原型
```
AddHeader(name string, value string)
```


### 描述

为请求添加请求头

### 代码范例
Javascript
```
Request.AddHeader("field1","more value")
```

Lua
```
Request.AddHeader("field1","more value")
```
### 返回值

无

## Request.DelHeader

删除请求头

### 原型

```
DelHeader(name string)
```

### 描述

删除请求的指定请求头

### 代码范例
Javascript
```
Request.DelHeader("myheader")
```
Lua
```
Request.DelHeader("myheader")
```

### 返回值

无

## Request.GetHeader

获取请求头

### 原型
```
GetHeader(name string) string
```

### 描述

获取指定的请求头

### 代码范例
Javascript
```
Request.GetHeader("myheader")
```

Lua
```
Request.GetHeader("myheader")
```
### 返回值

字符串

## Request.HeaderValues

获取请求头全部值

### 原型
```
HeaderValues(name string) []string
```

### 描述

获取指定请求头的所有值

### 代码范例
Javascript
```
var values=Request.HeaderValues("myheader")
values.forEach(function(value){
    world.Note(value)
})
```

Lua
```
local values=Request.HeaderValues("myheader")
for k, value in pairs(values) do
    Note(value)
end
```

### 返回值

字符串列表

## Request.HeaderFields

获取请求头全部字段

### 原型
```
HeaderFields() []string
```

### 描述

获取请求的所有请求头名

### 代码范例
Javascript
```
var fields=Request.HeaderFields()
fields.forEach(function(field){
    world.Note(field)
})
```

Lua
```
local fields=Request.HeaderFields()
for k, field in pairs(fields) do
    Note(field)
end
```

### 返回值

字符串列表
## Request.ResetHeader

重置请求头

### 原型
```
ResetHeader()
```

### 描述

重置请求头

### 代码范例
Javascript
```
Request.ResetHeader()
```

Lua
```
Request.ResetHeader()
```
### 返回值

无
## Request.AsyncExecute

异步执行请求

### 原型
```
AsyncExecute(script string)
```

### 描述

异步执行请求

* script 回调脚本代码

注意，如果
* 游戏没有HTTP授权
* 游戏没有信任URL或者HOST头中的地址

会报错
### 回调函数

* code 成功返回0,错误返回-1
* data 成功返回请求地址,错误返回错误内容
### 范例代码
Javascript
```
Request.AsyncExecute("handlerequest")
```

Lua
```
Request.AsyncExecute("handlerequest")
```

### 返回值

无

## Request.ExecuteStatus

获取执行状态

### 原型
```
ExecuteStauts() int
```

### 描述

返回请求的执行状态

### 范例代码
Javascript
```
world.Note(Request.ExecuteStatus)
```

Lua
```
Note(Request.ExecuteStatus)
```

### 返回值
* 0 准备状态
* 1 执行中
* 2 执行成功
* 3 执行失败

## Request.FinishedAt

获取请求成功时间

### 原型
```
FinishedAt() int64
```

### 描述

返回请求成功的unix时间戳

如果请求没有成功结束，会报错

### 代码范例
Javascript
```
world.Note(Request.FinishedAt())
```

Lua
```
Note(Request.FinishedAt())
```
### 返回值
Unix时间戳

## Request.ResponseStatusCode

获取相应状态码

### 原型
```
ResponseStatusCode() int
```

### 描述

返回请求的HTTP 响应码

如果请求没有成功结束，会报错

### 代码范例

Javscript
```
world.Note(Request.ResponseStatusCode())
```

Lua
```
Note(Request.ResponseStatusCode())
```
### 返回值
整数

## Request.ResponseBody

获取响应正文

### 原型
```
ResponseBody() string
```

### 描述

返回响应的正文

如果请求没有成功结束，会报错

### 代码范例

Javascript
```
world.Note(Request.ResponseBody())
```

Lua
```
Note(Request.ResponseBody())
```

### 返回值

字符串

## Request.ResponseHeader

获取响应头

### 原型

```
ResponseHeader(name string) string
```

### 描述

获取响应指定响应的响应头

如果请求没有成功结束，会报错

### 范例代码

Javascript
```
world.Note(Request.ResponseHeader("myheader))
```

Lua
```
Note(Request.ResponseHeader("myheader))
```

### 返回值

字符串

## Request.ResponseHeaderValues

获取响应头全部值

### 原型

```
ResponseHeaderValues(name string) []string
```

### 描述

返回指定响应字段的全部响应值

如果请求没有成功结束，会报错

### 代码范例
Javascript
```
var values=Request.ResponseHeaderValues("myheader")
values.forEach(function(value){
    world.Note(value)
})
```

Lua
```
local values=Request.ResponseHeaderValues("myheader")
for k, value in pairs(values) do
    Note(value)
end
```

### 返回值

字符串列表

## Request.ResponseHeaderFields

获取响应头全部字段

### 原型
```
ResponseHeaderFields() []string
```

### 描述

获取响应的所有响应头名

如果请求没有成功结束，会报错

### 代码范例
Javascript
```
var fields=Request.ResponseHeaderFields()
fields.forEach(function(field){
    world.Note(field)
})
```

Lua
```
local fields=Request.ResponseHeaderFields()
for k, field in pairs(fields) do
    Note(field)
end
```

### 返回值

字符串列表