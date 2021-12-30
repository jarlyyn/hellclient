# 杂项接口

## Version

版本信息

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Version

### 原型
```
Version() string
```

### 描述

返回版本信息

版本信息格式为

大版本号.YYYY.MM.DD

### 范例代码

Javascript
```
world.Note(world.Version());
```

Lua
```
Note(Version())
```

### 返回值

版本字符串


## Hash

摘要

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Hash

### 原型

```
Hash(text string) string
```

### 描述

返回给定字符串的小写sha摘要


### 代码范例

JavaScript
```
world.note (world.hash ("This Mud is running on the Dawn Codebase"));
```

Lua
```
Note (Hash("This Mud is running on the Dawn Codebase"))
```

### 返回值

字符串
## Base64Encode

Base64编码

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Base64Encode

### 原型

```
Base64Encode(text string, mutliline bool) string
```

### 描述

将给到的字符串编码为md5编码

* text 需要编码的内容
* mutliline 是否每76个字符插入一个换行

### 代码范例

Javascript
```
world.note (world.base64encode ("swordfish", 0));
```

Lua
```
Note (Base64Encode ("swordfish", 0))
```

### 返回值
字符串

## Base64Decode

Base64解码

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Base64Decode

### 原型

```
Base64Decode(text string) *string
```

### 描述

将给到的md5码结码

* text 需要解码的内容

### 代码范例

Javascript
```
world.note (world.Base64Decode ("TmljayBHYW1tb24="));
```

Lua
```
Note (Base64Decode ("TmljayBHYW1tb24="))
```

### 返回值
* 成功解码则返回解码后的字符串
* 解码失败返回空

## Trim

去空格

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Trim

### 原型

```
Trim(source string) string
```

### 描述

去除字符串两端的空格
* source 待处理的字符串

### 代码范例

Javascript
```
world.Note("*" + world.Trim("  mystring  ") + "*");
```
Lua
```
Note("*" .. Trim("  mystring  ") .. "*")
```

### 返回值

去除两对岸空格后的字符串

## GetUniqueNumber

获取唯一数值

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetUniqueNumber

### 原型
```
GetUniqueNumber() int
```

### 描述

获取一个唯一数字。该数值每次从程序启动开始从0自赠

### 代码范例

Javascript
```
world.note (world.GetUniqueNumber());
```

Lua
```
Note (GetUniqueNumber())
```

### 返回值

0到 2147483647 之间的整数

### GetUniqueID

获取唯一ID

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetUniqueID

### 原型

```
GetUniqueID() string
```

### 描述

返回一个唯一字符串，注意，和mush不同，返回的字符串是不定长的。

### 范例代码

Javascript
```
world.Note (world.GetUniqueID ());
```

Lua
```
Note (GetUniqueID ())
```

### 返回值

字符串


## CreateGUID

获取GUID

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=CreateGUID

### 原型
```
CreateGUID() string
```
### 描述

生成一个uuid v1形式的guid

格式为

```
6B29FC40-CA47-1067-B31D-00DD010662DA
```

### 范例代码

Javascript
```
world.Note (world.CreateGUID ());
```

Lua
```
Note (CreateGUID ())
```

### 返回值

36位uuid字符串

## SplitN

分割字符串

### 代码原型
```
SplitN(text string, sep string, n int) []string
```

### 描述

根据给到的信息分割字符串

* text 待分割的字符串
* sep 分割符
* n 分割数量

其中当n
* 大于 0 时，返回最多 n组数据，最后一组包含所有剩下的文字
* 等于0时，返回空
* 小于0时，返回全部可分割的数据

### 代码范例

Javascript
```
text=world.SplitN("a=4=5","=",2)
```

Lua
```
text=SplitN("a=4=5","=",2)

```
### 返回值

字符串

## UTF8Len

获取UTF8长度

### 原型

```
UTF8Len(text string)
```

### 描述

返回utf8字符串的长度

英文和中文的长度都计算为1

比如 "a甲b乙" 的长度 为4

### 代码范例

Javascript
```
world.Note(world.UTF8LEN("a甲b乙"))
```

Lua
```
Note(world.UTF8LEN("a甲b乙"))
```

### 返回值

给定的utf8字符串的长度

## UTF8Sub

获取UTF8子字符串

### 原型

```
UTF8Sub(text string, start int, end int)
```

### 描述

从给到的开始和结束位置截取字符串

* text 需要截取的原字符串
* start 开始截取的位置，小于0则当作0处理, 大于等于字符串长度则返回空字符串
* end 结束截取的位置，小于等于0则截取到字符串结束

英文和中文的长度都计算为1

比如 UTF8Sub("a甲b乙",1,3)的值为"甲b"

###  代码范例

Javascript
```
world.Note(world.UTF8Sub("a甲b乙",2,3))
```

Lua
```
Note(UTF8Sub("a甲b乙",2,3))
```
### 返回值

返回截取出的字符串

## ToUTF8

转换为UTF8字符串

### 原型
```
ToUTF8(code string, text string) *string
```

### 描述

将给定编码的字符串转码为utf8字符串

* code 编码，可选范围为 utf8,gbk,big5
* text 需要转换的字符串

### 范例代码

Javascript
```
world.Note(world.ToUTF8("gbk",rawtext))
```

Lua
```
Note(ToUTF8("gbk",rawtext))
```

### 返回值
* 转换成功返回对应字符串
* 返回失败(如数据问题或编码不支持)返回空

## FromUTF8

转换自UTF8字符串

### 原型
```
ToUTF8(code string, text string) *string
```

### 描述

将utf8字符串转码为给定编码的字符串

* code 编码，可选范围为 utf8,gbk,big5
* text 需要转换的字符串

### 范例代码

Javascript
```
data=world.FromUTF8("gbk","你好"))
```

Lua
```
data=FromUTF8("gbk","你好"))
```

### 返回值
* 转换成功返回对应字符串
* 返回失败(如数据问题或编码不支持)返回空

## Encrypt

加密字符串

### 原型
```
Encrypt(data, key string) *string
```

### 描述

用给到的数据和秘钥进行加密

* data 需要加密的数据
* key 密钥

加密算法为

生成随机16位IV,通过PKCS7补位的AES加密，将IV和加密数据拼接后通过base64编码

### 范例代码

Javascript
```
world.Note(world.Encrypt("hellclient","key"))
```

Lua
```
Note(Encrypt("hellclient","key"))
```

### 返回值
* 加密成功返回字符串
* 加密失败返回空

## Decrypt

解密字符串

### 原型
```
Decrypt(data, key string) *string
```

### 描述

用给到的数据和秘钥进行解密

* data 需要加密的数据
* key 密钥

解密算法为

通过base64解码，将数据前16位切出作为IV,剩余部分通过PKCS7补位的AES解密，

### 范例代码

Javascript
```
world.Note(world.Decrypt("ziLol5II2WEF4yAeQCf2HMrRDdWMu+Afp1u2ysVSgck=","key"))
```

Lua
```
Note(Decrypt("ziLol5II2WEF4yAeQCf2HMrRDdWMu+Afp1u2ysVSgck=","key"))
```

### 返回值
* 解密成功返回字符串
* 解密失败返回空