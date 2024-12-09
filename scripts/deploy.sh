#!/bin/bash

# Копирование systemd unit-файла
echo "Deploying systemd service..."
cp docker/systemd/hello-api.service /etc/systemd/system/

# Перезагрузка systemd
echo "Reloading systemd..."
systemctl daemon-reload

# Включение автозапуска сервиса
echo "Enabling hello-api service..."
systemctl enable hello-api.service

# Запуск сервиса
echo "Starting hello-api service..."
systemctl start hello-api.service

# Проверка статуса
systemctl status hello-api.service