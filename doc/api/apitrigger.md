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

## AddTriggerEx

高阶添加触发器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=AddTriggerEx

### 原型
```
AddTriggerEx(triggerName string, match string, responseText string, flags int, colour int, wildcard int, soundFileName string, scriptName string, sendTo int, sequence int) int
```

### 描述

添加触发器

* triggerName 触发器名
* match 匹配文字
* responseText 响应文字
* flags 标识
* colour 废弃
* wildcard 废弃
* soundFileName 废弃
* scriptName 脚本名
* sendTo 发送到
* sequence 优先级

可用的flags值

* var eEnabled = 1; // 是否激活
* var eOmitFromLog = 2; // 废弃
* var eOmitFromOutput = 4; // 屏蔽输出
* var eKeepEvaluating = 8; // 继续执行
* var eIgnoreCase = 16; // 忽视大小写
* var eTriggerRegularExpression = 32; // 正则表达式
* var eExpandVariables = 512; // 展开变量
* var eReplace = 1024; // 替换同名触发器
* var eTemporary = 16384; // 临时触发器

lua可以使用trigger_flag表

* trigger_flag.Enabled = 1
* trigger_flag.OmitFromLog = 2
* trigger_flag.OmitFromOutput = 4
* trigger_flag.KeepEvaluating = 8
* trigger_flag.IgnoreCase = 16
* trigger_flag.RegularExpression = 32
* trigger_flag.ExpandVariables = 512
* trigger_flag.Replace = 1024
* trigger_flag.Temporary = 16384
* trigger_flag.LowercaseWildcard = 2048

### 代码范例

Javascript:
```
world.AddTriggerEx ("", "* attacks", "You are under attack!", 1, 0, 0, "", "", 2, 50);
```

Lua:
```
AddTriggerEx ("", "* attacks", "You are under attack!", trigger_flag.Enabled, 0, 0, "", "", 2, 50)
```

### 返回值

* eTriggerAlreadyExists 触发器已存在
* eTriggerCannotBeEmpty 匹配文字不能为空
* eOK 添加成功

## DeleteTrigger

删除触发器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteTrigger

### 原型

```
DeleteTrigger(name string) int
```

### 描述

删除指定的用户触发器

* name 用户组名

### 代码范例

Javascript
```
world.DeleteTrigger("my_trigger");
```

Lua
```
DeleteTrigger("my_trigger")
```

### 返回值
* eTriggerNotFound: 触发器未找到
* eOK: 删除成功

## DeleteTriggerGroup

删除触发器组

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteTriggerGroup

### 原型

```
DeleteTriggerGroup(group string) int
```

### 描述

删除指定的用户触发器组

* group 用户组名

### 代码范例

Javscript
```
world.DeleteTriggerGroup ("groupname");
```
lua
```
DeleteTriggerGroup ("groupname")
```
### 返回值

涮出的触发器数量

## EnableTrigger

激活触发器

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=EnableTrigger

### 原型

```
EnableTrigger(name string, enabled bool) int
```

### 描述

激活脚本触发器

* name 触发器名
* enabled 是否激活
### Lua注意事项

enabled为空时，值为True

### 范例代码

Javascript
```
world.EnableTrigger("monster", true);  # enable trigger
world.EnableTrigger("monster", false);  # disable trigger
```

Lua
```
EnableTrigger("monster", true)  -- enable trigger
EnableTrigger("monster", false)  -- disable trigger
```

### 返回值
* eTriggerNotFound 触发器未找到
* eOK 执行成功

## EnableTriggerGroup

激活触发器组

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=EnableTriggerGroup

### 原型

```
EnableTriggerGroup(group string, enabled bool) int
```

### 描述

激活触发器组
* group 触发器组名
* enabeld 激活状态

### Lua注意事项

enabled为空时，值为True


### 代码范例

Javascript
```
world.DeleteTriggerGroup ("groupname");
```

Lua
```
DeleteTriggerGroup ("groupname")
```

### 返回值

组内触发器数量