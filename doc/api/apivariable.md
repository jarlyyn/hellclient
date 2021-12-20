# 变量接口

## GetVariable

获取变量值

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetVariable

### 原型
```
GetVariable(name string) string
```

### 描述

获取指定的变量值

* name 变量名

注意，由于与mushclient架构不同，空变量会返回空字符串值
### 代码范例
Javascript
```
MyName = world.GetVariable("MyName");
```
Lua
```
MyName = GetVariable("MyName")
```

### 返回值

变量值，空变量返回空字符串值


## SetVariable

设置变量

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=SetVariable

### 原型
```
SetVariable(name string, content string) int
```

### 描述

设置变量值

* name 变量名
* content 变量值

###  范例代码

Javascript
```
world.SetVariable("MyName", "Nick Gammon");
```

Lua
```
SetVariable ("MyName", "Nick Gammon")
```

### 返回值

eOK

## DeleteVariable

删除变量

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteVariable

### 原型
```
DeleteVariable(name string) int
```

### 描述

删除指定的变量

* name 变量名

### 代码范例

Javascript
```
world.DeleteVariable("myvariable");
```

Lua
```
DeleteVariable("myvariable")
```

### 返回值

eOK


## GetVariableList

获取变量列表

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetVariableList

### 原型
```
GetVariableList() map[string]string
```

### 描述 

返回变量名列表

### Lua注意事项

Lua中返回的是个键值对表

### 范例代码 

Javascript
```
variablelist = world.GetVariableList();

 for (i = 0; i < variablelist.length; i++)
   world.note(variablelist [i] + " = " + 
       world.GetVariable(variablelist [i]));

```

Lua
```
for k, v in pairs (GetVariableList()) do 
  Note (k, " = ", v) 
end

```

### 返回值

如描述

## GetVariableComment

获取变量备注

### 原型
```
GetVariableComment(name string) string
```

### 描述

获取指定变量的备注

* name 变量名
* 
### 范例代码

Javascript
```
comment = world.GetVariableComment("myvar")
```

Lua
```
comment = GetVariableComment("myvar")
```

### 返回值

变量备注，未设置则为空字符串


### SetVariableComment

设置变量备注

### 原型
```
SetVariableComment(name string, content string)
```

### 描述

设置指定变量的备注

* name 变量名
* content 备注
### 范例代码

Javascript
```
world.SetVariableComment("myvar"，"comment")
```

Lua
```
SetVariableComment("myvar"，"comment")
```

### 返回值

无