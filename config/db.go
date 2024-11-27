package config

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ConnectDB() *sqlx.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=users sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected")
	return db
}
