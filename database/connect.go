package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=postgres sslmode=disable password=yash"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	return db
}

func Close() {
	DB.Close()
}
