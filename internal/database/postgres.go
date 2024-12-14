// internal/database/database.go
package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
)

// DatabaseMiddlewareKey ключ для хранения соединения в контексте
const DatabaseMiddlewareKey = "dbConn"

// Database представляет клиент базы данных
type Database struct {
	Conn *pgx.Conn
}

// NewDatabaseConnection создает новое соединение с базой данных
func NewDatabaseConnection() (*Database, error) {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL")
	return &Database{Conn: conn}, nil
}

// Close закрывает соединение с базой данных
func (d *Database) Close() {
	if err := d.Conn.Close(context.Background()); err != nil {
		log.Fatalf("Unable to close database connection: %v\n", err)
	}
	log.Println("Database connection closed")
}

// DatabaseMiddleware добавляет соединение с базой данных в контекст
func DatabaseMiddleware(db *Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DatabaseMiddlewareKey, db.Conn)
			return next(c)
		}
	}
}

// GetDatabaseConnection возвращает соединение из контекста
func GetDatabaseConnection(c echo.Context) *pgx.Conn {
	conn, ok := c.Get(DatabaseMiddlewareKey).(*pgx.Conn)
	if !ok {
		panic("Database connection is not set in the context")
	}
	return conn
}
