# Auth Service with gRPC

## How to Generated new gRPC file from Proto (if the are any changes):

```bash
$ cd src/delivery/protos/
$ protoc --go_out=../pb/ --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative ./user.proto
```

## How to Run Server:
```bash
$ go mod tidy
$ cp config.yaml.example config.yaml
$ nano config.yaml # modify your .env here
$ go run main.go
```