package session

import (
	"flag"
	"log"

	"github.com/fasthttp/session"
	"github.com/fasthttp/session/memory"
)

//Start Session connection session
func Start() {

	var serverSession = session.New(session.NewDefaultConfig())
	var config session.ProviderConfig

	config = &memory.Config{}

	providerName := flag.String("provider", "memory", "Name of provider")
	if errSession := serverSession.SetProvider(*providerName, config); errSession != nil {
		log.Fatalf("Session could not start: %s", errSession)
	}

}
