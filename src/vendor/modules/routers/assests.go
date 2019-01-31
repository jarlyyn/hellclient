package routers

import "github.com/herb-go/herb/middleware"

//AssestsMiddlewares middlewares that should used in assests requests
var AssestsMiddlewares = func() middleware.Middlewares {
	return middleware.Middlewares{}
}
