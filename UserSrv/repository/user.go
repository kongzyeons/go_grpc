package repository

import "my-package/models"

type UserRepo interface {
	Create(user models.User) error
	GetQuery(user models.User) (users []models.User, err error)
	Update(user models.User, updateUser models.User) error
	Delete(id uint) error
}
