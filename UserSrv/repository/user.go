package repository

import (
	"fmt"
	"my-package/models"
	"strings"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user models.User) error
	GetQuery(user models.User) (users []models.User, err error)
	Update(user models.User, updateUser models.User) error
	Delete(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{db}
}

func (obj userRepo) Create(user models.User) error {
	return obj.db.Create(&user).Error
}

func (obj userRepo) GetQuery(user models.User) (users []models.User, err error) {
	query := "SELECT * FROM users"
	var condition []string
	if user.ID != 0 {
		condition = append(condition, fmt.Sprintf(`ID = "%v"`, user.ID))
	}
	if user.Username != "" {
		condition = append(condition, fmt.Sprintf(`username = "%v"`, user.Username))
	}
	if len(condition) > 0 {
		query = fmt.Sprintf(`
			SELECT * FROM users
			WHERE "%v"
		`, strings.Join(condition, " OR "))
	}
	err = obj.db.Raw(query).Scan(&users).Error
	return users, err
}

func (obj userRepo) Update(user models.User, updateUser models.User) error {
	// err := obj.db.Model(&user).Updates(updateUser).Error
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (obj userRepo) Delete(id uint) error {
	// err := obj.db.Delete(&models.User{}, id).Error
	// if err != nil {
	// 	return err
	// }
	return nil
}
