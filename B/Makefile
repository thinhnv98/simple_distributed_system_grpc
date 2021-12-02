gen:
	protoc --go_out=protobuffer --go_opt=paths=source_relative \
        --go-grpc_out=protobuffer --go-grpc_opt=paths=source_relative \
        proto/*.proto

run-client:
	go run main.go