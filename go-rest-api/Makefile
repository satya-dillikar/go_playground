
NAME:=go-rest-api
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	go build -o bin/server_main ./main.go

clean :
	@rm bin/*

test:
	# go test -v ./...
	./bin/server_main &
	sleep 5
	curl -s localhost:8080/events

build-container :
	@docker build -t dsatya6/$(NAME):$(VERSION) .

push-container :
	@docker push dsatya6/$(NAME):$(VERSION)	

test-container :
	@docker rm -f $(NAME) || true
	@docker run -dp 8080:8080 --name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@sleep 5
	@curl -s localhost:8080/events