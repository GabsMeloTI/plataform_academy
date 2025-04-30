package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"plataform_init/db"
)

type InterfaceRepository interface {
	CreateUser(ctx context.Context, user db.User) error
	FindUserByEmail(ctx context.Context, email string) (db.User, error) // Função para encontrar o usuário pelo e-mail
}

type Repository struct {
	dbGorm *gorm.DB
}

func NewUserRepository(dbGorm *gorm.DB) *Repository {
	return &Repository{
		dbGorm,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user db.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	result := r.dbGorm.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (db.User, error) {
	var user db.User
	result := r.dbGorm.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, result.Error
	}
	return user, nil
}
