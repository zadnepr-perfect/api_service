package middleware

import (
	"api/pkg/loggingdb"
	"log"

	"github.com/labstack/echo/v4"
)

func RequestLoggerMiddleware(model *loggingdb.RequestLogsModel) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			// Получение информации о запросе
			endpoint := c.Path()
			status := c.Response().Status
			request := c.Request().URL.String()
			response := c.Response().Header().Get("Content-Type") // Пример ответа
			ip := c.RealIP()

			// Логируем запрос
			if logErr := model.LogRequest(c.Request().Context(), endpoint, status, request, response, ip); logErr != nil {
				log.Printf("Failed to log request: %v", logErr)
			}

			return err
		}
	}
}
