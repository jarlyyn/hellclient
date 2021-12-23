# 色彩接口

## ColourNameToRGB

色彩名转rgb

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=ColourNameToRGB

### 原型
```
ColourNameToRGB(color string) int
```

### 描述

返回给定色彩名对应的RGB值

* color 色彩名，不区分大小写
* 
color的可选值

* black
* red
* green
* yellow
* blue
* magenta
* cyan
* white

对于未知的色彩名，返回-1

### 代码范例
Javascript
```
world.note world.colourNameToRGB ("black")
```

Lua
```
world.note world.colourNameToRGB ("black")
```

### 返回值

如描述所述
## NormalColour

获取高亮色彩rgb

部分兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=NormalColour

### 原型

```
NormalColour(WhichColour int) int
```

### 描述
 获取给定色彩的高亮rgb

 * WhichColour 1-8的色彩值
WhichColour可以取的值:
* 1: black
* 2: red 
* 3: green
* 4: yellow
* 5: blue
* 6: magenta
* 7: cyan
* 8: white

如果取不回值，返回-1

### 代码范例

Javascript
```
world.note(world.normalcolour(3)); 
```

Lua
```
Note (GetNormalColour (2))  
```
### 返回值

见描述
## BoldColour

获取高亮色彩rgb

部分兼容

对应MushclientAPI:https://www.gammon.com.au/scripts/doc.php?function=BoldColour

### 原型

```
BoldColour(WhichColour int) int
```

### 描述
 获取给定色彩的高亮rgb

 * WhichColour 1-8的色彩值
WhichColour可以取的值:
* 1: black
* 2: red 
* 3: green
* 4: yellow
* 5: blue
* 6: magenta
* 7: cyan
* 8: white

如果取不回值，返回-1

注意，实际返回的颜色和NormalColor一样

### 代码范例

Javascript
```
world.Note(world.boldcolour(2));  
```

Lua
```
Note (GetBoldColour (2))  
```
### 返回值

见描述

