@echo off
setlocal

set COMPOSE_FILE=..\..\deploy\docker-compose.yml

if "%1"=="off" (
    echo Stopping the project...
    docker-compose -f %COMPOSE_FILE% down
    exit /b
)

echo Starting the project...
docker-compose -f %COMPOSE_FILE% up -d --build

endlocal
