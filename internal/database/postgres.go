package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// PostgreSQL клиента
var DB *pgx.Conn

// InitDatabase инициализирует соединение с базой данных
func InitDatabase() {
	// Чтение переменных окружения для подключения
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Successfully connected to PostgreSQL")
}

// CloseDatabase закрывает соединение с базой данных
func CloseDatabase() {
	if err := DB.Close(context.Background()); err != nil {
		log.Fatalf("Unable to close database connection: %v\n", err)
	}
	log.Println("Database connection closed")
}
