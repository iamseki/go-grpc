package main

import (
	"log"
	"net"

	"github.com/iamseki/go-grpc/tutorial/proto"
	"github.com/iamseki/go-grpc/tutorial/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50501")
	if err != nil {
		log.Fatalf("Error to listen: %v", err)
	}

	s := grpc.NewServer()
	server := &server.AddressBookServer{Addresses: proto.AddressBook{}}
	proto.RegisterAddressBookInfoServer(s, server)

	log.Printf("Starting gRPC listener on port :50501")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
