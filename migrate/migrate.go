package main

import (
	"log"
	"nats-testing/initializers"
	"nats-testing/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Items{})
	if err != nil {
		return
	}
	log.Println("Start migration")

}
