package webserver

import (
	"flag"
	"fmt"

	"github.com/fasthttp/router"
	"github.com/google/logger"
	"github.com/valyala/fasthttp"
)

//Start a fasthttp server
func Start(port int, apiRouter *router.Router) {

	addr := flag.String("addr", fmt.Sprintf(":%d", port), "TCP address to listen to")

	flag.Parse()

	if err := fasthttp.ListenAndServe(*addr, apiRouter.Handler); err != nil {
		logger.Fatalf("Error in ListenAndServe: %s", err)
	}
}
