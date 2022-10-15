# Nginx RTMP service

## Requirement
- golang
- docker
- docker compose
- inotify-tools

 ## Run
 ```
 cd rtmp-server
 docker compose up -d
 ```

```
go build 
./go-rtmp
```

## Sample query
```
curl -X POST http://localhost:8000/create -H "Content-Type: application/json" -d '{"userId": "123", "secretKey": "123"}'
```
