# Используем официальный образ Golang для сборки
FROM golang:1.23.3

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем зависимости
RUN go install github.com/air-verse/air@latest

# Устанавливаем bash и wget
RUN apt-get update && apt-get install -y bash wget

# Скачиваем wait-for-it с исходного репозитория GitHub
RUN wget https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh -O /usr/local/bin/wait-for-it.sh && \
    chmod +x /usr/local/bin/wait-for-it.sh

# Открываем порт для приложения
EXPOSE 8080

# Запускаем приложение с автообновлением
CMD ["air"]