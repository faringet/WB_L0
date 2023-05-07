package main

import (
	"github.com/gin-gonic/gin"
	"nats-testing/initializers"

	"nats-testing/utils"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	utils.Migrate()
	go utils.Sub()
	utils.Pub()

	router := gin.Default()

	// Устанавливаем путь к папке с шаблонами (views)
	router.LoadHTMLGlob("views/*")

	// Отображение главной
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Обработки введенного ID от юзера и его отображения
	router.POST("/process", func(c *gin.Context) {
		id := c.PostForm("id")

		// Дополнительная логика
		if id == "" {
			// Если ID пустое то выбрасываем ошибку
			c.HTML(400, "errorIdNull.html", gin.H{"message": "Введите ID"})
			return
		}

		// мапа из которой будем по id доставать данные
		openCache := utils.GetAllDataFromDB()

		// достаем по order_uid (будет вводиться юзером через браузер)
		m, ok := openCache[id]

		// Проверка есть ли такой UID во всей мапе
		if ok {
			// Отправить ID на страницу результатов
			c.HTML(200, "result.html", gin.H{"id": m})
		} else {
			// Ошибка если такого id нет
			c.HTML(400, "errorIdNotExist.html", gin.H{"message": "Нет такого ID"})
		}

	})

	router.Run()
}
