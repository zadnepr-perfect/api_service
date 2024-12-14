package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Обработчик запроса на /hello
func IndexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello world! :-))",
	})
}
