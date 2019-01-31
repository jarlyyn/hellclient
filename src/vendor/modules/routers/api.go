package routers

import (
	"modules/middlewares"

	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/httprouter"
)

//APIMiddlewares middlewares that should used in api requests
var APIMiddlewares = func() middleware.Middlewares {
	return middleware.Middlewares{
		middlewares.MiddlewareCsrfVerifyHeader,
		middlewares.MiddlewareErrorPage.MiddlewareDisable,
	}
}

func newAPIRouter() router.Router {
	var Router = httprouter.New()
	//Put your router configure code here
	return Router
}
