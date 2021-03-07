package setup

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=hazel dbname=hazel sslmode=disable password=password")

	if err != nil {
		panic("failed to connect to database")
	}

	DB = database
}
