#!/bin/bash

echo "Stopping Docker Compose..."
docker compose down

echo "Deleting Kubernetes resources..."
kubectl delete -f kmanifests/pg-config.yaml
kubectl delete -f kmanifests/pg-statefulset.yaml
kubectl delete -f kmanifests/pg-service.yaml
kubectl delete -f kmanifests/pgadmin-deployment.yaml
kubectl delete -f kmanifests/pgadmin-service.yaml

echo "Cleanup completed."