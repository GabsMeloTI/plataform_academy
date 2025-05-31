package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"plataform_init/db"
	"plataform_init/internal/token"
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

	login, err := handler.InterfaceService.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
		"login":   login,
	})
}

func (handler *Handler) UpdateAvatar(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("failed to parse multipart form"))
	}

	payload := token.GetPayloadToken(c)
	result, err := handler.InterfaceService.UpdateAvatar(c.Request().Context(), form, payload.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"avatar": result})
}

func (handler *Handler) GetUserById(c echo.Context) error {
	payload := token.GetPayloadToken(c)

	result, err := handler.InterfaceService.GetUserById(c.Request().Context(), payload.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (handler *Handler) GetUsersByRole(c echo.Context) error {
	role := c.Param("role")
	if role == "" {
		return c.JSON(http.StatusBadRequest, errors.New("role parameter is required"))
	}

	result, err := handler.InterfaceService.GetUsersByRole(c.Request().Context(), role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
