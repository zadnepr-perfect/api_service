package middleware

import (
	"github.com/zadnepr-perfect/shared/pkg/loggingdb"

	"github.com/labstack/echo/v4"
)

// LoggingModelMiddleware добавляет модель loggingModel в контекст
func LoggingModelMiddleware(loggingModel *loggingdb.RequestLogsModel) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Добавляем модель в контекст
			c.Set("loggingModel", loggingModel)
			return next(c)
		}
	}
}
