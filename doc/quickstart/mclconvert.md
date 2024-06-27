# mclconvert使用说明
mclconvert是一个将mush的mcl文件转换为hellclient的toml文件的工具

## (一)MCL文件预处理

mush的mcl本身并不指定编码，是根据系统编码走的,而hellclient内部是使用utf8格式的编码，本身有编码概念。

因此，mcl文件，主要是gbk编码的mcl文件要先转换为utf编码。

转换编码可以使用各种文本编辑器，比如vscode的以gbk编码打开，以utf8编码保存的功能


## (二)MCL文件转换

在mclconvert目录下执行 mclconvert mcl文件名，屏幕会输出toml内容，这时候可以用管道符将其重命名为script.toml文件

例

```
./mclconvertor /tmp/aqiang.MCL >/tmp/script.toml
```

然后script.toml放在新目录下，将原程序放在该目录的script文件夹下，入口程序改为main.js或main.lua即可

同时，mclconvert也提供转换成world文件，即提取变量的功能

格式为

```
./mclconvertor -world /tmp/aqiang.MCL >/tmp/aqiang.toml
```

注意，所有的代码和引入的文本数据都要转码

## (三)代码修正

代码修正主要包括两个方面

首先是字符串修正。

在mush中，gbk编码是以一个汉字算两个字符，一个英文算一个字符的形式编码的

而在hellclient中，一个汉字和一个英文都是一个字符

所以计算文字长度，切割文字和 带{a,b}计算长度的正则都要进行修正。

除此之外，hellclient没有自定义颜色，所以不能直接判断颜色的grb值，需要用对应的mush接口获取预先定义的颜色值。

其他的就是部分不兼容或者没实现的接口了。

由于机制不同，又是黑箱开发的接口，可能有部分接口表现不一致，这可以联系我看看是否能调成和mush一致的接口

## (四)功能限制

由于设计上的理念问题，以下Mush功能无法在Hellclient中实现

* Lua加载dll,可能需要通过[request/response机制](../features/requestresponse.md)重写
* sqlite功能
* Plugin功能