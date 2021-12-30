package uiserver

import (
	"fmt"
	"hellclient/modules/app"
	"hellclient/modules/routers"
	"hellclient/modules/userpassword"
	"net/http"

	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/util"
	"github.com/herb-go/util/httpserver"
)

//ModuleName module name
const ModuleName = "900uiserver"

const BasicauthRealm = ""

func PasswordMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username, passowrd := userpassword.Load()
	if username == "" || passowrd == "" {
		username = app.System.Username
		passowrd = app.System.Password
	}
	if username != "" && passowrd != "" {
		u, p, ok := r.BasicAuth()
		if !ok || app.System.Username != u || app.System.Password != p {
			w.Header().Set("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\", charset=\"UTF-8\"", BasicauthRealm))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
	}
	next(w, r)
}

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		a := middleware.New()
		a.Use(
			httpserver.RecoverMiddleware(nil),
			PasswordMiddleware,
		).Handle(routers.RouterUIFactory.CreateRouter())
		s := http.Server{}
		s.Addr = app.System.Addr
		s.Handler = a
		go s.ListenAndServe()
		util.OnQuitAndLogError(s.Close)
		util.InitOrderByName(ModuleName)
	})
}
