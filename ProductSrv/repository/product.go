package repository

import "my-package/models"

type ProductRepo interface {
	Create(product models.Product) error
	GetQuery(product models.Product) (products []models.Product, err error)
	Update(product models.Product, update models.Product) error
	Delete(id uint) error
}
