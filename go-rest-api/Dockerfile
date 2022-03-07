FROM golang:1.16-alpine as builder

RUN mkdir -p /go-rest-api/

WORKDIR /go-rest-api

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -a -o bin/server_main ./main.go

FROM alpine:3.14

WORKDIR /home/app

COPY --from=builder /go-rest-api/bin/server_main .

CMD ["./server_main"]
