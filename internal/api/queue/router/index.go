package router

import "github.com/fasthttp/router"
import ctrDeploy "github.com/roger-russel/deploy-calendar/internal/api/queue/controller/deploy"

//API return router.Roter API routes
func API() *router.Router {
	r := router.New()
	r.POST("/deploy", ctrDeploy.Create)
	return r
}
