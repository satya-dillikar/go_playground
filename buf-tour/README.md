# Buf Tour

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started
```
buf config init

buf generate petapis
```


### Dependencies

* Describe any prerequisites, libraries, OS version, etc., needed before installing program.
* ex. Windows 10

```
https://buf.build/login
login  - use github account credentials.

export BUF_USER=satya-dillikar


Token
Token "dev" successfully created:
0a3f8e38b2384bed95e91ed6d7b487c2ecfe0aa5486e4f0499a86be854c5165a
```


### Installing

* How/where to download your program
* Any modifications needed to be made to files/folders



```
buf registry login

cat ~/.netrc <- this file is updated.

buf beta registry repository create buf.build/$BUF_USER/petapis --visibility public

cd /Users/sdillikar/myprojects/golang-projects/buf-tour/buf-tour/start/petapis
buf push
```
### Executing program

* How to run the program

```
cd .../buf-tour/petstore

mkdir petapis
cd petapis
buf config init
#this will create file buf.yaml
```

* Step-by-step bullets

```
buf build --exclude-source-info -o -#format=json | jq '.file[] | .package' | sort | uniq | head
"google.protobuf"
"google.type"
"pet.v1"
```



* Step-by-step bullets
```
buf ls-files

buf lint

create buf.gen.yaml

#in pet.proto file add below line
option go_package = "satya.com/buf-tour/petstore/gen/proto/go/pet/v1;petv1";

buf generate petapis
#this will create 'gen' folder

go mod init satya.com/buf-tour/petstore

update server and client main.go files with
import	petv1 "satya.com/buf-tour/petstore/gen/proto/go/pet/v1"

go mod tidy
```

* run server
```
go run server/main.go
2021/10/01 17:11:06 Listening on 127.0.0.1:8080
2021/10/01 17:11:17 Got a request to create a PET_TYPE_SNAKE named Ekans
```

* run client
```
go run client/main.go
2021/10/01 17:11:17 Connected to 127.0.0.1:8080
2021/10/01 17:11:17 Successfully PutPet
```

* XXX

```
curl -s http://127.0.0.1:8080/ | jq .
curl -s http://127.0.0.1:8080/
```

* XXX

```
grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto list
grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto describe pet.v1.PetStoreService
```

* XXX

```
grpcurl -plaintext 127.0.0.1:8080  list

grpcurl -plaintext -d '{"pet_type" : "2", "name" : "Sunny" }' 127.0.0.1:8080  pet.v1.PetStoreService.PutPet
```

* XXX

```
grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto list
pet.v1.PetStoreService

grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto describe pet.v1.PetStoreService
pet.v1.PetStoreService is a service:
service PetStoreService {
  rpc DeletePet ( .pet.v1.DeletePetRequest ) returns ( .pet.v1.DeletePetResponse );
  rpc GetPet ( .pet.v1.GetPetRequest ) returns ( .pet.v1.GetPetResponse );
  rpc PutPet ( .pet.v1.PutPetRequest ) returns ( .pet.v1.PutPetResponse );
}

grpcurl -plaintext 127.0.0.1:8080  list
grpc.reflection.v1alpha.ServerReflection
pet.v1.PetStoreService

grpcurl -plaintext 127.0.0.1:8080  pet.v1.PetStoreService.GetPet


grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto describe pet.v1.PetStoreService.GetPet
pet.v1.PetStoreService.GetPet is a method:
rpc GetPet ( .pet.v1.GetPetRequest ) returns ( .pet.v1.GetPetResponse );

grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto describe pet.v1.PetStoreService.PutPet
pet.v1.PetStoreService.PutPet is a method:
rpc PutPet ( .pet.v1.PutPetRequest ) returns ( .pet.v1.PutPetResponse );

grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto describe pet.v1.PutPetRequest
pet.v1.PutPetRequest is a message:
message PutPetRequest {
  .pet.v1.PetType pet_type = 1;
  string name = 2;
}

grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto describe pet.v1.GetPetRequest
pet.v1.GetPetRequest is a message:
message GetPetRequest {
  string pet_id = 1;
}

grpcurl --plaintext  --import-path ./petapis --proto petapis/pet/v1/pet.proto describe pet.v1.PutPetRequest
pet.v1.PutPetRequest is a message:
message PutPetRequest {
  .pet.v1.PetType pet_type = 1;
  string name = 2;
}

grpcurl -plaintext -d '{"pet_type" : "2", "name" : "Sunny" }' 127.0.0.1:8080  pet.v1.PetStoreService.PutPet
{

}

grpcurl -plaintext -d '{"pet_id" : "1" }' 127.0.0.1:8080  pet.v1.PetStoreService.GetPet
{

}

grpcurl -plaintext -d '{"pet_id" : "11" }' 127.0.0.1:8080  pet.v1.PetStoreService.GetPet
{

}

grpcurl -plaintext -d '{"pet_id" : "11" }' 127.0.0.1:8080  pet.v1.PetStoreService.DeletePet
{

}
```

* BloomPRC
```
step1: set "Import Paths"
/Users/sdillikar/github/projects/go_playground/buf-tour/petstore/petapis

step2: "Import Protos" click green "+" button and import pet.proto file
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

Inspiration, code snippets, etc.
* [Author](https://softbuild.dev)
* [Buf tour](https://github.com/bufbuild/buf-tour/tree/main/finish)
* [docs.buf.build](https://docs.buf.build/tour/generate-code)