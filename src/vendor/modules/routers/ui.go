package routers

import (
	"modules/scripts"

	"github.com/herb-go/herb/file/simplehttpserver"
	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/muxrouter"
	"github.com/herb-go/util"
)

//UIMiddlewares middlewares which should be used on router.
var UIMiddlewares = func() middleware.Middlewares {
	return middleware.Middlewares{}
}

//RouterUIFactory ui router factory.
var RouterUIFactory = router.NewFactory(func() router.Router {
	var Router = muxrouter.New()
	Router.StripPrefix("/public").
		HandleFunc(simplehttpserver.ServeFolder(util.Resources("public")))
	Router.HandleHomepage().HandleFunc(simplehttpserver.ServeFile(util.Resources("defaultui", "index.html")))
	Router.Handle(scripts.ScriptPrefix).Handle(scripts.NewWebdavServer())
	//Put your router configure code here
	return Router
})