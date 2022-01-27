package interceptors

import (
	"context"
	"log"
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

	log.Println("Identity: 1")
	time.Sleep(time.Second)

	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Println("Identity: 2")
	time.Sleep(time.Second)

	return err
}
