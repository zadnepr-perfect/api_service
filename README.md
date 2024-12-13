# Разработка и локальный запуск:

```
git clone git@github.com:zadnepr-perfect/go-api.git .
docker compose up -d --build
```

## Апи доступно по:
```
http://localhost:8080
```

### Для сборки образа
```
docker build -f ./docker/api/Dockerfile -t zadnepr/go-api:latest .
```

### Для пуша образа в hub
```
docker push zadnepr/go-api:latest    
```

# Kubernetes Для интеграционного тестирования
```
minikube start --driver=docker    
kubectl apply -f k8s/
```