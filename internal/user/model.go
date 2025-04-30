package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"plataform_init/db"
	"time"
)

type CreateUserRequest struct {
	gorm.Model
	ID          uuid.UUID          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string             `gorm:"type:varchar(100);not null" json:"name"`
	Email       string             `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password    string             `gorm:"type:varchar(256);not null" json:"password"`
	Phone       string             `gorm:"type:varchar(20)" json:"phone,omitempty"`
	DateOfBirth time.Time          `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender      *string            `gorm:"type:enum('M','F','O')" json:"gender,omitempty"`
	Role        string             `gorm:"type:enum('admin','personal','student');not null" json:"role"`
	Status      string             `gorm:"type:enum('active','inactive','banned');default:'active'" json:"status"`
	Weight      float64            `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	Height      float64            `gorm:"type:decimal(3,2)" json:"height,omitempty"`
	Avatar      *string            `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	TrainerID   *uuid.UUID         `gorm:"column:trainer_id" json:"trainer_id,omitempty"`
	Trainer     *CreateUserRequest `gorm:"foreignKey:TrainerID;constraint:OnDelete:SET NULL" json:"trainer,omitempty"`
	CreatedAt   time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
}

func (c *CreateUserRequest) ParseCreateToUser() db.User {
	arg := db.User{
		Model:       gorm.Model{},
		ID:          c.ID,
		Name:        c.Name,
		Email:       c.Email,
		Password:    c.Password,
		Phone:       c.Phone,
		DateOfBirth: c.DateOfBirth,
		Gender:      c.Gender,
		Role:        c.Role,
		Status:      c.Status,
		Weight:      c.Weight,
		Height:      c.Height,
		Avatar:      c.Avatar,
		TrainerID:   c.TrainerID,
	}

	return arg
}
