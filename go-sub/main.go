package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
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

		//fmt.Println(data)
		fmt.Println(dataFromNATS.OrderUid)

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
