proto.gen:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto

server.on:
	go run server/main.go

client.on:
	go run client/main.go