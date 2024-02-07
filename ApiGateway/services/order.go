package services

import (
	"context"
	"io"
	"my-package/grpcClient"
	"my-package/models"
	"my-package/utils"
	"net/http"
)

type OrderSrv interface {
	CreateOrder(user_id int, req models.CreateOrderReq) (res models.Response)
	GetOrderByUser(user_id int) (res models.Response)
	AddProduct(order_id string, req models.AddProductReq) (res models.Response)
}

type orderSrv struct {
	orderGrpc   grpcClient.OrderGrpcClient
	productGrpc grpcClient.ProductGrpcClient
}

func NewOrderSrv(orderGrpc grpcClient.OrderGrpcClient,
	productGrpc grpcClient.ProductGrpcClient) OrderSrv {
	return orderSrv{orderGrpc, productGrpc}
}

func (obj orderSrv) CreateOrder(user_id int, req models.CreateOrderReq) (res models.Response) {
	result_p, err := obj.productGrpc.GetProductID(context.Background(), &grpcClient.GetProductIDRequest{
		Id: uint64(req.ProductID),
	})
	if res := utils.HandlerErrGrpcCleint(result_p, err); res.Error {
		return res
	}

	result, err := obj.orderGrpc.CreateOrder(context.Background(), &grpcClient.CreateOrderRequest{
		UserId:    uint32(user_id),
		ProductId: uint32(req.ProductID),
		Amount:    int32(req.Amount),
	})
	if res = utils.HandlerErrGrpcCleint(result, err); res.Error {
		return res
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: res.Massage,
	}
	return res
}

func (obj orderSrv) GetOrderByUser(user_id int) (res models.Response) {
	result, err := obj.orderGrpc.GetOrderByUser(context.Background(), &grpcClient.GetOrderByUserRequest{
		UserId: uint32(user_id),
	})
	if res = utils.HandlerErrGrpcCleint(result, err); res.Error {
		return res
	}

	// connect grpc stream
	stream, err := obj.productGrpc.GetProductIDStream(context.Background())
	if err != nil {
		res = models.Response{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Massage: err.Error(),
		}
		return res
	}

	// // set config data
	data := make([]models.GetOrderByUserRes, len(result.Orders))

	var countProduct int
	for i := range result.Orders {
		countProduct += len(result.Orders[i].Products)
		// data[i].Products = make([]models.OrderProduct, len(result.Orders[i].Products))
	}

	// // set ch
	ch := make(chan *grpcClient.GetProductIDStreamResponse, countProduct)

	// send
	go func() {
		for io, order := range result.Orders {

			data[io] = models.GetOrderByUserRes{
				OrderID:     order.OrderId,
				UserID:      order.UserId,
				StatusOrder: order.StatusOrder,
				CreateTime:  order.CreateTime.AsTime(),
				UpdateTime:  order.UpdateTime.AsTime(),
			}
			data[io].Products = make([]models.OrderProduct, len(result.Orders[io].Products))

			for ip, p := range result.Orders[io].Products {

				stream.Send(&grpcClient.GetProductIDStreamRequest{
					Idx: &grpcClient.IndexOrder{
						IdxOrder:   int64(io),
						IdxProduct: int64(ip),
					},
					Id: uint64(p.ProductId),
				})
			}
		}
		stream.CloseSend()
	}()

	// receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				ch <- &grpcClient.GetProductIDStreamResponse{
					Error:   true,
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
			}
			ch <- res
		}
	}()

	for i := 0; i < countProduct; i++ {
		r := <-ch
		if r.Error {
			res = models.Response{
				Error:   r.Error,
				Status:  r.Status,
				Massage: r.Message,
			}
			return res
		}
		data[r.Idx.IdxOrder].Products[r.Idx.IdxProduct] = models.OrderProduct{
			Product: models.Product{
				ID:       uint(r.Product.Id),
				Name:     r.Product.Name,
				Price:    r.Product.Price,
				Category: r.Product.Category,
			},
			Amount: int(result.Orders[r.Idx.IdxOrder].Products[r.Idx.IdxProduct].Amount),
		}

	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
		Data:    data,
	}

	return res
}

func (obj orderSrv) AddProduct(order_id string, req models.AddProductReq) (res models.Response) {
	result_p, err := obj.productGrpc.GetProductID(context.Background(), &grpcClient.GetProductIDRequest{
		Id: uint64(req.ProductID),
	})
	if res := utils.HandlerErrGrpcCleint(result_p, err); res.Error {
		return res
	}
	result, err := obj.orderGrpc.AddProduct(context.Background(), &grpcClient.AddProductRequest{
		OrderId:   order_id,
		ProductId: uint32(req.ProductID),
		Amount:    int32(req.Amount),
	})
	if res := utils.HandlerErrGrpcCleint(result, err); res.Error {
		return res
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
	}

	return res
}
