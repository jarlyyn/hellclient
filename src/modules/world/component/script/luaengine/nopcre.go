//go:build !withpcre
// +build !withpcre

package luaengine

import "github.com/herb-go/herbplugin"

var ModuleRex = herbplugin.CreateModule("rex",
	nil,
	nil,
	nil,
)
