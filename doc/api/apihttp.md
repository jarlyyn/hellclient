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