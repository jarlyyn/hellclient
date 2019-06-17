package middlewares

import (
	"github.com/herb-go/herb/middleware/errorpage"
	"github.com/herb-go/util"
)

//ModuleName module name used in initing and debuging.
const ModuleName = "200Middleware"

//MiddlewareErrorPage middleware which used to show custom error page by status code.
var MiddlewareErrorPage = errorpage.New()

func init() {
	util.RegisterModule(ModuleName, func() {
		util.InitOrderByName(ModuleName)
	})
}
