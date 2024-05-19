package handler

import (
	"net/http"

	"github.com/Giafn/goMoneySaveAndManage/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := h.UserService.Register(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "registration successful"})
}

func (h *UserHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := h.UserService.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
