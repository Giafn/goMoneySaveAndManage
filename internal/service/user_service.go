package service

import (
	"errors"

	"github.com/Giafn/goMoneySaveAndManage/internal/entity"
	"github.com/Giafn/goMoneySaveAndManage/internal/repository"
	"github.com/Giafn/goMoneySaveAndManage/pkg/hash"
	"github.com/Giafn/goMoneySaveAndManage/pkg/jwt"

)

type UserService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(username, password string) error {
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username: username,
		Password: hashedPassword,
	}

	return s.userRepo.CreateUser(user)
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
