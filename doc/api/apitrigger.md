# 触发器接口

[返回](readme.md)

## AddTrigger

添加触发器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=AddTrigger

### 原型
```
AddTrigger(triggerName string, match string, responseText string, flags int, colour int, wildcard int, soundFileName string, scriptName string) int {
```

### 描述

添加脚本触发器

* triggerName 触发器名
* match 匹配文字
* responseText 匹配后发送的文字
* flags 标识位
* colour 废弃
* wildcard 废弃
* soundFileName 废弃
* scriptName 脚本名

flags的可选值为

* eEnabled = 1; // 激活触发器
* eOmitFromLog = 2; // 废弃
* eOmitFromOutput = 4; //  屏蔽输出
* eKeepEvaluating = 8; // 继续执行
* eIgnoreCase = 16; // 无视大小写
* eTriggerRegularExpression = 32; // 正则处罚
* eExpandVariables = 512; // 扩展变量
* eReplace = 1024; // 替换同名触发
* eTemporary = 16384; // 临时触发
* eTriggerOneShot = 32768; // 一次性触发

Lua时别名标记储存在trigger_flag表

* Enabled = 1
* OmitFromLog = 2
* OmitFromOutput = 4
* KeepEvaluating = 8
* IgnoreCase = 16
* RegularExpression = 32
* ExpandVariables = 512
* Replace = 1024
* Temporary = 16384
* LowercaseWildcard = 2048
* OneShot = 32768

### 代码范例
Javascript:
```
world.AddTrigger("monster", "* attacks", "flee", 1, 0, 0, "", "");
```

Lua:
```
AddTrigger("monster", "* attacks", "flee", trigger_flag.Enabled , 0, 0, "", "")

```

### 返回值

* eTriggerAlreadyExists 触发已存在
* eTriggerCannotBeEmpty 匹配文字为空
* eOK 成功