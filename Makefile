prepare:
	brew install protobuf
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	mkdir -p internal/pb
	protoc -I ./api ./api/user*.proto -I ./api \
 			--go-grpc_out=internal/pb \
 			--go_out=internal/pb \
 			--grpc-gateway_out=internal/pb \
 			--grpc-gateway_opt generate_unbound_methods=true

run:
	docker-compose up --build