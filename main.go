package main

import (
	"ejemplo2/src/view/controller"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	server := gin.New()

	server.POST("items", controller.CreateItem)
	server.GET("items", controller.ItemList)
	server.GET("items/:id", controller.FindItemById)

	server.Run(":8182")

	// item1 := *domain.NewItem(1, "Item 1", 100)
	// item2 := *domain.NewItem(2, "Item 2", 200)
	// item3 := *domain.NewItem(3, "Item 3", 300)

	// order1 := domain.NewOrder(1000)
	// order1.AddItem(item1, 7)
	// order1.AddItem(item3, 1)
	// order1.AddItem(item2, 1)

	// dataPayment := domain.DataPayment{
	// 	Number:    "1234-1234-1234-1234",
	// 	CVV:       123,
	// 	Quota:     2,
	// 	Email:     "aenoc.martinez@gmail.com",
	// 	Bank:      "Davivienda",
	// 	MethodPay: "TC",
	// }

	// var methodPay domain.MethodPay = domain.MethodPayFactory(dataPayment)
	// order1.Pay(methodPay)

}
