package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shravanth-drife/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetOrders)
	router.POST("/", controller.CreateOrder)
	router.DELETE("/:id", controller.DeleteOrder)
	router.PUT("/:id", controller.UpdateOrder)
}
