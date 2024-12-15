# Используем официальный образ Golang для сборки
FROM golang:1.23.3

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы конфигурации Air
COPY .air.toml /app/.air.toml

# Устанавливаем инструмент air (для hot-reload)
RUN go install github.com/cosmtrek/air@v1.28.0

# Устанавливаем bash и wget
RUN apt-get update && apt-get install -y bash wget

# Открываем порт для приложения
EXPOSE 8080

# Запускаем приложение с автообновлением через Air
CMD ["air"]
