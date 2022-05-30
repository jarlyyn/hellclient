# Mud输出结构体

Mud输出的结构以Line(行)和Word组成

## Line
结构体
```go
type Line struct {
	Words          []Word
	ID             string
	Time           int64
	Type           int
	OmitFromLog    bool
	OmitFromOutput bool
	Triggers       []string
	CreatorType    string
	Creator        string
}
```

字段解释:

* Word 行中构成显示的Word数组
* ID Line的唯一ID,随时间递增
* Time 行创建的Unix时间戳
* Type 行类型列表，具体见下文
* OmitFromLog 是否从日志中排除
* OmitFromOutput 是否从屏幕显示中排除
* Triggers 触发了哪些触发器
* Creator 行的由谁触发
* CreatorType 行的触发类型

Type可用值一览：

```go
//通过Print打印
const LineTypePrint = 0

//系统信息
const LineTypeSystem = 1

//收到的真实信息
const LineTypeReal = 2

//输入回显
const LineTypeEcho = 3

//输入行类型
const LineTypePrompt = 4

//发出的本地广播
const LineTypeLocalBroadcastOut = 5

//发出的全局广播
const LineTypeGlobalBroadcastOut = 6

//收到的本地广播
const LineTypeLocalBroadcastIn = 7

//收到的全局广播
const LineTypeGlobalBroadcastIn = 8

//Websocket发出的请求的信息
const LineTypeRequest = 9

//Websocket收到的响应的信息
const LineTypeResponse = 10
```

## Word

Word指带样式的一串文字，多个Word组成了一个实际的Line

结构:

```go
type Word struct {
	Text       string
	Color      string
	Background string
	Bold       bool
	Underlined bool
	Blinking   bool
	Inverse    bool
}
```

字段解释:

* Text 实际文本
* Color 文字颜色，见下文列表
* Background 背景颜色，见下文列表
* Bold 是否为粗体
* Underlined 是否有下划线 
* Blinking 是否闪烁
* Inverse 是否反转

### Color 列表

```go
	"Black":             0x000000,
	"Red":               0x7f0000,
	"Green":             0x009300,
	"Yellow":            0xfc7f00,
	"Blue":              0x00007f,
	"Magenta":           0x9c009c,
	"Cyan":              0x009393,
	"White":             0xd2d2d2,
	"BrightBlack":       0x7f7f7f,
	"BrightRed":         0xff0000,
	"BrightGreen":       0x00fc00,
	"BrightYellow":      0xffff00,
	"BrightBlue":        0x0000fc,
	"BrightMagenta":     0xff00ff,
	"BrightCyan":        0x00ffff,
	"BrightWhite":       0xffffff,

```
