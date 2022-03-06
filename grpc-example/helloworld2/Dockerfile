# FROM golang:1.16-alpine as builder
FROM --platform=linux/amd64 golang:1.16-alpine as builder

RUN mkdir -p /go_app/

WORKDIR /go_app

COPY . .

# System setup
# RUN apk update && apk add git curl build-base autoconf automake libtool

# # Install protoc
# ENV PROTOBUF_URL https://github.com/google/protobuf/releases/download/v3.3.0/protobuf-cpp-3.3.0.tar.gz
# RUN curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL
# WORKDIR /tmp/
# RUN tar xvzf protobuf.tar.gz
# WORKDIR /tmp/protobuf-3.3.0
# RUN mkdir /export
# RUN ./autogen.sh && \
#     ./configure --prefix=/export && \
#     make -j 3 && \
#     make check && \
#     make install

# # Install protoc-gen-go
# RUN go get github.com/golang/protobuf/protoc-gen-go
# RUN cp /go/bin/protoc-gen-go /export/bin/


# Get the source from GitHub
# RUN go get google.golang.org/grpc
# Install protoc-gen-go
# RUN go get github.com/golang/protobuf/protoc-gen-go

# Install protoc (cf. http://google.github.io/proto-lens/installing-protoc.html)
RUN apk add curl
ENV PROTOC_ZIP=protoc-3.13.0-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \ 
    && rm -f $PROTOC_ZIP

RUN apk add --no-cache protobuf git \
  && go get github.com/golang/protobuf/protoc-gen-go \
  && cp /go/bin/protoc-gen-go /usr/bin/

RUN go mod download

RUN protoc --go_out=gen --go_opt=paths=source_relative \
    --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto

RUN CGO_ENABLED=0 go build -o bin/greeter_server ./greeter_server/main.go
RUN CGO_ENABLED=0 go build -o bin/greeter_client ./greeter_client/main.go

FROM alpine:3.14

WORKDIR /home/app

COPY --from=builder /go_app/bin/greeter_server .

CMD ["./greeter_server"]
