package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func ConnectClientDatabase() *gorm.DB {
	LoadEnv()

	server := os.Getenv("CLIENT_DB_SERVER")
	// port := os.Getenv("CLIENT_DB_PORT") // If port is needed, uncomment this line
	user := os.Getenv("CLIENT_DB_USER")
	password := os.Getenv("CLIENT_DB_PASSWORD")
	database := os.Getenv("CLIENT_DB_NAME")

	// Connection string for SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, password, server, database)

	// Open the database connection with GORM
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Client database: %v", err)
	}

	return db
}
