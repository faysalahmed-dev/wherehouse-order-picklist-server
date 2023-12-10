package db

import (
	"fmt"
	"log"
	"os"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

type Store struct {
	User        store.UserStore
	Category    store.CategoryStore
	SubCategory store.SubCategoryStore
	Product     store.ProductStore
	Order       store.OrderStore
}

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
	db.AutoMigrate(&schema.User{})
	db.AutoMigrate(&schema.Category{})
	db.AutoMigrate(&schema.SubCategory{})
	db.AutoMigrate(&schema.Product{})
	db.AutoMigrate(&schema.Order{})
	DBClient = db
	fmt.Println("database connected...")
	return db
}
