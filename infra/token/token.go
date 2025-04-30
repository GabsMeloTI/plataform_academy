package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"time"
)

func GeneratedToken(c echo.Context) error {
	jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustomClaims{
		Name:  "test",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	return nil
}
