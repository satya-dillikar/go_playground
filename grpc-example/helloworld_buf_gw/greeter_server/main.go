/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "satya.com/helloworld_buf_gw/gen/proto"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "From server: Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "From server: Hello again " + in.GetName()}, nil
}

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance ",
	},
	{
		ID:          "2",
		Title:       "Introduction to Python",
		Description: "Come to learn ",
	},
	{
		ID:          "4",
		Title:       "Introduction to Java",
		Description: "try it out",
	},
}

func (s *server) HomeLink(ctx context.Context, in *pb.HomeLinkRequest) (*pb.HomeLinkResponse, error) {

	return &pb.HomeLinkResponse{Greetings: "Welcome home!"}, nil
}

func (s *server) CreateEvent(ctx context.Context, in *pb.PostRequest) (*pb.PostResponse, error) {
	newEvent := event{}

	postData := in.GetData()
	newEvent.ID = postData.ID
	newEvent.Title = postData.Title
	newEvent.Description = postData.Description

	log.Println("CreateEvent:", newEvent)
	events = append(events, newEvent)

	retEvent := &pb.EventStruct{}
	retEvent.ID = newEvent.ID
	retEvent.Title = newEvent.Title
	retEvent.Description = newEvent.Description

	return &pb.PostResponse{Data: retEvent}, nil
}

func (s *server) GetAllEvents(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	newEvents := []*pb.EventStruct{}
	for _, event := range events {
		retEvent := &pb.EventStruct{}
		retEvent.ID = event.ID
		retEvent.Title = event.Title
		retEvent.Description = event.Description
		newEvents = append(newEvents, retEvent)
	}
	return &pb.GetAllResponse{Data: newEvents}, nil
}

func (s *server) GetOneEvent(ctx context.Context, in *pb.GetOneRequest) (*pb.GetOneResponse, error) {

	eventID := in.GetID()
	retEvent := &pb.EventStruct{}
	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			retEvent.ID = singleEvent.ID
			retEvent.Title = singleEvent.Title
			retEvent.Description = singleEvent.Description
			break
		}
	}
	return &pb.GetOneResponse{Data: retEvent}, nil
}

func (s *server) UpdateEvent(ctx context.Context, in *pb.PatchOneRequest) (*pb.PatchOneResponse, error) {
	var updatedEvent event
	eventID := in.GetID()

	patchData := in.GetData()
	updatedEvent.Title = patchData.Title
	updatedEvent.Description = patchData.Description

	log.Println("UpdateEvent:", updatedEvent)
	log.Println("eventID:", eventID)
	retEvent := &pb.EventStruct{}
	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			retEvent.ID = singleEvent.ID
			retEvent.Title = singleEvent.Title
			retEvent.Description = singleEvent.Description
			events[i].Title = updatedEvent.Title
			events[i].Description = updatedEvent.Description
			break
		}
	}
	return &pb.PatchOneResponse{Data: retEvent}, nil
}

func (s *server) DeleteEvent(ctx context.Context, in *pb.DelOneRequest) (*pb.DelOneResponse, error) {
	eventID := in.GetID()

	retEvent := &pb.EventStruct{}
	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			retEvent.ID = singleEvent.ID
			retEvent.Title = singleEvent.Title
			retEvent.Description = singleEvent.Description
			events = append(events[:i], events[i+1:]...)
			log.Printf("The event with ID %v has been deleted successfully\n", eventID)
			break
		}
	}
	return &pb.DelOneResponse{Data: retEvent}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	listener_addr := fmt.Sprintf("%v", lis.Addr())
	log.Printf("server listening at %v", listener_addr)
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Register reflection service on gRPC server.
	log.Printf("Register reflection service on gRPC server")
	reflection.Register(s)

	/*
		// Create a listener on TCP port
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalln("Failed to listen:", err)
		}

		// Create a gRPC server object
		s := grpc.NewServer()
		// Attach the Greeter service to the server
		pb.RegisterGreeterServer(s, &server{})
		// Serve gRPC server
		log.Println("Serving gRPC on 0.0.0.0:8080")
		go func() {
			log.Fatalln(s.Serve(lis))
		}()
	*/

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	/* conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	*/
	listener_addr = fmt.Sprintf("0.0.0.0:%v", port)
	log.Printf("server GW listening at %v", listener_addr)
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving GW gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
