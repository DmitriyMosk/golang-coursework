#!/bin/sh

# Устанавливаем переменную для пути к файлу docker-compose
COMPOSE_FILE=../../deploy/docker-compose.yml

# Проверяем аргументы
if [ "$1" = "off" ]; then
    echo "Stopping the project..."
    docker-compose -f "$COMPOSE_FILE" down
    exit 0
fi

echo "Starting the project..."
docker-compose -f "$COMPOSE_FILE" up -d
