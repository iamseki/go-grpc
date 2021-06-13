package server

import (
	"context"

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
