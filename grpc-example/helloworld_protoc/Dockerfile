ARG ALPINE_VERSION=3.14
ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder


RUN mkdir -p /go_app/

WORKDIR /go_app

COPY . .

ENV GO111MODULE=on

RUN go mod download

# RUN protoc --go_out=gen --go_opt=paths=source_relative \
#     --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
#     proto/helloworld.proto

RUN CGO_ENABLED=0 go build -o bin/greeter_server ./greeter_server/main.go
# RUN CGO_ENABLED=0 go build -o bin/greeter_client ./greeter_client/main.go

FROM alpine:${ALPINE_VERSION}

WORKDIR /home/app

RUN apk add --no-cache --upgrade bash

COPY --from=builder /go_app/bin/greeter_server .
# COPY --from=builder /go_app/bin/greeter_client .

CMD ["./greeter_server"]
