package repository

import (
	"github.com/Giafn/goMoneySaveAndManage/internal/entity"
	"github.com/Giafn/goMoneySaveAndManage/internal/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: postgres.DB,
	}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
