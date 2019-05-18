package webserver

import (
	"flag"
	"fmt"

	"github.com/fasthttp/router"
	"github.com/roger-russel/deploy-calendar/pkg/logger"
	"github.com/valyala/fasthttp"
)

//Start a fasthttp server
func Start(port int, apiRouter *router.Router) {

	addr := flag.String("addr", fmt.Sprintf(":%d", port), "TCP address to listen to")

	flag.Parse()

	if err := fasthttp.ListenAndServe(*addr, apiRouter.Handler); err != nil {
		logger.Fatal("Error in ListenAndServe: %s", err)
	}
}
