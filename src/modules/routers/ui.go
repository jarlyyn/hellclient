package routers

import (
	"hellclient/modules/app"
	prophetactions "hellclient/modules/world/prophet/actions"
	"hellclient/modules/world/titan"
	"runtime"
	"time"

	"net/http/pprof"

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
	Router.Handle("/ws").HandleFunc(prophetactions.WebsocketAction)
	Router.Handle(titan.GamePrefix).Handle(titan.NewWebdavServer())
	if app.Development.Profiling {
		runtime.SetBlockProfileRate(int(10 * time.Second))
		runtime.SetMutexProfileFraction(int(10 * time.Second))
		Router.Handle("/debug/pprof/").HandleFunc(pprof.Index)
		Router.Handle("/debug/pprof/cmdline").HandleFunc(pprof.Cmdline)
		Router.Handle("/debug/pprof/profile").HandleFunc(pprof.Profile)
		Router.Handle("/debug/pprof/symbol").HandleFunc(pprof.Symbol)
		Router.Handle("/debug/pprof/trace").HandleFunc(pprof.Trace)
	}
	//Put your router configure code here
	return Router
})
