# 地图接口

## 术语

### Mapper

Mapper是每个游戏独立的地图对象,所有的操作都基于这个对象

### room

房间，地图中的每个位置的具体信息

原型为

```
type Room struct {
	ID    string
	Name  string
	Exits []*Path
}
```

* ID为房间ID,房间唯一识别代号，不重复
* Name为房间名，可以重复
* Exits是房间的出口列表

### path

path是路径,代表了房间和房间之间的行走关系

原型为

```
type Path struct {
	Command     string
	Delay       int
	From        string
	To          string
	Tags        map[string]bool
	ExcludeTags map[string]bool
}
```
* Command为路径前进的指令
* Delay是路径前进的延迟，计算路径会优先较短的路径
* From是出发房间
* To是到达房间
* Tags 是路径标签列表，需要Mapper具有所有这些标签才能进入
* ExclueTags 是黑名单标签列表，需要Mapper不局有所有这些标签才能进入

### step

移动，获取的每次移动的信息，原型为

```
type Step struct {
	To      string
	From    string
	Command string
	Delay   int
}
```
属性与Path里的部分一致
### tag

标签是字符串，往往由于代表游戏的状态，决定路径的可用性

### flylist

飞行列表,就是一个路径列表，用于只带那些从所有房间都可以进入的路径，不依附于特定房间

## Mapper.reset

重置地图

### 原型

```
Reset()
```

### 描述

重置当前地图

### 代码范例

Javascript
```
Mapper.reset()
```

Lua
```
Mapper:reset()
```

### 返回值

无

## Mapper.addtags

添加标签

### 原型
```
AddTags(tagnames []string)
```

### 描述

为Mapper追加指定的标签列表

### 代码范例

Javascript
```
Mapper.addtags("tag1","tag2","tag3")
```

Lua
```
Mapper:addtags("tag1","tag2","tag3")
```

### 返回值

无

## Mapper.settag

设置标签

### 原型

```
SetTag(tagname string, enabled bool)
```

### 描述

设置指定标签的状态

* tagname 标签名
* enabled 是否有效

### 范例代码

Javascript
```
Mapper.settag("mytag",true)
```

Lua
```
Mapper:settag("mytag",true)
```

### 返回值

无

## Mapper.settags

设置标签列表

### 原型
```
SetTags(tagnames []string)
```

### 描述

为Mapper添加指定的标签列表

### 范例代码

Javascript
```
Mapper.settags("tag1","tag2","tag3")
```

Lua
```
Mapper:settags("tag1","tag2","tag3")
```

### 返回值

无

### Mapper.flashtags

清理标签

### 原型

```
FlashTags()
```

### 描述

清理当前游戏地图的所有标签

### 范例代码

Javascript
```
Mapper.flashtags()
```
Lua
```
Mapper:flashtags()
```

### 返回值

无

## Mapper.tags

获取标签

### 原型
```
Tags() []string
```

### 描述

返回当前游戏的所有标签

### 代码范例

Javascript
```
world.Note(Mapper.Tags())
```

Lua
```
Note(Mapper:Tags)
```

### 返回值

字符串列表

## Mapper.getpath

获取路径

### 原型
```
GetPath(from string, fly bool, to []string) []*Step
```
### 描述

获取路径

* from 起点位置
* fly 是否使用flylist(0为false)
* to 重点列表

返回值为Step对象的列表

找不到路径返回空

### 代码范例

Javascript
```
var path=Mapper.GetPath("from",true,["to1","to2])
path.forEach(function (step) {
    world.Note(step.from) //起点
    world.Note(step.to) //目的地
    world.Note(step.delay) //延迟
    world.Note(step.command) //指令
})
```

Lua
```
local path=Mapper:GetPath("from",true,{"to1","to2"})
for k, step in pairs(path) do
    Note(step.from) -起点
    Note(step.to) -目的地
    Note(step.delay) -延迟
    Note(step.command) -指令
end
```

### 返回值

见描述

## Mapper.addpath

添加路径

### 原型

```
AddPath(id string, path *Path) bool
```

### 描述

将路径添加到指定的房间里

* id 房间id，房间必须存在
* path 路径对象，用Mapper.newpath创建

### 原型

Javascript
```
```

Lua
```
```

### 返回值

房间是否存在

## Mapper.newpath

新建路径

### 原型

```
NewPath() *Path
```

### 描述

创建一个新路径

### 代码范例

Javascript
```
var path=Mapper.newpath()
path.command="command"
path.from="form"
path.to="to"
path.delay=10
path.tags=["tag1","tag2","tag3"]
path.excludetags=["extag1","extag2"]
```

Lua
```
local path=Mapper:newpath()
path.command="command"
path.from="form"
path.to="to"
path.delay=10
path.tags={"tag1","tag2","tag3"}
path.excludetags={"extag1","extag2"}
```

### 返回值

path 对象

## Mapper.getroomid

获取房间id

### 原型
```
GetRoomID(name string) []string
```

### 描述
根据给到的房间名寻找房间ID

* name 房间名

### 范例代码

Javascript
```
var rooms=Mapper.GetRoomID("房间")
```

Lua
```
local rooms=Mapper:GetRoomID("房间")
```

### 返回值

房间名匹配的房间ID列表

## Mapper.getroomname

获取房间名

### 原型
```
GetRoomName(id string) string
```

### 描述

返回指定房间名的房间id

如果房间名未找到，返回空字符串

### 代码范例

Javascript
```
world.Note(Mapper.GetRoomName("房间名"))
```

Lua
```
Note(Mapper:GetRoomName("房间名"))
```

### 返回值

空

## Mapper.setroomname

设置房间

### 原型
```
SetRoomName(id string, name string)
```

### 描述

将指定id的房间的房间名设置为name

注意，若id不存在，将自动建立房间

### 代码范例
Javascript
```
Mapper.SetRoomName("start","开始")
```

Lua
```
Mapper:SetRoomName("start","开始")
```

### 返回值

无

## Mapper.clearroom

清理房间

### 原型

```
ClearRoom(id string)
```

### 描述

清理指定id的房间，重置房间为新房间

### 代码范例

Javascript
```
Mapper.clearroom("new")
```

Lua
```
Mapper:clearroom("new")
```

### 返回值

无

## Mapper.newarea 

新建区域

### 原型

```
NewArea(size int) []string
```

### 描述

新建一块指定大小的房间，并返回房间ID列表

一般用于创建临时区域

### 代码范例

Javascript
```
var rooms=Mapper.newarea(10)

rooms.forEach(function (id){
    world.Note(id)
})
```

Lua
```
local rooms=Mapper:newarea(10)
for k,id in pairs(rooms) do
    Note(id)
end
```

### 返回值

字符串列表

## Mapper.getexits

获取房间出口

### 原型
```
GetExits(id string, all bool) []*Path
```

### 描述

返回指定房间的出口列表

* id 房间id
* all 是否返回全部出口，为false则判断地图的标签属性进行过滤

如果房间未找到，返回空数组

### 代码范例

Javascript
```
var exits=Mapper.GetExits("room",true)

exits.forEach(function (exit){
    world.Note(exit.command) //出口命令
    world.Note(exit.delay) //出口延迟
    world.Note(exit.from) //出口出发房间
    world.Note(exit.to) //出口到达房间
    exit.tags.forEach(function(tag){
        world.Note(tag) //必须的标签
    })
    exit.excludetags(function(tag){
        world.Note(tag) //排除的标签
    })
})
```

Lua
```
local exits=Mapper:GetExits("room",true)
for k, exit in pairs(exits) do
    Note(exit.command) //出口命令
    Note(exit.delay) //出口延迟
    Note(exit.from) //出口出发房间
    Note(exit.to) //出口到达房间
    for k, tag in pairs(exit.tags) do
        Note(tag) //必须的标签
    end
    for k, tag in pairs(exit.excludetags) do
        Note(tag) //排除的标签
    end
end
```

### 返回值

path对象列表

## Mapper.flylist

获取飞行列表

### 原型
```
FlyList() []*Path
```

### 描述

获取当前地图的flylist

### 代码范例

Javascript
```
var flyList=Mapper.flyList()

flyList.forEach(function (fly){
    world.Note(fly.command) //出口命令
    world.Note(fly.delay) //出口延迟
    world.Note(fly.to) //出口到达房间
    fly.tags.forEach(function(tag){
        world.Note(tag) //必须的标签
    })
    fly.excludetags(function(tag){
        world.Note(tag) //排除的标签
    })
})
```

Lua
```
local flyList=Mapper:flyList()
for k, fly in pairs(flyList) do
    Note(fly.command) //出口命令
    Note(fly.delay) //出口延迟
    Note(fly.to) //出口到达房间
    for k, tag in pairs(fly.tags) do
        Note(tag) //必须的标签
    end
    for k, tag in pairs(fly.excludetags) do
        Note(tag) //排除的标签
    end
end
```

### 返回值

path对象列表

## Mapper.setflylist

设置飞行列表

### 原型
```
SetFlyList(fly []*Path)
```

### 描述

设置地图的飞行列表

* fly path数组，from属性将被忽略，应该通过Mapper.newpath创建

### 代码范例

Javascript
```
Mapper.setflylist(list)
```

Lua
```
Mapper:setflylist(list)
```

### 返回值

无