// package orders

// import (
// 	"log"
// 	"net/http"
// 	"reflect"
// 	"github.com/shravanth-drife/gin-gorm-rest/models"
// 	"github.com/gin-gonic/gin"
// )

// var dbmap = connectToDB() //Connect to DB and get the DbMap

// func Ping(c *gin.Context) {} //Check order service health

// func CreateOrder(c *gin.Context) {
// 	var orderReq models.OrderRequest
// 	c.Bind(&orderReq)

// 	order := &models.OrderResponse{
// 		Id: orderReq.Id,
// 		Amount: orderReq.Amount,
// 		Currency: orderReq.Currency,
// 	}

// 	err := dbmap.Insert(order)
// 	if err != nil {
// 		log.Fatal(err) //"New order creation failed in orders table"
// 	}

// 	c.JSON(http.StatusCreated,
// 		gin.H{"status": http.StatusCreated, "message": "Order Created Successfully!", "resourceId": order.Id})
// }