#!/bin/bash

# Environment
# export PGUSER="your_db_user"
# export PGPASSWORD="your_db_password"
# export PGHOST="localhost"
# export PGPORT=5432
# export PGDATABASE="your_db_name"

# # Initialize database
# psql -c 'CREATE DATABASE your_db_name;'

# Run the backend service
cd ../backend/cmd/connector
go build -o connector.out main.go 
./connector.out