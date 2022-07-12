build:
	go build cmd/hexcloud/hexcloud.go

run:
	go run cmd/hexcloud/hexcloud.go

protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc --go_out=./internal/pkg  --descriptor_set_out=api_descriptor.pb --go-grpc_out=./internal/pkg ./api/hexagon.proto
