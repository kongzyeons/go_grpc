package repository

import (
	"my-package/models"

	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

func NewuserRepoMock() *userRepoMock {
	return &userRepoMock{}
}

// Create mocks the Create method
func (m *userRepoMock) Create(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// GetQuery mocks the GetQuery method
func (m *userRepoMock) GetQuery(user models.User) ([]models.User, error) {
	args := m.Called(user)
	return args.Get(0).([]models.User), args.Error(1)
}

// Update mocks the Update method
func (m *userRepoMock) Update(user models.User, updateUser models.User) error {
	args := m.Called(user, updateUser)
	return args.Error(0)
}

// Delete mocks the Delete method
func (m *userRepoMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
