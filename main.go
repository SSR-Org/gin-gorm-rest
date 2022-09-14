package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shravanth-drife/gin-gorm-rest/config"
	"github.com/shravanth-drife/gin-gorm-rest/orders"
	"github.com/shravanth-drife/gin-gorm-rest/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New() //create new object
	config.Connect()
	routes.UserRoute(router)

	data := map[string]interface{}{
		"amount":          49900,
		"currency":        "INR",
		"receipt":         "some_receipt_id",
		"partial_payment": false,
	}

	orders.SendOrder(data)

	router.Run(":8080")

}
