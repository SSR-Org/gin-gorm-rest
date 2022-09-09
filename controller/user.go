package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shravanth-drife/gin-gorm-rest/config"
	"github.com/shravanth-drife/gin-gorm-rest/models"
)

func UserController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}
