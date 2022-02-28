rem order_management.bat
protoc --proto_path=./../app/do/stream/proto --go_out=../../ --go-grpc_out=../../ order_management.proto