package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	DB, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return DB
}
