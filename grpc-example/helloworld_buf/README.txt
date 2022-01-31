
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/simple_hello_world/
https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/generating_stubs/using_buf/


ACTUAL commands


cd github/projects/go_playground/grpc-example/helloworld_buf
rm go.mod go.sum
rm -rf gen/proto

tree .
.
├── README.txt
├── buf.gen.yaml
├── gen
├── greeter_client
│   └── main.go
├── greeter_server
│   └── main.go
└── proto
    ├── buf.yaml
    └── helloworld.proto


go mod init satya.com/helloworld_buf

change all
helloworld.proto
option go_package = "satya.com/helloworld_buf/gen/proto";

greeter_client/main.go
import pb "satya.com/helloworld_buf/gen/proto"

greeter_server/main.go
import pb "satya.com/helloworld_buf/gen/proto"

buf generate

tree  .
.
├── README.txt
├── buf.gen.yaml
├── gen
│   └── proto
│       ├── helloworld.pb.go
│       └── helloworld_grpc.pb.go
├── go.mod
├── greeter_client
│   └── main.go
├── greeter_server
│   └── main.go
└── proto
    ├── buf.yaml
    └── helloworld.proto


go mod tidy

go run greeter_server/main.go


go run greeter_client/main.go Alice