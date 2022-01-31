
https://grpc.io/docs/languages/go/quickstart/

$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$ git clone -b v1.41.0 https://github.com/grpc/grpc-go
----------------------------------------------------------------

ACTUAL commands


cd github/projects/go_playground/grpc-example/helloworld2
rm go.mod go.sum
rm -rf gen/proto

tree .
.
├── README.txt
├── gen
├── greeter_client
│   └── main.go
├── greeter_server
│   └── main.go
└── proto
    └── helloworld.proto

go mod init satya.com/helloworld2

protoc --go_out=gen --go_opt=paths=source_relative \
    --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto



change all
proto/helloworld.proto
option go_package = "satya.com/helloworld2/gen/proto";

greeter_client/main.go
import pb "satya.com/helloworld2/gen/proto"

greeter_server/main.go
import pb "satya.com/helloworld2/gen/proto"


go mod tidy

go run greeter_server/main.go


go run greeter_client/main.go Alice