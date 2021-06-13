package server

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/iamseki/go-grpc/patterns/proto"
)

type OrderManagementServer struct {
	orders map[string]*proto.Order
	proto.UnimplementedOrderManagementServer
}

func NewOrderManagementServer() *OrderManagementServer {
	return &OrderManagementServer{
		orders: map[string]*proto.Order{
			"1": {
				Id:          "1",
				Description: "Comida enlatada",
				Price:       214.17,
				Destination: "São Paulo",
				Items:       []string{"Espinafre", "Milho", "Leite moça"},
			},
			"2": {
				Id:          "2",
				Description: "Alimentos",
				Price:       17.9,
				Destination: "Rua João Gomes Batista",
				Items:       []string{"Arroz", "Feijão", "Macarrão"},
			},
		},
		UnimplementedOrderManagementServer: proto.UnimplementedOrderManagementServer{},
	}
}

func (s *OrderManagementServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*proto.Order, error) {
	order := s.orders[orderId.GetValue()]
	return order, nil
}

func (s *OrderManagementServer) SearchOrders(searchQuery *wrappers.StringValue, stream proto.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orders {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(
				itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf(
						"error sending message to stream : %v",
						err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}
