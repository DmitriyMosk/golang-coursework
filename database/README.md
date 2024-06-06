ConfigMap, StatefulSet, Service
kubectl apply -f pg-config.yaml
kubectl apply -f pg-statefulset.yaml
kubectl apply -f pg-service.yaml

Применение правил Deployment и Service для pgAdmin
kubectl apply -f pgadmin-deployment.yaml
kubectl apply -f pgadmin-service.yaml