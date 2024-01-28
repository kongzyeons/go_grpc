package repository

import (
	"my-package/models"

	"github.com/stretchr/testify/mock"
)

type productRepoMock struct {
	mock.Mock
}

func NewProductRepoMock() *productRepoMock {
	return &productRepoMock{}
}

// Create mocks the Create method of the ProductRepo interface.
func (m *productRepoMock) Create(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// GetQuery mocks the GetQuery method of the ProductRepo interface.
func (m *productRepoMock) GetQuery(product models.Product) ([]models.Product, error) {
	args := m.Called(product)
	return args.Get(0).([]models.Product), args.Error(1)
}

// Update mocks the Update method of the ProductRepo interface.
func (m *productRepoMock) Update(product models.Product, update models.Product) error {
	args := m.Called(product, update)
	return args.Error(0)
}

// Delete mocks the Delete method of the ProductRepo interface.
func (m *productRepoMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
