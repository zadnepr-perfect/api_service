package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// ResponseRecorder — это обертка для ответа, чтобы перехватывать тело ответа.
type ResponseRecorder struct {
	http.ResponseWriter // Используем стандартный http.ResponseWriter
	body                *bytes.Buffer
	status              int
}

// Write записывает данные в буфер
func (r *ResponseRecorder) Write(p []byte) (n int, err error) {
	log.Printf("Writing response body: %s", string(p)) // Логируем тело ответа
	n, err = r.body.Write(p)
	return n, err
}

// Header возвращает заголовки ответа
func (r *ResponseRecorder) Header() http.Header {
	return r.ResponseWriter.Header()
}

// WriteHeader записывает статус ответа
func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// Status возвращает статус ответа
func (r *ResponseRecorder) Status() int {
	return r.status
}

// Body возвращает тело ответа
func (r *ResponseRecorder) Body() io.Reader {
	return r.body
}
