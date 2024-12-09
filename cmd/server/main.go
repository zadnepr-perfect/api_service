package main

import (
	"github.com/labstack/echo/v4"

	"api/internal/api"
	"api/internal/database"
	"api/internal/logging"
	"api/internal/shutdown"
)

func main() {
	// Инициализация логирования
	logger := logging.NewLogger()

	// Инициализация базы данных
	database.InitDatabase()
	defer database.CloseDatabase()

	// Инициализация маршрутизатора
	e := echo.New()

	// Регистрация обработчика для endpoint /hello
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
