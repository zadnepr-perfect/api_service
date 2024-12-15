package api

import (
	"net/http"

	"github.com/zadnepr-perfect/api_service/config"

	"github.com/labstack/echo/v4"
)

// Обработчик запроса на /hello
func HelloHandler(c echo.Context) error {
	message := config.LoadConfig().Message // Получаем сообщение из конфигурации
	return c.JSON(http.StatusOK, map[string]string{
		"message": message,
	})
}
