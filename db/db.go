package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func InitDB() (*sql.DB, error) {
    err := godotenv.Load() // Load .env file
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    connStr := os.Getenv("DATABASE_URL") // Get the connection string from the .env file
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    return db, nil
}