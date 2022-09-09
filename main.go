package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shravanth-drife/gin-gorm-rest/config"
	"github.com/shravanth-drife/gin-gorm-rest/routes"
)

func main() {
	router := gin.New() //create new object
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
