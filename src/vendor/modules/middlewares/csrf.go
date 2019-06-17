package middlewares

import (
	"modules/app"

	"github.com/herb-go/herb/middleware/csrf"
	"github.com/herb-go/util"
)

//Csrf csrf module
var Csrf = csrf.New()

//MiddlewareCsrfSetToken middleware which sets csrf token to request.
var MiddlewareCsrfSetToken = Csrf.ServeSetCsrfTokenMiddleware

//MiddlewareCsrfVerifyHeader middleware which verifies csrf header
var MiddlewareCsrfVerifyHeader = Csrf.ServeVerifyHeaderMiddleware

//MiddlewareCsrfVerifyForm middleware which verifies csrf post data
var MiddlewareCsrfVerifyForm = Csrf.ServeVerifyFormMiddleware

func init() {
	util.RegisterInitiator(ModuleName, "csrf", func() {
		util.Must(app.Csrf.ApplyTo(Csrf))
	})
}
