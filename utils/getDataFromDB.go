package utils

import (
	"nats-testing/initializers"
	"nats-testing/models"
)

// Достаем данные из БД и пишем в мапу (кэш)

func GetAllDataFromDB() map[string]models.Order {
	cache := make(map[string]models.Order)
	var result []models.Order
	err := initializers.DB.Preload("Delivery").Preload("Payment").Preload("Items").Find(&result).Error
	if err != nil {
		// Обработка ошибки
	}

	for _, order := range result {
		cache[order.OrderUid] = order
	}

	return cache
}
