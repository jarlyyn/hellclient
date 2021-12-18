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

## GetTriggerInfo

获取触发器信息

部分兼容mushclient

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetTriggerInfo

### 原型

```
GetTriggerInfo(name string, infotype int) (string, int)
```

### 描述

获取指定的脚本触发器的信息

* name 触发器名称
* infotype 信息类型

可用的infotype值

* 1: 匹配内容 (string)
* 2: 发送内容 (string)
* 3: 废弃
* 4: 脚本名 (string)
* 5: 废弃
* 6: 是否屏蔽输出 (boolean)
* 7: 是否继续执行 (boolean)
* 8: 是否激活 (boolean)
* 9: 是否是正则表达式 (boolean)
* 10: 是否无视大小写 (boolean)
* 11: 是否在同一行重复 (boolean)
* 12: 废弃
* 13: 是否展开变量 (boolean)
* 14: 废弃
* 15: 发送到的位置 (int)
* 16: 优先级 (short)
* 17: 废弃
* 18: 废弃
* 19: 废弃
* 20: 废弃
* 21: 废弃
* 22: 废弃
* 23: 是否是临时触发 (boolean)
* 24: 废弃
* 25: 是否将匹配关键字转为小写 (boolean)
* 26: 组名 (string)
* 27: 发送到的变量名 (string)
* 28: 废弃
* 29: 废弃
* 30: 废弃
* 31: 最后一次匹配的正则表达式变量数量 (long)
* 32: 废弃
* 33: 废弃
* 34: 废弃
* 35: 废弃
* 36: 是否是一次性触发 (boolean)
* 37: 废弃
* 38: 废弃

### 范例代码

Javascript
```
world.note(world.gettriggerinfo ("monster", 2));
```

Lua
```
Note (GetTriggerInfo ("monster", 2))
```

### 返回值

* 成功获取的返回值
* 触发器没找到返回空
* Infotype无效返回空


## GetTriggerList

获取脚本触发器列表

不兼容mushclient

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetTriggerList

### 原型

```
GetTriggerList() []string
```

### 描述

返回脚本触发器名称列表

### 代码范例

Javascript
```
triggerlist = new VBArray(world.GetTriggerList()).toArray();

if (triggerlist)  // if not empty
 for (i = 0; i < triggerlist.length; i++)
   world.note(triggerlist [i]);
```

Lua
```
tl = GetTriggerList()
if tl then
  for k, v in ipairs (tl) do 
    Note (v) 
  end  -- for
end -- if we have any triggers
```

### 返回值

存有脚本触发器名的字符串列表

## GetTriggerOption 

获取触发器选项

部分兼容mushclisnt

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetTriggerOption

### 原型

```
GetTriggerOption(name string, option string) (string, int)
```

### 描述

返回给定名称的计时器的指定选项

* name 计时器名
* option 选项名

可用的option值

* "clipboard_arg": 废弃
* "colour_change_type": 废弃
* "custom_colour": 废弃
* "enabled": y/n - trigger is enabled
* "expand_variables": y/n - 是否展开变量
* "group": (string - 触发器组名)
* "ignore_case": y/n - 是否无视大小写
* "inverse":废弃
* "italic": 废弃
* "keep_evaluating": y/n - 是否继续执行
* "lines_to_match": 0 - 200 - 多行模式下的匹配行数
* "lowercase_wildcard": y/n - 是否将匹配转换为小写
* "match": (string - 匹配内容)
* "match_style":废弃
* "multi_line": y/n - 是否是多行匹配
* "name": (string - 触发器名)
* "new_style":废弃
* "omit_from_log": 废弃
* "omit_from_output": y/n - 是否屏蔽输出
* "one_shot": y/n - 是否是一次性触发
* "other_back_colour": 废弃
* "other_text_colour": 废弃
* "regexp": y/n - 是否是正则表达式
* "repeat": y/n - 是否重复触发
* "script": (string - 脚本名)
* "send": (string - 发送内容)
* "send_to": 0 - 14 - 发送到的位置
* "sequence": 0 - 10000 - 优先级
* "sound": 废弃
* "sound_if_inactive": 废弃
* "user": 废弃
* "variable": (string - 发送到的变量名)

对于数字型的值，传入的字符串应该能转换为数字

对于布尔型的值，传入的值必须是
* "y", "Y", or "1" 为真值
* "n", "N", or "0" 为假值

### 代码范例

Javasript
```
Note (world.GetTriggerOption ("mytrigger", "match"));
```

Lua
```
Note (GetTriggerOption ("mytrigger", "match"))
```

### 返回值

如描述所列

## GetTriggerWildcard

获取触发器匹配值

兼容mushclient

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetTriggerWildcard

### 原型

```
GetTriggerWildcard(triggername string, wildcard string) *string
```

### 描述

获取指定脚本触发器的最后一次触发的匹配内容

* triggername 触发器名
* wildcard 匹配名或匹配的数字序号

变量数量可以通过 GetTriggerInfo 13 获取

### 代码范例

Javascript
```
x = world.GetTriggerWildcard ("mytrigger", "who")  // get wildcard named 'who'
x = world.GetTriggerWildcard ("mytrigger", "what") // get wildcard named 'what'
x = world.GetTriggerWildcard ("mytrigger", "22")   //get wildcard 22 (if it exists)
```

Lua
```
x = GetTriggerWildcard ("mytrigger", "who")  -- get wildcard named 'who'
x = GetTriggerWildcard ("mytrigger", "what") -- get wildcard named 'what'
x = GetTriggerWildcard ("mytrigger", "22")   -- get wildcard 22 (if it exists)
```

### 返回值

* 触发器未找到返回空
* 匹配未找到返回空字符串
* 找到内容则返回对应匹配的文字 

## IsTrigger

判断触发器是否存在

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=

### 原型

```
IsTrigger(name string) int
```

### 描述

判断指定的脚本触发器是否存在

* name 触发器名

### 代码范例

Javascript
```
world.note(world.IsTrigger("myTrigger")); 
```

Lua
```
Note(IsTrigger("myTrigger"))
```

### 返回值

* eTriggerNotFound: 触发不存在
* eOK: 触发存在

