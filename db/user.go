package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Email       string    `gorm:"type:varchar(100);not null" json:"email"`
	Password    string    `gorm:"type:varchar(256);not null" json:"password"`
	Phone       string    `gorm:"type:varchar(20)" json:"phone,omitempty"`
	DateOfBirth time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender      *string   `gorm:"type:varchar(20)" json:"gender,omitempty"`
	Role        string    `gorm:"type:varchar(20);not null" json:"role"`
	Status      string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	Weight      float64   `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	Height      float64   `gorm:"type:decimal(3,2)" json:"height,omitempty"`
	Avatar      *string   `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	//TrainerID   *uuid.UUID `gorm:"column:trainer_id" json:"trainer_id,omitempty"`
	//Trainer     *User      `gorm:"foreignKey:TrainerID;constraint:OnDelete:SET NULL" json:"trainer,omitempty"`
}

type CreateUser struct {
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	Email           string    `gorm:"type:varchar(100);not null" json:"email"`
	Password        string    `gorm:"type:varchar(256);not null" json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	Phone           string    `gorm:"type:varchar(20)" json:"phone,omitempty"`
	DateOfBirth     time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender          *string   `gorm:"type:varchar(20)" json:"gender,omitempty"`
	Role            string    `gorm:"type:varchar(20);not null" json:"role"`
	Status          string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	Weight          float64   `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	Height          float64   `gorm:"type:decimal(3,2)" json:"height,omitempty"`
	Avatar          *string   `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	gorm.Model
}

type UpdateUser struct {
	Name        *string    `json:"name,omitempty"`
	Email       *string    `json:"email,omitempty"`
	Password    *string    `json:"password,omitempty" binding:"min=8"`
	Phone       *string    `json:"phone,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Role        *string    `json:"role,omitempty" binding:"oneof=admin user"`
	Weight      *float64   `json:"weight,omitempty"`
	Height      *float64   `json:"height,omitempty"`
	Avatar      *string    `json:"avatar,omitempty"`
	gorm.Model
}

type DeleteUser struct {
	ID     uuid.UUID `json:"id" binding:"required"`
	Reason *string   `json:"reason,omitempty"`
	gorm.Model
}

type LoginUser struct {
	Email    string `gorm:"type:varchar(100);not null" json:"email"`
	Password string `gorm:"type:varchar(256);not null" json:"password"`
}

func (c *CreateUser) ParseCreateToUser() User {
	arg := User{
		Model:       gorm.Model{},
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
	}

	return arg
}
