package main

import (
	"fmt"
	"nats-testing/initializers"
	"nats-testing/utils"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	// мапа из которой будем по id доставать данные
	openCache := utils.GetAllDataFromDB()

	// достаем по order_uid (будет вводиться юзером через браузер)
	m := openCache["dbeaacf8b6bdaa6test"]
	fmt.Println(m)

}
