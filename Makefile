authorizer:
	go run ./cmd/authorizer/main.go

proto-update:
	protoc --proto_path=pkg/proto --go_out=pkg/proto/pb --go_opt=paths=source_relative --go-grpc_out=pkg/proto/pb --go-grpc_opt=paths=source_relative pkg/proto/*.proto
