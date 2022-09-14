package models

import (
	//"encoding/json"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type OrderRequest struct {
	gorm.Model
	// Receipt  int64  `json:"receipt" gorm:"primaryKey`
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
}

type OrderResponse struct {
	gorm.Model
	Id       uint64 `db:"id" json:"id"`
	Amount   uint64 `db:"amount" json:"amount"`
	Currency string `db:"currency" json:"currency"`
}

// type RazorpayOrder struct {
// 	Items []RazorpayOrderItem `json:"items"`
// }

type RazorpayOrderItem struct {
	Id          uint64 `db:"id"`
	OrderId     string `db:"order_id" json:"id"`
	Amount      uint64 `db:"amount" json:"amount"`
	Amount_Due  uint64 `db:"amount_due" json:"amount_due"`
	Amount_Paid uint64 `db:"amount_paid" json:"amount_paid"`
	Currency    string `db:"currency" json:"currency"`
	Receipt     string `db:"receipt" json:"receipt"`
	Status      string `db:"status" json:"status"`
	Attempts    uint64 `db:"attempts" json:"attempts"`
	Created_At  uint64 `db:"created_at" json:"created_at"`
}
