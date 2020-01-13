build:
	protoc --proto_path=./proto/hello --micro_out=. --go_out=. greeter.proto