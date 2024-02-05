package services

import (
	context "context"
	"log"
	"my-package/models"
	"my-package/repository"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type orderGrpcServer struct {
	orderRepo repository.OrderRepo
}

func NewOrderGrpcServer(orderRepo repository.OrderRepo) OrderGrpcServer {
	return orderGrpcServer{orderRepo}
}

func (obj orderGrpcServer) mustEmbedUnimplementedOrderGrpcServer() {
	// Implement this method if required by your gRPC server setup
}

func (obj orderGrpcServer) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	if req == nil {
		res := &CreateOrderResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie request",
		}
		log.Println(res.Message)
		return res, nil
	}

	if req.UserId == 0 || req.ProductId == 0 || req.Amount == 0 {
		res := &CreateOrderResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "requie request",
		}
		log.Println(res.Message)
		return res, nil
	}

	err := obj.orderRepo.Create(models.Order{
		UserID: uint(req.UserId),
		Products: []struct {
			ProducrtID uint `bson:"product_id"`
			Amount     int  `bson:"amount"`
		}{
			{ProducrtID: uint(req.ProductId), Amount: int(req.Amount)},
		},
		StatusOrder: "create",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	})
	if err != nil {
		res := &CreateOrderResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	res := &CreateOrderResponse{
		Error:   false,
		Status:  http.StatusCreated,
		Message: "CreateOrder success",
	}
	log.Println(res.Message)
	return res, nil
}

func (obj orderGrpcServer) GetallOrder(ctx context.Context, req *GetallOrderRequest) (*GetallOrderResponse, error) {
	orders, err := obj.orderRepo.GetQuery(models.Order{})
	if err != nil {
		res := &GetallOrderResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	if len(orders) == 0 {
		res := &GetallOrderResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "order not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	res := &GetallOrderResponse{
		Error:   false,
		Status:  http.StatusOK,
		Message: "GetallOrder success",
	}
	res.Orders = make([]*Order, 0, len(orders))
	for i := range orders {
		res.Orders = append(res.Orders, &Order{
			OrderId:     orders[i].ID.String(),
			UserId:      uint32(orders[i].UserID),
			Products:    nil,
			StatusOrder: orders[i].StatusOrder,
			CreateTime:  timestamppb.New(orders[i].CreateTime),
			UpdateTime:  timestamppb.New(orders[i].UpdateTime),
		})
	}

	log.Println(res.Message)
	return res, nil
}

func (obj orderGrpcServer) GetOrderID(ctx context.Context, req *GetOrderIDRequest) (*GetOrderIDResponse, error) {

	order_id, _ := primitive.ObjectIDFromHex(req.OrderId)
	orders, err := obj.orderRepo.GetQuery(models.Order{
		ID: order_id,
	})
	if err != nil {
		res := &GetOrderIDResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}
	if len(orders) == 0 {
		res := &GetOrderIDResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "error order not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	res := &GetOrderIDResponse{
		Error:   false,
		Status:  http.StatusOK,
		Message: "GetOrderID success",
	}
	res.Order = &Order{
		OrderId:    orders[0].ID.String(),
		UserId:     uint32(orders[0].UserID),
		CreateTime: timestamppb.New(orders[0].CreateTime),
		UpdateTime: timestamppb.New(orders[0].UpdateTime),
	}
	res.Order.Products = make([]*Order_Product, 0, len(orders[0].Products))
	for _, p := range orders[0].Products {
		res.Order.Products = append(res.Order.Products, &Order_Product{
			ProductId: uint32(p.ProducrtID),
			Amount:    int32(p.Amount),
		})
	}
	log.Println(res.Message)
	return res, nil
}

func (obj orderGrpcServer) AddProduct(ctx context.Context, req *AddProductRequest) (*AddProductResponse, error) {
	// Implement your logic here, using obj.orderRepo or other necessary dependencies
	return nil, nil
}

func (obj orderGrpcServer) DeleteOrderID(ctx context.Context, req *DeleteOrderIDRequest) (*DeleteOrderIDResponse, error) {
	// Implement your logic here, using obj.orderRepo or other necessary dependencies
	return nil, nil
}
