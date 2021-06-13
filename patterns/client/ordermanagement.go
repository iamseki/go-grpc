package main

import (
	"context"
	"flag"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/iamseki/go-grpc/patterns/proto"
	"google.golang.org/grpc"
)

func main() {
	var callMethod string

	flag.StringVar(&callMethod, "call", "getOrder", "call method")

	flag.Parse()

	conn, err := grpc.Dial("localhost:50502", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Something went wrong when connect to grpc server :50502 : %v", err)
	}
	defer conn.Close()

	client := proto.NewOrderManagementClient(conn)

	switch callMethod {
	case "getOrder":
		callGetOrder(client)
	case "searchOrders":
		callSearchOrders(client)
	case "updateOrders":
		callUpdateOrders(client)
	default:
		log.Fatalln("Not allowed method")
	}
}

func callGetOrder(c proto.OrderManagementClient) {
	order, err := c.GetOrder(context.TODO(), &wrappers.StringValue{Value: "1"})
	if err != nil {
		log.Fatalf("Error on get Order ID: 1, %v", err)
	}

	log.Printf("Response from GetOrder 1: %v", order)
}

func callSearchOrders(c proto.OrderManagementClient) {
	searchStream, _ := c.SearchOrders(context.TODO(), &wrappers.StringValue{Value: "Arroz"})

	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Search Result:", searchOrder)
	}
}

func callUpdateOrders(c proto.OrderManagementClient) {
	updateStream, err := c.UpdateOrders(context.TODO())

	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _ %v", c, err)
	}

	// updating 1
	if err := updateStream.Send(&proto.Order{Id: "1", Price: 10}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, "Id = 1", err)
	}

	// updating 2
	if err := updateStream.Send(&proto.Order{Id: "2", Price: 100}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, "Id = 2", err)
	}
	// updating 3
	if err := updateStream.Send(&proto.Order{Id: "3", Price: 1000}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, "Id = 1", err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}

	log.Println("Update Orders Res:", updateRes)
}
