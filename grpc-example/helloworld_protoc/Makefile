NAME:=helloworld-protoc
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	protoc --go_out=gen --go_opt=paths=source_relative \
		--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
		proto/helloworld.proto
	go build -o bin/greeter_server ./greeter_server/main.go
	go build -o bin/greeter_client ./greeter_client/main.go

run :
	./bin/greeter_server &
	sleep 5
	./bin/greeter_client Alice

clean :
	@rm -rf bin/*
	@rm -rf gen/*

test:
	go test -v ./...

build-container :
	@docker build -t dsatya6/$(NAME):$(VERSION) .

push-container :
	@docker push dsatya6/$(NAME):$(VERSION)

test-container :
	@docker rm -f $(NAME) || true
	# greeter_server is running on 50051
	@docker run -d -p 50051:50051 --name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@sleep 5
	./bin/greeter_client Alice
	@echo  ""
	@sleep 2
	@docker stop $(NAME) || true
	@docker rm -f $(NAME) || true