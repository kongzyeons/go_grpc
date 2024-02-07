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
		req         *services.GetOrderIDRequest
		errGetquery error
		orders      []models.Order
		expect      bool
	}{
		{nameTest: "test GetOrderID success", req: &services.GetOrderIDRequest{OrderId: "60211c79d2b4eb29a4588313"}, errGetquery: nil, orders: []models.Order{{}}, expect: false},
		{nameTest: "test error repo getquery", req: &services.GetOrderIDRequest{OrderId: "60211c79d2b4eb29a4588313"}, errGetquery: errors.New(""), orders: []models.Order{{}}, expect: true},
		{nameTest: "test orders not found", req: &services.GetOrderIDRequest{OrderId: "60211c79d2b4eb29a4588313"}, errGetquery: nil, orders: []models.Order{}, expect: true},
		{nameTest: "test requie request", req: &services.GetOrderIDRequest{OrderId: "60211c79d2b4eb29a4588313"}, errGetquery: nil, orders: []models.Order{{}}, expect: false},
		{nameTest: "test requie request", req: &services.GetOrderIDRequest{}, errGetquery: nil, orders: []models.Order{{}}, expect: true},
		{nameTest: "test requie request", req: nil, errGetquery: nil, orders: []models.Order{{}}, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			orderRepo := repository.NewOrderRepoMock()
			orderRepo.On("GetQuery", mock.AnythingOfType("models.Order")).Return(
				tests[i].orders,
				tests[i].errGetquery,
			)
			orderSrv := services.NewOrderGrpcServer(orderRepo)
			res, _ := orderSrv.GetOrderID(context.Background(), tests[i].req)

			assert.Equal(t, tests[i].expect, res.Error)

		})
	}
}

func TestGetOrderByUser(t *testing.T) {
	tests := []struct {
		nameTest    string
		req         *services.GetOrderByUserRequest
		errGetquery error
		orders      []models.Order
		expect      bool
	}{
		{nameTest: "test GetOrderByUser success", req: &services.GetOrderByUserRequest{UserId: 1}, errGetquery: nil, orders: []models.Order{{}}, expect: false},
		{nameTest: "test error repo getquery", req: &services.GetOrderByUserRequest{UserId: 1}, errGetquery: errors.New(""), orders: []models.Order{{}}, expect: true},
		{nameTest: "test order not found", req: &services.GetOrderByUserRequest{UserId: 1}, errGetquery: nil, orders: []models.Order{}, expect: true},
		{nameTest: "test request requie", req: &services.GetOrderByUserRequest{}, errGetquery: nil, orders: []models.Order{{}}, expect: true},
		{nameTest: "test request requie is nill", req: nil, errGetquery: nil, orders: []models.Order{{}}, expect: true},
	}
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			orderRep := repository.NewOrderRepoMock()
			orderRep.On("GetQuery", mock.AnythingOfType("models.Order")).Return(
				tt.orders,
				tt.errGetquery,
			)

			orderSrv := services.NewOrderGrpcServer(orderRep)
			res, _ := orderSrv.GetOrderByUser(context.Background(), tt.req)
			assert.Equal(t, tt.expect, res.Error)
		})
	}
}

func TestAddProduct(t *testing.T) {
	tests := []struct {
		nameTest    string
		req         *services.AddProductRequest
		orders      []models.Order
		errGetQuery error
		errUpdate   error
		expect      bool
	}{
		{nameTest: "test AddProduct success", req: &services.AddProductRequest{OrderId: "60211c79d2b4eb29a4588313", ProductId: 1, Amount: 1}, orders: []models.Order{{}}, errGetQuery: nil, errUpdate: nil, expect: false},
		{nameTest: "test error repo update", req: &services.AddProductRequest{OrderId: "60211c79d2b4eb29a4588313", ProductId: 1, Amount: 1}, orders: []models.Order{{}}, errGetQuery: nil, errUpdate: errors.New(""), expect: true},
		{nameTest: "test AddProduct success", req: &services.AddProductRequest{OrderId: "60211c79d2b4eb29a4588313", ProductId: 1, Amount: 1}, orders: []models.Order{{}}, errGetQuery: errors.New(""), errUpdate: nil, expect: true},
		{nameTest: "test order id not found", req: &services.AddProductRequest{OrderId: "60211c79d2b4eb29a4588313", ProductId: 1, Amount: 1}, orders: []models.Order{}, errGetQuery: nil, errUpdate: nil, expect: true},
		{nameTest: "test requie request", req: &services.AddProductRequest{}, orders: []models.Order{{}}, errGetQuery: nil, errUpdate: nil, expect: true},
		{nameTest: "test requie request is null", req: nil, orders: []models.Order{{}}, errGetQuery: nil, errUpdate: nil, expect: true},
	}
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			orderRepo := repository.NewOrderRepoMock()
			orderRepo.On("GetQuery", mock.AnythingOfType("models.Order")).Return(
				tt.orders,
				tt.errGetQuery,
			)
			orderRepo.On("Update", mock.AnythingOfType("models.Order")).Return(
				tt.errUpdate,
			)
			orderSrv := services.NewOrderGrpcServer(orderRepo)
			res, _ := orderSrv.AddProduct(context.Background(), tt.req)
			assert.Equal(t, tt.expect, res.Error)
		})
	}
}
