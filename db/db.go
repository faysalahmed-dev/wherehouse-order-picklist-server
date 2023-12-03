package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func ConnectToDB() *gorm.DB {
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	PASSWORD := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", HOST, PORT, USER, DB_NAME, PASSWORD)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// db.AutoMigrate(&Product{})
	DBClient = db
	fmt.Println("database connected...")
	return db
}
