# 界面接口

界面/客户端 与 主服务的交互是通过websocket + json格式进行的。

Websocket连接地址为

\[服务器IP\]:\[服务器端口\]/ws

如有设置账号密码，需要进行http basic验证。具体交互信息如后

## 服务器发送信息

服务器发送信息代码位于 

* [/src/modules/world/prophet/adapter.go](/src/modules/world/prophet/adapter.go) 中的 initAdapter 函数
* [/src/modules/msg/msg.go](/src/modules/msg/msg.go)

## 客户端发送信息

服务器接受信息代码位于 

* [/src/modules/world/prophet/handlers.go](/src/modules/world/prophet/handlers.go) 中的 initHandlers 函数
