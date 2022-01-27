.PHONY: compile
compile:
	buf mod update pingpong
	buf generate pingpong

.PHONY: client
client:
	go run --race client/main.go

.PHONY: server
server:
	go run --race server/main.go server/server.go

.PHONY: clean
clean:
	rm -rf go.sum
	rm -rf pingpong/*pb.*

.PHONY: build
build:
	go mod tidy