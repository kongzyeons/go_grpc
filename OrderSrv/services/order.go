package services

import (
	context "context"
	"fmt"
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
			ProductID uint `bson:"product_id" json:"product_id"`
			Amount    int  `bson:"amount" json:"amount"`
		}{
			{ProductID: uint(req.ProductId), Amount: int(req.Amount)},
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
	if req == nil {
		res := &GetOrderIDResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "request requie",
		}
		log.Println(res.Message)
		return res, nil
	}

	if req.OrderId == "" {
		res := &GetOrderIDResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "request requie",
		}
		log.Println(res.Message)
		return res, nil
	}

	order_id, err := primitive.ObjectIDFromHex(req.OrderId)
	if err != nil {
		res := &GetOrderIDResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

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
		OrderId:    orders[0].ID.Hex(),
		UserId:     uint32(orders[0].UserID),
		CreateTime: timestamppb.New(orders[0].CreateTime),
		UpdateTime: timestamppb.New(orders[0].UpdateTime),
	}
	res.Order.Products = make([]*Order_Product, 0, len(orders[0].Products))
	for _, p := range orders[0].Products {
		res.Order.Products = append(res.Order.Products, &Order_Product{
			ProductId: uint32(p.ProductID),
			Amount:    int32(p.Amount),
		})
	}
	log.Println(res.Message)
	return res, nil
}

func (obj orderGrpcServer) GetOrderByUser(ctx context.Context, req *GetOrderByUserRequest) (*GetOrderByUserResponse, error) {
	if req == nil {
		res := &GetOrderByUserResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "request requie",
		}
		log.Println(res.Message)
		return res, nil
	}

	if req.UserId == 0 {
		res := &GetOrderByUserResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "request requie",
		}
		log.Println(res.Message)
		return res, nil
	}

	orders, err := obj.orderRepo.GetQuery(models.Order{
		UserID: uint(req.UserId),
	})
	if err != nil {
		res := &GetOrderByUserResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	if len(orders) == 0 {
		res := &GetOrderByUserResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "orders not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	res := &GetOrderByUserResponse{
		Error:   false,
		Status:  http.StatusOK,
		Message: "GetOrderByUser success",
		Orders:  nil,
	}
	res.Orders = make([]*Order, 0, len(orders))
	for i := range orders {
		res.Orders = append(res.Orders, &Order{
			OrderId:     orders[i].ID.Hex(),
			UserId:      uint32(orders[i].UserID),
			StatusOrder: orders[i].StatusOrder,
			CreateTime:  timestamppb.New(orders[i].CreateTime),
			UpdateTime:  timestamppb.New(orders[i].UpdateTime),
		})
		res.Orders[i].Products = make([]*Order_Product, 0, len(orders[i].Products))
		for _, product := range orders[i].Products {
			res.Orders[i].Products = append(res.Orders[i].Products, &Order_Product{
				ProductId: uint32(product.ProductID),
				Amount:    int32(product.Amount),
			})
		}
	}

	log.Println(res.Message)
	return res, nil
}

func (obj orderGrpcServer) AddProduct(ctx context.Context, req *AddProductRequest) (*AddProductResponse, error) {

	if req == nil {
		res := &AddProductResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "error requie request",
		}
		log.Println(res.Message)
		return res, nil
	}

	if req.OrderId == "" || req.ProductId == 0 || req.Amount == 0 {
		res := &AddProductResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "error requie request",
		}
		log.Println(res.Message)
		return res, nil
	}

	objectID, err := primitive.ObjectIDFromHex(req.OrderId)
	if err != nil {
		res := &AddProductResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	orders, err := obj.orderRepo.GetQuery(models.Order{
		ID: objectID,
	})
	if err != nil {
		res := &AddProductResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}
	if len(orders) == 0 {
		res := &AddProductResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "order id not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	orders[0].Products = append(orders[0].Products, struct {
		ProductID uint "bson:\"product_id\" json:\"product_id\""
		Amount    int  "bson:\"amount\" json:\"amount\""
	}{
		ProductID: uint(req.ProductId),
		Amount:    int(req.Amount),
	})
	fmt.Println("-------")
	fmt.Println(orders[0].UpdateTime)
	orders[0].UpdateTime = time.Now()
	fmt.Println(orders[0].UpdateTime)
	fmt.Println("-------")

	err = obj.orderRepo.Update(orders[0])
	if err != nil {
		res := &AddProductResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	res := &AddProductResponse{
		Error:   false,
		Status:  http.StatusOK,
		Message: "AddProduct success",
	}
	log.Println(res.Message)
	return res, nil
}

func (obj orderGrpcServer) DeleteOrderID(ctx context.Context, req *DeleteOrderIDRequest) (*DeleteOrderIDResponse, error) {
	// Implement your logic here, using obj.orderRepo or other necessary dependencies
	return nil, nil
}
