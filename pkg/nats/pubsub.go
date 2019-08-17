package nats

import nats "github.com/nats-io/go-nats"

//SubEventDeploy is the const for Subcribe Deploy Event
const SubEventDeploy string = "event.deploy"

func Pub(subject string, data *[]byte) error {

	nc, err := connect(nil)
	defer nc.Close()

	if err != nil {
		return err
	}

	nc.Publish(subject, *data)

	return nil

}

func Sub(subject string, callback func(m *nats.Msg)) {

	nc := connect(nil)
	defer nc.Close()

	nc.Subscribe(subject, callback)

}
