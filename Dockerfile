# Используем официальный образ Golang для сборки
FROM golang:1.23.3

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем зависимости
RUN go install github.com/air-verse/air@latest

# Открываем порт для приложения
EXPOSE 8080

# Запускаем приложение с автообновлением
CMD ["air"]