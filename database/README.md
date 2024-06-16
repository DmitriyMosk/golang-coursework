Start node
docker compose up -d 

ConfigMap, StatefulSet, Service<br>
kubectl apply -f pg-config.yaml<br>
kubectl apply -f pg-statefulset.yaml<br>
kubectl apply -f pg-service.yaml<br>
<br>
Применение правил Deployment и Service для pgAdmin<br>
kubectl apply -f pgadmin-deployment.yaml<br>
kubectl apply -f pgadmin-service.yaml<br>