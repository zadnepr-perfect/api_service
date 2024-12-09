package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

// GracefulShutdown обеспечивает корректное завершение работы сервера
func GracefulShutdown(e *echo.Echo, logger *log.Logger) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	<-sigchan
	logger.Println("Gracefully shutting down...")

	// Настройка таймера для завершения обработки запросов
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("Error shutting down:", err)
	}

	logger.Println("Server shut down successfully")
}
