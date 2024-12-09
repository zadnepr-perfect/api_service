#!/bin/bash

# Перезапуск через systemd
echo "Restarting hello-api service..."
systemctl restart hello-api.service

# Проверка статуса
systemctl status hello-api.service