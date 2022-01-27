package interceptors

import (
	"context"
	"log"
	"runtime"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Identity struct {
	ID string
}

func (i Identity) UnaryClient(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	md := metadata.Pairs()
	md.Set("client-id", i.ID)

	ctx = metadata.NewOutgoingContext(ctx, md)

	// Get the operating system the client is running on
	cos := runtime.GOOS
	// Append the OS info to the outgoing request
	ctx = metadata.AppendToOutgoingContext(ctx, "client-os", cos)

	log.Println("Identity: 1")
	time.Sleep(time.Second)

	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Println("Identity: 2")
	log.Printf("client interceptor hit: appending OS: '%v' to metadata", cos)
	time.Sleep(time.Second)

	return err
}
