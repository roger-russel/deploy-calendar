package session

import (
	"flag"

	"github.com/fasthttp/session"
	"github.com/fasthttp/session/memory"
	"github.com/google/logger"
)

//Start Session connection session
func Start(providerName string) {

	var serverSession = session.New(session.NewDefaultConfig())
	var config session.ProviderConfig

	config = &memory.Config{}

	providerNameFlag := flag.String("provider", "memory", providerName)
	if errSession := serverSession.SetProvider(*providerNameFlag, config); errSession != nil {
		logger.Fatalf("Session could not start: %s", errSession)
	}

}
