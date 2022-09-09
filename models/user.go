package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      int    `json:"id"gorm:"primary_key`
	RiderId string `json:"riderId"`
	Amount  string `json:"amount"`
}
