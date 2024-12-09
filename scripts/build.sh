#!/bin/bash

# Компиляция Go приложения
echo "Building Go application..."
go build -o hello-api ./cmd/server

# Создание Docker образа
echo "Building Docker image..."
docker build -t hello-api:latest .

# Примечание: добавить дополнительные шаги по сборке для CI/CD