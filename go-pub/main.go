package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"time"
)

var (
	rg = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {
	url := "nats://192.168.99.100:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	for i := 0; i < 1e5; i++ {
		s := fmt.Sprintf("Message %v: data: %v\n", i, rg.Intn(10000))

		nc.Publish("event.old", []byte(s))

	}

}
