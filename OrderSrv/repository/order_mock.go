package repository

import (
	"my-package/models"

	"github.com/stretchr/testify/mock"
)

type orderRepoMock struct {
	mock.Mock
}

func NewOrderRepoMock() *orderRepoMock {
	return &orderRepoMock{}
}

func (orm *orderRepoMock) Create(order models.Order) error {
	// Implement the mocked behavior for Create method
	args := orm.Called(order)
	return args.Error(0)
}

func (orm *orderRepoMock) GetQuery(order models.Order) ([]models.Order, error) {
	// Implement the mocked behavior for GetQuery method
	args := orm.Called(order)
	return args.Get(0).([]models.Order), args.Error(1)
}

func (orm *orderRepoMock) Update(order models.Order) error {
	// Implement the mocked behavior for Update method
	args := orm.Called(order)
	return args.Error(0)
}

func (orm *orderRepoMock) Delete(id string) error {
	// Implement the mocked behavior for Delete method
	args := orm.Called(id)
	return args.Error(0)
}
