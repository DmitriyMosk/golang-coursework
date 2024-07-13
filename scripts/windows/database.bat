@echo off
setlocal

rem Путь к директории с docker-compose.yaml относительно текущей директории скрипта
set COMPOSE_DIR=..\..\database

rem Переход к директории
pushd %COMPOSE_DIR%
if %ERRORLEVEL% neq 0 (
    echo Could not find directory %COMPOSE_DIR%
    exit /b 1
)

if "%1" == "off" (
    echo Stopping Docker Compose...
    docker compose down
    
    echo Deleting Kubernetes resources...
    kubectl delete -f kmanifests/pg-config.yaml
    kubectl delete -f kmanifests/pg-statefulset.yaml
    kubectl delete -f kmanifests/pg-service.yaml
    kubectl delete -f kmanifests/pgadmin-deployment.yaml
    kubectl delete -f kmanifests/pgadmin-service.yaml

    echo Cleanup completed.
) else (
    echo Starting Docker Compose...
    docker compose up -d --build
    
    echo Applying Kubernetes manifests...
    kubectl apply -f kmanifests/pg-config.yaml
    kubectl apply -f kmanifests/pg-statefulset.yaml
    kubectl apply -f kmanifests/pg-service.yaml
    kubectl apply -f kmanifests/pgadmin-deployment.yaml
    kubectl apply -f kmanifests/pgadmin-service.yaml

    echo Deployment completed.
)

echo successful
endlocal