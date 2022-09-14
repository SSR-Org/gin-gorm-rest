package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shravanth-drife/gin-gorm-rest/config"
	"github.com/shravanth-drife/gin-gorm-rest/models"
)

func CreateOrderDirect(orderJson models.RazorpayOrderItem) {
	config.DB.Create(&orderJson)
}

func GetOrders(c *gin.Context) {
	orders := []models.RazorpayOrderItem{}
	config.DB.Find(&orders)
	c.JSON(200, &orders)
}

func CreateOrder(c *gin.Context) {
	var order models.RazorpayOrderItem
	c.BindJSON(&order)
	config.DB.Create(&order)
	c.JSON(200, &order)
}

func DeleteOrder(c *gin.Context) {
	var order models.RazorpayOrderItem
	config.DB.Where("id = ?", c.Param("id")).Delete(&order)
	c.JSON(200, &order)
}

func UpdateOrder(c *gin.Context) {
	var order models.RazorpayOrderItem
	config.DB.Where("id =?", c.Param("id")).First(&order)
	c.BindJSON(&order)
	config.DB.Create(&order)
	c.JSON(200, &order)
}
