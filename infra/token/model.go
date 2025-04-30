package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtCustomClaims struct {
	UserId   uuid.UUID `json:"user_id"`
	Username string    `json:"name"`

	jwt.RegisteredClaims
}
