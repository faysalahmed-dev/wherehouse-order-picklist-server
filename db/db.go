package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent"
	_ "github.com/lib/pq"
)

var DBClient *ent.Client

func ConnectToDB() *ent.Client {
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	PASSWORD := os.Getenv("DB_PASSWORD")

	dbURL := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", HOST, PORT, USER, DB_NAME, PASSWORD)

	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	DBClient = client
	fmt.Println("database connected...")
	return client
}
