
NAME:=greeter_server
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	# go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	buf mod update proto
	buf generate proto
	go mod tidy
	go build -o bin/greeter_server ./greeter_server/main.go
	go build -o bin/greeter_client ./greeter_client/main.go

run :
	./bin/greeter_server &
	sleep 2
	grpcurl -plaintext 0.0.0.0:50051 list
	sleep 1
	curl -X POST -k http://127.0.0.1:8090/v1/example/echo -d '{"name" : "satya"}'

clean :
	@rm -rf bin/*
	@rm -rf gen/*
	@rm -rf proto/buf.lock
	@rm -rf proto/buf.lock


test:
	go test -v ./...

build-container :
	@docker build -t dsatya6/$(NAME):$(VERSION) .

push-container :
	@docker push dsatya6/$(NAME):$(VERSION)

test-container :
	@docker rm -f $(NAME) || true
	# greeter_server is running on 50051
	@docker run -dp 50051:9001 --name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@echo  ""
	@curl -X POST -k http://127.0.0.1:8090/v1/example/echo -d '{"name" : "satya"}'
	@echo  ""
	@grpcurl -plaintext 0.0.0.0:50051 list
	@echo  ""
	@grpcurl -plaintext  0.0.0.0:50051 describe Greeter
	@echo  ""
	@grpcurl -plaintext  0.0.0.0:50051 describe HelloRequest
	@echo  ""
	@grpcurl -plaintext  0.0.0.0:50051 describe HelloReply
	@echo  ""
	@grpcurl -plaintext -d '{ "name" : "sunny"}' 0.0.0.0:50051 Greeter.SayHello
	@echo  ""
	@grpcurl -plaintext -d '{ "name" : "sunny"}' 0.0.0.0:50051 Greeter.SayHelloAgain
	@echo  ""
	@grpcurl -plaintext  0.0.0.0:50051 Greeter.GetAllEvents
	@echo  ""
	@grpcurl -plaintext -d '{"ID" :"1"}'  0.0.0.0:50051 Greeter.GetOneEvent
	@echo  ""
	@grpcurl -plaintext -d '{"ID" :"4"}'  0.0.0.0:50051 Greeter.DeleteEvent
	@echo  ""
	@grpcurl -plaintext  0.0.0.0:50051 Greeter.GetAllEvents
	@echo  ""
	@grpcurl -plaintext -d '{"ID": "4", "data" :{ "Title" : "JAVA4", "Description":"JAVA 4"}}'  0.0.0.0:50051 Greeter.UpdateEvent
	@echo  ""
	@grpcurl -plaintext  0.0.0.0:50051 Greeter.GetAllEvents
