package webserver

import (
	"flag"
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//Start a fasthttp server
func Start(port int, apiRouter *router.Router) {

	addr := flag.String("addr", fmt.Sprintf(":%d", port), "TCP address to listen to")

	flag.Parse()

	if err := fasthttp.ListenAndServe(*addr, apiRouter.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
