# Auth Service with gRPC
Simple Auth Service build with gRPC in Go.

## Lib and Stack 📚
- [protobuf](https://pkg.go.dev/google.golang.org/protobuf) && [grpc](https://pkg.go.dev/google.golang.org/grpc) for gRPC library. 
- [validator](https://github.com/go-playground/validator) for validate incoming request.
- [mysql-driver](https://github.com/go-sql-driver/mysql) for realtional SQL driver.
- [uuid](https://github.com/google/uuid) for user unique identifier.
- [zerolog](https://github.com/rs/zerolog) for service logging.
- [viper](https://github.com/spf13/viper) for env configuration.
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) for password hashing.
- [testify](https://github.com/stretchr/testify)
 && [mockery](https://github.com/vektra/mockery) for unit testing and mocking.

## How to Update gRPC file from Proto (if the are any changes) 🤖
```bash
$ make update-proto
```

## How to Run Test 💉
```bash
$ make test
```
## How to Run Server 🏃‍♀️
Make sure you have MySQL `tb_user` table fisrt. Check in **/config/db/dump.sql** for definition.
```bash
$ go mod tidy
$ cp config.yaml.example config.yaml
$ nano config.yaml # modify your .env here
$ go run main.go
```