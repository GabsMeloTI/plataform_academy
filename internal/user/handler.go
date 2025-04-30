package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plataform_init/db"
)

type Handler struct {
	InterfaceService InterfaceService
}

func NewHandler(InterfaceService InterfaceService) *Handler {
	return &Handler{
		InterfaceService,
	}
}

func (handler *Handler) CreateUser(c echo.Context) error {
	var request db.CreateUser
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := handler.InterfaceService.CreateUser(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User created successfully")
}

func (handler *Handler) Login(c echo.Context) error {
	var request db.LoginUser
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := handler.InterfaceService.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
		"token":   token,
	})
}
