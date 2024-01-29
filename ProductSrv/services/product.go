package services

import (
	context "context"
	"log"
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
		log.Println(res.Message)
		return res, nil
	}

	if req.Name == "" || req.Price == 0 || req.Category == "" {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "error requie request",
		}
		log.Println(res.Message)
		return res, nil
	}

	products, err := obj.productRepo.GetQuery(models.Product{
		Name: req.Name,
	})

	if len(products) > 0 {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: "error product name and category alredy exits",
		}
		log.Println(res.Message)
		return res, nil
	}

	if err != nil {
		res := &CreateProductResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
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
		log.Println(res.Message)
		return res, nil
	}

	res := &CreateProductResponse{
		Status:  http.StatusCreated,
		Message: "CreateProduct success",
	}
	log.Println(res.Message)
	return res, nil
}

// GetAllProduct implements the GetAllProduct gRPC method.
func (obj productGrpcServer) GetAllProduct(ctx context.Context, req *GetAllProductRequest) (*GetAllProductResponse, error) {

	products, err := obj.productRepo.GetQuery(models.Product{})
	if err != nil {
		res := &GetAllProductResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	if len(products) == 0 {
		res := &GetAllProductResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "product not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	res := &GetAllProductResponse{
		Error:   false,
		Status:  http.StatusOK,
		Message: "GetAllProduct success",
	}
	res.Products = make([]*Product, 0, len(products))
	for i := range products {
		res.Products = append(res.Products, &Product{
			Id:       uint64(products[i].ID),
			Name:     products[i].Name,
			Price:    products[i].Price,
			Category: products[i].Category,
		})
	}

	log.Println(res.Message)
	return res, nil
}

// GetProductID implements the GetProductID gRPC method.
func (obj productGrpcServer) GetProductID(ctx context.Context, req *GetProductIDRequest) (*GetProductIDResponse, error) {
	if req == nil {
		res := &GetProductIDResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "error requie request",
		}
		log.Println(res.Message)
		return res, nil
	}

	if req.Id == 0 {
		res := &GetProductIDResponse{
			Error:   true,
			Status:  http.StatusBadRequest,
			Message: "error requie request",
		}
		log.Println(res.Message)
		return res, nil
	}

	products, err := obj.productRepo.GetQuery(models.Product{
		ID: uint(req.Id),
	})

	if len(products) == 0 {
		res := &GetProductIDResponse{
			Error:   true,
			Status:  http.StatusNotFound,
			Message: "product id not found",
		}
		log.Println(res.Message)
		return res, nil
	}

	if err != nil {
		res := &GetProductIDResponse{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		log.Println(res.Message)
		return res, nil
	}

	res := &GetProductIDResponse{
		Error:   false,
		Status:  http.StatusOK,
		Message: "GetProductID success",
		Product: &Product{
			Id:       uint64(products[0].ID),
			Name:     products[0].Name,
			Price:    products[0].Price,
			Category: products[0].Category,
		}}
	log.Println(res.Message)
	return res, nil
}

// mustEmbedUnimplementedProductGrpcServer is used to embed the unimplemented
// gRPC server methods. It is required by the grpc library.
