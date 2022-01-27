package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"satya.com/satya-dillikar/grpcexample/pingpong"
)

func main() {
	ctx := context.Background()
	// Create a new connection using the transport credentials
	conn, err := grpc.DialContext(ctx, "localhost:9990", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// A new GRPC client to use
	client := pingpong.NewPingPongClient(conn)

	pong, err := client.Ping(ctx, &pingpong.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("response received:", pong)
}
