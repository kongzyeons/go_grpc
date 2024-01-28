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
