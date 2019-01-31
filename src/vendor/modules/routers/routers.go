package routers

import (
	//"modules/actions"
	"modules/app"

	"github.com/herb-go/herb/file/simplehttpserver"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/httprouter"
	"github.com/herb-go/herbgo/util"
)

var AssestsURL = "/public"
var AssestsPath = "/public/"

func New() router.Router {
	var Router = httprouter.New()

	//Only host assests folder if folder exisits
	if app.Assets.URLPrefix != "" {
		Router.StripPrefix(app.Assets.URLPrefix).
			Use(AssestsMiddlewares()...).
			HandleFunc(simplehttpserver.ServeFolder(util.Resource(app.Assets.Location)))
	}
	var RouterAPI = newAPIRouter()
	Router.StripPrefix("/api").
		Use(APIMiddlewares()...).
		Handle(RouterAPI)

	//var RouterHTML = newHTMLRouter()
	//Router.StripPrefix("/page").Use(HTMLMiddlewares()...).Handle(RouterHTML)

	//Router.GET("/").Use(HTMLMiddlewares()...).HandleFunc(actions.IndexAction)

	return Router
}
