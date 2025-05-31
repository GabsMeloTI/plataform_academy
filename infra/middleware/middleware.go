package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"plataform_init/infra/token"
	"strings"
)

func CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerToken := c.Request().Header.Get("Authorization")
		tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

		jwtMaker, err := token.NewJwtMaker(os.Getenv("SIGNATURE_STRING"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		}

		payload, err := jwtMaker.VerifyToken(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		c.Set("token_user_id", payload.UserID)
		c.Set("token_username", payload.Username)
		c.Set("token_user_email", payload.Email)
		c.Set("token_expiry_at", payload.ExpiresAt)

		return next(c)
	}
}
