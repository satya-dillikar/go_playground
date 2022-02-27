# Helloworld Buf

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* Describe any prerequisites, libraries, OS version, etc., needed before installing program.
* ex. Windows 10

* XXX
```
cd github/projects/go_playground/grpc-example/helloworld_buf
rm go.mod go.sum
rm -rf gen/proto
```

* XXX
```
tree .
.
├── README.txt
├── buf.gen.yaml
├── gen
├── greeter_client
│   └── main.go
├── greeter_server
│   └── main.go
└── proto
    ├── buf.yaml
    └── helloworld.proto
```

* XXX
```
go mod init satya.com/helloworld_buf
```

### Installing

* How/where to download your program
* Any modifications needed to be made to files/folders

* change all
```
helloworld.proto
option go_package = "satya.com/helloworld_buf/gen/proto";

greeter_client/main.go
import pb "satya.com/helloworld_buf/gen/proto"

greeter_server/main.go
import pb "satya.com/helloworld_buf/gen/proto"
```

### Executing program

* How to run the program

* XXX
```
buf generate
```

* XXX
```
tree  .
.
├── README.txt
├── buf.gen.yaml
├── gen
│   └── proto
│       ├── helloworld.pb.go
│       └── helloworld_grpc.pb.go
├── go.mod
├── greeter_client
│   └── main.go
├── greeter_server
│   └── main.go
└── proto
    ├── buf.yaml
    └── helloworld.proto

```

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
* [GRPC Tutorial](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/simple_hello_world/)
* [Generating Stubs](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/generating_stubs/using_buf/)