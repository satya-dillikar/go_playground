package interceptors

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// PingCounter is a struct that keeps track of how many Pings that are performed
type PingCounter struct {
	Pings int
}

// ServerCount is a gRPC UnaryServerInterceptor that will count number of API calls.
func (pc *PingCounter) ServerCount(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
	// Append to PingCounts
	pc.Pings++
	// We want to extract metadata from the incomming context.
	// We dont create a new context since we dont wanna overwrite old metadata
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not grab metadata from context")
	}
	// Set ping-counts into the current ping value
	meta.Set("ping-counts", fmt.Sprintf("%d", pc.Pings))
	// Metadata is sent on its own, so we need to send the header. There is also something called Trailer
	grpc.SendHeader(ctx, meta)
	// Last but super important, execute the handler so that the actualy gRPC request is also performed
	return handler(ctx, req)
}

// ClientPingCounter is a UnaryClientInterceptor that will count the number of API calls on the Client side
func (pc *PingCounter) ClientPingCounter(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	pc.Pings++
	// Run regular gRPC call after
	// If you dont run the invoker, the gRPC call wont be sent to the server
	return invoker(ctx, method, req, reply, cc, opts...)
}

func VerifyUnaryServer(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("Verify: 1")
	time.Sleep(time.Second)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md["client-id"]) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if md["client-id"][0] != "client-1" {
		return nil, status.Error(codes.PermissionDenied, "unexpected client")
	}

	// Retrieve the client OS, this will be empty if it does not exist
	os := md.Get("client-os")
	// Get the client IP Address
	ip, err := getClientIP(ctx)
	if err != nil {
		return nil, err
	}

	res, err := handler(ctx, req)

	log.Println("Verify: 2")
	log.Printf("server interceptor hit: hydrating type with OS: '%v' and IP: '%v'", os[0], ip)
	time.Sleep(time.Second)

	return res, err
}

// GetClientIP inspects the context to retrieve the ip address of the client
func getClientIP(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("couldn't parse client IP address")
	}

	return p.Addr.String(), nil
}
