package services

import (
	"context"
	"my-package/grpcClient"
	"my-package/models"
	"my-package/utils"
)

type ProductSrv interface {
	CreateProduct(req models.CreateProductReq) (res models.Response)
	GetAllProduct() (res models.Response)
	GetProductID(id uint64) (res models.Response)
}

type productSrv struct {
	productGrpc grpcClient.ProductGrpcClient
}

func NewProductSrv(productGrpc grpcClient.ProductGrpcClient) ProductSrv {
	return productSrv{productGrpc}
}

func (obj productSrv) CreateProduct(req models.CreateProductReq) (res models.Response) {
	result, err := obj.productGrpc.CreateProduct(context.Background(), &grpcClient.CreateProductRequest{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
	})
	res = utils.HandlerErrGrpcCleint(result, err)
	if res.Error {
		return res
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
	}
	return res
}

func (obj productSrv) GetAllProduct() (res models.Response) {
	result, err := obj.productGrpc.GetAllProduct(context.Background(), &grpcClient.GetAllProductRequest{})
	res = utils.HandlerErrGrpcCleint(result, err)
	if res.Error {
		return res
	}

	var products []models.GetAllProductRes
	for i := range result.Products {
		products = append(products, models.GetAllProductRes{
			ID:       result.Products[i].Id,
			Name:     result.Products[i].Name,
			Price:    result.Products[i].Price,
			Category: result.Products[i].Category,
		})
	}
	res = models.Response{
		Error:   result.Error,
		Status:  result.Status,
		Massage: result.Message,
		Data:    products,
	}
	return res
}

func (obj productSrv) GetProductID(id uint64) (res models.Response) {
	result, err := obj.productGrpc.GetProductID(context.Background(), &grpcClient.GetProductIDRequest{
		Id: id,
	})
	res = utils.HandlerErrGrpcCleint(result, err)
	if res.Error {
		return res
	}
	res = models.Response{
		Error:   res.Error,
		Status:  res.Status,
		Massage: res.Massage,
		Data: models.GetProductIDRes{
			ID:       result.Product.Id,
			Name:     result.Product.Name,
			Price:    result.Product.Price,
			Category: result.Product.Category,
		},
	}
	return res
}
