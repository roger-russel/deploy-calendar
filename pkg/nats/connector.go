package nats

import (
	"os"

	n "github.com/nats-io/go-nats"
	"github.com/roger-russel/deploy-calendar/pkg/logger"
)

func connect(conn *string) (*n.Conn, error) {

	connectString := getConnString(conn)

	natsConnection, errConn := n.Connect(connectString)

	if errConn != nil {
		logger.Fatal("Could not connect at NATS", errConn)
	}

	return natsConnection, errConn

}

func getConnString(conn *string) string {
	if conn != nil {
		return *conn
	} else if os.Getenv("NATS_CONN") != "" {
		return os.Getenv("NATS_CONN")
	} else {
		return n.DefaultURL
	}
}
