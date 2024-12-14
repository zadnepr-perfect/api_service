// main.go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/internal/api"
	"api/internal/database"
	"api/internal/logging"
	middlewareLogging "api/internal/middleware"
	"api/internal/shutdown"
	"api/pkg/loggingdb"
)

func main() {
	// Инициализация логирования
	logger := logging.NewLogger()

	// Инициализация базы данных
	db, err := database.NewDatabaseConnection()
	if err != nil {
		logger.Fatal("Failed to initialize database", err)
	}
	defer db.Close()

	// Инициализация модели для работы с request_logs
	loggingModel := loggingdb.NewRequestLogsModel(db.Conn)

	// Инициализация маршрутизатора
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(database.DatabaseMiddleware(db))
	e.Use(middlewareLogging.RequestLoggerMiddleware(loggingModel)) // Подключение middleware для логирования
	e.Use(middlewareLogging.CounterMiddleware(loggingModel))       // Подключение middleware для логирования

	// Регистрация middleware для передачи базы данных в контекст
	e.Use(database.DatabaseMiddleware(db))

	// Регистрация обработчиков
	e.GET("/", api.IndexHandler)
	e.GET("/hello", api.HelloHandler)

	// Запуск сервера с graceful shutdown
	go func() {
		if err := e.Start(":8080"); err != nil {
			logger.Fatal("Error starting server", err)
		}
	}()

	// Настройка graceful shutdown
	shutdown.GracefulShutdown(e, logger)
}
