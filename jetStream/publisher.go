package jetStream

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"os"
	"time"
	"wbInternL0/config"
	"wbInternL0/models"
)

func publish(js nats.JetStreamContext) {
	orders, err := getOrders()
	if err != nil {
		log.Println(err)
		return
	}

	r := rand.Intn(1500)
	time.Sleep(time.Duration(r) * time.Millisecond)
	orderString, err := json.Marshal(orders)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = js.Publish(config.SubjectNameCreated, orderString)
	if err != nil {
		return
	} else {
		log.Printf("Publisher => Message: %s\n", orders.OrderUid)

	}

}

func getOrders() (models.Order, error) {
	rawOrders, _ := os.ReadFile("./models/model.json")
	var ordersObj models.Order
	err := json.Unmarshal(rawOrders, &ordersObj)
	return ordersObj, err
}
