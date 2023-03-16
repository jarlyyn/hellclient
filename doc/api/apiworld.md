# 游戏接口

## GetWorldById

废弃接口

### 代码原型

```
GetWorldById(WorldID string) interface{}
```

### 描述

已废弃接口，返回空

### 返回值

* 空

## GetWorld

废弃接口

### 代码原型

```
GetWorld(WorldName string)
```

### 描述

已废弃接口，返回空

### 返回值

* 空

## GetWorldID

获取当前游戏ID

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=GetWorldID

### 原型

```
GetWorldID() string
```

### 描述

返回当前游戏的ID

### 代码范例

Javascript
```
world.Note (world. GetWorldID ());
```

Lua
```
Note ( GetWorldID ())
```

### 返回值
游戏ID

## GetWorldList

废弃接口

### 代码原型

```
GetWorldList() []string
```

### 描述

已废弃接口，返回空字符串数组

### 返回值

* 空字符串数组

## GetWorldIdList

废弃接口

### 代码原型

```
GetWorldIdList() []string
```

### 描述

已废弃接口，返回空字符串数组

### 返回值

* 空字符串数组

## WorldName

获取游戏名

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=WorldName

### 原型

```
WorldName() string
```

### 描述

返回当前游戏的游戏名(ID)

返回值和GetWorldID一致

### 代码范例

Javascript
```
world.Note(world.WorldName());
```

Lua
```
Note(WorldName())
```

### 返回值

当前游戏的游戏ID

## WorldAddress

获取游戏网络地址

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=WorldAddress

### 原型

```
WorldAddress() string
```

### 描述

返回游戏地址(创建时填写的服务器网址)

### 代码范例

Javascript
```
world.Note(world.WorldAddress());
```

Lua
```
Note(WorldAddress())
```

### 返回值

服务器地址

## WorldPort

获取游戏网络端口

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=WorldPort

### 原型

```
WorldPort() int
```

### 描述

获取当前游戏连接的服务器端口


### 代码范例

Javascript
```
world.Note(world.WorldPort());
```

Lua
```
Note(WorldPort())
```

### 返回值

游戏的端口号


```
WorldProxy() string
```

### 描述

获取当前游戏使用的代理服务器地址


### 代码范例

Javascript
```
world.Note(world.WorldProxy());
```

Lua
```
Note(WorldProxy())
```

### 返回值

游戏的大理服务器信息

## DeleteGroup

按组删除元素

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=DeleteGroup


### 原型
```
DeleteGroup(group string) int
```

### 描述

按给到的分组删除触发器，计时器，别名

只删除分组内的脚本触发器，计时器，别名

### 代码范例

Javascript:
```
world.DeleteGroup ("groupname");
```
Lua:
```
DeleteGroup ("groupname")
```

### 返回值

删除的元素的数量
