package main

import (
	"log"
	"net"

	"github.com/iamseki/go-grpc/patterns/proto"
	"github.com/iamseki/go-grpc/patterns/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50502")
	if err != nil {
		log.Fatalf("Error to listen: %v", err)
	}

	s := grpc.NewServer()
	server := server.NewOrderManagementServer()
	proto.RegisterOrderManagementServer(s, server)

	log.Printf("Starting gRPC listener on port :50502")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
