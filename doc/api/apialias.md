# 别名接口

[返回](readme.md)

## AddAlias 

添加别名

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=AddAlias

### 原型:

```
AddAlias(aliasName string, match string, responseText string, flags int, scriptName string) int
```

### 描述：

在游戏内添加一个脚本别名

* aliasName 脚本名称
* match 匹配文本
* responseText 响应内容
* flags 标识位
* scriptName 触发脚本名

可用标识:

* eEnabled = 1; // 添加后别名是否有效
* eKeepEvaluating = 8; // 别名触发后是否继续执行
* eIgnoreAliasCase = 32; // 匹配时是否无视大小写
* eOmitFromLogFile = 64; // 废弃
* eAliasRegularExpression = 128; // 匹配内容是否是正则表达式
* eExpandVariables = 512; // 是否展开变量
* eReplace = 1024; // 是否替代同名别名
* eAliasSpeedWalk = 2048; // 废弃
* eAliasQueue = 4096; // 通过队列发送响应内容
* eAliasMenu = 8192; // 废弃
* eTemporary = 16384; // 临时别名，保存时不保留

Lua时别名标记储存在alias_flag表
* Enabled = 1
* KeepEvaluating = 8 
* IgnoreAliasCase = 32
* OmitFromLogFile = 64
* RegularExpression = 128
* ExpandVariables = 512
* Replace = 1024
* AliasSpeedWalk = 2048
* AliasQueue = 4096
* AliasMenu = 8192
* Temporary = 16384

### 代码范例

Javascript
```
world.AddAlias("food_alias", "eat", "eat food", eEnabled, "");
```
Lua
```
AddAlias("food_alias", "eat", "eat food", alias_flag.Enabled, "")

```
### 返回值

eAliasCannotBeEmpty 匹配内容为空
eAliasAlreadyExists 同名别名已经存在
eOK 添加成功

## DeleteAlias

删除别名

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteAlias

### 原型
```
DeleteAlias(name string) int
```
### 描述

删除指定名称的脚本别名

* name 别名名称

### 代码范例

Javascript:
```
world.DeleteAlias("my_alias");
```

Lua:
```
DeleteAlias("my_alias")
```
### 返回值

* eAliasNotFound 别名无法找到
* eOK 删除成功

## DeleteAliasGroup

删除别名组

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteAliasGroup

### 原型
```
DeleteAliasGroup(group string) int
```
### 描述

删除指定组别的别名
只删除组内的脚本别名

* group 别名组别

### 代码范例

Javascript:
```
world.DeleteAliasGroup ("groupname");
```

Lua:
```
DeleteAliasGroup ("groupname")
```

### 返回值

删除的别名数量

## DeleteTemporaryAliases

删除临时别名

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteTemporaryAliases

### 原型
```
DeleteTemporaryAliases() int
```

### 描述

删除所有临时别名

### 代码范例

Javascript:
```
world.DeleteTemporaryAliases ();
```

Lua
```
DeleteTemporaryAliases ()
```
### 返回值

被删除的别名数量

## EnableAlias 激活别名

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=EnableAlias

### 原型

```
EnableAlias(name string, enabled bool) int
```
### 描述

启用或禁用脚本别名

* name 别名名称
* enabled 激活状态

### 代码范例

Javasript:
```
world.EnableAlias("teleport", true);
```

Lua:
```
EnableAlias ("teleport", true)
```

### Lua注意事项

enabled为空时，值为True

### 返回值

* eAliasNotFound 别名未找到
* eOK 操作成功

## EnableAliasGroup

激活别名组

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=EnableAliasGroup

### 原型
```
EnableAliasGroup(group string, enabled bool) int
```

### 描述

按组启用或禁用别名
注意，将同时影响组内用户和脚本别名

* group 别名组别
* enabled 激活状态

### 代码范例获取别名信息

Javascript:
```
world.EnableAliasGroup ("groupname", 1);  // enable the group
world.EnableAliasGroup ("groupname", 0);  // disable the group
```

Lua:
```
EnableAliasGroup ("groupname", true)  -- enable the group
EnableAliasGroup ("groupname", false)  -- disable the group
```

### Lua注意事项

enabled为空时，值为True

### 返回值

组内别名数量

## GetAliasInfo

获取别名信息

不完全兼容Mushclient

对应Mushclient API:https://www.gammon.com.au/scripts/doc.php?function=GetAliasInfo

### 原型

```
GetAliasInfo(name string, infotype int) (string, int)
```

### 描述

获取给定别名的指定细节

* name 别名名
* infotype 信息类型

可用的infotype值

* 1: 匹配文字(string)
* 2: 发送内容 (string)
* 3: 脚本名 (string)
* 4: 废弃
* 5: 在输出中屏蔽 (boolean)
* 6: 是否有效 (boolean)
* 7: 是否是正则表达式 (boolean)
* 8: 是否无视大小写 (boolean)
* 9: 是否展开变量 (boolean)
* 10: 废弃
* 11: 废弃
* 12: 废弃
* 13: 废弃
* 14: 是否是临时别名 (boolean)
* 15: 废弃 (boolean)
* 16: 组名 (string)
* 17: 变量名 (string)
* 18: 发送到的位置 (long)
* 19: 是否继续执行 (boolean)
* 20: 优先级 (long)
* 21: 废弃
* 22: 废弃
* 23: 废弃
* 24: 废弃
* 25: 废弃
* 26: 废弃
* 27: 废弃
* 28: 废弃
* 29: 'one shot' flag (boolean)
* 30: 废弃
* 31: 废弃

### 代码范例

Javascript:
```
/world.note(world.GetAliasInfo ("my_alias", 2));
```

Lua:
```
Note (GetAliasInfo("my_alias", 2))
```
### 返回值
* 成功获取的返回值
* 别名没找到返回空
* Infotype无效返回空

## GetAliasList

获取脚本别名名称列表

注意，与Mushclient同名函数不完全兼容，返回值类型不同

### 原型
```
GetAliasList() []string
```

### 描述

获取脚本别名名字列表

### 代码范例

Javascript:
```
if (aliaslist)  // if not empty
 for (i = 0; i < aliaslist.length; i++)
   world.note(aliaslist [i]);
```

Lua:
```
al = GetAliasList()
if al then
  for k, v in ipairs (al) do 
    Note (v) 
  end  -- for
end -- if we have any aliases
```
### 返回值

存有脚本别名名的字符串列表

## GetAliasOption

获取别名选项

不完全兼容Mushclient

对应Mushclient API:https://www.gammon.com.au/scripts/doc.php?function=GetAliasOption

### 原型
```
GetAliasOption(name string, option string) (string, int)
```
### 描述

获取指定名称脚本别名的指定选项

* name 别名名
* option 选型名

可用的option值

* "echo_alias" 废弃
* "enabled": y/n  别名是否有效
* "expand_variables": y/n 是否展开变量
* "group": (string - 别民组名)
* "ignore_case": y/n - 无视大小写
* "keep_evaluating": y/n - 继续执行
* "match": (string - 匹配文字)
* "menu": 废弃
* "name": (string - 别名名称)
* "omit_from_command_history": 废弃 
* "omit_from_log": 废弃
* "omit_from_output": y/n - 是否在输出窗口屏蔽
* "one_shot": y/n - 是否是一次性别名
* "regexp": y/n - 是否是正则表达式
* "script": (string - 调用的脚本名)
* "send": (string - 发送的内容)
* "send_to": 0 - 13 - 发送到的位置
* "sequence": 0 - 10000 - 优先级
* "user": 废弃
* "variable": (string - sendto变量名)

布尔值会返回 0\(false\) 或 1\(true\)
### 代码范例

Javascript:
```
Note (GetAliasOption ("myalias", "match"));
```
Lua:
```
Note (GetAliasOption ("myalias", "match"))
```

### 返回值
如描述所列

## IsAlias

判断别名是否存在

对应Mushclient API:https://www.gammon.com.au/scripts/doc.php?function=GetAliasInfo

### 原型
```
IsAlias(name string) int
```

### 描述

获取给到名称的别名是否存在

* name 需要判断的别名名称

### 代码范例

Javascript:
```
world.note(world.IsAlias("myalias"));
```

Lua:
```
Note(IsAlias("myalias"))
```

### 返回值

* eAliasNotFound 别名不存在
* eOK 别名存在

## SetAliasOption

设置别名选项

不完全兼容Mushclient

对应Mushclient API:https://www.gammon.com.au/scripts/doc.php?function=SetAliasOption

### 原型
```
SetAliasOption(name string, option string, value string) int
```

### 描述

设置脚本别名的选项值

* name 别名名称
* option 选项名称
* value 选项值

可用的option值：


* "echo_alias": 废弃
* "enabled": y/n - 别名是否有效
* "expand_variables": y/n - 展开变量
* "group": (string - 别名组名)
* "ignore_case": y/n - 是否无视大小写
* "keep_evaluating": y/n - 是否继续执行
* "match": (string - 匹配文字)
* "menu":废弃
* "omit_from_command_history": 废弃
* "omit_from_log": 废弃
* "omit_from_output": y/n - 是否在输出中屏蔽
* "one_shot": y/n - 是否是一次性别名
* "regexp": y/n - 是否是正则表达式
* "script": (string - 脚本函数名)
* "send": (multi-line string - 发送的内容)
* "send_to": 0 - 13 - 发送到的位置
* "sequence": 0 - 10000 - 优先级
* "user":废弃
* "variable": (string - 发送到的变量名)

对于数字型的值，传入的字符串应该能转换为数字

对于布尔型的值，传入的值必须是
* "y", "Y", or "1" 为真值
* "n", "N", or "0" 为假值

### 返回值

* eAliasCannotBeEmpty 匹配文本不能为空
* eOK 成功

