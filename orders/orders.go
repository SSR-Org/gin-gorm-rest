package orders

import (
	"encoding/json"
	"log"

	"github.com/razorpay/razorpay-go"
	"github.com/shravanth-drife/gin-gorm-rest/controller"
	"github.com/shravanth-drife/gin-gorm-rest/models"
)

func SendOrder(data map[string]interface{}) {
	client := razorpay.NewClient("rzp_test_7MB1e7rfTp35VG", "YHDKuRvW0xvBvJb706mYfgDZ")

	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(body) //map

	jsonString, err := json.Marshal(body) //byte slice
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(jsonString)) // JSON string

	// convert json to struct
	orderJson := models.RazorpayOrderItem{}
	json.Unmarshal(jsonString, &orderJson)
	//fmt.Println(orderJson)

	controller.CreateOrderDirect(orderJson)

}
