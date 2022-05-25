# 请求响应服务

为了能更高效的扩展Hellclient的功能，同时保证必要的安全性，Hellclient通过暴露一个Websocket接口通信的方式实现功能的高级扩展。

脚本可以发送指定格式的请求到所有监听端口的服务，服务也口发送响应的响应到具体的脚本，通过简单的类似rpc的机制实现类似异步复杂的功能扩展

## 消息格式

消息正文是一个JSON结构体，形式为
```json
{
    "World":"testworld",
    "ID":"123456",
    "Data":"message data"
}
```
字段含义为
* World:具体的游戏ID,响应会发送到到对应的游戏
* ID:消息ID,脚本发出时自动生成，可以用于回调进行匹配
* Data:文字格式的数据内容，需要传递二进制数据的话建议使用Base64编码

## 发送

发送时在脚本里调用 

```Javascript
let msgid=world.Request("msgtype","data")
```
进行发送，Hellclient会自动加上world id 以及消息id 发送到服务 

### 响应

在脚本中设置响应回调函数，形式为
```Javascript
function onResponse(msgtype,id,data){
    if (msgtype="mytype" && id=lastid){
        world.Note("收到响应:"+data)
    }
}
```

### 关于安全性

由于服务已经脱离了hellclient本身，所以hellclient并不直接为服务的安全性负责。

唯一需要注意的就是由于服务能监听所有的消息，所以如果有设置敏感的信息，可以考虑进行加密。