package jetStream

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"wbInternL0/cache"
	"wbInternL0/config"
	"wbInternL0/models"
	"wbInternL0/repository/write"
)

func Consumer(ctx context.Context, db *sql.DB, c *cache.Cache) {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	log.Printf("Connected to NATS server on %s", nats.DefaultURL)

	subject := config.SubjectNameCreated
	messages := make(chan *nats.Msg, 1000)

	// we're subscribing to the subject
	// and assigning our channel as reference to receive messages there
	subscription, err := nc.ChanSubscribe(subject, messages)
	if err != nil {
		log.Fatal("Failed to subscribe to subject:", err)
	}

	defer func() {
		subscription.Unsubscribe()
		close(messages)
	}()

	log.Println("Subscribed to", subject)
	var order models.Order
	for {
		select {
		case <-ctx.Done():
			log.Println("Exiting from consumer")
			return
		case msg := <-messages:
			log.Println("Received", string(msg.Data))
			err := json.Unmarshal(msg.Data, &order)
			if err != nil {
				log.Fatal(err)
			}
			err = write.NewOrder(db, order)
			c.Set(order.OrderUid, order)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
