package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
	"log"
	"nats-testing/initializers"
	"nats-testing/models"
	"time"
)

func main() {
	url := "nats://192.168.99.100:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	count := 0
	sub, _ := nc.Subscribe("event.old", func(msg *nats.Msg) {
		count++
		fmt.Printf("message recieved on subject: %v, data: %v", msg.Subject, string(msg.Data))
		fmt.Println("__________________________________")

		// Из NATS переводим в структуру для дальнейшего извлечения данных
		var dataFromNATS models.Order
		err := json.Unmarshal(msg.Data, &dataFromNATS)
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println(dataFromNATS)
		fmt.Println(dataFromNATS.OrderUid)

		orderDB := models.Order{
			Model:       gorm.Model{},
			OrderUid:    dataFromNATS.OrderUid,
			TrackNumber: dataFromNATS.TrackNumber,
			Entry:       dataFromNATS.Entry,
			//Delivery:          models.Delivery{},
			//Payment:           models.Payment{},
			Items:             dataFromNATS.Items,
			Locale:            dataFromNATS.Locale,
			InternalSignature: dataFromNATS.InternalSignature,
			CustomerId:        dataFromNATS.CustomerId,
			DeliveryService:   dataFromNATS.DeliveryService,
			Shardkey:          dataFromNATS.Shardkey,
			SmId:              dataFromNATS.SmId,
			DateCreated:       dataFromNATS.DateCreated,
			OofShard:          dataFromNATS.OofShard,
		}

		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println(orderDB)
		initializers.DB.Create(orderDB)

	})

	defer sub.Unsubscribe()

	for {
		old := count
		time.Sleep(10 * time.Second) //если за 10 секунд не получили сообщения из NATS - вырубаемся
		if old == count {
			break
		}
	}

	fmt.Printf("processed %v messages\n", count)
}
