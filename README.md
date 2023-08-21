# Auth Service with gRPC

## How to Generated gRPC file from Proto:

```bash
protoc --proto_path=protos protos/*.proto --go_out=/d/session/go/src --go-grpc_out=/d/session/go/src
```
or
```bash
protoc --go_out=../pb/ --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative ./user.proto
```

## How to Run:
```bash
go mod tidy
cp config.yaml.example config.yaml
nano config.yaml # modify your .env here
go run main.go
```