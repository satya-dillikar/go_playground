package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	petv1 "satya.com/buf-tour/petstore/gen/proto/go/pet/v1"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	petv1.RegisterPetStoreServiceServer(server, &petStoreServiceServer{})
	// Register reflection service on gRPC server.
	reflection.Register(server)

	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// petStoreServiceServer implements the PetStoreService API.
type petStoreServiceServer struct {
	petv1.UnimplementedPetStoreServiceServer
}

// PutPet adds the pet associated with the given request into the PetStore.
func (s *petStoreServiceServer) PutPet(ctx context.Context, req *petv1.PutPetRequest) (*petv1.PutPetResponse, error) {
	name := req.GetName()
	petType := req.GetPetType()
	log.Println("Got a request to create a", petType, "named", name)

	return &petv1.PutPetResponse{}, nil
}

func (s *petStoreServiceServer) GetPet(ctx context.Context, req *petv1.GetPetRequest) (*petv1.GetPetResponse, error) {
	pet_id := req.GetPetId()
	log.Println("Got a GetPetRequest pet_id=", pet_id)

	return &petv1.GetPetResponse{}, nil
}

func (s *petStoreServiceServer) DeletePet(ctx context.Context, req *petv1.DeletePetRequest) (*petv1.DeletePetResponse, error) {
	pet_id := req.GetPetId()
	log.Println("Got a request DeletePet pet_id =", pet_id)

	return &petv1.DeletePetResponse{}, nil
}
