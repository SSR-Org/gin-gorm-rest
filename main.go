package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
	"github.com/shravanth-drife/gin-gorm-rest/config"
	"github.com/shravanth-drife/gin-gorm-rest/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New() //create new object
	config.Connect()
	routes.UserRoute(router)

	client := razorpay.NewClient("rzp_test_7MB1e7rfTp35VG", "YHDKuRvW0xvBvJb706mYfgDZ")
	data := map[string]interface{}{
		"amount":          50000,
		"currency":        "INR",
		"receipt":         "some_receipt_id",
		"partial_payment": false,
	}
	//extraHeader := make(map[string]string)

	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)

	response, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response))

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// b, ok := value.([]byte)
	// if !ok {
	// 	return errors.New("type assertion to []byte failed")
	// }
	// return json.Unmarshal(b, &c)

	// controller.CreateOrderDirect(body)
	// store := body["id"]
	// fmt.Print(store.(string))

}

// func (c *models.RazorpayOrder) Scan(value interface{}) error {

// }
