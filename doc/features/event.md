# 脚本事件处理

每个脚本可以在脚本设置中定义脚本事件处理函数。

当前版本可设置的内容如下

## 助理触发函数

助理按钮(输入框左侧图标)被点击后调用，无参数，无返回值，一般用作脚本的开始菜单。

## 快捷键触发函数

当系统端支持的快捷键被按下时触发。不同的客户端能实现的快捷键不一样。为一参数为传入的键值,无返回值。

## 广播触发函数

当接受到广播时触发。

必须广播方的 “广播频道”值，服务器Host和Post值完全一致才能触发。

函数有三个参数

msg, global, channel

msg为广播的信息

global为是否是全局事件，全局事件可能由hellswitch转发,channel为具体的广播频道值。

无返回值

## 响应触发函数

[请求响应机制](./requestresponse)中的处理函数

参数为 

msgtype, id, data

msgtype和id 与request发出时一致，data为响应方返回的数据。

无返回值。

## 加载触发函数

加载时触发，无参数，无返回值。

## 关闭触发函数

关闭时触发，无参数，无返回值。

## 连线触发函数

连接到游戏器服务时触发，无参数，无返回值。

## 断线触发函数

与游戏服务器断线时触发，无参数，无返回值。

## HUD点击函数

用户点击客户端的hud时的处理函数。

有两个参数，x,y,一般为想对显示区域的坐标百分比。与客户端强相关，一般不使用

## Buffer处理函数

当客户端接受到字符，但还没有接受到回车生成新行时调用。

会在 字符长度在 “Buffer处理函数最小响应字数“和 “Buffer处理函数最大响应字数”之间时触发

传入参数有两个，data,bytes,分别为当前字符的utf8字符信息和字节数组,lua下一般两者一致

返回真值，则会在当前位置插入回车断行。

## SubNegotiation处理函数

服务器发来非文字信息时的触发

参数有两个，code, data。

第一个code为一个byte,data为剩下信息。

一般用于处理gmcp信息等数据

无返回值

## 获得焦点函数

游戏成为当前焦点时调用，无参数，无返回值。

特别的，就算用户是再次点击游戏，没有发生切换，也会调用，方便客户端更新页面。

## 失去焦点函数

游戏失去当前焦点时调用，无参数，无返回值。

特别的，就算用户是再次点击游戏，没有发生切换，也会调用，方便客户端更新页面。