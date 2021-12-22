# 连接接口

## Connect

连接

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Connect

### 原型

```
Connect() int
```

### 描述

连接当前游戏到服务器

### 代码范例

Javascript
```
world.connect();
```

Lua
```
Connect()
```

### 返回值

eOK

## IsConnected

判断接连状态

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=IsConnected

### 原型

```
IsConnected() bool
```

### 描述

判断游戏是否连接到了服务器

### 代码范例

Javascript
```
world.note(world.IsConnected());
```

Lua
```
Note(IsConnected())
```

### 返回值

布尔值

## Disconnect

断开连接

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=Disconnect

### 原型

```
Disconnect() int
```

### 描述

 断开当前游戏到服务器的连接

### 代码范例

Javascript
```
world.disconnect();
```

Lua
```
Disconnect()
```

### 返回值

eOK
