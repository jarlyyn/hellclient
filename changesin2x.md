# 2.X版本主体变动

## hellclient.net支持

hellclient.net是为了解决go版本在windows下与msvc生态不兼容的问题开发的C#版本。功能与go版本一致。

## 加入asni支持

引入了asni(utf编码)字符串的支持，方便进行更底层的操作。

包括

* AddAnsi函数 添加一个Ansi行
* LastAnsi 函数，获取最新的Ansi行信息

## Mapper 进入废弃状态

建议使用HellMapManager相关算法

## 引入新的全局hook

* OnLine 在所有触发执行之前调用，可以终止触发执行
* OnAfterLine 在所有触发执行之后调用
* OnSend 在向mud发送数据时调用，可以终止发送

## 废弃prompt概念

原有prompt作为新行显示

## 破坏性变更：废除 OnBuffer的hook

对性能影响极大，功能极弱，可以通过OnLine和AddAnsi实现大部分功能