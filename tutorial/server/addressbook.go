package server

import (
	"context"

	"github.com/iamseki/go-grpc/tutorial/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AddressBookServer struct {
	Addresses proto.AddressBook
	proto.UnimplementedAddressBookInfoServer
}

func (s *AddressBookServer) AddAddress(ctx context.Context, in *proto.Person) (*proto.Void, error) {
	in.LastUpdated = timestamppb.Now()
	s.Addresses.People = append(s.Addresses.People, in)
	return &proto.Void{}, status.New(codes.OK, "").Err()
}

func (s *AddressBookServer) FindAddresses(ctx context.Context, in *proto.Void) (*proto.AddressBook, error) {
	return &s.Addresses, status.New(codes.OK, "").Err()
}
