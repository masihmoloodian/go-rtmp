# Nginx RTMP service

Automatic nginx rtmp configuraton
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

## OBS Configuration
* Stream Type: `Custom Streaming Server`
* URL: `rtmp://localhost:1935/<userId>`
* Stream Key: `<secretKey>`

## Watch Stream
```
http://localhost:8080/live/<userId>.m3u8
```

## TODO
- [-] Remove only bad configuration
