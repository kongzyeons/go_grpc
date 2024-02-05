package repository

import "my-package/models"

type OrderRepo interface {
	Create(order models.Order) error
	GetQuery(order models.Order) (orders []models.Order, err error)
	Update(order models.Order) error
	Delete(id string) error
}
