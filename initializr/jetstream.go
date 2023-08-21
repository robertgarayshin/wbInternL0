package initializr

import (
	"github.com/nats-io/nats.go"
	"log"
	"wbInternL0/config"
)

func JetStreamInit() (nats.JetStreamContext, error) {
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	err = CreateStream(js)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func CreateStream(ctx nats.JetStreamContext) error {
	stream, err := ctx.StreamInfo(config.StreamName)

	if stream == nil {
		log.Printf("Creating Stream: %s\n", config.StreamName)

		_, err = ctx.AddStream(&nats.StreamConfig{
			Name:     config.StreamName,
			Subjects: []string{config.StreamSubjects},
		})

		if err != nil {
			return err
		}
	}
	return nil
}
