package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dns := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns))
	if err != nil {
		log.Panic("Error to connect database")
	}
}
