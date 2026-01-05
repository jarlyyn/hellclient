# Hellclient mud客户端

Hellclient是一款支持Lua脚本和Javascript脚本的，采用B/S(浏览器/服务器)架构的Mud客户端。

客户端的目标是在个人电脑/服务器上进行长时间的稳定挂机，提高开发脚本效率，降低资源消耗。

追求更好的运行脚本胜过人工进行Mud操作

[讨论社区](https://forum.hellclient.com)

## 关于Mushclient

由于作者之前是使用Mushclient进行mud机器人制作的。

所以hellclient的机器人主体概念和mushclient一致，触发器/计时器/别名等大部分属性和mushclient一致，部分API接口与mushclient一致 [查看mushclient接口兼容性](doc/api/mush.md)

mushclient是mud活跃时代的一款非常杰出，优秀的客户端。

## 系统支持

当前的hellclient ui的系统支持为

* windows 7 sp2/windows 2008 +
* centos 7/debian 10 +

服务器自启动，Linux可以使用systemd,参考 [/system/system.d/hellclient.service](/system/system.d/hellclient.service),注意调整user,windows建议使用NSSM[https://nssm.cc/](https://nssm.cc/)。

## 界面

Hellclient本身是通过B/S架构提供操作控制的。

为了方便 多客户端管理/移动使用/操作体验/通知，也提供了跨平台的管理界面Hellclient UI。

Github地址为:[https://github.com/hellclient-scripts/hellclientui](https://github.com/hellclient-scripts/hellclientui)

提供了Windows/Linux/Android/Ios的客户端，以及Mac OS x 的实验性支持。

如有需要可以进行安装使用

## 特性

Hellclient拥有以下特性

* [游戏/脚本分离](doc/features/features.md#游戏/脚本分离)
* [Metronome限流器](doc/features/features.md#Metronome限流器)
* [Mapper地图组件](doc/features/features.md#Mapper地图组件)
* [Javascript脚本支持](doc/features/features.md#Javascript脚本支持)
* [Lua支持](doc/features/features.md#Lua支持)
* [用户授权](doc/features/features.md#用户授权)
* [HTTP组件](doc/features/features.md#HTTP组件)
* [广播消息](doc/features/features.md#广播消息)
## 文档

* [快速开始](doc/quickstart/quickstart.md)
* [接口文档](doc/api/readme.md)
* [事件处理函数](doc/features/event.md)
* [回调机制](doc/features/event.md)
* [开发说明](doc/develop/readme.md)
### 脚本范例

* [Helllua](https://github.com/hellclient-scripts/helllua) 原Mushclient机器人移植
* [zsz-self.jvs](https://github.com/hellclient-scripts/zsz-self.jvs) 原Mushclient机器人移植
* [helljs](https://github.com/hellclient-scripts/helljs) 针对 终极地狱 mud 全新制作的机器人
* [pkpxkx.noob](https://github.com/hellclient-scripts/pkuxkx.noob) 针对 北大侠客行 mud 全新制作的机器人