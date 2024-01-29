package repository

import (
	"my-package/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(product models.Product) error
	GetQuery(product models.Product) (products []models.Product, err error)
	Update(product models.Product, update models.Product) error
	Delete(id uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return productRepo{db}
}

func (obj productRepo) Create(product models.Product) error {
	return obj.db.Create(&product).Error
}

func (obj productRepo) GetQuery(product models.Product) (products []models.Product, err error) {
	err = obj.db.Where(&product).Find(&products).Error
	return products, err
}

func (obj productRepo) Update(product models.Product, update models.Product) error {
	return obj.db.Model(&product).Updates(update).Error
}

func (obj productRepo) Delete(id uint) error {
	return obj.db.Delete(&models.Product{ID: id}).Error
}
