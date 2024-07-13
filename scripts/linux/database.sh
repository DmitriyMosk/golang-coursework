#!/bin/bash

# Путь к директории с docker-compose.yaml относительно текущей директории скрипта
COMPOSE_DIR="../../database"

# Функция для запуска контейнеров с помощью docker-compose
start_containers() {
    echo "Запуск Docker контейнеров с использованием docker-compose..."

    cd "$COMPOSE_DIR" || { echo "Не удалось найти директорию $COMPOSE_DIR"; exit 1; }

    echo "Starting Docker Compose..."
    docker compose up -d --build

    echo "Applying Kubernetes manifests..."
    kubectl apply -f kmanifests/pg-config.yaml
    kubectl apply -f kmanifests/pg-statefulset.yaml
    kubectl apply -f kmanifests/pg-service.yaml
    kubectl apply -f kmanifests/pgadmin-deployment.yaml
    kubectl apply -f kmanifests/pgadmin-service.yaml

    echo "Deployment completed."
}

# Функция для остановки контейнеров с помощью docker-compose
stop_containers() {
    echo "Остановка Docker контейнеров с использованием docker-compose..."

    cd "$COMPOSE_DIR" || { echo "Не удалось найти директорию $COMPOSE_DIR"; exit 1; }

    echo "Stopping Docker Compose..."
    docker compose down

    echo "Deleting Kubernetes resources..."
    kubectl delete -f kmanifests/pg-config.yaml
    kubectl delete -f kmanifests/pg-statefulset.yaml
    kubectl delete -f kmanifests/pg-service.yaml
    kubectl delete -f kmanifests/pgadmin-deployment.yaml
    kubectl delete -f kmanifests/pgadmin-service.yaml

    echo "Cleanup completed."
}

# Основная логика скрипта
if [ "$1" == "off" ]; then
  stop_containers
else
  start_containers
fi