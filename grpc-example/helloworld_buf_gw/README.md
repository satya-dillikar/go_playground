# HelloWorld Buf GW

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* Describe any prerequisites, libraries, OS version, etc., needed before installing program.
* ex. Windows 10

### Installing

* How/where to download your program

```
cd github/projects/go_playground/grpc-example/helloworld_buf
rm go.mod go.sum
rm -rf gen/proto
```

* Any modifications needed to be made to files/folders

```
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
```


```
go mod init satya.com/helloworld_buf
```

#### DO NOT RUN "go mod tidy" YET. see below

### Executing program

* change all
```
helloworld.proto
helloworld.proto
option go_package = "satya.com/helloworld_buf_gw/gen/proto";

greeter_client/main.go
import pb "satya.com/helloworld_buf_gw/gen/proto"

greeter_server/main.go
import pb "satya.com/helloworld_buf_gw/gen/proto"
```

* Step-by-step bullets
```
#update buf.yaml
#update buf.gen.yaml

go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
which protoc-gen-grpc-gateway

buf mod update proto
buf generate proto


go mod tidy
```

* XXX

```
tree  .
.
├── README.txt
├── buf.gen.yaml
├── gen
│   └── proto
│       ├── helloworld.pb.go
│       ├── helloworld.pb.gw.go
│       └── helloworld_grpc.pb.go
├── go.mod
├── go.sum
├── greeter_client
│   └── main.go
├── greeter_server
│   └── main.go
└── proto
    ├── buf.lock
    ├── buf.yaml
    └── helloworld.proto
```

* 
```
go mod tidy
```

### Method-1
* run Server
```
go run greeter_server/main.go
```


* run Client 
```
go run greeter_client/main.go Alice
```

### Method-2

* run Server
```
go run greeter_server/main.go
2021/10/06 12:16:23 server listening at [::]:50051
2021/10/06 12:16:23 server listening at 0.0.0.0::50051
2021/10/06 12:16:23 Serving gRPC-Gateway on http://0.0.0.0:8090
2021/10/06 12:16:38 Received: sa
```

* run Client 
```
curl -X POST -k http://127.0.0.1:8090/v1/example/echo -d '{"name" : "satya"}'

```

* To enable grpcurl add below lines to greeter_server/main.go 
```
import "google.golang.org/grpc/reflection"

// Register reflection service on gRPC server.
reflection.Register(s)
```

```
grpcurl -plaintext 0.0.0.0:50051 list
Greeter
grpc.reflection.v1alpha.ServerReflection
```

```
grpcurl -plaintext  0.0.0.0:50051 describe Greeter
Greeter is a service:
service Greeter {
  rpc SayHello ( .HelloRequest ) returns ( .HelloReply ) {
    option (.google.api.http) = { post:"/v1/example/echo" body:"*"  };
  }
  rpc SayHelloAgain ( .HelloRequest ) returns ( .HelloReply );
}

grpcurl -plaintext  0.0.0.0:50051 describe HelloRequest
HelloRequest is a message:
message HelloRequest {
  string name = 1;
}
grpcurl -plaintext  0.0.0.0:50051 describe HelloReply
HelloReply is a message:
message HelloReply {
  string message = 1;
}

grpcurl -plaintext -d '{ "name" : "sunny"}' 0.0.0.0:50051 Greeter.SayHello
{
  "message": "From server: Hello sunny"
}

grpcurl -plaintext -d '{ "name" : "sunny"}' 0.0.0.0:50051 Greeter.SayHelloAgain
{
  "message": "From server: Hello again sunny"
}
```

```

grpcurl -plaintext  0.0.0.0:50051 Greeter.GetAllEvents
{
  "data": [
    {
      "ID": "1",
      "Title": "Introduction to Golang",
      "Description": "Come join us for a chance "
    },
    {
      "ID": "2",
      "Title": "Introduction to Python",
      "Description": "Come to learn "
    },
    {
      "ID": "4",
      "Title": "Introduction to Java",
      "Description": "try it out"
    }
  ]
}

 grpcurl -plaintext -d '{"ID" :"1"}'  0.0.0.0:50051 Greeter.GetOneEvent
{
  "data": {
    "ID": "1",
    "Title": "Introduction to Golang",
    "Description": "Come join us for a chance "
  }
}

grpcurl -plaintext -d '{"ID" :"4"}'  0.0.0.0:50051 Greeter.DeleteEvent


grpcurl -plaintext -d '{"data" :{"ID" :"4", "Title" : "TEST4"}}'  0.0.0.0:50051 Greeter.CreateEvent
{
  "data": {
    "ID": "4",
    "Title": "TEST4"
  }
}
➜  private_repo grpcurl -plaintext  0.0.0.0:50051 Greeter.GetAllEvents
{
  "data": [
    {
      "ID": "1",
      "Title": "Introduction to Golang",
      "Description": "Come join us for a chance "
    },
    {
      "ID": "2",
      "Title": "Introduction to Python",
      "Description": "Come to learn "
    },
    {
      "ID": "4",
      "Title": "TEST4"
    }
  ]
}

grpcurl -plaintext -d '{"ID": "4", "data" :{ "Title" : "JAVA4", "Description":"JAVA 4"}}'  0.0.0.0:50051 Greeter.UpdateEvent
{
  "data": {
    "ID": "4",
    "Title": "TEST4"
  }
}
grpcurl -plaintext  0.0.0.0:50051 Greeter.GetAllEvents
{
  "data": [
    {
      "ID": "1",
      "Title": "Introduction to Golang",
      "Description": "Come join us for a chance "
    },
    {
      "ID": "2",
      "Title": "Introduction to Python",
      "Description": "Come to learn "
    },
    {
      "ID": "4",
      "Title": "JAVA4",
      "Description": "JAVA 4"
    }
  ]
}

```

## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```

## Authors

Contributors names and contact info

[@SatyaDillikar](https://twitter.com/SatyaDillikar)

## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release

## License

N/A

## Acknowledgments
* [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/simple_hello_world/)
* [generating_stubs](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/generating_stubs/using_buf/)