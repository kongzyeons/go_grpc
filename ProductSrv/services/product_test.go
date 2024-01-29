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

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		nameTest        string
		req             *services.CreateProductRequest
		errRepoGetQuery error
		products        []models.Product
		errRepoCreate   error
		expect          bool
	}{
		{nameTest: "test CreateProduct success", req: &services.CreateProductRequest{Name: "A", Price: 100, Category: "C"}, errRepoGetQuery: nil, products: []models.Product{}, errRepoCreate: nil, expect: false},
		{nameTest: "test error reopo Create", req: &services.CreateProductRequest{Name: "A", Price: 100, Category: "C"}, errRepoGetQuery: nil, products: []models.Product{}, errRepoCreate: errors.New(""), expect: true},
		{nameTest: "test error repo Getquery", req: &services.CreateProductRequest{Name: "A", Price: 100, Category: "C"}, errRepoGetQuery: errors.New(""), products: []models.Product{}, errRepoCreate: nil, expect: true},
		{nameTest: "test name and cate alredy", req: &services.CreateProductRequest{Name: "A", Price: 100, Category: "C"}, errRepoGetQuery: nil, products: []models.Product{{}}, errRepoCreate: nil, expect: true},
		{nameTest: "test requie request", req: &services.CreateProductRequest{}, errRepoGetQuery: nil, products: []models.Product{}, errRepoCreate: nil, expect: true},
		{nameTest: "test requie request is nill", req: nil, errRepoGetQuery: nil, products: []models.Product{}, errRepoCreate: nil, expect: true},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {

			productRepo := repository.NewProductRepoMock()
			productRepo.On("GetQuery", mock.AnythingOfType("models.Product")).Return(
				tests[i].products,
				tests[i].errRepoGetQuery,
			)

			productRepo.On("Create", mock.AnythingOfType("models.Product")).Return(
				tests[i].errRepoCreate,
			)
			productSrv := services.NewProductGrpcServer(productRepo)

			res, _ := productSrv.CreateProduct(context.Background(), tests[i].req)

			assert.Equal(t, tests[i].expect, res.Error)

		})
	}
}

func TestGetAllProduct(t *testing.T) {
	tests := []struct {
		nameTest    string
		errGetquery error
		products    []models.Product
		expect      bool
	}{
		{nameTest: "test GetAllProduct success", errGetquery: nil, products: []models.Product{{}}, expect: false},
		{nameTest: "test error repo Getquery", errGetquery: errors.New(""), products: []models.Product{{}}, expect: true},
		{nameTest: "test product not found", errGetquery: nil, products: []models.Product{}, expect: true},
	}

	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {

			productRepo := repository.NewProductRepoMock()
			productRepo.On("GetQuery", mock.AnythingOfType("models.Product")).Return(
				tests[i].products,
				tests[i].errGetquery,
			)
			productSrv := services.NewProductGrpcServer(productRepo)
			res, _ := productSrv.GetAllProduct(context.Background(), &services.GetAllProductRequest{})
			assert.Equal(t, tests[i].expect, res.Error)

		})
	}
}

func TestGetProductID(t *testing.T) {
	tests := []struct {
		nameTest      string
		req           *services.GetProductIDRequest
		errGetquery   error
		products      []models.Product
		expectProduct *services.Product
		expect        bool
	}{
		{nameTest: "test GetProductID success", req: &services.GetProductIDRequest{Id: 1}, errGetquery: nil, products: []models.Product{{}}, expect: false},
		{nameTest: "test error repo Getquery", req: &services.GetProductIDRequest{Id: 1}, errGetquery: errors.New(""), products: []models.Product{{}}, expect: true},
		{nameTest: "test get product id notfound", req: &services.GetProductIDRequest{Id: 1}, errGetquery: nil, products: []models.Product{}, expect: true},
		{nameTest: "test requie request", req: &services.GetProductIDRequest{}, errGetquery: nil, products: []models.Product{{}}, expect: true},
		{nameTest: "test requie request is nill", req: nil, errGetquery: nil, products: []models.Product{{}}, expect: true},
		{nameTest: "test check product", req: &services.GetProductIDRequest{Id: 1}, errGetquery: nil, products: []models.Product{{ID: 1}}, expectProduct: &services.Product{Id: 1}, expect: false},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			productRepo := repository.NewProductRepoMock()
			productRepo.On("GetQuery", mock.AnythingOfType("models.Product")).Return(
				tests[i].products,
				tests[i].errGetquery,
			)
			productSrv := services.NewProductGrpcServer(productRepo)
			res, _ := productSrv.GetProductID(context.Background(), tests[i].req)

			if tests[i].nameTest == "test check product" {
				assert.Equal(t, tests[i].expectProduct, res.Product)
				return
			}

			assert.Equal(t, tests[i].expect, res.Error)
		})
	}

}
