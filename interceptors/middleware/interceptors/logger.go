package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// LogRequest is a gRPC UnaryServerInterceptor that will log the API call to stdOut
func LogRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {

	log.Println("Log: 1")
	time.Sleep(time.Second)

	log.Printf("Request for : %s\n", info.FullMethod)

	res, err := handler(ctx, req)

	log.Println("Log: 2")
	time.Sleep(time.Second)

	return res, err
}
