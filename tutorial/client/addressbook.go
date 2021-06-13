package main

import (
	"context"
	"log"

	"github.com/iamseki/go-grpc/tutorial/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50501", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Something went wrong when connect to grpc server :50501 : %v", err)
	}
	defer conn.Close()

	client := proto.NewAddressBookInfoClient(conn)

	_, err = client.AddAddress(context.TODO(), &proto.Person{})
	if err != nil {
		log.Fatalf("Something went wrong when trying to add a person to addressBook: %v", err)
	}

	addressBook, err := client.FindAddresses(context.TODO(), &proto.Void{})
	if err != nil {
		log.Fatalf("Something went wrong when trying to add a person to addressBook: %v", err)
	}

	log.Printf("People count on addressBook: %v", len(addressBook.People))
}
