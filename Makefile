build:
	protoc --proto_path=./proto/hello --micro_out=./proto/hello --go_out=./proto/hello greeter.proto