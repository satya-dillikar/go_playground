
export OUT_DIR=/tmp/proto
echo $OUT_DIR
mkdir -p $OUT_DIR

protoc --go_out=$OUT_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
    pingpong/service.proto

protoc -I pingpong \
    --grpc-gateway_out $OUT_DIR --grpc-gateway_opt paths=source_relative \
    pingpong/service.proto

protoc  pingpong/service.proto \
		--js_out=import_style=commonjs:$OUT_DIR \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT_DIR

protoc  service.proto \
   --js_out=import_style=commonjs:$OUT_DIR/proto \
   --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$OUT_DIR/proto


    go get -u google.golang.org/protobuf/cmd/protoc-gen-go  
    go install google.golang.org/protobuf/cmd/protoc-gen-go  
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc  
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc



#
http://www.inanzzz.com/index.php/post/7vsp/creating-grpc-unary-middleware-interceptor-for-client-and-server-golang-applications

https://github.com/ori-edge/grpc-interceptor-demo/pull/2/commits/53e47873e237e41b8cb828f53eb5a17028b75a4e
https://github.com/ori-edge/grpc-interceptor-demo

