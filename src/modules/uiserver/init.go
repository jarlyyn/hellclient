package uiserver

import (
	"fmt"
	"modules/app"
	"modules/routers"
	"modules/userpassword"
	"net/http"
	"time"

	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/util"
	"github.com/herb-go/util/httpserver"
)

// ModuleName module name
const ModuleName = "900uiserver"

const BasicauthRealm = ""

var ReadTimeout = 5 * time.Minute

func PasswordMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username, password := userpassword.Load()
	if username == "" || password == "" {
		username = app.System.Username
		password = app.System.Password
	}
	if username != "" && password != "" {
		u, p, ok := r.BasicAuth()
		if !ok || username != u || password != p {
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
		s.ReadTimeout = ReadTimeout
		s.Addr = app.System.Addr
		s.Handler = a
		go func() {
			err := s.ListenAndServe()
			if err != nil || err != http.ErrServerClosed {
				panic(err)
			}
		}()
		util.OnQuitAndLogError(s.Close)
		util.InitOrderByName(ModuleName)
	})
}
