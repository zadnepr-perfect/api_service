package middleware

import (
	"api/pkg/loggingdb"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CounterMiddleware добавляет поле "counter" в каждый JSON-ответ
func CounterMiddleware(loggingModel *loggingdb.RequestLogsModel) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Получаем контекст запроса
			ctx := c.Request().Context()

			// Получаем количество строк в таблице request_logs
			rowCount, err := loggingModel.GetRowCount(ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get row count: "+err.Error())
			}

			// Перехватываем оригинальный writer для изменения тела ответа
			rec := &ResponseRecorder{
				ResponseWriter: c.Response().Writer, // Используем http.ResponseWriter из Echo
				body:           new(bytes.Buffer),   // Буфер для тела ответа
			}
			c.Response().Writer = rec // Устанавливаем наш ResponseRecorder как Writer

			// Продолжаем выполнение запроса
			err = next(c)
			if err != nil {
				return err
			}

			// Проверяем, что статус ответа равен 200 OK перед добавлением counter
			if c.Response().Status == http.StatusOK {
				// Читаем тело ответа из буфера
				log.Println("Before modification, response body:", rec.body.String())

				var responseBody map[string]interface{}
				if err := json.Unmarshal(rec.body.Bytes(), &responseBody); err != nil {
					log.Println("Failed to parse response body:", err)
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to parse response body")
				}

				// Добавляем поле "counter" в ответ
				responseBody["counter"] = rowCount

				// Перезаписываем тело ответа с добавленным полем
				finalResponse, err := json.Marshal(responseBody)
				if err != nil {
					log.Println("Failed to marshal modified response:", err)
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal response body")
				}

				// Логируем модифицированный ответ
				log.Println("Modified response body with counter:", string(finalResponse))

				// Устанавливаем правильный Content-Length
				c.Response().Header().Set("Content-Length", string(len(finalResponse)))

				// Устанавливаем статус ответа
				c.Response().WriteHeader(c.Response().Status)

				// Отправляем новый ответ с добавленным полем
				_, err = c.Response().Write(finalResponse)
				if err != nil {
					log.Println("Failed to send response:", err)
					return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send response")
				}
			}

			return nil
		}
	}
}
