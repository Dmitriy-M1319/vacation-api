.PHONY: run proto

run:
	go run cmd/vacation-api/main.go

proto:
	protoc -I proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative proto/vacation-api.proto
