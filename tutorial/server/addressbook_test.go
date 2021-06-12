package server_test

import (
	"context"
	"testing"

	"github.com/iamseki/go-grpc/tutorial/proto"
	"github.com/iamseki/go-grpc/tutorial/server"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func makeFakePerson() *proto.Person {
	return &proto.Person{
		Name:  "Thamirets",
		Id:    1,
		Email: "thamirets69@gmail.com",
		Phones: []*proto.Person_PhoneNumber{
			{
				Number: "+55 (11) 9 1234 5678",
				Type:   proto.Person_HOME,
			},
		},
		LastUpdated: timestamppb.Now(),
	}
}

func TestAddAddressSuccess(t *testing.T) {
	p := makeFakePerson()

	sut := server.AddressBookServer{Addresses: proto.AddressBook{}}

	_, err := sut.AddAddress(context.TODO(), p)
	inserted := sut.Addresses.People[0]

	assert.Nil(t, err)
	assert.Equal(t, "thamirets69@gmail.com", inserted.Email)
	assert.Equal(t, int32(1), inserted.Id)
}

func TestFindAddressesSuccess(t *testing.T) {
	sut := server.AddressBookServer{Addresses: proto.AddressBook{}}
	sut.AddAddress(context.TODO(), makeFakePerson())

	addressBook, err := sut.FindAddresses(context.TODO(), &proto.Void{})

	assert.Nil(t, err)
	assert.NotEmpty(t, addressBook.People)

	expectedLen := 1
	actualLen := len(addressBook.People)
	assert.Equal(t, expectedLen, actualLen)
}
