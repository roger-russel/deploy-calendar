package main

import "github.com/roger-russel/deploy-calendar/pkg/webserver"

func main() {

	webserver.Start(8080, false)

}
