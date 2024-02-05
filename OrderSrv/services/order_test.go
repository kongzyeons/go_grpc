package services_test

import (
	"context"
	"errors"
	"my-package/models"
	"my-package/repository"
	"my-package/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		nameTest  string
		req       *services.CreateOrderRequest
		errCreate error
		expect    bool
	}{
		{nameTest: "test CreateOrder success", req: &services.CreateOrderRequest{UserId: 1, ProductId: 1, Amount: 1}, errCreate: nil, expect: false},
		{nameTest: "test error repo Create", req: &services.CreateOrderRequest{UserId: 1, ProductId: 1, Amount: 1}, errCreate: errors.New(""), expect: true},
		{nameTest: "test requie request", req: &services.CreateOrderRequest{}, errCreate: nil, expect: true},
		{nameTest: "test requie request is nill", req: nil, errCreate: nil, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {

			orderRepo := repository.NewOrderRepoMock()
			orderRepo.On("Create", mock.AnythingOfType("models.Order")).Return(
				tests[i].errCreate,
			)
			orderSrv := services.NewOrderGrpcServer(orderRepo)
			res, _ := orderSrv.CreateOrder(context.Background(), tests[i].req)
			assert.Equal(t, tests[i].expect, res.Error)
		})
	}
}

func TestGetallOrder(t *testing.T) {
	tests := []struct {
		nameTest    string
		errGetquery error
		orders      []models.Order
		expect      bool
	}{
		{nameTest: "test GetallOrder success", errGetquery: nil, orders: []models.Order{{}}, expect: false},
		{nameTest: "test error repo getquery", errGetquery: errors.New(""), orders: []models.Order{{}}, expect: true},
		{nameTest: "test orders not found", errGetquery: nil, orders: []models.Order{}, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {

			orderRepo := repository.NewOrderRepoMock()
			orderRepo.On("GetQuery", mock.AnythingOfType("models.Order")).Return(
				tests[i].orders,
				tests[i].errGetquery,
			)
			orderSrv := services.NewOrderGrpcServer(orderRepo)
			res, _ := orderSrv.GetallOrder(context.Background(), &services.GetallOrderRequest{})
			assert.Equal(t, tests[i].expect, res.Error)

		})
	}
}

func TestGetOrderID(t *testing.T) {
	tests := []struct {
		nameTest    string
		errGetquery error
		orders      []models.Order
		expect      bool
	}{
		{nameTest: "test GetOrderID success", errGetquery: nil, orders: []models.Order{{}}, expect: false},
		{nameTest: "test error repo getquery", errGetquery: errors.New(""), orders: []models.Order{{}}, expect: true},
		{nameTest: "test orders not found", errGetquery: nil, orders: []models.Order{}, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			orderRepo := repository.NewOrderRepoMock()
			orderRepo.On("GetQuery", mock.AnythingOfType("models.Order")).Return(
				tests[i].orders,
				tests[i].errGetquery,
			)
			orderSrv := services.NewOrderGrpcServer(orderRepo)
			res, _ := orderSrv.GetOrderID(context.Background(), &services.GetOrderIDRequest{})

			assert.Equal(t, tests[i].expect, res.Error)

		})
	}
}
