package token

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

func GetPayloadToken(c echo.Context) PayloadDTO {
	strUserID, _ := c.Get("token_user_id").(uuid.UUID)
	strUsername, _ := c.Get("token_username").(string)
	strEmail, _ := c.Get("token_email").(string)
	strExpiryAt, _ := c.Get("token_expiry_at").(time.Time)

	return PayloadDTO{
		UserID:    strUserID,
		Username:  strUsername,
		Email:     strEmail,
		ExpiresAt: strExpiryAt,
	}
}
