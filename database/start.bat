@echo off

echo "Starting Docker Compose..."
docker compose up -d --build

echo "Applying Kubernetes manifests..."
kubectl apply -f kmanifests/pg-config.yaml
kubectl apply -f kmanifests/pg-statefulset.yaml
kubectl apply -f kmanifests/pg-service.yaml
kubectl apply -f kmanifests/pgadmin-deployment.yaml
kubectl apply -f kmanifests/pgadmin-service.yaml

echo "Deployment completed."
pause