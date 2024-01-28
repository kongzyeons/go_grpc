package services

import (
	context "context"
	"my-package/models"
	"my-package/repository"
	"net/http"
)

type productGrpcServer struct {
	productRepo repository.ProductRepo
}

func NewProductGrpcServer(productRepo repository.ProductRepo) ProductGrpcServer {
	return productGrpcServer{productRepo}
}

func (obj productGrpcServer) mustEmbedUnimplementedProductGrpcServer() {}

// CreateProduct implements the CreateProduct gRPC method.
func (obj productGrpcServer) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	if req == nil {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "error requie request",
		}
		return res, nil
	}

	if req.Name == "" || req.Price == 0 || req.Category == "" {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "error requie request",
		}
		return res, nil
	}

	products, err := obj.productRepo.GetQuery(models.Product{
		Name:     req.Name,
		Category: req.Category,
	})

	if len(products) > 0 {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: "error product name and category alredy exits",
		}
		return res, nil
	}

	if err != nil {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return res, nil
	}

	err = obj.productRepo.Create(models.Product{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
	})

	if err != nil {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return res, nil
	}

	res := &CreateProductResponse{
		Status:  http.StatusCreated,
		Message: "CreateProduct success",
	}
	return res, nil
}

// GetAllProduct implements the GetAllProduct gRPC method.
func (obj productGrpcServer) GetAllProduct(ctx context.Context, req *GetAllProductRequest) (*GetAllProductResponse, error) {
	// Implement the logic for GetAllProduct method using productRepo
	// Example: s.productRepo.GetAllProduct(req)

	// Return a response based on the logic
	return &GetAllProductResponse{ /* fill with appropriate values */ }, nil
}

// GetProductID implements the GetProductID gRPC method.
func (obj productGrpcServer) GetProductID(ctx context.Context, req *GetProductIDRequest) (*GetProductIDResponse, error) {
	// Implement the logic for GetProductID method using productRepo
	// Example: s.productRepo.GetProductID(req)

	// Return a response based on the logic
	return &GetProductIDResponse{ /* fill with appropriate values */ }, nil
}

// mustEmbedUnimplementedProductGrpcServer is used to embed the unimplemented
// gRPC server methods. It is required by the grpc library.
