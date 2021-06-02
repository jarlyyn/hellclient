package routers

import (
	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/httprouter"
)

//UIMiddlewares middlewares which should be used on router.
var UIMiddlewares = func() middleware.Middlewares {
	return middleware.Middlewares{}
}

//RouterUIFactory ui router factory.
var RouterUIFactory = router.NewFactory(func() router.Router {
	var Router = httprouter.New()
	//Put your router configure code here
	return Router
})
