rem product_info.bat
protoc --proto_path=./../app/do/normal/proto --go_out=../../ --go-grpc_out=../../ product_info.proto