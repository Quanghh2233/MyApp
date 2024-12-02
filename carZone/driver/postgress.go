package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Waiting for the database start up...")
	time.Sleep(5 * time.Second)

	dtb, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = dtb.Ping()
	if err != nil {
		log.Fatalf("Error Connecting to the database: %v", err)
	}

	db = dtb
	fmt.Println("Successfully Connected to the database")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatalf("Error Closing the database: %v", err)

	}
}
