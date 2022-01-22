# 特性

## 游戏/脚本分离

## Metronome限流器

Mushclient自带一个固定间隔发送指令的speedwork queeu来实现限流的功能。

由于实际上很多mud采取的是心跳限流，即一个心跳内不能发送超过N个指令的方式限流。

提供一个根据心跳来限流的Metronome限流器,可以制定在固定心跳\(一般是1/2mud心跳\)里发送不超过N个指令的限流器。

并提供了一系列的队列维护指令

接口详见 [API文档](../api/apimetronome.md)

## Mapper地图组件

Hellclient 自带一个简单的地图组件

地图组件的可以供脚本添加房间和路径，设置标签，然后获取路径/搜索房间/获取房间出口信息

接口详见 [API文档](../api/apimapper.md)

## Javascript脚本支持

Hellclient通过使用 [goja库](https://github.com/dop251/goja) 提供 ECMAScript 5.1 的兼容

为了兼容部分Mushclient的现存Jscript库，提供了部分兼容代码

## Lua支持

Hellclient通过使用[gopher-lua库](github.com/yuin/gopher-lua)提供了Lua 5.1的支持

为了安全性，Hellclient对lua的功能进行了部分裁剪

## 用户授权

## HTTP组件

出于安全目的以及技术实现的方式，hellclient不支持传统的mushclient调用系统服务的功能\(lua潜入dll/jscript嵌入Activex对象\)

相应的，根据Mud本分的非及时性的特点，提供了HTTP访问支持，这样可以访问第三方的服务或者本地的自建服务 

接口详见[API文档](../api/apihttp.md)

## 广播消息

mud机器人有一个常见需求是跨客户端交互。

往往是用于全局寻找NPC等。

传统做法是利用Mud的公共频道进行交互。

但这样会受mud公共频道限制，容易干扰其他用户，还有性能压力。

所以hellclient提供了广播机制

默认情况下会在当前客户端的同Host,Port,Channel进行交互

如果在配置文件里开启了Switch，会通过Switch服务器进行全球交换信息

Switch服务器是一个Basic验证的Websocket广播服务器

可以通过 HellclientSwitch进行自建 [github地址](https://github.com/hellclient-scripts/hellclientswitch) [coding镜像](https://jarlyyn.coding.net/public/hellclient/hellclientswitch/git/files)

## 邮件通知

在配置文件中进行设置后，可以通过API发送邮件进行通知

借口详见[API文档](,,/api/apicommunication.md#Notify)