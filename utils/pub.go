package utils

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"nats-testing/models"
	"time"
)

func Pub() {

	rand.Seed(time.Now().UnixNano())

	// Коннектимся к NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	var order models.Order
	err = json.Unmarshal([]byte(models.ModelStr), &order)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// Циклом генерим кол-во json'ов
	for i := 0; i < 5; i++ {

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
			//fmt.Println(string(msg))
			if err != nil {
				log.Fatalf("Error during publish: %v\n", err)
			}
		}
	}

}
