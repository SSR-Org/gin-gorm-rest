package config

import (
	"github.com/shravanth-drife/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:aaah@localhost:5432/postgres"), &gorm.Config{}) //"postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.RazorpayOrderItem{})
	DB = db

	// construct a gorp DbMap
	// dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// dbmap.AddTableWithName(Order{}, "orders").SetKeys(true, "Id")
}
