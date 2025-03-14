# EyeSee Service

## Dev

### Generate gRPC code

Run from `EyeSee` dir

```bash
protoc --go_out=./dev/service/proto --go-grpc_out=./dev/service/proto ./protocols/*.proto
```

### Build

```bash
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ../../service/eyesee main.go
```

### Client

You can test the service with `grpc-client-cli` (https://github.com/vadimi/grpc-client-cli)

Run from `EyeSee` dir

```bash
grpc-client-cli --proto=protocols/service.proto "localhost:25910"
```
