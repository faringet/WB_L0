package utils

import (
	"log"
	"nats-testing/initializers"
	"nats-testing/models"
	"time"
)

func Migrate() {

	err := initializers.DB.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Items{})
	if err != nil {
		return
	}
	time.Sleep(3 * time.Second)
	log.Println("Start migration")

}
