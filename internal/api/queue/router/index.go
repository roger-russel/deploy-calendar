package router

import "github.com/fasthttp/router"
import c "github.com/roger-russel/deploy-calendar/internal/api/queue/controller"

//API return router.Roter API routes
func API() *router.Router {
	r := router.New()
	r.GET("/", c.Index)
	return r
}
