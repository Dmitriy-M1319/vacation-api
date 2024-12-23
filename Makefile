.PHONY: run proto proto_clean

run:
	go run cmd/vacation-api/main.go

proto:
	protoc --experimental_allow_proto3_optional \
		-I protos \
		--go_out=pkg --go_opt=paths=source_relative \
		--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pkg --grpc-gateway_opt=generate_unbound_methods=true --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=swagger \
		protos/vacation_api/v1/vacation_api.proto
proto_clean:
	rm -rf pkg/vacation_api
