package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/iamseki/go-grpc/patterns/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50502", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Something went wrong when connect to grpc server :50502 : %v", err)
	}
	defer conn.Close()

	client := proto.NewOrderManagementClient(conn)

	order, err := client.GetOrder(context.TODO(), &wrappers.StringValue{Value: "1"})
	if err != nil {
		log.Fatalf("Error on get Order ID: 1, %v", err)
	}

	log.Printf("Order 1: %v", order)

	order, err = client.GetOrder(context.TODO(), &wrappers.StringValue{Value: "2"})
	if err != nil {
		log.Fatalf("Error on get Order ID: 1, %v", err)
	}

	log.Printf("Order 2: %v", order)
}
