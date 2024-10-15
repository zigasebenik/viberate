package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(connectionString string) {
	var err error
	DB, err = pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
