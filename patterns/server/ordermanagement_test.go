package server_test

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/iamseki/go-grpc/patterns/server"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderById(t *testing.T) {
	sut := server.NewOrderManagementServer()
	order, err := sut.GetOrder(context.TODO(), &wrappers.StringValue{Value: "1"})
	assert.Nil(t, err)

	assert.Equal(t, "Comida enlatada", order.Description)
	assert.Equal(t, "1", order.Id)
	assert.Equal(t, "São Paulo", order.Destination)
	assert.Equal(t, 3, len(order.Items))
}
