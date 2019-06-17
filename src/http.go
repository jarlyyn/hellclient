package main

import (
	"modules/app"
	"modules/middlewares"
	"modules/routers"

	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/util"
	"github.com/herb-go/util/httpserver"
)

//App Main applactions. to serve http
var App = middleware.New()

var AppMiddlewares = func() middleware.Middlewares {
	return middleware.Middlewares{
		app.HTTP.Forwarded.ServeMiddleware,
		middlewares.MiddlewareErrorPage.ServeMiddleware,
		httpserver.RecoverMiddleware(nil),
		app.HTTP.Headers.ServeMiddleware,
	}
}

//RunHTTP run app as http server
var RunHTTP = func() {
	defer util.Quit()
	var Server = app.HTTP.Config.Server()
	httpserver.MustListenAndServeHTTP(Server, app.HTTP.Config, App)
	util.WaitingQuit()
	defer util.Bye()
	httpserver.ShutdownHTTP(Server)

}

func init() {

	util.RegisterModule("999HTTP", func() {
		App.
			Use(AppMiddlewares()...).
			Handle(routers.New())
	})
}
