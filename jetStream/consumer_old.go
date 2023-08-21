package jetStream

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"wbInternL0/config"
	"wbInternL0/models"
)

func Consume(js nats.JetStreamContext) {
	_, err := js.Subscribe(config.SubjectNameCreated, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		var order models.Order
		err = json.Unmarshal(m.Data, &order)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Consumer  =>  Subject: %s  -  ID: %s  -  Author: %s  -  Rating: %s\n",
			m.Subject, order.OrderUid, order.TrackNumber, order.Entry)
	})
	if err != nil {
		log.Println("Subscribe failed")
		return
	}
}
