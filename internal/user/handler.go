package user

import "C"
import (
	"github.com/labstack/echo/v4"
	"net/http"
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
	var request CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := handler.InterfaceService.CreateUser(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Success")
}
