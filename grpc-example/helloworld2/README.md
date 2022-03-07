# Helloworld2

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.



## Getting Started

### Dependencies

* protobuf

```
$ brew install protobuf 
$ protoc --version  # Ensure compiler version is 3+

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$ git clone -b v1.41.0 https://github.com/grpc/grpc-go

```

* XXX
```
cd github/projects/go_playground/grpc-example/helloworld2
rm go.mod go.sum
rm -rf gen/proto
```

* XXX
```
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
```

* Init
```
go mod init satya.com/helloworld2
```

### Installing

* gen files
```
protoc --go_out=gen --go_opt=paths=source_relative \
    --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto
```

* change all
```
helloworld.proto
option go_package = "satya.com/helloworld2/gen/proto";

greeter_client/main.go
import pb "satya.com/helloworld2/gen/proto"

greeter_server/main.go
import pb "satya.com/helloworld2/gen/proto"
```

### Executing program

* How to run the program

* XXX
```
go mod tidy
```

* run Server
```
go run greeter_server/main.go
```

* Run Client
```
go run greeter_client/main.go Alice
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
* [GRPC QuickStart](https://grpc.io/docs/languages/go/quickstart/)
