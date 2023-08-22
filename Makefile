test:
	@go test -v ./...

update-proto:
	@cd src/delivery/protos/ && protoc --go_out=../pb/ --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative ./user.proto