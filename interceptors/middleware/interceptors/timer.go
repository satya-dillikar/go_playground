package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func TimerUnaryClient(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()

	log.Println("Timer: 1")
	time.Sleep(time.Second)

	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Println("Timer: 2")
	time.Sleep(time.Second)

	end := time.Since(start)

	log.Printf("%s method call took %s\n", method, end)

	return err
}
