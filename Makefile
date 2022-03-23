all:
	mkdir -p "pkg/service"
	protoc \
	-Iapi/proto \
	-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:./pkg/service \
	--swagger_out=allow_merge=true,merge_file_name=api:./api/swagger \
	--go_out=plugins=grpc:./pkg/service ./api/proto/*.proto