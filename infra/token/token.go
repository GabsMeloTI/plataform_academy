package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

type JwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (*JwtMaker, error) {
	if len(secretKey) == 0 {
		return nil, fmt.Errorf("invalid key size: key must not be empty")
	}

	return &JwtMaker{
		secretKey: secretKey,
	}, nil
}

func (maker *JwtMaker) CreateToken(userID uuid.UUID, username, email string, expireAt time.Time) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":    userID,
		"username":   username,
		"email":      email,
		"issued_at":  time.Now().Unix(),
		"expired_at": expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func (maker *JwtMaker) VerifyToken(tokenStr string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(maker.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	rawUserID, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("user_id claim is not a string")
	}
	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id in token: %w", err)
	}
	rawExp, ok := claims["expired_at"].(float64)
	if !ok {
		return nil, fmt.Errorf("expired_at claim is not a number")
	}
	expiryTime := time.Unix(int64(rawExp), 0)
	if time.Now().After(expiryTime) {
		return nil, ErrExpiredToken
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("username claim is not a string")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("email claim is not a string")
	}

	return &JwtPayload{
		UserID:    userID,
		Username:  username,
		Email:     email,
		ExpiresAt: expiryTime,
	}, nil
}
