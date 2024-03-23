prepare:
	brew install protobuf
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	mkdir internal/pb
	protoc -I . ./api/*.proto -I ./api --go-grpc_out=internal/pb --go_out=internal/pb