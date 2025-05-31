package token

import (
	"github.com/google/uuid"
	"time"
)

type JwtPayload struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
}
