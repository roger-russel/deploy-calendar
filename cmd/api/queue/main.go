package main

import (
	"log"
	"os"
	"strconv"

	"github.com/roger-russel/deploy-calendar/internal/api/queue/router"

	"github.com/roger-russel/deploy-calendar/pkg/logger"
	"github.com/roger-russel/deploy-calendar/pkg/webserver"
)

func main() {

	initLogger()    // Start logger
	initWebserver() // Start API Webservice

}

func initLogger() {
	logger.Init("api-queue")
}

func initWebserver() {

	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	apiRouter := router.API()

	webserver.Start(apiPort, apiRouter)

}
