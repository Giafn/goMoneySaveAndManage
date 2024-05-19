package router

import (
	"github.com/Giafn/goMoneySaveAndManage/internal/handler"
	"github.com/Giafn/goMoneySaveAndManage/internal/repository"
	"github.com/Giafn/goMoneySaveAndManage/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
}
