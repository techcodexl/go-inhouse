package initialzers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	connectionStr := os.Getenv("DB_URI")
	DB,err = gorm.Open(postgres.Open(connectionStr),&gorm.Config{})

	if err!=nil{
		log.Fatal("Failed to connect to database")
	}



}