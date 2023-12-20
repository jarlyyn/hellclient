# 通讯接口

## Broadcast
广播

### 原型
```
Broadcast(msg string, gloabl bool)
```

### 描述

发送广播消息

* msg 广播内容,为空则不发送
* global 是否全局发送，为true会发送到config文件里指定的全局网关,转发给所有连入网关的客户端

注意，需要在脚本里设置Channel才会发送

同意Channel,同一个host和端口的脚本才会收到消息

### 代码范例

Javascript
```
world.Broadcast("found npc 1104",true)
```

Lua
```
Broadcast("found npc 1104",true)
```

### 返回值

无

## Notify

通知

### 原型

* Notify(title string, body string,link *string)

### 描述

通过配置文件里设置的通知方式(暂时只有邮件)发送通知

* title 通知标题
* body 通知内容
* link 链接 可选
### 代码范例

Javascript
```
world.Notify("ID挂了","被NPC xxx 杀死了",'http://www.google.com')
```

Lua
```
Notify("ID挂了","被NPC xxx 杀死了",'http://www.google.com')
```

### 返回值

无

## Request
发送请求

### 原型
```
Request(type string, data string) string
```

### 描述

发送请求

* type 请求类型，原则上相应请求的服务应该返回同样的type
* data 发送的数据

参考[文档](../features/requestresponse.md)
### 代码范例

Javascript
```
world.Note(world.Request("testmessage","messagedata"))
```

Lua
```
Note(Request("testmessage","messagedata")))
```

### 返回值

唯一ID,可用于响应出发函数里进行匹配