package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/iamseki/go-grpc/patterns/proto"
)

type OrderManagementServer struct {
	orders []proto.Order
}

func NewOrderManagementServer() OrderManagementServer {
	return OrderManagementServer{
		orders: []proto.Order{
			{
				Id:          "1",
				Description: "Comida enlatada",
				Price:       214.17,
				Destination: "São Paulo",
				Items:       []string{"Espinafre", "Milho", "Leite moça"},
			},
			{
				Id:          "2",
				Description: "Alimentos",
				Price:       17.9,
				Destination: "Rua João Gomes Batista",
				Items:       []string{"Arroz", "Feijão", "Macarrão"},
			},
		},
	}
}

func (s *OrderManagementServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*proto.Order, error) {
	order := s.findOrder(orderId.String())
	return order, nil
}

func (s *OrderManagementServer) findOrder(Id string) *proto.Order {
	order := proto.Order{}

	for _, o := range s.orders {
		if o.Id == Id {
			return &o
		}
	}

	return &order
}
