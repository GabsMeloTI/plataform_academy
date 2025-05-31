package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Username    string    `gorm:"type:varchar(100);not null" json:"username"`
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
	gorm.Model
}

type CreateUser struct {
	Name            string    `gorm:"type:varchar(100);not null" json:"name"`
	Email           string    `gorm:"type:varchar(100);not null" json:"email"`
	Username        string    `gorm:"type:varchar(100);not null" json:"username"`
	Password        string    `gorm:"type:varchar(256);not null" json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	Phone           string    `gorm:"type:varchar(20)" json:"phone,omitempty"`
	DateOfBirth     time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender          *string   `gorm:"type:varchar(20)" json:"gender,omitempty"`
	Role            string    `gorm:"type:varchar(20);not null" json:"role"`
	Status          string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	Weight          float64   `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	Height          float64   `gorm:"type:decimal(3,2)" json:"height,omitempty"`
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

type UpdateAvatar struct {
	Avatar string `form:"avatar,omitempty" gorm:"column:avatar"`
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

type UserList struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Username    string    `gorm:"type:varchar(100);not null" json:"username"`
	Email       string    `gorm:"type:varchar(100);not null" json:"email"`
	Phone       string    `gorm:"type:varchar(20)" json:"phone,omitempty"`
	DateOfBirth time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender      *string   `gorm:"type:varchar(20)" json:"gender,omitempty"`
	Role        string    `gorm:"type:varchar(20);not null" json:"role"`
	Status      string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	Weight      float64   `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	Height      float64   `gorm:"type:decimal(3,2)" json:"height,omitempty"`
	Avatar      *string   `gorm:"type:varchar(255)" json:"avatar,omitempty"`
}

type UsersCompleteResponse struct {
	UserID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	Name               string    `gorm:"type:varchar(100);not null" json:"name"`
	Username           string    `gorm:"type:varchar(100);not null" json:"username"`
	Email              string    `gorm:"type:varchar(100);not null" json:"email"`
	Phone              string    `gorm:"type:varchar(20)" json:"phone,omitempty"`
	DateOfBirth        time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Gender             *string   `gorm:"type:varchar(20)" json:"gender,omitempty"`
	Role               string    `gorm:"type:varchar(20);not null" json:"role"`
	Status             string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	Weight             float64   `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	Height             float64   `gorm:"type:decimal(3,2)" json:"height,omitempty"`
	Cref               string    `gorm:"type:varchar(255)" json:"cref,omitempty"`
	Avatar             *string   `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	TrainerID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"trainer_id"`
	College            string    `gorm:"type:varchar(255)" json:"college,omitempty"`
	GraduationDate     time.Time `gorm:"type:date" json:"graduation_date,omitempty"`
	ExperienceDuration string    `gorm:"type:varchar(50)" json:"experience_duration,omitempty"`
	Specialty          string    `gorm:"type:varchar(255)" json:"specialty,omitempty"`
	Certifications     string    `gorm:"type:varchar(255)" json:"certifications,omitempty"`
	PrimarySpecialty   string    `gorm:"type:varchar(255)" json:"primary_specialty,omitempty"`
	Bio                string    `gorm:"type:text" json:"bio,omitempty"`
}

func (c *CreateUser) ParseCreateToUser() User {
	arg := User{
		Name:        c.Name,
		Email:       c.Email,
		Username:    c.Username,
		Password:    c.Password,
		Phone:       c.Phone,
		DateOfBirth: c.DateOfBirth,
		Gender:      c.Gender,
		Role:        c.Role,
		Status:      c.Status,
		Weight:      c.Weight,
		Height:      c.Height,
	}

	return arg
}

func (c *UserList) ParseUserListToUser(result User) {
	c.ID = result.ID
	c.Name = result.Name
	c.Email = result.Email
	c.Username = result.Username
	c.Phone = result.Phone
	c.DateOfBirth = result.DateOfBirth
	c.Gender = result.Gender
	c.Role = result.Role
	c.Status = result.Status
	c.Weight = result.Weight
	c.Height = result.Height
	c.Avatar = result.Avatar
}
