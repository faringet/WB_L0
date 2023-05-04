package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"nats-testing/models"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Connect Options.
	//opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}

	// Connect to NATS
	url := "nats://192.168.99.100:4222"
	nc, err := nats.Connect(url) // потом поменяю на nats.Connect(stan.DefaultNatsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Старый код
	//sc, err := stan.Connect("test-cluster", "stan-pub", stan.NatsConn(nc))
	//if err != nil {
	//	log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	//}
	//defer sc.Close()
	//

	var order models.Order
	err = json.Unmarshal([]byte(models.ModelStr), &order)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// Генерируем случайное значение orderUid
	tmp := []byte(order.OrderUid)
	for i := 0; i < len(order.OrderUid)-4; i++ {
		if tmp[i] > 0x30 && tmp[i] < 0x39 {
			tmp[i] = byte(0x30 + rand.Intn(9))
		} else {
			tmp[i] = byte(0x61 + rand.Intn(6))
		}
	}
	order.OrderUid = string(tmp)
	fmt.Println("result orderUid:", order.OrderUid)

	// Генерируем случайное значение Transaction
	tmp = []byte(order.Payment.Transaction)
	for i := 0; i < len(order.OrderUid)-4; i++ {
		if tmp[i] > 0x30 && tmp[i] < 0x39 {
			tmp[i] = byte(0x30 + rand.Intn(9))
		} else {
			tmp[i] = byte(0x61 + rand.Intn(6))
		}
	}
	order.Payment.Transaction = string(tmp)
	fmt.Println("result Transaction:", order.Payment.Transaction)

	// Генерируем случайное значение Rid
	tmp = []byte(order.Items[0].Rid)
	for i := 0; i < len(order.Items[0].Rid)-4; i++ {
		if tmp[i] > 0x30 && tmp[i] < 0x39 {
			tmp[i] = byte(0x30 + rand.Intn(9))
		} else {
			tmp[i] = byte(0x61 + rand.Intn(6))
		}
	}
	order.Items[0].Rid = string(tmp)
	fmt.Println("result Rid:", order.Items[0].Rid)

	msg, err := json.MarshalIndent(order, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		err = nc.Publish("event.old", msg)
		fmt.Println(string(msg))
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
	}
}
